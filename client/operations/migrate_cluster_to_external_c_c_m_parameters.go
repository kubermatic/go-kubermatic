// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewMigrateClusterToExternalCCMParams creates a new MigrateClusterToExternalCCMParams object
// with the default values initialized.
func NewMigrateClusterToExternalCCMParams() *MigrateClusterToExternalCCMParams {
	var ()
	return &MigrateClusterToExternalCCMParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMigrateClusterToExternalCCMParamsWithTimeout creates a new MigrateClusterToExternalCCMParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMigrateClusterToExternalCCMParamsWithTimeout(timeout time.Duration) *MigrateClusterToExternalCCMParams {
	var ()
	return &MigrateClusterToExternalCCMParams{

		timeout: timeout,
	}
}

// NewMigrateClusterToExternalCCMParamsWithContext creates a new MigrateClusterToExternalCCMParams object
// with the default values initialized, and the ability to set a context for a request
func NewMigrateClusterToExternalCCMParamsWithContext(ctx context.Context) *MigrateClusterToExternalCCMParams {
	var ()
	return &MigrateClusterToExternalCCMParams{

		Context: ctx,
	}
}

// NewMigrateClusterToExternalCCMParamsWithHTTPClient creates a new MigrateClusterToExternalCCMParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMigrateClusterToExternalCCMParamsWithHTTPClient(client *http.Client) *MigrateClusterToExternalCCMParams {
	var ()
	return &MigrateClusterToExternalCCMParams{
		HTTPClient: client,
	}
}

/*MigrateClusterToExternalCCMParams contains all the parameters to send to the API endpoint
for the migrate cluster to external c c m operation typically these are written to a http.Request
*/
type MigrateClusterToExternalCCMParams struct {

	/*ClusterID*/
	ClusterID string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the migrate cluster to external c c m params
func (o *MigrateClusterToExternalCCMParams) WithTimeout(timeout time.Duration) *MigrateClusterToExternalCCMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the migrate cluster to external c c m params
func (o *MigrateClusterToExternalCCMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the migrate cluster to external c c m params
func (o *MigrateClusterToExternalCCMParams) WithContext(ctx context.Context) *MigrateClusterToExternalCCMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the migrate cluster to external c c m params
func (o *MigrateClusterToExternalCCMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the migrate cluster to external c c m params
func (o *MigrateClusterToExternalCCMParams) WithHTTPClient(client *http.Client) *MigrateClusterToExternalCCMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the migrate cluster to external c c m params
func (o *MigrateClusterToExternalCCMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the migrate cluster to external c c m params
func (o *MigrateClusterToExternalCCMParams) WithClusterID(clusterID string) *MigrateClusterToExternalCCMParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the migrate cluster to external c c m params
func (o *MigrateClusterToExternalCCMParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the migrate cluster to external c c m params
func (o *MigrateClusterToExternalCCMParams) WithProjectID(projectID string) *MigrateClusterToExternalCCMParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the migrate cluster to external c c m params
func (o *MigrateClusterToExternalCCMParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *MigrateClusterToExternalCCMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
