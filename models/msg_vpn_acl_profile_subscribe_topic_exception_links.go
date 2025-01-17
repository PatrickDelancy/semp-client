// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MsgVpnACLProfileSubscribeTopicExceptionLinks msg vpn Acl profile subscribe topic exception links
//
// swagger:model MsgVpnAclProfileSubscribeTopicExceptionLinks
type MsgVpnACLProfileSubscribeTopicExceptionLinks struct {

	// The URI of this Subscribe Topic Exception object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this msg vpn Acl profile subscribe topic exception links
func (m *MsgVpnACLProfileSubscribeTopicExceptionLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this msg vpn Acl profile subscribe topic exception links based on context it is used
func (m *MsgVpnACLProfileSubscribeTopicExceptionLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnACLProfileSubscribeTopicExceptionLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnACLProfileSubscribeTopicExceptionLinks) UnmarshalBinary(b []byte) error {
	var res MsgVpnACLProfileSubscribeTopicExceptionLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
