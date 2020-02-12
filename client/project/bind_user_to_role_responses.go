// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// BindUserToRoleReader is a Reader for the BindUserToRole structure.
type BindUserToRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BindUserToRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBindUserToRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewBindUserToRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBindUserToRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBindUserToRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBindUserToRoleOK creates a BindUserToRoleOK with default headers values
func NewBindUserToRoleOK() *BindUserToRoleOK {
	return &BindUserToRoleOK{}
}

/*BindUserToRoleOK handles this case with default header values.

RoleBinding
*/
type BindUserToRoleOK struct {
	Payload *models.RoleBinding
}

func (o *BindUserToRoleOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings][%d] bindUserToRoleOK  %+v", 200, o.Payload)
}

func (o *BindUserToRoleOK) GetPayload() *models.RoleBinding {
	return o.Payload
}

func (o *BindUserToRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBindUserToRoleUnauthorized creates a BindUserToRoleUnauthorized with default headers values
func NewBindUserToRoleUnauthorized() *BindUserToRoleUnauthorized {
	return &BindUserToRoleUnauthorized{}
}

/*BindUserToRoleUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type BindUserToRoleUnauthorized struct {
}

func (o *BindUserToRoleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings][%d] bindUserToRoleUnauthorized ", 401)
}

func (o *BindUserToRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBindUserToRoleForbidden creates a BindUserToRoleForbidden with default headers values
func NewBindUserToRoleForbidden() *BindUserToRoleForbidden {
	return &BindUserToRoleForbidden{}
}

/*BindUserToRoleForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type BindUserToRoleForbidden struct {
}

func (o *BindUserToRoleForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings][%d] bindUserToRoleForbidden ", 403)
}

func (o *BindUserToRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBindUserToRoleDefault creates a BindUserToRoleDefault with default headers values
func NewBindUserToRoleDefault(code int) *BindUserToRoleDefault {
	return &BindUserToRoleDefault{
		_statusCode: code,
	}
}

/*BindUserToRoleDefault handles this case with default header values.

errorResponse
*/
type BindUserToRoleDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the bind user to role default response
func (o *BindUserToRoleDefault) Code() int {
	return o._statusCode
}

func (o *BindUserToRoleDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings][%d] bindUserToRole default  %+v", o._statusCode, o.Payload)
}

func (o *BindUserToRoleDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BindUserToRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
