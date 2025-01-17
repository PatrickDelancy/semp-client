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

// NewGetDmrClusterParams creates a new GetDmrClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDmrClusterParams() *GetDmrClusterParams {
	return &GetDmrClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDmrClusterParamsWithTimeout creates a new GetDmrClusterParams object
// with the ability to set a timeout on a request.
func NewGetDmrClusterParamsWithTimeout(timeout time.Duration) *GetDmrClusterParams {
	return &GetDmrClusterParams{
		timeout: timeout,
	}
}

// NewGetDmrClusterParamsWithContext creates a new GetDmrClusterParams object
// with the ability to set a context for a request.
func NewGetDmrClusterParamsWithContext(ctx context.Context) *GetDmrClusterParams {
	return &GetDmrClusterParams{
		Context: ctx,
	}
}

// NewGetDmrClusterParamsWithHTTPClient creates a new GetDmrClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDmrClusterParamsWithHTTPClient(client *http.Client) *GetDmrClusterParams {
	return &GetDmrClusterParams{
		HTTPClient: client,
	}
}

/* GetDmrClusterParams contains all the parameters to send to the API endpoint
   for the get dmr cluster operation.

   Typically these are written to a http.Request.
*/
type GetDmrClusterParams struct {

	/* DmrClusterName.

	   The name of the Cluster.
	*/
	DmrClusterName string

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

// WithDefaults hydrates default values in the get dmr cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDmrClusterParams) WithDefaults() *GetDmrClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dmr cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDmrClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dmr cluster params
func (o *GetDmrClusterParams) WithTimeout(timeout time.Duration) *GetDmrClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dmr cluster params
func (o *GetDmrClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dmr cluster params
func (o *GetDmrClusterParams) WithContext(ctx context.Context) *GetDmrClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dmr cluster params
func (o *GetDmrClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dmr cluster params
func (o *GetDmrClusterParams) WithHTTPClient(client *http.Client) *GetDmrClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dmr cluster params
func (o *GetDmrClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDmrClusterName adds the dmrClusterName to the get dmr cluster params
func (o *GetDmrClusterParams) WithDmrClusterName(dmrClusterName string) *GetDmrClusterParams {
	o.SetDmrClusterName(dmrClusterName)
	return o
}

// SetDmrClusterName adds the dmrClusterName to the get dmr cluster params
func (o *GetDmrClusterParams) SetDmrClusterName(dmrClusterName string) {
	o.DmrClusterName = dmrClusterName
}

// WithOpaquePassword adds the opaquePassword to the get dmr cluster params
func (o *GetDmrClusterParams) WithOpaquePassword(opaquePassword *string) *GetDmrClusterParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the get dmr cluster params
func (o *GetDmrClusterParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the get dmr cluster params
func (o *GetDmrClusterParams) WithSelect(selectVar []string) *GetDmrClusterParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get dmr cluster params
func (o *GetDmrClusterParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetDmrClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamGetDmrCluster binds the parameter select
func (o *GetDmrClusterParams) bindParamSelect(formats strfmt.Registry) []string {
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
