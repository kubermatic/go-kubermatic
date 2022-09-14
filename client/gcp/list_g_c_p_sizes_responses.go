// Code generated by go-swagger; DO NOT EDIT.

package gcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListGCPSizesReader is a Reader for the ListGCPSizes structure.
type ListGCPSizesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGCPSizesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGCPSizesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGCPSizesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGCPSizesOK creates a ListGCPSizesOK with default headers values
func NewListGCPSizesOK() *ListGCPSizesOK {
	return &ListGCPSizesOK{}
}

/* ListGCPSizesOK describes a response with status code 200, with default header values.

GCPMachineSizeList
*/
type ListGCPSizesOK struct {
	Payload models.GCPMachineSizeList
}

// IsSuccess returns true when this list g c p sizes o k response has a 2xx status code
func (o *ListGCPSizesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list g c p sizes o k response has a 3xx status code
func (o *ListGCPSizesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list g c p sizes o k response has a 4xx status code
func (o *ListGCPSizesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list g c p sizes o k response has a 5xx status code
func (o *ListGCPSizesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list g c p sizes o k response a status code equal to that given
func (o *ListGCPSizesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListGCPSizesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/sizes][%d] listGCPSizesOK  %+v", 200, o.Payload)
}

func (o *ListGCPSizesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/sizes][%d] listGCPSizesOK  %+v", 200, o.Payload)
}

func (o *ListGCPSizesOK) GetPayload() models.GCPMachineSizeList {
	return o.Payload
}

func (o *ListGCPSizesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGCPSizesDefault creates a ListGCPSizesDefault with default headers values
func NewListGCPSizesDefault(code int) *ListGCPSizesDefault {
	return &ListGCPSizesDefault{
		_statusCode: code,
	}
}

/* ListGCPSizesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListGCPSizesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list g c p sizes default response
func (o *ListGCPSizesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list g c p sizes default response has a 2xx status code
func (o *ListGCPSizesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list g c p sizes default response has a 3xx status code
func (o *ListGCPSizesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list g c p sizes default response has a 4xx status code
func (o *ListGCPSizesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list g c p sizes default response has a 5xx status code
func (o *ListGCPSizesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list g c p sizes default response a status code equal to that given
func (o *ListGCPSizesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListGCPSizesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/sizes][%d] listGCPSizes default  %+v", o._statusCode, o.Payload)
}

func (o *ListGCPSizesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/sizes][%d] listGCPSizes default  %+v", o._statusCode, o.Payload)
}

func (o *ListGCPSizesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListGCPSizesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
