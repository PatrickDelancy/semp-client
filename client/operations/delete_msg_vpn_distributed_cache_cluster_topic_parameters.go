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

// NewDeleteMsgVpnDistributedCacheClusterTopicParams creates a new DeleteMsgVpnDistributedCacheClusterTopicParams object
// with the default values initialized.
func NewDeleteMsgVpnDistributedCacheClusterTopicParams() *DeleteMsgVpnDistributedCacheClusterTopicParams {
	var ()
	return &DeleteMsgVpnDistributedCacheClusterTopicParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMsgVpnDistributedCacheClusterTopicParamsWithTimeout creates a new DeleteMsgVpnDistributedCacheClusterTopicParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMsgVpnDistributedCacheClusterTopicParamsWithTimeout(timeout time.Duration) *DeleteMsgVpnDistributedCacheClusterTopicParams {
	var ()
	return &DeleteMsgVpnDistributedCacheClusterTopicParams{

		timeout: timeout,
	}
}

// NewDeleteMsgVpnDistributedCacheClusterTopicParamsWithContext creates a new DeleteMsgVpnDistributedCacheClusterTopicParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMsgVpnDistributedCacheClusterTopicParamsWithContext(ctx context.Context) *DeleteMsgVpnDistributedCacheClusterTopicParams {
	var ()
	return &DeleteMsgVpnDistributedCacheClusterTopicParams{

		Context: ctx,
	}
}

// NewDeleteMsgVpnDistributedCacheClusterTopicParamsWithHTTPClient creates a new DeleteMsgVpnDistributedCacheClusterTopicParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMsgVpnDistributedCacheClusterTopicParamsWithHTTPClient(client *http.Client) *DeleteMsgVpnDistributedCacheClusterTopicParams {
	var ()
	return &DeleteMsgVpnDistributedCacheClusterTopicParams{
		HTTPClient: client,
	}
}

/*DeleteMsgVpnDistributedCacheClusterTopicParams contains all the parameters to send to the API endpoint
for the delete msg vpn distributed cache cluster topic operation typically these are written to a http.Request
*/
type DeleteMsgVpnDistributedCacheClusterTopicParams struct {

	/*CacheName
	  The name of the Distributed Cache.

	*/
	CacheName string
	/*ClusterName
	  The name of the Cache Cluster.

	*/
	ClusterName string
	/*MsgVpnName
	  The name of the Message VPN.

	*/
	MsgVpnName string
	/*Topic
	  The value of the Topic in the form a/b/c.

	*/
	Topic string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete msg vpn distributed cache cluster topic params
func (o *DeleteMsgVpnDistributedCacheClusterTopicParams) WithTimeout(timeout time.Duration) *DeleteMsgVpnDistributedCacheClusterTopicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete msg vpn distributed cache cluster topic params
func (o *DeleteMsgVpnDistributedCacheClusterTopicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete msg vpn distributed cache cluster topic params
func (o *DeleteMsgVpnDistributedCacheClusterTopicParams) WithContext(ctx context.Context) *DeleteMsgVpnDistributedCacheClusterTopicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete msg vpn distributed cache cluster topic params
func (o *DeleteMsgVpnDistributedCacheClusterTopicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete msg vpn distributed cache cluster topic params
func (o *DeleteMsgVpnDistributedCacheClusterTopicParams) WithHTTPClient(client *http.Client) *DeleteMsgVpnDistributedCacheClusterTopicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete msg vpn distributed cache cluster topic params
func (o *DeleteMsgVpnDistributedCacheClusterTopicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCacheName adds the cacheName to the delete msg vpn distributed cache cluster topic params
func (o *DeleteMsgVpnDistributedCacheClusterTopicParams) WithCacheName(cacheName string) *DeleteMsgVpnDistributedCacheClusterTopicParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the delete msg vpn distributed cache cluster topic params
func (o *DeleteMsgVpnDistributedCacheClusterTopicParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithClusterName adds the clusterName to the delete msg vpn distributed cache cluster topic params
func (o *DeleteMsgVpnDistributedCacheClusterTopicParams) WithClusterName(clusterName string) *DeleteMsgVpnDistributedCacheClusterTopicParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the delete msg vpn distributed cache cluster topic params
func (o *DeleteMsgVpnDistributedCacheClusterTopicParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WithMsgVpnName adds the msgVpnName to the delete msg vpn distributed cache cluster topic params
func (o *DeleteMsgVpnDistributedCacheClusterTopicParams) WithMsgVpnName(msgVpnName string) *DeleteMsgVpnDistributedCacheClusterTopicParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the delete msg vpn distributed cache cluster topic params
func (o *DeleteMsgVpnDistributedCacheClusterTopicParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithTopic adds the topic to the delete msg vpn distributed cache cluster topic params
func (o *DeleteMsgVpnDistributedCacheClusterTopicParams) WithTopic(topic string) *DeleteMsgVpnDistributedCacheClusterTopicParams {
	o.SetTopic(topic)
	return o
}

// SetTopic adds the topic to the delete msg vpn distributed cache cluster topic params
func (o *DeleteMsgVpnDistributedCacheClusterTopicParams) SetTopic(topic string) {
	o.Topic = topic
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMsgVpnDistributedCacheClusterTopicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	// path param topic
	if err := r.SetPathParam("topic", o.Topic); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}