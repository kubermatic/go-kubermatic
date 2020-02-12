// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListVSphereNetworksNoCredentialsParams creates a new ListVSphereNetworksNoCredentialsParams object
// with the default values initialized.
func NewListVSphereNetworksNoCredentialsParams() *ListVSphereNetworksNoCredentialsParams {
	var ()
	return &ListVSphereNetworksNoCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListVSphereNetworksNoCredentialsParamsWithTimeout creates a new ListVSphereNetworksNoCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListVSphereNetworksNoCredentialsParamsWithTimeout(timeout time.Duration) *ListVSphereNetworksNoCredentialsParams {
	var ()
	return &ListVSphereNetworksNoCredentialsParams{

		timeout: timeout,
	}
}

// NewListVSphereNetworksNoCredentialsParamsWithContext creates a new ListVSphereNetworksNoCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListVSphereNetworksNoCredentialsParamsWithContext(ctx context.Context) *ListVSphereNetworksNoCredentialsParams {
	var ()
	return &ListVSphereNetworksNoCredentialsParams{

		Context: ctx,
	}
}

// NewListVSphereNetworksNoCredentialsParamsWithHTTPClient creates a new ListVSphereNetworksNoCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListVSphereNetworksNoCredentialsParamsWithHTTPClient(client *http.Client) *ListVSphereNetworksNoCredentialsParams {
	var ()
	return &ListVSphereNetworksNoCredentialsParams{
		HTTPClient: client,
	}
}

/*ListVSphereNetworksNoCredentialsParams contains all the parameters to send to the API endpoint
for the list v sphere networks no credentials operation typically these are written to a http.Request
*/
type ListVSphereNetworksNoCredentialsParams struct {

	/*ClusterID*/
	ClusterID string
	/*Dc*/
	DC string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list v sphere networks no credentials params
func (o *ListVSphereNetworksNoCredentialsParams) WithTimeout(timeout time.Duration) *ListVSphereNetworksNoCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list v sphere networks no credentials params
func (o *ListVSphereNetworksNoCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list v sphere networks no credentials params
func (o *ListVSphereNetworksNoCredentialsParams) WithContext(ctx context.Context) *ListVSphereNetworksNoCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list v sphere networks no credentials params
func (o *ListVSphereNetworksNoCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list v sphere networks no credentials params
func (o *ListVSphereNetworksNoCredentialsParams) WithHTTPClient(client *http.Client) *ListVSphereNetworksNoCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list v sphere networks no credentials params
func (o *ListVSphereNetworksNoCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list v sphere networks no credentials params
func (o *ListVSphereNetworksNoCredentialsParams) WithClusterID(clusterID string) *ListVSphereNetworksNoCredentialsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list v sphere networks no credentials params
func (o *ListVSphereNetworksNoCredentialsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the list v sphere networks no credentials params
func (o *ListVSphereNetworksNoCredentialsParams) WithDC(dc string) *ListVSphereNetworksNoCredentialsParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list v sphere networks no credentials params
func (o *ListVSphereNetworksNoCredentialsParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the list v sphere networks no credentials params
func (o *ListVSphereNetworksNoCredentialsParams) WithProjectID(projectID string) *ListVSphereNetworksNoCredentialsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list v sphere networks no credentials params
func (o *ListVSphereNetworksNoCredentialsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListVSphereNetworksNoCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
