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

// NewUpdateMsgVpnDistributedCacheClusterParams creates a new UpdateMsgVpnDistributedCacheClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMsgVpnDistributedCacheClusterParams() *UpdateMsgVpnDistributedCacheClusterParams {
	return &UpdateMsgVpnDistributedCacheClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMsgVpnDistributedCacheClusterParamsWithTimeout creates a new UpdateMsgVpnDistributedCacheClusterParams object
// with the ability to set a timeout on a request.
func NewUpdateMsgVpnDistributedCacheClusterParamsWithTimeout(timeout time.Duration) *UpdateMsgVpnDistributedCacheClusterParams {
	return &UpdateMsgVpnDistributedCacheClusterParams{
		timeout: timeout,
	}
}

// NewUpdateMsgVpnDistributedCacheClusterParamsWithContext creates a new UpdateMsgVpnDistributedCacheClusterParams object
// with the ability to set a context for a request.
func NewUpdateMsgVpnDistributedCacheClusterParamsWithContext(ctx context.Context) *UpdateMsgVpnDistributedCacheClusterParams {
	return &UpdateMsgVpnDistributedCacheClusterParams{
		Context: ctx,
	}
}

// NewUpdateMsgVpnDistributedCacheClusterParamsWithHTTPClient creates a new UpdateMsgVpnDistributedCacheClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMsgVpnDistributedCacheClusterParamsWithHTTPClient(client *http.Client) *UpdateMsgVpnDistributedCacheClusterParams {
	return &UpdateMsgVpnDistributedCacheClusterParams{
		HTTPClient: client,
	}
}

/* UpdateMsgVpnDistributedCacheClusterParams contains all the parameters to send to the API endpoint
   for the update msg vpn distributed cache cluster operation.

   Typically these are written to a http.Request.
*/
type UpdateMsgVpnDistributedCacheClusterParams struct {

	/* Body.

	   The Cache Cluster object's attributes.
	*/
	Body *models.MsgVpnDistributedCacheCluster

	/* CacheName.

	   The name of the Distributed Cache.
	*/
	CacheName string

	/* ClusterName.

	   The name of the Cache Cluster.
	*/
	ClusterName string

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

// WithDefaults hydrates default values in the update msg vpn distributed cache cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMsgVpnDistributedCacheClusterParams) WithDefaults() *UpdateMsgVpnDistributedCacheClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update msg vpn distributed cache cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMsgVpnDistributedCacheClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) WithTimeout(timeout time.Duration) *UpdateMsgVpnDistributedCacheClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) WithContext(ctx context.Context) *UpdateMsgVpnDistributedCacheClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) WithHTTPClient(client *http.Client) *UpdateMsgVpnDistributedCacheClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) WithBody(body *models.MsgVpnDistributedCacheCluster) *UpdateMsgVpnDistributedCacheClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) SetBody(body *models.MsgVpnDistributedCacheCluster) {
	o.Body = body
}

// WithCacheName adds the cacheName to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) WithCacheName(cacheName string) *UpdateMsgVpnDistributedCacheClusterParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithClusterName adds the clusterName to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) WithClusterName(clusterName string) *UpdateMsgVpnDistributedCacheClusterParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WithMsgVpnName adds the msgVpnName to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) WithMsgVpnName(msgVpnName string) *UpdateMsgVpnDistributedCacheClusterParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithOpaquePassword adds the opaquePassword to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) WithOpaquePassword(opaquePassword *string) *UpdateMsgVpnDistributedCacheClusterParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) WithSelect(selectVar []string) *UpdateMsgVpnDistributedCacheClusterParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMsgVpnDistributedCacheClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cacheName
	if err := r.SetPathParam("cacheName", o.CacheName); err != nil {
		return err
	}

	// path param clusterName
	if err := r.SetPathParam("clusterName", o.ClusterName); err != nil {
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

// bindParamUpdateMsgVpnDistributedCacheCluster binds the parameter select
func (o *UpdateMsgVpnDistributedCacheClusterParams) bindParamSelect(formats strfmt.Registry) []string {
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