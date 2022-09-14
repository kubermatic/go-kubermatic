// Code generated by go-swagger; DO NOT EDIT.

package etcdrestore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// CreateEtcdRestoreReader is a Reader for the CreateEtcdRestore structure.
type CreateEtcdRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEtcdRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateEtcdRestoreCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateEtcdRestoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateEtcdRestoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateEtcdRestoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateEtcdRestoreCreated creates a CreateEtcdRestoreCreated with default headers values
func NewCreateEtcdRestoreCreated() *CreateEtcdRestoreCreated {
	return &CreateEtcdRestoreCreated{}
}

/* CreateEtcdRestoreCreated describes a response with status code 201, with default header values.

EtcdBackupConfig
*/
type CreateEtcdRestoreCreated struct {
	Payload *models.EtcdBackupConfig
}

// IsSuccess returns true when this create etcd restore created response has a 2xx status code
func (o *CreateEtcdRestoreCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create etcd restore created response has a 3xx status code
func (o *CreateEtcdRestoreCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create etcd restore created response has a 4xx status code
func (o *CreateEtcdRestoreCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create etcd restore created response has a 5xx status code
func (o *CreateEtcdRestoreCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create etcd restore created response a status code equal to that given
func (o *CreateEtcdRestoreCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateEtcdRestoreCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores][%d] createEtcdRestoreCreated  %+v", 201, o.Payload)
}

func (o *CreateEtcdRestoreCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores][%d] createEtcdRestoreCreated  %+v", 201, o.Payload)
}

func (o *CreateEtcdRestoreCreated) GetPayload() *models.EtcdBackupConfig {
	return o.Payload
}

func (o *CreateEtcdRestoreCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EtcdBackupConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEtcdRestoreUnauthorized creates a CreateEtcdRestoreUnauthorized with default headers values
func NewCreateEtcdRestoreUnauthorized() *CreateEtcdRestoreUnauthorized {
	return &CreateEtcdRestoreUnauthorized{}
}

/* CreateEtcdRestoreUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateEtcdRestoreUnauthorized struct {
}

// IsSuccess returns true when this create etcd restore unauthorized response has a 2xx status code
func (o *CreateEtcdRestoreUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create etcd restore unauthorized response has a 3xx status code
func (o *CreateEtcdRestoreUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create etcd restore unauthorized response has a 4xx status code
func (o *CreateEtcdRestoreUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create etcd restore unauthorized response has a 5xx status code
func (o *CreateEtcdRestoreUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create etcd restore unauthorized response a status code equal to that given
func (o *CreateEtcdRestoreUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateEtcdRestoreUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores][%d] createEtcdRestoreUnauthorized ", 401)
}

func (o *CreateEtcdRestoreUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores][%d] createEtcdRestoreUnauthorized ", 401)
}

func (o *CreateEtcdRestoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateEtcdRestoreForbidden creates a CreateEtcdRestoreForbidden with default headers values
func NewCreateEtcdRestoreForbidden() *CreateEtcdRestoreForbidden {
	return &CreateEtcdRestoreForbidden{}
}

/* CreateEtcdRestoreForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateEtcdRestoreForbidden struct {
}

// IsSuccess returns true when this create etcd restore forbidden response has a 2xx status code
func (o *CreateEtcdRestoreForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create etcd restore forbidden response has a 3xx status code
func (o *CreateEtcdRestoreForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create etcd restore forbidden response has a 4xx status code
func (o *CreateEtcdRestoreForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create etcd restore forbidden response has a 5xx status code
func (o *CreateEtcdRestoreForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create etcd restore forbidden response a status code equal to that given
func (o *CreateEtcdRestoreForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateEtcdRestoreForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores][%d] createEtcdRestoreForbidden ", 403)
}

func (o *CreateEtcdRestoreForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores][%d] createEtcdRestoreForbidden ", 403)
}

func (o *CreateEtcdRestoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateEtcdRestoreDefault creates a CreateEtcdRestoreDefault with default headers values
func NewCreateEtcdRestoreDefault(code int) *CreateEtcdRestoreDefault {
	return &CreateEtcdRestoreDefault{
		_statusCode: code,
	}
}

/* CreateEtcdRestoreDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateEtcdRestoreDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create etcd restore default response
func (o *CreateEtcdRestoreDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create etcd restore default response has a 2xx status code
func (o *CreateEtcdRestoreDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create etcd restore default response has a 3xx status code
func (o *CreateEtcdRestoreDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create etcd restore default response has a 4xx status code
func (o *CreateEtcdRestoreDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create etcd restore default response has a 5xx status code
func (o *CreateEtcdRestoreDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create etcd restore default response a status code equal to that given
func (o *CreateEtcdRestoreDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateEtcdRestoreDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores][%d] createEtcdRestore default  %+v", o._statusCode, o.Payload)
}

func (o *CreateEtcdRestoreDefault) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores][%d] createEtcdRestore default  %+v", o._statusCode, o.Payload)
}

func (o *CreateEtcdRestoreDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateEtcdRestoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
