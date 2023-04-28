// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OSSToken oss token
//
// swagger:model OSSToken
type OSSToken struct {

	// access key Id
	// Read Only: true
	AccessKeyID string `json:"accessKeyId"`

	// access key secret
	// Read Only: true
	AccessKeySecret string `json:"accessKeySecret"`

	// expiration
	// Read Only: true
	Expiration string `json:"expiration"`

	// sts token
	// Read Only: true
	StsToken string `json:"stsToken"`
}

// Validate validates this o s s token
func (m *OSSToken) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this o s s token based on the context it is used
func (m *OSSToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessKeyID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAccessKeySecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpiration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStsToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OSSToken) contextValidateAccessKeyID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "accessKeyId", "body", string(m.AccessKeyID)); err != nil {
		return err
	}

	return nil
}

func (m *OSSToken) contextValidateAccessKeySecret(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "accessKeySecret", "body", string(m.AccessKeySecret)); err != nil {
		return err
	}

	return nil
}

func (m *OSSToken) contextValidateExpiration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "expiration", "body", string(m.Expiration)); err != nil {
		return err
	}

	return nil
}

func (m *OSSToken) contextValidateStsToken(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "stsToken", "body", string(m.StsToken)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OSSToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OSSToken) UnmarshalBinary(b []byte) error {
	var res OSSToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
