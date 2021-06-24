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

// GetDomainCertAuthoritiesReader is a Reader for the GetDomainCertAuthorities structure.
type GetDomainCertAuthoritiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomainCertAuthoritiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDomainCertAuthoritiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDomainCertAuthoritiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDomainCertAuthoritiesOK creates a GetDomainCertAuthoritiesOK with default headers values
func NewGetDomainCertAuthoritiesOK() *GetDomainCertAuthoritiesOK {
	return &GetDomainCertAuthoritiesOK{}
}

/* GetDomainCertAuthoritiesOK describes a response with status code 200, with default header values.

The list of Domain Certificate Authority objects' attributes, and the request metadata.
*/
type GetDomainCertAuthoritiesOK struct {
	Payload *models.DomainCertAuthoritiesResponse
}

func (o *GetDomainCertAuthoritiesOK) Error() string {
	return fmt.Sprintf("[GET /domainCertAuthorities][%d] getDomainCertAuthoritiesOK  %+v", 200, o.Payload)
}
func (o *GetDomainCertAuthoritiesOK) GetPayload() *models.DomainCertAuthoritiesResponse {
	return o.Payload
}

func (o *GetDomainCertAuthoritiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainCertAuthoritiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainCertAuthoritiesDefault creates a GetDomainCertAuthoritiesDefault with default headers values
func NewGetDomainCertAuthoritiesDefault(code int) *GetDomainCertAuthoritiesDefault {
	return &GetDomainCertAuthoritiesDefault{
		_statusCode: code,
	}
}

/* GetDomainCertAuthoritiesDefault describes a response with status code -1, with default header values.

The error response.
*/
type GetDomainCertAuthoritiesDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get domain cert authorities default response
func (o *GetDomainCertAuthoritiesDefault) Code() int {
	return o._statusCode
}

func (o *GetDomainCertAuthoritiesDefault) Error() string {
	return fmt.Sprintf("[GET /domainCertAuthorities][%d] getDomainCertAuthorities default  %+v", o._statusCode, o.Payload)
}
func (o *GetDomainCertAuthoritiesDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *GetDomainCertAuthoritiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
