// Code generated by go-swagger; DO NOT EDIT.

package alibaba

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

// NewListProjectAlibabaInstanceTypesParams creates a new ListProjectAlibabaInstanceTypesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectAlibabaInstanceTypesParams() *ListProjectAlibabaInstanceTypesParams {
	return &ListProjectAlibabaInstanceTypesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectAlibabaInstanceTypesParamsWithTimeout creates a new ListProjectAlibabaInstanceTypesParams object
// with the ability to set a timeout on a request.
func NewListProjectAlibabaInstanceTypesParamsWithTimeout(timeout time.Duration) *ListProjectAlibabaInstanceTypesParams {
	return &ListProjectAlibabaInstanceTypesParams{
		timeout: timeout,
	}
}

// NewListProjectAlibabaInstanceTypesParamsWithContext creates a new ListProjectAlibabaInstanceTypesParams object
// with the ability to set a context for a request.
func NewListProjectAlibabaInstanceTypesParamsWithContext(ctx context.Context) *ListProjectAlibabaInstanceTypesParams {
	return &ListProjectAlibabaInstanceTypesParams{
		Context: ctx,
	}
}

// NewListProjectAlibabaInstanceTypesParamsWithHTTPClient creates a new ListProjectAlibabaInstanceTypesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectAlibabaInstanceTypesParamsWithHTTPClient(client *http.Client) *ListProjectAlibabaInstanceTypesParams {
	return &ListProjectAlibabaInstanceTypesParams{
		HTTPClient: client,
	}
}

/*
ListProjectAlibabaInstanceTypesParams contains all the parameters to send to the API endpoint

	for the list project alibaba instance types operation.

	Typically these are written to a http.Request.
*/
type ListProjectAlibabaInstanceTypesParams struct {

	// AccessKeyID.
	AccessKeyID *string

	// AccessKeySecret.
	AccessKeySecret *string

	// Credential.
	Credential *string

	// DatacenterName.
	DatacenterName *string

	// Region.
	Region *string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project alibaba instance types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectAlibabaInstanceTypesParams) WithDefaults() *ListProjectAlibabaInstanceTypesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project alibaba instance types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectAlibabaInstanceTypesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project alibaba instance types params
func (o *ListProjectAlibabaInstanceTypesParams) WithTimeout(timeout time.Duration) *ListProjectAlibabaInstanceTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project alibaba instance types params
func (o *ListProjectAlibabaInstanceTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project alibaba instance types params
func (o *ListProjectAlibabaInstanceTypesParams) WithContext(ctx context.Context) *ListProjectAlibabaInstanceTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project alibaba instance types params
func (o *ListProjectAlibabaInstanceTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project alibaba instance types params
func (o *ListProjectAlibabaInstanceTypesParams) WithHTTPClient(client *http.Client) *ListProjectAlibabaInstanceTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project alibaba instance types params
func (o *ListProjectAlibabaInstanceTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessKeyID adds the accessKeyID to the list project alibaba instance types params
func (o *ListProjectAlibabaInstanceTypesParams) WithAccessKeyID(accessKeyID *string) *ListProjectAlibabaInstanceTypesParams {
	o.SetAccessKeyID(accessKeyID)
	return o
}

// SetAccessKeyID adds the accessKeyId to the list project alibaba instance types params
func (o *ListProjectAlibabaInstanceTypesParams) SetAccessKeyID(accessKeyID *string) {
	o.AccessKeyID = accessKeyID
}

// WithAccessKeySecret adds the accessKeySecret to the list project alibaba instance types params
func (o *ListProjectAlibabaInstanceTypesParams) WithAccessKeySecret(accessKeySecret *string) *ListProjectAlibabaInstanceTypesParams {
	o.SetAccessKeySecret(accessKeySecret)
	return o
}

// SetAccessKeySecret adds the accessKeySecret to the list project alibaba instance types params
func (o *ListProjectAlibabaInstanceTypesParams) SetAccessKeySecret(accessKeySecret *string) {
	o.AccessKeySecret = accessKeySecret
}

// WithCredential adds the credential to the list project alibaba instance types params
func (o *ListProjectAlibabaInstanceTypesParams) WithCredential(credential *string) *ListProjectAlibabaInstanceTypesParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list project alibaba instance types params
func (o *ListProjectAlibabaInstanceTypesParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithDatacenterName adds the datacenterName to the list project alibaba instance types params
func (o *ListProjectAlibabaInstanceTypesParams) WithDatacenterName(datacenterName *string) *ListProjectAlibabaInstanceTypesParams {
	o.SetDatacenterName(datacenterName)
	return o
}

// SetDatacenterName adds the datacenterName to the list project alibaba instance types params
func (o *ListProjectAlibabaInstanceTypesParams) SetDatacenterName(datacenterName *string) {
	o.DatacenterName = datacenterName
}

// WithRegion adds the region to the list project alibaba instance types params
func (o *ListProjectAlibabaInstanceTypesParams) WithRegion(region *string) *ListProjectAlibabaInstanceTypesParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the list project alibaba instance types params
func (o *ListProjectAlibabaInstanceTypesParams) SetRegion(region *string) {
	o.Region = region
}

// WithProjectID adds the projectID to the list project alibaba instance types params
func (o *ListProjectAlibabaInstanceTypesParams) WithProjectID(projectID string) *ListProjectAlibabaInstanceTypesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list project alibaba instance types params
func (o *ListProjectAlibabaInstanceTypesParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectAlibabaInstanceTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessKeyID != nil {

		// header param AccessKeyID
		if err := r.SetHeaderParam("AccessKeyID", *o.AccessKeyID); err != nil {
			return err
		}
	}

	if o.AccessKeySecret != nil {

		// header param AccessKeySecret
		if err := r.SetHeaderParam("AccessKeySecret", *o.AccessKeySecret); err != nil {
			return err
		}
	}

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

	if o.Region != nil {

		// header param Region
		if err := r.SetHeaderParam("Region", *o.Region); err != nil {
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
