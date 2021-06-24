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

// GetMsgVpnRestDeliveryPointRestConsumersReader is a Reader for the GetMsgVpnRestDeliveryPointRestConsumers structure.
type GetMsgVpnRestDeliveryPointRestConsumersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnRestDeliveryPointRestConsumersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMsgVpnRestDeliveryPointRestConsumersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMsgVpnRestDeliveryPointRestConsumersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnRestDeliveryPointRestConsumersOK creates a GetMsgVpnRestDeliveryPointRestConsumersOK with default headers values
func NewGetMsgVpnRestDeliveryPointRestConsumersOK() *GetMsgVpnRestDeliveryPointRestConsumersOK {
	return &GetMsgVpnRestDeliveryPointRestConsumersOK{}
}

/* GetMsgVpnRestDeliveryPointRestConsumersOK describes a response with status code 200, with default header values.

The list of REST Consumer objects' attributes, and the request metadata.
*/
type GetMsgVpnRestDeliveryPointRestConsumersOK struct {
	Payload *models.MsgVpnRestDeliveryPointRestConsumersResponse
}

func (o *GetMsgVpnRestDeliveryPointRestConsumersOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers][%d] getMsgVpnRestDeliveryPointRestConsumersOK  %+v", 200, o.Payload)
}
func (o *GetMsgVpnRestDeliveryPointRestConsumersOK) GetPayload() *models.MsgVpnRestDeliveryPointRestConsumersResponse {
	return o.Payload
}

func (o *GetMsgVpnRestDeliveryPointRestConsumersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnRestDeliveryPointRestConsumersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnRestDeliveryPointRestConsumersDefault creates a GetMsgVpnRestDeliveryPointRestConsumersDefault with default headers values
func NewGetMsgVpnRestDeliveryPointRestConsumersDefault(code int) *GetMsgVpnRestDeliveryPointRestConsumersDefault {
	return &GetMsgVpnRestDeliveryPointRestConsumersDefault{
		_statusCode: code,
	}
}

/* GetMsgVpnRestDeliveryPointRestConsumersDefault describes a response with status code -1, with default header values.

The error response.
*/
type GetMsgVpnRestDeliveryPointRestConsumersDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn rest delivery point rest consumers default response
func (o *GetMsgVpnRestDeliveryPointRestConsumersDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnRestDeliveryPointRestConsumersDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers][%d] getMsgVpnRestDeliveryPointRestConsumers default  %+v", o._statusCode, o.Payload)
}
func (o *GetMsgVpnRestDeliveryPointRestConsumersDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *GetMsgVpnRestDeliveryPointRestConsumersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}