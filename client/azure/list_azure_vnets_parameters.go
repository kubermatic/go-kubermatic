// Code generated by go-swagger; DO NOT EDIT.

package azure

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

// NewListAzureVnetsParams creates a new ListAzureVnetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAzureVnetsParams() *ListAzureVnetsParams {
	return &ListAzureVnetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAzureVnetsParamsWithTimeout creates a new ListAzureVnetsParams object
// with the ability to set a timeout on a request.
func NewListAzureVnetsParamsWithTimeout(timeout time.Duration) *ListAzureVnetsParams {
	return &ListAzureVnetsParams{
		timeout: timeout,
	}
}

// NewListAzureVnetsParamsWithContext creates a new ListAzureVnetsParams object
// with the ability to set a context for a request.
func NewListAzureVnetsParamsWithContext(ctx context.Context) *ListAzureVnetsParams {
	return &ListAzureVnetsParams{
		Context: ctx,
	}
}

// NewListAzureVnetsParamsWithHTTPClient creates a new ListAzureVnetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAzureVnetsParamsWithHTTPClient(client *http.Client) *ListAzureVnetsParams {
	return &ListAzureVnetsParams{
		HTTPClient: client,
	}
}

/*
ListAzureVnetsParams contains all the parameters to send to the API endpoint

	for the list azure vnets operation.

	Typically these are written to a http.Request.
*/
type ListAzureVnetsParams struct {

	// ClientID.
	ClientID *string

	// ClientSecret.
	ClientSecret *string

	// Credential.
	Credential *string

	// Location.
	Location *string

	// ResourceGroup.
	ResourceGroup *string

	// SubscriptionID.
	SubscriptionID *string

	// TenantID.
	TenantID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list azure vnets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAzureVnetsParams) WithDefaults() *ListAzureVnetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list azure vnets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAzureVnetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list azure vnets params
func (o *ListAzureVnetsParams) WithTimeout(timeout time.Duration) *ListAzureVnetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list azure vnets params
func (o *ListAzureVnetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list azure vnets params
func (o *ListAzureVnetsParams) WithContext(ctx context.Context) *ListAzureVnetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list azure vnets params
func (o *ListAzureVnetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list azure vnets params
func (o *ListAzureVnetsParams) WithHTTPClient(client *http.Client) *ListAzureVnetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list azure vnets params
func (o *ListAzureVnetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the list azure vnets params
func (o *ListAzureVnetsParams) WithClientID(clientID *string) *ListAzureVnetsParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the list azure vnets params
func (o *ListAzureVnetsParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithClientSecret adds the clientSecret to the list azure vnets params
func (o *ListAzureVnetsParams) WithClientSecret(clientSecret *string) *ListAzureVnetsParams {
	o.SetClientSecret(clientSecret)
	return o
}

// SetClientSecret adds the clientSecret to the list azure vnets params
func (o *ListAzureVnetsParams) SetClientSecret(clientSecret *string) {
	o.ClientSecret = clientSecret
}

// WithCredential adds the credential to the list azure vnets params
func (o *ListAzureVnetsParams) WithCredential(credential *string) *ListAzureVnetsParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list azure vnets params
func (o *ListAzureVnetsParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithLocation adds the location to the list azure vnets params
func (o *ListAzureVnetsParams) WithLocation(location *string) *ListAzureVnetsParams {
	o.SetLocation(location)
	return o
}

// SetLocation adds the location to the list azure vnets params
func (o *ListAzureVnetsParams) SetLocation(location *string) {
	o.Location = location
}

// WithResourceGroup adds the resourceGroup to the list azure vnets params
func (o *ListAzureVnetsParams) WithResourceGroup(resourceGroup *string) *ListAzureVnetsParams {
	o.SetResourceGroup(resourceGroup)
	return o
}

// SetResourceGroup adds the resourceGroup to the list azure vnets params
func (o *ListAzureVnetsParams) SetResourceGroup(resourceGroup *string) {
	o.ResourceGroup = resourceGroup
}

// WithSubscriptionID adds the subscriptionID to the list azure vnets params
func (o *ListAzureVnetsParams) WithSubscriptionID(subscriptionID *string) *ListAzureVnetsParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the list azure vnets params
func (o *ListAzureVnetsParams) SetSubscriptionID(subscriptionID *string) {
	o.SubscriptionID = subscriptionID
}

// WithTenantID adds the tenantID to the list azure vnets params
func (o *ListAzureVnetsParams) WithTenantID(tenantID *string) *ListAzureVnetsParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the list azure vnets params
func (o *ListAzureVnetsParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *ListAzureVnetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientID != nil {

		// header param ClientID
		if err := r.SetHeaderParam("ClientID", *o.ClientID); err != nil {
			return err
		}
	}

	if o.ClientSecret != nil {

		// header param ClientSecret
		if err := r.SetHeaderParam("ClientSecret", *o.ClientSecret); err != nil {
			return err
		}
	}

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	if o.Location != nil {

		// header param Location
		if err := r.SetHeaderParam("Location", *o.Location); err != nil {
			return err
		}
	}

	if o.ResourceGroup != nil {

		// header param ResourceGroup
		if err := r.SetHeaderParam("ResourceGroup", *o.ResourceGroup); err != nil {
			return err
		}
	}

	if o.SubscriptionID != nil {

		// header param SubscriptionID
		if err := r.SetHeaderParam("SubscriptionID", *o.SubscriptionID); err != nil {
			return err
		}
	}

	if o.TenantID != nil {

		// header param TenantID
		if err := r.SetHeaderParam("TenantID", *o.TenantID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
