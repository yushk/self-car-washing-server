// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// GetPhoneNumberHandlerFunc turns a function with the right signature into a get phone number handler
type GetPhoneNumberHandlerFunc func(GetPhoneNumberParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPhoneNumberHandlerFunc) Handle(params GetPhoneNumberParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetPhoneNumberHandler interface for that can handle valid get phone number params
type GetPhoneNumberHandler interface {
	Handle(GetPhoneNumberParams, *v1.Principal) middleware.Responder
}

// NewGetPhoneNumber creates a new http.Handler for the get phone number operation
func NewGetPhoneNumber(ctx *middleware.Context, handler GetPhoneNumberHandler) *GetPhoneNumber {
	return &GetPhoneNumber{Context: ctx, Handler: handler}
}

/* GetPhoneNumber swagger:route GET /v1/wx/phone user getPhoneNumber

获取手机号

获取手机号

*/
type GetPhoneNumber struct {
	Context *middleware.Context
	Handler GetPhoneNumberHandler
}

func (o *GetPhoneNumber) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPhoneNumberParams()
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