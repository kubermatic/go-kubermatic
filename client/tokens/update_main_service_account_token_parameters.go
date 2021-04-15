// Code generated by go-swagger; DO NOT EDIT.

package tokens

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

// NewUpdateMainServiceAccountTokenParams creates a new UpdateMainServiceAccountTokenParams object
// with the default values initialized.
func NewUpdateMainServiceAccountTokenParams() *UpdateMainServiceAccountTokenParams {
	var ()
	return &UpdateMainServiceAccountTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMainServiceAccountTokenParamsWithTimeout creates a new UpdateMainServiceAccountTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateMainServiceAccountTokenParamsWithTimeout(timeout time.Duration) *UpdateMainServiceAccountTokenParams {
	var ()
	return &UpdateMainServiceAccountTokenParams{

		timeout: timeout,
	}
}

// NewUpdateMainServiceAccountTokenParamsWithContext creates a new UpdateMainServiceAccountTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateMainServiceAccountTokenParamsWithContext(ctx context.Context) *UpdateMainServiceAccountTokenParams {
	var ()
	return &UpdateMainServiceAccountTokenParams{

		Context: ctx,
	}
}

// NewUpdateMainServiceAccountTokenParamsWithHTTPClient creates a new UpdateMainServiceAccountTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateMainServiceAccountTokenParamsWithHTTPClient(client *http.Client) *UpdateMainServiceAccountTokenParams {
	var ()
	return &UpdateMainServiceAccountTokenParams{
		HTTPClient: client,
	}
}

/*UpdateMainServiceAccountTokenParams contains all the parameters to send to the API endpoint
for the update main service account token operation typically these are written to a http.Request
*/
type UpdateMainServiceAccountTokenParams struct {

	/*Body*/
	Body *models.PublicServiceAccountToken
	/*ServiceaccountID*/
	ServiceAccountID string
	/*TokenID*/
	TokenID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update main service account token params
func (o *UpdateMainServiceAccountTokenParams) WithTimeout(timeout time.Duration) *UpdateMainServiceAccountTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update main service account token params
func (o *UpdateMainServiceAccountTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update main service account token params
func (o *UpdateMainServiceAccountTokenParams) WithContext(ctx context.Context) *UpdateMainServiceAccountTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update main service account token params
func (o *UpdateMainServiceAccountTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update main service account token params
func (o *UpdateMainServiceAccountTokenParams) WithHTTPClient(client *http.Client) *UpdateMainServiceAccountTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update main service account token params
func (o *UpdateMainServiceAccountTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update main service account token params
func (o *UpdateMainServiceAccountTokenParams) WithBody(body *models.PublicServiceAccountToken) *UpdateMainServiceAccountTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update main service account token params
func (o *UpdateMainServiceAccountTokenParams) SetBody(body *models.PublicServiceAccountToken) {
	o.Body = body
}

// WithServiceAccountID adds the serviceaccountID to the update main service account token params
func (o *UpdateMainServiceAccountTokenParams) WithServiceAccountID(serviceaccountID string) *UpdateMainServiceAccountTokenParams {
	o.SetServiceAccountID(serviceaccountID)
	return o
}

// SetServiceAccountID adds the serviceaccountId to the update main service account token params
func (o *UpdateMainServiceAccountTokenParams) SetServiceAccountID(serviceaccountID string) {
	o.ServiceAccountID = serviceaccountID
}

// WithTokenID adds the tokenID to the update main service account token params
func (o *UpdateMainServiceAccountTokenParams) WithTokenID(tokenID string) *UpdateMainServiceAccountTokenParams {
	o.SetTokenID(tokenID)
	return o
}

// SetTokenID adds the tokenId to the update main service account token params
func (o *UpdateMainServiceAccountTokenParams) SetTokenID(tokenID string) {
	o.TokenID = tokenID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMainServiceAccountTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param serviceaccount_id
	if err := r.SetPathParam("serviceaccount_id", o.ServiceAccountID); err != nil {
		return err
	}

	// path param token_id
	if err := r.SetPathParam("token_id", o.TokenID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}