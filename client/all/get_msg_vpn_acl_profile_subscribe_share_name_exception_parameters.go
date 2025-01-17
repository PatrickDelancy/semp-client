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

// NewGetMsgVpnACLProfileSubscribeShareNameExceptionParams creates a new GetMsgVpnACLProfileSubscribeShareNameExceptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMsgVpnACLProfileSubscribeShareNameExceptionParams() *GetMsgVpnACLProfileSubscribeShareNameExceptionParams {
	return &GetMsgVpnACLProfileSubscribeShareNameExceptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnACLProfileSubscribeShareNameExceptionParamsWithTimeout creates a new GetMsgVpnACLProfileSubscribeShareNameExceptionParams object
// with the ability to set a timeout on a request.
func NewGetMsgVpnACLProfileSubscribeShareNameExceptionParamsWithTimeout(timeout time.Duration) *GetMsgVpnACLProfileSubscribeShareNameExceptionParams {
	return &GetMsgVpnACLProfileSubscribeShareNameExceptionParams{
		timeout: timeout,
	}
}

// NewGetMsgVpnACLProfileSubscribeShareNameExceptionParamsWithContext creates a new GetMsgVpnACLProfileSubscribeShareNameExceptionParams object
// with the ability to set a context for a request.
func NewGetMsgVpnACLProfileSubscribeShareNameExceptionParamsWithContext(ctx context.Context) *GetMsgVpnACLProfileSubscribeShareNameExceptionParams {
	return &GetMsgVpnACLProfileSubscribeShareNameExceptionParams{
		Context: ctx,
	}
}

// NewGetMsgVpnACLProfileSubscribeShareNameExceptionParamsWithHTTPClient creates a new GetMsgVpnACLProfileSubscribeShareNameExceptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMsgVpnACLProfileSubscribeShareNameExceptionParamsWithHTTPClient(client *http.Client) *GetMsgVpnACLProfileSubscribeShareNameExceptionParams {
	return &GetMsgVpnACLProfileSubscribeShareNameExceptionParams{
		HTTPClient: client,
	}
}

/* GetMsgVpnACLProfileSubscribeShareNameExceptionParams contains all the parameters to send to the API endpoint
   for the get msg vpn Acl profile subscribe share name exception operation.

   Typically these are written to a http.Request.
*/
type GetMsgVpnACLProfileSubscribeShareNameExceptionParams struct {

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

	/* SubscribeShareNameException.

	   The subscribe share name exception to the default action taken. May include wildcard characters.
	*/
	SubscribeShareNameException string

	/* SubscribeShareNameExceptionSyntax.

	   The syntax of the subscribe share name for the exception to the default action taken.
	*/
	SubscribeShareNameExceptionSyntax string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get msg vpn Acl profile subscribe share name exception params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) WithDefaults() *GetMsgVpnACLProfileSubscribeShareNameExceptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get msg vpn Acl profile subscribe share name exception params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get msg vpn Acl profile subscribe share name exception params
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) WithTimeout(timeout time.Duration) *GetMsgVpnACLProfileSubscribeShareNameExceptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn Acl profile subscribe share name exception params
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn Acl profile subscribe share name exception params
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) WithContext(ctx context.Context) *GetMsgVpnACLProfileSubscribeShareNameExceptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn Acl profile subscribe share name exception params
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn Acl profile subscribe share name exception params
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) WithHTTPClient(client *http.Client) *GetMsgVpnACLProfileSubscribeShareNameExceptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn Acl profile subscribe share name exception params
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithACLProfileName adds the aCLProfileName to the get msg vpn Acl profile subscribe share name exception params
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) WithACLProfileName(aCLProfileName string) *GetMsgVpnACLProfileSubscribeShareNameExceptionParams {
	o.SetACLProfileName(aCLProfileName)
	return o
}

// SetACLProfileName adds the aclProfileName to the get msg vpn Acl profile subscribe share name exception params
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) SetACLProfileName(aCLProfileName string) {
	o.ACLProfileName = aCLProfileName
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn Acl profile subscribe share name exception params
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnACLProfileSubscribeShareNameExceptionParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn Acl profile subscribe share name exception params
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithOpaquePassword adds the opaquePassword to the get msg vpn Acl profile subscribe share name exception params
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) WithOpaquePassword(opaquePassword *string) *GetMsgVpnACLProfileSubscribeShareNameExceptionParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the get msg vpn Acl profile subscribe share name exception params
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the get msg vpn Acl profile subscribe share name exception params
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) WithSelect(selectVar []string) *GetMsgVpnACLProfileSubscribeShareNameExceptionParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn Acl profile subscribe share name exception params
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithSubscribeShareNameException adds the subscribeShareNameException to the get msg vpn Acl profile subscribe share name exception params
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) WithSubscribeShareNameException(subscribeShareNameException string) *GetMsgVpnACLProfileSubscribeShareNameExceptionParams {
	o.SetSubscribeShareNameException(subscribeShareNameException)
	return o
}

// SetSubscribeShareNameException adds the subscribeShareNameException to the get msg vpn Acl profile subscribe share name exception params
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) SetSubscribeShareNameException(subscribeShareNameException string) {
	o.SubscribeShareNameException = subscribeShareNameException
}

// WithSubscribeShareNameExceptionSyntax adds the subscribeShareNameExceptionSyntax to the get msg vpn Acl profile subscribe share name exception params
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) WithSubscribeShareNameExceptionSyntax(subscribeShareNameExceptionSyntax string) *GetMsgVpnACLProfileSubscribeShareNameExceptionParams {
	o.SetSubscribeShareNameExceptionSyntax(subscribeShareNameExceptionSyntax)
	return o
}

// SetSubscribeShareNameExceptionSyntax adds the subscribeShareNameExceptionSyntax to the get msg vpn Acl profile subscribe share name exception params
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) SetSubscribeShareNameExceptionSyntax(subscribeShareNameExceptionSyntax string) {
	o.SubscribeShareNameExceptionSyntax = subscribeShareNameExceptionSyntax
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param subscribeShareNameException
	if err := r.SetPathParam("subscribeShareNameException", o.SubscribeShareNameException); err != nil {
		return err
	}

	// path param subscribeShareNameExceptionSyntax
	if err := r.SetPathParam("subscribeShareNameExceptionSyntax", o.SubscribeShareNameExceptionSyntax); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetMsgVpnACLProfileSubscribeShareNameException binds the parameter select
func (o *GetMsgVpnACLProfileSubscribeShareNameExceptionParams) bindParamSelect(formats strfmt.Registry) []string {
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
