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

// NewReplaceMsgVpnAuthenticationOauthProviderParams creates a new ReplaceMsgVpnAuthenticationOauthProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceMsgVpnAuthenticationOauthProviderParams() *ReplaceMsgVpnAuthenticationOauthProviderParams {
	return &ReplaceMsgVpnAuthenticationOauthProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceMsgVpnAuthenticationOauthProviderParamsWithTimeout creates a new ReplaceMsgVpnAuthenticationOauthProviderParams object
// with the ability to set a timeout on a request.
func NewReplaceMsgVpnAuthenticationOauthProviderParamsWithTimeout(timeout time.Duration) *ReplaceMsgVpnAuthenticationOauthProviderParams {
	return &ReplaceMsgVpnAuthenticationOauthProviderParams{
		timeout: timeout,
	}
}

// NewReplaceMsgVpnAuthenticationOauthProviderParamsWithContext creates a new ReplaceMsgVpnAuthenticationOauthProviderParams object
// with the ability to set a context for a request.
func NewReplaceMsgVpnAuthenticationOauthProviderParamsWithContext(ctx context.Context) *ReplaceMsgVpnAuthenticationOauthProviderParams {
	return &ReplaceMsgVpnAuthenticationOauthProviderParams{
		Context: ctx,
	}
}

// NewReplaceMsgVpnAuthenticationOauthProviderParamsWithHTTPClient creates a new ReplaceMsgVpnAuthenticationOauthProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceMsgVpnAuthenticationOauthProviderParamsWithHTTPClient(client *http.Client) *ReplaceMsgVpnAuthenticationOauthProviderParams {
	return &ReplaceMsgVpnAuthenticationOauthProviderParams{
		HTTPClient: client,
	}
}

/* ReplaceMsgVpnAuthenticationOauthProviderParams contains all the parameters to send to the API endpoint
   for the replace msg vpn authentication oauth provider operation.

   Typically these are written to a http.Request.
*/
type ReplaceMsgVpnAuthenticationOauthProviderParams struct {

	/* Body.

	   The OAuth Provider object's attributes.
	*/
	Body *models.MsgVpnAuthenticationOauthProvider

	/* MsgVpnName.

	   The name of the Message VPN.
	*/
	MsgVpnName string

	/* OauthProviderName.

	   The name of the OAuth Provider.
	*/
	OauthProviderName string

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

// WithDefaults hydrates default values in the replace msg vpn authentication oauth provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) WithDefaults() *ReplaceMsgVpnAuthenticationOauthProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace msg vpn authentication oauth provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the replace msg vpn authentication oauth provider params
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) WithTimeout(timeout time.Duration) *ReplaceMsgVpnAuthenticationOauthProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace msg vpn authentication oauth provider params
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace msg vpn authentication oauth provider params
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) WithContext(ctx context.Context) *ReplaceMsgVpnAuthenticationOauthProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace msg vpn authentication oauth provider params
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace msg vpn authentication oauth provider params
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) WithHTTPClient(client *http.Client) *ReplaceMsgVpnAuthenticationOauthProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace msg vpn authentication oauth provider params
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace msg vpn authentication oauth provider params
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) WithBody(body *models.MsgVpnAuthenticationOauthProvider) *ReplaceMsgVpnAuthenticationOauthProviderParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace msg vpn authentication oauth provider params
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) SetBody(body *models.MsgVpnAuthenticationOauthProvider) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the replace msg vpn authentication oauth provider params
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) WithMsgVpnName(msgVpnName string) *ReplaceMsgVpnAuthenticationOauthProviderParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the replace msg vpn authentication oauth provider params
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithOauthProviderName adds the oauthProviderName to the replace msg vpn authentication oauth provider params
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) WithOauthProviderName(oauthProviderName string) *ReplaceMsgVpnAuthenticationOauthProviderParams {
	o.SetOauthProviderName(oauthProviderName)
	return o
}

// SetOauthProviderName adds the oauthProviderName to the replace msg vpn authentication oauth provider params
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) SetOauthProviderName(oauthProviderName string) {
	o.OauthProviderName = oauthProviderName
}

// WithOpaquePassword adds the opaquePassword to the replace msg vpn authentication oauth provider params
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) WithOpaquePassword(opaquePassword *string) *ReplaceMsgVpnAuthenticationOauthProviderParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the replace msg vpn authentication oauth provider params
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the replace msg vpn authentication oauth provider params
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) WithSelect(selectVar []string) *ReplaceMsgVpnAuthenticationOauthProviderParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the replace msg vpn authentication oauth provider params
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param oauthProviderName
	if err := r.SetPathParam("oauthProviderName", o.OauthProviderName); err != nil {
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

// bindParamReplaceMsgVpnAuthenticationOauthProvider binds the parameter select
func (o *ReplaceMsgVpnAuthenticationOauthProviderParams) bindParamSelect(formats strfmt.Registry) []string {
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
