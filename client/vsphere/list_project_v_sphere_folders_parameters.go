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

// NewListProjectVSphereFoldersParams creates a new ListProjectVSphereFoldersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectVSphereFoldersParams() *ListProjectVSphereFoldersParams {
	return &ListProjectVSphereFoldersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectVSphereFoldersParamsWithTimeout creates a new ListProjectVSphereFoldersParams object
// with the ability to set a timeout on a request.
func NewListProjectVSphereFoldersParamsWithTimeout(timeout time.Duration) *ListProjectVSphereFoldersParams {
	return &ListProjectVSphereFoldersParams{
		timeout: timeout,
	}
}

// NewListProjectVSphereFoldersParamsWithContext creates a new ListProjectVSphereFoldersParams object
// with the ability to set a context for a request.
func NewListProjectVSphereFoldersParamsWithContext(ctx context.Context) *ListProjectVSphereFoldersParams {
	return &ListProjectVSphereFoldersParams{
		Context: ctx,
	}
}

// NewListProjectVSphereFoldersParamsWithHTTPClient creates a new ListProjectVSphereFoldersParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectVSphereFoldersParamsWithHTTPClient(client *http.Client) *ListProjectVSphereFoldersParams {
	return &ListProjectVSphereFoldersParams{
		HTTPClient: client,
	}
}

/*
ListProjectVSphereFoldersParams contains all the parameters to send to the API endpoint

	for the list project v sphere folders operation.

	Typically these are written to a http.Request.
*/
type ListProjectVSphereFoldersParams struct {

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

// WithDefaults hydrates default values in the list project v sphere folders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectVSphereFoldersParams) WithDefaults() *ListProjectVSphereFoldersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project v sphere folders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectVSphereFoldersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project v sphere folders params
func (o *ListProjectVSphereFoldersParams) WithTimeout(timeout time.Duration) *ListProjectVSphereFoldersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project v sphere folders params
func (o *ListProjectVSphereFoldersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project v sphere folders params
func (o *ListProjectVSphereFoldersParams) WithContext(ctx context.Context) *ListProjectVSphereFoldersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project v sphere folders params
func (o *ListProjectVSphereFoldersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project v sphere folders params
func (o *ListProjectVSphereFoldersParams) WithHTTPClient(client *http.Client) *ListProjectVSphereFoldersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project v sphere folders params
func (o *ListProjectVSphereFoldersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the list project v sphere folders params
func (o *ListProjectVSphereFoldersParams) WithCredential(credential *string) *ListProjectVSphereFoldersParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list project v sphere folders params
func (o *ListProjectVSphereFoldersParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithDatacenterName adds the datacenterName to the list project v sphere folders params
func (o *ListProjectVSphereFoldersParams) WithDatacenterName(datacenterName *string) *ListProjectVSphereFoldersParams {
	o.SetDatacenterName(datacenterName)
	return o
}

// SetDatacenterName adds the datacenterName to the list project v sphere folders params
func (o *ListProjectVSphereFoldersParams) SetDatacenterName(datacenterName *string) {
	o.DatacenterName = datacenterName
}

// WithPassword adds the password to the list project v sphere folders params
func (o *ListProjectVSphereFoldersParams) WithPassword(password *string) *ListProjectVSphereFoldersParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the list project v sphere folders params
func (o *ListProjectVSphereFoldersParams) SetPassword(password *string) {
	o.Password = password
}

// WithUsername adds the username to the list project v sphere folders params
func (o *ListProjectVSphereFoldersParams) WithUsername(username *string) *ListProjectVSphereFoldersParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the list project v sphere folders params
func (o *ListProjectVSphereFoldersParams) SetUsername(username *string) {
	o.Username = username
}

// WithProjectID adds the projectID to the list project v sphere folders params
func (o *ListProjectVSphereFoldersParams) WithProjectID(projectID string) *ListProjectVSphereFoldersParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list project v sphere folders params
func (o *ListProjectVSphereFoldersParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectVSphereFoldersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
