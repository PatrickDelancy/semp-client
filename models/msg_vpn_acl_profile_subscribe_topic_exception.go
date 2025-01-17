// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MsgVpnACLProfileSubscribeTopicException msg vpn Acl profile subscribe topic exception
//
// swagger:model MsgVpnAclProfileSubscribeTopicException
type MsgVpnACLProfileSubscribeTopicException struct {

	// The name of the ACL Profile.
	ACLProfileName string `json:"aclProfileName,omitempty"`

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// The topic for the exception to the default action taken. May include wildcard characters.
	SubscribeTopicException string `json:"subscribeTopicException,omitempty"`

	// The syntax of the topic for the exception to the default action taken. The allowed values and their meaning are:
	//
	// <pre>
	// "smf" - Topic uses SMF syntax.
	// "mqtt" - Topic uses MQTT syntax.
	// </pre>
	//
	// Enum: [smf mqtt]
	SubscribeTopicExceptionSyntax string `json:"subscribeTopicExceptionSyntax,omitempty"`
}

// Validate validates this msg vpn Acl profile subscribe topic exception
func (m *MsgVpnACLProfileSubscribeTopicException) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubscribeTopicExceptionSyntax(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var msgVpnAclProfileSubscribeTopicExceptionTypeSubscribeTopicExceptionSyntaxPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["smf","mqtt"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnAclProfileSubscribeTopicExceptionTypeSubscribeTopicExceptionSyntaxPropEnum = append(msgVpnAclProfileSubscribeTopicExceptionTypeSubscribeTopicExceptionSyntaxPropEnum, v)
	}
}

const (

	// MsgVpnACLProfileSubscribeTopicExceptionSubscribeTopicExceptionSyntaxSmf captures enum value "smf"
	MsgVpnACLProfileSubscribeTopicExceptionSubscribeTopicExceptionSyntaxSmf string = "smf"

	// MsgVpnACLProfileSubscribeTopicExceptionSubscribeTopicExceptionSyntaxMqtt captures enum value "mqtt"
	MsgVpnACLProfileSubscribeTopicExceptionSubscribeTopicExceptionSyntaxMqtt string = "mqtt"
)

// prop value enum
func (m *MsgVpnACLProfileSubscribeTopicException) validateSubscribeTopicExceptionSyntaxEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, msgVpnAclProfileSubscribeTopicExceptionTypeSubscribeTopicExceptionSyntaxPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnACLProfileSubscribeTopicException) validateSubscribeTopicExceptionSyntax(formats strfmt.Registry) error {
	if swag.IsZero(m.SubscribeTopicExceptionSyntax) { // not required
		return nil
	}

	// value enum
	if err := m.validateSubscribeTopicExceptionSyntaxEnum("subscribeTopicExceptionSyntax", "body", m.SubscribeTopicExceptionSyntax); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this msg vpn Acl profile subscribe topic exception based on context it is used
func (m *MsgVpnACLProfileSubscribeTopicException) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnACLProfileSubscribeTopicException) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnACLProfileSubscribeTopicException) UnmarshalBinary(b []byte) error {
	var res MsgVpnACLProfileSubscribeTopicException
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
