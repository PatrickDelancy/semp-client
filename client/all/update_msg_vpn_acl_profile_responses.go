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

// UpdateMsgVpnACLProfileReader is a Reader for the UpdateMsgVpnACLProfile structure.
type UpdateMsgVpnACLProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMsgVpnACLProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMsgVpnACLProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateMsgVpnACLProfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMsgVpnACLProfileOK creates a UpdateMsgVpnACLProfileOK with default headers values
func NewUpdateMsgVpnACLProfileOK() *UpdateMsgVpnACLProfileOK {
	return &UpdateMsgVpnACLProfileOK{}
}

/* UpdateMsgVpnACLProfileOK describes a response with status code 200, with default header values.

The ACL Profile object's attributes after being updated, and the request metadata.
*/
type UpdateMsgVpnACLProfileOK struct {
	Payload *models.MsgVpnACLProfileResponse
}

func (o *UpdateMsgVpnACLProfileOK) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}][%d] updateMsgVpnAclProfileOK  %+v", 200, o.Payload)
}
func (o *UpdateMsgVpnACLProfileOK) GetPayload() *models.MsgVpnACLProfileResponse {
	return o.Payload
}

func (o *UpdateMsgVpnACLProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnACLProfileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMsgVpnACLProfileDefault creates a UpdateMsgVpnACLProfileDefault with default headers values
func NewUpdateMsgVpnACLProfileDefault(code int) *UpdateMsgVpnACLProfileDefault {
	return &UpdateMsgVpnACLProfileDefault{
		_statusCode: code,
	}
}

/* UpdateMsgVpnACLProfileDefault describes a response with status code -1, with default header values.

The error response.
*/
type UpdateMsgVpnACLProfileDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the update msg vpn Acl profile default response
func (o *UpdateMsgVpnACLProfileDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMsgVpnACLProfileDefault) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}][%d] updateMsgVpnAclProfile default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateMsgVpnACLProfileDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *UpdateMsgVpnACLProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}