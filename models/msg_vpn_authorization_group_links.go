// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MsgVpnAuthorizationGroupLinks msg vpn authorization group links
// swagger:model MsgVpnAuthorizationGroupLinks
type MsgVpnAuthorizationGroupLinks struct {

	// The URI of this MsgVpnAuthorizationGroup object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this msg vpn authorization group links
func (m *MsgVpnAuthorizationGroupLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnAuthorizationGroupLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnAuthorizationGroupLinks) UnmarshalBinary(b []byte) error {
	var res MsgVpnAuthorizationGroupLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
