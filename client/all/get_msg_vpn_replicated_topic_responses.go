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

// GetMsgVpnReplicatedTopicReader is a Reader for the GetMsgVpnReplicatedTopic structure.
type GetMsgVpnReplicatedTopicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnReplicatedTopicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMsgVpnReplicatedTopicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMsgVpnReplicatedTopicDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnReplicatedTopicOK creates a GetMsgVpnReplicatedTopicOK with default headers values
func NewGetMsgVpnReplicatedTopicOK() *GetMsgVpnReplicatedTopicOK {
	return &GetMsgVpnReplicatedTopicOK{}
}

/* GetMsgVpnReplicatedTopicOK describes a response with status code 200, with default header values.

The Replicated Topic object's attributes, and the request metadata.
*/
type GetMsgVpnReplicatedTopicOK struct {
	Payload *models.MsgVpnReplicatedTopicResponse
}

func (o *GetMsgVpnReplicatedTopicOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic}][%d] getMsgVpnReplicatedTopicOK  %+v", 200, o.Payload)
}
func (o *GetMsgVpnReplicatedTopicOK) GetPayload() *models.MsgVpnReplicatedTopicResponse {
	return o.Payload
}

func (o *GetMsgVpnReplicatedTopicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnReplicatedTopicResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnReplicatedTopicDefault creates a GetMsgVpnReplicatedTopicDefault with default headers values
func NewGetMsgVpnReplicatedTopicDefault(code int) *GetMsgVpnReplicatedTopicDefault {
	return &GetMsgVpnReplicatedTopicDefault{
		_statusCode: code,
	}
}

/* GetMsgVpnReplicatedTopicDefault describes a response with status code -1, with default header values.

The error response.
*/
type GetMsgVpnReplicatedTopicDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn replicated topic default response
func (o *GetMsgVpnReplicatedTopicDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnReplicatedTopicDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic}][%d] getMsgVpnReplicatedTopic default  %+v", o._statusCode, o.Payload)
}
func (o *GetMsgVpnReplicatedTopicDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *GetMsgVpnReplicatedTopicDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
