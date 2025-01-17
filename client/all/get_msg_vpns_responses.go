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

// GetMsgVpnsReader is a Reader for the GetMsgVpns structure.
type GetMsgVpnsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMsgVpnsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMsgVpnsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnsOK creates a GetMsgVpnsOK with default headers values
func NewGetMsgVpnsOK() *GetMsgVpnsOK {
	return &GetMsgVpnsOK{}
}

/* GetMsgVpnsOK describes a response with status code 200, with default header values.

The list of Message VPN objects' attributes, and the request metadata.
*/
type GetMsgVpnsOK struct {
	Payload *models.MsgVpnsResponse
}

func (o *GetMsgVpnsOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns][%d] getMsgVpnsOK  %+v", 200, o.Payload)
}
func (o *GetMsgVpnsOK) GetPayload() *models.MsgVpnsResponse {
	return o.Payload
}

func (o *GetMsgVpnsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnsDefault creates a GetMsgVpnsDefault with default headers values
func NewGetMsgVpnsDefault(code int) *GetMsgVpnsDefault {
	return &GetMsgVpnsDefault{
		_statusCode: code,
	}
}

/* GetMsgVpnsDefault describes a response with status code -1, with default header values.

The error response.
*/
type GetMsgVpnsDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpns default response
func (o *GetMsgVpnsDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnsDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns][%d] getMsgVpns default  %+v", o._statusCode, o.Payload)
}
func (o *GetMsgVpnsDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *GetMsgVpnsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
