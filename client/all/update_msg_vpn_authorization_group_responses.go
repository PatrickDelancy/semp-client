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

// UpdateMsgVpnAuthorizationGroupReader is a Reader for the UpdateMsgVpnAuthorizationGroup structure.
type UpdateMsgVpnAuthorizationGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMsgVpnAuthorizationGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMsgVpnAuthorizationGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateMsgVpnAuthorizationGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMsgVpnAuthorizationGroupOK creates a UpdateMsgVpnAuthorizationGroupOK with default headers values
func NewUpdateMsgVpnAuthorizationGroupOK() *UpdateMsgVpnAuthorizationGroupOK {
	return &UpdateMsgVpnAuthorizationGroupOK{}
}

/* UpdateMsgVpnAuthorizationGroupOK describes a response with status code 200, with default header values.

The LDAP Authorization Group object's attributes after being updated, and the request metadata.
*/
type UpdateMsgVpnAuthorizationGroupOK struct {
	Payload *models.MsgVpnAuthorizationGroupResponse
}

func (o *UpdateMsgVpnAuthorizationGroupOK) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName}][%d] updateMsgVpnAuthorizationGroupOK  %+v", 200, o.Payload)
}
func (o *UpdateMsgVpnAuthorizationGroupOK) GetPayload() *models.MsgVpnAuthorizationGroupResponse {
	return o.Payload
}

func (o *UpdateMsgVpnAuthorizationGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnAuthorizationGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMsgVpnAuthorizationGroupDefault creates a UpdateMsgVpnAuthorizationGroupDefault with default headers values
func NewUpdateMsgVpnAuthorizationGroupDefault(code int) *UpdateMsgVpnAuthorizationGroupDefault {
	return &UpdateMsgVpnAuthorizationGroupDefault{
		_statusCode: code,
	}
}

/* UpdateMsgVpnAuthorizationGroupDefault describes a response with status code -1, with default header values.

The error response.
*/
type UpdateMsgVpnAuthorizationGroupDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the update msg vpn authorization group default response
func (o *UpdateMsgVpnAuthorizationGroupDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMsgVpnAuthorizationGroupDefault) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName}][%d] updateMsgVpnAuthorizationGroup default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateMsgVpnAuthorizationGroupDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *UpdateMsgVpnAuthorizationGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
