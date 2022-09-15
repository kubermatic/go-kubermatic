// Code generated by go-swagger; DO NOT EDIT.

package ipampool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// DeleteIPAMPoolReader is a Reader for the DeleteIPAMPool structure.
type DeleteIPAMPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIPAMPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIPAMPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteIPAMPoolUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteIPAMPoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteIPAMPoolDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteIPAMPoolOK creates a DeleteIPAMPoolOK with default headers values
func NewDeleteIPAMPoolOK() *DeleteIPAMPoolOK {
	return &DeleteIPAMPoolOK{}
}

/*
DeleteIPAMPoolOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteIPAMPoolOK struct {
}

// IsSuccess returns true when this delete Ip a m pool o k response has a 2xx status code
func (o *DeleteIPAMPoolOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Ip a m pool o k response has a 3xx status code
func (o *DeleteIPAMPoolOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Ip a m pool o k response has a 4xx status code
func (o *DeleteIPAMPoolOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Ip a m pool o k response has a 5xx status code
func (o *DeleteIPAMPoolOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Ip a m pool o k response a status code equal to that given
func (o *DeleteIPAMPoolOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteIPAMPoolOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/seeds/{seed_name}/ipampools/{ipampool_name}][%d] deleteIpAMPoolOK ", 200)
}

func (o *DeleteIPAMPoolOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/seeds/{seed_name}/ipampools/{ipampool_name}][%d] deleteIpAMPoolOK ", 200)
}

func (o *DeleteIPAMPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIPAMPoolUnauthorized creates a DeleteIPAMPoolUnauthorized with default headers values
func NewDeleteIPAMPoolUnauthorized() *DeleteIPAMPoolUnauthorized {
	return &DeleteIPAMPoolUnauthorized{}
}

/*
DeleteIPAMPoolUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteIPAMPoolUnauthorized struct {
}

// IsSuccess returns true when this delete Ip a m pool unauthorized response has a 2xx status code
func (o *DeleteIPAMPoolUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Ip a m pool unauthorized response has a 3xx status code
func (o *DeleteIPAMPoolUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Ip a m pool unauthorized response has a 4xx status code
func (o *DeleteIPAMPoolUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Ip a m pool unauthorized response has a 5xx status code
func (o *DeleteIPAMPoolUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Ip a m pool unauthorized response a status code equal to that given
func (o *DeleteIPAMPoolUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteIPAMPoolUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/seeds/{seed_name}/ipampools/{ipampool_name}][%d] deleteIpAMPoolUnauthorized ", 401)
}

func (o *DeleteIPAMPoolUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/seeds/{seed_name}/ipampools/{ipampool_name}][%d] deleteIpAMPoolUnauthorized ", 401)
}

func (o *DeleteIPAMPoolUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIPAMPoolForbidden creates a DeleteIPAMPoolForbidden with default headers values
func NewDeleteIPAMPoolForbidden() *DeleteIPAMPoolForbidden {
	return &DeleteIPAMPoolForbidden{}
}

/*
DeleteIPAMPoolForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteIPAMPoolForbidden struct {
}

// IsSuccess returns true when this delete Ip a m pool forbidden response has a 2xx status code
func (o *DeleteIPAMPoolForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Ip a m pool forbidden response has a 3xx status code
func (o *DeleteIPAMPoolForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Ip a m pool forbidden response has a 4xx status code
func (o *DeleteIPAMPoolForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Ip a m pool forbidden response has a 5xx status code
func (o *DeleteIPAMPoolForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Ip a m pool forbidden response a status code equal to that given
func (o *DeleteIPAMPoolForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteIPAMPoolForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/seeds/{seed_name}/ipampools/{ipampool_name}][%d] deleteIpAMPoolForbidden ", 403)
}

func (o *DeleteIPAMPoolForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/seeds/{seed_name}/ipampools/{ipampool_name}][%d] deleteIpAMPoolForbidden ", 403)
}

func (o *DeleteIPAMPoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIPAMPoolDefault creates a DeleteIPAMPoolDefault with default headers values
func NewDeleteIPAMPoolDefault(code int) *DeleteIPAMPoolDefault {
	return &DeleteIPAMPoolDefault{
		_statusCode: code,
	}
}

/*
DeleteIPAMPoolDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteIPAMPoolDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete IP a m pool default response
func (o *DeleteIPAMPoolDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete IP a m pool default response has a 2xx status code
func (o *DeleteIPAMPoolDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete IP a m pool default response has a 3xx status code
func (o *DeleteIPAMPoolDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete IP a m pool default response has a 4xx status code
func (o *DeleteIPAMPoolDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete IP a m pool default response has a 5xx status code
func (o *DeleteIPAMPoolDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete IP a m pool default response a status code equal to that given
func (o *DeleteIPAMPoolDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteIPAMPoolDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/seeds/{seed_name}/ipampools/{ipampool_name}][%d] deleteIPAMPool default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteIPAMPoolDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v2/seeds/{seed_name}/ipampools/{ipampool_name}][%d] deleteIPAMPool default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteIPAMPoolDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteIPAMPoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
