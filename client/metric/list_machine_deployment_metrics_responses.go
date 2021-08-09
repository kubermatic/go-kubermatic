// Code generated by go-swagger; DO NOT EDIT.

package metric

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListMachineDeploymentMetricsReader is a Reader for the ListMachineDeploymentMetrics structure.
type ListMachineDeploymentMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListMachineDeploymentMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListMachineDeploymentMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListMachineDeploymentMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListMachineDeploymentMetricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListMachineDeploymentMetricsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListMachineDeploymentMetricsOK creates a ListMachineDeploymentMetricsOK with default headers values
func NewListMachineDeploymentMetricsOK() *ListMachineDeploymentMetricsOK {
	return &ListMachineDeploymentMetricsOK{}
}

/* ListMachineDeploymentMetricsOK describes a response with status code 200, with default header values.

NodeMetric
*/
type ListMachineDeploymentMetricsOK struct {
	Payload []*models.NodeMetric
}

func (o *ListMachineDeploymentMetricsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/metrics][%d] listMachineDeploymentMetricsOK  %+v", 200, o.Payload)
}
func (o *ListMachineDeploymentMetricsOK) GetPayload() []*models.NodeMetric {
	return o.Payload
}

func (o *ListMachineDeploymentMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMachineDeploymentMetricsUnauthorized creates a ListMachineDeploymentMetricsUnauthorized with default headers values
func NewListMachineDeploymentMetricsUnauthorized() *ListMachineDeploymentMetricsUnauthorized {
	return &ListMachineDeploymentMetricsUnauthorized{}
}

/* ListMachineDeploymentMetricsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListMachineDeploymentMetricsUnauthorized struct {
}

func (o *ListMachineDeploymentMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/metrics][%d] listMachineDeploymentMetricsUnauthorized ", 401)
}

func (o *ListMachineDeploymentMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListMachineDeploymentMetricsForbidden creates a ListMachineDeploymentMetricsForbidden with default headers values
func NewListMachineDeploymentMetricsForbidden() *ListMachineDeploymentMetricsForbidden {
	return &ListMachineDeploymentMetricsForbidden{}
}

/* ListMachineDeploymentMetricsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListMachineDeploymentMetricsForbidden struct {
}

func (o *ListMachineDeploymentMetricsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/metrics][%d] listMachineDeploymentMetricsForbidden ", 403)
}

func (o *ListMachineDeploymentMetricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListMachineDeploymentMetricsDefault creates a ListMachineDeploymentMetricsDefault with default headers values
func NewListMachineDeploymentMetricsDefault(code int) *ListMachineDeploymentMetricsDefault {
	return &ListMachineDeploymentMetricsDefault{
		_statusCode: code,
	}
}

/* ListMachineDeploymentMetricsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListMachineDeploymentMetricsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list machine deployment metrics default response
func (o *ListMachineDeploymentMetricsDefault) Code() int {
	return o._statusCode
}

func (o *ListMachineDeploymentMetricsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/metrics][%d] listMachineDeploymentMetrics default  %+v", o._statusCode, o.Payload)
}
func (o *ListMachineDeploymentMetricsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListMachineDeploymentMetricsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
