// Code generated by go-swagger; DO NOT EDIT.

package resource_quota

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

// NewDeleteResourceQuotaParams creates a new DeleteResourceQuotaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteResourceQuotaParams() *DeleteResourceQuotaParams {
	return &DeleteResourceQuotaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteResourceQuotaParamsWithTimeout creates a new DeleteResourceQuotaParams object
// with the ability to set a timeout on a request.
func NewDeleteResourceQuotaParamsWithTimeout(timeout time.Duration) *DeleteResourceQuotaParams {
	return &DeleteResourceQuotaParams{
		timeout: timeout,
	}
}

// NewDeleteResourceQuotaParamsWithContext creates a new DeleteResourceQuotaParams object
// with the ability to set a context for a request.
func NewDeleteResourceQuotaParamsWithContext(ctx context.Context) *DeleteResourceQuotaParams {
	return &DeleteResourceQuotaParams{
		Context: ctx,
	}
}

// NewDeleteResourceQuotaParamsWithHTTPClient creates a new DeleteResourceQuotaParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteResourceQuotaParamsWithHTTPClient(client *http.Client) *DeleteResourceQuotaParams {
	return &DeleteResourceQuotaParams{
		HTTPClient: client,
	}
}

/*
DeleteResourceQuotaParams contains all the parameters to send to the API endpoint

	for the delete resource quota operation.

	Typically these are written to a http.Request.
*/
type DeleteResourceQuotaParams struct {

	// QuotaName.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete resource quota params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteResourceQuotaParams) WithDefaults() *DeleteResourceQuotaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete resource quota params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteResourceQuotaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete resource quota params
func (o *DeleteResourceQuotaParams) WithTimeout(timeout time.Duration) *DeleteResourceQuotaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete resource quota params
func (o *DeleteResourceQuotaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete resource quota params
func (o *DeleteResourceQuotaParams) WithContext(ctx context.Context) *DeleteResourceQuotaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete resource quota params
func (o *DeleteResourceQuotaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete resource quota params
func (o *DeleteResourceQuotaParams) WithHTTPClient(client *http.Client) *DeleteResourceQuotaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete resource quota params
func (o *DeleteResourceQuotaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the quotaName to the delete resource quota params
func (o *DeleteResourceQuotaParams) WithName(quotaName string) *DeleteResourceQuotaParams {
	o.SetName(quotaName)
	return o
}

// SetName adds the quotaName to the delete resource quota params
func (o *DeleteResourceQuotaParams) SetName(quotaName string) {
	o.Name = quotaName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteResourceQuotaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param quota_name
	if err := r.SetPathParam("quota_name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
