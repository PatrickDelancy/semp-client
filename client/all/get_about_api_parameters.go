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

// NewGetAboutAPIParams creates a new GetAboutAPIParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAboutAPIParams() *GetAboutAPIParams {
	return &GetAboutAPIParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAboutAPIParamsWithTimeout creates a new GetAboutAPIParams object
// with the ability to set a timeout on a request.
func NewGetAboutAPIParamsWithTimeout(timeout time.Duration) *GetAboutAPIParams {
	return &GetAboutAPIParams{
		timeout: timeout,
	}
}

// NewGetAboutAPIParamsWithContext creates a new GetAboutAPIParams object
// with the ability to set a context for a request.
func NewGetAboutAPIParamsWithContext(ctx context.Context) *GetAboutAPIParams {
	return &GetAboutAPIParams{
		Context: ctx,
	}
}

// NewGetAboutAPIParamsWithHTTPClient creates a new GetAboutAPIParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAboutAPIParamsWithHTTPClient(client *http.Client) *GetAboutAPIParams {
	return &GetAboutAPIParams{
		HTTPClient: client,
	}
}

/* GetAboutAPIParams contains all the parameters to send to the API endpoint
   for the get about Api operation.

   Typically these are written to a http.Request.
*/
type GetAboutAPIParams struct {

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

// WithDefaults hydrates default values in the get about Api params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAboutAPIParams) WithDefaults() *GetAboutAPIParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get about Api params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAboutAPIParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get about Api params
func (o *GetAboutAPIParams) WithTimeout(timeout time.Duration) *GetAboutAPIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get about Api params
func (o *GetAboutAPIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get about Api params
func (o *GetAboutAPIParams) WithContext(ctx context.Context) *GetAboutAPIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get about Api params
func (o *GetAboutAPIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get about Api params
func (o *GetAboutAPIParams) WithHTTPClient(client *http.Client) *GetAboutAPIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get about Api params
func (o *GetAboutAPIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOpaquePassword adds the opaquePassword to the get about Api params
func (o *GetAboutAPIParams) WithOpaquePassword(opaquePassword *string) *GetAboutAPIParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the get about Api params
func (o *GetAboutAPIParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the get about Api params
func (o *GetAboutAPIParams) WithSelect(selectVar []string) *GetAboutAPIParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get about Api params
func (o *GetAboutAPIParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetAboutAPIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamGetAboutAPI binds the parameter select
func (o *GetAboutAPIParams) bindParamSelect(formats strfmt.Registry) []string {
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
