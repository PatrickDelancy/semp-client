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

// MsgVpnBridgeTLSTrustedCommonName msg vpn bridge Tls trusted common name
// swagger:model MsgVpnBridgeTlsTrustedCommonName
type MsgVpnBridgeTLSTrustedCommonName struct {

	// The name of the Bridge.
	BridgeName string `json:"bridgeName,omitempty"`

	// Specify whether the Bridge is configured for the primary or backup Virtual Router or auto configured. The allowed values and their meaning are:
	//
	// <pre>
	// "primary" - The Bridge is used for the primary Virtual Router.
	// "backup" - The Bridge is used for the backup Virtual Router.
	// "auto" - The Bridge is automatically assigned a Router.
	// </pre>
	//
	// Enum: [primary backup auto]
	BridgeVirtualRouter string `json:"bridgeVirtualRouter,omitempty"`

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// The expected trusted common name of the remote certificate.
	TLSTrustedCommonName string `json:"tlsTrustedCommonName,omitempty"`
}

// Validate validates this msg vpn bridge Tls trusted common name
func (m *MsgVpnBridgeTLSTrustedCommonName) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBridgeVirtualRouter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var msgVpnBridgeTlsTrustedCommonNameTypeBridgeVirtualRouterPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["primary","backup","auto"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnBridgeTlsTrustedCommonNameTypeBridgeVirtualRouterPropEnum = append(msgVpnBridgeTlsTrustedCommonNameTypeBridgeVirtualRouterPropEnum, v)
	}
}

const (

	// MsgVpnBridgeTLSTrustedCommonNameBridgeVirtualRouterPrimary captures enum value "primary"
	MsgVpnBridgeTLSTrustedCommonNameBridgeVirtualRouterPrimary string = "primary"

	// MsgVpnBridgeTLSTrustedCommonNameBridgeVirtualRouterBackup captures enum value "backup"
	MsgVpnBridgeTLSTrustedCommonNameBridgeVirtualRouterBackup string = "backup"

	// MsgVpnBridgeTLSTrustedCommonNameBridgeVirtualRouterAuto captures enum value "auto"
	MsgVpnBridgeTLSTrustedCommonNameBridgeVirtualRouterAuto string = "auto"
)

// prop value enum
func (m *MsgVpnBridgeTLSTrustedCommonName) validateBridgeVirtualRouterEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, msgVpnBridgeTlsTrustedCommonNameTypeBridgeVirtualRouterPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnBridgeTLSTrustedCommonName) validateBridgeVirtualRouter(formats strfmt.Registry) error {

	if swag.IsZero(m.BridgeVirtualRouter) { // not required
		return nil
	}

	// value enum
	if err := m.validateBridgeVirtualRouterEnum("bridgeVirtualRouter", "body", m.BridgeVirtualRouter); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnBridgeTLSTrustedCommonName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnBridgeTLSTrustedCommonName) UnmarshalBinary(b []byte) error {
	var res MsgVpnBridgeTLSTrustedCommonName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}