// Code generated by go-swagger; DO NOT EDIT.

package eks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListEKSSubnetsNoCredentialsReader is a Reader for the ListEKSSubnetsNoCredentials structure.
type ListEKSSubnetsNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEKSSubnetsNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEKSSubnetsNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListEKSSubnetsNoCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListEKSSubnetsNoCredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListEKSSubnetsNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListEKSSubnetsNoCredentialsOK creates a ListEKSSubnetsNoCredentialsOK with default headers values
func NewListEKSSubnetsNoCredentialsOK() *ListEKSSubnetsNoCredentialsOK {
	return &ListEKSSubnetsNoCredentialsOK{}
}

/*
ListEKSSubnetsNoCredentialsOK describes a response with status code 200, with default header values.

EKSSubnetList
*/
type ListEKSSubnetsNoCredentialsOK struct {
	Payload models.EKSSubnetList
}

// IsSuccess returns true when this list e k s subnets no credentials o k response has a 2xx status code
func (o *ListEKSSubnetsNoCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list e k s subnets no credentials o k response has a 3xx status code
func (o *ListEKSSubnetsNoCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s subnets no credentials o k response has a 4xx status code
func (o *ListEKSSubnetsNoCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list e k s subnets no credentials o k response has a 5xx status code
func (o *ListEKSSubnetsNoCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s subnets no credentials o k response a status code equal to that given
func (o *ListEKSSubnetsNoCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListEKSSubnetsNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/subnets][%d] listEKSSubnetsNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListEKSSubnetsNoCredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/subnets][%d] listEKSSubnetsNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListEKSSubnetsNoCredentialsOK) GetPayload() models.EKSSubnetList {
	return o.Payload
}

func (o *ListEKSSubnetsNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEKSSubnetsNoCredentialsUnauthorized creates a ListEKSSubnetsNoCredentialsUnauthorized with default headers values
func NewListEKSSubnetsNoCredentialsUnauthorized() *ListEKSSubnetsNoCredentialsUnauthorized {
	return &ListEKSSubnetsNoCredentialsUnauthorized{}
}

/*
ListEKSSubnetsNoCredentialsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListEKSSubnetsNoCredentialsUnauthorized struct {
}

// IsSuccess returns true when this list e k s subnets no credentials unauthorized response has a 2xx status code
func (o *ListEKSSubnetsNoCredentialsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list e k s subnets no credentials unauthorized response has a 3xx status code
func (o *ListEKSSubnetsNoCredentialsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s subnets no credentials unauthorized response has a 4xx status code
func (o *ListEKSSubnetsNoCredentialsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list e k s subnets no credentials unauthorized response has a 5xx status code
func (o *ListEKSSubnetsNoCredentialsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s subnets no credentials unauthorized response a status code equal to that given
func (o *ListEKSSubnetsNoCredentialsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListEKSSubnetsNoCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/subnets][%d] listEKSSubnetsNoCredentialsUnauthorized ", 401)
}

func (o *ListEKSSubnetsNoCredentialsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/subnets][%d] listEKSSubnetsNoCredentialsUnauthorized ", 401)
}

func (o *ListEKSSubnetsNoCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSSubnetsNoCredentialsForbidden creates a ListEKSSubnetsNoCredentialsForbidden with default headers values
func NewListEKSSubnetsNoCredentialsForbidden() *ListEKSSubnetsNoCredentialsForbidden {
	return &ListEKSSubnetsNoCredentialsForbidden{}
}

/*
ListEKSSubnetsNoCredentialsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListEKSSubnetsNoCredentialsForbidden struct {
}

// IsSuccess returns true when this list e k s subnets no credentials forbidden response has a 2xx status code
func (o *ListEKSSubnetsNoCredentialsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list e k s subnets no credentials forbidden response has a 3xx status code
func (o *ListEKSSubnetsNoCredentialsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s subnets no credentials forbidden response has a 4xx status code
func (o *ListEKSSubnetsNoCredentialsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list e k s subnets no credentials forbidden response has a 5xx status code
func (o *ListEKSSubnetsNoCredentialsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s subnets no credentials forbidden response a status code equal to that given
func (o *ListEKSSubnetsNoCredentialsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListEKSSubnetsNoCredentialsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/subnets][%d] listEKSSubnetsNoCredentialsForbidden ", 403)
}

func (o *ListEKSSubnetsNoCredentialsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/subnets][%d] listEKSSubnetsNoCredentialsForbidden ", 403)
}

func (o *ListEKSSubnetsNoCredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSSubnetsNoCredentialsDefault creates a ListEKSSubnetsNoCredentialsDefault with default headers values
func NewListEKSSubnetsNoCredentialsDefault(code int) *ListEKSSubnetsNoCredentialsDefault {
	return &ListEKSSubnetsNoCredentialsDefault{
		_statusCode: code,
	}
}

/*
ListEKSSubnetsNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListEKSSubnetsNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list e k s subnets no credentials default response
func (o *ListEKSSubnetsNoCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list e k s subnets no credentials default response has a 2xx status code
func (o *ListEKSSubnetsNoCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list e k s subnets no credentials default response has a 3xx status code
func (o *ListEKSSubnetsNoCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list e k s subnets no credentials default response has a 4xx status code
func (o *ListEKSSubnetsNoCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list e k s subnets no credentials default response has a 5xx status code
func (o *ListEKSSubnetsNoCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list e k s subnets no credentials default response a status code equal to that given
func (o *ListEKSSubnetsNoCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListEKSSubnetsNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/subnets][%d] listEKSSubnetsNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListEKSSubnetsNoCredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/subnets][%d] listEKSSubnetsNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListEKSSubnetsNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListEKSSubnetsNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
