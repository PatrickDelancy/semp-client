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

// NewGetMsgVpnACLProfileSubscribeExceptionParams creates a new GetMsgVpnACLProfileSubscribeExceptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMsgVpnACLProfileSubscribeExceptionParams() *GetMsgVpnACLProfileSubscribeExceptionParams {
	return &GetMsgVpnACLProfileSubscribeExceptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnACLProfileSubscribeExceptionParamsWithTimeout creates a new GetMsgVpnACLProfileSubscribeExceptionParams object
// with the ability to set a timeout on a request.
func NewGetMsgVpnACLProfileSubscribeExceptionParamsWithTimeout(timeout time.Duration) *GetMsgVpnACLProfileSubscribeExceptionParams {
	return &GetMsgVpnACLProfileSubscribeExceptionParams{
		timeout: timeout,
	}
}

// NewGetMsgVpnACLProfileSubscribeExceptionParamsWithContext creates a new GetMsgVpnACLProfileSubscribeExceptionParams object
// with the ability to set a context for a request.
func NewGetMsgVpnACLProfileSubscribeExceptionParamsWithContext(ctx context.Context) *GetMsgVpnACLProfileSubscribeExceptionParams {
	return &GetMsgVpnACLProfileSubscribeExceptionParams{
		Context: ctx,
	}
}

// NewGetMsgVpnACLProfileSubscribeExceptionParamsWithHTTPClient creates a new GetMsgVpnACLProfileSubscribeExceptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMsgVpnACLProfileSubscribeExceptionParamsWithHTTPClient(client *http.Client) *GetMsgVpnACLProfileSubscribeExceptionParams {
	return &GetMsgVpnACLProfileSubscribeExceptionParams{
		HTTPClient: client,
	}
}

/* GetMsgVpnACLProfileSubscribeExceptionParams contains all the parameters to send to the API endpoint
   for the get msg vpn Acl profile subscribe exception operation.

   Typically these are written to a http.Request.
*/
type GetMsgVpnACLProfileSubscribeExceptionParams struct {

	/* ACLProfileName.

	   The name of the ACL Profile.
	*/
	ACLProfileName string

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

	/* SubscribeExceptionTopic.

	   The topic for the exception to the default action taken. May include wildcard characters.
	*/
	SubscribeExceptionTopic string

	/* TopicSyntax.

	   The syntax of the topic for the exception to the default action taken.
	*/
	TopicSyntax string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get msg vpn Acl profile subscribe exception params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) WithDefaults() *GetMsgVpnACLProfileSubscribeExceptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get msg vpn Acl profile subscribe exception params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get msg vpn Acl profile subscribe exception params
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) WithTimeout(timeout time.Duration) *GetMsgVpnACLProfileSubscribeExceptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn Acl profile subscribe exception params
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn Acl profile subscribe exception params
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) WithContext(ctx context.Context) *GetMsgVpnACLProfileSubscribeExceptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn Acl profile subscribe exception params
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn Acl profile subscribe exception params
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) WithHTTPClient(client *http.Client) *GetMsgVpnACLProfileSubscribeExceptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn Acl profile subscribe exception params
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithACLProfileName adds the aCLProfileName to the get msg vpn Acl profile subscribe exception params
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) WithACLProfileName(aCLProfileName string) *GetMsgVpnACLProfileSubscribeExceptionParams {
	o.SetACLProfileName(aCLProfileName)
	return o
}

// SetACLProfileName adds the aclProfileName to the get msg vpn Acl profile subscribe exception params
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) SetACLProfileName(aCLProfileName string) {
	o.ACLProfileName = aCLProfileName
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn Acl profile subscribe exception params
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnACLProfileSubscribeExceptionParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn Acl profile subscribe exception params
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithOpaquePassword adds the opaquePassword to the get msg vpn Acl profile subscribe exception params
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) WithOpaquePassword(opaquePassword *string) *GetMsgVpnACLProfileSubscribeExceptionParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the get msg vpn Acl profile subscribe exception params
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the get msg vpn Acl profile subscribe exception params
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) WithSelect(selectVar []string) *GetMsgVpnACLProfileSubscribeExceptionParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn Acl profile subscribe exception params
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithSubscribeExceptionTopic adds the subscribeExceptionTopic to the get msg vpn Acl profile subscribe exception params
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) WithSubscribeExceptionTopic(subscribeExceptionTopic string) *GetMsgVpnACLProfileSubscribeExceptionParams {
	o.SetSubscribeExceptionTopic(subscribeExceptionTopic)
	return o
}

// SetSubscribeExceptionTopic adds the subscribeExceptionTopic to the get msg vpn Acl profile subscribe exception params
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) SetSubscribeExceptionTopic(subscribeExceptionTopic string) {
	o.SubscribeExceptionTopic = subscribeExceptionTopic
}

// WithTopicSyntax adds the topicSyntax to the get msg vpn Acl profile subscribe exception params
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) WithTopicSyntax(topicSyntax string) *GetMsgVpnACLProfileSubscribeExceptionParams {
	o.SetTopicSyntax(topicSyntax)
	return o
}

// SetTopicSyntax adds the topicSyntax to the get msg vpn Acl profile subscribe exception params
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) SetTopicSyntax(topicSyntax string) {
	o.TopicSyntax = topicSyntax
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param aclProfileName
	if err := r.SetPathParam("aclProfileName", o.ACLProfileName); err != nil {
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

	// path param subscribeExceptionTopic
	if err := r.SetPathParam("subscribeExceptionTopic", o.SubscribeExceptionTopic); err != nil {
		return err
	}

	// path param topicSyntax
	if err := r.SetPathParam("topicSyntax", o.TopicSyntax); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetMsgVpnACLProfileSubscribeException binds the parameter select
func (o *GetMsgVpnACLProfileSubscribeExceptionParams) bindParamSelect(formats strfmt.Registry) []string {
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