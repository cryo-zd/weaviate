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

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/weaviate/weaviate/entities/models"
)

// TenantsCreateHandlerFunc turns a function with the right signature into a tenants create handler
type TenantsCreateHandlerFunc func(TenantsCreateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn TenantsCreateHandlerFunc) Handle(params TenantsCreateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// TenantsCreateHandler interface for that can handle valid tenants create params
type TenantsCreateHandler interface {
	Handle(TenantsCreateParams, *models.Principal) middleware.Responder
}

// NewTenantsCreate creates a new http.Handler for the tenants create operation
func NewTenantsCreate(ctx *middleware.Context, handler TenantsCreateHandler) *TenantsCreate {
	return &TenantsCreate{Context: ctx, Handler: handler}
}

/*
	TenantsCreate swagger:route POST /schema/{className}/tenants schema tenantsCreate

# Create a new tenant

Create a new tenant for a collection. Multi-tenancy must be enabled in the collection definition.
*/
type TenantsCreate struct {
	Context *middleware.Context
	Handler TenantsCreateHandler
}

func (o *TenantsCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewTenantsCreateParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
