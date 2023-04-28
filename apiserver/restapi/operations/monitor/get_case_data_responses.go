// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// GetCaseDataOKCode is the HTTP code returned for type GetCaseDataOK
const GetCaseDataOKCode int = 200

/*GetCaseDataOK A successful response.

swagger:response getCaseDataOK
*/
type GetCaseDataOK struct {

	/*
	  In: Body
	*/
	Payload *v1.CaseData `json:"body,omitempty"`
}

// NewGetCaseDataOK creates GetCaseDataOK with default headers values
func NewGetCaseDataOK() *GetCaseDataOK {

	return &GetCaseDataOK{}
}

// WithPayload adds the payload to the get case data o k response
func (o *GetCaseDataOK) WithPayload(payload *v1.CaseData) *GetCaseDataOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get case data o k response
func (o *GetCaseDataOK) SetPayload(payload *v1.CaseData) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCaseDataOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}