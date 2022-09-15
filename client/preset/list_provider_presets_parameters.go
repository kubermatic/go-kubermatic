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
	"github.com/go-openapi/swag"
)

// NewListProviderPresetsParams creates a new ListProviderPresetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProviderPresetsParams() *ListProviderPresetsParams {
	return &ListProviderPresetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProviderPresetsParamsWithTimeout creates a new ListProviderPresetsParams object
// with the ability to set a timeout on a request.
func NewListProviderPresetsParamsWithTimeout(timeout time.Duration) *ListProviderPresetsParams {
	return &ListProviderPresetsParams{
		timeout: timeout,
	}
}

// NewListProviderPresetsParamsWithContext creates a new ListProviderPresetsParams object
// with the ability to set a context for a request.
func NewListProviderPresetsParamsWithContext(ctx context.Context) *ListProviderPresetsParams {
	return &ListProviderPresetsParams{
		Context: ctx,
	}
}

// NewListProviderPresetsParamsWithHTTPClient creates a new ListProviderPresetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProviderPresetsParamsWithHTTPClient(client *http.Client) *ListProviderPresetsParams {
	return &ListProviderPresetsParams{
		HTTPClient: client,
	}
}

/*
ListProviderPresetsParams contains all the parameters to send to the API endpoint

	for the list provider presets operation.

	Typically these are written to a http.Request.
*/
type ListProviderPresetsParams struct {

	// Datacenter.
	Datacenter *string

	// Disabled.
	Disabled *bool

	// ProviderName.
	ProviderName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list provider presets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProviderPresetsParams) WithDefaults() *ListProviderPresetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list provider presets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProviderPresetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list provider presets params
func (o *ListProviderPresetsParams) WithTimeout(timeout time.Duration) *ListProviderPresetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list provider presets params
func (o *ListProviderPresetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list provider presets params
func (o *ListProviderPresetsParams) WithContext(ctx context.Context) *ListProviderPresetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list provider presets params
func (o *ListProviderPresetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list provider presets params
func (o *ListProviderPresetsParams) WithHTTPClient(client *http.Client) *ListProviderPresetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list provider presets params
func (o *ListProviderPresetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatacenter adds the datacenter to the list provider presets params
func (o *ListProviderPresetsParams) WithDatacenter(datacenter *string) *ListProviderPresetsParams {
	o.SetDatacenter(datacenter)
	return o
}

// SetDatacenter adds the datacenter to the list provider presets params
func (o *ListProviderPresetsParams) SetDatacenter(datacenter *string) {
	o.Datacenter = datacenter
}

// WithDisabled adds the disabled to the list provider presets params
func (o *ListProviderPresetsParams) WithDisabled(disabled *bool) *ListProviderPresetsParams {
	o.SetDisabled(disabled)
	return o
}

// SetDisabled adds the disabled to the list provider presets params
func (o *ListProviderPresetsParams) SetDisabled(disabled *bool) {
	o.Disabled = disabled
}

// WithProviderName adds the providerName to the list provider presets params
func (o *ListProviderPresetsParams) WithProviderName(providerName string) *ListProviderPresetsParams {
	o.SetProviderName(providerName)
	return o
}

// SetProviderName adds the providerName to the list provider presets params
func (o *ListProviderPresetsParams) SetProviderName(providerName string) {
	o.ProviderName = providerName
}

// WriteToRequest writes these params to a swagger request
func (o *ListProviderPresetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Datacenter != nil {

		// query param datacenter
		var qrDatacenter string

		if o.Datacenter != nil {
			qrDatacenter = *o.Datacenter
		}
		qDatacenter := qrDatacenter
		if qDatacenter != "" {

			if err := r.SetQueryParam("datacenter", qDatacenter); err != nil {
				return err
			}
		}
	}

	if o.Disabled != nil {

		// query param disabled
		var qrDisabled bool

		if o.Disabled != nil {
			qrDisabled = *o.Disabled
		}
		qDisabled := swag.FormatBool(qrDisabled)
		if qDisabled != "" {

			if err := r.SetQueryParam("disabled", qDisabled); err != nil {
				return err
			}
		}
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
