// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SoleToken sole token
//
// swagger:model SoleToken
type SoleToken struct {

	// code
	Code int32 `json:"code,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this sole token
func (m *SoleToken) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sole token based on context it is used
func (m *SoleToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SoleToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoleToken) UnmarshalBinary(b []byte) error {
	var res SoleToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
