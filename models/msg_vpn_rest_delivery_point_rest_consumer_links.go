// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MsgVpnRestDeliveryPointRestConsumerLinks msg vpn rest delivery point rest consumer links
//
// swagger:model MsgVpnRestDeliveryPointRestConsumerLinks
type MsgVpnRestDeliveryPointRestConsumerLinks struct {

	// The URI of this REST Consumer's collection of Claim objects. Available since 2.21.
	OauthJwtClaimsURI string `json:"oauthJwtClaimsUri,omitempty"`

	// The URI of this REST Consumer's collection of Trusted Common Name objects. Deprecated since (will be deprecated in next SEMP version). Common Name validation has been replaced by Server Certificate Name validation.
	TLSTrustedCommonNamesURI string `json:"tlsTrustedCommonNamesUri,omitempty"`

	// The URI of this REST Consumer object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this msg vpn rest delivery point rest consumer links
func (m *MsgVpnRestDeliveryPointRestConsumerLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this msg vpn rest delivery point rest consumer links based on context it is used
func (m *MsgVpnRestDeliveryPointRestConsumerLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnRestDeliveryPointRestConsumerLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnRestDeliveryPointRestConsumerLinks) UnmarshalBinary(b []byte) error {
	var res MsgVpnRestDeliveryPointRestConsumerLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
