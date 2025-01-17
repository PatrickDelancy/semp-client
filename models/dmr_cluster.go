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

// DmrCluster dmr cluster
//
// swagger:model DmrCluster
type DmrCluster struct {

	// Enable or disable basic authentication for Cluster Links. The default value is `true`.
	AuthenticationBasicEnabled bool `json:"authenticationBasicEnabled,omitempty"`

	// The password used to authenticate incoming Cluster Links when using basic internal authentication. The same password is also used by outgoing Cluster Links if a per-Link password is not configured. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default value is `""`.
	AuthenticationBasicPassword string `json:"authenticationBasicPassword,omitempty"`

	// The type of basic authentication to use for Cluster Links. The default value is `"internal"`. The allowed values and their meaning are:
	//
	// <pre>
	// "internal" - Use locally configured password.
	// "none" - No authentication.
	// </pre>
	//
	// Enum: [internal none]
	AuthenticationBasicType string `json:"authenticationBasicType,omitempty"`

	// The PEM formatted content for the client certificate used to login to the remote node. It must consist of a private key and between one and three certificates comprising the certificate trust chain. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changing this attribute requires an HTTPS connection. The default value is `""`.
	AuthenticationClientCertContent string `json:"authenticationClientCertContent,omitempty"`

	// Enable or disable client certificate authentication for Cluster Links. The default value is `true`.
	AuthenticationClientCertEnabled bool `json:"authenticationClientCertEnabled,omitempty"`

	// The password for the client certificate. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changing this attribute requires an HTTPS connection. The default value is `""`.
	AuthenticationClientCertPassword string `json:"authenticationClientCertPassword,omitempty"`

	// Enable or disable direct messaging only. Guaranteed messages will not be transmitted through the cluster. The default value is `false`.
	DirectOnlyEnabled bool `json:"directOnlyEnabled,omitempty"`

	// The name of the Cluster.
	DmrClusterName string `json:"dmrClusterName,omitempty"`

	// Enable or disable the Cluster. The default value is `false`.
	Enabled bool `json:"enabled,omitempty"`

	// The name of this node in the Cluster. This is the name that this broker (or redundant group of brokers) is know by to other nodes in the Cluster. The name is chosen automatically to be either this broker's Router Name or Mate Router Name, depending on which Active Standby Role (primary or backup) this broker plays in its redundancy group.
	NodeName string `json:"nodeName,omitempty"`

	// Enable or disable the enforcing of the common name provided by the remote broker against the list of trusted common names configured for the Link. If enabled, the certificate's common name must match one of the trusted common names for the Link to be accepted. Common Name validation is not performed if Server Certificate Name Validation is enabled, even if Common Name validation is enabled. The default value is `true`. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.
	TLSServerCertEnforceTrustedCommonNameEnabled bool `json:"tlsServerCertEnforceTrustedCommonNameEnabled,omitempty"`

	// The maximum allowed depth of a certificate chain. The depth of a chain is defined as the number of signing CA certificates that are present in the chain back to a trusted self-signed root CA certificate. The default value is `3`.
	TLSServerCertMaxChainDepth int64 `json:"tlsServerCertMaxChainDepth,omitempty"`

	// Enable or disable the validation of the "Not Before" and "Not After" validity dates in the certificate. When disabled, the certificate is accepted even if the certificate is not valid based on these dates. The default value is `true`.
	TLSServerCertValidateDateEnabled bool `json:"tlsServerCertValidateDateEnabled,omitempty"`

	// Enable or disable the standard TLS authentication mechanism of verifying the name used to connect to the bridge. If enabled, the name used to connect to the bridge is checked against the names specified in the certificate returned by the remote router. Legacy Common Name validation is not performed if Server Certificate Name Validation is enabled, even if Common Name validation is also enabled. The default value is `true`. Available since 2.18.
	TLSServerCertValidateNameEnabled bool `json:"tlsServerCertValidateNameEnabled,omitempty"`
}

// Validate validates this dmr cluster
func (m *DmrCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticationBasicType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var dmrClusterTypeAuthenticationBasicTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["internal","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dmrClusterTypeAuthenticationBasicTypePropEnum = append(dmrClusterTypeAuthenticationBasicTypePropEnum, v)
	}
}

const (

	// DmrClusterAuthenticationBasicTypeInternal captures enum value "internal"
	DmrClusterAuthenticationBasicTypeInternal string = "internal"

	// DmrClusterAuthenticationBasicTypeNone captures enum value "none"
	DmrClusterAuthenticationBasicTypeNone string = "none"
)

// prop value enum
func (m *DmrCluster) validateAuthenticationBasicTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dmrClusterTypeAuthenticationBasicTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DmrCluster) validateAuthenticationBasicType(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticationBasicType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticationBasicTypeEnum("authenticationBasicType", "body", m.AuthenticationBasicType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this dmr cluster based on context it is used
func (m *DmrCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DmrCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DmrCluster) UnmarshalBinary(b []byte) error {
	var res DmrCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
