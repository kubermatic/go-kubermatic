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

// NewListProjectVSphereVMGroupsParams creates a new ListProjectVSphereVMGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectVSphereVMGroupsParams() *ListProjectVSphereVMGroupsParams {
	return &ListProjectVSphereVMGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectVSphereVMGroupsParamsWithTimeout creates a new ListProjectVSphereVMGroupsParams object
// with the ability to set a timeout on a request.
func NewListProjectVSphereVMGroupsParamsWithTimeout(timeout time.Duration) *ListProjectVSphereVMGroupsParams {
	return &ListProjectVSphereVMGroupsParams{
		timeout: timeout,
	}
}

// NewListProjectVSphereVMGroupsParamsWithContext creates a new ListProjectVSphereVMGroupsParams object
// with the ability to set a context for a request.
func NewListProjectVSphereVMGroupsParamsWithContext(ctx context.Context) *ListProjectVSphereVMGroupsParams {
	return &ListProjectVSphereVMGroupsParams{
		Context: ctx,
	}
}

// NewListProjectVSphereVMGroupsParamsWithHTTPClient creates a new ListProjectVSphereVMGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectVSphereVMGroupsParamsWithHTTPClient(client *http.Client) *ListProjectVSphereVMGroupsParams {
	return &ListProjectVSphereVMGroupsParams{
		HTTPClient: client,
	}
}

/*
ListProjectVSphereVMGroupsParams contains all the parameters to send to the API endpoint

	for the list project v sphere VM groups operation.

	Typically these are written to a http.Request.
*/
type ListProjectVSphereVMGroupsParams struct {

	// Credential.
	Credential *string

	// DatacenterName.
	DatacenterName *string

	// Password.
	Password *string

	// Username.
	Username *string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project v sphere VM groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectVSphereVMGroupsParams) WithDefaults() *ListProjectVSphereVMGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project v sphere VM groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectVSphereVMGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project v sphere VM groups params
func (o *ListProjectVSphereVMGroupsParams) WithTimeout(timeout time.Duration) *ListProjectVSphereVMGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project v sphere VM groups params
func (o *ListProjectVSphereVMGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project v sphere VM groups params
func (o *ListProjectVSphereVMGroupsParams) WithContext(ctx context.Context) *ListProjectVSphereVMGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project v sphere VM groups params
func (o *ListProjectVSphereVMGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project v sphere VM groups params
func (o *ListProjectVSphereVMGroupsParams) WithHTTPClient(client *http.Client) *ListProjectVSphereVMGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project v sphere VM groups params
func (o *ListProjectVSphereVMGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the list project v sphere VM groups params
func (o *ListProjectVSphereVMGroupsParams) WithCredential(credential *string) *ListProjectVSphereVMGroupsParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list project v sphere VM groups params
func (o *ListProjectVSphereVMGroupsParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithDatacenterName adds the datacenterName to the list project v sphere VM groups params
func (o *ListProjectVSphereVMGroupsParams) WithDatacenterName(datacenterName *string) *ListProjectVSphereVMGroupsParams {
	o.SetDatacenterName(datacenterName)
	return o
}

// SetDatacenterName adds the datacenterName to the list project v sphere VM groups params
func (o *ListProjectVSphereVMGroupsParams) SetDatacenterName(datacenterName *string) {
	o.DatacenterName = datacenterName
}

// WithPassword adds the password to the list project v sphere VM groups params
func (o *ListProjectVSphereVMGroupsParams) WithPassword(password *string) *ListProjectVSphereVMGroupsParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the list project v sphere VM groups params
func (o *ListProjectVSphereVMGroupsParams) SetPassword(password *string) {
	o.Password = password
}

// WithUsername adds the username to the list project v sphere VM groups params
func (o *ListProjectVSphereVMGroupsParams) WithUsername(username *string) *ListProjectVSphereVMGroupsParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the list project v sphere VM groups params
func (o *ListProjectVSphereVMGroupsParams) SetUsername(username *string) {
	o.Username = username
}

// WithProjectID adds the projectID to the list project v sphere VM groups params
func (o *ListProjectVSphereVMGroupsParams) WithProjectID(projectID string) *ListProjectVSphereVMGroupsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list project v sphere VM groups params
func (o *ListProjectVSphereVMGroupsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectVSphereVMGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	if o.DatacenterName != nil {

		// header param DatacenterName
		if err := r.SetHeaderParam("DatacenterName", *o.DatacenterName); err != nil {
			return err
		}
	}

	if o.Password != nil {

		// header param Password
		if err := r.SetHeaderParam("Password", *o.Password); err != nil {
			return err
		}
	}

	if o.Username != nil {

		// header param Username
		if err := r.SetHeaderParam("Username", *o.Username); err != nil {
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
