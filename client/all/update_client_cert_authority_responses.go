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

// UpdateClientCertAuthorityReader is a Reader for the UpdateClientCertAuthority structure.
type UpdateClientCertAuthorityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClientCertAuthorityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateClientCertAuthorityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateClientCertAuthorityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateClientCertAuthorityOK creates a UpdateClientCertAuthorityOK with default headers values
func NewUpdateClientCertAuthorityOK() *UpdateClientCertAuthorityOK {
	return &UpdateClientCertAuthorityOK{}
}

/* UpdateClientCertAuthorityOK describes a response with status code 200, with default header values.

The Client Certificate Authority object's attributes after being updated, and the request metadata.
*/
type UpdateClientCertAuthorityOK struct {
	Payload *models.ClientCertAuthorityResponse
}

func (o *UpdateClientCertAuthorityOK) Error() string {
	return fmt.Sprintf("[PATCH /clientCertAuthorities/{certAuthorityName}][%d] updateClientCertAuthorityOK  %+v", 200, o.Payload)
}
func (o *UpdateClientCertAuthorityOK) GetPayload() *models.ClientCertAuthorityResponse {
	return o.Payload
}

func (o *UpdateClientCertAuthorityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClientCertAuthorityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClientCertAuthorityDefault creates a UpdateClientCertAuthorityDefault with default headers values
func NewUpdateClientCertAuthorityDefault(code int) *UpdateClientCertAuthorityDefault {
	return &UpdateClientCertAuthorityDefault{
		_statusCode: code,
	}
}

/* UpdateClientCertAuthorityDefault describes a response with status code -1, with default header values.

The error response.
*/
type UpdateClientCertAuthorityDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the update client cert authority default response
func (o *UpdateClientCertAuthorityDefault) Code() int {
	return o._statusCode
}

func (o *UpdateClientCertAuthorityDefault) Error() string {
	return fmt.Sprintf("[PATCH /clientCertAuthorities/{certAuthorityName}][%d] updateClientCertAuthority default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateClientCertAuthorityDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *UpdateClientCertAuthorityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}