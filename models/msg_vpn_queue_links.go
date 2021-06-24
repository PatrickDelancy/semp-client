// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MsgVpnQueueLinks msg vpn queue links
//
// swagger:model MsgVpnQueueLinks
type MsgVpnQueueLinks struct {

	// The URI of this Queue's collection of Queue Subscription objects.
	SubscriptionsURI string `json:"subscriptionsUri,omitempty"`

	// The URI of this Queue object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this msg vpn queue links
func (m *MsgVpnQueueLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this msg vpn queue links based on context it is used
func (m *MsgVpnQueueLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnQueueLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnQueueLinks) UnmarshalBinary(b []byte) error {
	var res MsgVpnQueueLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
