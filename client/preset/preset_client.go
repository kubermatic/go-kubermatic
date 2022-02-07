// Code generated by go-swagger; DO NOT EDIT.

package preset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new preset API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for preset API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePreset(params *CreatePresetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePresetOK, error)

	DeletePreset(params *DeletePresetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePresetOK, error)

	DeletePresetProvider(params *DeletePresetProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePresetProviderOK, error)

	DeleteProviderPreset(params *DeleteProviderPresetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteProviderPresetOK, error)

	ListPresets(params *ListPresetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListPresetsOK, error)

	ListProviderPresets(params *ListProviderPresetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProviderPresetsOK, error)

	UpdatePreset(params *UpdatePresetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePresetOK, error)

	UpdatePresetStatus(params *UpdatePresetStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePresetStatusOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreatePreset Creates the preset
*/
func (a *Client) CreatePreset(params *CreatePresetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePresetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePresetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPreset",
		Method:             "POST",
		PathPattern:        "/api/v2/providers/{provider_name}/presets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePresetReader{formats: a.formats},
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
	success, ok := result.(*CreatePresetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreatePresetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeletePreset removes preset
*/
func (a *Client) DeletePreset(params *DeletePresetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePresetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePresetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePreset",
		Method:             "DELETE",
		PathPattern:        "/api/v2/presets/{preset_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePresetReader{formats: a.formats},
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
	success, ok := result.(*DeletePresetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeletePresetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeletePresetProvider removes selected preset s provider
*/
func (a *Client) DeletePresetProvider(params *DeletePresetProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePresetProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePresetProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePresetProvider",
		Method:             "DELETE",
		PathPattern:        "/api/v2/presets/{preset_name}/provider/{provider_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePresetProviderReader{formats: a.formats},
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
	success, ok := result.(*DeletePresetProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeletePresetProviderDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteProviderPreset deletes provider preset

  This endpoint has been depreciated in favour of /presets/{presets_name} and /presets/{preset_name}/providers/{provider_name}.
*/
func (a *Client) DeleteProviderPreset(params *DeleteProviderPresetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteProviderPresetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProviderPresetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteProviderPreset",
		Method:             "DELETE",
		PathPattern:        "/api/v2/providers/{provider_name}/presets/{preset_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteProviderPresetReader{formats: a.formats},
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
	success, ok := result.(*DeleteProviderPresetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteProviderPresetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListPresets Lists presets
*/
func (a *Client) ListPresets(params *ListPresetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListPresetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPresetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listPresets",
		Method:             "GET",
		PathPattern:        "/api/v2/presets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListPresetsReader{formats: a.formats},
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
	success, ok := result.(*ListPresetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListPresetsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListProviderPresets Lists presets for the provider
*/
func (a *Client) ListProviderPresets(params *ListProviderPresetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProviderPresetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProviderPresetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProviderPresets",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/{provider_name}/presets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProviderPresetsReader{formats: a.formats},
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
	success, ok := result.(*ListProviderPresetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProviderPresetsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdatePreset Updates provider preset
*/
func (a *Client) UpdatePreset(params *UpdatePresetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePresetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePresetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePreset",
		Method:             "PUT",
		PathPattern:        "/api/v2/providers/{provider_name}/presets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePresetReader{formats: a.formats},
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
	success, ok := result.(*UpdatePresetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdatePresetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdatePresetStatus updates the status of a preset it can enable or disable it so that it won t be listed by the list endpoints
*/
func (a *Client) UpdatePresetStatus(params *UpdatePresetStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePresetStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePresetStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePresetStatus",
		Method:             "PUT",
		PathPattern:        "/api/v2/presets/{preset_name}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePresetStatusReader{formats: a.formats},
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
	success, ok := result.(*UpdatePresetStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdatePresetStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
