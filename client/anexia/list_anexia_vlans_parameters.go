// Code generated by go-swagger; DO NOT EDIT.

package anexia

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

// NewListAnexiaVlansParams creates a new ListAnexiaVlansParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAnexiaVlansParams() *ListAnexiaVlansParams {
	return &ListAnexiaVlansParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAnexiaVlansParamsWithTimeout creates a new ListAnexiaVlansParams object
// with the ability to set a timeout on a request.
func NewListAnexiaVlansParamsWithTimeout(timeout time.Duration) *ListAnexiaVlansParams {
	return &ListAnexiaVlansParams{
		timeout: timeout,
	}
}

// NewListAnexiaVlansParamsWithContext creates a new ListAnexiaVlansParams object
// with the ability to set a context for a request.
func NewListAnexiaVlansParamsWithContext(ctx context.Context) *ListAnexiaVlansParams {
	return &ListAnexiaVlansParams{
		Context: ctx,
	}
}

// NewListAnexiaVlansParamsWithHTTPClient creates a new ListAnexiaVlansParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAnexiaVlansParamsWithHTTPClient(client *http.Client) *ListAnexiaVlansParams {
	return &ListAnexiaVlansParams{
		HTTPClient: client,
	}
}

/*
ListAnexiaVlansParams contains all the parameters to send to the API endpoint

	for the list anexia vlans operation.

	Typically these are written to a http.Request.
*/
type ListAnexiaVlansParams struct {

	// Credential.
	Credential *string

	// Token.
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list anexia vlans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAnexiaVlansParams) WithDefaults() *ListAnexiaVlansParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list anexia vlans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAnexiaVlansParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list anexia vlans params
func (o *ListAnexiaVlansParams) WithTimeout(timeout time.Duration) *ListAnexiaVlansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list anexia vlans params
func (o *ListAnexiaVlansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list anexia vlans params
func (o *ListAnexiaVlansParams) WithContext(ctx context.Context) *ListAnexiaVlansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list anexia vlans params
func (o *ListAnexiaVlansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list anexia vlans params
func (o *ListAnexiaVlansParams) WithHTTPClient(client *http.Client) *ListAnexiaVlansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list anexia vlans params
func (o *ListAnexiaVlansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the list anexia vlans params
func (o *ListAnexiaVlansParams) WithCredential(credential *string) *ListAnexiaVlansParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list anexia vlans params
func (o *ListAnexiaVlansParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithToken adds the token to the list anexia vlans params
func (o *ListAnexiaVlansParams) WithToken(token *string) *ListAnexiaVlansParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the list anexia vlans params
func (o *ListAnexiaVlansParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *ListAnexiaVlansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Token != nil {

		// header param Token
		if err := r.SetHeaderParam("Token", *o.Token); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
