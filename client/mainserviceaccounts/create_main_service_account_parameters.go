// Code generated by go-swagger; DO NOT EDIT.

package mainserviceaccounts

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

	"github.com/kubermatic/go-kubermatic/models"
)

// NewCreateMainServiceAccountParams creates a new CreateMainServiceAccountParams object
// with the default values initialized.
func NewCreateMainServiceAccountParams() *CreateMainServiceAccountParams {
	var ()
	return &CreateMainServiceAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMainServiceAccountParamsWithTimeout creates a new CreateMainServiceAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMainServiceAccountParamsWithTimeout(timeout time.Duration) *CreateMainServiceAccountParams {
	var ()
	return &CreateMainServiceAccountParams{

		timeout: timeout,
	}
}

// NewCreateMainServiceAccountParamsWithContext creates a new CreateMainServiceAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMainServiceAccountParamsWithContext(ctx context.Context) *CreateMainServiceAccountParams {
	var ()
	return &CreateMainServiceAccountParams{

		Context: ctx,
	}
}

// NewCreateMainServiceAccountParamsWithHTTPClient creates a new CreateMainServiceAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMainServiceAccountParamsWithHTTPClient(client *http.Client) *CreateMainServiceAccountParams {
	var ()
	return &CreateMainServiceAccountParams{
		HTTPClient: client,
	}
}

/*CreateMainServiceAccountParams contains all the parameters to send to the API endpoint
for the create main service account operation typically these are written to a http.Request
*/
type CreateMainServiceAccountParams struct {

	/*Body*/
	Body *models.ServiceAccount

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create main service account params
func (o *CreateMainServiceAccountParams) WithTimeout(timeout time.Duration) *CreateMainServiceAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create main service account params
func (o *CreateMainServiceAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create main service account params
func (o *CreateMainServiceAccountParams) WithContext(ctx context.Context) *CreateMainServiceAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create main service account params
func (o *CreateMainServiceAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create main service account params
func (o *CreateMainServiceAccountParams) WithHTTPClient(client *http.Client) *CreateMainServiceAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create main service account params
func (o *CreateMainServiceAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create main service account params
func (o *CreateMainServiceAccountParams) WithBody(body *models.ServiceAccount) *CreateMainServiceAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create main service account params
func (o *CreateMainServiceAccountParams) SetBody(body *models.ServiceAccount) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMainServiceAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
