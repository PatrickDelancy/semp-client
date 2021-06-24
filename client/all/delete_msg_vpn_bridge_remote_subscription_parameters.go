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

// NewDeleteMsgVpnBridgeRemoteSubscriptionParams creates a new DeleteMsgVpnBridgeRemoteSubscriptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMsgVpnBridgeRemoteSubscriptionParams() *DeleteMsgVpnBridgeRemoteSubscriptionParams {
	return &DeleteMsgVpnBridgeRemoteSubscriptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMsgVpnBridgeRemoteSubscriptionParamsWithTimeout creates a new DeleteMsgVpnBridgeRemoteSubscriptionParams object
// with the ability to set a timeout on a request.
func NewDeleteMsgVpnBridgeRemoteSubscriptionParamsWithTimeout(timeout time.Duration) *DeleteMsgVpnBridgeRemoteSubscriptionParams {
	return &DeleteMsgVpnBridgeRemoteSubscriptionParams{
		timeout: timeout,
	}
}

// NewDeleteMsgVpnBridgeRemoteSubscriptionParamsWithContext creates a new DeleteMsgVpnBridgeRemoteSubscriptionParams object
// with the ability to set a context for a request.
func NewDeleteMsgVpnBridgeRemoteSubscriptionParamsWithContext(ctx context.Context) *DeleteMsgVpnBridgeRemoteSubscriptionParams {
	return &DeleteMsgVpnBridgeRemoteSubscriptionParams{
		Context: ctx,
	}
}

// NewDeleteMsgVpnBridgeRemoteSubscriptionParamsWithHTTPClient creates a new DeleteMsgVpnBridgeRemoteSubscriptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMsgVpnBridgeRemoteSubscriptionParamsWithHTTPClient(client *http.Client) *DeleteMsgVpnBridgeRemoteSubscriptionParams {
	return &DeleteMsgVpnBridgeRemoteSubscriptionParams{
		HTTPClient: client,
	}
}

/* DeleteMsgVpnBridgeRemoteSubscriptionParams contains all the parameters to send to the API endpoint
   for the delete msg vpn bridge remote subscription operation.

   Typically these are written to a http.Request.
*/
type DeleteMsgVpnBridgeRemoteSubscriptionParams struct {

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

	/* RemoteSubscriptionTopic.

	   The topic of the Bridge remote subscription.
	*/
	RemoteSubscriptionTopic string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete msg vpn bridge remote subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMsgVpnBridgeRemoteSubscriptionParams) WithDefaults() *DeleteMsgVpnBridgeRemoteSubscriptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete msg vpn bridge remote subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMsgVpnBridgeRemoteSubscriptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete msg vpn bridge remote subscription params
func (o *DeleteMsgVpnBridgeRemoteSubscriptionParams) WithTimeout(timeout time.Duration) *DeleteMsgVpnBridgeRemoteSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete msg vpn bridge remote subscription params
func (o *DeleteMsgVpnBridgeRemoteSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete msg vpn bridge remote subscription params
func (o *DeleteMsgVpnBridgeRemoteSubscriptionParams) WithContext(ctx context.Context) *DeleteMsgVpnBridgeRemoteSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete msg vpn bridge remote subscription params
func (o *DeleteMsgVpnBridgeRemoteSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete msg vpn bridge remote subscription params
func (o *DeleteMsgVpnBridgeRemoteSubscriptionParams) WithHTTPClient(client *http.Client) *DeleteMsgVpnBridgeRemoteSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete msg vpn bridge remote subscription params
func (o *DeleteMsgVpnBridgeRemoteSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBridgeName adds the bridgeName to the delete msg vpn bridge remote subscription params
func (o *DeleteMsgVpnBridgeRemoteSubscriptionParams) WithBridgeName(bridgeName string) *DeleteMsgVpnBridgeRemoteSubscriptionParams {
	o.SetBridgeName(bridgeName)
	return o
}

// SetBridgeName adds the bridgeName to the delete msg vpn bridge remote subscription params
func (o *DeleteMsgVpnBridgeRemoteSubscriptionParams) SetBridgeName(bridgeName string) {
	o.BridgeName = bridgeName
}

// WithBridgeVirtualRouter adds the bridgeVirtualRouter to the delete msg vpn bridge remote subscription params
func (o *DeleteMsgVpnBridgeRemoteSubscriptionParams) WithBridgeVirtualRouter(bridgeVirtualRouter string) *DeleteMsgVpnBridgeRemoteSubscriptionParams {
	o.SetBridgeVirtualRouter(bridgeVirtualRouter)
	return o
}

// SetBridgeVirtualRouter adds the bridgeVirtualRouter to the delete msg vpn bridge remote subscription params
func (o *DeleteMsgVpnBridgeRemoteSubscriptionParams) SetBridgeVirtualRouter(bridgeVirtualRouter string) {
	o.BridgeVirtualRouter = bridgeVirtualRouter
}

// WithMsgVpnName adds the msgVpnName to the delete msg vpn bridge remote subscription params
func (o *DeleteMsgVpnBridgeRemoteSubscriptionParams) WithMsgVpnName(msgVpnName string) *DeleteMsgVpnBridgeRemoteSubscriptionParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the delete msg vpn bridge remote subscription params
func (o *DeleteMsgVpnBridgeRemoteSubscriptionParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithRemoteSubscriptionTopic adds the remoteSubscriptionTopic to the delete msg vpn bridge remote subscription params
func (o *DeleteMsgVpnBridgeRemoteSubscriptionParams) WithRemoteSubscriptionTopic(remoteSubscriptionTopic string) *DeleteMsgVpnBridgeRemoteSubscriptionParams {
	o.SetRemoteSubscriptionTopic(remoteSubscriptionTopic)
	return o
}

// SetRemoteSubscriptionTopic adds the remoteSubscriptionTopic to the delete msg vpn bridge remote subscription params
func (o *DeleteMsgVpnBridgeRemoteSubscriptionParams) SetRemoteSubscriptionTopic(remoteSubscriptionTopic string) {
	o.RemoteSubscriptionTopic = remoteSubscriptionTopic
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMsgVpnBridgeRemoteSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param remoteSubscriptionTopic
	if err := r.SetPathParam("remoteSubscriptionTopic", o.RemoteSubscriptionTopic); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}