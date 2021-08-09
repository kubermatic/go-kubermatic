// Code generated by go-swagger; DO NOT EDIT.

package constrainttemplates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// PatchConstraintTemplateReader is a Reader for the PatchConstraintTemplate structure.
type PatchConstraintTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConstraintTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchConstraintTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchConstraintTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConstraintTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchConstraintTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchConstraintTemplateOK creates a PatchConstraintTemplateOK with default headers values
func NewPatchConstraintTemplateOK() *PatchConstraintTemplateOK {
	return &PatchConstraintTemplateOK{}
}

/* PatchConstraintTemplateOK describes a response with status code 200, with default header values.

ConstraintTemplate
*/
type PatchConstraintTemplateOK struct {
	Payload *models.ConstraintTemplate
}

func (o *PatchConstraintTemplateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/constrainttemplates/{ct_name}][%d] patchConstraintTemplateOK  %+v", 200, o.Payload)
}
func (o *PatchConstraintTemplateOK) GetPayload() *models.ConstraintTemplate {
	return o.Payload
}

func (o *PatchConstraintTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConstraintTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConstraintTemplateUnauthorized creates a PatchConstraintTemplateUnauthorized with default headers values
func NewPatchConstraintTemplateUnauthorized() *PatchConstraintTemplateUnauthorized {
	return &PatchConstraintTemplateUnauthorized{}
}

/* PatchConstraintTemplateUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchConstraintTemplateUnauthorized struct {
}

func (o *PatchConstraintTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/constrainttemplates/{ct_name}][%d] patchConstraintTemplateUnauthorized ", 401)
}

func (o *PatchConstraintTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConstraintTemplateForbidden creates a PatchConstraintTemplateForbidden with default headers values
func NewPatchConstraintTemplateForbidden() *PatchConstraintTemplateForbidden {
	return &PatchConstraintTemplateForbidden{}
}

/* PatchConstraintTemplateForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchConstraintTemplateForbidden struct {
}

func (o *PatchConstraintTemplateForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/constrainttemplates/{ct_name}][%d] patchConstraintTemplateForbidden ", 403)
}

func (o *PatchConstraintTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConstraintTemplateDefault creates a PatchConstraintTemplateDefault with default headers values
func NewPatchConstraintTemplateDefault(code int) *PatchConstraintTemplateDefault {
	return &PatchConstraintTemplateDefault{
		_statusCode: code,
	}
}

/* PatchConstraintTemplateDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PatchConstraintTemplateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch constraint template default response
func (o *PatchConstraintTemplateDefault) Code() int {
	return o._statusCode
}

func (o *PatchConstraintTemplateDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/constrainttemplates/{ct_name}][%d] patchConstraintTemplate default  %+v", o._statusCode, o.Payload)
}
func (o *PatchConstraintTemplateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchConstraintTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
