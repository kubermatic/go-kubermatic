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
	"github.com/go-openapi/swag"
)

// NewPatchClusterV2Params creates a new PatchClusterV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchClusterV2Params() *PatchClusterV2Params {
	return &PatchClusterV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchClusterV2ParamsWithTimeout creates a new PatchClusterV2Params object
// with the ability to set a timeout on a request.
func NewPatchClusterV2ParamsWithTimeout(timeout time.Duration) *PatchClusterV2Params {
	return &PatchClusterV2Params{
		timeout: timeout,
	}
}

// NewPatchClusterV2ParamsWithContext creates a new PatchClusterV2Params object
// with the ability to set a context for a request.
func NewPatchClusterV2ParamsWithContext(ctx context.Context) *PatchClusterV2Params {
	return &PatchClusterV2Params{
		Context: ctx,
	}
}

// NewPatchClusterV2ParamsWithHTTPClient creates a new PatchClusterV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewPatchClusterV2ParamsWithHTTPClient(client *http.Client) *PatchClusterV2Params {
	return &PatchClusterV2Params{
		HTTPClient: client,
	}
}

/*
PatchClusterV2Params contains all the parameters to send to the API endpoint

	for the patch cluster v2 operation.

	Typically these are written to a http.Request.
*/
type PatchClusterV2Params struct {

	// Patch.
	Patch interface{}

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	// SkipKubeletVersionValidation.
	SkipKubeletVersionValidation *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch cluster v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchClusterV2Params) WithDefaults() *PatchClusterV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch cluster v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchClusterV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch cluster v2 params
func (o *PatchClusterV2Params) WithTimeout(timeout time.Duration) *PatchClusterV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch cluster v2 params
func (o *PatchClusterV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch cluster v2 params
func (o *PatchClusterV2Params) WithContext(ctx context.Context) *PatchClusterV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch cluster v2 params
func (o *PatchClusterV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch cluster v2 params
func (o *PatchClusterV2Params) WithHTTPClient(client *http.Client) *PatchClusterV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch cluster v2 params
func (o *PatchClusterV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatch adds the patch to the patch cluster v2 params
func (o *PatchClusterV2Params) WithPatch(patch interface{}) *PatchClusterV2Params {
	o.SetPatch(patch)
	return o
}

// SetPatch adds the patch to the patch cluster v2 params
func (o *PatchClusterV2Params) SetPatch(patch interface{}) {
	o.Patch = patch
}

// WithClusterID adds the clusterID to the patch cluster v2 params
func (o *PatchClusterV2Params) WithClusterID(clusterID string) *PatchClusterV2Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the patch cluster v2 params
func (o *PatchClusterV2Params) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the patch cluster v2 params
func (o *PatchClusterV2Params) WithProjectID(projectID string) *PatchClusterV2Params {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the patch cluster v2 params
func (o *PatchClusterV2Params) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithSkipKubeletVersionValidation adds the skipKubeletVersionValidation to the patch cluster v2 params
func (o *PatchClusterV2Params) WithSkipKubeletVersionValidation(skipKubeletVersionValidation *bool) *PatchClusterV2Params {
	o.SetSkipKubeletVersionValidation(skipKubeletVersionValidation)
	return o
}

// SetSkipKubeletVersionValidation adds the skipKubeletVersionValidation to the patch cluster v2 params
func (o *PatchClusterV2Params) SetSkipKubeletVersionValidation(skipKubeletVersionValidation *bool) {
	o.SkipKubeletVersionValidation = skipKubeletVersionValidation
}

// WriteToRequest writes these params to a swagger request
func (o *PatchClusterV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Patch != nil {
		if err := r.SetBodyParam(o.Patch); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if o.SkipKubeletVersionValidation != nil {

		// query param skip_kubelet_version_validation
		var qrSkipKubeletVersionValidation bool

		if o.SkipKubeletVersionValidation != nil {
			qrSkipKubeletVersionValidation = *o.SkipKubeletVersionValidation
		}
		qSkipKubeletVersionValidation := swag.FormatBool(qrSkipKubeletVersionValidation)
		if qSkipKubeletVersionValidation != "" {

			if err := r.SetQueryParam("skip_kubelet_version_validation", qSkipKubeletVersionValidation); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
