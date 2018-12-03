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

// UpdateMsgVpnRestDeliveryPointReader is a Reader for the UpdateMsgVpnRestDeliveryPoint structure.
type UpdateMsgVpnRestDeliveryPointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMsgVpnRestDeliveryPointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateMsgVpnRestDeliveryPointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateMsgVpnRestDeliveryPointDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMsgVpnRestDeliveryPointOK creates a UpdateMsgVpnRestDeliveryPointOK with default headers values
func NewUpdateMsgVpnRestDeliveryPointOK() *UpdateMsgVpnRestDeliveryPointOK {
	return &UpdateMsgVpnRestDeliveryPointOK{}
}

/*UpdateMsgVpnRestDeliveryPointOK handles this case with default header values.

The REST Delivery Point object's attributes after being updated, and the request metadata.
*/
type UpdateMsgVpnRestDeliveryPointOK struct {
	Payload *models.MsgVpnRestDeliveryPointResponse
}

func (o *UpdateMsgVpnRestDeliveryPointOK) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}][%d] updateMsgVpnRestDeliveryPointOK  %+v", 200, o.Payload)
}

func (o *UpdateMsgVpnRestDeliveryPointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnRestDeliveryPointResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMsgVpnRestDeliveryPointDefault creates a UpdateMsgVpnRestDeliveryPointDefault with default headers values
func NewUpdateMsgVpnRestDeliveryPointDefault(code int) *UpdateMsgVpnRestDeliveryPointDefault {
	return &UpdateMsgVpnRestDeliveryPointDefault{
		_statusCode: code,
	}
}

/*UpdateMsgVpnRestDeliveryPointDefault handles this case with default header values.

Error response
*/
type UpdateMsgVpnRestDeliveryPointDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the update msg vpn rest delivery point default response
func (o *UpdateMsgVpnRestDeliveryPointDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMsgVpnRestDeliveryPointDefault) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}][%d] updateMsgVpnRestDeliveryPoint default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateMsgVpnRestDeliveryPointDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
