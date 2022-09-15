// Code generated by go-swagger; DO NOT EDIT.

package preset

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

// NewDeletePresetProviderParams creates a new DeletePresetProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePresetProviderParams() *DeletePresetProviderParams {
	return &DeletePresetProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePresetProviderParamsWithTimeout creates a new DeletePresetProviderParams object
// with the ability to set a timeout on a request.
func NewDeletePresetProviderParamsWithTimeout(timeout time.Duration) *DeletePresetProviderParams {
	return &DeletePresetProviderParams{
		timeout: timeout,
	}
}

// NewDeletePresetProviderParamsWithContext creates a new DeletePresetProviderParams object
// with the ability to set a context for a request.
func NewDeletePresetProviderParamsWithContext(ctx context.Context) *DeletePresetProviderParams {
	return &DeletePresetProviderParams{
		Context: ctx,
	}
}

// NewDeletePresetProviderParamsWithHTTPClient creates a new DeletePresetProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePresetProviderParamsWithHTTPClient(client *http.Client) *DeletePresetProviderParams {
	return &DeletePresetProviderParams{
		HTTPClient: client,
	}
}

/*
DeletePresetProviderParams contains all the parameters to send to the API endpoint

	for the delete preset provider operation.

	Typically these are written to a http.Request.
*/
type DeletePresetProviderParams struct {

	// PresetName.
	PresetName string

	// ProviderName.
	ProviderName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete preset provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePresetProviderParams) WithDefaults() *DeletePresetProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete preset provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePresetProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete preset provider params
func (o *DeletePresetProviderParams) WithTimeout(timeout time.Duration) *DeletePresetProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete preset provider params
func (o *DeletePresetProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete preset provider params
func (o *DeletePresetProviderParams) WithContext(ctx context.Context) *DeletePresetProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete preset provider params
func (o *DeletePresetProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete preset provider params
func (o *DeletePresetProviderParams) WithHTTPClient(client *http.Client) *DeletePresetProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete preset provider params
func (o *DeletePresetProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPresetName adds the presetName to the delete preset provider params
func (o *DeletePresetProviderParams) WithPresetName(presetName string) *DeletePresetProviderParams {
	o.SetPresetName(presetName)
	return o
}

// SetPresetName adds the presetName to the delete preset provider params
func (o *DeletePresetProviderParams) SetPresetName(presetName string) {
	o.PresetName = presetName
}

// WithProviderName adds the providerName to the delete preset provider params
func (o *DeletePresetProviderParams) WithProviderName(providerName string) *DeletePresetProviderParams {
	o.SetProviderName(providerName)
	return o
}

// SetProviderName adds the providerName to the delete preset provider params
func (o *DeletePresetProviderParams) SetProviderName(providerName string) {
	o.ProviderName = providerName
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePresetProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param preset_name
	if err := r.SetPathParam("preset_name", o.PresetName); err != nil {
		return err
	}

	// path param provider_name
	if err := r.SetPathParam("provider_name", o.ProviderName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
