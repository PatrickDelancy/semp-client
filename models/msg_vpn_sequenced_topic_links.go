// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MsgVpnSequencedTopicLinks msg vpn sequenced topic links
//
// swagger:model MsgVpnSequencedTopicLinks
type MsgVpnSequencedTopicLinks struct {

	// The URI of this Sequenced Topic object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this msg vpn sequenced topic links
func (m *MsgVpnSequencedTopicLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this msg vpn sequenced topic links based on context it is used
func (m *MsgVpnSequencedTopicLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnSequencedTopicLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnSequencedTopicLinks) UnmarshalBinary(b []byte) error {
	var res MsgVpnSequencedTopicLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
