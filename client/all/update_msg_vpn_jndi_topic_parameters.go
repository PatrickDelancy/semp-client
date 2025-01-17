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

// NewUpdateMsgVpnJndiTopicParams creates a new UpdateMsgVpnJndiTopicParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMsgVpnJndiTopicParams() *UpdateMsgVpnJndiTopicParams {
	return &UpdateMsgVpnJndiTopicParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMsgVpnJndiTopicParamsWithTimeout creates a new UpdateMsgVpnJndiTopicParams object
// with the ability to set a timeout on a request.
func NewUpdateMsgVpnJndiTopicParamsWithTimeout(timeout time.Duration) *UpdateMsgVpnJndiTopicParams {
	return &UpdateMsgVpnJndiTopicParams{
		timeout: timeout,
	}
}

// NewUpdateMsgVpnJndiTopicParamsWithContext creates a new UpdateMsgVpnJndiTopicParams object
// with the ability to set a context for a request.
func NewUpdateMsgVpnJndiTopicParamsWithContext(ctx context.Context) *UpdateMsgVpnJndiTopicParams {
	return &UpdateMsgVpnJndiTopicParams{
		Context: ctx,
	}
}

// NewUpdateMsgVpnJndiTopicParamsWithHTTPClient creates a new UpdateMsgVpnJndiTopicParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMsgVpnJndiTopicParamsWithHTTPClient(client *http.Client) *UpdateMsgVpnJndiTopicParams {
	return &UpdateMsgVpnJndiTopicParams{
		HTTPClient: client,
	}
}

/* UpdateMsgVpnJndiTopicParams contains all the parameters to send to the API endpoint
   for the update msg vpn jndi topic operation.

   Typically these are written to a http.Request.
*/
type UpdateMsgVpnJndiTopicParams struct {

	/* Body.

	   The JNDI Topic object's attributes.
	*/
	Body *models.MsgVpnJndiTopic

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

	/* TopicName.

	   The JNDI name of the JMS Topic.
	*/
	TopicName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update msg vpn jndi topic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMsgVpnJndiTopicParams) WithDefaults() *UpdateMsgVpnJndiTopicParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update msg vpn jndi topic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMsgVpnJndiTopicParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update msg vpn jndi topic params
func (o *UpdateMsgVpnJndiTopicParams) WithTimeout(timeout time.Duration) *UpdateMsgVpnJndiTopicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update msg vpn jndi topic params
func (o *UpdateMsgVpnJndiTopicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update msg vpn jndi topic params
func (o *UpdateMsgVpnJndiTopicParams) WithContext(ctx context.Context) *UpdateMsgVpnJndiTopicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update msg vpn jndi topic params
func (o *UpdateMsgVpnJndiTopicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update msg vpn jndi topic params
func (o *UpdateMsgVpnJndiTopicParams) WithHTTPClient(client *http.Client) *UpdateMsgVpnJndiTopicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update msg vpn jndi topic params
func (o *UpdateMsgVpnJndiTopicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update msg vpn jndi topic params
func (o *UpdateMsgVpnJndiTopicParams) WithBody(body *models.MsgVpnJndiTopic) *UpdateMsgVpnJndiTopicParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update msg vpn jndi topic params
func (o *UpdateMsgVpnJndiTopicParams) SetBody(body *models.MsgVpnJndiTopic) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the update msg vpn jndi topic params
func (o *UpdateMsgVpnJndiTopicParams) WithMsgVpnName(msgVpnName string) *UpdateMsgVpnJndiTopicParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the update msg vpn jndi topic params
func (o *UpdateMsgVpnJndiTopicParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithOpaquePassword adds the opaquePassword to the update msg vpn jndi topic params
func (o *UpdateMsgVpnJndiTopicParams) WithOpaquePassword(opaquePassword *string) *UpdateMsgVpnJndiTopicParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the update msg vpn jndi topic params
func (o *UpdateMsgVpnJndiTopicParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the update msg vpn jndi topic params
func (o *UpdateMsgVpnJndiTopicParams) WithSelect(selectVar []string) *UpdateMsgVpnJndiTopicParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the update msg vpn jndi topic params
func (o *UpdateMsgVpnJndiTopicParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithTopicName adds the topicName to the update msg vpn jndi topic params
func (o *UpdateMsgVpnJndiTopicParams) WithTopicName(topicName string) *UpdateMsgVpnJndiTopicParams {
	o.SetTopicName(topicName)
	return o
}

// SetTopicName adds the topicName to the update msg vpn jndi topic params
func (o *UpdateMsgVpnJndiTopicParams) SetTopicName(topicName string) {
	o.TopicName = topicName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMsgVpnJndiTopicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param topicName
	if err := r.SetPathParam("topicName", o.TopicName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamUpdateMsgVpnJndiTopic binds the parameter select
func (o *UpdateMsgVpnJndiTopicParams) bindParamSelect(formats strfmt.Registry) []string {
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
