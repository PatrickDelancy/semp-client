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

// GetMsgVpnAuthenticationOauthProviderReader is a Reader for the GetMsgVpnAuthenticationOauthProvider structure.
type GetMsgVpnAuthenticationOauthProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnAuthenticationOauthProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMsgVpnAuthenticationOauthProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMsgVpnAuthenticationOauthProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnAuthenticationOauthProviderOK creates a GetMsgVpnAuthenticationOauthProviderOK with default headers values
func NewGetMsgVpnAuthenticationOauthProviderOK() *GetMsgVpnAuthenticationOauthProviderOK {
	return &GetMsgVpnAuthenticationOauthProviderOK{}
}

/* GetMsgVpnAuthenticationOauthProviderOK describes a response with status code 200, with default header values.

The OAuth Provider object's attributes, and the request metadata.
*/
type GetMsgVpnAuthenticationOauthProviderOK struct {
	Payload *models.MsgVpnAuthenticationOauthProviderResponse
}

func (o *GetMsgVpnAuthenticationOauthProviderOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName}][%d] getMsgVpnAuthenticationOauthProviderOK  %+v", 200, o.Payload)
}
func (o *GetMsgVpnAuthenticationOauthProviderOK) GetPayload() *models.MsgVpnAuthenticationOauthProviderResponse {
	return o.Payload
}

func (o *GetMsgVpnAuthenticationOauthProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnAuthenticationOauthProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnAuthenticationOauthProviderDefault creates a GetMsgVpnAuthenticationOauthProviderDefault with default headers values
func NewGetMsgVpnAuthenticationOauthProviderDefault(code int) *GetMsgVpnAuthenticationOauthProviderDefault {
	return &GetMsgVpnAuthenticationOauthProviderDefault{
		_statusCode: code,
	}
}

/* GetMsgVpnAuthenticationOauthProviderDefault describes a response with status code -1, with default header values.

The error response.
*/
type GetMsgVpnAuthenticationOauthProviderDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn authentication oauth provider default response
func (o *GetMsgVpnAuthenticationOauthProviderDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnAuthenticationOauthProviderDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName}][%d] getMsgVpnAuthenticationOauthProvider default  %+v", o._statusCode, o.Payload)
}
func (o *GetMsgVpnAuthenticationOauthProviderDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *GetMsgVpnAuthenticationOauthProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
