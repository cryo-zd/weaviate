package db

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/weaviate/weaviate/adapters/repos/db/vector/common"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/entities/vectorindex/hnsw"
)

func TestVectorIndexQueueBatchSize(t *testing.T) {
	t.Setenv("ASYNC_INDEXING", "true")
	os.Setenv("ASYNC_INDEXING_BATCH_SIZE", "6000")
	os.Setenv("ASYNC_INDEXING_STALE_TIMEOUT", "1ms")

	ctx := context.Background()
	className := "TestClass"
	shd, _ := testShardWithSettings(t, ctx, &models.Class{Class: className}, hnsw.UserConfig{}, false, true)

	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Println(err)
		}
	}(shd.Index().Config.RootPath)

	count := 10_000

	v := make([]float32, 1000)

	var vectors []common.VectorRecord
	for i := range count {
		vectors = append(vectors, &common.Vector[[]float32]{
			ID:     uint64(i),
			Vector: v,
		})
	}

	q := shd.Queue()
	err := q.Insert(ctx, vectors...)
	require.NoError(t, err)

	// wait for the batch to be stale
	time.Sleep(100 * time.Millisecond)

	b, err := q.DequeueBatch()
	require.NoError(t, err)
	require.NotNil(t, b)
	require.InEpsilon(t, 6000, len(b.Tasks), 0.4)
	size := len(b.Tasks)
	b.Done()
	require.EqualValues(t, 10000-size, q.Size())
}
