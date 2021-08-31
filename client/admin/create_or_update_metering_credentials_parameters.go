// Code generated by go-swagger; DO NOT EDIT.

package admin

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
)

// NewCreateOrUpdateMeteringCredentialsParams creates a new CreateOrUpdateMeteringCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrUpdateMeteringCredentialsParams() *CreateOrUpdateMeteringCredentialsParams {
	return &CreateOrUpdateMeteringCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrUpdateMeteringCredentialsParamsWithTimeout creates a new CreateOrUpdateMeteringCredentialsParams object
// with the ability to set a timeout on a request.
func NewCreateOrUpdateMeteringCredentialsParamsWithTimeout(timeout time.Duration) *CreateOrUpdateMeteringCredentialsParams {
	return &CreateOrUpdateMeteringCredentialsParams{
		timeout: timeout,
	}
}

// NewCreateOrUpdateMeteringCredentialsParamsWithContext creates a new CreateOrUpdateMeteringCredentialsParams object
// with the ability to set a context for a request.
func NewCreateOrUpdateMeteringCredentialsParamsWithContext(ctx context.Context) *CreateOrUpdateMeteringCredentialsParams {
	return &CreateOrUpdateMeteringCredentialsParams{
		Context: ctx,
	}
}

// NewCreateOrUpdateMeteringCredentialsParamsWithHTTPClient creates a new CreateOrUpdateMeteringCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrUpdateMeteringCredentialsParamsWithHTTPClient(client *http.Client) *CreateOrUpdateMeteringCredentialsParams {
	return &CreateOrUpdateMeteringCredentialsParams{
		HTTPClient: client,
	}
}

/* CreateOrUpdateMeteringCredentialsParams contains all the parameters to send to the API endpoint
   for the create or update metering credentials operation.

   Typically these are written to a http.Request.
*/
type CreateOrUpdateMeteringCredentialsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create or update metering credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrUpdateMeteringCredentialsParams) WithDefaults() *CreateOrUpdateMeteringCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create or update metering credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrUpdateMeteringCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create or update metering credentials params
func (o *CreateOrUpdateMeteringCredentialsParams) WithTimeout(timeout time.Duration) *CreateOrUpdateMeteringCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create or update metering credentials params
func (o *CreateOrUpdateMeteringCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create or update metering credentials params
func (o *CreateOrUpdateMeteringCredentialsParams) WithContext(ctx context.Context) *CreateOrUpdateMeteringCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create or update metering credentials params
func (o *CreateOrUpdateMeteringCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create or update metering credentials params
func (o *CreateOrUpdateMeteringCredentialsParams) WithHTTPClient(client *http.Client) *CreateOrUpdateMeteringCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create or update metering credentials params
func (o *CreateOrUpdateMeteringCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrUpdateMeteringCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
