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
)

// NewGetMsgVpnSequencedTopicParams creates a new GetMsgVpnSequencedTopicParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMsgVpnSequencedTopicParams() *GetMsgVpnSequencedTopicParams {
	return &GetMsgVpnSequencedTopicParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnSequencedTopicParamsWithTimeout creates a new GetMsgVpnSequencedTopicParams object
// with the ability to set a timeout on a request.
func NewGetMsgVpnSequencedTopicParamsWithTimeout(timeout time.Duration) *GetMsgVpnSequencedTopicParams {
	return &GetMsgVpnSequencedTopicParams{
		timeout: timeout,
	}
}

// NewGetMsgVpnSequencedTopicParamsWithContext creates a new GetMsgVpnSequencedTopicParams object
// with the ability to set a context for a request.
func NewGetMsgVpnSequencedTopicParamsWithContext(ctx context.Context) *GetMsgVpnSequencedTopicParams {
	return &GetMsgVpnSequencedTopicParams{
		Context: ctx,
	}
}

// NewGetMsgVpnSequencedTopicParamsWithHTTPClient creates a new GetMsgVpnSequencedTopicParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMsgVpnSequencedTopicParamsWithHTTPClient(client *http.Client) *GetMsgVpnSequencedTopicParams {
	return &GetMsgVpnSequencedTopicParams{
		HTTPClient: client,
	}
}

/* GetMsgVpnSequencedTopicParams contains all the parameters to send to the API endpoint
   for the get msg vpn sequenced topic operation.

   Typically these are written to a http.Request.
*/
type GetMsgVpnSequencedTopicParams struct {

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

	/* SequencedTopic.

	   Topic for applying sequence numbers.
	*/
	SequencedTopic string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get msg vpn sequenced topic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMsgVpnSequencedTopicParams) WithDefaults() *GetMsgVpnSequencedTopicParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get msg vpn sequenced topic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMsgVpnSequencedTopicParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get msg vpn sequenced topic params
func (o *GetMsgVpnSequencedTopicParams) WithTimeout(timeout time.Duration) *GetMsgVpnSequencedTopicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn sequenced topic params
func (o *GetMsgVpnSequencedTopicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn sequenced topic params
func (o *GetMsgVpnSequencedTopicParams) WithContext(ctx context.Context) *GetMsgVpnSequencedTopicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn sequenced topic params
func (o *GetMsgVpnSequencedTopicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn sequenced topic params
func (o *GetMsgVpnSequencedTopicParams) WithHTTPClient(client *http.Client) *GetMsgVpnSequencedTopicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn sequenced topic params
func (o *GetMsgVpnSequencedTopicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn sequenced topic params
func (o *GetMsgVpnSequencedTopicParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnSequencedTopicParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn sequenced topic params
func (o *GetMsgVpnSequencedTopicParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithOpaquePassword adds the opaquePassword to the get msg vpn sequenced topic params
func (o *GetMsgVpnSequencedTopicParams) WithOpaquePassword(opaquePassword *string) *GetMsgVpnSequencedTopicParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the get msg vpn sequenced topic params
func (o *GetMsgVpnSequencedTopicParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the get msg vpn sequenced topic params
func (o *GetMsgVpnSequencedTopicParams) WithSelect(selectVar []string) *GetMsgVpnSequencedTopicParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn sequenced topic params
func (o *GetMsgVpnSequencedTopicParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithSequencedTopic adds the sequencedTopic to the get msg vpn sequenced topic params
func (o *GetMsgVpnSequencedTopicParams) WithSequencedTopic(sequencedTopic string) *GetMsgVpnSequencedTopicParams {
	o.SetSequencedTopic(sequencedTopic)
	return o
}

// SetSequencedTopic adds the sequencedTopic to the get msg vpn sequenced topic params
func (o *GetMsgVpnSequencedTopicParams) SetSequencedTopic(sequencedTopic string) {
	o.SequencedTopic = sequencedTopic
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnSequencedTopicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param sequencedTopic
	if err := r.SetPathParam("sequencedTopic", o.SequencedTopic); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetMsgVpnSequencedTopic binds the parameter select
func (o *GetMsgVpnSequencedTopicParams) bindParamSelect(formats strfmt.Registry) []string {
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