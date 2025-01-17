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

// CreateMsgVpnBridgeRemoteMsgVpnReader is a Reader for the CreateMsgVpnBridgeRemoteMsgVpn structure.
type CreateMsgVpnBridgeRemoteMsgVpnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnBridgeRemoteMsgVpnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateMsgVpnBridgeRemoteMsgVpnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateMsgVpnBridgeRemoteMsgVpnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnBridgeRemoteMsgVpnOK creates a CreateMsgVpnBridgeRemoteMsgVpnOK with default headers values
func NewCreateMsgVpnBridgeRemoteMsgVpnOK() *CreateMsgVpnBridgeRemoteMsgVpnOK {
	return &CreateMsgVpnBridgeRemoteMsgVpnOK{}
}

/* CreateMsgVpnBridgeRemoteMsgVpnOK describes a response with status code 200, with default header values.

The Remote Message VPN object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnBridgeRemoteMsgVpnOK struct {
	Payload *models.MsgVpnBridgeRemoteMsgVpnResponse
}

func (o *CreateMsgVpnBridgeRemoteMsgVpnOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns][%d] createMsgVpnBridgeRemoteMsgVpnOK  %+v", 200, o.Payload)
}
func (o *CreateMsgVpnBridgeRemoteMsgVpnOK) GetPayload() *models.MsgVpnBridgeRemoteMsgVpnResponse {
	return o.Payload
}

func (o *CreateMsgVpnBridgeRemoteMsgVpnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnBridgeRemoteMsgVpnResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnBridgeRemoteMsgVpnDefault creates a CreateMsgVpnBridgeRemoteMsgVpnDefault with default headers values
func NewCreateMsgVpnBridgeRemoteMsgVpnDefault(code int) *CreateMsgVpnBridgeRemoteMsgVpnDefault {
	return &CreateMsgVpnBridgeRemoteMsgVpnDefault{
		_statusCode: code,
	}
}

/* CreateMsgVpnBridgeRemoteMsgVpnDefault describes a response with status code -1, with default header values.

The error response.
*/
type CreateMsgVpnBridgeRemoteMsgVpnDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn bridge remote msg vpn default response
func (o *CreateMsgVpnBridgeRemoteMsgVpnDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnBridgeRemoteMsgVpnDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns][%d] createMsgVpnBridgeRemoteMsgVpn default  %+v", o._statusCode, o.Payload)
}
func (o *CreateMsgVpnBridgeRemoteMsgVpnDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *CreateMsgVpnBridgeRemoteMsgVpnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
