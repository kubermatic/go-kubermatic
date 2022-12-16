// Code generated by go-swagger; DO NOT EDIT.

package seed

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new seed API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for seed API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetSeedOverview(params *GetSeedOverviewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSeedOverviewOK, error)

	GetSeedSettings(params *GetSeedSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSeedSettingsOK, error)

	ListSeedNames(params *ListSeedNamesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSeedNamesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetSeedOverview returns seed s overview
*/
func (a *Client) GetSeedOverview(params *GetSeedOverviewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSeedOverviewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSeedOverviewParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSeedOverview",
		Method:             "GET",
		PathPattern:        "/api/v2/seeds/{seed_name}/overview",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSeedOverviewReader{formats: a.formats},
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
	success, ok := result.(*GetSeedOverviewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSeedOverviewDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetSeedSettings gets the seed settings
*/
func (a *Client) GetSeedSettings(params *GetSeedSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSeedSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSeedSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSeedSettings",
		Method:             "GET",
		PathPattern:        "/api/v2/seeds/{seed_name}/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSeedSettingsReader{formats: a.formats},
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
	success, ok := result.(*GetSeedSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSeedSettingsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListSeedNames list seed names API
*/
func (a *Client) ListSeedNames(params *ListSeedNamesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSeedNamesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSeedNamesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listSeedNames",
		Method:             "GET",
		PathPattern:        "/api/v1/seed",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSeedNamesReader{formats: a.formats},
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
	success, ok := result.(*ListSeedNamesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListSeedNamesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
