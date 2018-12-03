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

// ReplaceMsgVpnTopicEndpointReader is a Reader for the ReplaceMsgVpnTopicEndpoint structure.
type ReplaceMsgVpnTopicEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceMsgVpnTopicEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceMsgVpnTopicEndpointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewReplaceMsgVpnTopicEndpointDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceMsgVpnTopicEndpointOK creates a ReplaceMsgVpnTopicEndpointOK with default headers values
func NewReplaceMsgVpnTopicEndpointOK() *ReplaceMsgVpnTopicEndpointOK {
	return &ReplaceMsgVpnTopicEndpointOK{}
}

/*ReplaceMsgVpnTopicEndpointOK handles this case with default header values.

The Topic Endpoint object's attributes after being replaced, and the request metadata.
*/
type ReplaceMsgVpnTopicEndpointOK struct {
	Payload *models.MsgVpnTopicEndpointResponse
}

func (o *ReplaceMsgVpnTopicEndpointOK) Error() string {
	return fmt.Sprintf("[PUT /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}][%d] replaceMsgVpnTopicEndpointOK  %+v", 200, o.Payload)
}

func (o *ReplaceMsgVpnTopicEndpointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnTopicEndpointResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceMsgVpnTopicEndpointDefault creates a ReplaceMsgVpnTopicEndpointDefault with default headers values
func NewReplaceMsgVpnTopicEndpointDefault(code int) *ReplaceMsgVpnTopicEndpointDefault {
	return &ReplaceMsgVpnTopicEndpointDefault{
		_statusCode: code,
	}
}

/*ReplaceMsgVpnTopicEndpointDefault handles this case with default header values.

Error response
*/
type ReplaceMsgVpnTopicEndpointDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the replace msg vpn topic endpoint default response
func (o *ReplaceMsgVpnTopicEndpointDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceMsgVpnTopicEndpointDefault) Error() string {
	return fmt.Sprintf("[PUT /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}][%d] replaceMsgVpnTopicEndpoint default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceMsgVpnTopicEndpointDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
