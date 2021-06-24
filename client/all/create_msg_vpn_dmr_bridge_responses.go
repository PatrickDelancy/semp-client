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

// CreateMsgVpnDmrBridgeReader is a Reader for the CreateMsgVpnDmrBridge structure.
type CreateMsgVpnDmrBridgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnDmrBridgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateMsgVpnDmrBridgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateMsgVpnDmrBridgeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnDmrBridgeOK creates a CreateMsgVpnDmrBridgeOK with default headers values
func NewCreateMsgVpnDmrBridgeOK() *CreateMsgVpnDmrBridgeOK {
	return &CreateMsgVpnDmrBridgeOK{}
}

/* CreateMsgVpnDmrBridgeOK describes a response with status code 200, with default header values.

The DMR Bridge object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnDmrBridgeOK struct {
	Payload *models.MsgVpnDmrBridgeResponse
}

func (o *CreateMsgVpnDmrBridgeOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/dmrBridges][%d] createMsgVpnDmrBridgeOK  %+v", 200, o.Payload)
}
func (o *CreateMsgVpnDmrBridgeOK) GetPayload() *models.MsgVpnDmrBridgeResponse {
	return o.Payload
}

func (o *CreateMsgVpnDmrBridgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnDmrBridgeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnDmrBridgeDefault creates a CreateMsgVpnDmrBridgeDefault with default headers values
func NewCreateMsgVpnDmrBridgeDefault(code int) *CreateMsgVpnDmrBridgeDefault {
	return &CreateMsgVpnDmrBridgeDefault{
		_statusCode: code,
	}
}

/* CreateMsgVpnDmrBridgeDefault describes a response with status code -1, with default header values.

The error response.
*/
type CreateMsgVpnDmrBridgeDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn dmr bridge default response
func (o *CreateMsgVpnDmrBridgeDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnDmrBridgeDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/dmrBridges][%d] createMsgVpnDmrBridge default  %+v", o._statusCode, o.Payload)
}
func (o *CreateMsgVpnDmrBridgeDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *CreateMsgVpnDmrBridgeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}