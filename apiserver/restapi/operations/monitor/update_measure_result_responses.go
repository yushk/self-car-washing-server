// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// UpdateMeasureResultOKCode is the HTTP code returned for type UpdateMeasureResultOK
const UpdateMeasureResultOKCode int = 200

/*UpdateMeasureResultOK A successful response.

swagger:response updateMeasureResultOK
*/
type UpdateMeasureResultOK struct {

	/*
	  In: Body
	*/
	Payload *v1.MeasureResult `json:"body,omitempty"`
}

// NewUpdateMeasureResultOK creates UpdateMeasureResultOK with default headers values
func NewUpdateMeasureResultOK() *UpdateMeasureResultOK {

	return &UpdateMeasureResultOK{}
}

// WithPayload adds the payload to the update measure result o k response
func (o *UpdateMeasureResultOK) WithPayload(payload *v1.MeasureResult) *UpdateMeasureResultOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update measure result o k response
func (o *UpdateMeasureResultOK) SetPayload(payload *v1.MeasureResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMeasureResultOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
