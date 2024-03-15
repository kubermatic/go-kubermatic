// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListProjectAzureVnetsReader is a Reader for the ListProjectAzureVnets structure.
type ListProjectAzureVnetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectAzureVnetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectAzureVnetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectAzureVnetsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectAzureVnetsOK creates a ListProjectAzureVnetsOK with default headers values
func NewListProjectAzureVnetsOK() *ListProjectAzureVnetsOK {
	return &ListProjectAzureVnetsOK{}
}

/*
ListProjectAzureVnetsOK describes a response with status code 200, with default header values.

AzureVirtualNetworksList
*/
type ListProjectAzureVnetsOK struct {
	Payload *models.AzureVirtualNetworksList
}

// IsSuccess returns true when this list project azure vnets o k response has a 2xx status code
func (o *ListProjectAzureVnetsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project azure vnets o k response has a 3xx status code
func (o *ListProjectAzureVnetsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project azure vnets o k response has a 4xx status code
func (o *ListProjectAzureVnetsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project azure vnets o k response has a 5xx status code
func (o *ListProjectAzureVnetsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project azure vnets o k response a status code equal to that given
func (o *ListProjectAzureVnetsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListProjectAzureVnetsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/azure/vnets][%d] listProjectAzureVnetsOK  %+v", 200, o.Payload)
}

func (o *ListProjectAzureVnetsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/azure/vnets][%d] listProjectAzureVnetsOK  %+v", 200, o.Payload)
}

func (o *ListProjectAzureVnetsOK) GetPayload() *models.AzureVirtualNetworksList {
	return o.Payload
}

func (o *ListProjectAzureVnetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureVirtualNetworksList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectAzureVnetsDefault creates a ListProjectAzureVnetsDefault with default headers values
func NewListProjectAzureVnetsDefault(code int) *ListProjectAzureVnetsDefault {
	return &ListProjectAzureVnetsDefault{
		_statusCode: code,
	}
}

/*
ListProjectAzureVnetsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectAzureVnetsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list project azure vnets default response
func (o *ListProjectAzureVnetsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list project azure vnets default response has a 2xx status code
func (o *ListProjectAzureVnetsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project azure vnets default response has a 3xx status code
func (o *ListProjectAzureVnetsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project azure vnets default response has a 4xx status code
func (o *ListProjectAzureVnetsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project azure vnets default response has a 5xx status code
func (o *ListProjectAzureVnetsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project azure vnets default response a status code equal to that given
func (o *ListProjectAzureVnetsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListProjectAzureVnetsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/azure/vnets][%d] listProjectAzureVnets default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectAzureVnetsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/azure/vnets][%d] listProjectAzureVnets default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectAzureVnetsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectAzureVnetsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
