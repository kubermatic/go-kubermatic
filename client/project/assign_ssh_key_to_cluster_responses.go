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

// AssignSSHKeyToClusterReader is a Reader for the AssignSSHKeyToCluster structure.
type AssignSSHKeyToClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignSSHKeyToClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAssignSSHKeyToClusterCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAssignSSHKeyToClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAssignSSHKeyToClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAssignSSHKeyToClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssignSSHKeyToClusterCreated creates a AssignSSHKeyToClusterCreated with default headers values
func NewAssignSSHKeyToClusterCreated() *AssignSSHKeyToClusterCreated {
	return &AssignSSHKeyToClusterCreated{}
}

/*
AssignSSHKeyToClusterCreated describes a response with status code 201, with default header values.

SSHKey
*/
type AssignSSHKeyToClusterCreated struct {
	Payload *models.SSHKey
}

// IsSuccess returns true when this assign Ssh key to cluster created response has a 2xx status code
func (o *AssignSSHKeyToClusterCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assign Ssh key to cluster created response has a 3xx status code
func (o *AssignSSHKeyToClusterCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign Ssh key to cluster created response has a 4xx status code
func (o *AssignSSHKeyToClusterCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign Ssh key to cluster created response has a 5xx status code
func (o *AssignSSHKeyToClusterCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this assign Ssh key to cluster created response a status code equal to that given
func (o *AssignSSHKeyToClusterCreated) IsCode(code int) bool {
	return code == 201
}

func (o *AssignSSHKeyToClusterCreated) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSshKeyToClusterCreated  %+v", 201, o.Payload)
}

func (o *AssignSSHKeyToClusterCreated) String() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSshKeyToClusterCreated  %+v", 201, o.Payload)
}

func (o *AssignSSHKeyToClusterCreated) GetPayload() *models.SSHKey {
	return o.Payload
}

func (o *AssignSSHKeyToClusterCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SSHKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignSSHKeyToClusterUnauthorized creates a AssignSSHKeyToClusterUnauthorized with default headers values
func NewAssignSSHKeyToClusterUnauthorized() *AssignSSHKeyToClusterUnauthorized {
	return &AssignSSHKeyToClusterUnauthorized{}
}

/*
AssignSSHKeyToClusterUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type AssignSSHKeyToClusterUnauthorized struct {
}

// IsSuccess returns true when this assign Ssh key to cluster unauthorized response has a 2xx status code
func (o *AssignSSHKeyToClusterUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign Ssh key to cluster unauthorized response has a 3xx status code
func (o *AssignSSHKeyToClusterUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign Ssh key to cluster unauthorized response has a 4xx status code
func (o *AssignSSHKeyToClusterUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this assign Ssh key to cluster unauthorized response has a 5xx status code
func (o *AssignSSHKeyToClusterUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this assign Ssh key to cluster unauthorized response a status code equal to that given
func (o *AssignSSHKeyToClusterUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AssignSSHKeyToClusterUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSshKeyToClusterUnauthorized ", 401)
}

func (o *AssignSSHKeyToClusterUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSshKeyToClusterUnauthorized ", 401)
}

func (o *AssignSSHKeyToClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignSSHKeyToClusterForbidden creates a AssignSSHKeyToClusterForbidden with default headers values
func NewAssignSSHKeyToClusterForbidden() *AssignSSHKeyToClusterForbidden {
	return &AssignSSHKeyToClusterForbidden{}
}

/*
AssignSSHKeyToClusterForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type AssignSSHKeyToClusterForbidden struct {
}

// IsSuccess returns true when this assign Ssh key to cluster forbidden response has a 2xx status code
func (o *AssignSSHKeyToClusterForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign Ssh key to cluster forbidden response has a 3xx status code
func (o *AssignSSHKeyToClusterForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign Ssh key to cluster forbidden response has a 4xx status code
func (o *AssignSSHKeyToClusterForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this assign Ssh key to cluster forbidden response has a 5xx status code
func (o *AssignSSHKeyToClusterForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this assign Ssh key to cluster forbidden response a status code equal to that given
func (o *AssignSSHKeyToClusterForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AssignSSHKeyToClusterForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSshKeyToClusterForbidden ", 403)
}

func (o *AssignSSHKeyToClusterForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSshKeyToClusterForbidden ", 403)
}

func (o *AssignSSHKeyToClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignSSHKeyToClusterDefault creates a AssignSSHKeyToClusterDefault with default headers values
func NewAssignSSHKeyToClusterDefault(code int) *AssignSSHKeyToClusterDefault {
	return &AssignSSHKeyToClusterDefault{
		_statusCode: code,
	}
}

/*
AssignSSHKeyToClusterDefault describes a response with status code -1, with default header values.

errorResponse
*/
type AssignSSHKeyToClusterDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the assign SSH key to cluster default response
func (o *AssignSSHKeyToClusterDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this assign SSH key to cluster default response has a 2xx status code
func (o *AssignSSHKeyToClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this assign SSH key to cluster default response has a 3xx status code
func (o *AssignSSHKeyToClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this assign SSH key to cluster default response has a 4xx status code
func (o *AssignSSHKeyToClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this assign SSH key to cluster default response has a 5xx status code
func (o *AssignSSHKeyToClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this assign SSH key to cluster default response a status code equal to that given
func (o *AssignSSHKeyToClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AssignSSHKeyToClusterDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSSHKeyToCluster default  %+v", o._statusCode, o.Payload)
}

func (o *AssignSSHKeyToClusterDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSSHKeyToCluster default  %+v", o._statusCode, o.Payload)
}

func (o *AssignSSHKeyToClusterDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AssignSSHKeyToClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
