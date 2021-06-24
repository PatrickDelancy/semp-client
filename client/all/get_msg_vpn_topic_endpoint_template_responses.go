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

// GetMsgVpnTopicEndpointTemplateReader is a Reader for the GetMsgVpnTopicEndpointTemplate structure.
type GetMsgVpnTopicEndpointTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnTopicEndpointTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMsgVpnTopicEndpointTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMsgVpnTopicEndpointTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnTopicEndpointTemplateOK creates a GetMsgVpnTopicEndpointTemplateOK with default headers values
func NewGetMsgVpnTopicEndpointTemplateOK() *GetMsgVpnTopicEndpointTemplateOK {
	return &GetMsgVpnTopicEndpointTemplateOK{}
}

/* GetMsgVpnTopicEndpointTemplateOK describes a response with status code 200, with default header values.

The Topic Endpoint Template object's attributes, and the request metadata.
*/
type GetMsgVpnTopicEndpointTemplateOK struct {
	Payload *models.MsgVpnTopicEndpointTemplateResponse
}

func (o *GetMsgVpnTopicEndpointTemplateOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName}][%d] getMsgVpnTopicEndpointTemplateOK  %+v", 200, o.Payload)
}
func (o *GetMsgVpnTopicEndpointTemplateOK) GetPayload() *models.MsgVpnTopicEndpointTemplateResponse {
	return o.Payload
}

func (o *GetMsgVpnTopicEndpointTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnTopicEndpointTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnTopicEndpointTemplateDefault creates a GetMsgVpnTopicEndpointTemplateDefault with default headers values
func NewGetMsgVpnTopicEndpointTemplateDefault(code int) *GetMsgVpnTopicEndpointTemplateDefault {
	return &GetMsgVpnTopicEndpointTemplateDefault{
		_statusCode: code,
	}
}

/* GetMsgVpnTopicEndpointTemplateDefault describes a response with status code -1, with default header values.

The error response.
*/
type GetMsgVpnTopicEndpointTemplateDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn topic endpoint template default response
func (o *GetMsgVpnTopicEndpointTemplateDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnTopicEndpointTemplateDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName}][%d] getMsgVpnTopicEndpointTemplate default  %+v", o._statusCode, o.Payload)
}
func (o *GetMsgVpnTopicEndpointTemplateDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *GetMsgVpnTopicEndpointTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
