// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClientCertAuthorityOcspTLSTrustedCommonName client cert authority ocsp Tls trusted common name
//
// swagger:model ClientCertAuthorityOcspTlsTrustedCommonName
type ClientCertAuthorityOcspTLSTrustedCommonName struct {

	// The name of the Certificate Authority.
	CertAuthorityName string `json:"certAuthorityName,omitempty"`

	// The expected Trusted Common Name of the OCSP responder remote certificate.
	OcspTLSTrustedCommonName string `json:"ocspTlsTrustedCommonName,omitempty"`
}

// Validate validates this client cert authority ocsp Tls trusted common name
func (m *ClientCertAuthorityOcspTLSTrustedCommonName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this client cert authority ocsp Tls trusted common name based on context it is used
func (m *ClientCertAuthorityOcspTLSTrustedCommonName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClientCertAuthorityOcspTLSTrustedCommonName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientCertAuthorityOcspTLSTrustedCommonName) UnmarshalBinary(b []byte) error {
	var res ClientCertAuthorityOcspTLSTrustedCommonName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}