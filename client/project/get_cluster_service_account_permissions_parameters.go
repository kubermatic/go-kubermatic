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

// NewGetClusterServiceAccountPermissionsParams creates a new GetClusterServiceAccountPermissionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterServiceAccountPermissionsParams() *GetClusterServiceAccountPermissionsParams {
	return &GetClusterServiceAccountPermissionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterServiceAccountPermissionsParamsWithTimeout creates a new GetClusterServiceAccountPermissionsParams object
// with the ability to set a timeout on a request.
func NewGetClusterServiceAccountPermissionsParamsWithTimeout(timeout time.Duration) *GetClusterServiceAccountPermissionsParams {
	return &GetClusterServiceAccountPermissionsParams{
		timeout: timeout,
	}
}

// NewGetClusterServiceAccountPermissionsParamsWithContext creates a new GetClusterServiceAccountPermissionsParams object
// with the ability to set a context for a request.
func NewGetClusterServiceAccountPermissionsParamsWithContext(ctx context.Context) *GetClusterServiceAccountPermissionsParams {
	return &GetClusterServiceAccountPermissionsParams{
		Context: ctx,
	}
}

// NewGetClusterServiceAccountPermissionsParamsWithHTTPClient creates a new GetClusterServiceAccountPermissionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterServiceAccountPermissionsParamsWithHTTPClient(client *http.Client) *GetClusterServiceAccountPermissionsParams {
	return &GetClusterServiceAccountPermissionsParams{
		HTTPClient: client,
	}
}

/*
GetClusterServiceAccountPermissionsParams contains all the parameters to send to the API endpoint

	for the get cluster service account permissions operation.

	Typically these are written to a http.Request.
*/
type GetClusterServiceAccountPermissionsParams struct {

	// ClusterID.
	ClusterID string

	// Namespace.
	Namespace string

	// ProjectID.
	ProjectID string

	// ServiceAccountID.
	ServiceAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster service account permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterServiceAccountPermissionsParams) WithDefaults() *GetClusterServiceAccountPermissionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster service account permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterServiceAccountPermissionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster service account permissions params
func (o *GetClusterServiceAccountPermissionsParams) WithTimeout(timeout time.Duration) *GetClusterServiceAccountPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster service account permissions params
func (o *GetClusterServiceAccountPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster service account permissions params
func (o *GetClusterServiceAccountPermissionsParams) WithContext(ctx context.Context) *GetClusterServiceAccountPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster service account permissions params
func (o *GetClusterServiceAccountPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster service account permissions params
func (o *GetClusterServiceAccountPermissionsParams) WithHTTPClient(client *http.Client) *GetClusterServiceAccountPermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster service account permissions params
func (o *GetClusterServiceAccountPermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get cluster service account permissions params
func (o *GetClusterServiceAccountPermissionsParams) WithClusterID(clusterID string) *GetClusterServiceAccountPermissionsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster service account permissions params
func (o *GetClusterServiceAccountPermissionsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithNamespace adds the namespace to the get cluster service account permissions params
func (o *GetClusterServiceAccountPermissionsParams) WithNamespace(namespace string) *GetClusterServiceAccountPermissionsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get cluster service account permissions params
func (o *GetClusterServiceAccountPermissionsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectID adds the projectID to the get cluster service account permissions params
func (o *GetClusterServiceAccountPermissionsParams) WithProjectID(projectID string) *GetClusterServiceAccountPermissionsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get cluster service account permissions params
func (o *GetClusterServiceAccountPermissionsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithServiceAccountID adds the serviceAccountID to the get cluster service account permissions params
func (o *GetClusterServiceAccountPermissionsParams) WithServiceAccountID(serviceAccountID string) *GetClusterServiceAccountPermissionsParams {
	o.SetServiceAccountID(serviceAccountID)
	return o
}

// SetServiceAccountID adds the serviceAccountId to the get cluster service account permissions params
func (o *GetClusterServiceAccountPermissionsParams) SetServiceAccountID(serviceAccountID string) {
	o.ServiceAccountID = serviceAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterServiceAccountPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param service_account_id
	if err := r.SetPathParam("service_account_id", o.ServiceAccountID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
