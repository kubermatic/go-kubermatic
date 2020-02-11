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

	models "github.com/kubermatic/go-kubermatic/models"
)

// NewCreateSSHKeyParams creates a new CreateSSHKeyParams object
// with the default values initialized.
func NewCreateSSHKeyParams() *CreateSSHKeyParams {
	var ()
	return &CreateSSHKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSSHKeyParamsWithTimeout creates a new CreateSSHKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSSHKeyParamsWithTimeout(timeout time.Duration) *CreateSSHKeyParams {
	var ()
	return &CreateSSHKeyParams{

		timeout: timeout,
	}
}

// NewCreateSSHKeyParamsWithContext creates a new CreateSSHKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSSHKeyParamsWithContext(ctx context.Context) *CreateSSHKeyParams {
	var ()
	return &CreateSSHKeyParams{

		Context: ctx,
	}
}

// NewCreateSSHKeyParamsWithHTTPClient creates a new CreateSSHKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSSHKeyParamsWithHTTPClient(client *http.Client) *CreateSSHKeyParams {
	var ()
	return &CreateSSHKeyParams{
		HTTPClient: client,
	}
}

/*CreateSSHKeyParams contains all the parameters to send to the API endpoint
for the create SSH key operation typically these are written to a http.Request
*/
type CreateSSHKeyParams struct {

	/*Key*/
	Key *models.SSHKey
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create SSH key params
func (o *CreateSSHKeyParams) WithTimeout(timeout time.Duration) *CreateSSHKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create SSH key params
func (o *CreateSSHKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create SSH key params
func (o *CreateSSHKeyParams) WithContext(ctx context.Context) *CreateSSHKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create SSH key params
func (o *CreateSSHKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create SSH key params
func (o *CreateSSHKeyParams) WithHTTPClient(client *http.Client) *CreateSSHKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create SSH key params
func (o *CreateSSHKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the create SSH key params
func (o *CreateSSHKeyParams) WithKey(key *models.SSHKey) *CreateSSHKeyParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the create SSH key params
func (o *CreateSSHKeyParams) SetKey(key *models.SSHKey) {
	o.Key = key
}

// WithProjectID adds the projectID to the create SSH key params
func (o *CreateSSHKeyParams) WithProjectID(projectID string) *CreateSSHKeyParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create SSH key params
func (o *CreateSSHKeyParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSSHKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Key != nil {
		if err := r.SetBodyParam(o.Key); err != nil {
			return err
		}
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
