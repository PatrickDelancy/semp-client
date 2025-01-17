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

// GetMsgVpnJndiQueuesReader is a Reader for the GetMsgVpnJndiQueues structure.
type GetMsgVpnJndiQueuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnJndiQueuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMsgVpnJndiQueuesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMsgVpnJndiQueuesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnJndiQueuesOK creates a GetMsgVpnJndiQueuesOK with default headers values
func NewGetMsgVpnJndiQueuesOK() *GetMsgVpnJndiQueuesOK {
	return &GetMsgVpnJndiQueuesOK{}
}

/* GetMsgVpnJndiQueuesOK describes a response with status code 200, with default header values.

The list of JNDI Queue objects' attributes, and the request metadata.
*/
type GetMsgVpnJndiQueuesOK struct {
	Payload *models.MsgVpnJndiQueuesResponse
}

func (o *GetMsgVpnJndiQueuesOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/jndiQueues][%d] getMsgVpnJndiQueuesOK  %+v", 200, o.Payload)
}
func (o *GetMsgVpnJndiQueuesOK) GetPayload() *models.MsgVpnJndiQueuesResponse {
	return o.Payload
}

func (o *GetMsgVpnJndiQueuesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnJndiQueuesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnJndiQueuesDefault creates a GetMsgVpnJndiQueuesDefault with default headers values
func NewGetMsgVpnJndiQueuesDefault(code int) *GetMsgVpnJndiQueuesDefault {
	return &GetMsgVpnJndiQueuesDefault{
		_statusCode: code,
	}
}

/* GetMsgVpnJndiQueuesDefault describes a response with status code -1, with default header values.

The error response.
*/
type GetMsgVpnJndiQueuesDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn jndi queues default response
func (o *GetMsgVpnJndiQueuesDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnJndiQueuesDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/jndiQueues][%d] getMsgVpnJndiQueues default  %+v", o._statusCode, o.Payload)
}
func (o *GetMsgVpnJndiQueuesDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *GetMsgVpnJndiQueuesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
