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

// CreateMsgVpnACLProfileSubscribeShareNameExceptionReader is a Reader for the CreateMsgVpnACLProfileSubscribeShareNameException structure.
type CreateMsgVpnACLProfileSubscribeShareNameExceptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnACLProfileSubscribeShareNameExceptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateMsgVpnACLProfileSubscribeShareNameExceptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateMsgVpnACLProfileSubscribeShareNameExceptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnACLProfileSubscribeShareNameExceptionOK creates a CreateMsgVpnACLProfileSubscribeShareNameExceptionOK with default headers values
func NewCreateMsgVpnACLProfileSubscribeShareNameExceptionOK() *CreateMsgVpnACLProfileSubscribeShareNameExceptionOK {
	return &CreateMsgVpnACLProfileSubscribeShareNameExceptionOK{}
}

/* CreateMsgVpnACLProfileSubscribeShareNameExceptionOK describes a response with status code 200, with default header values.

The Subscribe Share Name Exception object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnACLProfileSubscribeShareNameExceptionOK struct {
	Payload *models.MsgVpnACLProfileSubscribeShareNameExceptionResponse
}

func (o *CreateMsgVpnACLProfileSubscribeShareNameExceptionOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions][%d] createMsgVpnAclProfileSubscribeShareNameExceptionOK  %+v", 200, o.Payload)
}
func (o *CreateMsgVpnACLProfileSubscribeShareNameExceptionOK) GetPayload() *models.MsgVpnACLProfileSubscribeShareNameExceptionResponse {
	return o.Payload
}

func (o *CreateMsgVpnACLProfileSubscribeShareNameExceptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnACLProfileSubscribeShareNameExceptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnACLProfileSubscribeShareNameExceptionDefault creates a CreateMsgVpnACLProfileSubscribeShareNameExceptionDefault with default headers values
func NewCreateMsgVpnACLProfileSubscribeShareNameExceptionDefault(code int) *CreateMsgVpnACLProfileSubscribeShareNameExceptionDefault {
	return &CreateMsgVpnACLProfileSubscribeShareNameExceptionDefault{
		_statusCode: code,
	}
}

/* CreateMsgVpnACLProfileSubscribeShareNameExceptionDefault describes a response with status code -1, with default header values.

The error response.
*/
type CreateMsgVpnACLProfileSubscribeShareNameExceptionDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn Acl profile subscribe share name exception default response
func (o *CreateMsgVpnACLProfileSubscribeShareNameExceptionDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnACLProfileSubscribeShareNameExceptionDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions][%d] createMsgVpnAclProfileSubscribeShareNameException default  %+v", o._statusCode, o.Payload)
}
func (o *CreateMsgVpnACLProfileSubscribeShareNameExceptionDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *CreateMsgVpnACLProfileSubscribeShareNameExceptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
