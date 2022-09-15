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

// GetExternalClusterReader is a Reader for the GetExternalCluster structure.
type GetExternalClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetExternalClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetExternalClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetExternalClusterOK creates a GetExternalClusterOK with default headers values
func NewGetExternalClusterOK() *GetExternalClusterOK {
	return &GetExternalClusterOK{}
}

/*
GetExternalClusterOK describes a response with status code 200, with default header values.

ExternalCluster
*/
type GetExternalClusterOK struct {
	Payload *models.ExternalCluster
}

// IsSuccess returns true when this get external cluster o k response has a 2xx status code
func (o *GetExternalClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get external cluster o k response has a 3xx status code
func (o *GetExternalClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get external cluster o k response has a 4xx status code
func (o *GetExternalClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get external cluster o k response has a 5xx status code
func (o *GetExternalClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get external cluster o k response a status code equal to that given
func (o *GetExternalClusterOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetExternalClusterOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] getExternalClusterOK  %+v", 200, o.Payload)
}

func (o *GetExternalClusterOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] getExternalClusterOK  %+v", 200, o.Payload)
}

func (o *GetExternalClusterOK) GetPayload() *models.ExternalCluster {
	return o.Payload
}

func (o *GetExternalClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalClusterUnauthorized creates a GetExternalClusterUnauthorized with default headers values
func NewGetExternalClusterUnauthorized() *GetExternalClusterUnauthorized {
	return &GetExternalClusterUnauthorized{}
}

/*
GetExternalClusterUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetExternalClusterUnauthorized struct {
}

// IsSuccess returns true when this get external cluster unauthorized response has a 2xx status code
func (o *GetExternalClusterUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get external cluster unauthorized response has a 3xx status code
func (o *GetExternalClusterUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get external cluster unauthorized response has a 4xx status code
func (o *GetExternalClusterUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get external cluster unauthorized response has a 5xx status code
func (o *GetExternalClusterUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get external cluster unauthorized response a status code equal to that given
func (o *GetExternalClusterUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetExternalClusterUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] getExternalClusterUnauthorized ", 401)
}

func (o *GetExternalClusterUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] getExternalClusterUnauthorized ", 401)
}

func (o *GetExternalClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExternalClusterForbidden creates a GetExternalClusterForbidden with default headers values
func NewGetExternalClusterForbidden() *GetExternalClusterForbidden {
	return &GetExternalClusterForbidden{}
}

/*
GetExternalClusterForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetExternalClusterForbidden struct {
}

// IsSuccess returns true when this get external cluster forbidden response has a 2xx status code
func (o *GetExternalClusterForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get external cluster forbidden response has a 3xx status code
func (o *GetExternalClusterForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get external cluster forbidden response has a 4xx status code
func (o *GetExternalClusterForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get external cluster forbidden response has a 5xx status code
func (o *GetExternalClusterForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get external cluster forbidden response a status code equal to that given
func (o *GetExternalClusterForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetExternalClusterForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] getExternalClusterForbidden ", 403)
}

func (o *GetExternalClusterForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] getExternalClusterForbidden ", 403)
}

func (o *GetExternalClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExternalClusterDefault creates a GetExternalClusterDefault with default headers values
func NewGetExternalClusterDefault(code int) *GetExternalClusterDefault {
	return &GetExternalClusterDefault{
		_statusCode: code,
	}
}

/*
GetExternalClusterDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetExternalClusterDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get external cluster default response
func (o *GetExternalClusterDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get external cluster default response has a 2xx status code
func (o *GetExternalClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get external cluster default response has a 3xx status code
func (o *GetExternalClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get external cluster default response has a 4xx status code
func (o *GetExternalClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get external cluster default response has a 5xx status code
func (o *GetExternalClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get external cluster default response a status code equal to that given
func (o *GetExternalClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetExternalClusterDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] getExternalCluster default  %+v", o._statusCode, o.Payload)
}

func (o *GetExternalClusterDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] getExternalCluster default  %+v", o._statusCode, o.Payload)
}

func (o *GetExternalClusterDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetExternalClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
