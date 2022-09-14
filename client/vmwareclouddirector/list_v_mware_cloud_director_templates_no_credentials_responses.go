// Code generated by go-swagger; DO NOT EDIT.

package vmwareclouddirector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListVMwareCloudDirectorTemplatesNoCredentialsReader is a Reader for the ListVMwareCloudDirectorTemplatesNoCredentials structure.
type ListVMwareCloudDirectorTemplatesNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVMwareCloudDirectorTemplatesNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVMwareCloudDirectorTemplatesNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListVMwareCloudDirectorTemplatesNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVMwareCloudDirectorTemplatesNoCredentialsOK creates a ListVMwareCloudDirectorTemplatesNoCredentialsOK with default headers values
func NewListVMwareCloudDirectorTemplatesNoCredentialsOK() *ListVMwareCloudDirectorTemplatesNoCredentialsOK {
	return &ListVMwareCloudDirectorTemplatesNoCredentialsOK{}
}

/* ListVMwareCloudDirectorTemplatesNoCredentialsOK describes a response with status code 200, with default header values.

VMwareCloudDirectorTemplateList
*/
type ListVMwareCloudDirectorTemplatesNoCredentialsOK struct {
	Payload models.VMwareCloudDirectorTemplateList
}

// IsSuccess returns true when this list v mware cloud director templates no credentials o k response has a 2xx status code
func (o *ListVMwareCloudDirectorTemplatesNoCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list v mware cloud director templates no credentials o k response has a 3xx status code
func (o *ListVMwareCloudDirectorTemplatesNoCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list v mware cloud director templates no credentials o k response has a 4xx status code
func (o *ListVMwareCloudDirectorTemplatesNoCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list v mware cloud director templates no credentials o k response has a 5xx status code
func (o *ListVMwareCloudDirectorTemplatesNoCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list v mware cloud director templates no credentials o k response a status code equal to that given
func (o *ListVMwareCloudDirectorTemplatesNoCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListVMwareCloudDirectorTemplatesNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vmwareclouddirector/templates/{catalog_name}][%d] listVMwareCloudDirectorTemplatesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListVMwareCloudDirectorTemplatesNoCredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vmwareclouddirector/templates/{catalog_name}][%d] listVMwareCloudDirectorTemplatesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListVMwareCloudDirectorTemplatesNoCredentialsOK) GetPayload() models.VMwareCloudDirectorTemplateList {
	return o.Payload
}

func (o *ListVMwareCloudDirectorTemplatesNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVMwareCloudDirectorTemplatesNoCredentialsDefault creates a ListVMwareCloudDirectorTemplatesNoCredentialsDefault with default headers values
func NewListVMwareCloudDirectorTemplatesNoCredentialsDefault(code int) *ListVMwareCloudDirectorTemplatesNoCredentialsDefault {
	return &ListVMwareCloudDirectorTemplatesNoCredentialsDefault{
		_statusCode: code,
	}
}

/* ListVMwareCloudDirectorTemplatesNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListVMwareCloudDirectorTemplatesNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list v mware cloud director templates no credentials default response
func (o *ListVMwareCloudDirectorTemplatesNoCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list v mware cloud director templates no credentials default response has a 2xx status code
func (o *ListVMwareCloudDirectorTemplatesNoCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list v mware cloud director templates no credentials default response has a 3xx status code
func (o *ListVMwareCloudDirectorTemplatesNoCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list v mware cloud director templates no credentials default response has a 4xx status code
func (o *ListVMwareCloudDirectorTemplatesNoCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list v mware cloud director templates no credentials default response has a 5xx status code
func (o *ListVMwareCloudDirectorTemplatesNoCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list v mware cloud director templates no credentials default response a status code equal to that given
func (o *ListVMwareCloudDirectorTemplatesNoCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListVMwareCloudDirectorTemplatesNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vmwareclouddirector/templates/{catalog_name}][%d] listVMwareCloudDirectorTemplatesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListVMwareCloudDirectorTemplatesNoCredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vmwareclouddirector/templates/{catalog_name}][%d] listVMwareCloudDirectorTemplatesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListVMwareCloudDirectorTemplatesNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListVMwareCloudDirectorTemplatesNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
