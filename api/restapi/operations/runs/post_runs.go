// Code generated by go-swagger; DO NOT EDIT.

package runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostRunsHandlerFunc turns a function with the right signature into a post runs handler
type PostRunsHandlerFunc func(PostRunsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostRunsHandlerFunc) Handle(params PostRunsParams) middleware.Responder {
	return fn(params)
}

// PostRunsHandler interface for that can handle valid post runs params
type PostRunsHandler interface {
	Handle(PostRunsParams) middleware.Responder
}

// NewPostRuns creates a new http.Handler for the post runs operation
func NewPostRuns(ctx *middleware.Context, handler PostRunsHandler) *PostRuns {
	return &PostRuns{Context: ctx, Handler: handler}
}

/*
	PostRuns swagger:route POST /runs runs postRuns

Register an experiment run for monitoring.

Register an experiment run for monitoring
*/
type PostRuns struct {
	Context *middleware.Context
	Handler PostRunsHandler
}

func (o *PostRuns) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostRunsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}