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
)

// NewDeleteMsgVpnTopicEndpointTemplateParams creates a new DeleteMsgVpnTopicEndpointTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMsgVpnTopicEndpointTemplateParams() *DeleteMsgVpnTopicEndpointTemplateParams {
	return &DeleteMsgVpnTopicEndpointTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMsgVpnTopicEndpointTemplateParamsWithTimeout creates a new DeleteMsgVpnTopicEndpointTemplateParams object
// with the ability to set a timeout on a request.
func NewDeleteMsgVpnTopicEndpointTemplateParamsWithTimeout(timeout time.Duration) *DeleteMsgVpnTopicEndpointTemplateParams {
	return &DeleteMsgVpnTopicEndpointTemplateParams{
		timeout: timeout,
	}
}

// NewDeleteMsgVpnTopicEndpointTemplateParamsWithContext creates a new DeleteMsgVpnTopicEndpointTemplateParams object
// with the ability to set a context for a request.
func NewDeleteMsgVpnTopicEndpointTemplateParamsWithContext(ctx context.Context) *DeleteMsgVpnTopicEndpointTemplateParams {
	return &DeleteMsgVpnTopicEndpointTemplateParams{
		Context: ctx,
	}
}

// NewDeleteMsgVpnTopicEndpointTemplateParamsWithHTTPClient creates a new DeleteMsgVpnTopicEndpointTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMsgVpnTopicEndpointTemplateParamsWithHTTPClient(client *http.Client) *DeleteMsgVpnTopicEndpointTemplateParams {
	return &DeleteMsgVpnTopicEndpointTemplateParams{
		HTTPClient: client,
	}
}

/* DeleteMsgVpnTopicEndpointTemplateParams contains all the parameters to send to the API endpoint
   for the delete msg vpn topic endpoint template operation.

   Typically these are written to a http.Request.
*/
type DeleteMsgVpnTopicEndpointTemplateParams struct {

	/* MsgVpnName.

	   The name of the Message VPN.
	*/
	MsgVpnName string

	/* TopicEndpointTemplateName.

	   The name of the Topic Endpoint Template.
	*/
	TopicEndpointTemplateName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete msg vpn topic endpoint template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMsgVpnTopicEndpointTemplateParams) WithDefaults() *DeleteMsgVpnTopicEndpointTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete msg vpn topic endpoint template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMsgVpnTopicEndpointTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete msg vpn topic endpoint template params
func (o *DeleteMsgVpnTopicEndpointTemplateParams) WithTimeout(timeout time.Duration) *DeleteMsgVpnTopicEndpointTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete msg vpn topic endpoint template params
func (o *DeleteMsgVpnTopicEndpointTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete msg vpn topic endpoint template params
func (o *DeleteMsgVpnTopicEndpointTemplateParams) WithContext(ctx context.Context) *DeleteMsgVpnTopicEndpointTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete msg vpn topic endpoint template params
func (o *DeleteMsgVpnTopicEndpointTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete msg vpn topic endpoint template params
func (o *DeleteMsgVpnTopicEndpointTemplateParams) WithHTTPClient(client *http.Client) *DeleteMsgVpnTopicEndpointTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete msg vpn topic endpoint template params
func (o *DeleteMsgVpnTopicEndpointTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMsgVpnName adds the msgVpnName to the delete msg vpn topic endpoint template params
func (o *DeleteMsgVpnTopicEndpointTemplateParams) WithMsgVpnName(msgVpnName string) *DeleteMsgVpnTopicEndpointTemplateParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the delete msg vpn topic endpoint template params
func (o *DeleteMsgVpnTopicEndpointTemplateParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithTopicEndpointTemplateName adds the topicEndpointTemplateName to the delete msg vpn topic endpoint template params
func (o *DeleteMsgVpnTopicEndpointTemplateParams) WithTopicEndpointTemplateName(topicEndpointTemplateName string) *DeleteMsgVpnTopicEndpointTemplateParams {
	o.SetTopicEndpointTemplateName(topicEndpointTemplateName)
	return o
}

// SetTopicEndpointTemplateName adds the topicEndpointTemplateName to the delete msg vpn topic endpoint template params
func (o *DeleteMsgVpnTopicEndpointTemplateParams) SetTopicEndpointTemplateName(topicEndpointTemplateName string) {
	o.TopicEndpointTemplateName = topicEndpointTemplateName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMsgVpnTopicEndpointTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	// path param topicEndpointTemplateName
	if err := r.SetPathParam("topicEndpointTemplateName", o.TopicEndpointTemplateName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}