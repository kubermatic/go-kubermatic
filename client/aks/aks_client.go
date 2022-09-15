// Code generated by go-swagger; DO NOT EDIT.

package aks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new aks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for aks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ListAKSLocations(params *ListAKSLocationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAKSLocationsOK, error)

	ListAKSNodePoolModes(params *ListAKSNodePoolModesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAKSNodePoolModesOK, error)

	ListAKSNodeVersionsNoCredentials(params *ListAKSNodeVersionsNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAKSNodeVersionsNoCredentialsOK, error)

	ListAKSResourceGroups(params *ListAKSResourceGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAKSResourceGroupsOK, error)

	ListAKSVMSizes(params *ListAKSVMSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAKSVMSizesOK, error)

	ListAKSVMSizesNoCredentials(params *ListAKSVMSizesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAKSVMSizesNoCredentialsOK, error)

	ListAKSVersions(params *ListAKSVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAKSVersionsOK, error)

	ValidateAKSCredentials(params *ValidateAKSCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ValidateAKSCredentialsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ListAKSLocations lists a k s recommended locations
*/
func (a *Client) ListAKSLocations(params *ListAKSLocationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAKSLocationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAKSLocationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAKSLocations",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/aks/locations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAKSLocationsReader{formats: a.formats},
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
	success, ok := result.(*ListAKSLocationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAKSLocationsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListAKSNodePoolModes gets the a k s node pool modes
*/
func (a *Client) ListAKSNodePoolModes(params *ListAKSNodePoolModesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAKSNodePoolModesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAKSNodePoolModesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAKSNodePoolModes",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/aks/modes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAKSNodePoolModesReader{formats: a.formats},
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
	success, ok := result.(*ListAKSNodePoolModesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAKSNodePoolModesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListAKSNodeVersionsNoCredentials gets a k s nodepool available versions
*/
func (a *Client) ListAKSNodeVersionsNoCredentials(params *ListAKSNodeVersionsNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAKSNodeVersionsNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAKSNodeVersionsNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAKSNodeVersionsNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/aks/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAKSNodeVersionsNoCredentialsReader{formats: a.formats},
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
	success, ok := result.(*ListAKSNodeVersionsNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAKSNodeVersionsNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListAKSResourceGroups lists resource groups in an azure subscription
*/
func (a *Client) ListAKSResourceGroups(params *ListAKSResourceGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAKSResourceGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAKSResourceGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAKSResourceGroups",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/aks/resourcegroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAKSResourceGroupsReader{formats: a.formats},
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
	success, ok := result.(*ListAKSResourceGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAKSResourceGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListAKSVMSizes lists a k s available VM sizes in an azure region
*/
func (a *Client) ListAKSVMSizes(params *ListAKSVMSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAKSVMSizesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAKSVMSizesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAKSVMSizes",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/aks/vmsizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAKSVMSizesReader{formats: a.formats},
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
	success, ok := result.(*ListAKSVMSizesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAKSVMSizesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListAKSVMSizesNoCredentials gets a k s available VM sizes in an azure region
*/
func (a *Client) ListAKSVMSizesNoCredentials(params *ListAKSVMSizesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAKSVMSizesNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAKSVMSizesNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAKSVMSizesNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/aks/vmsizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAKSVMSizesNoCredentialsReader{formats: a.formats},
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
	success, ok := result.(*ListAKSVMSizesNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAKSVMSizesNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListAKSVersions Lists AKS versions
*/
func (a *Client) ListAKSVersions(params *ListAKSVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAKSVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAKSVersionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAKSVersions",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/aks/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAKSVersionsReader{formats: a.formats},
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
	success, ok := result.(*ListAKSVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAKSVersionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ValidateAKSCredentials Validates AKS credentials
*/
func (a *Client) ValidateAKSCredentials(params *ValidateAKSCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ValidateAKSCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateAKSCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "validateAKSCredentials",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/aks/validatecredentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateAKSCredentialsReader{formats: a.formats},
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
	success, ok := result.(*ValidateAKSCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ValidateAKSCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
