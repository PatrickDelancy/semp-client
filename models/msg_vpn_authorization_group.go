// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MsgVpnAuthorizationGroup msg vpn authorization group
//
// swagger:model MsgVpnAuthorizationGroup
type MsgVpnAuthorizationGroup struct {

	// The ACL Profile of the LDAP Authorization Group. The default value is `"default"`.
	ACLProfileName string `json:"aclProfileName,omitempty"`

	// The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as '#', '+', ';', '=' as the value of the group name returned from the LDAP server might prepend those characters with '\'. For example a group name called 'test#,lab,com' will be returned from the LDAP server as 'test\#,lab,com'.
	AuthorizationGroupName string `json:"authorizationGroupName,omitempty"`

	// The Client Profile of the LDAP Authorization Group. The default value is `"default"`.
	ClientProfileName string `json:"clientProfileName,omitempty"`

	// Enable or disable the LDAP Authorization Group in the Message VPN. The default value is `false`.
	Enabled bool `json:"enabled,omitempty"`

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// Lower the priority to be less than this group. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default is not applicable.
	OrderAfterAuthorizationGroupName string `json:"orderAfterAuthorizationGroupName,omitempty"`

	// Raise the priority to be greater than this group. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default is not applicable.
	OrderBeforeAuthorizationGroupName string `json:"orderBeforeAuthorizationGroupName,omitempty"`
}

// Validate validates this msg vpn authorization group
func (m *MsgVpnAuthorizationGroup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this msg vpn authorization group based on context it is used
func (m *MsgVpnAuthorizationGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnAuthorizationGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnAuthorizationGroup) UnmarshalBinary(b []byte) error {
	var res MsgVpnAuthorizationGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
