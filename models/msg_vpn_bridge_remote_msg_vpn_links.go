// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MsgVpnBridgeRemoteMsgVpnLinks msg vpn bridge remote msg vpn links
//
// swagger:model MsgVpnBridgeRemoteMsgVpnLinks
type MsgVpnBridgeRemoteMsgVpnLinks struct {

	// The URI of this Remote Message VPN object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this msg vpn bridge remote msg vpn links
func (m *MsgVpnBridgeRemoteMsgVpnLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this msg vpn bridge remote msg vpn links based on context it is used
func (m *MsgVpnBridgeRemoteMsgVpnLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnBridgeRemoteMsgVpnLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnBridgeRemoteMsgVpnLinks) UnmarshalBinary(b []byte) error {
	var res MsgVpnBridgeRemoteMsgVpnLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
