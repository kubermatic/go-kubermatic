// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vsphere API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vsphere API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ListProjectVSphereDatastores(params *ListProjectVSphereDatastoresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectVSphereDatastoresOK, error)

	ListProjectVSphereFolders(params *ListProjectVSphereFoldersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectVSphereFoldersOK, error)

	ListProjectVSphereNetworks(params *ListProjectVSphereNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectVSphereNetworksOK, error)

	ListProjectVSphereTagCategories(params *ListProjectVSphereTagCategoriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectVSphereTagCategoriesOK, error)

	ListProjectVSphereTagsForTagCategories(params *ListProjectVSphereTagsForTagCategoriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectVSphereTagsForTagCategoriesOK, error)

	ListVSphereDatastores(params *ListVSphereDatastoresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVSphereDatastoresOK, error)

	ListVSphereFolders(params *ListVSphereFoldersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVSphereFoldersOK, error)

	ListVSphereFoldersNoCredentials(params *ListVSphereFoldersNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVSphereFoldersNoCredentialsOK, error)

	ListVSphereFoldersNoCredentialsV2(params *ListVSphereFoldersNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVSphereFoldersNoCredentialsV2OK, error)

	ListVSphereNetworks(params *ListVSphereNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVSphereNetworksOK, error)

	ListVSphereNetworksNoCredentials(params *ListVSphereNetworksNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVSphereNetworksNoCredentialsOK, error)

	ListVSphereNetworksNoCredentialsV2(params *ListVSphereNetworksNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVSphereNetworksNoCredentialsV2OK, error)

	ListVSphereTagCategoriesNoCredentials(params *ListVSphereTagCategoriesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVSphereTagCategoriesNoCredentialsOK, error)

	ListVSphereTagsForTagCategoryNoCredentials(params *ListVSphereTagsForTagCategoryNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVSphereTagsForTagCategoryNoCredentialsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ListProjectVSphereDatastores lists datastores from v sphere datacenter
*/
func (a *Client) ListProjectVSphereDatastores(params *ListProjectVSphereDatastoresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectVSphereDatastoresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectVSphereDatastoresParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectVSphereDatastores",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/vsphere/datastores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectVSphereDatastoresReader{formats: a.formats},
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
	success, ok := result.(*ListProjectVSphereDatastoresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectVSphereDatastoresDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectVSphereFolders lists folders from v sphere datacenter
*/
func (a *Client) ListProjectVSphereFolders(params *ListProjectVSphereFoldersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectVSphereFoldersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectVSphereFoldersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectVSphereFolders",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/vsphere/folders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectVSphereFoldersReader{formats: a.formats},
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
	success, ok := result.(*ListProjectVSphereFoldersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectVSphereFoldersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectVSphereNetworks lists networks from v sphere datacenter
*/
func (a *Client) ListProjectVSphereNetworks(params *ListProjectVSphereNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectVSphereNetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectVSphereNetworksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectVSphereNetworks",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/vsphere/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectVSphereNetworksReader{formats: a.formats},
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
	success, ok := result.(*ListProjectVSphereNetworksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectVSphereNetworksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectVSphereTagCategories lists tag categories from v sphere datacenter
*/
func (a *Client) ListProjectVSphereTagCategories(params *ListProjectVSphereTagCategoriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectVSphereTagCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectVSphereTagCategoriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectVSphereTagCategories",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/vsphere/tagcategories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectVSphereTagCategoriesReader{formats: a.formats},
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
	success, ok := result.(*ListProjectVSphereTagCategoriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectVSphereTagCategoriesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectVSphereTagsForTagCategories lists tags for a tag category from v sphere datacenter
*/
func (a *Client) ListProjectVSphereTagsForTagCategories(params *ListProjectVSphereTagsForTagCategoriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectVSphereTagsForTagCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectVSphereTagsForTagCategoriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectVSphereTagsForTagCategories",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/vsphere/tagcategories/{tag_category}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectVSphereTagsForTagCategoriesReader{formats: a.formats},
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
	success, ok := result.(*ListProjectVSphereTagsForTagCategoriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectVSphereTagsForTagCategoriesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListVSphereDatastores Lists datastores from vsphere datacenter
*/
func (a *Client) ListVSphereDatastores(params *ListVSphereDatastoresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVSphereDatastoresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVSphereDatastoresParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listVSphereDatastores",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/vsphere/datastores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVSphereDatastoresReader{formats: a.formats},
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
	success, ok := result.(*ListVSphereDatastoresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVSphereDatastoresDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListVSphereFolders Lists folders from vsphere datacenter
*/
func (a *Client) ListVSphereFolders(params *ListVSphereFoldersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVSphereFoldersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVSphereFoldersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listVSphereFolders",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/vsphere/folders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVSphereFoldersReader{formats: a.formats},
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
	success, ok := result.(*ListVSphereFoldersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVSphereFoldersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListVSphereFoldersNoCredentials Lists folders from vsphere datacenter
*/
func (a *Client) ListVSphereFoldersNoCredentials(params *ListVSphereFoldersNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVSphereFoldersNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVSphereFoldersNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listVSphereFoldersNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/vsphere/folders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVSphereFoldersNoCredentialsReader{formats: a.formats},
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
	success, ok := result.(*ListVSphereFoldersNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVSphereFoldersNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListVSphereFoldersNoCredentialsV2 Lists folders from vsphere datacenter
*/
func (a *Client) ListVSphereFoldersNoCredentialsV2(params *ListVSphereFoldersNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVSphereFoldersNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVSphereFoldersNoCredentialsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listVSphereFoldersNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vsphere/folders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVSphereFoldersNoCredentialsV2Reader{formats: a.formats},
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
	success, ok := result.(*ListVSphereFoldersNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVSphereFoldersNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListVSphereNetworks Lists networks from vsphere datacenter
*/
func (a *Client) ListVSphereNetworks(params *ListVSphereNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVSphereNetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVSphereNetworksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listVSphereNetworks",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/vsphere/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVSphereNetworksReader{formats: a.formats},
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
	success, ok := result.(*ListVSphereNetworksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVSphereNetworksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListVSphereNetworksNoCredentials Lists networks from vsphere datacenter
*/
func (a *Client) ListVSphereNetworksNoCredentials(params *ListVSphereNetworksNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVSphereNetworksNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVSphereNetworksNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listVSphereNetworksNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/vsphere/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVSphereNetworksNoCredentialsReader{formats: a.formats},
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
	success, ok := result.(*ListVSphereNetworksNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVSphereNetworksNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListVSphereNetworksNoCredentialsV2 Lists networks from vsphere datacenter
*/
func (a *Client) ListVSphereNetworksNoCredentialsV2(params *ListVSphereNetworksNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVSphereNetworksNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVSphereNetworksNoCredentialsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listVSphereNetworksNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vsphere/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVSphereNetworksNoCredentialsV2Reader{formats: a.formats},
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
	success, ok := result.(*ListVSphereNetworksNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVSphereNetworksNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListVSphereTagCategoriesNoCredentials lists tag categories from v sphere datacenter
*/
func (a *Client) ListVSphereTagCategoriesNoCredentials(params *ListVSphereTagCategoriesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVSphereTagCategoriesNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVSphereTagCategoriesNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listVSphereTagCategoriesNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vsphere/tagcategories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVSphereTagCategoriesNoCredentialsReader{formats: a.formats},
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
	success, ok := result.(*ListVSphereTagCategoriesNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVSphereTagCategoriesNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListVSphereTagsForTagCategoryNoCredentials lists tags for a tag category from v sphere datacenter
*/
func (a *Client) ListVSphereTagsForTagCategoryNoCredentials(params *ListVSphereTagsForTagCategoryNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListVSphereTagsForTagCategoryNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVSphereTagsForTagCategoryNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listVSphereTagsForTagCategoryNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vsphere/tagcategories/{tag_category}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVSphereTagsForTagCategoryNoCredentialsReader{formats: a.formats},
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
	success, ok := result.(*ListVSphereTagsForTagCategoryNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVSphereTagsForTagCategoryNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
