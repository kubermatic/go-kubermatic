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

// PatchConstraintReader is a Reader for the PatchConstraint structure.
type PatchConstraintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConstraintReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchConstraintOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchConstraintUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConstraintForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchConstraintDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchConstraintOK creates a PatchConstraintOK with default headers values
func NewPatchConstraintOK() *PatchConstraintOK {
	return &PatchConstraintOK{}
}

/*PatchConstraintOK handles this case with default header values.

Constraint
*/
type PatchConstraintOK struct {
	Payload *models.Constraint
}

func (o *PatchConstraintOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints/{constraint_name}][%d] patchConstraintOK  %+v", 200, o.Payload)
}

func (o *PatchConstraintOK) GetPayload() *models.Constraint {
	return o.Payload
}

func (o *PatchConstraintOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Constraint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConstraintUnauthorized creates a PatchConstraintUnauthorized with default headers values
func NewPatchConstraintUnauthorized() *PatchConstraintUnauthorized {
	return &PatchConstraintUnauthorized{}
}

/*PatchConstraintUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type PatchConstraintUnauthorized struct {
}

func (o *PatchConstraintUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints/{constraint_name}][%d] patchConstraintUnauthorized ", 401)
}

func (o *PatchConstraintUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConstraintForbidden creates a PatchConstraintForbidden with default headers values
func NewPatchConstraintForbidden() *PatchConstraintForbidden {
	return &PatchConstraintForbidden{}
}

/*PatchConstraintForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type PatchConstraintForbidden struct {
}

func (o *PatchConstraintForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints/{constraint_name}][%d] patchConstraintForbidden ", 403)
}

func (o *PatchConstraintForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConstraintDefault creates a PatchConstraintDefault with default headers values
func NewPatchConstraintDefault(code int) *PatchConstraintDefault {
	return &PatchConstraintDefault{
		_statusCode: code,
	}
}

/*PatchConstraintDefault handles this case with default header values.

errorResponse
*/
type PatchConstraintDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch constraint default response
func (o *PatchConstraintDefault) Code() int {
	return o._statusCode
}

func (o *PatchConstraintDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints/{constraint_name}][%d] patchConstraint default  %+v", o._statusCode, o.Payload)
}

func (o *PatchConstraintDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchConstraintDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
