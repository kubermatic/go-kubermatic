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
	"github.com/go-openapi/strfmt"
)

// NewListVSphereFoldersNoCredentialsParams creates a new ListVSphereFoldersNoCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListVSphereFoldersNoCredentialsParams() *ListVSphereFoldersNoCredentialsParams {
	return &ListVSphereFoldersNoCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListVSphereFoldersNoCredentialsParamsWithTimeout creates a new ListVSphereFoldersNoCredentialsParams object
// with the ability to set a timeout on a request.
func NewListVSphereFoldersNoCredentialsParamsWithTimeout(timeout time.Duration) *ListVSphereFoldersNoCredentialsParams {
	return &ListVSphereFoldersNoCredentialsParams{
		timeout: timeout,
	}
}

// NewListVSphereFoldersNoCredentialsParamsWithContext creates a new ListVSphereFoldersNoCredentialsParams object
// with the ability to set a context for a request.
func NewListVSphereFoldersNoCredentialsParamsWithContext(ctx context.Context) *ListVSphereFoldersNoCredentialsParams {
	return &ListVSphereFoldersNoCredentialsParams{
		Context: ctx,
	}
}

// NewListVSphereFoldersNoCredentialsParamsWithHTTPClient creates a new ListVSphereFoldersNoCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListVSphereFoldersNoCredentialsParamsWithHTTPClient(client *http.Client) *ListVSphereFoldersNoCredentialsParams {
	return &ListVSphereFoldersNoCredentialsParams{
		HTTPClient: client,
	}
}

/*
ListVSphereFoldersNoCredentialsParams contains all the parameters to send to the API endpoint

	for the list v sphere folders no credentials operation.

	Typically these are written to a http.Request.
*/
type ListVSphereFoldersNoCredentialsParams struct {

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

// WithDefaults hydrates default values in the list v sphere folders no credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVSphereFoldersNoCredentialsParams) WithDefaults() *ListVSphereFoldersNoCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list v sphere folders no credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVSphereFoldersNoCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list v sphere folders no credentials params
func (o *ListVSphereFoldersNoCredentialsParams) WithTimeout(timeout time.Duration) *ListVSphereFoldersNoCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list v sphere folders no credentials params
func (o *ListVSphereFoldersNoCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list v sphere folders no credentials params
func (o *ListVSphereFoldersNoCredentialsParams) WithContext(ctx context.Context) *ListVSphereFoldersNoCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list v sphere folders no credentials params
func (o *ListVSphereFoldersNoCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list v sphere folders no credentials params
func (o *ListVSphereFoldersNoCredentialsParams) WithHTTPClient(client *http.Client) *ListVSphereFoldersNoCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list v sphere folders no credentials params
func (o *ListVSphereFoldersNoCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list v sphere folders no credentials params
func (o *ListVSphereFoldersNoCredentialsParams) WithClusterID(clusterID string) *ListVSphereFoldersNoCredentialsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list v sphere folders no credentials params
func (o *ListVSphereFoldersNoCredentialsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the list v sphere folders no credentials params
func (o *ListVSphereFoldersNoCredentialsParams) WithDC(dc string) *ListVSphereFoldersNoCredentialsParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list v sphere folders no credentials params
func (o *ListVSphereFoldersNoCredentialsParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the list v sphere folders no credentials params
func (o *ListVSphereFoldersNoCredentialsParams) WithProjectID(projectID string) *ListVSphereFoldersNoCredentialsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list v sphere folders no credentials params
func (o *ListVSphereFoldersNoCredentialsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListVSphereFoldersNoCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
