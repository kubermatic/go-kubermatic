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

// DeleteExternalClusterMachineDeploymentReader is a Reader for the DeleteExternalClusterMachineDeployment structure.
type DeleteExternalClusterMachineDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteExternalClusterMachineDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteExternalClusterMachineDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteExternalClusterMachineDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteExternalClusterMachineDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteExternalClusterMachineDeploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteExternalClusterMachineDeploymentOK creates a DeleteExternalClusterMachineDeploymentOK with default headers values
func NewDeleteExternalClusterMachineDeploymentOK() *DeleteExternalClusterMachineDeploymentOK {
	return &DeleteExternalClusterMachineDeploymentOK{}
}

/*
DeleteExternalClusterMachineDeploymentOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteExternalClusterMachineDeploymentOK struct {
}

// IsSuccess returns true when this delete external cluster machine deployment o k response has a 2xx status code
func (o *DeleteExternalClusterMachineDeploymentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete external cluster machine deployment o k response has a 3xx status code
func (o *DeleteExternalClusterMachineDeploymentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete external cluster machine deployment o k response has a 4xx status code
func (o *DeleteExternalClusterMachineDeploymentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete external cluster machine deployment o k response has a 5xx status code
func (o *DeleteExternalClusterMachineDeploymentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete external cluster machine deployment o k response a status code equal to that given
func (o *DeleteExternalClusterMachineDeploymentOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteExternalClusterMachineDeploymentOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteExternalClusterMachineDeploymentOK ", 200)
}

func (o *DeleteExternalClusterMachineDeploymentOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteExternalClusterMachineDeploymentOK ", 200)
}

func (o *DeleteExternalClusterMachineDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteExternalClusterMachineDeploymentUnauthorized creates a DeleteExternalClusterMachineDeploymentUnauthorized with default headers values
func NewDeleteExternalClusterMachineDeploymentUnauthorized() *DeleteExternalClusterMachineDeploymentUnauthorized {
	return &DeleteExternalClusterMachineDeploymentUnauthorized{}
}

/*
DeleteExternalClusterMachineDeploymentUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteExternalClusterMachineDeploymentUnauthorized struct {
}

// IsSuccess returns true when this delete external cluster machine deployment unauthorized response has a 2xx status code
func (o *DeleteExternalClusterMachineDeploymentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete external cluster machine deployment unauthorized response has a 3xx status code
func (o *DeleteExternalClusterMachineDeploymentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete external cluster machine deployment unauthorized response has a 4xx status code
func (o *DeleteExternalClusterMachineDeploymentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete external cluster machine deployment unauthorized response has a 5xx status code
func (o *DeleteExternalClusterMachineDeploymentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete external cluster machine deployment unauthorized response a status code equal to that given
func (o *DeleteExternalClusterMachineDeploymentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteExternalClusterMachineDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteExternalClusterMachineDeploymentUnauthorized ", 401)
}

func (o *DeleteExternalClusterMachineDeploymentUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteExternalClusterMachineDeploymentUnauthorized ", 401)
}

func (o *DeleteExternalClusterMachineDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteExternalClusterMachineDeploymentForbidden creates a DeleteExternalClusterMachineDeploymentForbidden with default headers values
func NewDeleteExternalClusterMachineDeploymentForbidden() *DeleteExternalClusterMachineDeploymentForbidden {
	return &DeleteExternalClusterMachineDeploymentForbidden{}
}

/*
DeleteExternalClusterMachineDeploymentForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteExternalClusterMachineDeploymentForbidden struct {
}

// IsSuccess returns true when this delete external cluster machine deployment forbidden response has a 2xx status code
func (o *DeleteExternalClusterMachineDeploymentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete external cluster machine deployment forbidden response has a 3xx status code
func (o *DeleteExternalClusterMachineDeploymentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete external cluster machine deployment forbidden response has a 4xx status code
func (o *DeleteExternalClusterMachineDeploymentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete external cluster machine deployment forbidden response has a 5xx status code
func (o *DeleteExternalClusterMachineDeploymentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete external cluster machine deployment forbidden response a status code equal to that given
func (o *DeleteExternalClusterMachineDeploymentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteExternalClusterMachineDeploymentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteExternalClusterMachineDeploymentForbidden ", 403)
}

func (o *DeleteExternalClusterMachineDeploymentForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteExternalClusterMachineDeploymentForbidden ", 403)
}

func (o *DeleteExternalClusterMachineDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteExternalClusterMachineDeploymentDefault creates a DeleteExternalClusterMachineDeploymentDefault with default headers values
func NewDeleteExternalClusterMachineDeploymentDefault(code int) *DeleteExternalClusterMachineDeploymentDefault {
	return &DeleteExternalClusterMachineDeploymentDefault{
		_statusCode: code,
	}
}

/*
DeleteExternalClusterMachineDeploymentDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteExternalClusterMachineDeploymentDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete external cluster machine deployment default response
func (o *DeleteExternalClusterMachineDeploymentDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete external cluster machine deployment default response has a 2xx status code
func (o *DeleteExternalClusterMachineDeploymentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete external cluster machine deployment default response has a 3xx status code
func (o *DeleteExternalClusterMachineDeploymentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete external cluster machine deployment default response has a 4xx status code
func (o *DeleteExternalClusterMachineDeploymentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete external cluster machine deployment default response has a 5xx status code
func (o *DeleteExternalClusterMachineDeploymentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete external cluster machine deployment default response a status code equal to that given
func (o *DeleteExternalClusterMachineDeploymentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteExternalClusterMachineDeploymentDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteExternalClusterMachineDeployment default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteExternalClusterMachineDeploymentDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteExternalClusterMachineDeployment default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteExternalClusterMachineDeploymentDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteExternalClusterMachineDeploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
