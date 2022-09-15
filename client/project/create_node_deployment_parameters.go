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

// NewCreateNodeDeploymentParams creates a new CreateNodeDeploymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNodeDeploymentParams() *CreateNodeDeploymentParams {
	return &CreateNodeDeploymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNodeDeploymentParamsWithTimeout creates a new CreateNodeDeploymentParams object
// with the ability to set a timeout on a request.
func NewCreateNodeDeploymentParamsWithTimeout(timeout time.Duration) *CreateNodeDeploymentParams {
	return &CreateNodeDeploymentParams{
		timeout: timeout,
	}
}

// NewCreateNodeDeploymentParamsWithContext creates a new CreateNodeDeploymentParams object
// with the ability to set a context for a request.
func NewCreateNodeDeploymentParamsWithContext(ctx context.Context) *CreateNodeDeploymentParams {
	return &CreateNodeDeploymentParams{
		Context: ctx,
	}
}

// NewCreateNodeDeploymentParamsWithHTTPClient creates a new CreateNodeDeploymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNodeDeploymentParamsWithHTTPClient(client *http.Client) *CreateNodeDeploymentParams {
	return &CreateNodeDeploymentParams{
		HTTPClient: client,
	}
}

/*
CreateNodeDeploymentParams contains all the parameters to send to the API endpoint

	for the create node deployment operation.

	Typically these are written to a http.Request.
*/
type CreateNodeDeploymentParams struct {

	// Body.
	Body *models.NodeDeployment

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

// WithDefaults hydrates default values in the create node deployment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNodeDeploymentParams) WithDefaults() *CreateNodeDeploymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create node deployment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNodeDeploymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create node deployment params
func (o *CreateNodeDeploymentParams) WithTimeout(timeout time.Duration) *CreateNodeDeploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create node deployment params
func (o *CreateNodeDeploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create node deployment params
func (o *CreateNodeDeploymentParams) WithContext(ctx context.Context) *CreateNodeDeploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create node deployment params
func (o *CreateNodeDeploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create node deployment params
func (o *CreateNodeDeploymentParams) WithHTTPClient(client *http.Client) *CreateNodeDeploymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create node deployment params
func (o *CreateNodeDeploymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create node deployment params
func (o *CreateNodeDeploymentParams) WithBody(body *models.NodeDeployment) *CreateNodeDeploymentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create node deployment params
func (o *CreateNodeDeploymentParams) SetBody(body *models.NodeDeployment) {
	o.Body = body
}

// WithClusterID adds the clusterID to the create node deployment params
func (o *CreateNodeDeploymentParams) WithClusterID(clusterID string) *CreateNodeDeploymentParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the create node deployment params
func (o *CreateNodeDeploymentParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the create node deployment params
func (o *CreateNodeDeploymentParams) WithDC(dc string) *CreateNodeDeploymentParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the create node deployment params
func (o *CreateNodeDeploymentParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the create node deployment params
func (o *CreateNodeDeploymentParams) WithProjectID(projectID string) *CreateNodeDeploymentParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create node deployment params
func (o *CreateNodeDeploymentParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNodeDeploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
