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

// CreateMsgVpnDistributedCacheReader is a Reader for the CreateMsgVpnDistributedCache structure.
type CreateMsgVpnDistributedCacheReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnDistributedCacheReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateMsgVpnDistributedCacheOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateMsgVpnDistributedCacheDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnDistributedCacheOK creates a CreateMsgVpnDistributedCacheOK with default headers values
func NewCreateMsgVpnDistributedCacheOK() *CreateMsgVpnDistributedCacheOK {
	return &CreateMsgVpnDistributedCacheOK{}
}

/* CreateMsgVpnDistributedCacheOK describes a response with status code 200, with default header values.

The Distributed Cache object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnDistributedCacheOK struct {
	Payload *models.MsgVpnDistributedCacheResponse
}

func (o *CreateMsgVpnDistributedCacheOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/distributedCaches][%d] createMsgVpnDistributedCacheOK  %+v", 200, o.Payload)
}
func (o *CreateMsgVpnDistributedCacheOK) GetPayload() *models.MsgVpnDistributedCacheResponse {
	return o.Payload
}

func (o *CreateMsgVpnDistributedCacheOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnDistributedCacheResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnDistributedCacheDefault creates a CreateMsgVpnDistributedCacheDefault with default headers values
func NewCreateMsgVpnDistributedCacheDefault(code int) *CreateMsgVpnDistributedCacheDefault {
	return &CreateMsgVpnDistributedCacheDefault{
		_statusCode: code,
	}
}

/* CreateMsgVpnDistributedCacheDefault describes a response with status code -1, with default header values.

The error response.
*/
type CreateMsgVpnDistributedCacheDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn distributed cache default response
func (o *CreateMsgVpnDistributedCacheDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnDistributedCacheDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/distributedCaches][%d] createMsgVpnDistributedCache default  %+v", o._statusCode, o.Payload)
}
func (o *CreateMsgVpnDistributedCacheDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *CreateMsgVpnDistributedCacheDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}