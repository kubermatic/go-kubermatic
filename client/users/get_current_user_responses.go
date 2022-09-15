// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// GetCurrentUserReader is a Reader for the GetCurrentUser structure.
type GetCurrentUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCurrentUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCurrentUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCurrentUserOK creates a GetCurrentUserOK with default headers values
func NewGetCurrentUserOK() *GetCurrentUserOK {
	return &GetCurrentUserOK{}
}

/*
GetCurrentUserOK describes a response with status code 200, with default header values.

User
*/
type GetCurrentUserOK struct {
	Payload *models.User
}

// IsSuccess returns true when this get current user o k response has a 2xx status code
func (o *GetCurrentUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get current user o k response has a 3xx status code
func (o *GetCurrentUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current user o k response has a 4xx status code
func (o *GetCurrentUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get current user o k response has a 5xx status code
func (o *GetCurrentUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get current user o k response a status code equal to that given
func (o *GetCurrentUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCurrentUserOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/me][%d] getCurrentUserOK  %+v", 200, o.Payload)
}

func (o *GetCurrentUserOK) String() string {
	return fmt.Sprintf("[GET /api/v1/me][%d] getCurrentUserOK  %+v", 200, o.Payload)
}

func (o *GetCurrentUserOK) GetPayload() *models.User {
	return o.Payload
}

func (o *GetCurrentUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserUnauthorized creates a GetCurrentUserUnauthorized with default headers values
func NewGetCurrentUserUnauthorized() *GetCurrentUserUnauthorized {
	return &GetCurrentUserUnauthorized{}
}

/*
GetCurrentUserUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetCurrentUserUnauthorized struct {
}

// IsSuccess returns true when this get current user unauthorized response has a 2xx status code
func (o *GetCurrentUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get current user unauthorized response has a 3xx status code
func (o *GetCurrentUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current user unauthorized response has a 4xx status code
func (o *GetCurrentUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get current user unauthorized response has a 5xx status code
func (o *GetCurrentUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get current user unauthorized response a status code equal to that given
func (o *GetCurrentUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetCurrentUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/me][%d] getCurrentUserUnauthorized ", 401)
}

func (o *GetCurrentUserUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/me][%d] getCurrentUserUnauthorized ", 401)
}

func (o *GetCurrentUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCurrentUserDefault creates a GetCurrentUserDefault with default headers values
func NewGetCurrentUserDefault(code int) *GetCurrentUserDefault {
	return &GetCurrentUserDefault{
		_statusCode: code,
	}
}

/*
GetCurrentUserDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetCurrentUserDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get current user default response
func (o *GetCurrentUserDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get current user default response has a 2xx status code
func (o *GetCurrentUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get current user default response has a 3xx status code
func (o *GetCurrentUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get current user default response has a 4xx status code
func (o *GetCurrentUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get current user default response has a 5xx status code
func (o *GetCurrentUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get current user default response a status code equal to that given
func (o *GetCurrentUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetCurrentUserDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/me][%d] getCurrentUser default  %+v", o._statusCode, o.Payload)
}

func (o *GetCurrentUserDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/me][%d] getCurrentUser default  %+v", o._statusCode, o.Payload)
}

func (o *GetCurrentUserDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetCurrentUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
