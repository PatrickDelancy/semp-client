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

// NewDeleteMsgVpnBridgeParams creates a new DeleteMsgVpnBridgeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMsgVpnBridgeParams() *DeleteMsgVpnBridgeParams {
	return &DeleteMsgVpnBridgeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMsgVpnBridgeParamsWithTimeout creates a new DeleteMsgVpnBridgeParams object
// with the ability to set a timeout on a request.
func NewDeleteMsgVpnBridgeParamsWithTimeout(timeout time.Duration) *DeleteMsgVpnBridgeParams {
	return &DeleteMsgVpnBridgeParams{
		timeout: timeout,
	}
}

// NewDeleteMsgVpnBridgeParamsWithContext creates a new DeleteMsgVpnBridgeParams object
// with the ability to set a context for a request.
func NewDeleteMsgVpnBridgeParamsWithContext(ctx context.Context) *DeleteMsgVpnBridgeParams {
	return &DeleteMsgVpnBridgeParams{
		Context: ctx,
	}
}

// NewDeleteMsgVpnBridgeParamsWithHTTPClient creates a new DeleteMsgVpnBridgeParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMsgVpnBridgeParamsWithHTTPClient(client *http.Client) *DeleteMsgVpnBridgeParams {
	return &DeleteMsgVpnBridgeParams{
		HTTPClient: client,
	}
}

/* DeleteMsgVpnBridgeParams contains all the parameters to send to the API endpoint
   for the delete msg vpn bridge operation.

   Typically these are written to a http.Request.
*/
type DeleteMsgVpnBridgeParams struct {

	/* BridgeName.

	   The name of the Bridge.
	*/
	BridgeName string

	/* BridgeVirtualRouter.

	   The virtual router of the Bridge.
	*/
	BridgeVirtualRouter string

	/* MsgVpnName.

	   The name of the Message VPN.
	*/
	MsgVpnName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete msg vpn bridge params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMsgVpnBridgeParams) WithDefaults() *DeleteMsgVpnBridgeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete msg vpn bridge params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMsgVpnBridgeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete msg vpn bridge params
func (o *DeleteMsgVpnBridgeParams) WithTimeout(timeout time.Duration) *DeleteMsgVpnBridgeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete msg vpn bridge params
func (o *DeleteMsgVpnBridgeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete msg vpn bridge params
func (o *DeleteMsgVpnBridgeParams) WithContext(ctx context.Context) *DeleteMsgVpnBridgeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete msg vpn bridge params
func (o *DeleteMsgVpnBridgeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete msg vpn bridge params
func (o *DeleteMsgVpnBridgeParams) WithHTTPClient(client *http.Client) *DeleteMsgVpnBridgeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete msg vpn bridge params
func (o *DeleteMsgVpnBridgeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBridgeName adds the bridgeName to the delete msg vpn bridge params
func (o *DeleteMsgVpnBridgeParams) WithBridgeName(bridgeName string) *DeleteMsgVpnBridgeParams {
	o.SetBridgeName(bridgeName)
	return o
}

// SetBridgeName adds the bridgeName to the delete msg vpn bridge params
func (o *DeleteMsgVpnBridgeParams) SetBridgeName(bridgeName string) {
	o.BridgeName = bridgeName
}

// WithBridgeVirtualRouter adds the bridgeVirtualRouter to the delete msg vpn bridge params
func (o *DeleteMsgVpnBridgeParams) WithBridgeVirtualRouter(bridgeVirtualRouter string) *DeleteMsgVpnBridgeParams {
	o.SetBridgeVirtualRouter(bridgeVirtualRouter)
	return o
}

// SetBridgeVirtualRouter adds the bridgeVirtualRouter to the delete msg vpn bridge params
func (o *DeleteMsgVpnBridgeParams) SetBridgeVirtualRouter(bridgeVirtualRouter string) {
	o.BridgeVirtualRouter = bridgeVirtualRouter
}

// WithMsgVpnName adds the msgVpnName to the delete msg vpn bridge params
func (o *DeleteMsgVpnBridgeParams) WithMsgVpnName(msgVpnName string) *DeleteMsgVpnBridgeParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the delete msg vpn bridge params
func (o *DeleteMsgVpnBridgeParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMsgVpnBridgeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bridgeName
	if err := r.SetPathParam("bridgeName", o.BridgeName); err != nil {
		return err
	}

	// path param bridgeVirtualRouter
	if err := r.SetPathParam("bridgeVirtualRouter", o.BridgeVirtualRouter); err != nil {
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
