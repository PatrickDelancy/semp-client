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

// CreateMsgVpnJndiConnectionFactoryReader is a Reader for the CreateMsgVpnJndiConnectionFactory structure.
type CreateMsgVpnJndiConnectionFactoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnJndiConnectionFactoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateMsgVpnJndiConnectionFactoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateMsgVpnJndiConnectionFactoryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnJndiConnectionFactoryOK creates a CreateMsgVpnJndiConnectionFactoryOK with default headers values
func NewCreateMsgVpnJndiConnectionFactoryOK() *CreateMsgVpnJndiConnectionFactoryOK {
	return &CreateMsgVpnJndiConnectionFactoryOK{}
}

/* CreateMsgVpnJndiConnectionFactoryOK describes a response with status code 200, with default header values.

The JNDI Connection Factory object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnJndiConnectionFactoryOK struct {
	Payload *models.MsgVpnJndiConnectionFactoryResponse
}

func (o *CreateMsgVpnJndiConnectionFactoryOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/jndiConnectionFactories][%d] createMsgVpnJndiConnectionFactoryOK  %+v", 200, o.Payload)
}
func (o *CreateMsgVpnJndiConnectionFactoryOK) GetPayload() *models.MsgVpnJndiConnectionFactoryResponse {
	return o.Payload
}

func (o *CreateMsgVpnJndiConnectionFactoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnJndiConnectionFactoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnJndiConnectionFactoryDefault creates a CreateMsgVpnJndiConnectionFactoryDefault with default headers values
func NewCreateMsgVpnJndiConnectionFactoryDefault(code int) *CreateMsgVpnJndiConnectionFactoryDefault {
	return &CreateMsgVpnJndiConnectionFactoryDefault{
		_statusCode: code,
	}
}

/* CreateMsgVpnJndiConnectionFactoryDefault describes a response with status code -1, with default header values.

The error response.
*/
type CreateMsgVpnJndiConnectionFactoryDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn jndi connection factory default response
func (o *CreateMsgVpnJndiConnectionFactoryDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnJndiConnectionFactoryDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/jndiConnectionFactories][%d] createMsgVpnJndiConnectionFactory default  %+v", o._statusCode, o.Payload)
}
func (o *CreateMsgVpnJndiConnectionFactoryDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *CreateMsgVpnJndiConnectionFactoryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
