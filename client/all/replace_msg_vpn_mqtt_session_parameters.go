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
	"github.com/go-openapi/swag"

	"github.com/PatrickDelancy/semp-client/models"
)

// NewReplaceMsgVpnMqttSessionParams creates a new ReplaceMsgVpnMqttSessionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceMsgVpnMqttSessionParams() *ReplaceMsgVpnMqttSessionParams {
	return &ReplaceMsgVpnMqttSessionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceMsgVpnMqttSessionParamsWithTimeout creates a new ReplaceMsgVpnMqttSessionParams object
// with the ability to set a timeout on a request.
func NewReplaceMsgVpnMqttSessionParamsWithTimeout(timeout time.Duration) *ReplaceMsgVpnMqttSessionParams {
	return &ReplaceMsgVpnMqttSessionParams{
		timeout: timeout,
	}
}

// NewReplaceMsgVpnMqttSessionParamsWithContext creates a new ReplaceMsgVpnMqttSessionParams object
// with the ability to set a context for a request.
func NewReplaceMsgVpnMqttSessionParamsWithContext(ctx context.Context) *ReplaceMsgVpnMqttSessionParams {
	return &ReplaceMsgVpnMqttSessionParams{
		Context: ctx,
	}
}

// NewReplaceMsgVpnMqttSessionParamsWithHTTPClient creates a new ReplaceMsgVpnMqttSessionParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceMsgVpnMqttSessionParamsWithHTTPClient(client *http.Client) *ReplaceMsgVpnMqttSessionParams {
	return &ReplaceMsgVpnMqttSessionParams{
		HTTPClient: client,
	}
}

/* ReplaceMsgVpnMqttSessionParams contains all the parameters to send to the API endpoint
   for the replace msg vpn mqtt session operation.

   Typically these are written to a http.Request.
*/
type ReplaceMsgVpnMqttSessionParams struct {

	/* Body.

	   The MQTT Session object's attributes.
	*/
	Body *models.MsgVpnMqttSession

	/* MqttSessionClientID.

	   The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.
	*/
	MqttSessionClientID string

	/* MqttSessionVirtualRouter.

	   The virtual router of the MQTT Session.
	*/
	MqttSessionVirtualRouter string

	/* MsgVpnName.

	   The name of the Message VPN.
	*/
	MsgVpnName string

	/* OpaquePassword.

	   Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter.
	*/
	OpaquePassword *string

	/* Select.

	   Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.
	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the replace msg vpn mqtt session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceMsgVpnMqttSessionParams) WithDefaults() *ReplaceMsgVpnMqttSessionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace msg vpn mqtt session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceMsgVpnMqttSessionParams) SetDefaults() {
	// no default values defined for this parameter
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

// WithOpaquePassword adds the opaquePassword to the replace msg vpn mqtt session params
func (o *ReplaceMsgVpnMqttSessionParams) WithOpaquePassword(opaquePassword *string) *ReplaceMsgVpnMqttSessionParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the replace msg vpn mqtt session params
func (o *ReplaceMsgVpnMqttSessionParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
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

	if o.OpaquePassword != nil {

		// query param opaquePassword
		var qrOpaquePassword string

		if o.OpaquePassword != nil {
			qrOpaquePassword = *o.OpaquePassword
		}
		qOpaquePassword := qrOpaquePassword
		if qOpaquePassword != "" {

			if err := r.SetQueryParam("opaquePassword", qOpaquePassword); err != nil {
				return err
			}
		}
	}

	if o.Select != nil {

		// binding items for select
		joinedSelect := o.bindParamSelect(reg)

		// query array param select
		if err := r.SetQueryParam("select", joinedSelect...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamReplaceMsgVpnMqttSession binds the parameter select
func (o *ReplaceMsgVpnMqttSessionParams) bindParamSelect(formats strfmt.Registry) []string {
	selectIR := o.Select

	var selectIC []string
	for _, selectIIR := range selectIR { // explode []string

		selectIIV := selectIIR // string as string
		selectIC = append(selectIC, selectIIV)
	}

	// items.CollectionFormat: "csv"
	selectIS := swag.JoinByFormat(selectIC, "csv")

	return selectIS
}
