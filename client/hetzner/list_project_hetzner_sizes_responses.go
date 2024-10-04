// Code generated by go-swagger; DO NOT EDIT.

package hetzner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListProjectHetznerSizesReader is a Reader for the ListProjectHetznerSizes structure.
type ListProjectHetznerSizesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectHetznerSizesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectHetznerSizesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectHetznerSizesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectHetznerSizesOK creates a ListProjectHetznerSizesOK with default headers values
func NewListProjectHetznerSizesOK() *ListProjectHetznerSizesOK {
	return &ListProjectHetznerSizesOK{}
}

/*
ListProjectHetznerSizesOK describes a response with status code 200, with default header values.

HetznerSizeList
*/
type ListProjectHetznerSizesOK struct {
	Payload *models.HetznerSizeList
}

// IsSuccess returns true when this list project hetzner sizes o k response has a 2xx status code
func (o *ListProjectHetznerSizesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project hetzner sizes o k response has a 3xx status code
func (o *ListProjectHetznerSizesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project hetzner sizes o k response has a 4xx status code
func (o *ListProjectHetznerSizesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project hetzner sizes o k response has a 5xx status code
func (o *ListProjectHetznerSizesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project hetzner sizes o k response a status code equal to that given
func (o *ListProjectHetznerSizesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListProjectHetznerSizesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/hetzner/sizes][%d] listProjectHetznerSizesOK  %+v", 200, o.Payload)
}

func (o *ListProjectHetznerSizesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/hetzner/sizes][%d] listProjectHetznerSizesOK  %+v", 200, o.Payload)
}

func (o *ListProjectHetznerSizesOK) GetPayload() *models.HetznerSizeList {
	return o.Payload
}

func (o *ListProjectHetznerSizesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HetznerSizeList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectHetznerSizesDefault creates a ListProjectHetznerSizesDefault with default headers values
func NewListProjectHetznerSizesDefault(code int) *ListProjectHetznerSizesDefault {
	return &ListProjectHetznerSizesDefault{
		_statusCode: code,
	}
}

/*
ListProjectHetznerSizesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectHetznerSizesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list project hetzner sizes default response
func (o *ListProjectHetznerSizesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list project hetzner sizes default response has a 2xx status code
func (o *ListProjectHetznerSizesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project hetzner sizes default response has a 3xx status code
func (o *ListProjectHetznerSizesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project hetzner sizes default response has a 4xx status code
func (o *ListProjectHetznerSizesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project hetzner sizes default response has a 5xx status code
func (o *ListProjectHetznerSizesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project hetzner sizes default response a status code equal to that given
func (o *ListProjectHetznerSizesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListProjectHetznerSizesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/hetzner/sizes][%d] listProjectHetznerSizes default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectHetznerSizesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/hetzner/sizes][%d] listProjectHetznerSizes default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectHetznerSizesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectHetznerSizesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
