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

// DeleteMsgVpnDistributedCacheClusterInstanceReader is a Reader for the DeleteMsgVpnDistributedCacheClusterInstance structure.
type DeleteMsgVpnDistributedCacheClusterInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMsgVpnDistributedCacheClusterInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMsgVpnDistributedCacheClusterInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteMsgVpnDistributedCacheClusterInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMsgVpnDistributedCacheClusterInstanceOK creates a DeleteMsgVpnDistributedCacheClusterInstanceOK with default headers values
func NewDeleteMsgVpnDistributedCacheClusterInstanceOK() *DeleteMsgVpnDistributedCacheClusterInstanceOK {
	return &DeleteMsgVpnDistributedCacheClusterInstanceOK{}
}

/* DeleteMsgVpnDistributedCacheClusterInstanceOK describes a response with status code 200, with default header values.

The request metadata.
*/
type DeleteMsgVpnDistributedCacheClusterInstanceOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteMsgVpnDistributedCacheClusterInstanceOK) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}][%d] deleteMsgVpnDistributedCacheClusterInstanceOK  %+v", 200, o.Payload)
}
func (o *DeleteMsgVpnDistributedCacheClusterInstanceOK) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *DeleteMsgVpnDistributedCacheClusterInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMsgVpnDistributedCacheClusterInstanceDefault creates a DeleteMsgVpnDistributedCacheClusterInstanceDefault with default headers values
func NewDeleteMsgVpnDistributedCacheClusterInstanceDefault(code int) *DeleteMsgVpnDistributedCacheClusterInstanceDefault {
	return &DeleteMsgVpnDistributedCacheClusterInstanceDefault{
		_statusCode: code,
	}
}

/* DeleteMsgVpnDistributedCacheClusterInstanceDefault describes a response with status code -1, with default header values.

The error response.
*/
type DeleteMsgVpnDistributedCacheClusterInstanceDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete msg vpn distributed cache cluster instance default response
func (o *DeleteMsgVpnDistributedCacheClusterInstanceDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMsgVpnDistributedCacheClusterInstanceDefault) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}][%d] deleteMsgVpnDistributedCacheClusterInstance default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteMsgVpnDistributedCacheClusterInstanceDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *DeleteMsgVpnDistributedCacheClusterInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
