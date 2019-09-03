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

// ReplaceMsgVpnDmrBridgeReader is a Reader for the ReplaceMsgVpnDmrBridge structure.
type ReplaceMsgVpnDmrBridgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceMsgVpnDmrBridgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceMsgVpnDmrBridgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewReplaceMsgVpnDmrBridgeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceMsgVpnDmrBridgeOK creates a ReplaceMsgVpnDmrBridgeOK with default headers values
func NewReplaceMsgVpnDmrBridgeOK() *ReplaceMsgVpnDmrBridgeOK {
	return &ReplaceMsgVpnDmrBridgeOK{}
}

/*ReplaceMsgVpnDmrBridgeOK handles this case with default header values.

The DMR Bridge object's attributes after being replaced, and the request metadata.
*/
type ReplaceMsgVpnDmrBridgeOK struct {
	Payload *models.MsgVpnDmrBridgeResponse
}

func (o *ReplaceMsgVpnDmrBridgeOK) Error() string {
	return fmt.Sprintf("[PUT /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName}][%d] replaceMsgVpnDmrBridgeOK  %+v", 200, o.Payload)
}

func (o *ReplaceMsgVpnDmrBridgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnDmrBridgeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceMsgVpnDmrBridgeDefault creates a ReplaceMsgVpnDmrBridgeDefault with default headers values
func NewReplaceMsgVpnDmrBridgeDefault(code int) *ReplaceMsgVpnDmrBridgeDefault {
	return &ReplaceMsgVpnDmrBridgeDefault{
		_statusCode: code,
	}
}

/*ReplaceMsgVpnDmrBridgeDefault handles this case with default header values.

The error response.
*/
type ReplaceMsgVpnDmrBridgeDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the replace msg vpn dmr bridge default response
func (o *ReplaceMsgVpnDmrBridgeDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceMsgVpnDmrBridgeDefault) Error() string {
	return fmt.Sprintf("[PUT /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName}][%d] replaceMsgVpnDmrBridge default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceMsgVpnDmrBridgeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}