// Code generated by go-swagger; DO NOT EDIT.

package openstack

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

// NewListOpenstackSizesNoCredentialsParams creates a new ListOpenstackSizesNoCredentialsParams object
// with the default values initialized.
func NewListOpenstackSizesNoCredentialsParams() *ListOpenstackSizesNoCredentialsParams {
	var ()
	return &ListOpenstackSizesNoCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListOpenstackSizesNoCredentialsParamsWithTimeout creates a new ListOpenstackSizesNoCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListOpenstackSizesNoCredentialsParamsWithTimeout(timeout time.Duration) *ListOpenstackSizesNoCredentialsParams {
	var ()
	return &ListOpenstackSizesNoCredentialsParams{

		timeout: timeout,
	}
}

// NewListOpenstackSizesNoCredentialsParamsWithContext creates a new ListOpenstackSizesNoCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListOpenstackSizesNoCredentialsParamsWithContext(ctx context.Context) *ListOpenstackSizesNoCredentialsParams {
	var ()
	return &ListOpenstackSizesNoCredentialsParams{

		Context: ctx,
	}
}

// NewListOpenstackSizesNoCredentialsParamsWithHTTPClient creates a new ListOpenstackSizesNoCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListOpenstackSizesNoCredentialsParamsWithHTTPClient(client *http.Client) *ListOpenstackSizesNoCredentialsParams {
	var ()
	return &ListOpenstackSizesNoCredentialsParams{
		HTTPClient: client,
	}
}

/*ListOpenstackSizesNoCredentialsParams contains all the parameters to send to the API endpoint
for the list openstack sizes no credentials operation typically these are written to a http.Request
*/
type ListOpenstackSizesNoCredentialsParams struct {

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

// WithTimeout adds the timeout to the list openstack sizes no credentials params
func (o *ListOpenstackSizesNoCredentialsParams) WithTimeout(timeout time.Duration) *ListOpenstackSizesNoCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list openstack sizes no credentials params
func (o *ListOpenstackSizesNoCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list openstack sizes no credentials params
func (o *ListOpenstackSizesNoCredentialsParams) WithContext(ctx context.Context) *ListOpenstackSizesNoCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list openstack sizes no credentials params
func (o *ListOpenstackSizesNoCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list openstack sizes no credentials params
func (o *ListOpenstackSizesNoCredentialsParams) WithHTTPClient(client *http.Client) *ListOpenstackSizesNoCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list openstack sizes no credentials params
func (o *ListOpenstackSizesNoCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list openstack sizes no credentials params
func (o *ListOpenstackSizesNoCredentialsParams) WithClusterID(clusterID string) *ListOpenstackSizesNoCredentialsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list openstack sizes no credentials params
func (o *ListOpenstackSizesNoCredentialsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the list openstack sizes no credentials params
func (o *ListOpenstackSizesNoCredentialsParams) WithDC(dc string) *ListOpenstackSizesNoCredentialsParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list openstack sizes no credentials params
func (o *ListOpenstackSizesNoCredentialsParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the list openstack sizes no credentials params
func (o *ListOpenstackSizesNoCredentialsParams) WithProjectID(projectID string) *ListOpenstackSizesNoCredentialsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list openstack sizes no credentials params
func (o *ListOpenstackSizesNoCredentialsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListOpenstackSizesNoCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
