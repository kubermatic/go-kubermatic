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

// NewCreateConstraintParams creates a new CreateConstraintParams object
// with the default values initialized.
func NewCreateConstraintParams() *CreateConstraintParams {
	var ()
	return &CreateConstraintParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateConstraintParamsWithTimeout creates a new CreateConstraintParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateConstraintParamsWithTimeout(timeout time.Duration) *CreateConstraintParams {
	var ()
	return &CreateConstraintParams{

		timeout: timeout,
	}
}

// NewCreateConstraintParamsWithContext creates a new CreateConstraintParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateConstraintParamsWithContext(ctx context.Context) *CreateConstraintParams {
	var ()
	return &CreateConstraintParams{

		Context: ctx,
	}
}

// NewCreateConstraintParamsWithHTTPClient creates a new CreateConstraintParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateConstraintParamsWithHTTPClient(client *http.Client) *CreateConstraintParams {
	var ()
	return &CreateConstraintParams{
		HTTPClient: client,
	}
}

/*CreateConstraintParams contains all the parameters to send to the API endpoint
for the create constraint operation typically these are written to a http.Request
*/
type CreateConstraintParams struct {

	/*Body*/
	Body *models.ConstraintBody
	/*ClusterID*/
	ClusterID string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create constraint params
func (o *CreateConstraintParams) WithTimeout(timeout time.Duration) *CreateConstraintParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create constraint params
func (o *CreateConstraintParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create constraint params
func (o *CreateConstraintParams) WithContext(ctx context.Context) *CreateConstraintParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create constraint params
func (o *CreateConstraintParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create constraint params
func (o *CreateConstraintParams) WithHTTPClient(client *http.Client) *CreateConstraintParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create constraint params
func (o *CreateConstraintParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create constraint params
func (o *CreateConstraintParams) WithBody(body *models.ConstraintBody) *CreateConstraintParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create constraint params
func (o *CreateConstraintParams) SetBody(body *models.ConstraintBody) {
	o.Body = body
}

// WithClusterID adds the clusterID to the create constraint params
func (o *CreateConstraintParams) WithClusterID(clusterID string) *CreateConstraintParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the create constraint params
func (o *CreateConstraintParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the create constraint params
func (o *CreateConstraintParams) WithProjectID(projectID string) *CreateConstraintParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create constraint params
func (o *CreateConstraintParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateConstraintParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
