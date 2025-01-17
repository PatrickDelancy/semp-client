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

// CreateMsgVpnQueueSubscriptionReader is a Reader for the CreateMsgVpnQueueSubscription structure.
type CreateMsgVpnQueueSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnQueueSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateMsgVpnQueueSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateMsgVpnQueueSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnQueueSubscriptionOK creates a CreateMsgVpnQueueSubscriptionOK with default headers values
func NewCreateMsgVpnQueueSubscriptionOK() *CreateMsgVpnQueueSubscriptionOK {
	return &CreateMsgVpnQueueSubscriptionOK{}
}

/* CreateMsgVpnQueueSubscriptionOK describes a response with status code 200, with default header values.

The Queue Subscription object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnQueueSubscriptionOK struct {
	Payload *models.MsgVpnQueueSubscriptionResponse
}

func (o *CreateMsgVpnQueueSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions][%d] createMsgVpnQueueSubscriptionOK  %+v", 200, o.Payload)
}
func (o *CreateMsgVpnQueueSubscriptionOK) GetPayload() *models.MsgVpnQueueSubscriptionResponse {
	return o.Payload
}

func (o *CreateMsgVpnQueueSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnQueueSubscriptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnQueueSubscriptionDefault creates a CreateMsgVpnQueueSubscriptionDefault with default headers values
func NewCreateMsgVpnQueueSubscriptionDefault(code int) *CreateMsgVpnQueueSubscriptionDefault {
	return &CreateMsgVpnQueueSubscriptionDefault{
		_statusCode: code,
	}
}

/* CreateMsgVpnQueueSubscriptionDefault describes a response with status code -1, with default header values.

The error response.
*/
type CreateMsgVpnQueueSubscriptionDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn queue subscription default response
func (o *CreateMsgVpnQueueSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnQueueSubscriptionDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions][%d] createMsgVpnQueueSubscription default  %+v", o._statusCode, o.Payload)
}
func (o *CreateMsgVpnQueueSubscriptionDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *CreateMsgVpnQueueSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
