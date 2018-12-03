// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ExalDraen/semp-client/models"
)

// NewCreateMsgVpnRestDeliveryPointRestConsumerParams creates a new CreateMsgVpnRestDeliveryPointRestConsumerParams object
// with the default values initialized.
func NewCreateMsgVpnRestDeliveryPointRestConsumerParams() *CreateMsgVpnRestDeliveryPointRestConsumerParams {
	var ()
	return &CreateMsgVpnRestDeliveryPointRestConsumerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMsgVpnRestDeliveryPointRestConsumerParamsWithTimeout creates a new CreateMsgVpnRestDeliveryPointRestConsumerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMsgVpnRestDeliveryPointRestConsumerParamsWithTimeout(timeout time.Duration) *CreateMsgVpnRestDeliveryPointRestConsumerParams {
	var ()
	return &CreateMsgVpnRestDeliveryPointRestConsumerParams{

		timeout: timeout,
	}
}

// NewCreateMsgVpnRestDeliveryPointRestConsumerParamsWithContext creates a new CreateMsgVpnRestDeliveryPointRestConsumerParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMsgVpnRestDeliveryPointRestConsumerParamsWithContext(ctx context.Context) *CreateMsgVpnRestDeliveryPointRestConsumerParams {
	var ()
	return &CreateMsgVpnRestDeliveryPointRestConsumerParams{

		Context: ctx,
	}
}

// NewCreateMsgVpnRestDeliveryPointRestConsumerParamsWithHTTPClient creates a new CreateMsgVpnRestDeliveryPointRestConsumerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMsgVpnRestDeliveryPointRestConsumerParamsWithHTTPClient(client *http.Client) *CreateMsgVpnRestDeliveryPointRestConsumerParams {
	var ()
	return &CreateMsgVpnRestDeliveryPointRestConsumerParams{
		HTTPClient: client,
	}
}

/*CreateMsgVpnRestDeliveryPointRestConsumerParams contains all the parameters to send to the API endpoint
for the create msg vpn rest delivery point rest consumer operation typically these are written to a http.Request
*/
type CreateMsgVpnRestDeliveryPointRestConsumerParams struct {

	/*Body
	  The REST Consumer object's attributes.

	*/
	Body *models.MsgVpnRestDeliveryPointRestConsumer
	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string
	/*RestDeliveryPointName
	  The restDeliveryPointName of the REST Delivery Point.

	*/
	RestDeliveryPointName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See [Select](#select "Description of the syntax of the `select` parameter").

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create msg vpn rest delivery point rest consumer params
func (o *CreateMsgVpnRestDeliveryPointRestConsumerParams) WithTimeout(timeout time.Duration) *CreateMsgVpnRestDeliveryPointRestConsumerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create msg vpn rest delivery point rest consumer params
func (o *CreateMsgVpnRestDeliveryPointRestConsumerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create msg vpn rest delivery point rest consumer params
func (o *CreateMsgVpnRestDeliveryPointRestConsumerParams) WithContext(ctx context.Context) *CreateMsgVpnRestDeliveryPointRestConsumerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create msg vpn rest delivery point rest consumer params
func (o *CreateMsgVpnRestDeliveryPointRestConsumerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create msg vpn rest delivery point rest consumer params
func (o *CreateMsgVpnRestDeliveryPointRestConsumerParams) WithHTTPClient(client *http.Client) *CreateMsgVpnRestDeliveryPointRestConsumerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create msg vpn rest delivery point rest consumer params
func (o *CreateMsgVpnRestDeliveryPointRestConsumerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create msg vpn rest delivery point rest consumer params
func (o *CreateMsgVpnRestDeliveryPointRestConsumerParams) WithBody(body *models.MsgVpnRestDeliveryPointRestConsumer) *CreateMsgVpnRestDeliveryPointRestConsumerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create msg vpn rest delivery point rest consumer params
func (o *CreateMsgVpnRestDeliveryPointRestConsumerParams) SetBody(body *models.MsgVpnRestDeliveryPointRestConsumer) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the create msg vpn rest delivery point rest consumer params
func (o *CreateMsgVpnRestDeliveryPointRestConsumerParams) WithMsgVpnName(msgVpnName string) *CreateMsgVpnRestDeliveryPointRestConsumerParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the create msg vpn rest delivery point rest consumer params
func (o *CreateMsgVpnRestDeliveryPointRestConsumerParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithRestDeliveryPointName adds the restDeliveryPointName to the create msg vpn rest delivery point rest consumer params
func (o *CreateMsgVpnRestDeliveryPointRestConsumerParams) WithRestDeliveryPointName(restDeliveryPointName string) *CreateMsgVpnRestDeliveryPointRestConsumerParams {
	o.SetRestDeliveryPointName(restDeliveryPointName)
	return o
}

// SetRestDeliveryPointName adds the restDeliveryPointName to the create msg vpn rest delivery point rest consumer params
func (o *CreateMsgVpnRestDeliveryPointRestConsumerParams) SetRestDeliveryPointName(restDeliveryPointName string) {
	o.RestDeliveryPointName = restDeliveryPointName
}

// WithSelect adds the selectVar to the create msg vpn rest delivery point rest consumer params
func (o *CreateMsgVpnRestDeliveryPointRestConsumerParams) WithSelect(selectVar []string) *CreateMsgVpnRestDeliveryPointRestConsumerParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the create msg vpn rest delivery point rest consumer params
func (o *CreateMsgVpnRestDeliveryPointRestConsumerParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMsgVpnRestDeliveryPointRestConsumerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	// path param restDeliveryPointName
	if err := r.SetPathParam("restDeliveryPointName", o.RestDeliveryPointName); err != nil {
		return err
	}

	valuesSelect := o.Select

	joinedSelect := swag.JoinByFormat(valuesSelect, "csv")
	// query array param select
	if err := r.SetQueryParam("select", joinedSelect...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}