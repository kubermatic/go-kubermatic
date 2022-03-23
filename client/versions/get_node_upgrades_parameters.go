// Code generated by go-swagger; DO NOT EDIT.

package versions

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

// NewGetNodeUpgradesParams creates a new GetNodeUpgradesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNodeUpgradesParams() *GetNodeUpgradesParams {
	return &GetNodeUpgradesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodeUpgradesParamsWithTimeout creates a new GetNodeUpgradesParams object
// with the ability to set a timeout on a request.
func NewGetNodeUpgradesParamsWithTimeout(timeout time.Duration) *GetNodeUpgradesParams {
	return &GetNodeUpgradesParams{
		timeout: timeout,
	}
}

// NewGetNodeUpgradesParamsWithContext creates a new GetNodeUpgradesParams object
// with the ability to set a context for a request.
func NewGetNodeUpgradesParamsWithContext(ctx context.Context) *GetNodeUpgradesParams {
	return &GetNodeUpgradesParams{
		Context: ctx,
	}
}

// NewGetNodeUpgradesParamsWithHTTPClient creates a new GetNodeUpgradesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNodeUpgradesParamsWithHTTPClient(client *http.Client) *GetNodeUpgradesParams {
	return &GetNodeUpgradesParams{
		HTTPClient: client,
	}
}

/* GetNodeUpgradesParams contains all the parameters to send to the API endpoint
   for the get node upgrades operation.

   Typically these are written to a http.Request.
*/
type GetNodeUpgradesParams struct {

	// ControlPlaneVersion.
	ControlPlaneVersion *string

	/* Type.

	   Type is deprecated and not used anymore.
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get node upgrades params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNodeUpgradesParams) WithDefaults() *GetNodeUpgradesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get node upgrades params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNodeUpgradesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get node upgrades params
func (o *GetNodeUpgradesParams) WithTimeout(timeout time.Duration) *GetNodeUpgradesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get node upgrades params
func (o *GetNodeUpgradesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get node upgrades params
func (o *GetNodeUpgradesParams) WithContext(ctx context.Context) *GetNodeUpgradesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get node upgrades params
func (o *GetNodeUpgradesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get node upgrades params
func (o *GetNodeUpgradesParams) WithHTTPClient(client *http.Client) *GetNodeUpgradesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get node upgrades params
func (o *GetNodeUpgradesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithControlPlaneVersion adds the controlPlaneVersion to the get node upgrades params
func (o *GetNodeUpgradesParams) WithControlPlaneVersion(controlPlaneVersion *string) *GetNodeUpgradesParams {
	o.SetControlPlaneVersion(controlPlaneVersion)
	return o
}

// SetControlPlaneVersion adds the controlPlaneVersion to the get node upgrades params
func (o *GetNodeUpgradesParams) SetControlPlaneVersion(controlPlaneVersion *string) {
	o.ControlPlaneVersion = controlPlaneVersion
}

// WithType adds the typeVar to the get node upgrades params
func (o *GetNodeUpgradesParams) WithType(typeVar *string) *GetNodeUpgradesParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get node upgrades params
func (o *GetNodeUpgradesParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodeUpgradesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ControlPlaneVersion != nil {

		// query param control_plane_version
		var qrControlPlaneVersion string

		if o.ControlPlaneVersion != nil {
			qrControlPlaneVersion = *o.ControlPlaneVersion
		}
		qControlPlaneVersion := qrControlPlaneVersion
		if qControlPlaneVersion != "" {

			if err := r.SetQueryParam("control_plane_version", qControlPlaneVersion); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
