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

// NewGetAboutParams creates a new GetAboutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAboutParams() *GetAboutParams {
	return &GetAboutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAboutParamsWithTimeout creates a new GetAboutParams object
// with the ability to set a timeout on a request.
func NewGetAboutParamsWithTimeout(timeout time.Duration) *GetAboutParams {
	return &GetAboutParams{
		timeout: timeout,
	}
}

// NewGetAboutParamsWithContext creates a new GetAboutParams object
// with the ability to set a context for a request.
func NewGetAboutParamsWithContext(ctx context.Context) *GetAboutParams {
	return &GetAboutParams{
		Context: ctx,
	}
}

// NewGetAboutParamsWithHTTPClient creates a new GetAboutParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAboutParamsWithHTTPClient(client *http.Client) *GetAboutParams {
	return &GetAboutParams{
		HTTPClient: client,
	}
}

/* GetAboutParams contains all the parameters to send to the API endpoint
   for the get about operation.

   Typically these are written to a http.Request.
*/
type GetAboutParams struct {

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

// WithDefaults hydrates default values in the get about params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAboutParams) WithDefaults() *GetAboutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get about params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAboutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get about params
func (o *GetAboutParams) WithTimeout(timeout time.Duration) *GetAboutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get about params
func (o *GetAboutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get about params
func (o *GetAboutParams) WithContext(ctx context.Context) *GetAboutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get about params
func (o *GetAboutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get about params
func (o *GetAboutParams) WithHTTPClient(client *http.Client) *GetAboutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get about params
func (o *GetAboutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOpaquePassword adds the opaquePassword to the get about params
func (o *GetAboutParams) WithOpaquePassword(opaquePassword *string) *GetAboutParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the get about params
func (o *GetAboutParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the get about params
func (o *GetAboutParams) WithSelect(selectVar []string) *GetAboutParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get about params
func (o *GetAboutParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetAboutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

// bindParamGetAbout binds the parameter select
func (o *GetAboutParams) bindParamSelect(formats strfmt.Registry) []string {
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
