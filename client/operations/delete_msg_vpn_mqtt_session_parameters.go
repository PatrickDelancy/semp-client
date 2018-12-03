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

// NewDeleteMsgVpnMqttSessionParams creates a new DeleteMsgVpnMqttSessionParams object
// with the default values initialized.
func NewDeleteMsgVpnMqttSessionParams() *DeleteMsgVpnMqttSessionParams {
	var ()
	return &DeleteMsgVpnMqttSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMsgVpnMqttSessionParamsWithTimeout creates a new DeleteMsgVpnMqttSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMsgVpnMqttSessionParamsWithTimeout(timeout time.Duration) *DeleteMsgVpnMqttSessionParams {
	var ()
	return &DeleteMsgVpnMqttSessionParams{

		timeout: timeout,
	}
}

// NewDeleteMsgVpnMqttSessionParamsWithContext creates a new DeleteMsgVpnMqttSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMsgVpnMqttSessionParamsWithContext(ctx context.Context) *DeleteMsgVpnMqttSessionParams {
	var ()
	return &DeleteMsgVpnMqttSessionParams{

		Context: ctx,
	}
}

// NewDeleteMsgVpnMqttSessionParamsWithHTTPClient creates a new DeleteMsgVpnMqttSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMsgVpnMqttSessionParamsWithHTTPClient(client *http.Client) *DeleteMsgVpnMqttSessionParams {
	var ()
	return &DeleteMsgVpnMqttSessionParams{
		HTTPClient: client,
	}
}

/*DeleteMsgVpnMqttSessionParams contains all the parameters to send to the API endpoint
for the delete msg vpn mqtt session operation typically these are written to a http.Request
*/
type DeleteMsgVpnMqttSessionParams struct {

	/*MqttSessionClientID
	  The mqttSessionClientId of the MQTT Session.

	*/
	MqttSessionClientID string
	/*MqttSessionVirtualRouter
	  The mqttSessionVirtualRouter of the MQTT Session.

	*/
	MqttSessionVirtualRouter string
	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete msg vpn mqtt session params
func (o *DeleteMsgVpnMqttSessionParams) WithTimeout(timeout time.Duration) *DeleteMsgVpnMqttSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete msg vpn mqtt session params
func (o *DeleteMsgVpnMqttSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete msg vpn mqtt session params
func (o *DeleteMsgVpnMqttSessionParams) WithContext(ctx context.Context) *DeleteMsgVpnMqttSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete msg vpn mqtt session params
func (o *DeleteMsgVpnMqttSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete msg vpn mqtt session params
func (o *DeleteMsgVpnMqttSessionParams) WithHTTPClient(client *http.Client) *DeleteMsgVpnMqttSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete msg vpn mqtt session params
func (o *DeleteMsgVpnMqttSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMqttSessionClientID adds the mqttSessionClientID to the delete msg vpn mqtt session params
func (o *DeleteMsgVpnMqttSessionParams) WithMqttSessionClientID(mqttSessionClientID string) *DeleteMsgVpnMqttSessionParams {
	o.SetMqttSessionClientID(mqttSessionClientID)
	return o
}

// SetMqttSessionClientID adds the mqttSessionClientId to the delete msg vpn mqtt session params
func (o *DeleteMsgVpnMqttSessionParams) SetMqttSessionClientID(mqttSessionClientID string) {
	o.MqttSessionClientID = mqttSessionClientID
}

// WithMqttSessionVirtualRouter adds the mqttSessionVirtualRouter to the delete msg vpn mqtt session params
func (o *DeleteMsgVpnMqttSessionParams) WithMqttSessionVirtualRouter(mqttSessionVirtualRouter string) *DeleteMsgVpnMqttSessionParams {
	o.SetMqttSessionVirtualRouter(mqttSessionVirtualRouter)
	return o
}

// SetMqttSessionVirtualRouter adds the mqttSessionVirtualRouter to the delete msg vpn mqtt session params
func (o *DeleteMsgVpnMqttSessionParams) SetMqttSessionVirtualRouter(mqttSessionVirtualRouter string) {
	o.MqttSessionVirtualRouter = mqttSessionVirtualRouter
}

// WithMsgVpnName adds the msgVpnName to the delete msg vpn mqtt session params
func (o *DeleteMsgVpnMqttSessionParams) WithMsgVpnName(msgVpnName string) *DeleteMsgVpnMqttSessionParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the delete msg vpn mqtt session params
func (o *DeleteMsgVpnMqttSessionParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMsgVpnMqttSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mqttSessionClientId
	if err := r.SetPathParam("mqttSessionClientId", o.MqttSessionClientID); err != nil {
		return err
	}

	// path param mqttSessionVirtualRouter
	if err := r.SetPathParam("mqttSessionVirtualRouter", o.MqttSessionVirtualRouter); err != nil {
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