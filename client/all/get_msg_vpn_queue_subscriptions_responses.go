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

// GetMsgVpnQueueSubscriptionsReader is a Reader for the GetMsgVpnQueueSubscriptions structure.
type GetMsgVpnQueueSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnQueueSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMsgVpnQueueSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMsgVpnQueueSubscriptionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnQueueSubscriptionsOK creates a GetMsgVpnQueueSubscriptionsOK with default headers values
func NewGetMsgVpnQueueSubscriptionsOK() *GetMsgVpnQueueSubscriptionsOK {
	return &GetMsgVpnQueueSubscriptionsOK{}
}

/* GetMsgVpnQueueSubscriptionsOK describes a response with status code 200, with default header values.

The list of Queue Subscription objects' attributes, and the request metadata.
*/
type GetMsgVpnQueueSubscriptionsOK struct {
	Payload *models.MsgVpnQueueSubscriptionsResponse
}

func (o *GetMsgVpnQueueSubscriptionsOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions][%d] getMsgVpnQueueSubscriptionsOK  %+v", 200, o.Payload)
}
func (o *GetMsgVpnQueueSubscriptionsOK) GetPayload() *models.MsgVpnQueueSubscriptionsResponse {
	return o.Payload
}

func (o *GetMsgVpnQueueSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnQueueSubscriptionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnQueueSubscriptionsDefault creates a GetMsgVpnQueueSubscriptionsDefault with default headers values
func NewGetMsgVpnQueueSubscriptionsDefault(code int) *GetMsgVpnQueueSubscriptionsDefault {
	return &GetMsgVpnQueueSubscriptionsDefault{
		_statusCode: code,
	}
}

/* GetMsgVpnQueueSubscriptionsDefault describes a response with status code -1, with default header values.

The error response.
*/
type GetMsgVpnQueueSubscriptionsDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn queue subscriptions default response
func (o *GetMsgVpnQueueSubscriptionsDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnQueueSubscriptionsDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions][%d] getMsgVpnQueueSubscriptions default  %+v", o._statusCode, o.Payload)
}
func (o *GetMsgVpnQueueSubscriptionsDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *GetMsgVpnQueueSubscriptionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
