// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListAWSSizesReader is a Reader for the ListAWSSizes structure.
type ListAWSSizesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAWSSizesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAWSSizesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAWSSizesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAWSSizesOK creates a ListAWSSizesOK with default headers values
func NewListAWSSizesOK() *ListAWSSizesOK {
	return &ListAWSSizesOK{}
}

/*
ListAWSSizesOK describes a response with status code 200, with default header values.

AWSSizeList
*/
type ListAWSSizesOK struct {
	Payload models.AWSSizeList
}

// IsSuccess returns true when this list a w s sizes o k response has a 2xx status code
func (o *ListAWSSizesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list a w s sizes o k response has a 3xx status code
func (o *ListAWSSizesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list a w s sizes o k response has a 4xx status code
func (o *ListAWSSizesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list a w s sizes o k response has a 5xx status code
func (o *ListAWSSizesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list a w s sizes o k response a status code equal to that given
func (o *ListAWSSizesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListAWSSizesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/aws/sizes][%d] listAWSSizesOK  %+v", 200, o.Payload)
}

func (o *ListAWSSizesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/aws/sizes][%d] listAWSSizesOK  %+v", 200, o.Payload)
}

func (o *ListAWSSizesOK) GetPayload() models.AWSSizeList {
	return o.Payload
}

func (o *ListAWSSizesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAWSSizesDefault creates a ListAWSSizesDefault with default headers values
func NewListAWSSizesDefault(code int) *ListAWSSizesDefault {
	return &ListAWSSizesDefault{
		_statusCode: code,
	}
}

/*
ListAWSSizesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListAWSSizesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list a w s sizes default response
func (o *ListAWSSizesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list a w s sizes default response has a 2xx status code
func (o *ListAWSSizesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list a w s sizes default response has a 3xx status code
func (o *ListAWSSizesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list a w s sizes default response has a 4xx status code
func (o *ListAWSSizesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list a w s sizes default response has a 5xx status code
func (o *ListAWSSizesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list a w s sizes default response a status code equal to that given
func (o *ListAWSSizesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListAWSSizesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/aws/sizes][%d] listAWSSizes default  %+v", o._statusCode, o.Payload)
}

func (o *ListAWSSizesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/aws/sizes][%d] listAWSSizes default  %+v", o._statusCode, o.Payload)
}

func (o *ListAWSSizesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAWSSizesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
