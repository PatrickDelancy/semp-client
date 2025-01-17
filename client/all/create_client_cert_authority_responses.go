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

// CreateClientCertAuthorityReader is a Reader for the CreateClientCertAuthority structure.
type CreateClientCertAuthorityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClientCertAuthorityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateClientCertAuthorityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateClientCertAuthorityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateClientCertAuthorityOK creates a CreateClientCertAuthorityOK with default headers values
func NewCreateClientCertAuthorityOK() *CreateClientCertAuthorityOK {
	return &CreateClientCertAuthorityOK{}
}

/* CreateClientCertAuthorityOK describes a response with status code 200, with default header values.

The Client Certificate Authority object's attributes after being created, and the request metadata.
*/
type CreateClientCertAuthorityOK struct {
	Payload *models.ClientCertAuthorityResponse
}

func (o *CreateClientCertAuthorityOK) Error() string {
	return fmt.Sprintf("[POST /clientCertAuthorities][%d] createClientCertAuthorityOK  %+v", 200, o.Payload)
}
func (o *CreateClientCertAuthorityOK) GetPayload() *models.ClientCertAuthorityResponse {
	return o.Payload
}

func (o *CreateClientCertAuthorityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClientCertAuthorityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClientCertAuthorityDefault creates a CreateClientCertAuthorityDefault with default headers values
func NewCreateClientCertAuthorityDefault(code int) *CreateClientCertAuthorityDefault {
	return &CreateClientCertAuthorityDefault{
		_statusCode: code,
	}
}

/* CreateClientCertAuthorityDefault describes a response with status code -1, with default header values.

The error response.
*/
type CreateClientCertAuthorityDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create client cert authority default response
func (o *CreateClientCertAuthorityDefault) Code() int {
	return o._statusCode
}

func (o *CreateClientCertAuthorityDefault) Error() string {
	return fmt.Sprintf("[POST /clientCertAuthorities][%d] createClientCertAuthority default  %+v", o._statusCode, o.Payload)
}
func (o *CreateClientCertAuthorityDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *CreateClientCertAuthorityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
