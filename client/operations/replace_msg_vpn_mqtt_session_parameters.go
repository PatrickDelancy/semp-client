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

// NewReplaceMsgVpnMqttSessionParams creates a new ReplaceMsgVpnMqttSessionParams object
// with the default values initialized.
func NewReplaceMsgVpnMqttSessionParams() *ReplaceMsgVpnMqttSessionParams {
	var ()
	return &ReplaceMsgVpnMqttSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceMsgVpnMqttSessionParamsWithTimeout creates a new ReplaceMsgVpnMqttSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceMsgVpnMqttSessionParamsWithTimeout(timeout time.Duration) *ReplaceMsgVpnMqttSessionParams {
	var ()
	return &ReplaceMsgVpnMqttSessionParams{

		timeout: timeout,
	}
}

// NewReplaceMsgVpnMqttSessionParamsWithContext creates a new ReplaceMsgVpnMqttSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceMsgVpnMqttSessionParamsWithContext(ctx context.Context) *ReplaceMsgVpnMqttSessionParams {
	var ()
	return &ReplaceMsgVpnMqttSessionParams{

		Context: ctx,
	}
}

// NewReplaceMsgVpnMqttSessionParamsWithHTTPClient creates a new ReplaceMsgVpnMqttSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceMsgVpnMqttSessionParamsWithHTTPClient(client *http.Client) *ReplaceMsgVpnMqttSessionParams {
	var ()
	return &ReplaceMsgVpnMqttSessionParams{
		HTTPClient: client,
	}
}

/*ReplaceMsgVpnMqttSessionParams contains all the parameters to send to the API endpoint
for the replace msg vpn mqtt session operation typically these are written to a http.Request
*/
type ReplaceMsgVpnMqttSessionParams struct {

	/*Body
	  The MQTT Session object's attributes.

	*/
	Body *models.MsgVpnMqttSession
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
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See [Select](#select "Description of the syntax of the `select` parameter").

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace msg vpn mqtt session params
func (o *ReplaceMsgVpnMqttSessionParams) WithTimeout(timeout time.Duration) *ReplaceMsgVpnMqttSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace msg vpn mqtt session params
func (o *ReplaceMsgVpnMqttSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace msg vpn mqtt session params
func (o *ReplaceMsgVpnMqttSessionParams) WithContext(ctx context.Context) *ReplaceMsgVpnMqttSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace msg vpn mqtt session params
func (o *ReplaceMsgVpnMqttSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace msg vpn mqtt session params
func (o *ReplaceMsgVpnMqttSessionParams) WithHTTPClient(client *http.Client) *ReplaceMsgVpnMqttSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace msg vpn mqtt session params
func (o *ReplaceMsgVpnMqttSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace msg vpn mqtt session params
func (o *ReplaceMsgVpnMqttSessionParams) WithBody(body *models.MsgVpnMqttSession) *ReplaceMsgVpnMqttSessionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace msg vpn mqtt session params
func (o *ReplaceMsgVpnMqttSessionParams) SetBody(body *models.MsgVpnMqttSession) {
	o.Body = body
}

// WithMqttSessionClientID adds the mqttSessionClientID to the replace msg vpn mqtt session params
func (o *ReplaceMsgVpnMqttSessionParams) WithMqttSessionClientID(mqttSessionClientID string) *ReplaceMsgVpnMqttSessionParams {
	o.SetMqttSessionClientID(mqttSessionClientID)
	return o
}

// SetMqttSessionClientID adds the mqttSessionClientId to the replace msg vpn mqtt session params
func (o *ReplaceMsgVpnMqttSessionParams) SetMqttSessionClientID(mqttSessionClientID string) {
	o.MqttSessionClientID = mqttSessionClientID
}

// WithMqttSessionVirtualRouter adds the mqttSessionVirtualRouter to the replace msg vpn mqtt session params
func (o *ReplaceMsgVpnMqttSessionParams) WithMqttSessionVirtualRouter(mqttSessionVirtualRouter string) *ReplaceMsgVpnMqttSessionParams {
	o.SetMqttSessionVirtualRouter(mqttSessionVirtualRouter)
	return o
}

// SetMqttSessionVirtualRouter adds the mqttSessionVirtualRouter to the replace msg vpn mqtt session params
func (o *ReplaceMsgVpnMqttSessionParams) SetMqttSessionVirtualRouter(mqttSessionVirtualRouter string) {
	o.MqttSessionVirtualRouter = mqttSessionVirtualRouter
}

// WithMsgVpnName adds the msgVpnName to the replace msg vpn mqtt session params
func (o *ReplaceMsgVpnMqttSessionParams) WithMsgVpnName(msgVpnName string) *ReplaceMsgVpnMqttSessionParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the replace msg vpn mqtt session params
func (o *ReplaceMsgVpnMqttSessionParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the replace msg vpn mqtt session params
func (o *ReplaceMsgVpnMqttSessionParams) WithSelect(selectVar []string) *ReplaceMsgVpnMqttSessionParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the replace msg vpn mqtt session params
func (o *ReplaceMsgVpnMqttSessionParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceMsgVpnMqttSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
