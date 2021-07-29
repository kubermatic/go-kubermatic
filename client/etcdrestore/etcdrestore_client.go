// Code generated by go-swagger; DO NOT EDIT.

package etcdrestore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new etcdrestore API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for etcdrestore API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateEtcdRestore(params *CreateEtcdRestoreParams, authInfo runtime.ClientAuthInfoWriter) (*CreateEtcdRestoreCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateEtcdRestore Creates a etcd backup restore for a given cluster
*/
func (a *Client) CreateEtcdRestore(params *CreateEtcdRestoreParams, authInfo runtime.ClientAuthInfoWriter) (*CreateEtcdRestoreCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEtcdRestoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createEtcdRestore",
		Method:             "POST",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateEtcdRestoreReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateEtcdRestoreCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateEtcdRestoreDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
