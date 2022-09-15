// Code generated by go-swagger; DO NOT EDIT.

package applications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// CreateApplicationInstallationReader is a Reader for the CreateApplicationInstallation structure.
type CreateApplicationInstallationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateApplicationInstallationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateApplicationInstallationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateApplicationInstallationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateApplicationInstallationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateApplicationInstallationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateApplicationInstallationCreated creates a CreateApplicationInstallationCreated with default headers values
func NewCreateApplicationInstallationCreated() *CreateApplicationInstallationCreated {
	return &CreateApplicationInstallationCreated{}
}

/*
CreateApplicationInstallationCreated describes a response with status code 201, with default header values.

ApplicationInstallation
*/
type CreateApplicationInstallationCreated struct {
	Payload *models.ApplicationInstallation
}

// IsSuccess returns true when this create application installation created response has a 2xx status code
func (o *CreateApplicationInstallationCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create application installation created response has a 3xx status code
func (o *CreateApplicationInstallationCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create application installation created response has a 4xx status code
func (o *CreateApplicationInstallationCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create application installation created response has a 5xx status code
func (o *CreateApplicationInstallationCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create application installation created response a status code equal to that given
func (o *CreateApplicationInstallationCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateApplicationInstallationCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations][%d] createApplicationInstallationCreated  %+v", 201, o.Payload)
}

func (o *CreateApplicationInstallationCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations][%d] createApplicationInstallationCreated  %+v", 201, o.Payload)
}

func (o *CreateApplicationInstallationCreated) GetPayload() *models.ApplicationInstallation {
	return o.Payload
}

func (o *CreateApplicationInstallationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationInstallation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateApplicationInstallationUnauthorized creates a CreateApplicationInstallationUnauthorized with default headers values
func NewCreateApplicationInstallationUnauthorized() *CreateApplicationInstallationUnauthorized {
	return &CreateApplicationInstallationUnauthorized{}
}

/*
CreateApplicationInstallationUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateApplicationInstallationUnauthorized struct {
}

// IsSuccess returns true when this create application installation unauthorized response has a 2xx status code
func (o *CreateApplicationInstallationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create application installation unauthorized response has a 3xx status code
func (o *CreateApplicationInstallationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create application installation unauthorized response has a 4xx status code
func (o *CreateApplicationInstallationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create application installation unauthorized response has a 5xx status code
func (o *CreateApplicationInstallationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create application installation unauthorized response a status code equal to that given
func (o *CreateApplicationInstallationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateApplicationInstallationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations][%d] createApplicationInstallationUnauthorized ", 401)
}

func (o *CreateApplicationInstallationUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations][%d] createApplicationInstallationUnauthorized ", 401)
}

func (o *CreateApplicationInstallationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateApplicationInstallationForbidden creates a CreateApplicationInstallationForbidden with default headers values
func NewCreateApplicationInstallationForbidden() *CreateApplicationInstallationForbidden {
	return &CreateApplicationInstallationForbidden{}
}

/*
CreateApplicationInstallationForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateApplicationInstallationForbidden struct {
}

// IsSuccess returns true when this create application installation forbidden response has a 2xx status code
func (o *CreateApplicationInstallationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create application installation forbidden response has a 3xx status code
func (o *CreateApplicationInstallationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create application installation forbidden response has a 4xx status code
func (o *CreateApplicationInstallationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create application installation forbidden response has a 5xx status code
func (o *CreateApplicationInstallationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create application installation forbidden response a status code equal to that given
func (o *CreateApplicationInstallationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateApplicationInstallationForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations][%d] createApplicationInstallationForbidden ", 403)
}

func (o *CreateApplicationInstallationForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations][%d] createApplicationInstallationForbidden ", 403)
}

func (o *CreateApplicationInstallationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateApplicationInstallationDefault creates a CreateApplicationInstallationDefault with default headers values
func NewCreateApplicationInstallationDefault(code int) *CreateApplicationInstallationDefault {
	return &CreateApplicationInstallationDefault{
		_statusCode: code,
	}
}

/*
CreateApplicationInstallationDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateApplicationInstallationDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create application installation default response
func (o *CreateApplicationInstallationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create application installation default response has a 2xx status code
func (o *CreateApplicationInstallationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create application installation default response has a 3xx status code
func (o *CreateApplicationInstallationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create application installation default response has a 4xx status code
func (o *CreateApplicationInstallationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create application installation default response has a 5xx status code
func (o *CreateApplicationInstallationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create application installation default response a status code equal to that given
func (o *CreateApplicationInstallationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateApplicationInstallationDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations][%d] createApplicationInstallation default  %+v", o._statusCode, o.Payload)
}

func (o *CreateApplicationInstallationDefault) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations][%d] createApplicationInstallation default  %+v", o._statusCode, o.Payload)
}

func (o *CreateApplicationInstallationDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateApplicationInstallationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
