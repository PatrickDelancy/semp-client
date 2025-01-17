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

// DeleteCertAuthorityOcspTLSTrustedCommonNameReader is a Reader for the DeleteCertAuthorityOcspTLSTrustedCommonName structure.
type DeleteCertAuthorityOcspTLSTrustedCommonNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCertAuthorityOcspTLSTrustedCommonNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCertAuthorityOcspTLSTrustedCommonNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteCertAuthorityOcspTLSTrustedCommonNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCertAuthorityOcspTLSTrustedCommonNameOK creates a DeleteCertAuthorityOcspTLSTrustedCommonNameOK with default headers values
func NewDeleteCertAuthorityOcspTLSTrustedCommonNameOK() *DeleteCertAuthorityOcspTLSTrustedCommonNameOK {
	return &DeleteCertAuthorityOcspTLSTrustedCommonNameOK{}
}

/* DeleteCertAuthorityOcspTLSTrustedCommonNameOK describes a response with status code 200, with default header values.

The request metadata.
*/
type DeleteCertAuthorityOcspTLSTrustedCommonNameOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteCertAuthorityOcspTLSTrustedCommonNameOK) Error() string {
	return fmt.Sprintf("[DELETE /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName}][%d] deleteCertAuthorityOcspTlsTrustedCommonNameOK  %+v", 200, o.Payload)
}
func (o *DeleteCertAuthorityOcspTLSTrustedCommonNameOK) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *DeleteCertAuthorityOcspTLSTrustedCommonNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCertAuthorityOcspTLSTrustedCommonNameDefault creates a DeleteCertAuthorityOcspTLSTrustedCommonNameDefault with default headers values
func NewDeleteCertAuthorityOcspTLSTrustedCommonNameDefault(code int) *DeleteCertAuthorityOcspTLSTrustedCommonNameDefault {
	return &DeleteCertAuthorityOcspTLSTrustedCommonNameDefault{
		_statusCode: code,
	}
}

/* DeleteCertAuthorityOcspTLSTrustedCommonNameDefault describes a response with status code -1, with default header values.

The error response.
*/
type DeleteCertAuthorityOcspTLSTrustedCommonNameDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete cert authority ocsp Tls trusted common name default response
func (o *DeleteCertAuthorityOcspTLSTrustedCommonNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCertAuthorityOcspTLSTrustedCommonNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName}][%d] deleteCertAuthorityOcspTlsTrustedCommonName default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteCertAuthorityOcspTLSTrustedCommonNameDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *DeleteCertAuthorityOcspTLSTrustedCommonNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
