// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListExternalClusterMachineDeploymentNodesReader is a Reader for the ListExternalClusterMachineDeploymentNodes structure.
type ListExternalClusterMachineDeploymentNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListExternalClusterMachineDeploymentNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListExternalClusterMachineDeploymentNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListExternalClusterMachineDeploymentNodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListExternalClusterMachineDeploymentNodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListExternalClusterMachineDeploymentNodesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListExternalClusterMachineDeploymentNodesOK creates a ListExternalClusterMachineDeploymentNodesOK with default headers values
func NewListExternalClusterMachineDeploymentNodesOK() *ListExternalClusterMachineDeploymentNodesOK {
	return &ListExternalClusterMachineDeploymentNodesOK{}
}

/*
ListExternalClusterMachineDeploymentNodesOK describes a response with status code 200, with default header values.

ExternalClusterNode
*/
type ListExternalClusterMachineDeploymentNodesOK struct {
	Payload []*models.ExternalClusterNode
}

// IsSuccess returns true when this list external cluster machine deployment nodes o k response has a 2xx status code
func (o *ListExternalClusterMachineDeploymentNodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list external cluster machine deployment nodes o k response has a 3xx status code
func (o *ListExternalClusterMachineDeploymentNodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list external cluster machine deployment nodes o k response has a 4xx status code
func (o *ListExternalClusterMachineDeploymentNodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list external cluster machine deployment nodes o k response has a 5xx status code
func (o *ListExternalClusterMachineDeploymentNodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list external cluster machine deployment nodes o k response a status code equal to that given
func (o *ListExternalClusterMachineDeploymentNodesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListExternalClusterMachineDeploymentNodesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listExternalClusterMachineDeploymentNodesOK  %+v", 200, o.Payload)
}

func (o *ListExternalClusterMachineDeploymentNodesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listExternalClusterMachineDeploymentNodesOK  %+v", 200, o.Payload)
}

func (o *ListExternalClusterMachineDeploymentNodesOK) GetPayload() []*models.ExternalClusterNode {
	return o.Payload
}

func (o *ListExternalClusterMachineDeploymentNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListExternalClusterMachineDeploymentNodesUnauthorized creates a ListExternalClusterMachineDeploymentNodesUnauthorized with default headers values
func NewListExternalClusterMachineDeploymentNodesUnauthorized() *ListExternalClusterMachineDeploymentNodesUnauthorized {
	return &ListExternalClusterMachineDeploymentNodesUnauthorized{}
}

/*
ListExternalClusterMachineDeploymentNodesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListExternalClusterMachineDeploymentNodesUnauthorized struct {
}

// IsSuccess returns true when this list external cluster machine deployment nodes unauthorized response has a 2xx status code
func (o *ListExternalClusterMachineDeploymentNodesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list external cluster machine deployment nodes unauthorized response has a 3xx status code
func (o *ListExternalClusterMachineDeploymentNodesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list external cluster machine deployment nodes unauthorized response has a 4xx status code
func (o *ListExternalClusterMachineDeploymentNodesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list external cluster machine deployment nodes unauthorized response has a 5xx status code
func (o *ListExternalClusterMachineDeploymentNodesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list external cluster machine deployment nodes unauthorized response a status code equal to that given
func (o *ListExternalClusterMachineDeploymentNodesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListExternalClusterMachineDeploymentNodesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listExternalClusterMachineDeploymentNodesUnauthorized ", 401)
}

func (o *ListExternalClusterMachineDeploymentNodesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listExternalClusterMachineDeploymentNodesUnauthorized ", 401)
}

func (o *ListExternalClusterMachineDeploymentNodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListExternalClusterMachineDeploymentNodesForbidden creates a ListExternalClusterMachineDeploymentNodesForbidden with default headers values
func NewListExternalClusterMachineDeploymentNodesForbidden() *ListExternalClusterMachineDeploymentNodesForbidden {
	return &ListExternalClusterMachineDeploymentNodesForbidden{}
}

/*
ListExternalClusterMachineDeploymentNodesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListExternalClusterMachineDeploymentNodesForbidden struct {
}

// IsSuccess returns true when this list external cluster machine deployment nodes forbidden response has a 2xx status code
func (o *ListExternalClusterMachineDeploymentNodesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list external cluster machine deployment nodes forbidden response has a 3xx status code
func (o *ListExternalClusterMachineDeploymentNodesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list external cluster machine deployment nodes forbidden response has a 4xx status code
func (o *ListExternalClusterMachineDeploymentNodesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list external cluster machine deployment nodes forbidden response has a 5xx status code
func (o *ListExternalClusterMachineDeploymentNodesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list external cluster machine deployment nodes forbidden response a status code equal to that given
func (o *ListExternalClusterMachineDeploymentNodesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListExternalClusterMachineDeploymentNodesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listExternalClusterMachineDeploymentNodesForbidden ", 403)
}

func (o *ListExternalClusterMachineDeploymentNodesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listExternalClusterMachineDeploymentNodesForbidden ", 403)
}

func (o *ListExternalClusterMachineDeploymentNodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListExternalClusterMachineDeploymentNodesDefault creates a ListExternalClusterMachineDeploymentNodesDefault with default headers values
func NewListExternalClusterMachineDeploymentNodesDefault(code int) *ListExternalClusterMachineDeploymentNodesDefault {
	return &ListExternalClusterMachineDeploymentNodesDefault{
		_statusCode: code,
	}
}

/*
ListExternalClusterMachineDeploymentNodesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListExternalClusterMachineDeploymentNodesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list external cluster machine deployment nodes default response
func (o *ListExternalClusterMachineDeploymentNodesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list external cluster machine deployment nodes default response has a 2xx status code
func (o *ListExternalClusterMachineDeploymentNodesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list external cluster machine deployment nodes default response has a 3xx status code
func (o *ListExternalClusterMachineDeploymentNodesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list external cluster machine deployment nodes default response has a 4xx status code
func (o *ListExternalClusterMachineDeploymentNodesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list external cluster machine deployment nodes default response has a 5xx status code
func (o *ListExternalClusterMachineDeploymentNodesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list external cluster machine deployment nodes default response a status code equal to that given
func (o *ListExternalClusterMachineDeploymentNodesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListExternalClusterMachineDeploymentNodesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listExternalClusterMachineDeploymentNodes default  %+v", o._statusCode, o.Payload)
}

func (o *ListExternalClusterMachineDeploymentNodesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listExternalClusterMachineDeploymentNodes default  %+v", o._statusCode, o.Payload)
}

func (o *ListExternalClusterMachineDeploymentNodesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListExternalClusterMachineDeploymentNodesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
