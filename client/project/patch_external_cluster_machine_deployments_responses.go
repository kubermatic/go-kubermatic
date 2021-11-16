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

// PatchExternalClusterMachineDeploymentsReader is a Reader for the PatchExternalClusterMachineDeployments structure.
type PatchExternalClusterMachineDeploymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchExternalClusterMachineDeploymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchExternalClusterMachineDeploymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchExternalClusterMachineDeploymentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchExternalClusterMachineDeploymentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchExternalClusterMachineDeploymentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchExternalClusterMachineDeploymentsOK creates a PatchExternalClusterMachineDeploymentsOK with default headers values
func NewPatchExternalClusterMachineDeploymentsOK() *PatchExternalClusterMachineDeploymentsOK {
	return &PatchExternalClusterMachineDeploymentsOK{}
}

/* PatchExternalClusterMachineDeploymentsOK describes a response with status code 200, with default header values.

ExternalClusterMachineDeployment
*/
type PatchExternalClusterMachineDeploymentsOK struct {
	Payload *models.ExternalClusterMachineDeployment
}

func (o *PatchExternalClusterMachineDeploymentsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] patchExternalClusterMachineDeploymentsOK  %+v", 200, o.Payload)
}
func (o *PatchExternalClusterMachineDeploymentsOK) GetPayload() *models.ExternalClusterMachineDeployment {
	return o.Payload
}

func (o *PatchExternalClusterMachineDeploymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalClusterMachineDeployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchExternalClusterMachineDeploymentsUnauthorized creates a PatchExternalClusterMachineDeploymentsUnauthorized with default headers values
func NewPatchExternalClusterMachineDeploymentsUnauthorized() *PatchExternalClusterMachineDeploymentsUnauthorized {
	return &PatchExternalClusterMachineDeploymentsUnauthorized{}
}

/* PatchExternalClusterMachineDeploymentsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchExternalClusterMachineDeploymentsUnauthorized struct {
}

func (o *PatchExternalClusterMachineDeploymentsUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] patchExternalClusterMachineDeploymentsUnauthorized ", 401)
}

func (o *PatchExternalClusterMachineDeploymentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchExternalClusterMachineDeploymentsForbidden creates a PatchExternalClusterMachineDeploymentsForbidden with default headers values
func NewPatchExternalClusterMachineDeploymentsForbidden() *PatchExternalClusterMachineDeploymentsForbidden {
	return &PatchExternalClusterMachineDeploymentsForbidden{}
}

/* PatchExternalClusterMachineDeploymentsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchExternalClusterMachineDeploymentsForbidden struct {
}

func (o *PatchExternalClusterMachineDeploymentsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] patchExternalClusterMachineDeploymentsForbidden ", 403)
}

func (o *PatchExternalClusterMachineDeploymentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchExternalClusterMachineDeploymentsDefault creates a PatchExternalClusterMachineDeploymentsDefault with default headers values
func NewPatchExternalClusterMachineDeploymentsDefault(code int) *PatchExternalClusterMachineDeploymentsDefault {
	return &PatchExternalClusterMachineDeploymentsDefault{
		_statusCode: code,
	}
}

/* PatchExternalClusterMachineDeploymentsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PatchExternalClusterMachineDeploymentsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch external cluster machine deployments default response
func (o *PatchExternalClusterMachineDeploymentsDefault) Code() int {
	return o._statusCode
}

func (o *PatchExternalClusterMachineDeploymentsDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] patchExternalClusterMachineDeployments default  %+v", o._statusCode, o.Payload)
}
func (o *PatchExternalClusterMachineDeploymentsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchExternalClusterMachineDeploymentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
