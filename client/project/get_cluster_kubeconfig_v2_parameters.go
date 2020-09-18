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

// NewGetClusterKubeconfigV2Params creates a new GetClusterKubeconfigV2Params object
// with the default values initialized.
func NewGetClusterKubeconfigV2Params() *GetClusterKubeconfigV2Params {
	var ()
	return &GetClusterKubeconfigV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterKubeconfigV2ParamsWithTimeout creates a new GetClusterKubeconfigV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterKubeconfigV2ParamsWithTimeout(timeout time.Duration) *GetClusterKubeconfigV2Params {
	var ()
	return &GetClusterKubeconfigV2Params{

		timeout: timeout,
	}
}

// NewGetClusterKubeconfigV2ParamsWithContext creates a new GetClusterKubeconfigV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterKubeconfigV2ParamsWithContext(ctx context.Context) *GetClusterKubeconfigV2Params {
	var ()
	return &GetClusterKubeconfigV2Params{

		Context: ctx,
	}
}

// NewGetClusterKubeconfigV2ParamsWithHTTPClient creates a new GetClusterKubeconfigV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterKubeconfigV2ParamsWithHTTPClient(client *http.Client) *GetClusterKubeconfigV2Params {
	var ()
	return &GetClusterKubeconfigV2Params{
		HTTPClient: client,
	}
}

/*GetClusterKubeconfigV2Params contains all the parameters to send to the API endpoint
for the get cluster kubeconfig v2 operation typically these are written to a http.Request
*/
type GetClusterKubeconfigV2Params struct {

	/*ClusterID*/
	ClusterID string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster kubeconfig v2 params
func (o *GetClusterKubeconfigV2Params) WithTimeout(timeout time.Duration) *GetClusterKubeconfigV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster kubeconfig v2 params
func (o *GetClusterKubeconfigV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster kubeconfig v2 params
func (o *GetClusterKubeconfigV2Params) WithContext(ctx context.Context) *GetClusterKubeconfigV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster kubeconfig v2 params
func (o *GetClusterKubeconfigV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster kubeconfig v2 params
func (o *GetClusterKubeconfigV2Params) WithHTTPClient(client *http.Client) *GetClusterKubeconfigV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster kubeconfig v2 params
func (o *GetClusterKubeconfigV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get cluster kubeconfig v2 params
func (o *GetClusterKubeconfigV2Params) WithClusterID(clusterID string) *GetClusterKubeconfigV2Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster kubeconfig v2 params
func (o *GetClusterKubeconfigV2Params) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the get cluster kubeconfig v2 params
func (o *GetClusterKubeconfigV2Params) WithProjectID(projectID string) *GetClusterKubeconfigV2Params {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get cluster kubeconfig v2 params
func (o *GetClusterKubeconfigV2Params) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterKubeconfigV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
