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

// NewUpdateMsgVpnDistributedCacheParams creates a new UpdateMsgVpnDistributedCacheParams object
// with the default values initialized.
func NewUpdateMsgVpnDistributedCacheParams() *UpdateMsgVpnDistributedCacheParams {
	var ()
	return &UpdateMsgVpnDistributedCacheParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMsgVpnDistributedCacheParamsWithTimeout creates a new UpdateMsgVpnDistributedCacheParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateMsgVpnDistributedCacheParamsWithTimeout(timeout time.Duration) *UpdateMsgVpnDistributedCacheParams {
	var ()
	return &UpdateMsgVpnDistributedCacheParams{

		timeout: timeout,
	}
}

// NewUpdateMsgVpnDistributedCacheParamsWithContext creates a new UpdateMsgVpnDistributedCacheParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateMsgVpnDistributedCacheParamsWithContext(ctx context.Context) *UpdateMsgVpnDistributedCacheParams {
	var ()
	return &UpdateMsgVpnDistributedCacheParams{

		Context: ctx,
	}
}

// NewUpdateMsgVpnDistributedCacheParamsWithHTTPClient creates a new UpdateMsgVpnDistributedCacheParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateMsgVpnDistributedCacheParamsWithHTTPClient(client *http.Client) *UpdateMsgVpnDistributedCacheParams {
	var ()
	return &UpdateMsgVpnDistributedCacheParams{
		HTTPClient: client,
	}
}

/*UpdateMsgVpnDistributedCacheParams contains all the parameters to send to the API endpoint
for the update msg vpn distributed cache operation typically these are written to a http.Request
*/
type UpdateMsgVpnDistributedCacheParams struct {

	/*Body
	  The Distributed Cache object's attributes.

	*/
	Body *models.MsgVpnDistributedCache
	/*CacheName
	  The name of the Distributed Cache.

	*/
	CacheName string
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

// WithTimeout adds the timeout to the update msg vpn distributed cache params
func (o *UpdateMsgVpnDistributedCacheParams) WithTimeout(timeout time.Duration) *UpdateMsgVpnDistributedCacheParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update msg vpn distributed cache params
func (o *UpdateMsgVpnDistributedCacheParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update msg vpn distributed cache params
func (o *UpdateMsgVpnDistributedCacheParams) WithContext(ctx context.Context) *UpdateMsgVpnDistributedCacheParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update msg vpn distributed cache params
func (o *UpdateMsgVpnDistributedCacheParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update msg vpn distributed cache params
func (o *UpdateMsgVpnDistributedCacheParams) WithHTTPClient(client *http.Client) *UpdateMsgVpnDistributedCacheParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update msg vpn distributed cache params
func (o *UpdateMsgVpnDistributedCacheParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update msg vpn distributed cache params
func (o *UpdateMsgVpnDistributedCacheParams) WithBody(body *models.MsgVpnDistributedCache) *UpdateMsgVpnDistributedCacheParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update msg vpn distributed cache params
func (o *UpdateMsgVpnDistributedCacheParams) SetBody(body *models.MsgVpnDistributedCache) {
	o.Body = body
}

// WithCacheName adds the cacheName to the update msg vpn distributed cache params
func (o *UpdateMsgVpnDistributedCacheParams) WithCacheName(cacheName string) *UpdateMsgVpnDistributedCacheParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the update msg vpn distributed cache params
func (o *UpdateMsgVpnDistributedCacheParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithMsgVpnName adds the msgVpnName to the update msg vpn distributed cache params
func (o *UpdateMsgVpnDistributedCacheParams) WithMsgVpnName(msgVpnName string) *UpdateMsgVpnDistributedCacheParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the update msg vpn distributed cache params
func (o *UpdateMsgVpnDistributedCacheParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the update msg vpn distributed cache params
func (o *UpdateMsgVpnDistributedCacheParams) WithSelect(selectVar []string) *UpdateMsgVpnDistributedCacheParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the update msg vpn distributed cache params
func (o *UpdateMsgVpnDistributedCacheParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMsgVpnDistributedCacheParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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