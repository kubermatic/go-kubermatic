// Code generated by go-swagger; DO NOT EDIT.

package kubevirt

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

// NewListProjectKubevirtVPCsParams creates a new ListProjectKubevirtVPCsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectKubevirtVPCsParams() *ListProjectKubevirtVPCsParams {
	return &ListProjectKubevirtVPCsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectKubevirtVPCsParamsWithTimeout creates a new ListProjectKubevirtVPCsParams object
// with the ability to set a timeout on a request.
func NewListProjectKubevirtVPCsParamsWithTimeout(timeout time.Duration) *ListProjectKubevirtVPCsParams {
	return &ListProjectKubevirtVPCsParams{
		timeout: timeout,
	}
}

// NewListProjectKubevirtVPCsParamsWithContext creates a new ListProjectKubevirtVPCsParams object
// with the ability to set a context for a request.
func NewListProjectKubevirtVPCsParamsWithContext(ctx context.Context) *ListProjectKubevirtVPCsParams {
	return &ListProjectKubevirtVPCsParams{
		Context: ctx,
	}
}

// NewListProjectKubevirtVPCsParamsWithHTTPClient creates a new ListProjectKubevirtVPCsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectKubevirtVPCsParamsWithHTTPClient(client *http.Client) *ListProjectKubevirtVPCsParams {
	return &ListProjectKubevirtVPCsParams{
		HTTPClient: client,
	}
}

/*
ListProjectKubevirtVPCsParams contains all the parameters to send to the API endpoint

	for the list project kubevirt v p cs operation.

	Typically these are written to a http.Request.
*/
type ListProjectKubevirtVPCsParams struct {

	// Credential.
	Credential *string

	// DatacenterName.
	DatacenterName *string

	// Kubeconfig.
	Kubeconfig *string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project kubevirt v p cs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectKubevirtVPCsParams) WithDefaults() *ListProjectKubevirtVPCsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project kubevirt v p cs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectKubevirtVPCsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project kubevirt v p cs params
func (o *ListProjectKubevirtVPCsParams) WithTimeout(timeout time.Duration) *ListProjectKubevirtVPCsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project kubevirt v p cs params
func (o *ListProjectKubevirtVPCsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project kubevirt v p cs params
func (o *ListProjectKubevirtVPCsParams) WithContext(ctx context.Context) *ListProjectKubevirtVPCsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project kubevirt v p cs params
func (o *ListProjectKubevirtVPCsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project kubevirt v p cs params
func (o *ListProjectKubevirtVPCsParams) WithHTTPClient(client *http.Client) *ListProjectKubevirtVPCsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project kubevirt v p cs params
func (o *ListProjectKubevirtVPCsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the list project kubevirt v p cs params
func (o *ListProjectKubevirtVPCsParams) WithCredential(credential *string) *ListProjectKubevirtVPCsParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list project kubevirt v p cs params
func (o *ListProjectKubevirtVPCsParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithDatacenterName adds the datacenterName to the list project kubevirt v p cs params
func (o *ListProjectKubevirtVPCsParams) WithDatacenterName(datacenterName *string) *ListProjectKubevirtVPCsParams {
	o.SetDatacenterName(datacenterName)
	return o
}

// SetDatacenterName adds the datacenterName to the list project kubevirt v p cs params
func (o *ListProjectKubevirtVPCsParams) SetDatacenterName(datacenterName *string) {
	o.DatacenterName = datacenterName
}

// WithKubeconfig adds the kubeconfig to the list project kubevirt v p cs params
func (o *ListProjectKubevirtVPCsParams) WithKubeconfig(kubeconfig *string) *ListProjectKubevirtVPCsParams {
	o.SetKubeconfig(kubeconfig)
	return o
}

// SetKubeconfig adds the kubeconfig to the list project kubevirt v p cs params
func (o *ListProjectKubevirtVPCsParams) SetKubeconfig(kubeconfig *string) {
	o.Kubeconfig = kubeconfig
}

// WithProjectID adds the projectID to the list project kubevirt v p cs params
func (o *ListProjectKubevirtVPCsParams) WithProjectID(projectID string) *ListProjectKubevirtVPCsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list project kubevirt v p cs params
func (o *ListProjectKubevirtVPCsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectKubevirtVPCsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Kubeconfig != nil {

		// header param Kubeconfig
		if err := r.SetHeaderParam("Kubeconfig", *o.Kubeconfig); err != nil {
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