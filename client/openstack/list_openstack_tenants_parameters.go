// Code generated by go-swagger; DO NOT EDIT.

package openstack

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

// NewListOpenstackTenantsParams creates a new ListOpenstackTenantsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListOpenstackTenantsParams() *ListOpenstackTenantsParams {
	return &ListOpenstackTenantsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListOpenstackTenantsParamsWithTimeout creates a new ListOpenstackTenantsParams object
// with the ability to set a timeout on a request.
func NewListOpenstackTenantsParamsWithTimeout(timeout time.Duration) *ListOpenstackTenantsParams {
	return &ListOpenstackTenantsParams{
		timeout: timeout,
	}
}

// NewListOpenstackTenantsParamsWithContext creates a new ListOpenstackTenantsParams object
// with the ability to set a context for a request.
func NewListOpenstackTenantsParamsWithContext(ctx context.Context) *ListOpenstackTenantsParams {
	return &ListOpenstackTenantsParams{
		Context: ctx,
	}
}

// NewListOpenstackTenantsParamsWithHTTPClient creates a new ListOpenstackTenantsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListOpenstackTenantsParamsWithHTTPClient(client *http.Client) *ListOpenstackTenantsParams {
	return &ListOpenstackTenantsParams{
		HTTPClient: client,
	}
}

/*
ListOpenstackTenantsParams contains all the parameters to send to the API endpoint

	for the list openstack tenants operation.

	Typically these are written to a http.Request.
*/
type ListOpenstackTenantsParams struct {

	// ApplicationCredentialID.
	ApplicationCredentialID *string

	// ApplicationCredentialSecret.
	ApplicationCredentialSecret *string

	// Credential.
	Credential *string

	// DatacenterName.
	DatacenterName *string

	// Domain.
	Domain *string

	// OIDCAuthentication.
	OIDCAuthentication *bool

	// Password.
	Password *string

	// Username.
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list openstack tenants params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOpenstackTenantsParams) WithDefaults() *ListOpenstackTenantsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list openstack tenants params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOpenstackTenantsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list openstack tenants params
func (o *ListOpenstackTenantsParams) WithTimeout(timeout time.Duration) *ListOpenstackTenantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list openstack tenants params
func (o *ListOpenstackTenantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list openstack tenants params
func (o *ListOpenstackTenantsParams) WithContext(ctx context.Context) *ListOpenstackTenantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list openstack tenants params
func (o *ListOpenstackTenantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list openstack tenants params
func (o *ListOpenstackTenantsParams) WithHTTPClient(client *http.Client) *ListOpenstackTenantsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list openstack tenants params
func (o *ListOpenstackTenantsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationCredentialID adds the applicationCredentialID to the list openstack tenants params
func (o *ListOpenstackTenantsParams) WithApplicationCredentialID(applicationCredentialID *string) *ListOpenstackTenantsParams {
	o.SetApplicationCredentialID(applicationCredentialID)
	return o
}

// SetApplicationCredentialID adds the applicationCredentialId to the list openstack tenants params
func (o *ListOpenstackTenantsParams) SetApplicationCredentialID(applicationCredentialID *string) {
	o.ApplicationCredentialID = applicationCredentialID
}

// WithApplicationCredentialSecret adds the applicationCredentialSecret to the list openstack tenants params
func (o *ListOpenstackTenantsParams) WithApplicationCredentialSecret(applicationCredentialSecret *string) *ListOpenstackTenantsParams {
	o.SetApplicationCredentialSecret(applicationCredentialSecret)
	return o
}

// SetApplicationCredentialSecret adds the applicationCredentialSecret to the list openstack tenants params
func (o *ListOpenstackTenantsParams) SetApplicationCredentialSecret(applicationCredentialSecret *string) {
	o.ApplicationCredentialSecret = applicationCredentialSecret
}

// WithCredential adds the credential to the list openstack tenants params
func (o *ListOpenstackTenantsParams) WithCredential(credential *string) *ListOpenstackTenantsParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list openstack tenants params
func (o *ListOpenstackTenantsParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithDatacenterName adds the datacenterName to the list openstack tenants params
func (o *ListOpenstackTenantsParams) WithDatacenterName(datacenterName *string) *ListOpenstackTenantsParams {
	o.SetDatacenterName(datacenterName)
	return o
}

// SetDatacenterName adds the datacenterName to the list openstack tenants params
func (o *ListOpenstackTenantsParams) SetDatacenterName(datacenterName *string) {
	o.DatacenterName = datacenterName
}

// WithDomain adds the domain to the list openstack tenants params
func (o *ListOpenstackTenantsParams) WithDomain(domain *string) *ListOpenstackTenantsParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the list openstack tenants params
func (o *ListOpenstackTenantsParams) SetDomain(domain *string) {
	o.Domain = domain
}

// WithOIDCAuthentication adds the oIDCAuthentication to the list openstack tenants params
func (o *ListOpenstackTenantsParams) WithOIDCAuthentication(oIDCAuthentication *bool) *ListOpenstackTenantsParams {
	o.SetOIDCAuthentication(oIDCAuthentication)
	return o
}

// SetOIDCAuthentication adds the oIdCAuthentication to the list openstack tenants params
func (o *ListOpenstackTenantsParams) SetOIDCAuthentication(oIDCAuthentication *bool) {
	o.OIDCAuthentication = oIDCAuthentication
}

// WithPassword adds the password to the list openstack tenants params
func (o *ListOpenstackTenantsParams) WithPassword(password *string) *ListOpenstackTenantsParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the list openstack tenants params
func (o *ListOpenstackTenantsParams) SetPassword(password *string) {
	o.Password = password
}

// WithUsername adds the username to the list openstack tenants params
func (o *ListOpenstackTenantsParams) WithUsername(username *string) *ListOpenstackTenantsParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the list openstack tenants params
func (o *ListOpenstackTenantsParams) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *ListOpenstackTenantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicationCredentialID != nil {

		// header param ApplicationCredentialID
		if err := r.SetHeaderParam("ApplicationCredentialID", *o.ApplicationCredentialID); err != nil {
			return err
		}
	}

	if o.ApplicationCredentialSecret != nil {

		// header param ApplicationCredentialSecret
		if err := r.SetHeaderParam("ApplicationCredentialSecret", *o.ApplicationCredentialSecret); err != nil {
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

	if o.Domain != nil {

		// header param Domain
		if err := r.SetHeaderParam("Domain", *o.Domain); err != nil {
			return err
		}
	}

	if o.OIDCAuthentication != nil {

		// header param OIDCAuthentication
		if err := r.SetHeaderParam("OIDCAuthentication", swag.FormatBool(*o.OIDCAuthentication)); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
