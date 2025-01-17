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

// NewGetDmrClusterLinkTLSTrustedCommonNamesParams creates a new GetDmrClusterLinkTLSTrustedCommonNamesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDmrClusterLinkTLSTrustedCommonNamesParams() *GetDmrClusterLinkTLSTrustedCommonNamesParams {
	return &GetDmrClusterLinkTLSTrustedCommonNamesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDmrClusterLinkTLSTrustedCommonNamesParamsWithTimeout creates a new GetDmrClusterLinkTLSTrustedCommonNamesParams object
// with the ability to set a timeout on a request.
func NewGetDmrClusterLinkTLSTrustedCommonNamesParamsWithTimeout(timeout time.Duration) *GetDmrClusterLinkTLSTrustedCommonNamesParams {
	return &GetDmrClusterLinkTLSTrustedCommonNamesParams{
		timeout: timeout,
	}
}

// NewGetDmrClusterLinkTLSTrustedCommonNamesParamsWithContext creates a new GetDmrClusterLinkTLSTrustedCommonNamesParams object
// with the ability to set a context for a request.
func NewGetDmrClusterLinkTLSTrustedCommonNamesParamsWithContext(ctx context.Context) *GetDmrClusterLinkTLSTrustedCommonNamesParams {
	return &GetDmrClusterLinkTLSTrustedCommonNamesParams{
		Context: ctx,
	}
}

// NewGetDmrClusterLinkTLSTrustedCommonNamesParamsWithHTTPClient creates a new GetDmrClusterLinkTLSTrustedCommonNamesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDmrClusterLinkTLSTrustedCommonNamesParamsWithHTTPClient(client *http.Client) *GetDmrClusterLinkTLSTrustedCommonNamesParams {
	return &GetDmrClusterLinkTLSTrustedCommonNamesParams{
		HTTPClient: client,
	}
}

/* GetDmrClusterLinkTLSTrustedCommonNamesParams contains all the parameters to send to the API endpoint
   for the get dmr cluster link Tls trusted common names operation.

   Typically these are written to a http.Request.
*/
type GetDmrClusterLinkTLSTrustedCommonNamesParams struct {

	/* DmrClusterName.

	   The name of the Cluster.
	*/
	DmrClusterName string

	/* OpaquePassword.

	   Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter.
	*/
	OpaquePassword *string

	/* RemoteNodeName.

	   The name of the node at the remote end of the Link.
	*/
	RemoteNodeName string

	/* Select.

	   Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.
	*/
	Select []string

	/* Where.

	   Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter.
	*/
	Where []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dmr cluster link Tls trusted common names params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) WithDefaults() *GetDmrClusterLinkTLSTrustedCommonNamesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dmr cluster link Tls trusted common names params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dmr cluster link Tls trusted common names params
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) WithTimeout(timeout time.Duration) *GetDmrClusterLinkTLSTrustedCommonNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dmr cluster link Tls trusted common names params
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dmr cluster link Tls trusted common names params
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) WithContext(ctx context.Context) *GetDmrClusterLinkTLSTrustedCommonNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dmr cluster link Tls trusted common names params
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dmr cluster link Tls trusted common names params
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) WithHTTPClient(client *http.Client) *GetDmrClusterLinkTLSTrustedCommonNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dmr cluster link Tls trusted common names params
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDmrClusterName adds the dmrClusterName to the get dmr cluster link Tls trusted common names params
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) WithDmrClusterName(dmrClusterName string) *GetDmrClusterLinkTLSTrustedCommonNamesParams {
	o.SetDmrClusterName(dmrClusterName)
	return o
}

// SetDmrClusterName adds the dmrClusterName to the get dmr cluster link Tls trusted common names params
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) SetDmrClusterName(dmrClusterName string) {
	o.DmrClusterName = dmrClusterName
}

// WithOpaquePassword adds the opaquePassword to the get dmr cluster link Tls trusted common names params
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) WithOpaquePassword(opaquePassword *string) *GetDmrClusterLinkTLSTrustedCommonNamesParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the get dmr cluster link Tls trusted common names params
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithRemoteNodeName adds the remoteNodeName to the get dmr cluster link Tls trusted common names params
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) WithRemoteNodeName(remoteNodeName string) *GetDmrClusterLinkTLSTrustedCommonNamesParams {
	o.SetRemoteNodeName(remoteNodeName)
	return o
}

// SetRemoteNodeName adds the remoteNodeName to the get dmr cluster link Tls trusted common names params
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) SetRemoteNodeName(remoteNodeName string) {
	o.RemoteNodeName = remoteNodeName
}

// WithSelect adds the selectVar to the get dmr cluster link Tls trusted common names params
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) WithSelect(selectVar []string) *GetDmrClusterLinkTLSTrustedCommonNamesParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get dmr cluster link Tls trusted common names params
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithWhere adds the where to the get dmr cluster link Tls trusted common names params
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) WithWhere(where []string) *GetDmrClusterLinkTLSTrustedCommonNamesParams {
	o.SetWhere(where)
	return o
}

// SetWhere adds the where to the get dmr cluster link Tls trusted common names params
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) SetWhere(where []string) {
	o.Where = where
}

// WriteToRequest writes these params to a swagger request
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dmrClusterName
	if err := r.SetPathParam("dmrClusterName", o.DmrClusterName); err != nil {
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

	// path param remoteNodeName
	if err := r.SetPathParam("remoteNodeName", o.RemoteNodeName); err != nil {
		return err
	}

	if o.Select != nil {

		// binding items for select
		joinedSelect := o.bindParamSelect(reg)

		// query array param select
		if err := r.SetQueryParam("select", joinedSelect...); err != nil {
			return err
		}
	}

	if o.Where != nil {

		// binding items for where
		joinedWhere := o.bindParamWhere(reg)

		// query array param where
		if err := r.SetQueryParam("where", joinedWhere...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetDmrClusterLinkTLSTrustedCommonNames binds the parameter select
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) bindParamSelect(formats strfmt.Registry) []string {
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

// bindParamGetDmrClusterLinkTLSTrustedCommonNames binds the parameter where
func (o *GetDmrClusterLinkTLSTrustedCommonNamesParams) bindParamWhere(formats strfmt.Registry) []string {
	whereIR := o.Where

	var whereIC []string
	for _, whereIIR := range whereIR { // explode []string

		whereIIV := whereIIR // string as string
		whereIC = append(whereIC, whereIIV)
	}

	// items.CollectionFormat: "csv"
	whereIS := swag.JoinByFormat(whereIC, "csv")

	return whereIS
}
