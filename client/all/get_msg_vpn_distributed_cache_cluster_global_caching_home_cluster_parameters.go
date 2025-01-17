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

// NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams creates a new GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams() *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	return &GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParamsWithTimeout creates a new GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams object
// with the ability to set a timeout on a request.
func NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParamsWithTimeout(timeout time.Duration) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	return &GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams{
		timeout: timeout,
	}
}

// NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParamsWithContext creates a new GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams object
// with the ability to set a context for a request.
func NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParamsWithContext(ctx context.Context) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	return &GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams{
		Context: ctx,
	}
}

// NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParamsWithHTTPClient creates a new GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParamsWithHTTPClient(client *http.Client) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	return &GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams{
		HTTPClient: client,
	}
}

/* GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams contains all the parameters to send to the API endpoint
   for the get msg vpn distributed cache cluster global caching home cluster operation.

   Typically these are written to a http.Request.
*/
type GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams struct {

	/* CacheName.

	   The name of the Distributed Cache.
	*/
	CacheName string

	/* ClusterName.

	   The name of the Cache Cluster.
	*/
	ClusterName string

	/* HomeClusterName.

	   The name of the remote Home Cache Cluster.
	*/
	HomeClusterName string

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

// WithDefaults hydrates default values in the get msg vpn distributed cache cluster global caching home cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithDefaults() *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get msg vpn distributed cache cluster global caching home cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get msg vpn distributed cache cluster global caching home cluster params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithTimeout(timeout time.Duration) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn distributed cache cluster global caching home cluster params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn distributed cache cluster global caching home cluster params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithContext(ctx context.Context) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn distributed cache cluster global caching home cluster params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn distributed cache cluster global caching home cluster params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithHTTPClient(client *http.Client) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn distributed cache cluster global caching home cluster params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCacheName adds the cacheName to the get msg vpn distributed cache cluster global caching home cluster params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithCacheName(cacheName string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the get msg vpn distributed cache cluster global caching home cluster params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithClusterName adds the clusterName to the get msg vpn distributed cache cluster global caching home cluster params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithClusterName(clusterName string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the get msg vpn distributed cache cluster global caching home cluster params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WithHomeClusterName adds the homeClusterName to the get msg vpn distributed cache cluster global caching home cluster params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithHomeClusterName(homeClusterName string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetHomeClusterName(homeClusterName)
	return o
}

// SetHomeClusterName adds the homeClusterName to the get msg vpn distributed cache cluster global caching home cluster params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetHomeClusterName(homeClusterName string) {
	o.HomeClusterName = homeClusterName
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn distributed cache cluster global caching home cluster params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn distributed cache cluster global caching home cluster params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithOpaquePassword adds the opaquePassword to the get msg vpn distributed cache cluster global caching home cluster params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithOpaquePassword(opaquePassword *string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the get msg vpn distributed cache cluster global caching home cluster params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the get msg vpn distributed cache cluster global caching home cluster params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithSelect(selectVar []string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn distributed cache cluster global caching home cluster params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cacheName
	if err := r.SetPathParam("cacheName", o.CacheName); err != nil {
		return err
	}

	// path param clusterName
	if err := r.SetPathParam("clusterName", o.ClusterName); err != nil {
		return err
	}

	// path param homeClusterName
	if err := r.SetPathParam("homeClusterName", o.HomeClusterName); err != nil {
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

// bindParamGetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster binds the parameter select
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) bindParamSelect(formats strfmt.Registry) []string {
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
