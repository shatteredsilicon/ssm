// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/percona/pmm/api/json/models"
)

// NewAddGenericNodeParams creates a new AddGenericNodeParams object
// with the default values initialized.
func NewAddGenericNodeParams() *AddGenericNodeParams {
	var ()
	return &AddGenericNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddGenericNodeParamsWithTimeout creates a new AddGenericNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddGenericNodeParamsWithTimeout(timeout time.Duration) *AddGenericNodeParams {
	var ()
	return &AddGenericNodeParams{

		timeout: timeout,
	}
}

// NewAddGenericNodeParamsWithContext creates a new AddGenericNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddGenericNodeParamsWithContext(ctx context.Context) *AddGenericNodeParams {
	var ()
	return &AddGenericNodeParams{

		Context: ctx,
	}
}

// NewAddGenericNodeParamsWithHTTPClient creates a new AddGenericNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddGenericNodeParamsWithHTTPClient(client *http.Client) *AddGenericNodeParams {
	var ()
	return &AddGenericNodeParams{
		HTTPClient: client,
	}
}

/*AddGenericNodeParams contains all the parameters to send to the API endpoint
for the add generic node operation typically these are written to a http.Request
*/
type AddGenericNodeParams struct {

	/*Body*/
	Body *models.InventoryAddGenericNodeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add generic node params
func (o *AddGenericNodeParams) WithTimeout(timeout time.Duration) *AddGenericNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add generic node params
func (o *AddGenericNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add generic node params
func (o *AddGenericNodeParams) WithContext(ctx context.Context) *AddGenericNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add generic node params
func (o *AddGenericNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add generic node params
func (o *AddGenericNodeParams) WithHTTPClient(client *http.Client) *AddGenericNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add generic node params
func (o *AddGenericNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add generic node params
func (o *AddGenericNodeParams) WithBody(body *models.InventoryAddGenericNodeRequest) *AddGenericNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add generic node params
func (o *AddGenericNodeParams) SetBody(body *models.InventoryAddGenericNodeRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddGenericNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
