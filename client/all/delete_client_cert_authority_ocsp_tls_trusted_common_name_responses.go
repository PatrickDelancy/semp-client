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

// DeleteClientCertAuthorityOcspTLSTrustedCommonNameReader is a Reader for the DeleteClientCertAuthorityOcspTLSTrustedCommonName structure.
type DeleteClientCertAuthorityOcspTLSTrustedCommonNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteClientCertAuthorityOcspTLSTrustedCommonNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteClientCertAuthorityOcspTLSTrustedCommonNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteClientCertAuthorityOcspTLSTrustedCommonNameOK creates a DeleteClientCertAuthorityOcspTLSTrustedCommonNameOK with default headers values
func NewDeleteClientCertAuthorityOcspTLSTrustedCommonNameOK() *DeleteClientCertAuthorityOcspTLSTrustedCommonNameOK {
	return &DeleteClientCertAuthorityOcspTLSTrustedCommonNameOK{}
}

/* DeleteClientCertAuthorityOcspTLSTrustedCommonNameOK describes a response with status code 200, with default header values.

The request metadata.
*/
type DeleteClientCertAuthorityOcspTLSTrustedCommonNameOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameOK) Error() string {
	return fmt.Sprintf("[DELETE /clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName}][%d] deleteClientCertAuthorityOcspTlsTrustedCommonNameOK  %+v", 200, o.Payload)
}
func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameOK) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClientCertAuthorityOcspTLSTrustedCommonNameDefault creates a DeleteClientCertAuthorityOcspTLSTrustedCommonNameDefault with default headers values
func NewDeleteClientCertAuthorityOcspTLSTrustedCommonNameDefault(code int) *DeleteClientCertAuthorityOcspTLSTrustedCommonNameDefault {
	return &DeleteClientCertAuthorityOcspTLSTrustedCommonNameDefault{
		_statusCode: code,
	}
}

/* DeleteClientCertAuthorityOcspTLSTrustedCommonNameDefault describes a response with status code -1, with default header values.

The error response.
*/
type DeleteClientCertAuthorityOcspTLSTrustedCommonNameDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete client cert authority ocsp Tls trusted common name default response
func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName}][%d] deleteClientCertAuthorityOcspTlsTrustedCommonName default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
