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

	strfmt "github.com/go-openapi/strfmt"
)

// NewRevokeClusterViewerTokenParams creates a new RevokeClusterViewerTokenParams object
// with the default values initialized.
func NewRevokeClusterViewerTokenParams() *RevokeClusterViewerTokenParams {
	var ()
	return &RevokeClusterViewerTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRevokeClusterViewerTokenParamsWithTimeout creates a new RevokeClusterViewerTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRevokeClusterViewerTokenParamsWithTimeout(timeout time.Duration) *RevokeClusterViewerTokenParams {
	var ()
	return &RevokeClusterViewerTokenParams{

		timeout: timeout,
	}
}

// NewRevokeClusterViewerTokenParamsWithContext creates a new RevokeClusterViewerTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewRevokeClusterViewerTokenParamsWithContext(ctx context.Context) *RevokeClusterViewerTokenParams {
	var ()
	return &RevokeClusterViewerTokenParams{

		Context: ctx,
	}
}

// NewRevokeClusterViewerTokenParamsWithHTTPClient creates a new RevokeClusterViewerTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRevokeClusterViewerTokenParamsWithHTTPClient(client *http.Client) *RevokeClusterViewerTokenParams {
	var ()
	return &RevokeClusterViewerTokenParams{
		HTTPClient: client,
	}
}

/*RevokeClusterViewerTokenParams contains all the parameters to send to the API endpoint
for the revoke cluster viewer token operation typically these are written to a http.Request
*/
type RevokeClusterViewerTokenParams struct {

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

// WithTimeout adds the timeout to the revoke cluster viewer token params
func (o *RevokeClusterViewerTokenParams) WithTimeout(timeout time.Duration) *RevokeClusterViewerTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revoke cluster viewer token params
func (o *RevokeClusterViewerTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revoke cluster viewer token params
func (o *RevokeClusterViewerTokenParams) WithContext(ctx context.Context) *RevokeClusterViewerTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revoke cluster viewer token params
func (o *RevokeClusterViewerTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revoke cluster viewer token params
func (o *RevokeClusterViewerTokenParams) WithHTTPClient(client *http.Client) *RevokeClusterViewerTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revoke cluster viewer token params
func (o *RevokeClusterViewerTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the revoke cluster viewer token params
func (o *RevokeClusterViewerTokenParams) WithClusterID(clusterID string) *RevokeClusterViewerTokenParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the revoke cluster viewer token params
func (o *RevokeClusterViewerTokenParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the revoke cluster viewer token params
func (o *RevokeClusterViewerTokenParams) WithDC(dc string) *RevokeClusterViewerTokenParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the revoke cluster viewer token params
func (o *RevokeClusterViewerTokenParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the revoke cluster viewer token params
func (o *RevokeClusterViewerTokenParams) WithProjectID(projectID string) *RevokeClusterViewerTokenParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the revoke cluster viewer token params
func (o *RevokeClusterViewerTokenParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *RevokeClusterViewerTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
