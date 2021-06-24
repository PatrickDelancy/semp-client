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

// NewGetClientCertAuthoritiesParams creates a new GetClientCertAuthoritiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClientCertAuthoritiesParams() *GetClientCertAuthoritiesParams {
	return &GetClientCertAuthoritiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClientCertAuthoritiesParamsWithTimeout creates a new GetClientCertAuthoritiesParams object
// with the ability to set a timeout on a request.
func NewGetClientCertAuthoritiesParamsWithTimeout(timeout time.Duration) *GetClientCertAuthoritiesParams {
	return &GetClientCertAuthoritiesParams{
		timeout: timeout,
	}
}

// NewGetClientCertAuthoritiesParamsWithContext creates a new GetClientCertAuthoritiesParams object
// with the ability to set a context for a request.
func NewGetClientCertAuthoritiesParamsWithContext(ctx context.Context) *GetClientCertAuthoritiesParams {
	return &GetClientCertAuthoritiesParams{
		Context: ctx,
	}
}

// NewGetClientCertAuthoritiesParamsWithHTTPClient creates a new GetClientCertAuthoritiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClientCertAuthoritiesParamsWithHTTPClient(client *http.Client) *GetClientCertAuthoritiesParams {
	return &GetClientCertAuthoritiesParams{
		HTTPClient: client,
	}
}

/* GetClientCertAuthoritiesParams contains all the parameters to send to the API endpoint
   for the get client cert authorities operation.

   Typically these are written to a http.Request.
*/
type GetClientCertAuthoritiesParams struct {

	/* Count.

	   Limit the count of objects in the response. See the documentation for the `count` parameter.

	   Default: 10
	*/
	Count *int64

	/* Cursor.

	   The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter.
	*/
	Cursor *string

	/* OpaquePassword.

	   Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter.
	*/
	OpaquePassword *string

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

// WithDefaults hydrates default values in the get client cert authorities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClientCertAuthoritiesParams) WithDefaults() *GetClientCertAuthoritiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get client cert authorities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClientCertAuthoritiesParams) SetDefaults() {
	var (
		countDefault = int64(10)
	)

	val := GetClientCertAuthoritiesParams{
		Count: &countDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get client cert authorities params
func (o *GetClientCertAuthoritiesParams) WithTimeout(timeout time.Duration) *GetClientCertAuthoritiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get client cert authorities params
func (o *GetClientCertAuthoritiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get client cert authorities params
func (o *GetClientCertAuthoritiesParams) WithContext(ctx context.Context) *GetClientCertAuthoritiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get client cert authorities params
func (o *GetClientCertAuthoritiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get client cert authorities params
func (o *GetClientCertAuthoritiesParams) WithHTTPClient(client *http.Client) *GetClientCertAuthoritiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get client cert authorities params
func (o *GetClientCertAuthoritiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get client cert authorities params
func (o *GetClientCertAuthoritiesParams) WithCount(count *int64) *GetClientCertAuthoritiesParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get client cert authorities params
func (o *GetClientCertAuthoritiesParams) SetCount(count *int64) {
	o.Count = count
}

// WithCursor adds the cursor to the get client cert authorities params
func (o *GetClientCertAuthoritiesParams) WithCursor(cursor *string) *GetClientCertAuthoritiesParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the get client cert authorities params
func (o *GetClientCertAuthoritiesParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithOpaquePassword adds the opaquePassword to the get client cert authorities params
func (o *GetClientCertAuthoritiesParams) WithOpaquePassword(opaquePassword *string) *GetClientCertAuthoritiesParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the get client cert authorities params
func (o *GetClientCertAuthoritiesParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the get client cert authorities params
func (o *GetClientCertAuthoritiesParams) WithSelect(selectVar []string) *GetClientCertAuthoritiesParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get client cert authorities params
func (o *GetClientCertAuthoritiesParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithWhere adds the where to the get client cert authorities params
func (o *GetClientCertAuthoritiesParams) WithWhere(where []string) *GetClientCertAuthoritiesParams {
	o.SetWhere(where)
	return o
}

// SetWhere adds the where to the get client cert authorities params
func (o *GetClientCertAuthoritiesParams) SetWhere(where []string) {
	o.Where = where
}

// WriteToRequest writes these params to a swagger request
func (o *GetClientCertAuthoritiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int64

		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {

			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}
	}

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string

		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {

			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}
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

// bindParamGetClientCertAuthorities binds the parameter select
func (o *GetClientCertAuthoritiesParams) bindParamSelect(formats strfmt.Registry) []string {
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

// bindParamGetClientCertAuthorities binds the parameter where
func (o *GetClientCertAuthoritiesParams) bindParamWhere(formats strfmt.Registry) []string {
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