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

// UpdateMsgVpnQueueReader is a Reader for the UpdateMsgVpnQueue structure.
type UpdateMsgVpnQueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMsgVpnQueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateMsgVpnQueueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateMsgVpnQueueDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMsgVpnQueueOK creates a UpdateMsgVpnQueueOK with default headers values
func NewUpdateMsgVpnQueueOK() *UpdateMsgVpnQueueOK {
	return &UpdateMsgVpnQueueOK{}
}

/*UpdateMsgVpnQueueOK handles this case with default header values.

The Queue object's attributes after being updated, and the request metadata.
*/
type UpdateMsgVpnQueueOK struct {
	Payload *models.MsgVpnQueueResponse
}

func (o *UpdateMsgVpnQueueOK) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/queues/{queueName}][%d] updateMsgVpnQueueOK  %+v", 200, o.Payload)
}

func (o *UpdateMsgVpnQueueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnQueueResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMsgVpnQueueDefault creates a UpdateMsgVpnQueueDefault with default headers values
func NewUpdateMsgVpnQueueDefault(code int) *UpdateMsgVpnQueueDefault {
	return &UpdateMsgVpnQueueDefault{
		_statusCode: code,
	}
}

/*UpdateMsgVpnQueueDefault handles this case with default header values.

Error response
*/
type UpdateMsgVpnQueueDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the update msg vpn queue default response
func (o *UpdateMsgVpnQueueDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMsgVpnQueueDefault) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/queues/{queueName}][%d] updateMsgVpnQueue default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateMsgVpnQueueDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
