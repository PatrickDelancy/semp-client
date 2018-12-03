// Code generated by go-swagger; DO NOT EDIT.

package about

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ExalDraen/semp-client/models"
)

// GetAboutUserReader is a Reader for the GetAboutUser structure.
type GetAboutUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAboutUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAboutUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAboutUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAboutUserOK creates a GetAboutUserOK with default headers values
func NewGetAboutUserOK() *GetAboutUserOK {
	return &GetAboutUserOK{}
}

/*GetAboutUserOK handles this case with default header values.

The Current User object's attributes, and the request metadata.
*/
type GetAboutUserOK struct {
	Payload *models.AboutUserResponse
}

func (o *GetAboutUserOK) Error() string {
	return fmt.Sprintf("[GET /about/user][%d] getAboutUserOK  %+v", 200, o.Payload)
}

func (o *GetAboutUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AboutUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAboutUserDefault creates a GetAboutUserDefault with default headers values
func NewGetAboutUserDefault(code int) *GetAboutUserDefault {
	return &GetAboutUserDefault{
		_statusCode: code,
	}
}

/*GetAboutUserDefault handles this case with default header values.

Error response
*/
type GetAboutUserDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get about user default response
func (o *GetAboutUserDefault) Code() int {
	return o._statusCode
}

func (o *GetAboutUserDefault) Error() string {
	return fmt.Sprintf("[GET /about/user][%d] getAboutUser default  %+v", o._statusCode, o.Payload)
}

func (o *GetAboutUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
