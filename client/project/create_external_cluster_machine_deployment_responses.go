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

// CreateExternalClusterMachineDeploymentReader is a Reader for the CreateExternalClusterMachineDeployment structure.
type CreateExternalClusterMachineDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateExternalClusterMachineDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateExternalClusterMachineDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateExternalClusterMachineDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateExternalClusterMachineDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateExternalClusterMachineDeploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateExternalClusterMachineDeploymentOK creates a CreateExternalClusterMachineDeploymentOK with default headers values
func NewCreateExternalClusterMachineDeploymentOK() *CreateExternalClusterMachineDeploymentOK {
	return &CreateExternalClusterMachineDeploymentOK{}
}

/* CreateExternalClusterMachineDeploymentOK describes a response with status code 200, with default header values.

ExternalClusterMachineDeployment
*/
type CreateExternalClusterMachineDeploymentOK struct {
	Payload *models.ExternalClusterMachineDeployment
}

func (o *CreateExternalClusterMachineDeploymentOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments][%d] createExternalClusterMachineDeploymentOK  %+v", 200, o.Payload)
}
func (o *CreateExternalClusterMachineDeploymentOK) GetPayload() *models.ExternalClusterMachineDeployment {
	return o.Payload
}

func (o *CreateExternalClusterMachineDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalClusterMachineDeployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateExternalClusterMachineDeploymentUnauthorized creates a CreateExternalClusterMachineDeploymentUnauthorized with default headers values
func NewCreateExternalClusterMachineDeploymentUnauthorized() *CreateExternalClusterMachineDeploymentUnauthorized {
	return &CreateExternalClusterMachineDeploymentUnauthorized{}
}

/* CreateExternalClusterMachineDeploymentUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateExternalClusterMachineDeploymentUnauthorized struct {
}

func (o *CreateExternalClusterMachineDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments][%d] createExternalClusterMachineDeploymentUnauthorized ", 401)
}

func (o *CreateExternalClusterMachineDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateExternalClusterMachineDeploymentForbidden creates a CreateExternalClusterMachineDeploymentForbidden with default headers values
func NewCreateExternalClusterMachineDeploymentForbidden() *CreateExternalClusterMachineDeploymentForbidden {
	return &CreateExternalClusterMachineDeploymentForbidden{}
}

/* CreateExternalClusterMachineDeploymentForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateExternalClusterMachineDeploymentForbidden struct {
}

func (o *CreateExternalClusterMachineDeploymentForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments][%d] createExternalClusterMachineDeploymentForbidden ", 403)
}

func (o *CreateExternalClusterMachineDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateExternalClusterMachineDeploymentDefault creates a CreateExternalClusterMachineDeploymentDefault with default headers values
func NewCreateExternalClusterMachineDeploymentDefault(code int) *CreateExternalClusterMachineDeploymentDefault {
	return &CreateExternalClusterMachineDeploymentDefault{
		_statusCode: code,
	}
}

/* CreateExternalClusterMachineDeploymentDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateExternalClusterMachineDeploymentDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create external cluster machine deployment default response
func (o *CreateExternalClusterMachineDeploymentDefault) Code() int {
	return o._statusCode
}

func (o *CreateExternalClusterMachineDeploymentDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments][%d] createExternalClusterMachineDeployment default  %+v", o._statusCode, o.Payload)
}
func (o *CreateExternalClusterMachineDeploymentDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateExternalClusterMachineDeploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}