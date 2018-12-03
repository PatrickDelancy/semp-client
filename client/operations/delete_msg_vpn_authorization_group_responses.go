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

// DeleteMsgVpnAuthorizationGroupReader is a Reader for the DeleteMsgVpnAuthorizationGroup structure.
type DeleteMsgVpnAuthorizationGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMsgVpnAuthorizationGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMsgVpnAuthorizationGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteMsgVpnAuthorizationGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMsgVpnAuthorizationGroupOK creates a DeleteMsgVpnAuthorizationGroupOK with default headers values
func NewDeleteMsgVpnAuthorizationGroupOK() *DeleteMsgVpnAuthorizationGroupOK {
	return &DeleteMsgVpnAuthorizationGroupOK{}
}

/*DeleteMsgVpnAuthorizationGroupOK handles this case with default header values.

The request metadata.
*/
type DeleteMsgVpnAuthorizationGroupOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteMsgVpnAuthorizationGroupOK) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName}][%d] deleteMsgVpnAuthorizationGroupOK  %+v", 200, o.Payload)
}

func (o *DeleteMsgVpnAuthorizationGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMsgVpnAuthorizationGroupDefault creates a DeleteMsgVpnAuthorizationGroupDefault with default headers values
func NewDeleteMsgVpnAuthorizationGroupDefault(code int) *DeleteMsgVpnAuthorizationGroupDefault {
	return &DeleteMsgVpnAuthorizationGroupDefault{
		_statusCode: code,
	}
}

/*DeleteMsgVpnAuthorizationGroupDefault handles this case with default header values.

Error response
*/
type DeleteMsgVpnAuthorizationGroupDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete msg vpn authorization group default response
func (o *DeleteMsgVpnAuthorizationGroupDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMsgVpnAuthorizationGroupDefault) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName}][%d] deleteMsgVpnAuthorizationGroup default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMsgVpnAuthorizationGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}