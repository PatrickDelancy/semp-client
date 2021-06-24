// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MsgVpnBridgeRemoteSubscriptionLinks msg vpn bridge remote subscription links
//
// swagger:model MsgVpnBridgeRemoteSubscriptionLinks
type MsgVpnBridgeRemoteSubscriptionLinks struct {

	// The URI of this Remote Subscription object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this msg vpn bridge remote subscription links
func (m *MsgVpnBridgeRemoteSubscriptionLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this msg vpn bridge remote subscription links based on context it is used
func (m *MsgVpnBridgeRemoteSubscriptionLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnBridgeRemoteSubscriptionLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnBridgeRemoteSubscriptionLinks) UnmarshalBinary(b []byte) error {
	var res MsgVpnBridgeRemoteSubscriptionLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
