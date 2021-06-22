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

// NewListOpenstackSecurityGroupsParams creates a new ListOpenstackSecurityGroupsParams object
// with the default values initialized.
func NewListOpenstackSecurityGroupsParams() *ListOpenstackSecurityGroupsParams {
	var ()
	return &ListOpenstackSecurityGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListOpenstackSecurityGroupsParamsWithTimeout creates a new ListOpenstackSecurityGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListOpenstackSecurityGroupsParamsWithTimeout(timeout time.Duration) *ListOpenstackSecurityGroupsParams {
	var ()
	return &ListOpenstackSecurityGroupsParams{

		timeout: timeout,
	}
}

// NewListOpenstackSecurityGroupsParamsWithContext creates a new ListOpenstackSecurityGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListOpenstackSecurityGroupsParamsWithContext(ctx context.Context) *ListOpenstackSecurityGroupsParams {
	var ()
	return &ListOpenstackSecurityGroupsParams{

		Context: ctx,
	}
}

// NewListOpenstackSecurityGroupsParamsWithHTTPClient creates a new ListOpenstackSecurityGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListOpenstackSecurityGroupsParamsWithHTTPClient(client *http.Client) *ListOpenstackSecurityGroupsParams {
	var ()
	return &ListOpenstackSecurityGroupsParams{
		HTTPClient: client,
	}
}

/*ListOpenstackSecurityGroupsParams contains all the parameters to send to the API endpoint
for the list openstack security groups operation typically these are written to a http.Request
*/
type ListOpenstackSecurityGroupsParams struct {

	/*ApplicationCredentialID*/
	ApplicationCredentialID *string
	/*ApplicationCredentialSecret*/
	ApplicationCredentialSecret *string
	/*Credential*/
	Credential *string
	/*DatacenterName*/
	DatacenterName *string
	/*Domain*/
	Domain *string
	/*OIDCAuthentication*/
	OIDCAuthentication *bool
	/*Password*/
	Password *string
	/*Tenant*/
	Tenant *string
	/*TenantID*/
	TenantID *string
	/*Username*/
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) WithTimeout(timeout time.Duration) *ListOpenstackSecurityGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) WithContext(ctx context.Context) *ListOpenstackSecurityGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) WithHTTPClient(client *http.Client) *ListOpenstackSecurityGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationCredentialID adds the applicationCredentialID to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) WithApplicationCredentialID(applicationCredentialID *string) *ListOpenstackSecurityGroupsParams {
	o.SetApplicationCredentialID(applicationCredentialID)
	return o
}

// SetApplicationCredentialID adds the applicationCredentialId to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) SetApplicationCredentialID(applicationCredentialID *string) {
	o.ApplicationCredentialID = applicationCredentialID
}

// WithApplicationCredentialSecret adds the applicationCredentialSecret to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) WithApplicationCredentialSecret(applicationCredentialSecret *string) *ListOpenstackSecurityGroupsParams {
	o.SetApplicationCredentialSecret(applicationCredentialSecret)
	return o
}

// SetApplicationCredentialSecret adds the applicationCredentialSecret to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) SetApplicationCredentialSecret(applicationCredentialSecret *string) {
	o.ApplicationCredentialSecret = applicationCredentialSecret
}

// WithCredential adds the credential to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) WithCredential(credential *string) *ListOpenstackSecurityGroupsParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithDatacenterName adds the datacenterName to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) WithDatacenterName(datacenterName *string) *ListOpenstackSecurityGroupsParams {
	o.SetDatacenterName(datacenterName)
	return o
}

// SetDatacenterName adds the datacenterName to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) SetDatacenterName(datacenterName *string) {
	o.DatacenterName = datacenterName
}

// WithDomain adds the domain to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) WithDomain(domain *string) *ListOpenstackSecurityGroupsParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) SetDomain(domain *string) {
	o.Domain = domain
}

// WithOIDCAuthentication adds the oIDCAuthentication to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) WithOIDCAuthentication(oIDCAuthentication *bool) *ListOpenstackSecurityGroupsParams {
	o.SetOIDCAuthentication(oIDCAuthentication)
	return o
}

// SetOIDCAuthentication adds the oIdCAuthentication to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) SetOIDCAuthentication(oIDCAuthentication *bool) {
	o.OIDCAuthentication = oIDCAuthentication
}

// WithPassword adds the password to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) WithPassword(password *string) *ListOpenstackSecurityGroupsParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) SetPassword(password *string) {
	o.Password = password
}

// WithTenant adds the tenant to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) WithTenant(tenant *string) *ListOpenstackSecurityGroupsParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantID adds the tenantID to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) WithTenantID(tenantID *string) *ListOpenstackSecurityGroupsParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithUsername adds the username to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) WithUsername(username *string) *ListOpenstackSecurityGroupsParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the list openstack security groups params
func (o *ListOpenstackSecurityGroupsParams) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *ListOpenstackSecurityGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Tenant != nil {

		// header param Tenant
		if err := r.SetHeaderParam("Tenant", *o.Tenant); err != nil {
			return err
		}

	}

	if o.TenantID != nil {

		// header param TenantID
		if err := r.SetHeaderParam("TenantID", *o.TenantID); err != nil {
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
