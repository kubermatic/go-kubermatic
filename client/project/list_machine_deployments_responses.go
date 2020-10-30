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

// ListMachineDeploymentsReader is a Reader for the ListMachineDeployments structure.
type ListMachineDeploymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListMachineDeploymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListMachineDeploymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListMachineDeploymentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListMachineDeploymentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListMachineDeploymentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListMachineDeploymentsOK creates a ListMachineDeploymentsOK with default headers values
func NewListMachineDeploymentsOK() *ListMachineDeploymentsOK {
	return &ListMachineDeploymentsOK{}
}

/*ListMachineDeploymentsOK handles this case with default header values.

NodeDeployment
*/
type ListMachineDeploymentsOK struct {
	Payload []*models.NodeDeployment
}

func (o *ListMachineDeploymentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments][%d] listMachineDeploymentsOK  %+v", 200, o.Payload)
}

func (o *ListMachineDeploymentsOK) GetPayload() []*models.NodeDeployment {
	return o.Payload
}

func (o *ListMachineDeploymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMachineDeploymentsUnauthorized creates a ListMachineDeploymentsUnauthorized with default headers values
func NewListMachineDeploymentsUnauthorized() *ListMachineDeploymentsUnauthorized {
	return &ListMachineDeploymentsUnauthorized{}
}

/*ListMachineDeploymentsUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type ListMachineDeploymentsUnauthorized struct {
}

func (o *ListMachineDeploymentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments][%d] listMachineDeploymentsUnauthorized ", 401)
}

func (o *ListMachineDeploymentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListMachineDeploymentsForbidden creates a ListMachineDeploymentsForbidden with default headers values
func NewListMachineDeploymentsForbidden() *ListMachineDeploymentsForbidden {
	return &ListMachineDeploymentsForbidden{}
}

/*ListMachineDeploymentsForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type ListMachineDeploymentsForbidden struct {
}

func (o *ListMachineDeploymentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments][%d] listMachineDeploymentsForbidden ", 403)
}

func (o *ListMachineDeploymentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListMachineDeploymentsDefault creates a ListMachineDeploymentsDefault with default headers values
func NewListMachineDeploymentsDefault(code int) *ListMachineDeploymentsDefault {
	return &ListMachineDeploymentsDefault{
		_statusCode: code,
	}
}

/*ListMachineDeploymentsDefault handles this case with default header values.

errorResponse
*/
type ListMachineDeploymentsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list machine deployments default response
func (o *ListMachineDeploymentsDefault) Code() int {
	return o._statusCode
}

func (o *ListMachineDeploymentsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments][%d] listMachineDeployments default  %+v", o._statusCode, o.Payload)
}

func (o *ListMachineDeploymentsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListMachineDeploymentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
