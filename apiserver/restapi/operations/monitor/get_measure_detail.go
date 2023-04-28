// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// GetMeasureDetailHandlerFunc turns a function with the right signature into a get measure detail handler
type GetMeasureDetailHandlerFunc func(GetMeasureDetailParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMeasureDetailHandlerFunc) Handle(params GetMeasureDetailParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetMeasureDetailHandler interface for that can handle valid get measure detail params
type GetMeasureDetailHandler interface {
	Handle(GetMeasureDetailParams, *v1.Principal) middleware.Responder
}

// NewGetMeasureDetail creates a new http.Handler for the get measure detail operation
func NewGetMeasureDetail(ctx *middleware.Context, handler GetMeasureDetailHandler) *GetMeasureDetail {
	return &GetMeasureDetail{Context: ctx, Handler: handler}
}

/* GetMeasureDetail swagger:route GET /v1/measureDetail/{id} monitor getMeasureDetail

获取测评数据明细信息

获取测评数据明细信息

*/
type GetMeasureDetail struct {
	Context *middleware.Context
	Handler GetMeasureDetailHandler
}

func (o *GetMeasureDetail) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetMeasureDetailParams()
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
