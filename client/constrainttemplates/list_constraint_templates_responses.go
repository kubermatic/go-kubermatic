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

/*
ListConstraintTemplatesOK describes a response with status code 200, with default header values.

ConstraintTemplate
*/
type ListConstraintTemplatesOK struct {
	Payload []*models.ConstraintTemplate
}

// IsSuccess returns true when this list constraint templates o k response has a 2xx status code
func (o *ListConstraintTemplatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list constraint templates o k response has a 3xx status code
func (o *ListConstraintTemplatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list constraint templates o k response has a 4xx status code
func (o *ListConstraintTemplatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list constraint templates o k response has a 5xx status code
func (o *ListConstraintTemplatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list constraint templates o k response a status code equal to that given
func (o *ListConstraintTemplatesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListConstraintTemplatesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/constrainttemplates][%d] listConstraintTemplatesOK  %+v", 200, o.Payload)
}

func (o *ListConstraintTemplatesOK) String() string {
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

/*
ListConstraintTemplatesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListConstraintTemplatesUnauthorized struct {
}

// IsSuccess returns true when this list constraint templates unauthorized response has a 2xx status code
func (o *ListConstraintTemplatesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list constraint templates unauthorized response has a 3xx status code
func (o *ListConstraintTemplatesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list constraint templates unauthorized response has a 4xx status code
func (o *ListConstraintTemplatesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list constraint templates unauthorized response has a 5xx status code
func (o *ListConstraintTemplatesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list constraint templates unauthorized response a status code equal to that given
func (o *ListConstraintTemplatesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListConstraintTemplatesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/constrainttemplates][%d] listConstraintTemplatesUnauthorized ", 401)
}

func (o *ListConstraintTemplatesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/constrainttemplates][%d] listConstraintTemplatesUnauthorized ", 401)
}

func (o *ListConstraintTemplatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListConstraintTemplatesForbidden creates a ListConstraintTemplatesForbidden with default headers values
func NewListConstraintTemplatesForbidden() *ListConstraintTemplatesForbidden {
	return &ListConstraintTemplatesForbidden{}
}

/*
ListConstraintTemplatesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListConstraintTemplatesForbidden struct {
}

// IsSuccess returns true when this list constraint templates forbidden response has a 2xx status code
func (o *ListConstraintTemplatesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list constraint templates forbidden response has a 3xx status code
func (o *ListConstraintTemplatesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list constraint templates forbidden response has a 4xx status code
func (o *ListConstraintTemplatesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list constraint templates forbidden response has a 5xx status code
func (o *ListConstraintTemplatesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list constraint templates forbidden response a status code equal to that given
func (o *ListConstraintTemplatesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListConstraintTemplatesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/constrainttemplates][%d] listConstraintTemplatesForbidden ", 403)
}

func (o *ListConstraintTemplatesForbidden) String() string {
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

/*
ListConstraintTemplatesDefault describes a response with status code -1, with default header values.

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

// IsSuccess returns true when this list constraint templates default response has a 2xx status code
func (o *ListConstraintTemplatesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list constraint templates default response has a 3xx status code
func (o *ListConstraintTemplatesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list constraint templates default response has a 4xx status code
func (o *ListConstraintTemplatesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list constraint templates default response has a 5xx status code
func (o *ListConstraintTemplatesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list constraint templates default response a status code equal to that given
func (o *ListConstraintTemplatesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListConstraintTemplatesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/constrainttemplates][%d] listConstraintTemplates default  %+v", o._statusCode, o.Payload)
}

func (o *ListConstraintTemplatesDefault) String() string {
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
