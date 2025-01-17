// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AboutAPILinks about Api links
//
// swagger:model AboutApiLinks
type AboutAPILinks struct {

	// The URI of this API Description object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this about Api links
func (m *AboutAPILinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this about Api links based on context it is used
func (m *AboutAPILinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AboutAPILinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AboutAPILinks) UnmarshalBinary(b []byte) error {
	var res AboutAPILinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
