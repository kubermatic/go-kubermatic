// Code generated by go-swagger; DO NOT EDIT.

package rulegroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new rulegroup API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for rulegroup API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAdminRuleGroup(params *CreateAdminRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAdminRuleGroupCreated, error)

	CreateRuleGroup(params *CreateRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRuleGroupCreated, error)

	DeleteAdminRuleGroup(params *DeleteAdminRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAdminRuleGroupOK, error)

	DeleteRuleGroup(params *DeleteRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRuleGroupOK, error)

	GetAdminRuleGroup(params *GetAdminRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAdminRuleGroupOK, error)

	GetRuleGroup(params *GetRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRuleGroupOK, error)

	ListAdminRuleGroups(params *ListAdminRuleGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAdminRuleGroupsOK, error)

	ListRuleGroups(params *ListRuleGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListRuleGroupsOK, error)

	UpdateAdminRuleGroup(params *UpdateAdminRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAdminRuleGroupOK, error)

	UpdateRuleGroup(params *UpdateRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRuleGroupOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateAdminRuleGroup Creates a rule group that will belong to the given Seed
*/
func (a *Client) CreateAdminRuleGroup(params *CreateAdminRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAdminRuleGroupCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAdminRuleGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createAdminRuleGroup",
		Method:             "POST",
		PathPattern:        "/api/v2/seeds/{seed_name}/rulegroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAdminRuleGroupReader{formats: a.formats},
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
	success, ok := result.(*CreateAdminRuleGroupCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateAdminRuleGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CreateRuleGroup Creates a rule group that will belong to the given cluster
*/
func (a *Client) CreateRuleGroup(params *CreateRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRuleGroupCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRuleGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createRuleGroup",
		Method:             "POST",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateRuleGroupReader{formats: a.formats},
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
	success, ok := result.(*CreateRuleGroupCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateRuleGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteAdminRuleGroup deletes the given rule group that belongs to the seed
*/
func (a *Client) DeleteAdminRuleGroup(params *DeleteAdminRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAdminRuleGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAdminRuleGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAdminRuleGroup",
		Method:             "DELETE",
		PathPattern:        "/api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAdminRuleGroupReader{formats: a.formats},
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
	success, ok := result.(*DeleteAdminRuleGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteAdminRuleGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteRuleGroup deletes the given rule group that belongs to the cluster
*/
func (a *Client) DeleteRuleGroup(params *DeleteRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRuleGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRuleGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteRuleGroup",
		Method:             "DELETE",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRuleGroupReader{formats: a.formats},
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
	success, ok := result.(*DeleteRuleGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteRuleGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetAdminRuleGroup gets a specified rule group for a given seed
*/
func (a *Client) GetAdminRuleGroup(params *GetAdminRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAdminRuleGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdminRuleGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAdminRuleGroup",
		Method:             "GET",
		PathPattern:        "/api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAdminRuleGroupReader{formats: a.formats},
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
	success, ok := result.(*GetAdminRuleGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAdminRuleGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetRuleGroup gets a specified rule group for the given cluster
*/
func (a *Client) GetRuleGroup(params *GetRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRuleGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRuleGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRuleGroup",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRuleGroupReader{formats: a.formats},
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
	success, ok := result.(*GetRuleGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRuleGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListAdminRuleGroups lists rule groups that belong to a given seed
*/
func (a *Client) ListAdminRuleGroups(params *ListAdminRuleGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAdminRuleGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAdminRuleGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAdminRuleGroups",
		Method:             "GET",
		PathPattern:        "/api/v2/seeds/{seed_name}/rulegroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAdminRuleGroupsReader{formats: a.formats},
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
	success, ok := result.(*ListAdminRuleGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAdminRuleGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListRuleGroups Lists rule groups that belong to the given cluster
*/
func (a *Client) ListRuleGroups(params *ListRuleGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListRuleGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRuleGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listRuleGroups",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListRuleGroupsReader{formats: a.formats},
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
	success, ok := result.(*ListRuleGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListRuleGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateAdminRuleGroup updates the specified rule group for the given seed
*/
func (a *Client) UpdateAdminRuleGroup(params *UpdateAdminRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAdminRuleGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAdminRuleGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateAdminRuleGroup",
		Method:             "PUT",
		PathPattern:        "/api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAdminRuleGroupReader{formats: a.formats},
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
	success, ok := result.(*UpdateAdminRuleGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateAdminRuleGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateRuleGroup updates the specified rule group for the given cluster
*/
func (a *Client) UpdateRuleGroup(params *UpdateRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRuleGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRuleGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateRuleGroup",
		Method:             "PUT",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRuleGroupReader{formats: a.formats},
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
	success, ok := result.(*UpdateRuleGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateRuleGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
