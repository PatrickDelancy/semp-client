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

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteMsgVpnClientProfileParams creates a new DeleteMsgVpnClientProfileParams object
// with the default values initialized.
func NewDeleteMsgVpnClientProfileParams() *DeleteMsgVpnClientProfileParams {
	var ()
	return &DeleteMsgVpnClientProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMsgVpnClientProfileParamsWithTimeout creates a new DeleteMsgVpnClientProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMsgVpnClientProfileParamsWithTimeout(timeout time.Duration) *DeleteMsgVpnClientProfileParams {
	var ()
	return &DeleteMsgVpnClientProfileParams{

		timeout: timeout,
	}
}

// NewDeleteMsgVpnClientProfileParamsWithContext creates a new DeleteMsgVpnClientProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMsgVpnClientProfileParamsWithContext(ctx context.Context) *DeleteMsgVpnClientProfileParams {
	var ()
	return &DeleteMsgVpnClientProfileParams{

		Context: ctx,
	}
}

// NewDeleteMsgVpnClientProfileParamsWithHTTPClient creates a new DeleteMsgVpnClientProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMsgVpnClientProfileParamsWithHTTPClient(client *http.Client) *DeleteMsgVpnClientProfileParams {
	var ()
	return &DeleteMsgVpnClientProfileParams{
		HTTPClient: client,
	}
}

/*DeleteMsgVpnClientProfileParams contains all the parameters to send to the API endpoint
for the delete msg vpn client profile operation typically these are written to a http.Request
*/
type DeleteMsgVpnClientProfileParams struct {

	/*ClientProfileName
	  The clientProfileName of the Client Profile.

	*/
	ClientProfileName string
	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete msg vpn client profile params
func (o *DeleteMsgVpnClientProfileParams) WithTimeout(timeout time.Duration) *DeleteMsgVpnClientProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete msg vpn client profile params
func (o *DeleteMsgVpnClientProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete msg vpn client profile params
func (o *DeleteMsgVpnClientProfileParams) WithContext(ctx context.Context) *DeleteMsgVpnClientProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete msg vpn client profile params
func (o *DeleteMsgVpnClientProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete msg vpn client profile params
func (o *DeleteMsgVpnClientProfileParams) WithHTTPClient(client *http.Client) *DeleteMsgVpnClientProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete msg vpn client profile params
func (o *DeleteMsgVpnClientProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientProfileName adds the clientProfileName to the delete msg vpn client profile params
func (o *DeleteMsgVpnClientProfileParams) WithClientProfileName(clientProfileName string) *DeleteMsgVpnClientProfileParams {
	o.SetClientProfileName(clientProfileName)
	return o
}

// SetClientProfileName adds the clientProfileName to the delete msg vpn client profile params
func (o *DeleteMsgVpnClientProfileParams) SetClientProfileName(clientProfileName string) {
	o.ClientProfileName = clientProfileName
}

// WithMsgVpnName adds the msgVpnName to the delete msg vpn client profile params
func (o *DeleteMsgVpnClientProfileParams) WithMsgVpnName(msgVpnName string) *DeleteMsgVpnClientProfileParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the delete msg vpn client profile params
func (o *DeleteMsgVpnClientProfileParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMsgVpnClientProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clientProfileName
	if err := r.SetPathParam("clientProfileName", o.ClientProfileName); err != nil {
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
