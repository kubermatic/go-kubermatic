// Code generated by go-swagger; DO NOT EDIT.

package aks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListAKSNodeVersionsNoCredentialsReader is a Reader for the ListAKSNodeVersionsNoCredentials structure.
type ListAKSNodeVersionsNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAKSNodeVersionsNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAKSNodeVersionsNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAKSNodeVersionsNoCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAKSNodeVersionsNoCredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListAKSNodeVersionsNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAKSNodeVersionsNoCredentialsOK creates a ListAKSNodeVersionsNoCredentialsOK with default headers values
func NewListAKSNodeVersionsNoCredentialsOK() *ListAKSNodeVersionsNoCredentialsOK {
	return &ListAKSNodeVersionsNoCredentialsOK{}
}

/*
ListAKSNodeVersionsNoCredentialsOK describes a response with status code 200, with default header values.

MasterVersion
*/
type ListAKSNodeVersionsNoCredentialsOK struct {
	Payload []*models.MasterVersion
}

// IsSuccess returns true when this list a k s node versions no credentials o k response has a 2xx status code
func (o *ListAKSNodeVersionsNoCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list a k s node versions no credentials o k response has a 3xx status code
func (o *ListAKSNodeVersionsNoCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list a k s node versions no credentials o k response has a 4xx status code
func (o *ListAKSNodeVersionsNoCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list a k s node versions no credentials o k response has a 5xx status code
func (o *ListAKSNodeVersionsNoCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list a k s node versions no credentials o k response a status code equal to that given
func (o *ListAKSNodeVersionsNoCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListAKSNodeVersionsNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/aks/versions][%d] listAKSNodeVersionsNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListAKSNodeVersionsNoCredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/aks/versions][%d] listAKSNodeVersionsNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListAKSNodeVersionsNoCredentialsOK) GetPayload() []*models.MasterVersion {
	return o.Payload
}

func (o *ListAKSNodeVersionsNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAKSNodeVersionsNoCredentialsUnauthorized creates a ListAKSNodeVersionsNoCredentialsUnauthorized with default headers values
func NewListAKSNodeVersionsNoCredentialsUnauthorized() *ListAKSNodeVersionsNoCredentialsUnauthorized {
	return &ListAKSNodeVersionsNoCredentialsUnauthorized{}
}

/*
ListAKSNodeVersionsNoCredentialsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListAKSNodeVersionsNoCredentialsUnauthorized struct {
}

// IsSuccess returns true when this list a k s node versions no credentials unauthorized response has a 2xx status code
func (o *ListAKSNodeVersionsNoCredentialsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list a k s node versions no credentials unauthorized response has a 3xx status code
func (o *ListAKSNodeVersionsNoCredentialsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list a k s node versions no credentials unauthorized response has a 4xx status code
func (o *ListAKSNodeVersionsNoCredentialsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list a k s node versions no credentials unauthorized response has a 5xx status code
func (o *ListAKSNodeVersionsNoCredentialsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list a k s node versions no credentials unauthorized response a status code equal to that given
func (o *ListAKSNodeVersionsNoCredentialsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListAKSNodeVersionsNoCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/aks/versions][%d] listAKSNodeVersionsNoCredentialsUnauthorized ", 401)
}

func (o *ListAKSNodeVersionsNoCredentialsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/aks/versions][%d] listAKSNodeVersionsNoCredentialsUnauthorized ", 401)
}

func (o *ListAKSNodeVersionsNoCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAKSNodeVersionsNoCredentialsForbidden creates a ListAKSNodeVersionsNoCredentialsForbidden with default headers values
func NewListAKSNodeVersionsNoCredentialsForbidden() *ListAKSNodeVersionsNoCredentialsForbidden {
	return &ListAKSNodeVersionsNoCredentialsForbidden{}
}

/*
ListAKSNodeVersionsNoCredentialsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListAKSNodeVersionsNoCredentialsForbidden struct {
}

// IsSuccess returns true when this list a k s node versions no credentials forbidden response has a 2xx status code
func (o *ListAKSNodeVersionsNoCredentialsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list a k s node versions no credentials forbidden response has a 3xx status code
func (o *ListAKSNodeVersionsNoCredentialsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list a k s node versions no credentials forbidden response has a 4xx status code
func (o *ListAKSNodeVersionsNoCredentialsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list a k s node versions no credentials forbidden response has a 5xx status code
func (o *ListAKSNodeVersionsNoCredentialsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list a k s node versions no credentials forbidden response a status code equal to that given
func (o *ListAKSNodeVersionsNoCredentialsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListAKSNodeVersionsNoCredentialsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/aks/versions][%d] listAKSNodeVersionsNoCredentialsForbidden ", 403)
}

func (o *ListAKSNodeVersionsNoCredentialsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/aks/versions][%d] listAKSNodeVersionsNoCredentialsForbidden ", 403)
}

func (o *ListAKSNodeVersionsNoCredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAKSNodeVersionsNoCredentialsDefault creates a ListAKSNodeVersionsNoCredentialsDefault with default headers values
func NewListAKSNodeVersionsNoCredentialsDefault(code int) *ListAKSNodeVersionsNoCredentialsDefault {
	return &ListAKSNodeVersionsNoCredentialsDefault{
		_statusCode: code,
	}
}

/*
ListAKSNodeVersionsNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListAKSNodeVersionsNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list a k s node versions no credentials default response
func (o *ListAKSNodeVersionsNoCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list a k s node versions no credentials default response has a 2xx status code
func (o *ListAKSNodeVersionsNoCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list a k s node versions no credentials default response has a 3xx status code
func (o *ListAKSNodeVersionsNoCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list a k s node versions no credentials default response has a 4xx status code
func (o *ListAKSNodeVersionsNoCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list a k s node versions no credentials default response has a 5xx status code
func (o *ListAKSNodeVersionsNoCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list a k s node versions no credentials default response a status code equal to that given
func (o *ListAKSNodeVersionsNoCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListAKSNodeVersionsNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/aks/versions][%d] listAKSNodeVersionsNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListAKSNodeVersionsNoCredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/aks/versions][%d] listAKSNodeVersionsNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListAKSNodeVersionsNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAKSNodeVersionsNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
