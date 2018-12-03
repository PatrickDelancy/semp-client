// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MsgVpnACLProfileSubscribeException msg vpn Acl profile subscribe exception
// swagger:model MsgVpnAclProfileSubscribeException
type MsgVpnACLProfileSubscribeException struct {

	// The name of the ACL Profile.
	ACLProfileName string `json:"aclProfileName,omitempty"`

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// The name of the Topic for the Exception to the default action taken. May include wildcard characters.
	SubscribeExceptionTopic string `json:"subscribeExceptionTopic,omitempty"`

	// The syntax of the Topic for the Exception to the default action taken. The allowed values and their meaning are:
	//
	// <pre>
	// "smf" - Topic uses SMF syntax.
	// "mqtt" - Topic uses MQTT syntax.
	// </pre>
	//
	// Enum: [smf mqtt]
	TopicSyntax string `json:"topicSyntax,omitempty"`
}

// Validate validates this msg vpn Acl profile subscribe exception
func (m *MsgVpnACLProfileSubscribeException) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTopicSyntax(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var msgVpnAclProfileSubscribeExceptionTypeTopicSyntaxPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["smf","mqtt"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnAclProfileSubscribeExceptionTypeTopicSyntaxPropEnum = append(msgVpnAclProfileSubscribeExceptionTypeTopicSyntaxPropEnum, v)
	}
}

const (

	// MsgVpnACLProfileSubscribeExceptionTopicSyntaxSmf captures enum value "smf"
	MsgVpnACLProfileSubscribeExceptionTopicSyntaxSmf string = "smf"

	// MsgVpnACLProfileSubscribeExceptionTopicSyntaxMqtt captures enum value "mqtt"
	MsgVpnACLProfileSubscribeExceptionTopicSyntaxMqtt string = "mqtt"
)

// prop value enum
func (m *MsgVpnACLProfileSubscribeException) validateTopicSyntaxEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, msgVpnAclProfileSubscribeExceptionTypeTopicSyntaxPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnACLProfileSubscribeException) validateTopicSyntax(formats strfmt.Registry) error {

	if swag.IsZero(m.TopicSyntax) { // not required
		return nil
	}

	// value enum
	if err := m.validateTopicSyntaxEnum("topicSyntax", "body", m.TopicSyntax); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnACLProfileSubscribeException) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnACLProfileSubscribeException) UnmarshalBinary(b []byte) error {
	var res MsgVpnACLProfileSubscribeException
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
