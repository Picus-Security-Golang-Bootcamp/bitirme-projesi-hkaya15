// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APICart API cart
//
// swagger:model APICart
type APICart struct {

	// cartitems
	Cartitems []*CartItemToResponse `json:"cartitems"`

	// cartlength
	Cartlength int64 `json:"cartlength,omitempty"`

	// complete order
	CompleteOrder bool `json:"completeOrder,omitempty"`

	// total price
	TotalPrice float64 `json:"totalPrice,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`
}

// Validate validates this API cart
func (m *APICart) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCartitems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APICart) validateCartitems(formats strfmt.Registry) error {
	if swag.IsZero(m.Cartitems) { // not required
		return nil
	}

	for i := 0; i < len(m.Cartitems); i++ {
		if swag.IsZero(m.Cartitems[i]) { // not required
			continue
		}

		if m.Cartitems[i] != nil {
			if err := m.Cartitems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cartitems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cartitems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this API cart based on the context it is used
func (m *APICart) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCartitems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APICart) contextValidateCartitems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Cartitems); i++ {

		if m.Cartitems[i] != nil {
			if err := m.Cartitems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cartitems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cartitems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APICart) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APICart) UnmarshalBinary(b []byte) error {
	var res APICart
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}