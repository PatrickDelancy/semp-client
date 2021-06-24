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

// ReplaceDomainCertAuthorityReader is a Reader for the ReplaceDomainCertAuthority structure.
type ReplaceDomainCertAuthorityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceDomainCertAuthorityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceDomainCertAuthorityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReplaceDomainCertAuthorityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceDomainCertAuthorityOK creates a ReplaceDomainCertAuthorityOK with default headers values
func NewReplaceDomainCertAuthorityOK() *ReplaceDomainCertAuthorityOK {
	return &ReplaceDomainCertAuthorityOK{}
}

/* ReplaceDomainCertAuthorityOK describes a response with status code 200, with default header values.

The Domain Certificate Authority object's attributes after being replaced, and the request metadata.
*/
type ReplaceDomainCertAuthorityOK struct {
	Payload *models.DomainCertAuthorityResponse
}

func (o *ReplaceDomainCertAuthorityOK) Error() string {
	return fmt.Sprintf("[PUT /domainCertAuthorities/{certAuthorityName}][%d] replaceDomainCertAuthorityOK  %+v", 200, o.Payload)
}
func (o *ReplaceDomainCertAuthorityOK) GetPayload() *models.DomainCertAuthorityResponse {
	return o.Payload
}

func (o *ReplaceDomainCertAuthorityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainCertAuthorityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceDomainCertAuthorityDefault creates a ReplaceDomainCertAuthorityDefault with default headers values
func NewReplaceDomainCertAuthorityDefault(code int) *ReplaceDomainCertAuthorityDefault {
	return &ReplaceDomainCertAuthorityDefault{
		_statusCode: code,
	}
}

/* ReplaceDomainCertAuthorityDefault describes a response with status code -1, with default header values.

The error response.
*/
type ReplaceDomainCertAuthorityDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the replace domain cert authority default response
func (o *ReplaceDomainCertAuthorityDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceDomainCertAuthorityDefault) Error() string {
	return fmt.Sprintf("[PUT /domainCertAuthorities/{certAuthorityName}][%d] replaceDomainCertAuthority default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceDomainCertAuthorityDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *ReplaceDomainCertAuthorityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
