// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// DeleteHomeWorkHandlerFunc turns a function with the right signature into a delete home work handler
type DeleteHomeWorkHandlerFunc func(DeleteHomeWorkParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteHomeWorkHandlerFunc) Handle(params DeleteHomeWorkParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteHomeWorkHandler interface for that can handle valid delete home work params
type DeleteHomeWorkHandler interface {
	Handle(DeleteHomeWorkParams, *v1.Principal) middleware.Responder
}

// NewDeleteHomeWork creates a new http.Handler for the delete home work operation
func NewDeleteHomeWork(ctx *middleware.Context, handler DeleteHomeWorkHandler) *DeleteHomeWork {
	return &DeleteHomeWork{Context: ctx, Handler: handler}
}

/* DeleteHomeWork swagger:route DELETE /v1/homeWork/{id} monitor deleteHomeWork

删除作业信息

删除作业信息

*/
type DeleteHomeWork struct {
	Context *middleware.Context
	Handler DeleteHomeWorkHandler
}

func (o *DeleteHomeWork) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteHomeWorkParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *v1.Principal
	if uprinc != nil {
		principal = uprinc.(*v1.Principal) // this is really a v1.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
