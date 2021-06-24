// Code generated by go-swagger; DO NOT EDIT.

package all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PatrickDelancy/semp-client/models"
)

// GetCertAuthorityOcspTLSTrustedCommonNameReader is a Reader for the GetCertAuthorityOcspTLSTrustedCommonName structure.
type GetCertAuthorityOcspTLSTrustedCommonNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCertAuthorityOcspTLSTrustedCommonNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCertAuthorityOcspTLSTrustedCommonNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCertAuthorityOcspTLSTrustedCommonNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCertAuthorityOcspTLSTrustedCommonNameOK creates a GetCertAuthorityOcspTLSTrustedCommonNameOK with default headers values
func NewGetCertAuthorityOcspTLSTrustedCommonNameOK() *GetCertAuthorityOcspTLSTrustedCommonNameOK {
	return &GetCertAuthorityOcspTLSTrustedCommonNameOK{}
}

/* GetCertAuthorityOcspTLSTrustedCommonNameOK describes a response with status code 200, with default header values.

The OCSP Responder Trusted Common Name object's attributes, and the request metadata.
*/
type GetCertAuthorityOcspTLSTrustedCommonNameOK struct {
	Payload *models.CertAuthorityOcspTLSTrustedCommonNameResponse
}

func (o *GetCertAuthorityOcspTLSTrustedCommonNameOK) Error() string {
	return fmt.Sprintf("[GET /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName}][%d] getCertAuthorityOcspTlsTrustedCommonNameOK  %+v", 200, o.Payload)
}
func (o *GetCertAuthorityOcspTLSTrustedCommonNameOK) GetPayload() *models.CertAuthorityOcspTLSTrustedCommonNameResponse {
	return o.Payload
}

func (o *GetCertAuthorityOcspTLSTrustedCommonNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertAuthorityOcspTLSTrustedCommonNameResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertAuthorityOcspTLSTrustedCommonNameDefault creates a GetCertAuthorityOcspTLSTrustedCommonNameDefault with default headers values
func NewGetCertAuthorityOcspTLSTrustedCommonNameDefault(code int) *GetCertAuthorityOcspTLSTrustedCommonNameDefault {
	return &GetCertAuthorityOcspTLSTrustedCommonNameDefault{
		_statusCode: code,
	}
}

/* GetCertAuthorityOcspTLSTrustedCommonNameDefault describes a response with status code -1, with default header values.

The error response.
*/
type GetCertAuthorityOcspTLSTrustedCommonNameDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get cert authority ocsp Tls trusted common name default response
func (o *GetCertAuthorityOcspTLSTrustedCommonNameDefault) Code() int {
	return o._statusCode
}

func (o *GetCertAuthorityOcspTLSTrustedCommonNameDefault) Error() string {
	return fmt.Sprintf("[GET /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName}][%d] getCertAuthorityOcspTlsTrustedCommonName default  %+v", o._statusCode, o.Payload)
}
func (o *GetCertAuthorityOcspTLSTrustedCommonNameDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *GetCertAuthorityOcspTLSTrustedCommonNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}