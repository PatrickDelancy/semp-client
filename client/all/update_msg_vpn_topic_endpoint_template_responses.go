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

// UpdateMsgVpnTopicEndpointTemplateReader is a Reader for the UpdateMsgVpnTopicEndpointTemplate structure.
type UpdateMsgVpnTopicEndpointTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMsgVpnTopicEndpointTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMsgVpnTopicEndpointTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateMsgVpnTopicEndpointTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMsgVpnTopicEndpointTemplateOK creates a UpdateMsgVpnTopicEndpointTemplateOK with default headers values
func NewUpdateMsgVpnTopicEndpointTemplateOK() *UpdateMsgVpnTopicEndpointTemplateOK {
	return &UpdateMsgVpnTopicEndpointTemplateOK{}
}

/* UpdateMsgVpnTopicEndpointTemplateOK describes a response with status code 200, with default header values.

The Topic Endpoint Template object's attributes after being updated, and the request metadata.
*/
type UpdateMsgVpnTopicEndpointTemplateOK struct {
	Payload *models.MsgVpnTopicEndpointTemplateResponse
}

func (o *UpdateMsgVpnTopicEndpointTemplateOK) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName}][%d] updateMsgVpnTopicEndpointTemplateOK  %+v", 200, o.Payload)
}
func (o *UpdateMsgVpnTopicEndpointTemplateOK) GetPayload() *models.MsgVpnTopicEndpointTemplateResponse {
	return o.Payload
}

func (o *UpdateMsgVpnTopicEndpointTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnTopicEndpointTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMsgVpnTopicEndpointTemplateDefault creates a UpdateMsgVpnTopicEndpointTemplateDefault with default headers values
func NewUpdateMsgVpnTopicEndpointTemplateDefault(code int) *UpdateMsgVpnTopicEndpointTemplateDefault {
	return &UpdateMsgVpnTopicEndpointTemplateDefault{
		_statusCode: code,
	}
}

/* UpdateMsgVpnTopicEndpointTemplateDefault describes a response with status code -1, with default header values.

The error response.
*/
type UpdateMsgVpnTopicEndpointTemplateDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the update msg vpn topic endpoint template default response
func (o *UpdateMsgVpnTopicEndpointTemplateDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMsgVpnTopicEndpointTemplateDefault) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName}][%d] updateMsgVpnTopicEndpointTemplate default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateMsgVpnTopicEndpointTemplateDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *UpdateMsgVpnTopicEndpointTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
