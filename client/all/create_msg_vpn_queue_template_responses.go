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

// CreateMsgVpnQueueTemplateReader is a Reader for the CreateMsgVpnQueueTemplate structure.
type CreateMsgVpnQueueTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnQueueTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateMsgVpnQueueTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateMsgVpnQueueTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnQueueTemplateOK creates a CreateMsgVpnQueueTemplateOK with default headers values
func NewCreateMsgVpnQueueTemplateOK() *CreateMsgVpnQueueTemplateOK {
	return &CreateMsgVpnQueueTemplateOK{}
}

/* CreateMsgVpnQueueTemplateOK describes a response with status code 200, with default header values.

The Queue Template object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnQueueTemplateOK struct {
	Payload *models.MsgVpnQueueTemplateResponse
}

func (o *CreateMsgVpnQueueTemplateOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/queueTemplates][%d] createMsgVpnQueueTemplateOK  %+v", 200, o.Payload)
}
func (o *CreateMsgVpnQueueTemplateOK) GetPayload() *models.MsgVpnQueueTemplateResponse {
	return o.Payload
}

func (o *CreateMsgVpnQueueTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnQueueTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnQueueTemplateDefault creates a CreateMsgVpnQueueTemplateDefault with default headers values
func NewCreateMsgVpnQueueTemplateDefault(code int) *CreateMsgVpnQueueTemplateDefault {
	return &CreateMsgVpnQueueTemplateDefault{
		_statusCode: code,
	}
}

/* CreateMsgVpnQueueTemplateDefault describes a response with status code -1, with default header values.

The error response.
*/
type CreateMsgVpnQueueTemplateDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn queue template default response
func (o *CreateMsgVpnQueueTemplateDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnQueueTemplateDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/queueTemplates][%d] createMsgVpnQueueTemplate default  %+v", o._statusCode, o.Payload)
}
func (o *CreateMsgVpnQueueTemplateDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *CreateMsgVpnQueueTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
