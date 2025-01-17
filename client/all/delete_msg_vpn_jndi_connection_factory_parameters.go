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

// NewDeleteMsgVpnJndiConnectionFactoryParams creates a new DeleteMsgVpnJndiConnectionFactoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMsgVpnJndiConnectionFactoryParams() *DeleteMsgVpnJndiConnectionFactoryParams {
	return &DeleteMsgVpnJndiConnectionFactoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMsgVpnJndiConnectionFactoryParamsWithTimeout creates a new DeleteMsgVpnJndiConnectionFactoryParams object
// with the ability to set a timeout on a request.
func NewDeleteMsgVpnJndiConnectionFactoryParamsWithTimeout(timeout time.Duration) *DeleteMsgVpnJndiConnectionFactoryParams {
	return &DeleteMsgVpnJndiConnectionFactoryParams{
		timeout: timeout,
	}
}

// NewDeleteMsgVpnJndiConnectionFactoryParamsWithContext creates a new DeleteMsgVpnJndiConnectionFactoryParams object
// with the ability to set a context for a request.
func NewDeleteMsgVpnJndiConnectionFactoryParamsWithContext(ctx context.Context) *DeleteMsgVpnJndiConnectionFactoryParams {
	return &DeleteMsgVpnJndiConnectionFactoryParams{
		Context: ctx,
	}
}

// NewDeleteMsgVpnJndiConnectionFactoryParamsWithHTTPClient creates a new DeleteMsgVpnJndiConnectionFactoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMsgVpnJndiConnectionFactoryParamsWithHTTPClient(client *http.Client) *DeleteMsgVpnJndiConnectionFactoryParams {
	return &DeleteMsgVpnJndiConnectionFactoryParams{
		HTTPClient: client,
	}
}

/* DeleteMsgVpnJndiConnectionFactoryParams contains all the parameters to send to the API endpoint
   for the delete msg vpn jndi connection factory operation.

   Typically these are written to a http.Request.
*/
type DeleteMsgVpnJndiConnectionFactoryParams struct {

	/* ConnectionFactoryName.

	   The name of the JMS Connection Factory.
	*/
	ConnectionFactoryName string

	/* MsgVpnName.

	   The name of the Message VPN.
	*/
	MsgVpnName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete msg vpn jndi connection factory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMsgVpnJndiConnectionFactoryParams) WithDefaults() *DeleteMsgVpnJndiConnectionFactoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete msg vpn jndi connection factory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMsgVpnJndiConnectionFactoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete msg vpn jndi connection factory params
func (o *DeleteMsgVpnJndiConnectionFactoryParams) WithTimeout(timeout time.Duration) *DeleteMsgVpnJndiConnectionFactoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete msg vpn jndi connection factory params
func (o *DeleteMsgVpnJndiConnectionFactoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete msg vpn jndi connection factory params
func (o *DeleteMsgVpnJndiConnectionFactoryParams) WithContext(ctx context.Context) *DeleteMsgVpnJndiConnectionFactoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete msg vpn jndi connection factory params
func (o *DeleteMsgVpnJndiConnectionFactoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete msg vpn jndi connection factory params
func (o *DeleteMsgVpnJndiConnectionFactoryParams) WithHTTPClient(client *http.Client) *DeleteMsgVpnJndiConnectionFactoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete msg vpn jndi connection factory params
func (o *DeleteMsgVpnJndiConnectionFactoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionFactoryName adds the connectionFactoryName to the delete msg vpn jndi connection factory params
func (o *DeleteMsgVpnJndiConnectionFactoryParams) WithConnectionFactoryName(connectionFactoryName string) *DeleteMsgVpnJndiConnectionFactoryParams {
	o.SetConnectionFactoryName(connectionFactoryName)
	return o
}

// SetConnectionFactoryName adds the connectionFactoryName to the delete msg vpn jndi connection factory params
func (o *DeleteMsgVpnJndiConnectionFactoryParams) SetConnectionFactoryName(connectionFactoryName string) {
	o.ConnectionFactoryName = connectionFactoryName
}

// WithMsgVpnName adds the msgVpnName to the delete msg vpn jndi connection factory params
func (o *DeleteMsgVpnJndiConnectionFactoryParams) WithMsgVpnName(msgVpnName string) *DeleteMsgVpnJndiConnectionFactoryParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the delete msg vpn jndi connection factory params
func (o *DeleteMsgVpnJndiConnectionFactoryParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMsgVpnJndiConnectionFactoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connectionFactoryName
	if err := r.SetPathParam("connectionFactoryName", o.ConnectionFactoryName); err != nil {
		return err
	}

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
