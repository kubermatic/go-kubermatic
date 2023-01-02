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

// GetClusterServiceAccountKubeconfigReader is a Reader for the GetClusterServiceAccountKubeconfig structure.
type GetClusterServiceAccountKubeconfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterServiceAccountKubeconfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterServiceAccountKubeconfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClusterServiceAccountKubeconfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetClusterServiceAccountKubeconfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClusterServiceAccountKubeconfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterServiceAccountKubeconfigOK creates a GetClusterServiceAccountKubeconfigOK with default headers values
func NewGetClusterServiceAccountKubeconfigOK() *GetClusterServiceAccountKubeconfigOK {
	return &GetClusterServiceAccountKubeconfigOK{}
}

/*
GetClusterServiceAccountKubeconfigOK describes a response with status code 200, with default header values.

Kubeconfig is a clusters kubeconfig
*/
type GetClusterServiceAccountKubeconfigOK struct {
	Payload []uint8
}

// IsSuccess returns true when this get cluster service account kubeconfig o k response has a 2xx status code
func (o *GetClusterServiceAccountKubeconfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster service account kubeconfig o k response has a 3xx status code
func (o *GetClusterServiceAccountKubeconfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster service account kubeconfig o k response has a 4xx status code
func (o *GetClusterServiceAccountKubeconfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster service account kubeconfig o k response has a 5xx status code
func (o *GetClusterServiceAccountKubeconfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster service account kubeconfig o k response a status code equal to that given
func (o *GetClusterServiceAccountKubeconfigOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetClusterServiceAccountKubeconfigOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount/{namespace}/{service_account_id}/kubeconfig][%d] getClusterServiceAccountKubeconfigOK  %+v", 200, o.Payload)
}

func (o *GetClusterServiceAccountKubeconfigOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount/{namespace}/{service_account_id}/kubeconfig][%d] getClusterServiceAccountKubeconfigOK  %+v", 200, o.Payload)
}

func (o *GetClusterServiceAccountKubeconfigOK) GetPayload() []uint8 {
	return o.Payload
}

func (o *GetClusterServiceAccountKubeconfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterServiceAccountKubeconfigUnauthorized creates a GetClusterServiceAccountKubeconfigUnauthorized with default headers values
func NewGetClusterServiceAccountKubeconfigUnauthorized() *GetClusterServiceAccountKubeconfigUnauthorized {
	return &GetClusterServiceAccountKubeconfigUnauthorized{}
}

/*
GetClusterServiceAccountKubeconfigUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetClusterServiceAccountKubeconfigUnauthorized struct {
}

// IsSuccess returns true when this get cluster service account kubeconfig unauthorized response has a 2xx status code
func (o *GetClusterServiceAccountKubeconfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster service account kubeconfig unauthorized response has a 3xx status code
func (o *GetClusterServiceAccountKubeconfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster service account kubeconfig unauthorized response has a 4xx status code
func (o *GetClusterServiceAccountKubeconfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster service account kubeconfig unauthorized response has a 5xx status code
func (o *GetClusterServiceAccountKubeconfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster service account kubeconfig unauthorized response a status code equal to that given
func (o *GetClusterServiceAccountKubeconfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetClusterServiceAccountKubeconfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount/{namespace}/{service_account_id}/kubeconfig][%d] getClusterServiceAccountKubeconfigUnauthorized ", 401)
}

func (o *GetClusterServiceAccountKubeconfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount/{namespace}/{service_account_id}/kubeconfig][%d] getClusterServiceAccountKubeconfigUnauthorized ", 401)
}

func (o *GetClusterServiceAccountKubeconfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterServiceAccountKubeconfigForbidden creates a GetClusterServiceAccountKubeconfigForbidden with default headers values
func NewGetClusterServiceAccountKubeconfigForbidden() *GetClusterServiceAccountKubeconfigForbidden {
	return &GetClusterServiceAccountKubeconfigForbidden{}
}

/*
GetClusterServiceAccountKubeconfigForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetClusterServiceAccountKubeconfigForbidden struct {
}

// IsSuccess returns true when this get cluster service account kubeconfig forbidden response has a 2xx status code
func (o *GetClusterServiceAccountKubeconfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster service account kubeconfig forbidden response has a 3xx status code
func (o *GetClusterServiceAccountKubeconfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster service account kubeconfig forbidden response has a 4xx status code
func (o *GetClusterServiceAccountKubeconfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster service account kubeconfig forbidden response has a 5xx status code
func (o *GetClusterServiceAccountKubeconfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster service account kubeconfig forbidden response a status code equal to that given
func (o *GetClusterServiceAccountKubeconfigForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetClusterServiceAccountKubeconfigForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount/{namespace}/{service_account_id}/kubeconfig][%d] getClusterServiceAccountKubeconfigForbidden ", 403)
}

func (o *GetClusterServiceAccountKubeconfigForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount/{namespace}/{service_account_id}/kubeconfig][%d] getClusterServiceAccountKubeconfigForbidden ", 403)
}

func (o *GetClusterServiceAccountKubeconfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterServiceAccountKubeconfigDefault creates a GetClusterServiceAccountKubeconfigDefault with default headers values
func NewGetClusterServiceAccountKubeconfigDefault(code int) *GetClusterServiceAccountKubeconfigDefault {
	return &GetClusterServiceAccountKubeconfigDefault{
		_statusCode: code,
	}
}

/*
GetClusterServiceAccountKubeconfigDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetClusterServiceAccountKubeconfigDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster service account kubeconfig default response
func (o *GetClusterServiceAccountKubeconfigDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get cluster service account kubeconfig default response has a 2xx status code
func (o *GetClusterServiceAccountKubeconfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cluster service account kubeconfig default response has a 3xx status code
func (o *GetClusterServiceAccountKubeconfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cluster service account kubeconfig default response has a 4xx status code
func (o *GetClusterServiceAccountKubeconfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cluster service account kubeconfig default response has a 5xx status code
func (o *GetClusterServiceAccountKubeconfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cluster service account kubeconfig default response a status code equal to that given
func (o *GetClusterServiceAccountKubeconfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetClusterServiceAccountKubeconfigDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount/{namespace}/{service_account_id}/kubeconfig][%d] getClusterServiceAccountKubeconfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterServiceAccountKubeconfigDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount/{namespace}/{service_account_id}/kubeconfig][%d] getClusterServiceAccountKubeconfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterServiceAccountKubeconfigDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterServiceAccountKubeconfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}