// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewDeleteMachineDeploymentParams creates a new DeleteMachineDeploymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMachineDeploymentParams() *DeleteMachineDeploymentParams {
	return &DeleteMachineDeploymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMachineDeploymentParamsWithTimeout creates a new DeleteMachineDeploymentParams object
// with the ability to set a timeout on a request.
func NewDeleteMachineDeploymentParamsWithTimeout(timeout time.Duration) *DeleteMachineDeploymentParams {
	return &DeleteMachineDeploymentParams{
		timeout: timeout,
	}
}

// NewDeleteMachineDeploymentParamsWithContext creates a new DeleteMachineDeploymentParams object
// with the ability to set a context for a request.
func NewDeleteMachineDeploymentParamsWithContext(ctx context.Context) *DeleteMachineDeploymentParams {
	return &DeleteMachineDeploymentParams{
		Context: ctx,
	}
}

// NewDeleteMachineDeploymentParamsWithHTTPClient creates a new DeleteMachineDeploymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMachineDeploymentParamsWithHTTPClient(client *http.Client) *DeleteMachineDeploymentParams {
	return &DeleteMachineDeploymentParams{
		HTTPClient: client,
	}
}

/*
DeleteMachineDeploymentParams contains all the parameters to send to the API endpoint

	for the delete machine deployment operation.

	Typically these are written to a http.Request.
*/
type DeleteMachineDeploymentParams struct {

	// ClusterID.
	ClusterID string

	// MachinedeploymentID.
	MachineDeploymentID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete machine deployment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMachineDeploymentParams) WithDefaults() *DeleteMachineDeploymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete machine deployment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMachineDeploymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete machine deployment params
func (o *DeleteMachineDeploymentParams) WithTimeout(timeout time.Duration) *DeleteMachineDeploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete machine deployment params
func (o *DeleteMachineDeploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete machine deployment params
func (o *DeleteMachineDeploymentParams) WithContext(ctx context.Context) *DeleteMachineDeploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete machine deployment params
func (o *DeleteMachineDeploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete machine deployment params
func (o *DeleteMachineDeploymentParams) WithHTTPClient(client *http.Client) *DeleteMachineDeploymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete machine deployment params
func (o *DeleteMachineDeploymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the delete machine deployment params
func (o *DeleteMachineDeploymentParams) WithClusterID(clusterID string) *DeleteMachineDeploymentParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete machine deployment params
func (o *DeleteMachineDeploymentParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithMachineDeploymentID adds the machinedeploymentID to the delete machine deployment params
func (o *DeleteMachineDeploymentParams) WithMachineDeploymentID(machinedeploymentID string) *DeleteMachineDeploymentParams {
	o.SetMachineDeploymentID(machinedeploymentID)
	return o
}

// SetMachineDeploymentID adds the machinedeploymentId to the delete machine deployment params
func (o *DeleteMachineDeploymentParams) SetMachineDeploymentID(machinedeploymentID string) {
	o.MachineDeploymentID = machinedeploymentID
}

// WithProjectID adds the projectID to the delete machine deployment params
func (o *DeleteMachineDeploymentParams) WithProjectID(projectID string) *DeleteMachineDeploymentParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete machine deployment params
func (o *DeleteMachineDeploymentParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMachineDeploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param machinedeployment_id
	if err := r.SetPathParam("machinedeployment_id", o.MachineDeploymentID); err != nil {
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
