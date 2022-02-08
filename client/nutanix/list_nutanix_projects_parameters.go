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

// NewListNutanixProjectsParams creates a new ListNutanixProjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListNutanixProjectsParams() *ListNutanixProjectsParams {
	return &ListNutanixProjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListNutanixProjectsParamsWithTimeout creates a new ListNutanixProjectsParams object
// with the ability to set a timeout on a request.
func NewListNutanixProjectsParamsWithTimeout(timeout time.Duration) *ListNutanixProjectsParams {
	return &ListNutanixProjectsParams{
		timeout: timeout,
	}
}

// NewListNutanixProjectsParamsWithContext creates a new ListNutanixProjectsParams object
// with the ability to set a context for a request.
func NewListNutanixProjectsParamsWithContext(ctx context.Context) *ListNutanixProjectsParams {
	return &ListNutanixProjectsParams{
		Context: ctx,
	}
}

// NewListNutanixProjectsParamsWithHTTPClient creates a new ListNutanixProjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListNutanixProjectsParamsWithHTTPClient(client *http.Client) *ListNutanixProjectsParams {
	return &ListNutanixProjectsParams{
		HTTPClient: client,
	}
}

/* ListNutanixProjectsParams contains all the parameters to send to the API endpoint
   for the list nutanix projects operation.

   Typically these are written to a http.Request.
*/
type ListNutanixProjectsParams struct {

	// Credential.
	Credential *string

	// NutanixPassword.
	NutanixPassword *string

	// NutanixUsername.
	NutanixUsername *string

	// ProxyURL.
	ProxyURL *string

	/* Dc.

	   KKP Datacenter to use for endpoint
	*/
	DC string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list nutanix projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNutanixProjectsParams) WithDefaults() *ListNutanixProjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list nutanix projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNutanixProjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list nutanix projects params
func (o *ListNutanixProjectsParams) WithTimeout(timeout time.Duration) *ListNutanixProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list nutanix projects params
func (o *ListNutanixProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list nutanix projects params
func (o *ListNutanixProjectsParams) WithContext(ctx context.Context) *ListNutanixProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list nutanix projects params
func (o *ListNutanixProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list nutanix projects params
func (o *ListNutanixProjectsParams) WithHTTPClient(client *http.Client) *ListNutanixProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list nutanix projects params
func (o *ListNutanixProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the list nutanix projects params
func (o *ListNutanixProjectsParams) WithCredential(credential *string) *ListNutanixProjectsParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list nutanix projects params
func (o *ListNutanixProjectsParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithNutanixPassword adds the nutanixPassword to the list nutanix projects params
func (o *ListNutanixProjectsParams) WithNutanixPassword(nutanixPassword *string) *ListNutanixProjectsParams {
	o.SetNutanixPassword(nutanixPassword)
	return o
}

// SetNutanixPassword adds the nutanixPassword to the list nutanix projects params
func (o *ListNutanixProjectsParams) SetNutanixPassword(nutanixPassword *string) {
	o.NutanixPassword = nutanixPassword
}

// WithNutanixUsername adds the nutanixUsername to the list nutanix projects params
func (o *ListNutanixProjectsParams) WithNutanixUsername(nutanixUsername *string) *ListNutanixProjectsParams {
	o.SetNutanixUsername(nutanixUsername)
	return o
}

// SetNutanixUsername adds the nutanixUsername to the list nutanix projects params
func (o *ListNutanixProjectsParams) SetNutanixUsername(nutanixUsername *string) {
	o.NutanixUsername = nutanixUsername
}

// WithProxyURL adds the proxyURL to the list nutanix projects params
func (o *ListNutanixProjectsParams) WithProxyURL(proxyURL *string) *ListNutanixProjectsParams {
	o.SetProxyURL(proxyURL)
	return o
}

// SetProxyURL adds the proxyUrl to the list nutanix projects params
func (o *ListNutanixProjectsParams) SetProxyURL(proxyURL *string) {
	o.ProxyURL = proxyURL
}

// WithDC adds the dc to the list nutanix projects params
func (o *ListNutanixProjectsParams) WithDC(dc string) *ListNutanixProjectsParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list nutanix projects params
func (o *ListNutanixProjectsParams) SetDC(dc string) {
	o.DC = dc
}

// WriteToRequest writes these params to a swagger request
func (o *ListNutanixProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.NutanixUsername != nil {

		// header param NutanixUsername
		if err := r.SetHeaderParam("NutanixUsername", *o.NutanixUsername); err != nil {
			return err
		}
	}

	if o.ProxyURL != nil {

		// header param ProxyURL
		if err := r.SetHeaderParam("ProxyURL", *o.ProxyURL); err != nil {
			return err
		}
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
