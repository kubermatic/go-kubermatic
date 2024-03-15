// Code generated by go-swagger; DO NOT EDIT.

package resource_quota

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

// NewListResourceQuotasParams creates a new ListResourceQuotasParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListResourceQuotasParams() *ListResourceQuotasParams {
	return &ListResourceQuotasParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListResourceQuotasParamsWithTimeout creates a new ListResourceQuotasParams object
// with the ability to set a timeout on a request.
func NewListResourceQuotasParamsWithTimeout(timeout time.Duration) *ListResourceQuotasParams {
	return &ListResourceQuotasParams{
		timeout: timeout,
	}
}

// NewListResourceQuotasParamsWithContext creates a new ListResourceQuotasParams object
// with the ability to set a context for a request.
func NewListResourceQuotasParamsWithContext(ctx context.Context) *ListResourceQuotasParams {
	return &ListResourceQuotasParams{
		Context: ctx,
	}
}

// NewListResourceQuotasParamsWithHTTPClient creates a new ListResourceQuotasParams object
// with the ability to set a custom HTTPClient for a request.
func NewListResourceQuotasParamsWithHTTPClient(client *http.Client) *ListResourceQuotasParams {
	return &ListResourceQuotasParams{
		HTTPClient: client,
	}
}

/*
ListResourceQuotasParams contains all the parameters to send to the API endpoint

	for the list resource quotas operation.

	Typically these are written to a http.Request.
*/
type ListResourceQuotasParams struct {

	// Accumulate.
	Accumulate *bool

	// SubjectKind.
	SubjectKind *string

	// SubjectName.
	SubjectName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list resource quotas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListResourceQuotasParams) WithDefaults() *ListResourceQuotasParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list resource quotas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListResourceQuotasParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list resource quotas params
func (o *ListResourceQuotasParams) WithTimeout(timeout time.Duration) *ListResourceQuotasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list resource quotas params
func (o *ListResourceQuotasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list resource quotas params
func (o *ListResourceQuotasParams) WithContext(ctx context.Context) *ListResourceQuotasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list resource quotas params
func (o *ListResourceQuotasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list resource quotas params
func (o *ListResourceQuotasParams) WithHTTPClient(client *http.Client) *ListResourceQuotasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list resource quotas params
func (o *ListResourceQuotasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccumulate adds the accumulate to the list resource quotas params
func (o *ListResourceQuotasParams) WithAccumulate(accumulate *bool) *ListResourceQuotasParams {
	o.SetAccumulate(accumulate)
	return o
}

// SetAccumulate adds the accumulate to the list resource quotas params
func (o *ListResourceQuotasParams) SetAccumulate(accumulate *bool) {
	o.Accumulate = accumulate
}

// WithSubjectKind adds the subjectKind to the list resource quotas params
func (o *ListResourceQuotasParams) WithSubjectKind(subjectKind *string) *ListResourceQuotasParams {
	o.SetSubjectKind(subjectKind)
	return o
}

// SetSubjectKind adds the subjectKind to the list resource quotas params
func (o *ListResourceQuotasParams) SetSubjectKind(subjectKind *string) {
	o.SubjectKind = subjectKind
}

// WithSubjectName adds the subjectName to the list resource quotas params
func (o *ListResourceQuotasParams) WithSubjectName(subjectName *string) *ListResourceQuotasParams {
	o.SetSubjectName(subjectName)
	return o
}

// SetSubjectName adds the subjectName to the list resource quotas params
func (o *ListResourceQuotasParams) SetSubjectName(subjectName *string) {
	o.SubjectName = subjectName
}

// WriteToRequest writes these params to a swagger request
func (o *ListResourceQuotasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Accumulate != nil {

		// query param accumulate
		var qrAccumulate bool

		if o.Accumulate != nil {
			qrAccumulate = *o.Accumulate
		}
		qAccumulate := swag.FormatBool(qrAccumulate)
		if qAccumulate != "" {

			if err := r.SetQueryParam("accumulate", qAccumulate); err != nil {
				return err
			}
		}
	}

	if o.SubjectKind != nil {

		// query param subject_kind
		var qrSubjectKind string

		if o.SubjectKind != nil {
			qrSubjectKind = *o.SubjectKind
		}
		qSubjectKind := qrSubjectKind
		if qSubjectKind != "" {

			if err := r.SetQueryParam("subject_kind", qSubjectKind); err != nil {
				return err
			}
		}
	}

	if o.SubjectName != nil {

		// query param subject_name
		var qrSubjectName string

		if o.SubjectName != nil {
			qrSubjectName = *o.SubjectName
		}
		qSubjectName := qrSubjectName
		if qSubjectName != "" {

			if err := r.SetQueryParam("subject_name", qSubjectName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
