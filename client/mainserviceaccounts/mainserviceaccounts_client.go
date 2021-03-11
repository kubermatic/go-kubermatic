// Code generated by go-swagger; DO NOT EDIT.

package mainserviceaccounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new mainserviceaccounts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for mainserviceaccounts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateMainServiceAccount(params *CreateMainServiceAccountParams, authInfo runtime.ClientAuthInfoWriter) (*CreateMainServiceAccountCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateMainServiceAccount Creates the given service account
*/
func (a *Client) CreateMainServiceAccount(params *CreateMainServiceAccountParams, authInfo runtime.ClientAuthInfoWriter) (*CreateMainServiceAccountCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMainServiceAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createMainServiceAccount",
		Method:             "POST",
		PathPattern:        "/api/v2/serviceaccounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMainServiceAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateMainServiceAccountCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateMainServiceAccountDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
