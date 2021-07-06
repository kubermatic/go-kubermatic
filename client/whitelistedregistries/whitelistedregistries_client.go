// Code generated by go-swagger; DO NOT EDIT.

package whitelistedregistries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new whitelistedregistries API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for whitelistedregistries API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteWhitelistedRegistry(params *DeleteWhitelistedRegistryParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteWhitelistedRegistryOK, error)

	GetWhitelistedRegistry(params *GetWhitelistedRegistryParams, authInfo runtime.ClientAuthInfoWriter) (*GetWhitelistedRegistryOK, error)

	PatchWhitelistedRegistry(params *PatchWhitelistedRegistryParams, authInfo runtime.ClientAuthInfoWriter) (*PatchWhitelistedRegistryOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteWhitelistedRegistry deletes the given whitelisted registry
*/
func (a *Client) DeleteWhitelistedRegistry(params *DeleteWhitelistedRegistryParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteWhitelistedRegistryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWhitelistedRegistryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteWhitelistedRegistry",
		Method:             "DELETE",
		PathPattern:        "/api/v2/whitelistedregistries/{whitelisted_registry}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteWhitelistedRegistryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteWhitelistedRegistryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteWhitelistedRegistryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetWhitelistedRegistry Get whitelisted registries specified by name
*/
func (a *Client) GetWhitelistedRegistry(params *GetWhitelistedRegistryParams, authInfo runtime.ClientAuthInfoWriter) (*GetWhitelistedRegistryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWhitelistedRegistryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWhitelistedRegistry",
		Method:             "GET",
		PathPattern:        "/api/v2/whitelistedregistries/{whitelisted_registry}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWhitelistedRegistryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWhitelistedRegistryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetWhitelistedRegistryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PatchWhitelistedRegistry Patch a specified whitelisted registry
*/
func (a *Client) PatchWhitelistedRegistry(params *PatchWhitelistedRegistryParams, authInfo runtime.ClientAuthInfoWriter) (*PatchWhitelistedRegistryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchWhitelistedRegistryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchWhitelistedRegistry",
		Method:             "PATCH",
		PathPattern:        "/api/v2/whitelistedregistries/{whitelisted_registry}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchWhitelistedRegistryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchWhitelistedRegistryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchWhitelistedRegistryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
