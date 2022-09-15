// Code generated by go-swagger; DO NOT EDIT.

package cniversion

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cniversion API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cniversion API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ListVersionsByCNIPlugin(params *ListVersionsByCNIPluginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVersionsByCNIPluginOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ListVersionsByCNIPlugin Lists all CNI Plugin versions that are supported for a given CNI plugin type
*/
func (a *Client) ListVersionsByCNIPlugin(params *ListVersionsByCNIPluginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVersionsByCNIPluginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVersionsByCNIPluginParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listVersionsByCNIPlugin",
		Method:             "GET",
		PathPattern:        "/api/v2/cni/{cni_plugin_type}/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVersionsByCNIPluginReader{formats: a.formats},
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
	success, ok := result.(*ListVersionsByCNIPluginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVersionsByCNIPluginDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
