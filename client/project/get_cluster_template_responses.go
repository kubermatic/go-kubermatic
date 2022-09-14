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

// GetClusterTemplateReader is a Reader for the GetClusterTemplate structure.
type GetClusterTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClusterTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetClusterTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClusterTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterTemplateOK creates a GetClusterTemplateOK with default headers values
func NewGetClusterTemplateOK() *GetClusterTemplateOK {
	return &GetClusterTemplateOK{}
}

/* GetClusterTemplateOK describes a response with status code 200, with default header values.

ClusterTemplate
*/
type GetClusterTemplateOK struct {
	Payload *models.ClusterTemplate
}

// IsSuccess returns true when this get cluster template o k response has a 2xx status code
func (o *GetClusterTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster template o k response has a 3xx status code
func (o *GetClusterTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster template o k response has a 4xx status code
func (o *GetClusterTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster template o k response has a 5xx status code
func (o *GetClusterTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster template o k response a status code equal to that given
func (o *GetClusterTemplateOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetClusterTemplateOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clustertemplates/{template_id}][%d] getClusterTemplateOK  %+v", 200, o.Payload)
}

func (o *GetClusterTemplateOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clustertemplates/{template_id}][%d] getClusterTemplateOK  %+v", 200, o.Payload)
}

func (o *GetClusterTemplateOK) GetPayload() *models.ClusterTemplate {
	return o.Payload
}

func (o *GetClusterTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterTemplateUnauthorized creates a GetClusterTemplateUnauthorized with default headers values
func NewGetClusterTemplateUnauthorized() *GetClusterTemplateUnauthorized {
	return &GetClusterTemplateUnauthorized{}
}

/* GetClusterTemplateUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetClusterTemplateUnauthorized struct {
}

// IsSuccess returns true when this get cluster template unauthorized response has a 2xx status code
func (o *GetClusterTemplateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster template unauthorized response has a 3xx status code
func (o *GetClusterTemplateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster template unauthorized response has a 4xx status code
func (o *GetClusterTemplateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster template unauthorized response has a 5xx status code
func (o *GetClusterTemplateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster template unauthorized response a status code equal to that given
func (o *GetClusterTemplateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetClusterTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clustertemplates/{template_id}][%d] getClusterTemplateUnauthorized ", 401)
}

func (o *GetClusterTemplateUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clustertemplates/{template_id}][%d] getClusterTemplateUnauthorized ", 401)
}

func (o *GetClusterTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterTemplateForbidden creates a GetClusterTemplateForbidden with default headers values
func NewGetClusterTemplateForbidden() *GetClusterTemplateForbidden {
	return &GetClusterTemplateForbidden{}
}

/* GetClusterTemplateForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetClusterTemplateForbidden struct {
}

// IsSuccess returns true when this get cluster template forbidden response has a 2xx status code
func (o *GetClusterTemplateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster template forbidden response has a 3xx status code
func (o *GetClusterTemplateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster template forbidden response has a 4xx status code
func (o *GetClusterTemplateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster template forbidden response has a 5xx status code
func (o *GetClusterTemplateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster template forbidden response a status code equal to that given
func (o *GetClusterTemplateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetClusterTemplateForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clustertemplates/{template_id}][%d] getClusterTemplateForbidden ", 403)
}

func (o *GetClusterTemplateForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clustertemplates/{template_id}][%d] getClusterTemplateForbidden ", 403)
}

func (o *GetClusterTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterTemplateDefault creates a GetClusterTemplateDefault with default headers values
func NewGetClusterTemplateDefault(code int) *GetClusterTemplateDefault {
	return &GetClusterTemplateDefault{
		_statusCode: code,
	}
}

/* GetClusterTemplateDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetClusterTemplateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster template default response
func (o *GetClusterTemplateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get cluster template default response has a 2xx status code
func (o *GetClusterTemplateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cluster template default response has a 3xx status code
func (o *GetClusterTemplateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cluster template default response has a 4xx status code
func (o *GetClusterTemplateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cluster template default response has a 5xx status code
func (o *GetClusterTemplateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cluster template default response a status code equal to that given
func (o *GetClusterTemplateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetClusterTemplateDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clustertemplates/{template_id}][%d] getClusterTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterTemplateDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clustertemplates/{template_id}][%d] getClusterTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterTemplateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
