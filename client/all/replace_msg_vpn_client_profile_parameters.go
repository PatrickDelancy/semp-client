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

// NewReplaceMsgVpnClientProfileParams creates a new ReplaceMsgVpnClientProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceMsgVpnClientProfileParams() *ReplaceMsgVpnClientProfileParams {
	return &ReplaceMsgVpnClientProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceMsgVpnClientProfileParamsWithTimeout creates a new ReplaceMsgVpnClientProfileParams object
// with the ability to set a timeout on a request.
func NewReplaceMsgVpnClientProfileParamsWithTimeout(timeout time.Duration) *ReplaceMsgVpnClientProfileParams {
	return &ReplaceMsgVpnClientProfileParams{
		timeout: timeout,
	}
}

// NewReplaceMsgVpnClientProfileParamsWithContext creates a new ReplaceMsgVpnClientProfileParams object
// with the ability to set a context for a request.
func NewReplaceMsgVpnClientProfileParamsWithContext(ctx context.Context) *ReplaceMsgVpnClientProfileParams {
	return &ReplaceMsgVpnClientProfileParams{
		Context: ctx,
	}
}

// NewReplaceMsgVpnClientProfileParamsWithHTTPClient creates a new ReplaceMsgVpnClientProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceMsgVpnClientProfileParamsWithHTTPClient(client *http.Client) *ReplaceMsgVpnClientProfileParams {
	return &ReplaceMsgVpnClientProfileParams{
		HTTPClient: client,
	}
}

/* ReplaceMsgVpnClientProfileParams contains all the parameters to send to the API endpoint
   for the replace msg vpn client profile operation.

   Typically these are written to a http.Request.
*/
type ReplaceMsgVpnClientProfileParams struct {

	/* Body.

	   The Client Profile object's attributes.
	*/
	Body *models.MsgVpnClientProfile

	/* ClientProfileName.

	   The name of the Client Profile.
	*/
	ClientProfileName string

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

// WithDefaults hydrates default values in the replace msg vpn client profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceMsgVpnClientProfileParams) WithDefaults() *ReplaceMsgVpnClientProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace msg vpn client profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceMsgVpnClientProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) WithTimeout(timeout time.Duration) *ReplaceMsgVpnClientProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) WithContext(ctx context.Context) *ReplaceMsgVpnClientProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) WithHTTPClient(client *http.Client) *ReplaceMsgVpnClientProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) WithBody(body *models.MsgVpnClientProfile) *ReplaceMsgVpnClientProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) SetBody(body *models.MsgVpnClientProfile) {
	o.Body = body
}

// WithClientProfileName adds the clientProfileName to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) WithClientProfileName(clientProfileName string) *ReplaceMsgVpnClientProfileParams {
	o.SetClientProfileName(clientProfileName)
	return o
}

// SetClientProfileName adds the clientProfileName to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) SetClientProfileName(clientProfileName string) {
	o.ClientProfileName = clientProfileName
}

// WithMsgVpnName adds the msgVpnName to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) WithMsgVpnName(msgVpnName string) *ReplaceMsgVpnClientProfileParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithOpaquePassword adds the opaquePassword to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) WithOpaquePassword(opaquePassword *string) *ReplaceMsgVpnClientProfileParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) WithSelect(selectVar []string) *ReplaceMsgVpnClientProfileParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceMsgVpnClientProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param clientProfileName
	if err := r.SetPathParam("clientProfileName", o.ClientProfileName); err != nil {
		return err
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

// bindParamReplaceMsgVpnClientProfile binds the parameter select
func (o *ReplaceMsgVpnClientProfileParams) bindParamSelect(formats strfmt.Registry) []string {
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
