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

// NewGetClusterTemplateInstanceParams creates a new GetClusterTemplateInstanceParams object
// with the default values initialized.
func NewGetClusterTemplateInstanceParams() *GetClusterTemplateInstanceParams {
	var ()
	return &GetClusterTemplateInstanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterTemplateInstanceParamsWithTimeout creates a new GetClusterTemplateInstanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterTemplateInstanceParamsWithTimeout(timeout time.Duration) *GetClusterTemplateInstanceParams {
	var ()
	return &GetClusterTemplateInstanceParams{

		timeout: timeout,
	}
}

// NewGetClusterTemplateInstanceParamsWithContext creates a new GetClusterTemplateInstanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterTemplateInstanceParamsWithContext(ctx context.Context) *GetClusterTemplateInstanceParams {
	var ()
	return &GetClusterTemplateInstanceParams{

		Context: ctx,
	}
}

// NewGetClusterTemplateInstanceParamsWithHTTPClient creates a new GetClusterTemplateInstanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterTemplateInstanceParamsWithHTTPClient(client *http.Client) *GetClusterTemplateInstanceParams {
	var ()
	return &GetClusterTemplateInstanceParams{
		HTTPClient: client,
	}
}

/*GetClusterTemplateInstanceParams contains all the parameters to send to the API endpoint
for the get cluster template instance operation typically these are written to a http.Request
*/
type GetClusterTemplateInstanceParams struct {

	/*InstanceID*/
	ClusterTemplateInstanceID string
	/*ProjectID*/
	ProjectID string
	/*TemplateID*/
	ClusterTemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster template instance params
func (o *GetClusterTemplateInstanceParams) WithTimeout(timeout time.Duration) *GetClusterTemplateInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster template instance params
func (o *GetClusterTemplateInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster template instance params
func (o *GetClusterTemplateInstanceParams) WithContext(ctx context.Context) *GetClusterTemplateInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster template instance params
func (o *GetClusterTemplateInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster template instance params
func (o *GetClusterTemplateInstanceParams) WithHTTPClient(client *http.Client) *GetClusterTemplateInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster template instance params
func (o *GetClusterTemplateInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterTemplateInstanceID adds the instanceID to the get cluster template instance params
func (o *GetClusterTemplateInstanceParams) WithClusterTemplateInstanceID(instanceID string) *GetClusterTemplateInstanceParams {
	o.SetClusterTemplateInstanceID(instanceID)
	return o
}

// SetClusterTemplateInstanceID adds the instanceId to the get cluster template instance params
func (o *GetClusterTemplateInstanceParams) SetClusterTemplateInstanceID(instanceID string) {
	o.ClusterTemplateInstanceID = instanceID
}

// WithProjectID adds the projectID to the get cluster template instance params
func (o *GetClusterTemplateInstanceParams) WithProjectID(projectID string) *GetClusterTemplateInstanceParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get cluster template instance params
func (o *GetClusterTemplateInstanceParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithClusterTemplateID adds the templateID to the get cluster template instance params
func (o *GetClusterTemplateInstanceParams) WithClusterTemplateID(templateID string) *GetClusterTemplateInstanceParams {
	o.SetClusterTemplateID(templateID)
	return o
}

// SetClusterTemplateID adds the templateId to the get cluster template instance params
func (o *GetClusterTemplateInstanceParams) SetClusterTemplateID(templateID string) {
	o.ClusterTemplateID = templateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterTemplateInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param instance_id
	if err := r.SetPathParam("instance_id", o.ClusterTemplateInstanceID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	// path param template_id
	if err := r.SetPathParam("template_id", o.ClusterTemplateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
