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

// UpdateMsgVpnJndiQueueReader is a Reader for the UpdateMsgVpnJndiQueue structure.
type UpdateMsgVpnJndiQueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMsgVpnJndiQueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMsgVpnJndiQueueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateMsgVpnJndiQueueDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMsgVpnJndiQueueOK creates a UpdateMsgVpnJndiQueueOK with default headers values
func NewUpdateMsgVpnJndiQueueOK() *UpdateMsgVpnJndiQueueOK {
	return &UpdateMsgVpnJndiQueueOK{}
}

/* UpdateMsgVpnJndiQueueOK describes a response with status code 200, with default header values.

The JNDI Queue object's attributes after being updated, and the request metadata.
*/
type UpdateMsgVpnJndiQueueOK struct {
	Payload *models.MsgVpnJndiQueueResponse
}

func (o *UpdateMsgVpnJndiQueueOK) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/jndiQueues/{queueName}][%d] updateMsgVpnJndiQueueOK  %+v", 200, o.Payload)
}
func (o *UpdateMsgVpnJndiQueueOK) GetPayload() *models.MsgVpnJndiQueueResponse {
	return o.Payload
}

func (o *UpdateMsgVpnJndiQueueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnJndiQueueResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMsgVpnJndiQueueDefault creates a UpdateMsgVpnJndiQueueDefault with default headers values
func NewUpdateMsgVpnJndiQueueDefault(code int) *UpdateMsgVpnJndiQueueDefault {
	return &UpdateMsgVpnJndiQueueDefault{
		_statusCode: code,
	}
}

/* UpdateMsgVpnJndiQueueDefault describes a response with status code -1, with default header values.

The error response.
*/
type UpdateMsgVpnJndiQueueDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the update msg vpn jndi queue default response
func (o *UpdateMsgVpnJndiQueueDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMsgVpnJndiQueueDefault) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/jndiQueues/{queueName}][%d] updateMsgVpnJndiQueue default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateMsgVpnJndiQueueDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *UpdateMsgVpnJndiQueueDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}