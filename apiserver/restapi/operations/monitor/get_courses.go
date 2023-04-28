// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// GetCoursesHandlerFunc turns a function with the right signature into a get courses handler
type GetCoursesHandlerFunc func(GetCoursesParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCoursesHandlerFunc) Handle(params GetCoursesParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetCoursesHandler interface for that can handle valid get courses params
type GetCoursesHandler interface {
	Handle(GetCoursesParams, *v1.Principal) middleware.Responder
}

// NewGetCourses creates a new http.Handler for the get courses operation
func NewGetCourses(ctx *middleware.Context, handler GetCoursesHandler) *GetCourses {
	return &GetCourses{Context: ctx, Handler: handler}
}

/* GetCourses swagger:route GET /v1/courses monitor getCourses

获取课程信息列表

获取课程信息列表

*/
type GetCourses struct {
	Context *middleware.Context
	Handler GetCoursesHandler
}

func (o *GetCourses) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetCoursesParams()
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
