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

// DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimReader is a Reader for the DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim structure.
type DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOK creates a DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOK with default headers values
func NewDeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOK() *DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOK {
	return &DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOK{}
}

/* DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOK describes a response with status code 200, with default header values.

The request metadata.
*/
type DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOK) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims/{oauthJwtClaimName}][%d] deleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOK  %+v", 200, o.Payload)
}
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOK) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimDefault creates a DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimDefault with default headers values
func NewDeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimDefault(code int) *DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimDefault {
	return &DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimDefault{
		_statusCode: code,
	}
}

/* DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimDefault describes a response with status code -1, with default header values.

The error response.
*/
type DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete msg vpn rest delivery point rest consumer oauth jwt claim default response
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimDefault) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims/{oauthJwtClaimName}][%d] deleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
