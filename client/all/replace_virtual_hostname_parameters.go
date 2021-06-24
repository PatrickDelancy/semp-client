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

// NewReplaceVirtualHostnameParams creates a new ReplaceVirtualHostnameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceVirtualHostnameParams() *ReplaceVirtualHostnameParams {
	return &ReplaceVirtualHostnameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceVirtualHostnameParamsWithTimeout creates a new ReplaceVirtualHostnameParams object
// with the ability to set a timeout on a request.
func NewReplaceVirtualHostnameParamsWithTimeout(timeout time.Duration) *ReplaceVirtualHostnameParams {
	return &ReplaceVirtualHostnameParams{
		timeout: timeout,
	}
}

// NewReplaceVirtualHostnameParamsWithContext creates a new ReplaceVirtualHostnameParams object
// with the ability to set a context for a request.
func NewReplaceVirtualHostnameParamsWithContext(ctx context.Context) *ReplaceVirtualHostnameParams {
	return &ReplaceVirtualHostnameParams{
		Context: ctx,
	}
}

// NewReplaceVirtualHostnameParamsWithHTTPClient creates a new ReplaceVirtualHostnameParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceVirtualHostnameParamsWithHTTPClient(client *http.Client) *ReplaceVirtualHostnameParams {
	return &ReplaceVirtualHostnameParams{
		HTTPClient: client,
	}
}

/* ReplaceVirtualHostnameParams contains all the parameters to send to the API endpoint
   for the replace virtual hostname operation.

   Typically these are written to a http.Request.
*/
type ReplaceVirtualHostnameParams struct {

	/* Body.

	   The Virtual Hostname object's attributes.
	*/
	Body *models.VirtualHostname

	/* OpaquePassword.

	   Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter.
	*/
	OpaquePassword *string

	/* Select.

	   Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.
	*/
	Select []string

	/* VirtualHostname.

	   The virtual hostname.
	*/
	VirtualHostname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the replace virtual hostname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceVirtualHostnameParams) WithDefaults() *ReplaceVirtualHostnameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace virtual hostname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceVirtualHostnameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the replace virtual hostname params
func (o *ReplaceVirtualHostnameParams) WithTimeout(timeout time.Duration) *ReplaceVirtualHostnameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace virtual hostname params
func (o *ReplaceVirtualHostnameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace virtual hostname params
func (o *ReplaceVirtualHostnameParams) WithContext(ctx context.Context) *ReplaceVirtualHostnameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace virtual hostname params
func (o *ReplaceVirtualHostnameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace virtual hostname params
func (o *ReplaceVirtualHostnameParams) WithHTTPClient(client *http.Client) *ReplaceVirtualHostnameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace virtual hostname params
func (o *ReplaceVirtualHostnameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace virtual hostname params
func (o *ReplaceVirtualHostnameParams) WithBody(body *models.VirtualHostname) *ReplaceVirtualHostnameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace virtual hostname params
func (o *ReplaceVirtualHostnameParams) SetBody(body *models.VirtualHostname) {
	o.Body = body
}

// WithOpaquePassword adds the opaquePassword to the replace virtual hostname params
func (o *ReplaceVirtualHostnameParams) WithOpaquePassword(opaquePassword *string) *ReplaceVirtualHostnameParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the replace virtual hostname params
func (o *ReplaceVirtualHostnameParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the replace virtual hostname params
func (o *ReplaceVirtualHostnameParams) WithSelect(selectVar []string) *ReplaceVirtualHostnameParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the replace virtual hostname params
func (o *ReplaceVirtualHostnameParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithVirtualHostname adds the virtualHostname to the replace virtual hostname params
func (o *ReplaceVirtualHostnameParams) WithVirtualHostname(virtualHostname string) *ReplaceVirtualHostnameParams {
	o.SetVirtualHostname(virtualHostname)
	return o
}

// SetVirtualHostname adds the virtualHostname to the replace virtual hostname params
func (o *ReplaceVirtualHostnameParams) SetVirtualHostname(virtualHostname string) {
	o.VirtualHostname = virtualHostname
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceVirtualHostnameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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

	// path param virtualHostname
	if err := r.SetPathParam("virtualHostname", o.VirtualHostname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamReplaceVirtualHostname binds the parameter select
func (o *ReplaceVirtualHostnameParams) bindParamSelect(formats strfmt.Registry) []string {
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