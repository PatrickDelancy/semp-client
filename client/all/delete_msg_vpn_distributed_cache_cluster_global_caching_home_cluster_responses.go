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

// DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterReader is a Reader for the DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster structure.
type DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK creates a DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK with default headers values
func NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK() *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK {
	return &DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK{}
}

/* DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK describes a response with status code 200, with default header values.

The request metadata.
*/
type DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}][%d] deleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK  %+v", 200, o.Payload)
}
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault creates a DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault with default headers values
func NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault(code int) *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault {
	return &DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault{
		_statusCode: code,
	}
}

/* DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault describes a response with status code -1, with default header values.

The error response.
*/
type DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete msg vpn distributed cache cluster global caching home cluster default response
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}][%d] deleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
