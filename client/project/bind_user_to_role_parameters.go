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

	"github.com/kubermatic/go-kubermatic/models"
)

// NewBindUserToRoleParams creates a new BindUserToRoleParams object
// with the default values initialized.
func NewBindUserToRoleParams() *BindUserToRoleParams {
	var ()
	return &BindUserToRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBindUserToRoleParamsWithTimeout creates a new BindUserToRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBindUserToRoleParamsWithTimeout(timeout time.Duration) *BindUserToRoleParams {
	var ()
	return &BindUserToRoleParams{

		timeout: timeout,
	}
}

// NewBindUserToRoleParamsWithContext creates a new BindUserToRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewBindUserToRoleParamsWithContext(ctx context.Context) *BindUserToRoleParams {
	var ()
	return &BindUserToRoleParams{

		Context: ctx,
	}
}

// NewBindUserToRoleParamsWithHTTPClient creates a new BindUserToRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBindUserToRoleParamsWithHTTPClient(client *http.Client) *BindUserToRoleParams {
	var ()
	return &BindUserToRoleParams{
		HTTPClient: client,
	}
}

/*BindUserToRoleParams contains all the parameters to send to the API endpoint
for the bind user to role operation typically these are written to a http.Request
*/
type BindUserToRoleParams struct {

	/*Body*/
	Body *models.RoleUser
	/*ClusterID*/
	ClusterID string
	/*Dc*/
	DC string
	/*Namespace*/
	Namespace string
	/*ProjectID*/
	ProjectID string
	/*RoleID*/
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the bind user to role params
func (o *BindUserToRoleParams) WithTimeout(timeout time.Duration) *BindUserToRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bind user to role params
func (o *BindUserToRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bind user to role params
func (o *BindUserToRoleParams) WithContext(ctx context.Context) *BindUserToRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bind user to role params
func (o *BindUserToRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bind user to role params
func (o *BindUserToRoleParams) WithHTTPClient(client *http.Client) *BindUserToRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bind user to role params
func (o *BindUserToRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the bind user to role params
func (o *BindUserToRoleParams) WithBody(body *models.RoleUser) *BindUserToRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the bind user to role params
func (o *BindUserToRoleParams) SetBody(body *models.RoleUser) {
	o.Body = body
}

// WithClusterID adds the clusterID to the bind user to role params
func (o *BindUserToRoleParams) WithClusterID(clusterID string) *BindUserToRoleParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the bind user to role params
func (o *BindUserToRoleParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the bind user to role params
func (o *BindUserToRoleParams) WithDC(dc string) *BindUserToRoleParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the bind user to role params
func (o *BindUserToRoleParams) SetDC(dc string) {
	o.DC = dc
}

// WithNamespace adds the namespace to the bind user to role params
func (o *BindUserToRoleParams) WithNamespace(namespace string) *BindUserToRoleParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the bind user to role params
func (o *BindUserToRoleParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectID adds the projectID to the bind user to role params
func (o *BindUserToRoleParams) WithProjectID(projectID string) *BindUserToRoleParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the bind user to role params
func (o *BindUserToRoleParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithRoleID adds the roleID to the bind user to role params
func (o *BindUserToRoleParams) WithRoleID(roleID string) *BindUserToRoleParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the bind user to role params
func (o *BindUserToRoleParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *BindUserToRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	// path param role_id
	if err := r.SetPathParam("role_id", o.RoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
