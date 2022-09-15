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

// DeleteConstraintTemplateReader is a Reader for the DeleteConstraintTemplate structure.
type DeleteConstraintTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConstraintTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteConstraintTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteConstraintTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteConstraintTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteConstraintTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteConstraintTemplateOK creates a DeleteConstraintTemplateOK with default headers values
func NewDeleteConstraintTemplateOK() *DeleteConstraintTemplateOK {
	return &DeleteConstraintTemplateOK{}
}

/*
DeleteConstraintTemplateOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteConstraintTemplateOK struct {
}

// IsSuccess returns true when this delete constraint template o k response has a 2xx status code
func (o *DeleteConstraintTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete constraint template o k response has a 3xx status code
func (o *DeleteConstraintTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete constraint template o k response has a 4xx status code
func (o *DeleteConstraintTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete constraint template o k response has a 5xx status code
func (o *DeleteConstraintTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete constraint template o k response a status code equal to that given
func (o *DeleteConstraintTemplateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteConstraintTemplateOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/constrainttemplates/{ct_name}][%d] deleteConstraintTemplateOK ", 200)
}

func (o *DeleteConstraintTemplateOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/constrainttemplates/{ct_name}][%d] deleteConstraintTemplateOK ", 200)
}

func (o *DeleteConstraintTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConstraintTemplateUnauthorized creates a DeleteConstraintTemplateUnauthorized with default headers values
func NewDeleteConstraintTemplateUnauthorized() *DeleteConstraintTemplateUnauthorized {
	return &DeleteConstraintTemplateUnauthorized{}
}

/*
DeleteConstraintTemplateUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteConstraintTemplateUnauthorized struct {
}

// IsSuccess returns true when this delete constraint template unauthorized response has a 2xx status code
func (o *DeleteConstraintTemplateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete constraint template unauthorized response has a 3xx status code
func (o *DeleteConstraintTemplateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete constraint template unauthorized response has a 4xx status code
func (o *DeleteConstraintTemplateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete constraint template unauthorized response has a 5xx status code
func (o *DeleteConstraintTemplateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete constraint template unauthorized response a status code equal to that given
func (o *DeleteConstraintTemplateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteConstraintTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/constrainttemplates/{ct_name}][%d] deleteConstraintTemplateUnauthorized ", 401)
}

func (o *DeleteConstraintTemplateUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/constrainttemplates/{ct_name}][%d] deleteConstraintTemplateUnauthorized ", 401)
}

func (o *DeleteConstraintTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConstraintTemplateForbidden creates a DeleteConstraintTemplateForbidden with default headers values
func NewDeleteConstraintTemplateForbidden() *DeleteConstraintTemplateForbidden {
	return &DeleteConstraintTemplateForbidden{}
}

/*
DeleteConstraintTemplateForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteConstraintTemplateForbidden struct {
}

// IsSuccess returns true when this delete constraint template forbidden response has a 2xx status code
func (o *DeleteConstraintTemplateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete constraint template forbidden response has a 3xx status code
func (o *DeleteConstraintTemplateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete constraint template forbidden response has a 4xx status code
func (o *DeleteConstraintTemplateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete constraint template forbidden response has a 5xx status code
func (o *DeleteConstraintTemplateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete constraint template forbidden response a status code equal to that given
func (o *DeleteConstraintTemplateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteConstraintTemplateForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/constrainttemplates/{ct_name}][%d] deleteConstraintTemplateForbidden ", 403)
}

func (o *DeleteConstraintTemplateForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/constrainttemplates/{ct_name}][%d] deleteConstraintTemplateForbidden ", 403)
}

func (o *DeleteConstraintTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConstraintTemplateDefault creates a DeleteConstraintTemplateDefault with default headers values
func NewDeleteConstraintTemplateDefault(code int) *DeleteConstraintTemplateDefault {
	return &DeleteConstraintTemplateDefault{
		_statusCode: code,
	}
}

/*
DeleteConstraintTemplateDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteConstraintTemplateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete constraint template default response
func (o *DeleteConstraintTemplateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete constraint template default response has a 2xx status code
func (o *DeleteConstraintTemplateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete constraint template default response has a 3xx status code
func (o *DeleteConstraintTemplateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete constraint template default response has a 4xx status code
func (o *DeleteConstraintTemplateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete constraint template default response has a 5xx status code
func (o *DeleteConstraintTemplateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete constraint template default response a status code equal to that given
func (o *DeleteConstraintTemplateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteConstraintTemplateDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/constrainttemplates/{ct_name}][%d] deleteConstraintTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteConstraintTemplateDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v2/constrainttemplates/{ct_name}][%d] deleteConstraintTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteConstraintTemplateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteConstraintTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
