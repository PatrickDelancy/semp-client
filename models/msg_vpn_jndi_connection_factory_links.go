// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MsgVpnJndiConnectionFactoryLinks msg vpn jndi connection factory links
//
// swagger:model MsgVpnJndiConnectionFactoryLinks
type MsgVpnJndiConnectionFactoryLinks struct {

	// The URI of this JNDI Connection Factory object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this msg vpn jndi connection factory links
func (m *MsgVpnJndiConnectionFactoryLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this msg vpn jndi connection factory links based on context it is used
func (m *MsgVpnJndiConnectionFactoryLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnJndiConnectionFactoryLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnJndiConnectionFactoryLinks) UnmarshalBinary(b []byte) error {
	var res MsgVpnJndiConnectionFactoryLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
