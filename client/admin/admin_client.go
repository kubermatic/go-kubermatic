// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new admin API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for admin API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAdmissionPlugin(params *DeleteAdmissionPluginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAdmissionPluginOK, error)

	DeleteSeed(params *DeleteSeedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSeedOK, error)

	GetAdmins(params *GetAdminsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAdminsOK, error)

	GetAdmissionPlugin(params *GetAdmissionPluginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAdmissionPluginOK, error)

	GetKubermaticCustomLinks(params *GetKubermaticCustomLinksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetKubermaticCustomLinksOK, error)

	GetKubermaticSettings(params *GetKubermaticSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetKubermaticSettingsOK, error)

	GetSeed(params *GetSeedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSeedOK, error)

	ListAdmissionPlugins(params *ListAdmissionPluginsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAdmissionPluginsOK, error)

	ListSeeds(params *ListSeedsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSeedsOK, error)

	PatchKubermaticSettings(params *PatchKubermaticSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchKubermaticSettingsOK, error)

	SetAdmin(params *SetAdminParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SetAdminOK, error)

	UpdateAdmissionPlugin(params *UpdateAdmissionPluginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAdmissionPluginOK, error)

	UpdateSeed(params *UpdateSeedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSeedOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteAdmissionPlugin deletes the admission plugin
*/
func (a *Client) DeleteAdmissionPlugin(params *DeleteAdmissionPluginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAdmissionPluginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAdmissionPluginParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAdmissionPlugin",
		Method:             "DELETE",
		PathPattern:        "/api/v1/admin/admission/plugins/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAdmissionPluginReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAdmissionPluginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteAdmissionPluginDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteSeed deletes the seed c r d object from the kubermatic
*/
func (a *Client) DeleteSeed(params *DeleteSeedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSeedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSeedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSeed",
		Method:             "DELETE",
		PathPattern:        "/api/v1/admin/seeds/{seed_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSeedReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteSeedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSeedDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetAdmins returns list of admin users
*/
func (a *Client) GetAdmins(params *GetAdminsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAdminsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdminsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAdmins",
		Method:             "GET",
		PathPattern:        "/api/v1/admin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAdminsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAdminsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAdminsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetAdmissionPlugin gets the admission plugin
*/
func (a *Client) GetAdmissionPlugin(params *GetAdmissionPluginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAdmissionPluginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdmissionPluginParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAdmissionPlugin",
		Method:             "GET",
		PathPattern:        "/api/v1/admin/admission/plugins/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAdmissionPluginReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAdmissionPluginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAdmissionPluginDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetKubermaticCustomLinks gets the custom links
*/
func (a *Client) GetKubermaticCustomLinks(params *GetKubermaticCustomLinksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetKubermaticCustomLinksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKubermaticCustomLinksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getKubermaticCustomLinks",
		Method:             "GET",
		PathPattern:        "/api/v1/admin/settings/customlinks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKubermaticCustomLinksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetKubermaticCustomLinksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetKubermaticCustomLinksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetKubermaticSettings gets the global settings
*/
func (a *Client) GetKubermaticSettings(params *GetKubermaticSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetKubermaticSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKubermaticSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getKubermaticSettings",
		Method:             "GET",
		PathPattern:        "/api/v1/admin/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKubermaticSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetKubermaticSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetKubermaticSettingsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSeed returns the seed object
*/
func (a *Client) GetSeed(params *GetSeedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSeedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSeedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSeed",
		Method:             "GET",
		PathPattern:        "/api/v1/admin/seeds/{seed_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSeedReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSeedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSeedDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListAdmissionPlugins returns all admission plugins from the c r ds
*/
func (a *Client) ListAdmissionPlugins(params *ListAdmissionPluginsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAdmissionPluginsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAdmissionPluginsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAdmissionPlugins",
		Method:             "GET",
		PathPattern:        "/api/v1/admin/admission/plugins",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAdmissionPluginsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAdmissionPluginsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAdmissionPluginsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListSeeds returns all seeds from the c r ds
*/
func (a *Client) ListSeeds(params *ListSeedsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSeedsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSeedsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listSeeds",
		Method:             "GET",
		PathPattern:        "/api/v1/admin/seeds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSeedsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListSeedsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListSeedsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PatchKubermaticSettings patches the global settings
*/
func (a *Client) PatchKubermaticSettings(params *PatchKubermaticSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchKubermaticSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchKubermaticSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchKubermaticSettings",
		Method:             "PATCH",
		PathPattern:        "/api/v1/admin/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchKubermaticSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchKubermaticSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchKubermaticSettingsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  SetAdmin allows setting and clearing admin role for users
*/
func (a *Client) SetAdmin(params *SetAdminParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SetAdminOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetAdminParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setAdmin",
		Method:             "PUT",
		PathPattern:        "/api/v1/admin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetAdminReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetAdminOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SetAdminDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateAdmissionPlugin updates the admission plugin
*/
func (a *Client) UpdateAdmissionPlugin(params *UpdateAdmissionPluginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAdmissionPluginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAdmissionPluginParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateAdmissionPlugin",
		Method:             "PATCH",
		PathPattern:        "/api/v1/admin/admission/plugins/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAdmissionPluginReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateAdmissionPluginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateAdmissionPluginDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateSeed updates the seed
*/
func (a *Client) UpdateSeed(params *UpdateSeedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSeedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSeedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateSeed",
		Method:             "PATCH",
		PathPattern:        "/api/v1/admin/seeds/{seed_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSeedReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateSeedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateSeedDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
