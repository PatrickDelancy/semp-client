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

// UpdateMsgVpnAuthenticationOauthProviderReader is a Reader for the UpdateMsgVpnAuthenticationOauthProvider structure.
type UpdateMsgVpnAuthenticationOauthProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMsgVpnAuthenticationOauthProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMsgVpnAuthenticationOauthProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateMsgVpnAuthenticationOauthProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMsgVpnAuthenticationOauthProviderOK creates a UpdateMsgVpnAuthenticationOauthProviderOK with default headers values
func NewUpdateMsgVpnAuthenticationOauthProviderOK() *UpdateMsgVpnAuthenticationOauthProviderOK {
	return &UpdateMsgVpnAuthenticationOauthProviderOK{}
}

/* UpdateMsgVpnAuthenticationOauthProviderOK describes a response with status code 200, with default header values.

The OAuth Provider object's attributes after being updated, and the request metadata.
*/
type UpdateMsgVpnAuthenticationOauthProviderOK struct {
	Payload *models.MsgVpnAuthenticationOauthProviderResponse
}

func (o *UpdateMsgVpnAuthenticationOauthProviderOK) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName}][%d] updateMsgVpnAuthenticationOauthProviderOK  %+v", 200, o.Payload)
}
func (o *UpdateMsgVpnAuthenticationOauthProviderOK) GetPayload() *models.MsgVpnAuthenticationOauthProviderResponse {
	return o.Payload
}

func (o *UpdateMsgVpnAuthenticationOauthProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnAuthenticationOauthProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMsgVpnAuthenticationOauthProviderDefault creates a UpdateMsgVpnAuthenticationOauthProviderDefault with default headers values
func NewUpdateMsgVpnAuthenticationOauthProviderDefault(code int) *UpdateMsgVpnAuthenticationOauthProviderDefault {
	return &UpdateMsgVpnAuthenticationOauthProviderDefault{
		_statusCode: code,
	}
}

/* UpdateMsgVpnAuthenticationOauthProviderDefault describes a response with status code -1, with default header values.

The error response.
*/
type UpdateMsgVpnAuthenticationOauthProviderDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the update msg vpn authentication oauth provider default response
func (o *UpdateMsgVpnAuthenticationOauthProviderDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMsgVpnAuthenticationOauthProviderDefault) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName}][%d] updateMsgVpnAuthenticationOauthProvider default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateMsgVpnAuthenticationOauthProviderDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *UpdateMsgVpnAuthenticationOauthProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}