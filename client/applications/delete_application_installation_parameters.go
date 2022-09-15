// Code generated by go-swagger; DO NOT EDIT.

package applications

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

// NewDeleteApplicationInstallationParams creates a new DeleteApplicationInstallationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteApplicationInstallationParams() *DeleteApplicationInstallationParams {
	return &DeleteApplicationInstallationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteApplicationInstallationParamsWithTimeout creates a new DeleteApplicationInstallationParams object
// with the ability to set a timeout on a request.
func NewDeleteApplicationInstallationParamsWithTimeout(timeout time.Duration) *DeleteApplicationInstallationParams {
	return &DeleteApplicationInstallationParams{
		timeout: timeout,
	}
}

// NewDeleteApplicationInstallationParamsWithContext creates a new DeleteApplicationInstallationParams object
// with the ability to set a context for a request.
func NewDeleteApplicationInstallationParamsWithContext(ctx context.Context) *DeleteApplicationInstallationParams {
	return &DeleteApplicationInstallationParams{
		Context: ctx,
	}
}

// NewDeleteApplicationInstallationParamsWithHTTPClient creates a new DeleteApplicationInstallationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteApplicationInstallationParamsWithHTTPClient(client *http.Client) *DeleteApplicationInstallationParams {
	return &DeleteApplicationInstallationParams{
		HTTPClient: client,
	}
}

/*
DeleteApplicationInstallationParams contains all the parameters to send to the API endpoint

	for the delete application installation operation.

	Typically these are written to a http.Request.
*/
type DeleteApplicationInstallationParams struct {

	// AppinstallName.
	ApplicationInstallationName string

	// ClusterID.
	ClusterID string

	// Namespace.
	Namespace string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete application installation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteApplicationInstallationParams) WithDefaults() *DeleteApplicationInstallationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete application installation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteApplicationInstallationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete application installation params
func (o *DeleteApplicationInstallationParams) WithTimeout(timeout time.Duration) *DeleteApplicationInstallationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete application installation params
func (o *DeleteApplicationInstallationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete application installation params
func (o *DeleteApplicationInstallationParams) WithContext(ctx context.Context) *DeleteApplicationInstallationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete application installation params
func (o *DeleteApplicationInstallationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete application installation params
func (o *DeleteApplicationInstallationParams) WithHTTPClient(client *http.Client) *DeleteApplicationInstallationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete application installation params
func (o *DeleteApplicationInstallationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationInstallationName adds the appinstallName to the delete application installation params
func (o *DeleteApplicationInstallationParams) WithApplicationInstallationName(appinstallName string) *DeleteApplicationInstallationParams {
	o.SetApplicationInstallationName(appinstallName)
	return o
}

// SetApplicationInstallationName adds the appinstallName to the delete application installation params
func (o *DeleteApplicationInstallationParams) SetApplicationInstallationName(appinstallName string) {
	o.ApplicationInstallationName = appinstallName
}

// WithClusterID adds the clusterID to the delete application installation params
func (o *DeleteApplicationInstallationParams) WithClusterID(clusterID string) *DeleteApplicationInstallationParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete application installation params
func (o *DeleteApplicationInstallationParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithNamespace adds the namespace to the delete application installation params
func (o *DeleteApplicationInstallationParams) WithNamespace(namespace string) *DeleteApplicationInstallationParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete application installation params
func (o *DeleteApplicationInstallationParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectID adds the projectID to the delete application installation params
func (o *DeleteApplicationInstallationParams) WithProjectID(projectID string) *DeleteApplicationInstallationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete application installation params
func (o *DeleteApplicationInstallationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteApplicationInstallationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appinstall_name
	if err := r.SetPathParam("appinstall_name", o.ApplicationInstallationName); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
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
