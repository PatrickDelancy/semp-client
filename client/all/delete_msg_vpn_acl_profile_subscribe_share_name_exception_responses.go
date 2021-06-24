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

// DeleteMsgVpnACLProfileSubscribeShareNameExceptionReader is a Reader for the DeleteMsgVpnACLProfileSubscribeShareNameException structure.
type DeleteMsgVpnACLProfileSubscribeShareNameExceptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMsgVpnACLProfileSubscribeShareNameExceptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMsgVpnACLProfileSubscribeShareNameExceptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteMsgVpnACLProfileSubscribeShareNameExceptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMsgVpnACLProfileSubscribeShareNameExceptionOK creates a DeleteMsgVpnACLProfileSubscribeShareNameExceptionOK with default headers values
func NewDeleteMsgVpnACLProfileSubscribeShareNameExceptionOK() *DeleteMsgVpnACLProfileSubscribeShareNameExceptionOK {
	return &DeleteMsgVpnACLProfileSubscribeShareNameExceptionOK{}
}

/* DeleteMsgVpnACLProfileSubscribeShareNameExceptionOK describes a response with status code 200, with default header values.

The request metadata.
*/
type DeleteMsgVpnACLProfileSubscribeShareNameExceptionOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteMsgVpnACLProfileSubscribeShareNameExceptionOK) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions/{subscribeShareNameExceptionSyntax},{subscribeShareNameException}][%d] deleteMsgVpnAclProfileSubscribeShareNameExceptionOK  %+v", 200, o.Payload)
}
func (o *DeleteMsgVpnACLProfileSubscribeShareNameExceptionOK) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *DeleteMsgVpnACLProfileSubscribeShareNameExceptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMsgVpnACLProfileSubscribeShareNameExceptionDefault creates a DeleteMsgVpnACLProfileSubscribeShareNameExceptionDefault with default headers values
func NewDeleteMsgVpnACLProfileSubscribeShareNameExceptionDefault(code int) *DeleteMsgVpnACLProfileSubscribeShareNameExceptionDefault {
	return &DeleteMsgVpnACLProfileSubscribeShareNameExceptionDefault{
		_statusCode: code,
	}
}

/* DeleteMsgVpnACLProfileSubscribeShareNameExceptionDefault describes a response with status code -1, with default header values.

The error response.
*/
type DeleteMsgVpnACLProfileSubscribeShareNameExceptionDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete msg vpn Acl profile subscribe share name exception default response
func (o *DeleteMsgVpnACLProfileSubscribeShareNameExceptionDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMsgVpnACLProfileSubscribeShareNameExceptionDefault) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions/{subscribeShareNameExceptionSyntax},{subscribeShareNameException}][%d] deleteMsgVpnAclProfileSubscribeShareNameException default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteMsgVpnACLProfileSubscribeShareNameExceptionDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *DeleteMsgVpnACLProfileSubscribeShareNameExceptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
