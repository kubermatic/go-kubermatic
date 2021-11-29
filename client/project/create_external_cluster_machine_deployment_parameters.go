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

	"github.com/kubermatic/go-kubermatic/models"
)

// NewCreateExternalClusterMachineDeploymentParams creates a new CreateExternalClusterMachineDeploymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateExternalClusterMachineDeploymentParams() *CreateExternalClusterMachineDeploymentParams {
	return &CreateExternalClusterMachineDeploymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateExternalClusterMachineDeploymentParamsWithTimeout creates a new CreateExternalClusterMachineDeploymentParams object
// with the ability to set a timeout on a request.
func NewCreateExternalClusterMachineDeploymentParamsWithTimeout(timeout time.Duration) *CreateExternalClusterMachineDeploymentParams {
	return &CreateExternalClusterMachineDeploymentParams{
		timeout: timeout,
	}
}

// NewCreateExternalClusterMachineDeploymentParamsWithContext creates a new CreateExternalClusterMachineDeploymentParams object
// with the ability to set a context for a request.
func NewCreateExternalClusterMachineDeploymentParamsWithContext(ctx context.Context) *CreateExternalClusterMachineDeploymentParams {
	return &CreateExternalClusterMachineDeploymentParams{
		Context: ctx,
	}
}

// NewCreateExternalClusterMachineDeploymentParamsWithHTTPClient creates a new CreateExternalClusterMachineDeploymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateExternalClusterMachineDeploymentParamsWithHTTPClient(client *http.Client) *CreateExternalClusterMachineDeploymentParams {
	return &CreateExternalClusterMachineDeploymentParams{
		HTTPClient: client,
	}
}

/* CreateExternalClusterMachineDeploymentParams contains all the parameters to send to the API endpoint
   for the create external cluster machine deployment operation.

   Typically these are written to a http.Request.
*/
type CreateExternalClusterMachineDeploymentParams struct {

	// Body.
	Body *models.ExternalClusterMachineDeployment

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create external cluster machine deployment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateExternalClusterMachineDeploymentParams) WithDefaults() *CreateExternalClusterMachineDeploymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create external cluster machine deployment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateExternalClusterMachineDeploymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create external cluster machine deployment params
func (o *CreateExternalClusterMachineDeploymentParams) WithTimeout(timeout time.Duration) *CreateExternalClusterMachineDeploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create external cluster machine deployment params
func (o *CreateExternalClusterMachineDeploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create external cluster machine deployment params
func (o *CreateExternalClusterMachineDeploymentParams) WithContext(ctx context.Context) *CreateExternalClusterMachineDeploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create external cluster machine deployment params
func (o *CreateExternalClusterMachineDeploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create external cluster machine deployment params
func (o *CreateExternalClusterMachineDeploymentParams) WithHTTPClient(client *http.Client) *CreateExternalClusterMachineDeploymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create external cluster machine deployment params
func (o *CreateExternalClusterMachineDeploymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create external cluster machine deployment params
func (o *CreateExternalClusterMachineDeploymentParams) WithBody(body *models.ExternalClusterMachineDeployment) *CreateExternalClusterMachineDeploymentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create external cluster machine deployment params
func (o *CreateExternalClusterMachineDeploymentParams) SetBody(body *models.ExternalClusterMachineDeployment) {
	o.Body = body
}

// WithClusterID adds the clusterID to the create external cluster machine deployment params
func (o *CreateExternalClusterMachineDeploymentParams) WithClusterID(clusterID string) *CreateExternalClusterMachineDeploymentParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the create external cluster machine deployment params
func (o *CreateExternalClusterMachineDeploymentParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the create external cluster machine deployment params
func (o *CreateExternalClusterMachineDeploymentParams) WithProjectID(projectID string) *CreateExternalClusterMachineDeploymentParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create external cluster machine deployment params
func (o *CreateExternalClusterMachineDeploymentParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateExternalClusterMachineDeploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
