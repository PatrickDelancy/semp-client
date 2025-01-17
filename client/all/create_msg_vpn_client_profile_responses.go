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

// CreateMsgVpnClientProfileReader is a Reader for the CreateMsgVpnClientProfile structure.
type CreateMsgVpnClientProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnClientProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateMsgVpnClientProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateMsgVpnClientProfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnClientProfileOK creates a CreateMsgVpnClientProfileOK with default headers values
func NewCreateMsgVpnClientProfileOK() *CreateMsgVpnClientProfileOK {
	return &CreateMsgVpnClientProfileOK{}
}

/* CreateMsgVpnClientProfileOK describes a response with status code 200, with default header values.

The Client Profile object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnClientProfileOK struct {
	Payload *models.MsgVpnClientProfileResponse
}

func (o *CreateMsgVpnClientProfileOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/clientProfiles][%d] createMsgVpnClientProfileOK  %+v", 200, o.Payload)
}
func (o *CreateMsgVpnClientProfileOK) GetPayload() *models.MsgVpnClientProfileResponse {
	return o.Payload
}

func (o *CreateMsgVpnClientProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnClientProfileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnClientProfileDefault creates a CreateMsgVpnClientProfileDefault with default headers values
func NewCreateMsgVpnClientProfileDefault(code int) *CreateMsgVpnClientProfileDefault {
	return &CreateMsgVpnClientProfileDefault{
		_statusCode: code,
	}
}

/* CreateMsgVpnClientProfileDefault describes a response with status code -1, with default header values.

The error response.
*/
type CreateMsgVpnClientProfileDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn client profile default response
func (o *CreateMsgVpnClientProfileDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnClientProfileDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/clientProfiles][%d] createMsgVpnClientProfile default  %+v", o._statusCode, o.Payload)
}
func (o *CreateMsgVpnClientProfileDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *CreateMsgVpnClientProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
