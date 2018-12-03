// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ExalDraen/semp-client/models"
)

// CreateMsgVpnBridgeRemoteSubscriptionReader is a Reader for the CreateMsgVpnBridgeRemoteSubscription structure.
type CreateMsgVpnBridgeRemoteSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnBridgeRemoteSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateMsgVpnBridgeRemoteSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateMsgVpnBridgeRemoteSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnBridgeRemoteSubscriptionOK creates a CreateMsgVpnBridgeRemoteSubscriptionOK with default headers values
func NewCreateMsgVpnBridgeRemoteSubscriptionOK() *CreateMsgVpnBridgeRemoteSubscriptionOK {
	return &CreateMsgVpnBridgeRemoteSubscriptionOK{}
}

/*CreateMsgVpnBridgeRemoteSubscriptionOK handles this case with default header values.

The Remote Subscription object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnBridgeRemoteSubscriptionOK struct {
	Payload *models.MsgVpnBridgeRemoteSubscriptionResponse
}

func (o *CreateMsgVpnBridgeRemoteSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions][%d] createMsgVpnBridgeRemoteSubscriptionOK  %+v", 200, o.Payload)
}

func (o *CreateMsgVpnBridgeRemoteSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnBridgeRemoteSubscriptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnBridgeRemoteSubscriptionDefault creates a CreateMsgVpnBridgeRemoteSubscriptionDefault with default headers values
func NewCreateMsgVpnBridgeRemoteSubscriptionDefault(code int) *CreateMsgVpnBridgeRemoteSubscriptionDefault {
	return &CreateMsgVpnBridgeRemoteSubscriptionDefault{
		_statusCode: code,
	}
}

/*CreateMsgVpnBridgeRemoteSubscriptionDefault handles this case with default header values.

Error response
*/
type CreateMsgVpnBridgeRemoteSubscriptionDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn bridge remote subscription default response
func (o *CreateMsgVpnBridgeRemoteSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnBridgeRemoteSubscriptionDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions][%d] createMsgVpnBridgeRemoteSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *CreateMsgVpnBridgeRemoteSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
