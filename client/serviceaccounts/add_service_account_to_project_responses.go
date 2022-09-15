// Code generated by go-swagger; DO NOT EDIT.

package serviceaccounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// AddServiceAccountToProjectReader is a Reader for the AddServiceAccountToProject structure.
type AddServiceAccountToProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddServiceAccountToProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddServiceAccountToProjectCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddServiceAccountToProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddServiceAccountToProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddServiceAccountToProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddServiceAccountToProjectCreated creates a AddServiceAccountToProjectCreated with default headers values
func NewAddServiceAccountToProjectCreated() *AddServiceAccountToProjectCreated {
	return &AddServiceAccountToProjectCreated{}
}

/*
AddServiceAccountToProjectCreated describes a response with status code 201, with default header values.

ServiceAccount
*/
type AddServiceAccountToProjectCreated struct {
	Payload *models.ServiceAccount
}

// IsSuccess returns true when this add service account to project created response has a 2xx status code
func (o *AddServiceAccountToProjectCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add service account to project created response has a 3xx status code
func (o *AddServiceAccountToProjectCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add service account to project created response has a 4xx status code
func (o *AddServiceAccountToProjectCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add service account to project created response has a 5xx status code
func (o *AddServiceAccountToProjectCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add service account to project created response a status code equal to that given
func (o *AddServiceAccountToProjectCreated) IsCode(code int) bool {
	return code == 201
}

func (o *AddServiceAccountToProjectCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/serviceaccounts][%d] addServiceAccountToProjectCreated  %+v", 201, o.Payload)
}

func (o *AddServiceAccountToProjectCreated) String() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/serviceaccounts][%d] addServiceAccountToProjectCreated  %+v", 201, o.Payload)
}

func (o *AddServiceAccountToProjectCreated) GetPayload() *models.ServiceAccount {
	return o.Payload
}

func (o *AddServiceAccountToProjectCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddServiceAccountToProjectUnauthorized creates a AddServiceAccountToProjectUnauthorized with default headers values
func NewAddServiceAccountToProjectUnauthorized() *AddServiceAccountToProjectUnauthorized {
	return &AddServiceAccountToProjectUnauthorized{}
}

/*
AddServiceAccountToProjectUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type AddServiceAccountToProjectUnauthorized struct {
}

// IsSuccess returns true when this add service account to project unauthorized response has a 2xx status code
func (o *AddServiceAccountToProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add service account to project unauthorized response has a 3xx status code
func (o *AddServiceAccountToProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add service account to project unauthorized response has a 4xx status code
func (o *AddServiceAccountToProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add service account to project unauthorized response has a 5xx status code
func (o *AddServiceAccountToProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add service account to project unauthorized response a status code equal to that given
func (o *AddServiceAccountToProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AddServiceAccountToProjectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/serviceaccounts][%d] addServiceAccountToProjectUnauthorized ", 401)
}

func (o *AddServiceAccountToProjectUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/serviceaccounts][%d] addServiceAccountToProjectUnauthorized ", 401)
}

func (o *AddServiceAccountToProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddServiceAccountToProjectForbidden creates a AddServiceAccountToProjectForbidden with default headers values
func NewAddServiceAccountToProjectForbidden() *AddServiceAccountToProjectForbidden {
	return &AddServiceAccountToProjectForbidden{}
}

/*
AddServiceAccountToProjectForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type AddServiceAccountToProjectForbidden struct {
}

// IsSuccess returns true when this add service account to project forbidden response has a 2xx status code
func (o *AddServiceAccountToProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add service account to project forbidden response has a 3xx status code
func (o *AddServiceAccountToProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add service account to project forbidden response has a 4xx status code
func (o *AddServiceAccountToProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add service account to project forbidden response has a 5xx status code
func (o *AddServiceAccountToProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add service account to project forbidden response a status code equal to that given
func (o *AddServiceAccountToProjectForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AddServiceAccountToProjectForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/serviceaccounts][%d] addServiceAccountToProjectForbidden ", 403)
}

func (o *AddServiceAccountToProjectForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/serviceaccounts][%d] addServiceAccountToProjectForbidden ", 403)
}

func (o *AddServiceAccountToProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddServiceAccountToProjectDefault creates a AddServiceAccountToProjectDefault with default headers values
func NewAddServiceAccountToProjectDefault(code int) *AddServiceAccountToProjectDefault {
	return &AddServiceAccountToProjectDefault{
		_statusCode: code,
	}
}

/*
AddServiceAccountToProjectDefault describes a response with status code -1, with default header values.

errorResponse
*/
type AddServiceAccountToProjectDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add service account to project default response
func (o *AddServiceAccountToProjectDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this add service account to project default response has a 2xx status code
func (o *AddServiceAccountToProjectDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add service account to project default response has a 3xx status code
func (o *AddServiceAccountToProjectDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add service account to project default response has a 4xx status code
func (o *AddServiceAccountToProjectDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add service account to project default response has a 5xx status code
func (o *AddServiceAccountToProjectDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add service account to project default response a status code equal to that given
func (o *AddServiceAccountToProjectDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AddServiceAccountToProjectDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/serviceaccounts][%d] addServiceAccountToProject default  %+v", o._statusCode, o.Payload)
}

func (o *AddServiceAccountToProjectDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/serviceaccounts][%d] addServiceAccountToProject default  %+v", o._statusCode, o.Payload)
}

func (o *AddServiceAccountToProjectDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddServiceAccountToProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
