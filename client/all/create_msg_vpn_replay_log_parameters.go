// Code generated by go-swagger; DO NOT EDIT.

package all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/PatrickDelancy/semp-client/models"
)

// NewCreateMsgVpnReplayLogParams creates a new CreateMsgVpnReplayLogParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateMsgVpnReplayLogParams() *CreateMsgVpnReplayLogParams {
	return &CreateMsgVpnReplayLogParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMsgVpnReplayLogParamsWithTimeout creates a new CreateMsgVpnReplayLogParams object
// with the ability to set a timeout on a request.
func NewCreateMsgVpnReplayLogParamsWithTimeout(timeout time.Duration) *CreateMsgVpnReplayLogParams {
	return &CreateMsgVpnReplayLogParams{
		timeout: timeout,
	}
}

// NewCreateMsgVpnReplayLogParamsWithContext creates a new CreateMsgVpnReplayLogParams object
// with the ability to set a context for a request.
func NewCreateMsgVpnReplayLogParamsWithContext(ctx context.Context) *CreateMsgVpnReplayLogParams {
	return &CreateMsgVpnReplayLogParams{
		Context: ctx,
	}
}

// NewCreateMsgVpnReplayLogParamsWithHTTPClient creates a new CreateMsgVpnReplayLogParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateMsgVpnReplayLogParamsWithHTTPClient(client *http.Client) *CreateMsgVpnReplayLogParams {
	return &CreateMsgVpnReplayLogParams{
		HTTPClient: client,
	}
}

/* CreateMsgVpnReplayLogParams contains all the parameters to send to the API endpoint
   for the create msg vpn replay log operation.

   Typically these are written to a http.Request.
*/
type CreateMsgVpnReplayLogParams struct {

	/* Body.

	   The Replay Log object's attributes.
	*/
	Body *models.MsgVpnReplayLog

	/* MsgVpnName.

	   The name of the Message VPN.
	*/
	MsgVpnName string

	/* OpaquePassword.

	   Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter.
	*/
	OpaquePassword *string

	/* Select.

	   Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.
	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create msg vpn replay log params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMsgVpnReplayLogParams) WithDefaults() *CreateMsgVpnReplayLogParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create msg vpn replay log params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMsgVpnReplayLogParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create msg vpn replay log params
func (o *CreateMsgVpnReplayLogParams) WithTimeout(timeout time.Duration) *CreateMsgVpnReplayLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create msg vpn replay log params
func (o *CreateMsgVpnReplayLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create msg vpn replay log params
func (o *CreateMsgVpnReplayLogParams) WithContext(ctx context.Context) *CreateMsgVpnReplayLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create msg vpn replay log params
func (o *CreateMsgVpnReplayLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create msg vpn replay log params
func (o *CreateMsgVpnReplayLogParams) WithHTTPClient(client *http.Client) *CreateMsgVpnReplayLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create msg vpn replay log params
func (o *CreateMsgVpnReplayLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create msg vpn replay log params
func (o *CreateMsgVpnReplayLogParams) WithBody(body *models.MsgVpnReplayLog) *CreateMsgVpnReplayLogParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create msg vpn replay log params
func (o *CreateMsgVpnReplayLogParams) SetBody(body *models.MsgVpnReplayLog) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the create msg vpn replay log params
func (o *CreateMsgVpnReplayLogParams) WithMsgVpnName(msgVpnName string) *CreateMsgVpnReplayLogParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the create msg vpn replay log params
func (o *CreateMsgVpnReplayLogParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithOpaquePassword adds the opaquePassword to the create msg vpn replay log params
func (o *CreateMsgVpnReplayLogParams) WithOpaquePassword(opaquePassword *string) *CreateMsgVpnReplayLogParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the create msg vpn replay log params
func (o *CreateMsgVpnReplayLogParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the create msg vpn replay log params
func (o *CreateMsgVpnReplayLogParams) WithSelect(selectVar []string) *CreateMsgVpnReplayLogParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the create msg vpn replay log params
func (o *CreateMsgVpnReplayLogParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMsgVpnReplayLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.OpaquePassword != nil {

		// query param opaquePassword
		var qrOpaquePassword string

		if o.OpaquePassword != nil {
			qrOpaquePassword = *o.OpaquePassword
		}
		qOpaquePassword := qrOpaquePassword
		if qOpaquePassword != "" {

			if err := r.SetQueryParam("opaquePassword", qOpaquePassword); err != nil {
				return err
			}
		}
	}

	if o.Select != nil {

		// binding items for select
		joinedSelect := o.bindParamSelect(reg)

		// query array param select
		if err := r.SetQueryParam("select", joinedSelect...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCreateMsgVpnReplayLog binds the parameter select
func (o *CreateMsgVpnReplayLogParams) bindParamSelect(formats strfmt.Registry) []string {
	selectIR := o.Select

	var selectIC []string
	for _, selectIIR := range selectIR { // explode []string

		selectIIV := selectIIR // string as string
		selectIC = append(selectIC, selectIIV)
	}

	// items.CollectionFormat: "csv"
	selectIS := swag.JoinByFormat(selectIC, "csv")

	return selectIS
}
