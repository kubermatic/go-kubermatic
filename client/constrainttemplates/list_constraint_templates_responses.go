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

// ListConstraintTemplatesReader is a Reader for the ListConstraintTemplates structure.
type ListConstraintTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListConstraintTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListConstraintTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListConstraintTemplatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListConstraintTemplatesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListConstraintTemplatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListConstraintTemplatesOK creates a ListConstraintTemplatesOK with default headers values
func NewListConstraintTemplatesOK() *ListConstraintTemplatesOK {
	return &ListConstraintTemplatesOK{}
}

/* ListConstraintTemplatesOK describes a response with status code 200, with default header values.

ConstraintTemplate
*/
type ListConstraintTemplatesOK struct {
	Payload []*models.ConstraintTemplate
}

func (o *ListConstraintTemplatesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/constrainttemplates][%d] listConstraintTemplatesOK  %+v", 200, o.Payload)
}
func (o *ListConstraintTemplatesOK) GetPayload() []*models.ConstraintTemplate {
	return o.Payload
}

func (o *ListConstraintTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConstraintTemplatesUnauthorized creates a ListConstraintTemplatesUnauthorized with default headers values
func NewListConstraintTemplatesUnauthorized() *ListConstraintTemplatesUnauthorized {
	return &ListConstraintTemplatesUnauthorized{}
}

/* ListConstraintTemplatesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListConstraintTemplatesUnauthorized struct {
}

func (o *ListConstraintTemplatesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/constrainttemplates][%d] listConstraintTemplatesUnauthorized ", 401)
}

func (o *ListConstraintTemplatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListConstraintTemplatesForbidden creates a ListConstraintTemplatesForbidden with default headers values
func NewListConstraintTemplatesForbidden() *ListConstraintTemplatesForbidden {
	return &ListConstraintTemplatesForbidden{}
}

/* ListConstraintTemplatesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListConstraintTemplatesForbidden struct {
}

func (o *ListConstraintTemplatesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/constrainttemplates][%d] listConstraintTemplatesForbidden ", 403)
}

func (o *ListConstraintTemplatesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListConstraintTemplatesDefault creates a ListConstraintTemplatesDefault with default headers values
func NewListConstraintTemplatesDefault(code int) *ListConstraintTemplatesDefault {
	return &ListConstraintTemplatesDefault{
		_statusCode: code,
	}
}

/* ListConstraintTemplatesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListConstraintTemplatesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list constraint templates default response
func (o *ListConstraintTemplatesDefault) Code() int {
	return o._statusCode
}

func (o *ListConstraintTemplatesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/constrainttemplates][%d] listConstraintTemplates default  %+v", o._statusCode, o.Payload)
}
func (o *ListConstraintTemplatesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListConstraintTemplatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
