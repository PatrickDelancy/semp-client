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
)

// NewGetMsgVpnACLProfileSubscribeExceptionParams creates a new GetMsgVpnACLProfileSubscribeExceptionParams object
// with the default values initialized.
func NewGetMsgVpnACLProfileSubscribeExceptionParams() *GetMsgVpnACLProfileSubscribeExceptionParams {
	var ()
	return &GetMsgVpnACLProfileSubscribeExceptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnACLProfileSubscribeExceptionParamsWithTimeout creates a new GetMsgVpnACLProfileSubscribeExceptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMsgVpnACLProfileSubscribeExceptionParamsWithTimeout(timeout time.Duration) *GetMsgVpnACLProfileSubscribeExceptionParams {
	var ()
	return &GetMsgVpnACLProfileSubscribeExceptionParams{

		timeout: timeout,
	}
}

// NewGetMsgVpnACLProfileSubscribeExceptionParamsWithContext creates a new GetMsgVpnACLProfileSubscribeExceptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMsgVpnACLProfileSubscribeExceptionParamsWithContext(ctx context.Context) *GetMsgVpnACLProfileSubscribeExceptionParams {
	var ()
	return &GetMsgVpnACLProfileSubscribeExceptionParams{

		Context: ctx,
	}
}

// NewGetMsgVpnACLProfileSubscribeExceptionParamsWithHTTPClient creates a new GetMsgVpnACLProfileSubscribeExceptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMsgVpnACLProfileSubscribeExceptionParamsWithHTTPClient(client *http.Client) *GetMsgVpnACLProfileSubscribeExceptionParams {
	var ()
	return &GetMsgVpnACLProfileSubscribeExceptionParams{
		HTTPClient: client,
	}
}

/*GetMsgVpnACLProfileSubscribeExceptionParams contains all the parameters to send to the API endpoint
for the get msg vpn Acl profile subscribe exception operation typically these are written to a http.Request
*/
type GetMsgVpnACLProfileSubscribeExceptionParams struct {

	/*ACLProfileName
	  The aclProfileName of the ACL Profile.

	*/
	ACLProfileName string
	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See [Select](#select "Description of the syntax of the `select` parameter").

	*/
	Select []string
	/*SubscribeExceptionTopic
	  The subscribeExceptionTopic of the Subscribe Topic Exception.

	*/
	SubscribeExceptionTopic string
	/*TopicSyntax
	  The topicSyntax of the Subscribe Topic Exception.

	*/
	TopicSyntax string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
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

	valuesSelect := o.Select

	joinedSelect := swag.JoinByFormat(valuesSelect, "csv")
	// query array param select
	if err := r.SetQueryParam("select", joinedSelect...); err != nil {
		return err
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
