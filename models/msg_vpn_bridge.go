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

// MsgVpnBridge msg vpn bridge
//
// swagger:model MsgVpnBridge
type MsgVpnBridge struct {

	// The name of the Bridge.
	BridgeName string `json:"bridgeName,omitempty"`

	// The virtual router of the Bridge. The allowed values and their meaning are:
	//
	// <pre>
	// "primary" - The Bridge is used for the primary virtual router.
	// "backup" - The Bridge is used for the backup virtual router.
	// "auto" - The Bridge is automatically assigned a virtual router at creation, depending on the broker's active-standby role.
	// </pre>
	//
	// Enum: [primary backup auto]
	BridgeVirtualRouter string `json:"bridgeVirtualRouter,omitempty"`

	// Enable or disable the Bridge. The default value is `false`.
	Enabled bool `json:"enabled,omitempty"`

	// The maximum time-to-live (TTL) in hops. Messages are discarded if their TTL exceeds this value. The default value is `8`.
	MaxTTL int64 `json:"maxTtl,omitempty"`

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// The Client Username the Bridge uses to login to the remote Message VPN. The default value is `""`.
	RemoteAuthenticationBasicClientUsername string `json:"remoteAuthenticationBasicClientUsername,omitempty"`

	// The password for the Client Username. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default value is `""`.
	RemoteAuthenticationBasicPassword string `json:"remoteAuthenticationBasicPassword,omitempty"`

	// The PEM formatted content for the client certificate used by the Bridge to login to the remote Message VPN. It must consist of a private key and between one and three certificates comprising the certificate trust chain. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changing this attribute requires an HTTPS connection. The default value is `""`. Available since 2.9.
	RemoteAuthenticationClientCertContent string `json:"remoteAuthenticationClientCertContent,omitempty"`

	// The password for the client certificate. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changing this attribute requires an HTTPS connection. The default value is `""`. Available since 2.9.
	RemoteAuthenticationClientCertPassword string `json:"remoteAuthenticationClientCertPassword,omitempty"`

	// The authentication scheme for the remote Message VPN. The default value is `"basic"`. The allowed values and their meaning are:
	//
	// <pre>
	// "basic" - Basic Authentication Scheme (via username and password).
	// "client-certificate" - Client Certificate Authentication Scheme (via certificate file or content).
	// </pre>
	//
	// Enum: [basic client-certificate]
	RemoteAuthenticationScheme string `json:"remoteAuthenticationScheme,omitempty"`

	// The maximum number of retry attempts to establish a connection to the remote Message VPN. A value of 0 means to retry forever. The default value is `0`.
	RemoteConnectionRetryCount int64 `json:"remoteConnectionRetryCount,omitempty"`

	// The number of seconds the broker waits for the bridge connection to be established before attempting a new connection. The default value is `3`.
	RemoteConnectionRetryDelay int64 `json:"remoteConnectionRetryDelay,omitempty"`

	// The priority for deliver-to-one (DTO) messages transmitted from the remote Message VPN. The default value is `"p1"`. The allowed values and their meaning are:
	//
	// <pre>
	// "p1" - The 1st or highest priority.
	// "p2" - The 2nd highest priority.
	// "p3" - The 3rd highest priority.
	// "p4" - The 4th highest priority.
	// "da" - Ignore priority and deliver always.
	// </pre>
	//
	// Enum: [p1 p2 p3 p4 da]
	RemoteDeliverToOnePriority string `json:"remoteDeliverToOnePriority,omitempty"`

	// The colon-separated list of cipher suites supported for TLS connections to the remote Message VPN. The value "default" implies all supported suites ordered from most secure to least secure. The default value is `"default"`.
	TLSCipherSuiteList string `json:"tlsCipherSuiteList,omitempty"`
}

// Validate validates this msg vpn bridge
func (m *MsgVpnBridge) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBridgeVirtualRouter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteAuthenticationScheme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteDeliverToOnePriority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var msgVpnBridgeTypeBridgeVirtualRouterPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["primary","backup","auto"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnBridgeTypeBridgeVirtualRouterPropEnum = append(msgVpnBridgeTypeBridgeVirtualRouterPropEnum, v)
	}
}

const (

	// MsgVpnBridgeBridgeVirtualRouterPrimary captures enum value "primary"
	MsgVpnBridgeBridgeVirtualRouterPrimary string = "primary"

	// MsgVpnBridgeBridgeVirtualRouterBackup captures enum value "backup"
	MsgVpnBridgeBridgeVirtualRouterBackup string = "backup"

	// MsgVpnBridgeBridgeVirtualRouterAuto captures enum value "auto"
	MsgVpnBridgeBridgeVirtualRouterAuto string = "auto"
)

// prop value enum
func (m *MsgVpnBridge) validateBridgeVirtualRouterEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, msgVpnBridgeTypeBridgeVirtualRouterPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnBridge) validateBridgeVirtualRouter(formats strfmt.Registry) error {
	if swag.IsZero(m.BridgeVirtualRouter) { // not required
		return nil
	}

	// value enum
	if err := m.validateBridgeVirtualRouterEnum("bridgeVirtualRouter", "body", m.BridgeVirtualRouter); err != nil {
		return err
	}

	return nil
}

var msgVpnBridgeTypeRemoteAuthenticationSchemePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["basic","client-certificate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnBridgeTypeRemoteAuthenticationSchemePropEnum = append(msgVpnBridgeTypeRemoteAuthenticationSchemePropEnum, v)
	}
}

const (

	// MsgVpnBridgeRemoteAuthenticationSchemeBasic captures enum value "basic"
	MsgVpnBridgeRemoteAuthenticationSchemeBasic string = "basic"

	// MsgVpnBridgeRemoteAuthenticationSchemeClientDashCertificate captures enum value "client-certificate"
	MsgVpnBridgeRemoteAuthenticationSchemeClientDashCertificate string = "client-certificate"
)

// prop value enum
func (m *MsgVpnBridge) validateRemoteAuthenticationSchemeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, msgVpnBridgeTypeRemoteAuthenticationSchemePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnBridge) validateRemoteAuthenticationScheme(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteAuthenticationScheme) { // not required
		return nil
	}

	// value enum
	if err := m.validateRemoteAuthenticationSchemeEnum("remoteAuthenticationScheme", "body", m.RemoteAuthenticationScheme); err != nil {
		return err
	}

	return nil
}

var msgVpnBridgeTypeRemoteDeliverToOnePriorityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["p1","p2","p3","p4","da"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnBridgeTypeRemoteDeliverToOnePriorityPropEnum = append(msgVpnBridgeTypeRemoteDeliverToOnePriorityPropEnum, v)
	}
}

const (

	// MsgVpnBridgeRemoteDeliverToOnePriorityP1 captures enum value "p1"
	MsgVpnBridgeRemoteDeliverToOnePriorityP1 string = "p1"

	// MsgVpnBridgeRemoteDeliverToOnePriorityP2 captures enum value "p2"
	MsgVpnBridgeRemoteDeliverToOnePriorityP2 string = "p2"

	// MsgVpnBridgeRemoteDeliverToOnePriorityP3 captures enum value "p3"
	MsgVpnBridgeRemoteDeliverToOnePriorityP3 string = "p3"

	// MsgVpnBridgeRemoteDeliverToOnePriorityP4 captures enum value "p4"
	MsgVpnBridgeRemoteDeliverToOnePriorityP4 string = "p4"

	// MsgVpnBridgeRemoteDeliverToOnePriorityDa captures enum value "da"
	MsgVpnBridgeRemoteDeliverToOnePriorityDa string = "da"
)

// prop value enum
func (m *MsgVpnBridge) validateRemoteDeliverToOnePriorityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, msgVpnBridgeTypeRemoteDeliverToOnePriorityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnBridge) validateRemoteDeliverToOnePriority(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteDeliverToOnePriority) { // not required
		return nil
	}

	// value enum
	if err := m.validateRemoteDeliverToOnePriorityEnum("remoteDeliverToOnePriority", "body", m.RemoteDeliverToOnePriority); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this msg vpn bridge based on context it is used
func (m *MsgVpnBridge) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnBridge) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnBridge) UnmarshalBinary(b []byte) error {
	var res MsgVpnBridge
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
