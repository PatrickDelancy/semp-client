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

// NewDeleteMsgVpnClientUsernameParams creates a new DeleteMsgVpnClientUsernameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMsgVpnClientUsernameParams() *DeleteMsgVpnClientUsernameParams {
	return &DeleteMsgVpnClientUsernameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMsgVpnClientUsernameParamsWithTimeout creates a new DeleteMsgVpnClientUsernameParams object
// with the ability to set a timeout on a request.
func NewDeleteMsgVpnClientUsernameParamsWithTimeout(timeout time.Duration) *DeleteMsgVpnClientUsernameParams {
	return &DeleteMsgVpnClientUsernameParams{
		timeout: timeout,
	}
}

// NewDeleteMsgVpnClientUsernameParamsWithContext creates a new DeleteMsgVpnClientUsernameParams object
// with the ability to set a context for a request.
func NewDeleteMsgVpnClientUsernameParamsWithContext(ctx context.Context) *DeleteMsgVpnClientUsernameParams {
	return &DeleteMsgVpnClientUsernameParams{
		Context: ctx,
	}
}

// NewDeleteMsgVpnClientUsernameParamsWithHTTPClient creates a new DeleteMsgVpnClientUsernameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMsgVpnClientUsernameParamsWithHTTPClient(client *http.Client) *DeleteMsgVpnClientUsernameParams {
	return &DeleteMsgVpnClientUsernameParams{
		HTTPClient: client,
	}
}

/* DeleteMsgVpnClientUsernameParams contains all the parameters to send to the API endpoint
   for the delete msg vpn client username operation.

   Typically these are written to a http.Request.
*/
type DeleteMsgVpnClientUsernameParams struct {

	/* ClientUsername.

	   The name of the Client Username.
	*/
	ClientUsername string

	/* MsgVpnName.

	   The name of the Message VPN.
	*/
	MsgVpnName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete msg vpn client username params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMsgVpnClientUsernameParams) WithDefaults() *DeleteMsgVpnClientUsernameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete msg vpn client username params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMsgVpnClientUsernameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete msg vpn client username params
func (o *DeleteMsgVpnClientUsernameParams) WithTimeout(timeout time.Duration) *DeleteMsgVpnClientUsernameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete msg vpn client username params
func (o *DeleteMsgVpnClientUsernameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete msg vpn client username params
func (o *DeleteMsgVpnClientUsernameParams) WithContext(ctx context.Context) *DeleteMsgVpnClientUsernameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete msg vpn client username params
func (o *DeleteMsgVpnClientUsernameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete msg vpn client username params
func (o *DeleteMsgVpnClientUsernameParams) WithHTTPClient(client *http.Client) *DeleteMsgVpnClientUsernameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete msg vpn client username params
func (o *DeleteMsgVpnClientUsernameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientUsername adds the clientUsername to the delete msg vpn client username params
func (o *DeleteMsgVpnClientUsernameParams) WithClientUsername(clientUsername string) *DeleteMsgVpnClientUsernameParams {
	o.SetClientUsername(clientUsername)
	return o
}

// SetClientUsername adds the clientUsername to the delete msg vpn client username params
func (o *DeleteMsgVpnClientUsernameParams) SetClientUsername(clientUsername string) {
	o.ClientUsername = clientUsername
}

// WithMsgVpnName adds the msgVpnName to the delete msg vpn client username params
func (o *DeleteMsgVpnClientUsernameParams) WithMsgVpnName(msgVpnName string) *DeleteMsgVpnClientUsernameParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the delete msg vpn client username params
func (o *DeleteMsgVpnClientUsernameParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMsgVpnClientUsernameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clientUsername
	if err := r.SetPathParam("clientUsername", o.ClientUsername); err != nil {
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