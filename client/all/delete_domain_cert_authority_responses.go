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

// DeleteDomainCertAuthorityReader is a Reader for the DeleteDomainCertAuthority structure.
type DeleteDomainCertAuthorityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDomainCertAuthorityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDomainCertAuthorityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteDomainCertAuthorityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDomainCertAuthorityOK creates a DeleteDomainCertAuthorityOK with default headers values
func NewDeleteDomainCertAuthorityOK() *DeleteDomainCertAuthorityOK {
	return &DeleteDomainCertAuthorityOK{}
}

/* DeleteDomainCertAuthorityOK describes a response with status code 200, with default header values.

The request metadata.
*/
type DeleteDomainCertAuthorityOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteDomainCertAuthorityOK) Error() string {
	return fmt.Sprintf("[DELETE /domainCertAuthorities/{certAuthorityName}][%d] deleteDomainCertAuthorityOK  %+v", 200, o.Payload)
}
func (o *DeleteDomainCertAuthorityOK) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *DeleteDomainCertAuthorityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDomainCertAuthorityDefault creates a DeleteDomainCertAuthorityDefault with default headers values
func NewDeleteDomainCertAuthorityDefault(code int) *DeleteDomainCertAuthorityDefault {
	return &DeleteDomainCertAuthorityDefault{
		_statusCode: code,
	}
}

/* DeleteDomainCertAuthorityDefault describes a response with status code -1, with default header values.

The error response.
*/
type DeleteDomainCertAuthorityDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete domain cert authority default response
func (o *DeleteDomainCertAuthorityDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDomainCertAuthorityDefault) Error() string {
	return fmt.Sprintf("[DELETE /domainCertAuthorities/{certAuthorityName}][%d] deleteDomainCertAuthority default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteDomainCertAuthorityDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *DeleteDomainCertAuthorityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}