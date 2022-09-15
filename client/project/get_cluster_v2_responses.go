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

// GetClusterV2Reader is a Reader for the GetClusterV2 structure.
type GetClusterV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClusterV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetClusterV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClusterV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterV2OK creates a GetClusterV2OK with default headers values
func NewGetClusterV2OK() *GetClusterV2OK {
	return &GetClusterV2OK{}
}

/*
GetClusterV2OK describes a response with status code 200, with default header values.

Cluster
*/
type GetClusterV2OK struct {
	Payload *models.Cluster
}

// IsSuccess returns true when this get cluster v2 o k response has a 2xx status code
func (o *GetClusterV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster v2 o k response has a 3xx status code
func (o *GetClusterV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster v2 o k response has a 4xx status code
func (o *GetClusterV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster v2 o k response has a 5xx status code
func (o *GetClusterV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster v2 o k response a status code equal to that given
func (o *GetClusterV2OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetClusterV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}][%d] getClusterV2OK  %+v", 200, o.Payload)
}

func (o *GetClusterV2OK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}][%d] getClusterV2OK  %+v", 200, o.Payload)
}

func (o *GetClusterV2OK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *GetClusterV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterV2Unauthorized creates a GetClusterV2Unauthorized with default headers values
func NewGetClusterV2Unauthorized() *GetClusterV2Unauthorized {
	return &GetClusterV2Unauthorized{}
}

/*
GetClusterV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetClusterV2Unauthorized struct {
}

// IsSuccess returns true when this get cluster v2 unauthorized response has a 2xx status code
func (o *GetClusterV2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster v2 unauthorized response has a 3xx status code
func (o *GetClusterV2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster v2 unauthorized response has a 4xx status code
func (o *GetClusterV2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster v2 unauthorized response has a 5xx status code
func (o *GetClusterV2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster v2 unauthorized response a status code equal to that given
func (o *GetClusterV2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetClusterV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}][%d] getClusterV2Unauthorized ", 401)
}

func (o *GetClusterV2Unauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}][%d] getClusterV2Unauthorized ", 401)
}

func (o *GetClusterV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterV2Forbidden creates a GetClusterV2Forbidden with default headers values
func NewGetClusterV2Forbidden() *GetClusterV2Forbidden {
	return &GetClusterV2Forbidden{}
}

/*
GetClusterV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetClusterV2Forbidden struct {
}

// IsSuccess returns true when this get cluster v2 forbidden response has a 2xx status code
func (o *GetClusterV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster v2 forbidden response has a 3xx status code
func (o *GetClusterV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster v2 forbidden response has a 4xx status code
func (o *GetClusterV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster v2 forbidden response has a 5xx status code
func (o *GetClusterV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster v2 forbidden response a status code equal to that given
func (o *GetClusterV2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetClusterV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}][%d] getClusterV2Forbidden ", 403)
}

func (o *GetClusterV2Forbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}][%d] getClusterV2Forbidden ", 403)
}

func (o *GetClusterV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterV2Default creates a GetClusterV2Default with default headers values
func NewGetClusterV2Default(code int) *GetClusterV2Default {
	return &GetClusterV2Default{
		_statusCode: code,
	}
}

/*
GetClusterV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type GetClusterV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster v2 default response
func (o *GetClusterV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get cluster v2 default response has a 2xx status code
func (o *GetClusterV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cluster v2 default response has a 3xx status code
func (o *GetClusterV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cluster v2 default response has a 4xx status code
func (o *GetClusterV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cluster v2 default response has a 5xx status code
func (o *GetClusterV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cluster v2 default response a status code equal to that given
func (o *GetClusterV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetClusterV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}][%d] getClusterV2 default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterV2Default) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}][%d] getClusterV2 default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
