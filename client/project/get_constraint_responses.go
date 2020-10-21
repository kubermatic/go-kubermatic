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

// GetConstraintReader is a Reader for the GetConstraint structure.
type GetConstraintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConstraintReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConstraintOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetConstraintUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConstraintForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetConstraintDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetConstraintOK creates a GetConstraintOK with default headers values
func NewGetConstraintOK() *GetConstraintOK {
	return &GetConstraintOK{}
}

/*GetConstraintOK handles this case with default header values.

Constraint
*/
type GetConstraintOK struct {
	Payload *models.Constraint
}

func (o *GetConstraintOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints/{constraint_name}][%d] getConstraintOK  %+v", 200, o.Payload)
}

func (o *GetConstraintOK) GetPayload() *models.Constraint {
	return o.Payload
}

func (o *GetConstraintOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Constraint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConstraintUnauthorized creates a GetConstraintUnauthorized with default headers values
func NewGetConstraintUnauthorized() *GetConstraintUnauthorized {
	return &GetConstraintUnauthorized{}
}

/*GetConstraintUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type GetConstraintUnauthorized struct {
}

func (o *GetConstraintUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints/{constraint_name}][%d] getConstraintUnauthorized ", 401)
}

func (o *GetConstraintUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConstraintForbidden creates a GetConstraintForbidden with default headers values
func NewGetConstraintForbidden() *GetConstraintForbidden {
	return &GetConstraintForbidden{}
}

/*GetConstraintForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type GetConstraintForbidden struct {
}

func (o *GetConstraintForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints/{constraint_name}][%d] getConstraintForbidden ", 403)
}

func (o *GetConstraintForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConstraintDefault creates a GetConstraintDefault with default headers values
func NewGetConstraintDefault(code int) *GetConstraintDefault {
	return &GetConstraintDefault{
		_statusCode: code,
	}
}

/*GetConstraintDefault handles this case with default header values.

errorResponse
*/
type GetConstraintDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get constraint default response
func (o *GetConstraintDefault) Code() int {
	return o._statusCode
}

func (o *GetConstraintDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints/{constraint_name}][%d] getConstraint default  %+v", o._statusCode, o.Payload)
}

func (o *GetConstraintDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetConstraintDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
