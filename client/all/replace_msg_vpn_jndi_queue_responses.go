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

// ReplaceMsgVpnJndiQueueReader is a Reader for the ReplaceMsgVpnJndiQueue structure.
type ReplaceMsgVpnJndiQueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceMsgVpnJndiQueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceMsgVpnJndiQueueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReplaceMsgVpnJndiQueueDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceMsgVpnJndiQueueOK creates a ReplaceMsgVpnJndiQueueOK with default headers values
func NewReplaceMsgVpnJndiQueueOK() *ReplaceMsgVpnJndiQueueOK {
	return &ReplaceMsgVpnJndiQueueOK{}
}

/* ReplaceMsgVpnJndiQueueOK describes a response with status code 200, with default header values.

The JNDI Queue object's attributes after being replaced, and the request metadata.
*/
type ReplaceMsgVpnJndiQueueOK struct {
	Payload *models.MsgVpnJndiQueueResponse
}

func (o *ReplaceMsgVpnJndiQueueOK) Error() string {
	return fmt.Sprintf("[PUT /msgVpns/{msgVpnName}/jndiQueues/{queueName}][%d] replaceMsgVpnJndiQueueOK  %+v", 200, o.Payload)
}
func (o *ReplaceMsgVpnJndiQueueOK) GetPayload() *models.MsgVpnJndiQueueResponse {
	return o.Payload
}

func (o *ReplaceMsgVpnJndiQueueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnJndiQueueResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceMsgVpnJndiQueueDefault creates a ReplaceMsgVpnJndiQueueDefault with default headers values
func NewReplaceMsgVpnJndiQueueDefault(code int) *ReplaceMsgVpnJndiQueueDefault {
	return &ReplaceMsgVpnJndiQueueDefault{
		_statusCode: code,
	}
}

/* ReplaceMsgVpnJndiQueueDefault describes a response with status code -1, with default header values.

The error response.
*/
type ReplaceMsgVpnJndiQueueDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the replace msg vpn jndi queue default response
func (o *ReplaceMsgVpnJndiQueueDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceMsgVpnJndiQueueDefault) Error() string {
	return fmt.Sprintf("[PUT /msgVpns/{msgVpnName}/jndiQueues/{queueName}][%d] replaceMsgVpnJndiQueue default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceMsgVpnJndiQueueDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *ReplaceMsgVpnJndiQueueDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
