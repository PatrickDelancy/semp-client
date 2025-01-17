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

// NewDeleteMsgVpnParams creates a new DeleteMsgVpnParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMsgVpnParams() *DeleteMsgVpnParams {
	return &DeleteMsgVpnParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMsgVpnParamsWithTimeout creates a new DeleteMsgVpnParams object
// with the ability to set a timeout on a request.
func NewDeleteMsgVpnParamsWithTimeout(timeout time.Duration) *DeleteMsgVpnParams {
	return &DeleteMsgVpnParams{
		timeout: timeout,
	}
}

// NewDeleteMsgVpnParamsWithContext creates a new DeleteMsgVpnParams object
// with the ability to set a context for a request.
func NewDeleteMsgVpnParamsWithContext(ctx context.Context) *DeleteMsgVpnParams {
	return &DeleteMsgVpnParams{
		Context: ctx,
	}
}

// NewDeleteMsgVpnParamsWithHTTPClient creates a new DeleteMsgVpnParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMsgVpnParamsWithHTTPClient(client *http.Client) *DeleteMsgVpnParams {
	return &DeleteMsgVpnParams{
		HTTPClient: client,
	}
}

/* DeleteMsgVpnParams contains all the parameters to send to the API endpoint
   for the delete msg vpn operation.

   Typically these are written to a http.Request.
*/
type DeleteMsgVpnParams struct {

	/* MsgVpnName.

	   The name of the Message VPN.
	*/
	MsgVpnName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete msg vpn params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMsgVpnParams) WithDefaults() *DeleteMsgVpnParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete msg vpn params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMsgVpnParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete msg vpn params
func (o *DeleteMsgVpnParams) WithTimeout(timeout time.Duration) *DeleteMsgVpnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete msg vpn params
func (o *DeleteMsgVpnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete msg vpn params
func (o *DeleteMsgVpnParams) WithContext(ctx context.Context) *DeleteMsgVpnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete msg vpn params
func (o *DeleteMsgVpnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete msg vpn params
func (o *DeleteMsgVpnParams) WithHTTPClient(client *http.Client) *DeleteMsgVpnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete msg vpn params
func (o *DeleteMsgVpnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMsgVpnName adds the msgVpnName to the delete msg vpn params
func (o *DeleteMsgVpnParams) WithMsgVpnName(msgVpnName string) *DeleteMsgVpnParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the delete msg vpn params
func (o *DeleteMsgVpnParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMsgVpnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
