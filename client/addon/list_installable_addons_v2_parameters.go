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

// NewListInstallableAddonsV2Params creates a new ListInstallableAddonsV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListInstallableAddonsV2Params() *ListInstallableAddonsV2Params {
	return &ListInstallableAddonsV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewListInstallableAddonsV2ParamsWithTimeout creates a new ListInstallableAddonsV2Params object
// with the ability to set a timeout on a request.
func NewListInstallableAddonsV2ParamsWithTimeout(timeout time.Duration) *ListInstallableAddonsV2Params {
	return &ListInstallableAddonsV2Params{
		timeout: timeout,
	}
}

// NewListInstallableAddonsV2ParamsWithContext creates a new ListInstallableAddonsV2Params object
// with the ability to set a context for a request.
func NewListInstallableAddonsV2ParamsWithContext(ctx context.Context) *ListInstallableAddonsV2Params {
	return &ListInstallableAddonsV2Params{
		Context: ctx,
	}
}

// NewListInstallableAddonsV2ParamsWithHTTPClient creates a new ListInstallableAddonsV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewListInstallableAddonsV2ParamsWithHTTPClient(client *http.Client) *ListInstallableAddonsV2Params {
	return &ListInstallableAddonsV2Params{
		HTTPClient: client,
	}
}

/* ListInstallableAddonsV2Params contains all the parameters to send to the API endpoint
   for the list installable addons v2 operation.

   Typically these are written to a http.Request.
*/
type ListInstallableAddonsV2Params struct {

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list installable addons v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListInstallableAddonsV2Params) WithDefaults() *ListInstallableAddonsV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list installable addons v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListInstallableAddonsV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list installable addons v2 params
func (o *ListInstallableAddonsV2Params) WithTimeout(timeout time.Duration) *ListInstallableAddonsV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list installable addons v2 params
func (o *ListInstallableAddonsV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list installable addons v2 params
func (o *ListInstallableAddonsV2Params) WithContext(ctx context.Context) *ListInstallableAddonsV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list installable addons v2 params
func (o *ListInstallableAddonsV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list installable addons v2 params
func (o *ListInstallableAddonsV2Params) WithHTTPClient(client *http.Client) *ListInstallableAddonsV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list installable addons v2 params
func (o *ListInstallableAddonsV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list installable addons v2 params
func (o *ListInstallableAddonsV2Params) WithClusterID(clusterID string) *ListInstallableAddonsV2Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list installable addons v2 params
func (o *ListInstallableAddonsV2Params) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the list installable addons v2 params
func (o *ListInstallableAddonsV2Params) WithProjectID(projectID string) *ListInstallableAddonsV2Params {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list installable addons v2 params
func (o *ListInstallableAddonsV2Params) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListInstallableAddonsV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
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
