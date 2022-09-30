// Code generated by go-swagger; DO NOT EDIT.

package kubevirt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new kubevirt API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for kubevirt API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ListKubeVirtInstancetypes(params *ListKubeVirtInstancetypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListKubeVirtInstancetypesOK, error)

	ListKubeVirtInstancetypesNoCredentials(params *ListKubeVirtInstancetypesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListKubeVirtInstancetypesNoCredentialsOK, error)

	ListKubeVirtPreferences(params *ListKubeVirtPreferencesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListKubeVirtPreferencesOK, error)

	ListKubeVirtPreferencesNoCredentials(params *ListKubeVirtPreferencesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListKubeVirtPreferencesNoCredentialsOK, error)

	ListKubeVirtVMIPresets(params *ListKubeVirtVMIPresetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListKubeVirtVMIPresetsOK, error)

	ListKubeVirtVMIPresetsNoCredentials(params *ListKubeVirtVMIPresetsNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListKubeVirtVMIPresetsNoCredentialsOK, error)

	ListKubevirtStorageClasses(params *ListKubevirtStorageClassesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListKubevirtStorageClassesOK, error)

	ListKubevirtStorageClassesNoCredentials(params *ListKubevirtStorageClassesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListKubevirtStorageClassesNoCredentialsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ListKubeVirtInstancetypes lists available kube virt virtual machine instancetype
*/
func (a *Client) ListKubeVirtInstancetypes(params *ListKubeVirtInstancetypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListKubeVirtInstancetypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListKubeVirtInstancetypesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listKubeVirtInstancetypes",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/kubevirt/instancetypes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListKubeVirtInstancetypesReader{formats: a.formats},
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
	success, ok := result.(*ListKubeVirtInstancetypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListKubeVirtInstancetypesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListKubeVirtInstancetypesNoCredentials Lists available VirtualMachineInstancetype
*/
func (a *Client) ListKubeVirtInstancetypesNoCredentials(params *ListKubeVirtInstancetypesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListKubeVirtInstancetypesNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListKubeVirtInstancetypesNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listKubeVirtInstancetypesNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/kubevirt/instancetypes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListKubeVirtInstancetypesNoCredentialsReader{formats: a.formats},
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
	success, ok := result.(*ListKubeVirtInstancetypesNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListKubeVirtInstancetypesNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListKubeVirtPreferences lists available kube virt virtual machine preference
*/
func (a *Client) ListKubeVirtPreferences(params *ListKubeVirtPreferencesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListKubeVirtPreferencesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListKubeVirtPreferencesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listKubeVirtPreferences",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/kubevirt/preferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListKubeVirtPreferencesReader{formats: a.formats},
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
	success, ok := result.(*ListKubeVirtPreferencesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListKubeVirtPreferencesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListKubeVirtPreferencesNoCredentials Lists available VirtualMachinePreference
*/
func (a *Client) ListKubeVirtPreferencesNoCredentials(params *ListKubeVirtPreferencesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListKubeVirtPreferencesNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListKubeVirtPreferencesNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listKubeVirtPreferencesNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/kubevirt/preferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListKubeVirtPreferencesNoCredentialsReader{formats: a.formats},
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
	success, ok := result.(*ListKubeVirtPreferencesNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListKubeVirtPreferencesNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListKubeVirtVMIPresets lists available kube virt virtual machine instance preset
*/
func (a *Client) ListKubeVirtVMIPresets(params *ListKubeVirtVMIPresetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListKubeVirtVMIPresetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListKubeVirtVMIPresetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listKubeVirtVMIPresets",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/kubevirt/vmflavors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListKubeVirtVMIPresetsReader{formats: a.formats},
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
	success, ok := result.(*ListKubeVirtVMIPresetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListKubeVirtVMIPresetsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListKubeVirtVMIPresetsNoCredentials Lists available VirtualMachineInstancePreset
*/
func (a *Client) ListKubeVirtVMIPresetsNoCredentials(params *ListKubeVirtVMIPresetsNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListKubeVirtVMIPresetsNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListKubeVirtVMIPresetsNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listKubeVirtVMIPresetsNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/kubevirt/vmflavors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListKubeVirtVMIPresetsNoCredentialsReader{formats: a.formats},
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
	success, ok := result.(*ListKubeVirtVMIPresetsNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListKubeVirtVMIPresetsNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListKubevirtStorageClasses lists available k8s storage classes in the kubevirt cluster
*/
func (a *Client) ListKubevirtStorageClasses(params *ListKubevirtStorageClassesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListKubevirtStorageClassesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListKubevirtStorageClassesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listKubevirtStorageClasses",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/kubevirt/storageclasses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListKubevirtStorageClassesReader{formats: a.formats},
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
	success, ok := result.(*ListKubevirtStorageClassesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListKubevirtStorageClassesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListKubevirtStorageClassesNoCredentials List Storage Classes
*/
func (a *Client) ListKubevirtStorageClassesNoCredentials(params *ListKubevirtStorageClassesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListKubevirtStorageClassesNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListKubevirtStorageClassesNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listKubevirtStorageClassesNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/kubevirt/storageclasses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListKubevirtStorageClassesNoCredentialsReader{formats: a.formats},
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
	success, ok := result.(*ListKubevirtStorageClassesNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListKubevirtStorageClassesNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
