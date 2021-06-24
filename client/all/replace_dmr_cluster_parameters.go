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

// NewReplaceDmrClusterParams creates a new ReplaceDmrClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceDmrClusterParams() *ReplaceDmrClusterParams {
	return &ReplaceDmrClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceDmrClusterParamsWithTimeout creates a new ReplaceDmrClusterParams object
// with the ability to set a timeout on a request.
func NewReplaceDmrClusterParamsWithTimeout(timeout time.Duration) *ReplaceDmrClusterParams {
	return &ReplaceDmrClusterParams{
		timeout: timeout,
	}
}

// NewReplaceDmrClusterParamsWithContext creates a new ReplaceDmrClusterParams object
// with the ability to set a context for a request.
func NewReplaceDmrClusterParamsWithContext(ctx context.Context) *ReplaceDmrClusterParams {
	return &ReplaceDmrClusterParams{
		Context: ctx,
	}
}

// NewReplaceDmrClusterParamsWithHTTPClient creates a new ReplaceDmrClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceDmrClusterParamsWithHTTPClient(client *http.Client) *ReplaceDmrClusterParams {
	return &ReplaceDmrClusterParams{
		HTTPClient: client,
	}
}

/* ReplaceDmrClusterParams contains all the parameters to send to the API endpoint
   for the replace dmr cluster operation.

   Typically these are written to a http.Request.
*/
type ReplaceDmrClusterParams struct {

	/* Body.

	   The Cluster object's attributes.
	*/
	Body *models.DmrCluster

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

// WithDefaults hydrates default values in the replace dmr cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceDmrClusterParams) WithDefaults() *ReplaceDmrClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace dmr cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceDmrClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the replace dmr cluster params
func (o *ReplaceDmrClusterParams) WithTimeout(timeout time.Duration) *ReplaceDmrClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace dmr cluster params
func (o *ReplaceDmrClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace dmr cluster params
func (o *ReplaceDmrClusterParams) WithContext(ctx context.Context) *ReplaceDmrClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace dmr cluster params
func (o *ReplaceDmrClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace dmr cluster params
func (o *ReplaceDmrClusterParams) WithHTTPClient(client *http.Client) *ReplaceDmrClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace dmr cluster params
func (o *ReplaceDmrClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace dmr cluster params
func (o *ReplaceDmrClusterParams) WithBody(body *models.DmrCluster) *ReplaceDmrClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace dmr cluster params
func (o *ReplaceDmrClusterParams) SetBody(body *models.DmrCluster) {
	o.Body = body
}

// WithDmrClusterName adds the dmrClusterName to the replace dmr cluster params
func (o *ReplaceDmrClusterParams) WithDmrClusterName(dmrClusterName string) *ReplaceDmrClusterParams {
	o.SetDmrClusterName(dmrClusterName)
	return o
}

// SetDmrClusterName adds the dmrClusterName to the replace dmr cluster params
func (o *ReplaceDmrClusterParams) SetDmrClusterName(dmrClusterName string) {
	o.DmrClusterName = dmrClusterName
}

// WithOpaquePassword adds the opaquePassword to the replace dmr cluster params
func (o *ReplaceDmrClusterParams) WithOpaquePassword(opaquePassword *string) *ReplaceDmrClusterParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the replace dmr cluster params
func (o *ReplaceDmrClusterParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the replace dmr cluster params
func (o *ReplaceDmrClusterParams) WithSelect(selectVar []string) *ReplaceDmrClusterParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the replace dmr cluster params
func (o *ReplaceDmrClusterParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceDmrClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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

// bindParamReplaceDmrCluster binds the parameter select
func (o *ReplaceDmrClusterParams) bindParamSelect(formats strfmt.Registry) []string {
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