// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MsgVpnTopicEndpointLinks msg vpn topic endpoint links
//
// swagger:model MsgVpnTopicEndpointLinks
type MsgVpnTopicEndpointLinks struct {

	// The URI of this Topic Endpoint object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this msg vpn topic endpoint links
func (m *MsgVpnTopicEndpointLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this msg vpn topic endpoint links based on context it is used
func (m *MsgVpnTopicEndpointLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnTopicEndpointLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnTopicEndpointLinks) UnmarshalBinary(b []byte) error {
	var res MsgVpnTopicEndpointLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
