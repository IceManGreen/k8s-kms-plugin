// Code generated by go-swagger; DO NOT EDIT.

package operation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SimplereenrollHandlerFunc turns a function with the right signature into a simplereenroll handler
type SimplereenrollHandlerFunc func(SimplereenrollParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SimplereenrollHandlerFunc) Handle(params SimplereenrollParams) middleware.Responder {
	return fn(params)
}

// SimplereenrollHandler interface for that can handle valid simplereenroll params
type SimplereenrollHandler interface {
	Handle(SimplereenrollParams) middleware.Responder
}

// NewSimplereenroll creates a new http.Handler for the simplereenroll operation
func NewSimplereenroll(ctx *middleware.Context, handler SimplereenrollHandler) *Simplereenroll {
	return &Simplereenroll{Context: ctx, Handler: handler}
}

/*Simplereenroll swagger:route POST /simplereenroll operation simplereenroll

Enrollment of Clients (requires mutual-tls)

*/
type Simplereenroll struct {
	Context *middleware.Context
	Handler SimplereenrollHandler
}

func (o *Simplereenroll) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSimplereenrollParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}