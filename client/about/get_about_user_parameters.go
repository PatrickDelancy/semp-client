// Code generated by go-swagger; DO NOT EDIT.

package about

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
)

// NewGetAboutUserParams creates a new GetAboutUserParams object
// with the default values initialized.
func NewGetAboutUserParams() *GetAboutUserParams {
	var ()
	return &GetAboutUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAboutUserParamsWithTimeout creates a new GetAboutUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAboutUserParamsWithTimeout(timeout time.Duration) *GetAboutUserParams {
	var ()
	return &GetAboutUserParams{

		timeout: timeout,
	}
}

// NewGetAboutUserParamsWithContext creates a new GetAboutUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAboutUserParamsWithContext(ctx context.Context) *GetAboutUserParams {
	var ()
	return &GetAboutUserParams{

		Context: ctx,
	}
}

// NewGetAboutUserParamsWithHTTPClient creates a new GetAboutUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAboutUserParamsWithHTTPClient(client *http.Client) *GetAboutUserParams {
	var ()
	return &GetAboutUserParams{
		HTTPClient: client,
	}
}

/*GetAboutUserParams contains all the parameters to send to the API endpoint
for the get about user operation typically these are written to a http.Request
*/
type GetAboutUserParams struct {

	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See [Select](#select "Description of the syntax of the `select` parameter").

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get about user params
func (o *GetAboutUserParams) WithTimeout(timeout time.Duration) *GetAboutUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get about user params
func (o *GetAboutUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get about user params
func (o *GetAboutUserParams) WithContext(ctx context.Context) *GetAboutUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get about user params
func (o *GetAboutUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get about user params
func (o *GetAboutUserParams) WithHTTPClient(client *http.Client) *GetAboutUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get about user params
func (o *GetAboutUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSelect adds the selectVar to the get about user params
func (o *GetAboutUserParams) WithSelect(selectVar []string) *GetAboutUserParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get about user params
func (o *GetAboutUserParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetAboutUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
