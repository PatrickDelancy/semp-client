// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimLinks msg vpn rest delivery point rest consumer oauth jwt claim links
//
// swagger:model MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimLinks
type MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimLinks struct {

	// The URI of this Claim object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this msg vpn rest delivery point rest consumer oauth jwt claim links
func (m *MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this msg vpn rest delivery point rest consumer oauth jwt claim links based on context it is used
func (m *MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimLinks) UnmarshalBinary(b []byte) error {
	var res MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
