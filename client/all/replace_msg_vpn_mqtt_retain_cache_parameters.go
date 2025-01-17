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

// NewReplaceMsgVpnMqttRetainCacheParams creates a new ReplaceMsgVpnMqttRetainCacheParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceMsgVpnMqttRetainCacheParams() *ReplaceMsgVpnMqttRetainCacheParams {
	return &ReplaceMsgVpnMqttRetainCacheParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceMsgVpnMqttRetainCacheParamsWithTimeout creates a new ReplaceMsgVpnMqttRetainCacheParams object
// with the ability to set a timeout on a request.
func NewReplaceMsgVpnMqttRetainCacheParamsWithTimeout(timeout time.Duration) *ReplaceMsgVpnMqttRetainCacheParams {
	return &ReplaceMsgVpnMqttRetainCacheParams{
		timeout: timeout,
	}
}

// NewReplaceMsgVpnMqttRetainCacheParamsWithContext creates a new ReplaceMsgVpnMqttRetainCacheParams object
// with the ability to set a context for a request.
func NewReplaceMsgVpnMqttRetainCacheParamsWithContext(ctx context.Context) *ReplaceMsgVpnMqttRetainCacheParams {
	return &ReplaceMsgVpnMqttRetainCacheParams{
		Context: ctx,
	}
}

// NewReplaceMsgVpnMqttRetainCacheParamsWithHTTPClient creates a new ReplaceMsgVpnMqttRetainCacheParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceMsgVpnMqttRetainCacheParamsWithHTTPClient(client *http.Client) *ReplaceMsgVpnMqttRetainCacheParams {
	return &ReplaceMsgVpnMqttRetainCacheParams{
		HTTPClient: client,
	}
}

/* ReplaceMsgVpnMqttRetainCacheParams contains all the parameters to send to the API endpoint
   for the replace msg vpn mqtt retain cache operation.

   Typically these are written to a http.Request.
*/
type ReplaceMsgVpnMqttRetainCacheParams struct {

	/* Body.

	   The MQTT Retain Cache object's attributes.
	*/
	Body *models.MsgVpnMqttRetainCache

	/* CacheName.

	   The name of the MQTT Retain Cache.
	*/
	CacheName string

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

// WithDefaults hydrates default values in the replace msg vpn mqtt retain cache params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceMsgVpnMqttRetainCacheParams) WithDefaults() *ReplaceMsgVpnMqttRetainCacheParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace msg vpn mqtt retain cache params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceMsgVpnMqttRetainCacheParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the replace msg vpn mqtt retain cache params
func (o *ReplaceMsgVpnMqttRetainCacheParams) WithTimeout(timeout time.Duration) *ReplaceMsgVpnMqttRetainCacheParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace msg vpn mqtt retain cache params
func (o *ReplaceMsgVpnMqttRetainCacheParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace msg vpn mqtt retain cache params
func (o *ReplaceMsgVpnMqttRetainCacheParams) WithContext(ctx context.Context) *ReplaceMsgVpnMqttRetainCacheParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace msg vpn mqtt retain cache params
func (o *ReplaceMsgVpnMqttRetainCacheParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace msg vpn mqtt retain cache params
func (o *ReplaceMsgVpnMqttRetainCacheParams) WithHTTPClient(client *http.Client) *ReplaceMsgVpnMqttRetainCacheParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace msg vpn mqtt retain cache params
func (o *ReplaceMsgVpnMqttRetainCacheParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace msg vpn mqtt retain cache params
func (o *ReplaceMsgVpnMqttRetainCacheParams) WithBody(body *models.MsgVpnMqttRetainCache) *ReplaceMsgVpnMqttRetainCacheParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace msg vpn mqtt retain cache params
func (o *ReplaceMsgVpnMqttRetainCacheParams) SetBody(body *models.MsgVpnMqttRetainCache) {
	o.Body = body
}

// WithCacheName adds the cacheName to the replace msg vpn mqtt retain cache params
func (o *ReplaceMsgVpnMqttRetainCacheParams) WithCacheName(cacheName string) *ReplaceMsgVpnMqttRetainCacheParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the replace msg vpn mqtt retain cache params
func (o *ReplaceMsgVpnMqttRetainCacheParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithMsgVpnName adds the msgVpnName to the replace msg vpn mqtt retain cache params
func (o *ReplaceMsgVpnMqttRetainCacheParams) WithMsgVpnName(msgVpnName string) *ReplaceMsgVpnMqttRetainCacheParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the replace msg vpn mqtt retain cache params
func (o *ReplaceMsgVpnMqttRetainCacheParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithOpaquePassword adds the opaquePassword to the replace msg vpn mqtt retain cache params
func (o *ReplaceMsgVpnMqttRetainCacheParams) WithOpaquePassword(opaquePassword *string) *ReplaceMsgVpnMqttRetainCacheParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the replace msg vpn mqtt retain cache params
func (o *ReplaceMsgVpnMqttRetainCacheParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the replace msg vpn mqtt retain cache params
func (o *ReplaceMsgVpnMqttRetainCacheParams) WithSelect(selectVar []string) *ReplaceMsgVpnMqttRetainCacheParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the replace msg vpn mqtt retain cache params
func (o *ReplaceMsgVpnMqttRetainCacheParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceMsgVpnMqttRetainCacheParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cacheName
	if err := r.SetPathParam("cacheName", o.CacheName); err != nil {
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

// bindParamReplaceMsgVpnMqttRetainCache binds the parameter select
func (o *ReplaceMsgVpnMqttRetainCacheParams) bindParamSelect(formats strfmt.Registry) []string {
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
