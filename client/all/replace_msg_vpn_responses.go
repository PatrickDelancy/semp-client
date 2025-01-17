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

// ReplaceMsgVpnReader is a Reader for the ReplaceMsgVpn structure.
type ReplaceMsgVpnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceMsgVpnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceMsgVpnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReplaceMsgVpnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceMsgVpnOK creates a ReplaceMsgVpnOK with default headers values
func NewReplaceMsgVpnOK() *ReplaceMsgVpnOK {
	return &ReplaceMsgVpnOK{}
}

/* ReplaceMsgVpnOK describes a response with status code 200, with default header values.

The Message VPN object's attributes after being replaced, and the request metadata.
*/
type ReplaceMsgVpnOK struct {
	Payload *models.MsgVpnResponse
}

func (o *ReplaceMsgVpnOK) Error() string {
	return fmt.Sprintf("[PUT /msgVpns/{msgVpnName}][%d] replaceMsgVpnOK  %+v", 200, o.Payload)
}
func (o *ReplaceMsgVpnOK) GetPayload() *models.MsgVpnResponse {
	return o.Payload
}

func (o *ReplaceMsgVpnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceMsgVpnDefault creates a ReplaceMsgVpnDefault with default headers values
func NewReplaceMsgVpnDefault(code int) *ReplaceMsgVpnDefault {
	return &ReplaceMsgVpnDefault{
		_statusCode: code,
	}
}

/* ReplaceMsgVpnDefault describes a response with status code -1, with default header values.

The error response.
*/
type ReplaceMsgVpnDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the replace msg vpn default response
func (o *ReplaceMsgVpnDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceMsgVpnDefault) Error() string {
	return fmt.Sprintf("[PUT /msgVpns/{msgVpnName}][%d] replaceMsgVpn default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceMsgVpnDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *ReplaceMsgVpnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
