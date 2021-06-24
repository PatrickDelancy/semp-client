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

// GetMsgVpnDistributedCacheClusterInstancesReader is a Reader for the GetMsgVpnDistributedCacheClusterInstances structure.
type GetMsgVpnDistributedCacheClusterInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnDistributedCacheClusterInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMsgVpnDistributedCacheClusterInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMsgVpnDistributedCacheClusterInstancesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnDistributedCacheClusterInstancesOK creates a GetMsgVpnDistributedCacheClusterInstancesOK with default headers values
func NewGetMsgVpnDistributedCacheClusterInstancesOK() *GetMsgVpnDistributedCacheClusterInstancesOK {
	return &GetMsgVpnDistributedCacheClusterInstancesOK{}
}

/* GetMsgVpnDistributedCacheClusterInstancesOK describes a response with status code 200, with default header values.

The list of Cache Instance objects' attributes, and the request metadata.
*/
type GetMsgVpnDistributedCacheClusterInstancesOK struct {
	Payload *models.MsgVpnDistributedCacheClusterInstancesResponse
}

func (o *GetMsgVpnDistributedCacheClusterInstancesOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances][%d] getMsgVpnDistributedCacheClusterInstancesOK  %+v", 200, o.Payload)
}
func (o *GetMsgVpnDistributedCacheClusterInstancesOK) GetPayload() *models.MsgVpnDistributedCacheClusterInstancesResponse {
	return o.Payload
}

func (o *GetMsgVpnDistributedCacheClusterInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnDistributedCacheClusterInstancesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnDistributedCacheClusterInstancesDefault creates a GetMsgVpnDistributedCacheClusterInstancesDefault with default headers values
func NewGetMsgVpnDistributedCacheClusterInstancesDefault(code int) *GetMsgVpnDistributedCacheClusterInstancesDefault {
	return &GetMsgVpnDistributedCacheClusterInstancesDefault{
		_statusCode: code,
	}
}

/* GetMsgVpnDistributedCacheClusterInstancesDefault describes a response with status code -1, with default header values.

The error response.
*/
type GetMsgVpnDistributedCacheClusterInstancesDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn distributed cache cluster instances default response
func (o *GetMsgVpnDistributedCacheClusterInstancesDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnDistributedCacheClusterInstancesDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances][%d] getMsgVpnDistributedCacheClusterInstances default  %+v", o._statusCode, o.Payload)
}
func (o *GetMsgVpnDistributedCacheClusterInstancesDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *GetMsgVpnDistributedCacheClusterInstancesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}