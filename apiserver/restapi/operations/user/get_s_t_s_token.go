// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSTSTokenHandlerFunc turns a function with the right signature into a get s t s token handler
type GetSTSTokenHandlerFunc func(GetSTSTokenParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSTSTokenHandlerFunc) Handle(params GetSTSTokenParams) middleware.Responder {
	return fn(params)
}

// GetSTSTokenHandler interface for that can handle valid get s t s token params
type GetSTSTokenHandler interface {
	Handle(GetSTSTokenParams) middleware.Responder
}

// NewGetSTSToken creates a new http.Handler for the get s t s token operation
func NewGetSTSToken(ctx *middleware.Context, handler GetSTSTokenHandler) *GetSTSToken {
	return &GetSTSToken{Context: ctx, Handler: handler}
}

/* GetSTSToken swagger:route GET /v1/osssts user getSTSToken

oss ststoken

获取oss 临时token

*/
type GetSTSToken struct {
	Context *middleware.Context
	Handler GetSTSTokenHandler
}

func (o *GetSTSToken) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSTSTokenParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
