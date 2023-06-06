// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new azure API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for azure API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ListAzureAvailabilityZones(params *ListAzureAvailabilityZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureAvailabilityZonesOK, error)

	ListAzureAvailabilityZonesNoCredentials(params *ListAzureAvailabilityZonesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureAvailabilityZonesNoCredentialsOK, error)

	ListAzureAvailabilityZonesNoCredentialsV2(params *ListAzureAvailabilityZonesNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureAvailabilityZonesNoCredentialsV2OK, error)

	ListAzureResourceGroups(params *ListAzureResourceGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureResourceGroupsOK, error)

	ListAzureRouteTables(params *ListAzureRouteTablesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureRouteTablesOK, error)

	ListAzureSecurityGroups(params *ListAzureSecurityGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureSecurityGroupsOK, error)

	ListAzureSizes(params *ListAzureSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureSizesOK, error)

	ListAzureSizesNoCredentials(params *ListAzureSizesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureSizesNoCredentialsOK, error)

	ListAzureSizesNoCredentialsV2(params *ListAzureSizesNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureSizesNoCredentialsV2OK, error)

	ListAzureSubnets(params *ListAzureSubnetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureSubnetsOK, error)

	ListAzureVnets(params *ListAzureVnetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureVnetsOK, error)

	ListProjectAzureAvailabilityZones(params *ListProjectAzureAvailabilityZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectAzureAvailabilityZonesOK, error)

	ListProjectAzureResourceGroups(params *ListProjectAzureResourceGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectAzureResourceGroupsOK, error)

	ListProjectAzureRouteTables(params *ListProjectAzureRouteTablesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectAzureRouteTablesOK, error)

	ListProjectAzureSecurityGroups(params *ListProjectAzureSecurityGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectAzureSecurityGroupsOK, error)

	ListProjectAzureSizes(params *ListProjectAzureSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectAzureSizesOK, error)

	ListProjectAzureSubnets(params *ListProjectAzureSubnetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectAzureSubnetsOK, error)

	ListProjectAzureVnets(params *ListProjectAzureVnetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectAzureVnetsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ListAzureAvailabilityZones Lists VM availability zones in an Azure region
*/
func (a *Client) ListAzureAvailabilityZones(params *ListAzureAvailabilityZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureAvailabilityZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAzureAvailabilityZonesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAzureAvailabilityZones",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/azure/availabilityzones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAzureAvailabilityZonesReader{formats: a.formats},
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
	success, ok := result.(*ListAzureAvailabilityZonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAzureAvailabilityZonesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListAzureAvailabilityZonesNoCredentials Lists available VM availability zones in an Azure region
*/
func (a *Client) ListAzureAvailabilityZonesNoCredentials(params *ListAzureAvailabilityZonesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureAvailabilityZonesNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAzureAvailabilityZonesNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAzureAvailabilityZonesNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/azure/availabilityzones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAzureAvailabilityZonesNoCredentialsReader{formats: a.formats},
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
	success, ok := result.(*ListAzureAvailabilityZonesNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAzureAvailabilityZonesNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListAzureAvailabilityZonesNoCredentialsV2 Lists available VM availability zones in an Azure region
*/
func (a *Client) ListAzureAvailabilityZonesNoCredentialsV2(params *ListAzureAvailabilityZonesNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureAvailabilityZonesNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAzureAvailabilityZonesNoCredentialsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAzureAvailabilityZonesNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/azure/availabilityzones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAzureAvailabilityZonesNoCredentialsV2Reader{formats: a.formats},
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
	success, ok := result.(*ListAzureAvailabilityZonesNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAzureAvailabilityZonesNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListAzureResourceGroups Lists available VM resource groups
*/
func (a *Client) ListAzureResourceGroups(params *ListAzureResourceGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureResourceGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAzureResourceGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAzureResourceGroups",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/azure/resourcegroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAzureResourceGroupsReader{formats: a.formats},
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
	success, ok := result.(*ListAzureResourceGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAzureResourceGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListAzureRouteTables Lists available VM route tables
*/
func (a *Client) ListAzureRouteTables(params *ListAzureRouteTablesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureRouteTablesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAzureRouteTablesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAzureRouteTables",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/azure/routetables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAzureRouteTablesReader{formats: a.formats},
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
	success, ok := result.(*ListAzureRouteTablesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAzureRouteTablesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListAzureSecurityGroups Lists available VM security groups
*/
func (a *Client) ListAzureSecurityGroups(params *ListAzureSecurityGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureSecurityGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAzureSecurityGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAzureSecurityGroups",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/azure/securitygroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAzureSecurityGroupsReader{formats: a.formats},
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
	success, ok := result.(*ListAzureSecurityGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAzureSecurityGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListAzureSizes Lists available VM sizes in an Azure region
*/
func (a *Client) ListAzureSizes(params *ListAzureSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureSizesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAzureSizesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAzureSizes",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/azure/sizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAzureSizesReader{formats: a.formats},
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
	success, ok := result.(*ListAzureSizesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAzureSizesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListAzureSizesNoCredentials Lists available VM sizes in an Azure region
*/
func (a *Client) ListAzureSizesNoCredentials(params *ListAzureSizesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureSizesNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAzureSizesNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAzureSizesNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/azure/sizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAzureSizesNoCredentialsReader{formats: a.formats},
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
	success, ok := result.(*ListAzureSizesNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAzureSizesNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListAzureSizesNoCredentialsV2 Lists available VM sizes in an Azure region
*/
func (a *Client) ListAzureSizesNoCredentialsV2(params *ListAzureSizesNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureSizesNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAzureSizesNoCredentialsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAzureSizesNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/azure/sizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAzureSizesNoCredentialsV2Reader{formats: a.formats},
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
	success, ok := result.(*ListAzureSizesNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAzureSizesNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListAzureSubnets Lists available VM subnets
*/
func (a *Client) ListAzureSubnets(params *ListAzureSubnetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureSubnetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAzureSubnetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAzureSubnets",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/azure/subnets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAzureSubnetsReader{formats: a.formats},
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
	success, ok := result.(*ListAzureSubnetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAzureSubnetsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListAzureVnets Lists available VM virtual networks
*/
func (a *Client) ListAzureVnets(params *ListAzureVnetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAzureVnetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAzureVnetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAzureVnets",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/azure/vnets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAzureVnetsReader{formats: a.formats},
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
	success, ok := result.(*ListAzureVnetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAzureVnetsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectAzureAvailabilityZones Lists VM availability zones in an Azure region
*/
func (a *Client) ListProjectAzureAvailabilityZones(params *ListProjectAzureAvailabilityZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectAzureAvailabilityZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectAzureAvailabilityZonesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectAzureAvailabilityZones",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/azure/availabilityzones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectAzureAvailabilityZonesReader{formats: a.formats},
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
	success, ok := result.(*ListProjectAzureAvailabilityZonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectAzureAvailabilityZonesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectAzureResourceGroups Lists available VM resource groups
*/
func (a *Client) ListProjectAzureResourceGroups(params *ListProjectAzureResourceGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectAzureResourceGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectAzureResourceGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectAzureResourceGroups",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/azure/resourcegroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectAzureResourceGroupsReader{formats: a.formats},
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
	success, ok := result.(*ListProjectAzureResourceGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectAzureResourceGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectAzureRouteTables Lists available VM route tables
*/
func (a *Client) ListProjectAzureRouteTables(params *ListProjectAzureRouteTablesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectAzureRouteTablesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectAzureRouteTablesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectAzureRouteTables",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/azure/routetables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectAzureRouteTablesReader{formats: a.formats},
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
	success, ok := result.(*ListProjectAzureRouteTablesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectAzureRouteTablesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectAzureSecurityGroups Lists available VM security groups
*/
func (a *Client) ListProjectAzureSecurityGroups(params *ListProjectAzureSecurityGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectAzureSecurityGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectAzureSecurityGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectAzureSecurityGroups",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/azure/securitygroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectAzureSecurityGroupsReader{formats: a.formats},
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
	success, ok := result.(*ListProjectAzureSecurityGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectAzureSecurityGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectAzureSizes Lists available VM sizes in an Azure region
*/
func (a *Client) ListProjectAzureSizes(params *ListProjectAzureSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectAzureSizesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectAzureSizesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectAzureSizes",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/azure/sizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectAzureSizesReader{formats: a.formats},
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
	success, ok := result.(*ListProjectAzureSizesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectAzureSizesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectAzureSubnets Lists available VM subnets
*/
func (a *Client) ListProjectAzureSubnets(params *ListProjectAzureSubnetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectAzureSubnetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectAzureSubnetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectAzureSubnets",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/azure/subnets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectAzureSubnetsReader{formats: a.formats},
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
	success, ok := result.(*ListProjectAzureSubnetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectAzureSubnetsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectAzureVnets Lists available VM virtual networks
*/
func (a *Client) ListProjectAzureVnets(params *ListProjectAzureVnetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectAzureVnetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectAzureVnetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectAzureVnets",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/azure/vnets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectAzureVnetsReader{formats: a.formats},
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
	success, ok := result.(*ListProjectAzureVnetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectAzureVnetsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
