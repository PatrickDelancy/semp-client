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

// CreateMsgVpnMqttSessionSubscriptionReader is a Reader for the CreateMsgVpnMqttSessionSubscription structure.
type CreateMsgVpnMqttSessionSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnMqttSessionSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateMsgVpnMqttSessionSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateMsgVpnMqttSessionSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnMqttSessionSubscriptionOK creates a CreateMsgVpnMqttSessionSubscriptionOK with default headers values
func NewCreateMsgVpnMqttSessionSubscriptionOK() *CreateMsgVpnMqttSessionSubscriptionOK {
	return &CreateMsgVpnMqttSessionSubscriptionOK{}
}

/* CreateMsgVpnMqttSessionSubscriptionOK describes a response with status code 200, with default header values.

The Subscription object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnMqttSessionSubscriptionOK struct {
	Payload *models.MsgVpnMqttSessionSubscriptionResponse
}

func (o *CreateMsgVpnMqttSessionSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions][%d] createMsgVpnMqttSessionSubscriptionOK  %+v", 200, o.Payload)
}
func (o *CreateMsgVpnMqttSessionSubscriptionOK) GetPayload() *models.MsgVpnMqttSessionSubscriptionResponse {
	return o.Payload
}

func (o *CreateMsgVpnMqttSessionSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnMqttSessionSubscriptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnMqttSessionSubscriptionDefault creates a CreateMsgVpnMqttSessionSubscriptionDefault with default headers values
func NewCreateMsgVpnMqttSessionSubscriptionDefault(code int) *CreateMsgVpnMqttSessionSubscriptionDefault {
	return &CreateMsgVpnMqttSessionSubscriptionDefault{
		_statusCode: code,
	}
}

/* CreateMsgVpnMqttSessionSubscriptionDefault describes a response with status code -1, with default header values.

The error response.
*/
type CreateMsgVpnMqttSessionSubscriptionDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn mqtt session subscription default response
func (o *CreateMsgVpnMqttSessionSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnMqttSessionSubscriptionDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions][%d] createMsgVpnMqttSessionSubscription default  %+v", o._statusCode, o.Payload)
}
func (o *CreateMsgVpnMqttSessionSubscriptionDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *CreateMsgVpnMqttSessionSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}