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

// CreateMsgVpnAuthorizationGroupReader is a Reader for the CreateMsgVpnAuthorizationGroup structure.
type CreateMsgVpnAuthorizationGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnAuthorizationGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateMsgVpnAuthorizationGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateMsgVpnAuthorizationGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnAuthorizationGroupOK creates a CreateMsgVpnAuthorizationGroupOK with default headers values
func NewCreateMsgVpnAuthorizationGroupOK() *CreateMsgVpnAuthorizationGroupOK {
	return &CreateMsgVpnAuthorizationGroupOK{}
}

/*CreateMsgVpnAuthorizationGroupOK handles this case with default header values.

The LDAP Authorization Group object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnAuthorizationGroupOK struct {
	Payload *models.MsgVpnAuthorizationGroupResponse
}

func (o *CreateMsgVpnAuthorizationGroupOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/authorizationGroups][%d] createMsgVpnAuthorizationGroupOK  %+v", 200, o.Payload)
}

func (o *CreateMsgVpnAuthorizationGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnAuthorizationGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnAuthorizationGroupDefault creates a CreateMsgVpnAuthorizationGroupDefault with default headers values
func NewCreateMsgVpnAuthorizationGroupDefault(code int) *CreateMsgVpnAuthorizationGroupDefault {
	return &CreateMsgVpnAuthorizationGroupDefault{
		_statusCode: code,
	}
}

/*CreateMsgVpnAuthorizationGroupDefault handles this case with default header values.

Error response
*/
type CreateMsgVpnAuthorizationGroupDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn authorization group default response
func (o *CreateMsgVpnAuthorizationGroupDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnAuthorizationGroupDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/authorizationGroups][%d] createMsgVpnAuthorizationGroup default  %+v", o._statusCode, o.Payload)
}

func (o *CreateMsgVpnAuthorizationGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
