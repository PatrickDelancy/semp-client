// Code generated by go-swagger; DO NOT EDIT.

package dmr_cluster

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

// NewCreateDmrClusterLinkParams creates a new CreateDmrClusterLinkParams object
// with the default values initialized.
func NewCreateDmrClusterLinkParams() *CreateDmrClusterLinkParams {
	var ()
	return &CreateDmrClusterLinkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDmrClusterLinkParamsWithTimeout creates a new CreateDmrClusterLinkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDmrClusterLinkParamsWithTimeout(timeout time.Duration) *CreateDmrClusterLinkParams {
	var ()
	return &CreateDmrClusterLinkParams{

		timeout: timeout,
	}
}

// NewCreateDmrClusterLinkParamsWithContext creates a new CreateDmrClusterLinkParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDmrClusterLinkParamsWithContext(ctx context.Context) *CreateDmrClusterLinkParams {
	var ()
	return &CreateDmrClusterLinkParams{

		Context: ctx,
	}
}

// NewCreateDmrClusterLinkParamsWithHTTPClient creates a new CreateDmrClusterLinkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDmrClusterLinkParamsWithHTTPClient(client *http.Client) *CreateDmrClusterLinkParams {
	var ()
	return &CreateDmrClusterLinkParams{
		HTTPClient: client,
	}
}

/*CreateDmrClusterLinkParams contains all the parameters to send to the API endpoint
for the create dmr cluster link operation typically these are written to a http.Request
*/
type CreateDmrClusterLinkParams struct {

	/*Body
	  The Link object's attributes.

	*/
	Body *models.DmrClusterLink
	/*DmrClusterName
	  The name of the Cluster.

	*/
	DmrClusterName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create dmr cluster link params
func (o *CreateDmrClusterLinkParams) WithTimeout(timeout time.Duration) *CreateDmrClusterLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create dmr cluster link params
func (o *CreateDmrClusterLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create dmr cluster link params
func (o *CreateDmrClusterLinkParams) WithContext(ctx context.Context) *CreateDmrClusterLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create dmr cluster link params
func (o *CreateDmrClusterLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create dmr cluster link params
func (o *CreateDmrClusterLinkParams) WithHTTPClient(client *http.Client) *CreateDmrClusterLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create dmr cluster link params
func (o *CreateDmrClusterLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create dmr cluster link params
func (o *CreateDmrClusterLinkParams) WithBody(body *models.DmrClusterLink) *CreateDmrClusterLinkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create dmr cluster link params
func (o *CreateDmrClusterLinkParams) SetBody(body *models.DmrClusterLink) {
	o.Body = body
}

// WithDmrClusterName adds the dmrClusterName to the create dmr cluster link params
func (o *CreateDmrClusterLinkParams) WithDmrClusterName(dmrClusterName string) *CreateDmrClusterLinkParams {
	o.SetDmrClusterName(dmrClusterName)
	return o
}

// SetDmrClusterName adds the dmrClusterName to the create dmr cluster link params
func (o *CreateDmrClusterLinkParams) SetDmrClusterName(dmrClusterName string) {
	o.DmrClusterName = dmrClusterName
}

// WithSelect adds the selectVar to the create dmr cluster link params
func (o *CreateDmrClusterLinkParams) WithSelect(selectVar []string) *CreateDmrClusterLinkParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the create dmr cluster link params
func (o *CreateDmrClusterLinkParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDmrClusterLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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