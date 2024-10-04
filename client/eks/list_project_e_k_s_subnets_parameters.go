// Code generated by go-swagger; DO NOT EDIT.

package eks

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

// NewListProjectEKSSubnetsParams creates a new ListProjectEKSSubnetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectEKSSubnetsParams() *ListProjectEKSSubnetsParams {
	return &ListProjectEKSSubnetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectEKSSubnetsParamsWithTimeout creates a new ListProjectEKSSubnetsParams object
// with the ability to set a timeout on a request.
func NewListProjectEKSSubnetsParamsWithTimeout(timeout time.Duration) *ListProjectEKSSubnetsParams {
	return &ListProjectEKSSubnetsParams{
		timeout: timeout,
	}
}

// NewListProjectEKSSubnetsParamsWithContext creates a new ListProjectEKSSubnetsParams object
// with the ability to set a context for a request.
func NewListProjectEKSSubnetsParamsWithContext(ctx context.Context) *ListProjectEKSSubnetsParams {
	return &ListProjectEKSSubnetsParams{
		Context: ctx,
	}
}

// NewListProjectEKSSubnetsParamsWithHTTPClient creates a new ListProjectEKSSubnetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectEKSSubnetsParamsWithHTTPClient(client *http.Client) *ListProjectEKSSubnetsParams {
	return &ListProjectEKSSubnetsParams{
		HTTPClient: client,
	}
}

/*
ListProjectEKSSubnetsParams contains all the parameters to send to the API endpoint

	for the list project e k s subnets operation.

	Typically these are written to a http.Request.
*/
type ListProjectEKSSubnetsParams struct {

	// AccessKeyID.
	AccessKeyID *string

	// AssumeRoleARN.
	AssumeRoleARN *string

	// AssumeRoleExternalID.
	AssumeRoleExternalID *string

	// Credential.
	Credential *string

	// Region.
	Region *string

	// SecretAccessKey.
	SecretAccessKey *string

	// VpcID.
	VpcID *string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project e k s subnets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectEKSSubnetsParams) WithDefaults() *ListProjectEKSSubnetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project e k s subnets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectEKSSubnetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) WithTimeout(timeout time.Duration) *ListProjectEKSSubnetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) WithContext(ctx context.Context) *ListProjectEKSSubnetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) WithHTTPClient(client *http.Client) *ListProjectEKSSubnetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessKeyID adds the accessKeyID to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) WithAccessKeyID(accessKeyID *string) *ListProjectEKSSubnetsParams {
	o.SetAccessKeyID(accessKeyID)
	return o
}

// SetAccessKeyID adds the accessKeyId to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) SetAccessKeyID(accessKeyID *string) {
	o.AccessKeyID = accessKeyID
}

// WithAssumeRoleARN adds the assumeRoleARN to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) WithAssumeRoleARN(assumeRoleARN *string) *ListProjectEKSSubnetsParams {
	o.SetAssumeRoleARN(assumeRoleARN)
	return o
}

// SetAssumeRoleARN adds the assumeRoleARN to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) SetAssumeRoleARN(assumeRoleARN *string) {
	o.AssumeRoleARN = assumeRoleARN
}

// WithAssumeRoleExternalID adds the assumeRoleExternalID to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) WithAssumeRoleExternalID(assumeRoleExternalID *string) *ListProjectEKSSubnetsParams {
	o.SetAssumeRoleExternalID(assumeRoleExternalID)
	return o
}

// SetAssumeRoleExternalID adds the assumeRoleExternalId to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) SetAssumeRoleExternalID(assumeRoleExternalID *string) {
	o.AssumeRoleExternalID = assumeRoleExternalID
}

// WithCredential adds the credential to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) WithCredential(credential *string) *ListProjectEKSSubnetsParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithRegion adds the region to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) WithRegion(region *string) *ListProjectEKSSubnetsParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) SetRegion(region *string) {
	o.Region = region
}

// WithSecretAccessKey adds the secretAccessKey to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) WithSecretAccessKey(secretAccessKey *string) *ListProjectEKSSubnetsParams {
	o.SetSecretAccessKey(secretAccessKey)
	return o
}

// SetSecretAccessKey adds the secretAccessKey to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) SetSecretAccessKey(secretAccessKey *string) {
	o.SecretAccessKey = secretAccessKey
}

// WithVpcID adds the vpcID to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) WithVpcID(vpcID *string) *ListProjectEKSSubnetsParams {
	o.SetVpcID(vpcID)
	return o
}

// SetVpcID adds the vpcId to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) SetVpcID(vpcID *string) {
	o.VpcID = vpcID
}

// WithProjectID adds the projectID to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) WithProjectID(projectID string) *ListProjectEKSSubnetsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list project e k s subnets params
func (o *ListProjectEKSSubnetsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectEKSSubnetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.AssumeRoleARN != nil {

		// header param AssumeRoleARN
		if err := r.SetHeaderParam("AssumeRoleARN", *o.AssumeRoleARN); err != nil {
			return err
		}
	}

	if o.AssumeRoleExternalID != nil {

		// header param AssumeRoleExternalID
		if err := r.SetHeaderParam("AssumeRoleExternalID", *o.AssumeRoleExternalID); err != nil {
			return err
		}
	}

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	if o.Region != nil {

		// header param Region
		if err := r.SetHeaderParam("Region", *o.Region); err != nil {
			return err
		}
	}

	if o.SecretAccessKey != nil {

		// header param SecretAccessKey
		if err := r.SetHeaderParam("SecretAccessKey", *o.SecretAccessKey); err != nil {
			return err
		}
	}

	if o.VpcID != nil {

		// header param VpcId
		if err := r.SetHeaderParam("VpcId", *o.VpcID); err != nil {
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
