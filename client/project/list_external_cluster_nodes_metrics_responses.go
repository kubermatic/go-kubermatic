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

// ListExternalClusterNodesMetricsReader is a Reader for the ListExternalClusterNodesMetrics structure.
type ListExternalClusterNodesMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListExternalClusterNodesMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListExternalClusterNodesMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListExternalClusterNodesMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListExternalClusterNodesMetricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListExternalClusterNodesMetricsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListExternalClusterNodesMetricsOK creates a ListExternalClusterNodesMetricsOK with default headers values
func NewListExternalClusterNodesMetricsOK() *ListExternalClusterNodesMetricsOK {
	return &ListExternalClusterNodesMetricsOK{}
}

/* ListExternalClusterNodesMetricsOK describes a response with status code 200, with default header values.

NodeMetric
*/
type ListExternalClusterNodesMetricsOK struct {
	Payload []*models.NodeMetric
}

func (o *ListExternalClusterNodesMetricsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodesmetrics][%d] listExternalClusterNodesMetricsOK  %+v", 200, o.Payload)
}
func (o *ListExternalClusterNodesMetricsOK) GetPayload() []*models.NodeMetric {
	return o.Payload
}

func (o *ListExternalClusterNodesMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListExternalClusterNodesMetricsUnauthorized creates a ListExternalClusterNodesMetricsUnauthorized with default headers values
func NewListExternalClusterNodesMetricsUnauthorized() *ListExternalClusterNodesMetricsUnauthorized {
	return &ListExternalClusterNodesMetricsUnauthorized{}
}

/* ListExternalClusterNodesMetricsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListExternalClusterNodesMetricsUnauthorized struct {
}

func (o *ListExternalClusterNodesMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodesmetrics][%d] listExternalClusterNodesMetricsUnauthorized ", 401)
}

func (o *ListExternalClusterNodesMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListExternalClusterNodesMetricsForbidden creates a ListExternalClusterNodesMetricsForbidden with default headers values
func NewListExternalClusterNodesMetricsForbidden() *ListExternalClusterNodesMetricsForbidden {
	return &ListExternalClusterNodesMetricsForbidden{}
}

/* ListExternalClusterNodesMetricsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListExternalClusterNodesMetricsForbidden struct {
}

func (o *ListExternalClusterNodesMetricsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodesmetrics][%d] listExternalClusterNodesMetricsForbidden ", 403)
}

func (o *ListExternalClusterNodesMetricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListExternalClusterNodesMetricsDefault creates a ListExternalClusterNodesMetricsDefault with default headers values
func NewListExternalClusterNodesMetricsDefault(code int) *ListExternalClusterNodesMetricsDefault {
	return &ListExternalClusterNodesMetricsDefault{
		_statusCode: code,
	}
}

/* ListExternalClusterNodesMetricsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListExternalClusterNodesMetricsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list external cluster nodes metrics default response
func (o *ListExternalClusterNodesMetricsDefault) Code() int {
	return o._statusCode
}

func (o *ListExternalClusterNodesMetricsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodesmetrics][%d] listExternalClusterNodesMetrics default  %+v", o._statusCode, o.Payload)
}
func (o *ListExternalClusterNodesMetricsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListExternalClusterNodesMetricsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
