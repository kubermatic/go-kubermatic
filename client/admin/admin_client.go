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

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAdmissionPlugin(params *DeleteAdmissionPluginParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAdmissionPluginOK, error)

	GetAdmins(params *GetAdminsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAdminsOK, error)

	GetAdmissionPlugin(params *GetAdmissionPluginParams, authInfo runtime.ClientAuthInfoWriter) (*GetAdmissionPluginOK, error)

	GetKubermaticSettings(params *GetKubermaticSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetKubermaticSettingsOK, error)

	ListAdmissionPlugins(params *ListAdmissionPluginsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAdmissionPluginsOK, error)

	PatchKubermaticSettings(params *PatchKubermaticSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*PatchKubermaticSettingsOK, error)

	SetAdmin(params *SetAdminParams, authInfo runtime.ClientAuthInfoWriter) (*SetAdminOK, error)

	UpdateAdmissionPlugin(params *UpdateAdmissionPluginParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAdmissionPluginOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteAdmissionPlugin deletes the admission plugin
*/
func (a *Client) DeleteAdmissionPlugin(params *DeleteAdmissionPluginParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAdmissionPluginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAdmissionPluginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
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
	})
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
  GetAdmins returns list of admin users
*/
func (a *Client) GetAdmins(params *GetAdminsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAdminsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdminsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
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
	})
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
func (a *Client) GetAdmissionPlugin(params *GetAdmissionPluginParams, authInfo runtime.ClientAuthInfoWriter) (*GetAdmissionPluginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdmissionPluginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
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
	})
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
  GetKubermaticSettings gets the global settings
*/
func (a *Client) GetKubermaticSettings(params *GetKubermaticSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetKubermaticSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKubermaticSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
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
	})
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
  ListAdmissionPlugins returns all admission plugins from the c r ds
*/
func (a *Client) ListAdmissionPlugins(params *ListAdmissionPluginsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAdmissionPluginsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAdmissionPluginsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
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
	})
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
  PatchKubermaticSettings patches the global settings
*/
func (a *Client) PatchKubermaticSettings(params *PatchKubermaticSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*PatchKubermaticSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchKubermaticSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
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
	})
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
func (a *Client) SetAdmin(params *SetAdminParams, authInfo runtime.ClientAuthInfoWriter) (*SetAdminOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetAdminParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
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
	})
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
func (a *Client) UpdateAdmissionPlugin(params *UpdateAdmissionPluginParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAdmissionPluginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAdmissionPluginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
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
	})
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

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
