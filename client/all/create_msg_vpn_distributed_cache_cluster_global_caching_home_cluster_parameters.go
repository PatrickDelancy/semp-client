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

// NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams creates a new CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams() *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	return &CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParamsWithTimeout creates a new CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams object
// with the ability to set a timeout on a request.
func NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParamsWithTimeout(timeout time.Duration) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	return &CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams{
		timeout: timeout,
	}
}

// NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParamsWithContext creates a new CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams object
// with the ability to set a context for a request.
func NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParamsWithContext(ctx context.Context) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	return &CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams{
		Context: ctx,
	}
}

// NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParamsWithHTTPClient creates a new CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParamsWithHTTPClient(client *http.Client) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	return &CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams{
		HTTPClient: client,
	}
}

/* CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams contains all the parameters to send to the API endpoint
   for the create msg vpn distributed cache cluster global caching home cluster operation.

   Typically these are written to a http.Request.
*/
type CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams struct {

	/* Body.

	   The Home Cache Cluster object's attributes.
	*/
	Body *models.MsgVpnDistributedCacheClusterGlobalCachingHomeCluster

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

// WithDefaults hydrates default values in the create msg vpn distributed cache cluster global caching home cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithDefaults() *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create msg vpn distributed cache cluster global caching home cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create msg vpn distributed cache cluster global caching home cluster params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithTimeout(timeout time.Duration) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create msg vpn distributed cache cluster global caching home cluster params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create msg vpn distributed cache cluster global caching home cluster params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithContext(ctx context.Context) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create msg vpn distributed cache cluster global caching home cluster params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create msg vpn distributed cache cluster global caching home cluster params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithHTTPClient(client *http.Client) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create msg vpn distributed cache cluster global caching home cluster params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create msg vpn distributed cache cluster global caching home cluster params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithBody(body *models.MsgVpnDistributedCacheClusterGlobalCachingHomeCluster) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create msg vpn distributed cache cluster global caching home cluster params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetBody(body *models.MsgVpnDistributedCacheClusterGlobalCachingHomeCluster) {
	o.Body = body
}

// WithCacheName adds the cacheName to the create msg vpn distributed cache cluster global caching home cluster params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithCacheName(cacheName string) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the create msg vpn distributed cache cluster global caching home cluster params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithClusterName adds the clusterName to the create msg vpn distributed cache cluster global caching home cluster params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithClusterName(clusterName string) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the create msg vpn distributed cache cluster global caching home cluster params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WithMsgVpnName adds the msgVpnName to the create msg vpn distributed cache cluster global caching home cluster params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithMsgVpnName(msgVpnName string) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the create msg vpn distributed cache cluster global caching home cluster params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithOpaquePassword adds the opaquePassword to the create msg vpn distributed cache cluster global caching home cluster params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithOpaquePassword(opaquePassword *string) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the create msg vpn distributed cache cluster global caching home cluster params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the create msg vpn distributed cache cluster global caching home cluster params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithSelect(selectVar []string) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the create msg vpn distributed cache cluster global caching home cluster params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamCreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster binds the parameter select
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) bindParamSelect(formats strfmt.Registry) []string {
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