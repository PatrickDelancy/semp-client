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

// ReplaceMsgVpnJndiTopicReader is a Reader for the ReplaceMsgVpnJndiTopic structure.
type ReplaceMsgVpnJndiTopicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceMsgVpnJndiTopicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceMsgVpnJndiTopicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewReplaceMsgVpnJndiTopicDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceMsgVpnJndiTopicOK creates a ReplaceMsgVpnJndiTopicOK with default headers values
func NewReplaceMsgVpnJndiTopicOK() *ReplaceMsgVpnJndiTopicOK {
	return &ReplaceMsgVpnJndiTopicOK{}
}

/*ReplaceMsgVpnJndiTopicOK handles this case with default header values.

The JNDI Topic object's attributes after being replaced, and the request metadata.
*/
type ReplaceMsgVpnJndiTopicOK struct {
	Payload *models.MsgVpnJndiTopicResponse
}

func (o *ReplaceMsgVpnJndiTopicOK) Error() string {
	return fmt.Sprintf("[PUT /msgVpns/{msgVpnName}/jndiTopics/{topicName}][%d] replaceMsgVpnJndiTopicOK  %+v", 200, o.Payload)
}

func (o *ReplaceMsgVpnJndiTopicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnJndiTopicResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceMsgVpnJndiTopicDefault creates a ReplaceMsgVpnJndiTopicDefault with default headers values
func NewReplaceMsgVpnJndiTopicDefault(code int) *ReplaceMsgVpnJndiTopicDefault {
	return &ReplaceMsgVpnJndiTopicDefault{
		_statusCode: code,
	}
}

/*ReplaceMsgVpnJndiTopicDefault handles this case with default header values.

Error response
*/
type ReplaceMsgVpnJndiTopicDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the replace msg vpn jndi topic default response
func (o *ReplaceMsgVpnJndiTopicDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceMsgVpnJndiTopicDefault) Error() string {
	return fmt.Sprintf("[PUT /msgVpns/{msgVpnName}/jndiTopics/{topicName}][%d] replaceMsgVpnJndiTopic default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceMsgVpnJndiTopicDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
