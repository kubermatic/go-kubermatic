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

	"github.com/kubermatic/go-kubermatic/models"
)

// NewCreateAddonParams creates a new CreateAddonParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAddonParams() *CreateAddonParams {
	return &CreateAddonParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAddonParamsWithTimeout creates a new CreateAddonParams object
// with the ability to set a timeout on a request.
func NewCreateAddonParamsWithTimeout(timeout time.Duration) *CreateAddonParams {
	return &CreateAddonParams{
		timeout: timeout,
	}
}

// NewCreateAddonParamsWithContext creates a new CreateAddonParams object
// with the ability to set a context for a request.
func NewCreateAddonParamsWithContext(ctx context.Context) *CreateAddonParams {
	return &CreateAddonParams{
		Context: ctx,
	}
}

// NewCreateAddonParamsWithHTTPClient creates a new CreateAddonParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAddonParamsWithHTTPClient(client *http.Client) *CreateAddonParams {
	return &CreateAddonParams{
		HTTPClient: client,
	}
}

/* CreateAddonParams contains all the parameters to send to the API endpoint
   for the create addon operation.

   Typically these are written to a http.Request.
*/
type CreateAddonParams struct {

	// Body.
	Body *models.Addon

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

// WithDefaults hydrates default values in the create addon params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAddonParams) WithDefaults() *CreateAddonParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create addon params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAddonParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create addon params
func (o *CreateAddonParams) WithTimeout(timeout time.Duration) *CreateAddonParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create addon params
func (o *CreateAddonParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create addon params
func (o *CreateAddonParams) WithContext(ctx context.Context) *CreateAddonParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create addon params
func (o *CreateAddonParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create addon params
func (o *CreateAddonParams) WithHTTPClient(client *http.Client) *CreateAddonParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create addon params
func (o *CreateAddonParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create addon params
func (o *CreateAddonParams) WithBody(body *models.Addon) *CreateAddonParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create addon params
func (o *CreateAddonParams) SetBody(body *models.Addon) {
	o.Body = body
}

// WithClusterID adds the clusterID to the create addon params
func (o *CreateAddonParams) WithClusterID(clusterID string) *CreateAddonParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the create addon params
func (o *CreateAddonParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the create addon params
func (o *CreateAddonParams) WithDC(dc string) *CreateAddonParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the create addon params
func (o *CreateAddonParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the create addon params
func (o *CreateAddonParams) WithProjectID(projectID string) *CreateAddonParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create addon params
func (o *CreateAddonParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAddonParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
