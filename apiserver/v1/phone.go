// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Phone phone
//
// swagger:model Phone
type Phone struct {

	// jwt token
	Phone string `json:"phone"`
}

// Validate validates this phone
func (m *Phone) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this phone based on context it is used
func (m *Phone) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Phone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Phone) UnmarshalBinary(b []byte) error {
	var res Phone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
