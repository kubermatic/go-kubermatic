// Code generated by go-swagger; DO NOT EDIT.

package addon

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

// NewListAddonsV2Params creates a new ListAddonsV2Params object
// with the default values initialized.
func NewListAddonsV2Params() *ListAddonsV2Params {
	var ()
	return &ListAddonsV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAddonsV2ParamsWithTimeout creates a new ListAddonsV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAddonsV2ParamsWithTimeout(timeout time.Duration) *ListAddonsV2Params {
	var ()
	return &ListAddonsV2Params{

		timeout: timeout,
	}
}

// NewListAddonsV2ParamsWithContext creates a new ListAddonsV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewListAddonsV2ParamsWithContext(ctx context.Context) *ListAddonsV2Params {
	var ()
	return &ListAddonsV2Params{

		Context: ctx,
	}
}

// NewListAddonsV2ParamsWithHTTPClient creates a new ListAddonsV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAddonsV2ParamsWithHTTPClient(client *http.Client) *ListAddonsV2Params {
	var ()
	return &ListAddonsV2Params{
		HTTPClient: client,
	}
}

/*ListAddonsV2Params contains all the parameters to send to the API endpoint
for the list addons v2 operation typically these are written to a http.Request
*/
type ListAddonsV2Params struct {

	/*ClusterID*/
	ClusterID string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list addons v2 params
func (o *ListAddonsV2Params) WithTimeout(timeout time.Duration) *ListAddonsV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list addons v2 params
func (o *ListAddonsV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list addons v2 params
func (o *ListAddonsV2Params) WithContext(ctx context.Context) *ListAddonsV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list addons v2 params
func (o *ListAddonsV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list addons v2 params
func (o *ListAddonsV2Params) WithHTTPClient(client *http.Client) *ListAddonsV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list addons v2 params
func (o *ListAddonsV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list addons v2 params
func (o *ListAddonsV2Params) WithClusterID(clusterID string) *ListAddonsV2Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list addons v2 params
func (o *ListAddonsV2Params) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the list addons v2 params
func (o *ListAddonsV2Params) WithProjectID(projectID string) *ListAddonsV2Params {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list addons v2 params
func (o *ListAddonsV2Params) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListAddonsV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
