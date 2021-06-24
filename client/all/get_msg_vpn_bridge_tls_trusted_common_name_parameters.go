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
)

// NewGetMsgVpnBridgeTLSTrustedCommonNameParams creates a new GetMsgVpnBridgeTLSTrustedCommonNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMsgVpnBridgeTLSTrustedCommonNameParams() *GetMsgVpnBridgeTLSTrustedCommonNameParams {
	return &GetMsgVpnBridgeTLSTrustedCommonNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnBridgeTLSTrustedCommonNameParamsWithTimeout creates a new GetMsgVpnBridgeTLSTrustedCommonNameParams object
// with the ability to set a timeout on a request.
func NewGetMsgVpnBridgeTLSTrustedCommonNameParamsWithTimeout(timeout time.Duration) *GetMsgVpnBridgeTLSTrustedCommonNameParams {
	return &GetMsgVpnBridgeTLSTrustedCommonNameParams{
		timeout: timeout,
	}
}

// NewGetMsgVpnBridgeTLSTrustedCommonNameParamsWithContext creates a new GetMsgVpnBridgeTLSTrustedCommonNameParams object
// with the ability to set a context for a request.
func NewGetMsgVpnBridgeTLSTrustedCommonNameParamsWithContext(ctx context.Context) *GetMsgVpnBridgeTLSTrustedCommonNameParams {
	return &GetMsgVpnBridgeTLSTrustedCommonNameParams{
		Context: ctx,
	}
}

// NewGetMsgVpnBridgeTLSTrustedCommonNameParamsWithHTTPClient creates a new GetMsgVpnBridgeTLSTrustedCommonNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMsgVpnBridgeTLSTrustedCommonNameParamsWithHTTPClient(client *http.Client) *GetMsgVpnBridgeTLSTrustedCommonNameParams {
	return &GetMsgVpnBridgeTLSTrustedCommonNameParams{
		HTTPClient: client,
	}
}

/* GetMsgVpnBridgeTLSTrustedCommonNameParams contains all the parameters to send to the API endpoint
   for the get msg vpn bridge Tls trusted common name operation.

   Typically these are written to a http.Request.
*/
type GetMsgVpnBridgeTLSTrustedCommonNameParams struct {

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

	/* OpaquePassword.

	   Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter.
	*/
	OpaquePassword *string

	/* Select.

	   Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.
	*/
	Select []string

	/* TLSTrustedCommonName.

	   The expected trusted common name of the remote certificate.
	*/
	TLSTrustedCommonName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get msg vpn bridge Tls trusted common name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) WithDefaults() *GetMsgVpnBridgeTLSTrustedCommonNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get msg vpn bridge Tls trusted common name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get msg vpn bridge Tls trusted common name params
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) WithTimeout(timeout time.Duration) *GetMsgVpnBridgeTLSTrustedCommonNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn bridge Tls trusted common name params
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn bridge Tls trusted common name params
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) WithContext(ctx context.Context) *GetMsgVpnBridgeTLSTrustedCommonNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn bridge Tls trusted common name params
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn bridge Tls trusted common name params
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) WithHTTPClient(client *http.Client) *GetMsgVpnBridgeTLSTrustedCommonNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn bridge Tls trusted common name params
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBridgeName adds the bridgeName to the get msg vpn bridge Tls trusted common name params
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) WithBridgeName(bridgeName string) *GetMsgVpnBridgeTLSTrustedCommonNameParams {
	o.SetBridgeName(bridgeName)
	return o
}

// SetBridgeName adds the bridgeName to the get msg vpn bridge Tls trusted common name params
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) SetBridgeName(bridgeName string) {
	o.BridgeName = bridgeName
}

// WithBridgeVirtualRouter adds the bridgeVirtualRouter to the get msg vpn bridge Tls trusted common name params
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) WithBridgeVirtualRouter(bridgeVirtualRouter string) *GetMsgVpnBridgeTLSTrustedCommonNameParams {
	o.SetBridgeVirtualRouter(bridgeVirtualRouter)
	return o
}

// SetBridgeVirtualRouter adds the bridgeVirtualRouter to the get msg vpn bridge Tls trusted common name params
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) SetBridgeVirtualRouter(bridgeVirtualRouter string) {
	o.BridgeVirtualRouter = bridgeVirtualRouter
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn bridge Tls trusted common name params
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnBridgeTLSTrustedCommonNameParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn bridge Tls trusted common name params
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithOpaquePassword adds the opaquePassword to the get msg vpn bridge Tls trusted common name params
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) WithOpaquePassword(opaquePassword *string) *GetMsgVpnBridgeTLSTrustedCommonNameParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the get msg vpn bridge Tls trusted common name params
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the get msg vpn bridge Tls trusted common name params
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) WithSelect(selectVar []string) *GetMsgVpnBridgeTLSTrustedCommonNameParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn bridge Tls trusted common name params
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithTLSTrustedCommonName adds the tLSTrustedCommonName to the get msg vpn bridge Tls trusted common name params
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) WithTLSTrustedCommonName(tLSTrustedCommonName string) *GetMsgVpnBridgeTLSTrustedCommonNameParams {
	o.SetTLSTrustedCommonName(tLSTrustedCommonName)
	return o
}

// SetTLSTrustedCommonName adds the tlsTrustedCommonName to the get msg vpn bridge Tls trusted common name params
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) SetTLSTrustedCommonName(tLSTrustedCommonName string) {
	o.TLSTrustedCommonName = tLSTrustedCommonName
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tlsTrustedCommonName
	if err := r.SetPathParam("tlsTrustedCommonName", o.TLSTrustedCommonName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetMsgVpnBridgeTLSTrustedCommonName binds the parameter select
func (o *GetMsgVpnBridgeTLSTrustedCommonNameParams) bindParamSelect(formats strfmt.Registry) []string {
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