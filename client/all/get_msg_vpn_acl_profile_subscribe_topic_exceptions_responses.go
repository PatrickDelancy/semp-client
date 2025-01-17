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

// GetMsgVpnACLProfileSubscribeTopicExceptionsReader is a Reader for the GetMsgVpnACLProfileSubscribeTopicExceptions structure.
type GetMsgVpnACLProfileSubscribeTopicExceptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnACLProfileSubscribeTopicExceptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMsgVpnACLProfileSubscribeTopicExceptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMsgVpnACLProfileSubscribeTopicExceptionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnACLProfileSubscribeTopicExceptionsOK creates a GetMsgVpnACLProfileSubscribeTopicExceptionsOK with default headers values
func NewGetMsgVpnACLProfileSubscribeTopicExceptionsOK() *GetMsgVpnACLProfileSubscribeTopicExceptionsOK {
	return &GetMsgVpnACLProfileSubscribeTopicExceptionsOK{}
}

/* GetMsgVpnACLProfileSubscribeTopicExceptionsOK describes a response with status code 200, with default header values.

The list of Subscribe Topic Exception objects' attributes, and the request metadata.
*/
type GetMsgVpnACLProfileSubscribeTopicExceptionsOK struct {
	Payload *models.MsgVpnACLProfileSubscribeTopicExceptionsResponse
}

func (o *GetMsgVpnACLProfileSubscribeTopicExceptionsOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions][%d] getMsgVpnAclProfileSubscribeTopicExceptionsOK  %+v", 200, o.Payload)
}
func (o *GetMsgVpnACLProfileSubscribeTopicExceptionsOK) GetPayload() *models.MsgVpnACLProfileSubscribeTopicExceptionsResponse {
	return o.Payload
}

func (o *GetMsgVpnACLProfileSubscribeTopicExceptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnACLProfileSubscribeTopicExceptionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnACLProfileSubscribeTopicExceptionsDefault creates a GetMsgVpnACLProfileSubscribeTopicExceptionsDefault with default headers values
func NewGetMsgVpnACLProfileSubscribeTopicExceptionsDefault(code int) *GetMsgVpnACLProfileSubscribeTopicExceptionsDefault {
	return &GetMsgVpnACLProfileSubscribeTopicExceptionsDefault{
		_statusCode: code,
	}
}

/* GetMsgVpnACLProfileSubscribeTopicExceptionsDefault describes a response with status code -1, with default header values.

The error response.
*/
type GetMsgVpnACLProfileSubscribeTopicExceptionsDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn Acl profile subscribe topic exceptions default response
func (o *GetMsgVpnACLProfileSubscribeTopicExceptionsDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnACLProfileSubscribeTopicExceptionsDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions][%d] getMsgVpnAclProfileSubscribeTopicExceptions default  %+v", o._statusCode, o.Payload)
}
func (o *GetMsgVpnACLProfileSubscribeTopicExceptionsDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *GetMsgVpnACLProfileSubscribeTopicExceptionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
