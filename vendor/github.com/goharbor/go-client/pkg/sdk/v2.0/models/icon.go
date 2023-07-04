// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Icon icon
//
// swagger:model Icon
type Icon struct {

	// The base64 encoded content of the icon
	Content string `json:"content,omitempty"`

	// The content type of the icon
	ContentType string `json:"content-type,omitempty"`
}

// Validate validates this icon
func (m *Icon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this icon based on context it is used
func (m *Icon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Icon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Icon) UnmarshalBinary(b []byte) error {
	var res Icon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}