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

// GetMsgVpnQueuesReader is a Reader for the GetMsgVpnQueues structure.
type GetMsgVpnQueuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnQueuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMsgVpnQueuesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMsgVpnQueuesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnQueuesOK creates a GetMsgVpnQueuesOK with default headers values
func NewGetMsgVpnQueuesOK() *GetMsgVpnQueuesOK {
	return &GetMsgVpnQueuesOK{}
}

/* GetMsgVpnQueuesOK describes a response with status code 200, with default header values.

The list of Queue objects' attributes, and the request metadata.
*/
type GetMsgVpnQueuesOK struct {
	Payload *models.MsgVpnQueuesResponse
}

func (o *GetMsgVpnQueuesOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/queues][%d] getMsgVpnQueuesOK  %+v", 200, o.Payload)
}
func (o *GetMsgVpnQueuesOK) GetPayload() *models.MsgVpnQueuesResponse {
	return o.Payload
}

func (o *GetMsgVpnQueuesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnQueuesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnQueuesDefault creates a GetMsgVpnQueuesDefault with default headers values
func NewGetMsgVpnQueuesDefault(code int) *GetMsgVpnQueuesDefault {
	return &GetMsgVpnQueuesDefault{
		_statusCode: code,
	}
}

/* GetMsgVpnQueuesDefault describes a response with status code -1, with default header values.

The error response.
*/
type GetMsgVpnQueuesDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn queues default response
func (o *GetMsgVpnQueuesDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnQueuesDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/queues][%d] getMsgVpnQueues default  %+v", o._statusCode, o.Payload)
}
func (o *GetMsgVpnQueuesDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *GetMsgVpnQueuesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}