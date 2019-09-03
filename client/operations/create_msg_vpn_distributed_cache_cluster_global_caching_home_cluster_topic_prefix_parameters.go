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

// NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams creates a new CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams object
// with the default values initialized.
func NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams() *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	var ()
	return &CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParamsWithTimeout creates a new CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParamsWithTimeout(timeout time.Duration) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	var ()
	return &CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams{

		timeout: timeout,
	}
}

// NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParamsWithContext creates a new CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParamsWithContext(ctx context.Context) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	var ()
	return &CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams{

		Context: ctx,
	}
}

// NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParamsWithHTTPClient creates a new CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParamsWithHTTPClient(client *http.Client) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	var ()
	return &CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams{
		HTTPClient: client,
	}
}

/*CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams contains all the parameters to send to the API endpoint
for the create msg vpn distributed cache cluster global caching home cluster topic prefix operation typically these are written to a http.Request
*/
type CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams struct {

	/*Body
	  The Topic Prefix object's attributes.

	*/
	Body *models.MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix
	/*CacheName
	  The name of the Distributed Cache.

	*/
	CacheName string
	/*ClusterName
	  The name of the Cache Cluster.

	*/
	ClusterName string
	/*HomeClusterName
	  The name of the remote Home Cache Cluster.

	*/
	HomeClusterName string
	/*MsgVpnName
	  The name of the Message VPN.

	*/
	MsgVpnName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithTimeout(timeout time.Duration) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithContext(ctx context.Context) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithHTTPClient(client *http.Client) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithBody(body *models.MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetBody(body *models.MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix) {
	o.Body = body
}

// WithCacheName adds the cacheName to the create msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithCacheName(cacheName string) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the create msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithClusterName adds the clusterName to the create msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithClusterName(clusterName string) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the create msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WithHomeClusterName adds the homeClusterName to the create msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithHomeClusterName(homeClusterName string) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetHomeClusterName(homeClusterName)
	return o
}

// SetHomeClusterName adds the homeClusterName to the create msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetHomeClusterName(homeClusterName string) {
	o.HomeClusterName = homeClusterName
}

// WithMsgVpnName adds the msgVpnName to the create msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithMsgVpnName(msgVpnName string) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the create msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the create msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithSelect(selectVar []string) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the create msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param homeClusterName
	if err := r.SetPathParam("homeClusterName", o.HomeClusterName); err != nil {
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