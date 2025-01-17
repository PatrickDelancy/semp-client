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

// DeleteMsgVpnMqttRetainCacheReader is a Reader for the DeleteMsgVpnMqttRetainCache structure.
type DeleteMsgVpnMqttRetainCacheReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMsgVpnMqttRetainCacheReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMsgVpnMqttRetainCacheOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteMsgVpnMqttRetainCacheDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMsgVpnMqttRetainCacheOK creates a DeleteMsgVpnMqttRetainCacheOK with default headers values
func NewDeleteMsgVpnMqttRetainCacheOK() *DeleteMsgVpnMqttRetainCacheOK {
	return &DeleteMsgVpnMqttRetainCacheOK{}
}

/* DeleteMsgVpnMqttRetainCacheOK describes a response with status code 200, with default header values.

The request metadata.
*/
type DeleteMsgVpnMqttRetainCacheOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteMsgVpnMqttRetainCacheOK) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/mqttRetainCaches/{cacheName}][%d] deleteMsgVpnMqttRetainCacheOK  %+v", 200, o.Payload)
}
func (o *DeleteMsgVpnMqttRetainCacheOK) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *DeleteMsgVpnMqttRetainCacheOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMsgVpnMqttRetainCacheDefault creates a DeleteMsgVpnMqttRetainCacheDefault with default headers values
func NewDeleteMsgVpnMqttRetainCacheDefault(code int) *DeleteMsgVpnMqttRetainCacheDefault {
	return &DeleteMsgVpnMqttRetainCacheDefault{
		_statusCode: code,
	}
}

/* DeleteMsgVpnMqttRetainCacheDefault describes a response with status code -1, with default header values.

The error response.
*/
type DeleteMsgVpnMqttRetainCacheDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete msg vpn mqtt retain cache default response
func (o *DeleteMsgVpnMqttRetainCacheDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMsgVpnMqttRetainCacheDefault) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/mqttRetainCaches/{cacheName}][%d] deleteMsgVpnMqttRetainCache default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteMsgVpnMqttRetainCacheDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *DeleteMsgVpnMqttRetainCacheDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
