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

// CreateMsgVpnDistributedCacheClusterReader is a Reader for the CreateMsgVpnDistributedCacheCluster structure.
type CreateMsgVpnDistributedCacheClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnDistributedCacheClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateMsgVpnDistributedCacheClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateMsgVpnDistributedCacheClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnDistributedCacheClusterOK creates a CreateMsgVpnDistributedCacheClusterOK with default headers values
func NewCreateMsgVpnDistributedCacheClusterOK() *CreateMsgVpnDistributedCacheClusterOK {
	return &CreateMsgVpnDistributedCacheClusterOK{}
}

/* CreateMsgVpnDistributedCacheClusterOK describes a response with status code 200, with default header values.

The Cache Cluster object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnDistributedCacheClusterOK struct {
	Payload *models.MsgVpnDistributedCacheClusterResponse
}

func (o *CreateMsgVpnDistributedCacheClusterOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters][%d] createMsgVpnDistributedCacheClusterOK  %+v", 200, o.Payload)
}
func (o *CreateMsgVpnDistributedCacheClusterOK) GetPayload() *models.MsgVpnDistributedCacheClusterResponse {
	return o.Payload
}

func (o *CreateMsgVpnDistributedCacheClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnDistributedCacheClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnDistributedCacheClusterDefault creates a CreateMsgVpnDistributedCacheClusterDefault with default headers values
func NewCreateMsgVpnDistributedCacheClusterDefault(code int) *CreateMsgVpnDistributedCacheClusterDefault {
	return &CreateMsgVpnDistributedCacheClusterDefault{
		_statusCode: code,
	}
}

/* CreateMsgVpnDistributedCacheClusterDefault describes a response with status code -1, with default header values.

The error response.
*/
type CreateMsgVpnDistributedCacheClusterDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn distributed cache cluster default response
func (o *CreateMsgVpnDistributedCacheClusterDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnDistributedCacheClusterDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters][%d] createMsgVpnDistributedCacheCluster default  %+v", o._statusCode, o.Payload)
}
func (o *CreateMsgVpnDistributedCacheClusterDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *CreateMsgVpnDistributedCacheClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
