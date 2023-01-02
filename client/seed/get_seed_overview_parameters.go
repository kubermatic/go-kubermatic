// Code generated by go-swagger; DO NOT EDIT.

package seed

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

// NewGetSeedOverviewParams creates a new GetSeedOverviewParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSeedOverviewParams() *GetSeedOverviewParams {
	return &GetSeedOverviewParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSeedOverviewParamsWithTimeout creates a new GetSeedOverviewParams object
// with the ability to set a timeout on a request.
func NewGetSeedOverviewParamsWithTimeout(timeout time.Duration) *GetSeedOverviewParams {
	return &GetSeedOverviewParams{
		timeout: timeout,
	}
}

// NewGetSeedOverviewParamsWithContext creates a new GetSeedOverviewParams object
// with the ability to set a context for a request.
func NewGetSeedOverviewParamsWithContext(ctx context.Context) *GetSeedOverviewParams {
	return &GetSeedOverviewParams{
		Context: ctx,
	}
}

// NewGetSeedOverviewParamsWithHTTPClient creates a new GetSeedOverviewParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSeedOverviewParamsWithHTTPClient(client *http.Client) *GetSeedOverviewParams {
	return &GetSeedOverviewParams{
		HTTPClient: client,
	}
}

/*
GetSeedOverviewParams contains all the parameters to send to the API endpoint

	for the get seed overview operation.

	Typically these are written to a http.Request.
*/
type GetSeedOverviewParams struct {

	// SeedName.
	SeedName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get seed overview params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSeedOverviewParams) WithDefaults() *GetSeedOverviewParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get seed overview params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSeedOverviewParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get seed overview params
func (o *GetSeedOverviewParams) WithTimeout(timeout time.Duration) *GetSeedOverviewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get seed overview params
func (o *GetSeedOverviewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get seed overview params
func (o *GetSeedOverviewParams) WithContext(ctx context.Context) *GetSeedOverviewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get seed overview params
func (o *GetSeedOverviewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get seed overview params
func (o *GetSeedOverviewParams) WithHTTPClient(client *http.Client) *GetSeedOverviewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get seed overview params
func (o *GetSeedOverviewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSeedName adds the seedName to the get seed overview params
func (o *GetSeedOverviewParams) WithSeedName(seedName string) *GetSeedOverviewParams {
	o.SetSeedName(seedName)
	return o
}

// SetSeedName adds the seedName to the get seed overview params
func (o *GetSeedOverviewParams) SetSeedName(seedName string) {
	o.SeedName = seedName
}

// WriteToRequest writes these params to a swagger request
func (o *GetSeedOverviewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param seed_name
	if err := r.SetPathParam("seed_name", o.SeedName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}