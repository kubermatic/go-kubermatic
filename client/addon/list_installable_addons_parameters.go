// Code generated by go-swagger; DO NOT EDIT.

package addon

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

// NewListInstallableAddonsParams creates a new ListInstallableAddonsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListInstallableAddonsParams() *ListInstallableAddonsParams {
	return &ListInstallableAddonsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListInstallableAddonsParamsWithTimeout creates a new ListInstallableAddonsParams object
// with the ability to set a timeout on a request.
func NewListInstallableAddonsParamsWithTimeout(timeout time.Duration) *ListInstallableAddonsParams {
	return &ListInstallableAddonsParams{
		timeout: timeout,
	}
}

// NewListInstallableAddonsParamsWithContext creates a new ListInstallableAddonsParams object
// with the ability to set a context for a request.
func NewListInstallableAddonsParamsWithContext(ctx context.Context) *ListInstallableAddonsParams {
	return &ListInstallableAddonsParams{
		Context: ctx,
	}
}

// NewListInstallableAddonsParamsWithHTTPClient creates a new ListInstallableAddonsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListInstallableAddonsParamsWithHTTPClient(client *http.Client) *ListInstallableAddonsParams {
	return &ListInstallableAddonsParams{
		HTTPClient: client,
	}
}

/* ListInstallableAddonsParams contains all the parameters to send to the API endpoint
   for the list installable addons operation.

   Typically these are written to a http.Request.
*/
type ListInstallableAddonsParams struct {

	// ClusterID.
	ClusterID string

	// Dc.
	DC string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list installable addons params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListInstallableAddonsParams) WithDefaults() *ListInstallableAddonsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list installable addons params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListInstallableAddonsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list installable addons params
func (o *ListInstallableAddonsParams) WithTimeout(timeout time.Duration) *ListInstallableAddonsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list installable addons params
func (o *ListInstallableAddonsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list installable addons params
func (o *ListInstallableAddonsParams) WithContext(ctx context.Context) *ListInstallableAddonsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list installable addons params
func (o *ListInstallableAddonsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list installable addons params
func (o *ListInstallableAddonsParams) WithHTTPClient(client *http.Client) *ListInstallableAddonsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list installable addons params
func (o *ListInstallableAddonsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list installable addons params
func (o *ListInstallableAddonsParams) WithClusterID(clusterID string) *ListInstallableAddonsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list installable addons params
func (o *ListInstallableAddonsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the list installable addons params
func (o *ListInstallableAddonsParams) WithDC(dc string) *ListInstallableAddonsParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list installable addons params
func (o *ListInstallableAddonsParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the list installable addons params
func (o *ListInstallableAddonsParams) WithProjectID(projectID string) *ListInstallableAddonsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list installable addons params
func (o *ListInstallableAddonsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListInstallableAddonsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
