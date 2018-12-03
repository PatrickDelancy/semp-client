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

	models "github.com/ExalDraen/semp-client/models"
)

// NewReplaceMsgVpnBridgeRemoteMsgVpnParams creates a new ReplaceMsgVpnBridgeRemoteMsgVpnParams object
// with the default values initialized.
func NewReplaceMsgVpnBridgeRemoteMsgVpnParams() *ReplaceMsgVpnBridgeRemoteMsgVpnParams {
	var ()
	return &ReplaceMsgVpnBridgeRemoteMsgVpnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceMsgVpnBridgeRemoteMsgVpnParamsWithTimeout creates a new ReplaceMsgVpnBridgeRemoteMsgVpnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceMsgVpnBridgeRemoteMsgVpnParamsWithTimeout(timeout time.Duration) *ReplaceMsgVpnBridgeRemoteMsgVpnParams {
	var ()
	return &ReplaceMsgVpnBridgeRemoteMsgVpnParams{

		timeout: timeout,
	}
}

// NewReplaceMsgVpnBridgeRemoteMsgVpnParamsWithContext creates a new ReplaceMsgVpnBridgeRemoteMsgVpnParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceMsgVpnBridgeRemoteMsgVpnParamsWithContext(ctx context.Context) *ReplaceMsgVpnBridgeRemoteMsgVpnParams {
	var ()
	return &ReplaceMsgVpnBridgeRemoteMsgVpnParams{

		Context: ctx,
	}
}

// NewReplaceMsgVpnBridgeRemoteMsgVpnParamsWithHTTPClient creates a new ReplaceMsgVpnBridgeRemoteMsgVpnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceMsgVpnBridgeRemoteMsgVpnParamsWithHTTPClient(client *http.Client) *ReplaceMsgVpnBridgeRemoteMsgVpnParams {
	var ()
	return &ReplaceMsgVpnBridgeRemoteMsgVpnParams{
		HTTPClient: client,
	}
}

/*ReplaceMsgVpnBridgeRemoteMsgVpnParams contains all the parameters to send to the API endpoint
for the replace msg vpn bridge remote msg vpn operation typically these are written to a http.Request
*/
type ReplaceMsgVpnBridgeRemoteMsgVpnParams struct {

	/*Body
	  The Remote Message VPN object's attributes.

	*/
	Body *models.MsgVpnBridgeRemoteMsgVpn
	/*BridgeName
	  The bridgeName of the Bridge.

	*/
	BridgeName string
	/*BridgeVirtualRouter
	  The bridgeVirtualRouter of the Bridge.

	*/
	BridgeVirtualRouter string
	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string
	/*RemoteMsgVpnInterface
	  The remoteMsgVpnInterface of the Remote Message VPN.

	*/
	RemoteMsgVpnInterface string
	/*RemoteMsgVpnLocation
	  The remoteMsgVpnLocation of the Remote Message VPN.

	*/
	RemoteMsgVpnLocation string
	/*RemoteMsgVpnName
	  The remoteMsgVpnName of the Remote Message VPN.

	*/
	RemoteMsgVpnName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See [Select](#select "Description of the syntax of the `select` parameter").

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) WithTimeout(timeout time.Duration) *ReplaceMsgVpnBridgeRemoteMsgVpnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) WithContext(ctx context.Context) *ReplaceMsgVpnBridgeRemoteMsgVpnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) WithHTTPClient(client *http.Client) *ReplaceMsgVpnBridgeRemoteMsgVpnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) WithBody(body *models.MsgVpnBridgeRemoteMsgVpn) *ReplaceMsgVpnBridgeRemoteMsgVpnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) SetBody(body *models.MsgVpnBridgeRemoteMsgVpn) {
	o.Body = body
}

// WithBridgeName adds the bridgeName to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) WithBridgeName(bridgeName string) *ReplaceMsgVpnBridgeRemoteMsgVpnParams {
	o.SetBridgeName(bridgeName)
	return o
}

// SetBridgeName adds the bridgeName to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) SetBridgeName(bridgeName string) {
	o.BridgeName = bridgeName
}

// WithBridgeVirtualRouter adds the bridgeVirtualRouter to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) WithBridgeVirtualRouter(bridgeVirtualRouter string) *ReplaceMsgVpnBridgeRemoteMsgVpnParams {
	o.SetBridgeVirtualRouter(bridgeVirtualRouter)
	return o
}

// SetBridgeVirtualRouter adds the bridgeVirtualRouter to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) SetBridgeVirtualRouter(bridgeVirtualRouter string) {
	o.BridgeVirtualRouter = bridgeVirtualRouter
}

// WithMsgVpnName adds the msgVpnName to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) WithMsgVpnName(msgVpnName string) *ReplaceMsgVpnBridgeRemoteMsgVpnParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithRemoteMsgVpnInterface adds the remoteMsgVpnInterface to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) WithRemoteMsgVpnInterface(remoteMsgVpnInterface string) *ReplaceMsgVpnBridgeRemoteMsgVpnParams {
	o.SetRemoteMsgVpnInterface(remoteMsgVpnInterface)
	return o
}

// SetRemoteMsgVpnInterface adds the remoteMsgVpnInterface to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) SetRemoteMsgVpnInterface(remoteMsgVpnInterface string) {
	o.RemoteMsgVpnInterface = remoteMsgVpnInterface
}

// WithRemoteMsgVpnLocation adds the remoteMsgVpnLocation to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) WithRemoteMsgVpnLocation(remoteMsgVpnLocation string) *ReplaceMsgVpnBridgeRemoteMsgVpnParams {
	o.SetRemoteMsgVpnLocation(remoteMsgVpnLocation)
	return o
}

// SetRemoteMsgVpnLocation adds the remoteMsgVpnLocation to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) SetRemoteMsgVpnLocation(remoteMsgVpnLocation string) {
	o.RemoteMsgVpnLocation = remoteMsgVpnLocation
}

// WithRemoteMsgVpnName adds the remoteMsgVpnName to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) WithRemoteMsgVpnName(remoteMsgVpnName string) *ReplaceMsgVpnBridgeRemoteMsgVpnParams {
	o.SetRemoteMsgVpnName(remoteMsgVpnName)
	return o
}

// SetRemoteMsgVpnName adds the remoteMsgVpnName to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) SetRemoteMsgVpnName(remoteMsgVpnName string) {
	o.RemoteMsgVpnName = remoteMsgVpnName
}

// WithSelect adds the selectVar to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) WithSelect(selectVar []string) *ReplaceMsgVpnBridgeRemoteMsgVpnParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the replace msg vpn bridge remote msg vpn params
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceMsgVpnBridgeRemoteMsgVpnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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

	// path param remoteMsgVpnInterface
	if err := r.SetPathParam("remoteMsgVpnInterface", o.RemoteMsgVpnInterface); err != nil {
		return err
	}

	// path param remoteMsgVpnLocation
	if err := r.SetPathParam("remoteMsgVpnLocation", o.RemoteMsgVpnLocation); err != nil {
		return err
	}

	// path param remoteMsgVpnName
	if err := r.SetPathParam("remoteMsgVpnName", o.RemoteMsgVpnName); err != nil {
		return err
	}

	valuesSelect := o.Select

	joinedSelect := swag.JoinByFormat(valuesSelect, "csv")
	// query array param select
	if err := r.SetQueryParam("select", joinedSelect...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
