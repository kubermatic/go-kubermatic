// Code generated by go-swagger; DO NOT EDIT.

package nutanix

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

// NewListProjectNutanixCategoriesParams creates a new ListProjectNutanixCategoriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectNutanixCategoriesParams() *ListProjectNutanixCategoriesParams {
	return &ListProjectNutanixCategoriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectNutanixCategoriesParamsWithTimeout creates a new ListProjectNutanixCategoriesParams object
// with the ability to set a timeout on a request.
func NewListProjectNutanixCategoriesParamsWithTimeout(timeout time.Duration) *ListProjectNutanixCategoriesParams {
	return &ListProjectNutanixCategoriesParams{
		timeout: timeout,
	}
}

// NewListProjectNutanixCategoriesParamsWithContext creates a new ListProjectNutanixCategoriesParams object
// with the ability to set a context for a request.
func NewListProjectNutanixCategoriesParamsWithContext(ctx context.Context) *ListProjectNutanixCategoriesParams {
	return &ListProjectNutanixCategoriesParams{
		Context: ctx,
	}
}

// NewListProjectNutanixCategoriesParamsWithHTTPClient creates a new ListProjectNutanixCategoriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectNutanixCategoriesParamsWithHTTPClient(client *http.Client) *ListProjectNutanixCategoriesParams {
	return &ListProjectNutanixCategoriesParams{
		HTTPClient: client,
	}
}

/*
ListProjectNutanixCategoriesParams contains all the parameters to send to the API endpoint

	for the list project nutanix categories operation.

	Typically these are written to a http.Request.
*/
type ListProjectNutanixCategoriesParams struct {

	// Credential.
	Credential *string

	// NutanixPassword.
	NutanixPassword *string

	// NutanixProxyURL.
	NutanixProxyURL *string

	// NutanixUsername.
	NutanixUsername *string

	/* Dc.

	   KKP Datacenter to use for endpoint
	*/
	DC string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project nutanix categories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectNutanixCategoriesParams) WithDefaults() *ListProjectNutanixCategoriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project nutanix categories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectNutanixCategoriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project nutanix categories params
func (o *ListProjectNutanixCategoriesParams) WithTimeout(timeout time.Duration) *ListProjectNutanixCategoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project nutanix categories params
func (o *ListProjectNutanixCategoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project nutanix categories params
func (o *ListProjectNutanixCategoriesParams) WithContext(ctx context.Context) *ListProjectNutanixCategoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project nutanix categories params
func (o *ListProjectNutanixCategoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project nutanix categories params
func (o *ListProjectNutanixCategoriesParams) WithHTTPClient(client *http.Client) *ListProjectNutanixCategoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project nutanix categories params
func (o *ListProjectNutanixCategoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the list project nutanix categories params
func (o *ListProjectNutanixCategoriesParams) WithCredential(credential *string) *ListProjectNutanixCategoriesParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list project nutanix categories params
func (o *ListProjectNutanixCategoriesParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithNutanixPassword adds the nutanixPassword to the list project nutanix categories params
func (o *ListProjectNutanixCategoriesParams) WithNutanixPassword(nutanixPassword *string) *ListProjectNutanixCategoriesParams {
	o.SetNutanixPassword(nutanixPassword)
	return o
}

// SetNutanixPassword adds the nutanixPassword to the list project nutanix categories params
func (o *ListProjectNutanixCategoriesParams) SetNutanixPassword(nutanixPassword *string) {
	o.NutanixPassword = nutanixPassword
}

// WithNutanixProxyURL adds the nutanixProxyURL to the list project nutanix categories params
func (o *ListProjectNutanixCategoriesParams) WithNutanixProxyURL(nutanixProxyURL *string) *ListProjectNutanixCategoriesParams {
	o.SetNutanixProxyURL(nutanixProxyURL)
	return o
}

// SetNutanixProxyURL adds the nutanixProxyUrl to the list project nutanix categories params
func (o *ListProjectNutanixCategoriesParams) SetNutanixProxyURL(nutanixProxyURL *string) {
	o.NutanixProxyURL = nutanixProxyURL
}

// WithNutanixUsername adds the nutanixUsername to the list project nutanix categories params
func (o *ListProjectNutanixCategoriesParams) WithNutanixUsername(nutanixUsername *string) *ListProjectNutanixCategoriesParams {
	o.SetNutanixUsername(nutanixUsername)
	return o
}

// SetNutanixUsername adds the nutanixUsername to the list project nutanix categories params
func (o *ListProjectNutanixCategoriesParams) SetNutanixUsername(nutanixUsername *string) {
	o.NutanixUsername = nutanixUsername
}

// WithDC adds the dc to the list project nutanix categories params
func (o *ListProjectNutanixCategoriesParams) WithDC(dc string) *ListProjectNutanixCategoriesParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list project nutanix categories params
func (o *ListProjectNutanixCategoriesParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the list project nutanix categories params
func (o *ListProjectNutanixCategoriesParams) WithProjectID(projectID string) *ListProjectNutanixCategoriesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list project nutanix categories params
func (o *ListProjectNutanixCategoriesParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectNutanixCategoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.NutanixPassword != nil {

		// header param NutanixPassword
		if err := r.SetHeaderParam("NutanixPassword", *o.NutanixPassword); err != nil {
			return err
		}
	}

	if o.NutanixProxyURL != nil {

		// header param NutanixProxyURL
		if err := r.SetHeaderParam("NutanixProxyURL", *o.NutanixProxyURL); err != nil {
			return err
		}
	}

	if o.NutanixUsername != nil {

		// header param NutanixUsername
		if err := r.SetHeaderParam("NutanixUsername", *o.NutanixUsername); err != nil {
			return err
		}
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
