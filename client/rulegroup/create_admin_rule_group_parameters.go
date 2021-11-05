// Code generated by go-swagger; DO NOT EDIT.

package rulegroup

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

	"github.com/kubermatic/go-kubermatic/models"
)

// NewCreateAdminRuleGroupParams creates a new CreateAdminRuleGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAdminRuleGroupParams() *CreateAdminRuleGroupParams {
	return &CreateAdminRuleGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAdminRuleGroupParamsWithTimeout creates a new CreateAdminRuleGroupParams object
// with the ability to set a timeout on a request.
func NewCreateAdminRuleGroupParamsWithTimeout(timeout time.Duration) *CreateAdminRuleGroupParams {
	return &CreateAdminRuleGroupParams{
		timeout: timeout,
	}
}

// NewCreateAdminRuleGroupParamsWithContext creates a new CreateAdminRuleGroupParams object
// with the ability to set a context for a request.
func NewCreateAdminRuleGroupParamsWithContext(ctx context.Context) *CreateAdminRuleGroupParams {
	return &CreateAdminRuleGroupParams{
		Context: ctx,
	}
}

// NewCreateAdminRuleGroupParamsWithHTTPClient creates a new CreateAdminRuleGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAdminRuleGroupParamsWithHTTPClient(client *http.Client) *CreateAdminRuleGroupParams {
	return &CreateAdminRuleGroupParams{
		HTTPClient: client,
	}
}

/* CreateAdminRuleGroupParams contains all the parameters to send to the API endpoint
   for the create admin rule group operation.

   Typically these are written to a http.Request.
*/
type CreateAdminRuleGroupParams struct {

	// Body.
	Body *models.RuleGroup

	// SeedName.
	SeedName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create admin rule group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAdminRuleGroupParams) WithDefaults() *CreateAdminRuleGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create admin rule group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAdminRuleGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create admin rule group params
func (o *CreateAdminRuleGroupParams) WithTimeout(timeout time.Duration) *CreateAdminRuleGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create admin rule group params
func (o *CreateAdminRuleGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create admin rule group params
func (o *CreateAdminRuleGroupParams) WithContext(ctx context.Context) *CreateAdminRuleGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create admin rule group params
func (o *CreateAdminRuleGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create admin rule group params
func (o *CreateAdminRuleGroupParams) WithHTTPClient(client *http.Client) *CreateAdminRuleGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create admin rule group params
func (o *CreateAdminRuleGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create admin rule group params
func (o *CreateAdminRuleGroupParams) WithBody(body *models.RuleGroup) *CreateAdminRuleGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create admin rule group params
func (o *CreateAdminRuleGroupParams) SetBody(body *models.RuleGroup) {
	o.Body = body
}

// WithSeedName adds the seedName to the create admin rule group params
func (o *CreateAdminRuleGroupParams) WithSeedName(seedName string) *CreateAdminRuleGroupParams {
	o.SetSeedName(seedName)
	return o
}

// SetSeedName adds the seedName to the create admin rule group params
func (o *CreateAdminRuleGroupParams) SetSeedName(seedName string) {
	o.SeedName = seedName
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAdminRuleGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param seed_name
	if err := r.SetPathParam("seed_name", o.SeedName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
