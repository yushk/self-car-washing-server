// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// GetPersonalOKCode is the HTTP code returned for type GetPersonalOK
const GetPersonalOKCode int = 200

/*GetPersonalOK A successful response.

swagger:response getPersonalOK
*/
type GetPersonalOK struct {

	/*
	  In: Body
	*/
	Payload *v1.Personal `json:"body,omitempty"`
}

// NewGetPersonalOK creates GetPersonalOK with default headers values
func NewGetPersonalOK() *GetPersonalOK {

	return &GetPersonalOK{}
}

// WithPayload adds the payload to the get personal o k response
func (o *GetPersonalOK) WithPayload(payload *v1.Personal) *GetPersonalOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get personal o k response
func (o *GetPersonalOK) SetPayload(payload *v1.Personal) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPersonalOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
