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

// UpdateAlertmanagerReader is a Reader for the UpdateAlertmanager structure.
type UpdateAlertmanagerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAlertmanagerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAlertmanagerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateAlertmanagerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAlertmanagerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateAlertmanagerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAlertmanagerOK creates a UpdateAlertmanagerOK with default headers values
func NewUpdateAlertmanagerOK() *UpdateAlertmanagerOK {
	return &UpdateAlertmanagerOK{}
}

/*
UpdateAlertmanagerOK describes a response with status code 200, with default header values.

Alertmanager
*/
type UpdateAlertmanagerOK struct {
	Payload *models.Alertmanager
}

// IsSuccess returns true when this update alertmanager o k response has a 2xx status code
func (o *UpdateAlertmanagerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update alertmanager o k response has a 3xx status code
func (o *UpdateAlertmanagerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update alertmanager o k response has a 4xx status code
func (o *UpdateAlertmanagerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update alertmanager o k response has a 5xx status code
func (o *UpdateAlertmanagerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update alertmanager o k response a status code equal to that given
func (o *UpdateAlertmanagerOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateAlertmanagerOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config][%d] updateAlertmanagerOK  %+v", 200, o.Payload)
}

func (o *UpdateAlertmanagerOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config][%d] updateAlertmanagerOK  %+v", 200, o.Payload)
}

func (o *UpdateAlertmanagerOK) GetPayload() *models.Alertmanager {
	return o.Payload
}

func (o *UpdateAlertmanagerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Alertmanager)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAlertmanagerUnauthorized creates a UpdateAlertmanagerUnauthorized with default headers values
func NewUpdateAlertmanagerUnauthorized() *UpdateAlertmanagerUnauthorized {
	return &UpdateAlertmanagerUnauthorized{}
}

/*
UpdateAlertmanagerUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type UpdateAlertmanagerUnauthorized struct {
}

// IsSuccess returns true when this update alertmanager unauthorized response has a 2xx status code
func (o *UpdateAlertmanagerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update alertmanager unauthorized response has a 3xx status code
func (o *UpdateAlertmanagerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update alertmanager unauthorized response has a 4xx status code
func (o *UpdateAlertmanagerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update alertmanager unauthorized response has a 5xx status code
func (o *UpdateAlertmanagerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update alertmanager unauthorized response a status code equal to that given
func (o *UpdateAlertmanagerUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateAlertmanagerUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config][%d] updateAlertmanagerUnauthorized ", 401)
}

func (o *UpdateAlertmanagerUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config][%d] updateAlertmanagerUnauthorized ", 401)
}

func (o *UpdateAlertmanagerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAlertmanagerForbidden creates a UpdateAlertmanagerForbidden with default headers values
func NewUpdateAlertmanagerForbidden() *UpdateAlertmanagerForbidden {
	return &UpdateAlertmanagerForbidden{}
}

/*
UpdateAlertmanagerForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type UpdateAlertmanagerForbidden struct {
}

// IsSuccess returns true when this update alertmanager forbidden response has a 2xx status code
func (o *UpdateAlertmanagerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update alertmanager forbidden response has a 3xx status code
func (o *UpdateAlertmanagerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update alertmanager forbidden response has a 4xx status code
func (o *UpdateAlertmanagerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update alertmanager forbidden response has a 5xx status code
func (o *UpdateAlertmanagerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update alertmanager forbidden response a status code equal to that given
func (o *UpdateAlertmanagerForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateAlertmanagerForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config][%d] updateAlertmanagerForbidden ", 403)
}

func (o *UpdateAlertmanagerForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config][%d] updateAlertmanagerForbidden ", 403)
}

func (o *UpdateAlertmanagerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAlertmanagerDefault creates a UpdateAlertmanagerDefault with default headers values
func NewUpdateAlertmanagerDefault(code int) *UpdateAlertmanagerDefault {
	return &UpdateAlertmanagerDefault{
		_statusCode: code,
	}
}

/*
UpdateAlertmanagerDefault describes a response with status code -1, with default header values.

errorResponse
*/
type UpdateAlertmanagerDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update alertmanager default response
func (o *UpdateAlertmanagerDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update alertmanager default response has a 2xx status code
func (o *UpdateAlertmanagerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update alertmanager default response has a 3xx status code
func (o *UpdateAlertmanagerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update alertmanager default response has a 4xx status code
func (o *UpdateAlertmanagerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update alertmanager default response has a 5xx status code
func (o *UpdateAlertmanagerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update alertmanager default response a status code equal to that given
func (o *UpdateAlertmanagerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateAlertmanagerDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config][%d] updateAlertmanager default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAlertmanagerDefault) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config][%d] updateAlertmanager default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAlertmanagerDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateAlertmanagerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
