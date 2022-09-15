// Code generated by go-swagger; DO NOT EDIT.

package gke

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ValidateGKECredentialsReader is a Reader for the ValidateGKECredentials structure.
type ValidateGKECredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateGKECredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateGKECredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewValidateGKECredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewValidateGKECredentialsOK creates a ValidateGKECredentialsOK with default headers values
func NewValidateGKECredentialsOK() *ValidateGKECredentialsOK {
	return &ValidateGKECredentialsOK{}
}

/*
ValidateGKECredentialsOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type ValidateGKECredentialsOK struct {
}

// IsSuccess returns true when this validate g k e credentials o k response has a 2xx status code
func (o *ValidateGKECredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate g k e credentials o k response has a 3xx status code
func (o *ValidateGKECredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate g k e credentials o k response has a 4xx status code
func (o *ValidateGKECredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate g k e credentials o k response has a 5xx status code
func (o *ValidateGKECredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this validate g k e credentials o k response a status code equal to that given
func (o *ValidateGKECredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ValidateGKECredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/gke/validatecredentials][%d] validateGKECredentialsOK ", 200)
}

func (o *ValidateGKECredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/gke/validatecredentials][%d] validateGKECredentialsOK ", 200)
}

func (o *ValidateGKECredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateGKECredentialsDefault creates a ValidateGKECredentialsDefault with default headers values
func NewValidateGKECredentialsDefault(code int) *ValidateGKECredentialsDefault {
	return &ValidateGKECredentialsDefault{
		_statusCode: code,
	}
}

/*
ValidateGKECredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ValidateGKECredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the validate g k e credentials default response
func (o *ValidateGKECredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this validate g k e credentials default response has a 2xx status code
func (o *ValidateGKECredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this validate g k e credentials default response has a 3xx status code
func (o *ValidateGKECredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this validate g k e credentials default response has a 4xx status code
func (o *ValidateGKECredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this validate g k e credentials default response has a 5xx status code
func (o *ValidateGKECredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this validate g k e credentials default response a status code equal to that given
func (o *ValidateGKECredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ValidateGKECredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/gke/validatecredentials][%d] validateGKECredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ValidateGKECredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/gke/validatecredentials][%d] validateGKECredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ValidateGKECredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ValidateGKECredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
