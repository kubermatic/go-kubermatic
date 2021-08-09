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

// GetClusterMetricsV2Reader is a Reader for the GetClusterMetricsV2 structure.
type GetClusterMetricsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterMetricsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterMetricsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClusterMetricsV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetClusterMetricsV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClusterMetricsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterMetricsV2OK creates a GetClusterMetricsV2OK with default headers values
func NewGetClusterMetricsV2OK() *GetClusterMetricsV2OK {
	return &GetClusterMetricsV2OK{}
}

/* GetClusterMetricsV2OK describes a response with status code 200, with default header values.

ClusterMetrics
*/
type GetClusterMetricsV2OK struct {
	Payload *models.ClusterMetrics
}

func (o *GetClusterMetricsV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/metrics][%d] getClusterMetricsV2OK  %+v", 200, o.Payload)
}
func (o *GetClusterMetricsV2OK) GetPayload() *models.ClusterMetrics {
	return o.Payload
}

func (o *GetClusterMetricsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterMetrics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterMetricsV2Unauthorized creates a GetClusterMetricsV2Unauthorized with default headers values
func NewGetClusterMetricsV2Unauthorized() *GetClusterMetricsV2Unauthorized {
	return &GetClusterMetricsV2Unauthorized{}
}

/* GetClusterMetricsV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetClusterMetricsV2Unauthorized struct {
}

func (o *GetClusterMetricsV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/metrics][%d] getClusterMetricsV2Unauthorized ", 401)
}

func (o *GetClusterMetricsV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterMetricsV2Forbidden creates a GetClusterMetricsV2Forbidden with default headers values
func NewGetClusterMetricsV2Forbidden() *GetClusterMetricsV2Forbidden {
	return &GetClusterMetricsV2Forbidden{}
}

/* GetClusterMetricsV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetClusterMetricsV2Forbidden struct {
}

func (o *GetClusterMetricsV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/metrics][%d] getClusterMetricsV2Forbidden ", 403)
}

func (o *GetClusterMetricsV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterMetricsV2Default creates a GetClusterMetricsV2Default with default headers values
func NewGetClusterMetricsV2Default(code int) *GetClusterMetricsV2Default {
	return &GetClusterMetricsV2Default{
		_statusCode: code,
	}
}

/* GetClusterMetricsV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type GetClusterMetricsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster metrics v2 default response
func (o *GetClusterMetricsV2Default) Code() int {
	return o._statusCode
}

func (o *GetClusterMetricsV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/metrics][%d] getClusterMetricsV2 default  %+v", o._statusCode, o.Payload)
}
func (o *GetClusterMetricsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterMetricsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
