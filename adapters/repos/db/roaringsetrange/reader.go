//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package roaringsetrange

import (
	"context"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/weaviate/sroar"
	"github.com/weaviate/weaviate/adapters/repos/db/helpers"
	"github.com/weaviate/weaviate/adapters/repos/db/roaringset"
	"github.com/weaviate/weaviate/entities/concurrency"
	"github.com/weaviate/weaviate/entities/errors"
	"github.com/weaviate/weaviate/entities/filters"
)

type InnerReader interface {
	Read(ctx context.Context, value uint64, operator filters.Operator) (roaringset.BitmapLayer, error)
}

type CombinedReader struct {
	logger         logrus.FieldLogger
	readers        []InnerReader
	releaseReaders func()
	concurrency    int
}

func NewCombinedReader(readers []InnerReader, releaseReaders func(), concurrency int,
	logger logrus.FieldLogger,
) *CombinedReader {
	return &CombinedReader{
		logger:         logger,
		readers:        readers,
		releaseReaders: releaseReaders,
		concurrency:    concurrency,
	}
}

func (r *CombinedReader) Read(ctx context.Context, value uint64, operator filters.Operator,
) (*sroar.Bitmap, error) {
	before := time.Now()
	count := len(r.readers)

	var subresultsReadSum time.Duration
	var mergingSum time.Duration
	lock := new(sync.Mutex)

	defer func() {
		took := time.Since(before)
		vals := map[string]any{
			"readers":                         count,
			"subresults_read_sum_took":        subresultsReadSum,
			"subresults_read_sum_took_string": subresultsReadSum.String(),
			"merging_sum_took":                mergingSum,
			"merging_sum_took_string":         mergingSum.String(),
			"took":                            took,
			"took_string":                     took.String(),
		}

		helpers.AnnotateSlowQueryLogAppend(ctx, "build_allow_list_doc_bitmap_roaringrange", vals)
	}()

	switch count {
	case 0:
		return sroar.NewBitmap(), nil
	case 1:
		t := time.Now()
		layer, err := r.readers[0].Read(ctx, value, operator)
		subresultsReadSum = time.Since(t)

		if err != nil {
			return nil, err
		}
		return layer.Additions, nil
	}

	// all readers but last one. it will be processed by current goroutine
	responseChans := make([]chan *readerResponse, count-1)
	for i := range responseChans {
		responseChans[i] = make(chan *readerResponse, 1)
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	errors.GoWrapper(func() {
		eg, gctx := errors.NewErrorGroupWithContextWrapper(r.logger, ctx)
		eg.SetLimit(r.concurrency)

		for i := 1; i < count; i++ {
			i := i
			eg.Go(func() error {
				t := time.Now()
				layer, err := r.readers[i].Read(gctx, value, operator)
				lock.Lock()
				subresultsReadSum += time.Since(t)
				lock.Unlock()
				responseChans[i-1] <- &readerResponse{layer, err}
				return err
			})
		}
	}, r.logger)

	t := time.Now()
	layer, err := r.readers[0].Read(ctx, value, operator)
	lock.Lock()
	subresultsReadSum += time.Since(t)
	lock.Unlock()
	if err != nil {
		return nil, err
	}

	for i := 1; i < count; i++ {
		response := <-responseChans[i-1]
		if response.err != nil {
			return nil, response.err
		}

		t := time.Now()
		layer.Additions.AndNotConc(response.layer.Deletions, concurrency.SROAR_MERGE)
		layer.Additions.OrConc(response.layer.Additions, concurrency.SROAR_MERGE)
		mergingSum += time.Since(t)
	}

	return layer.Additions, nil
}

func (r *CombinedReader) Close() {
	r.releaseReaders()
}

type readerResponse struct {
	layer roaringset.BitmapLayer
	err   error
}
