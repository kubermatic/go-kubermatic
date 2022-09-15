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

// GetIPAMPoolReader is a Reader for the GetIPAMPool structure.
type GetIPAMPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPAMPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIPAMPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetIPAMPoolUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIPAMPoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetIPAMPoolDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIPAMPoolOK creates a GetIPAMPoolOK with default headers values
func NewGetIPAMPoolOK() *GetIPAMPoolOK {
	return &GetIPAMPoolOK{}
}

/*
GetIPAMPoolOK describes a response with status code 200, with default header values.

IPAMPool
*/
type GetIPAMPoolOK struct {
	Payload *models.IPAMPool
}

// IsSuccess returns true when this get Ip a m pool o k response has a 2xx status code
func (o *GetIPAMPoolOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Ip a m pool o k response has a 3xx status code
func (o *GetIPAMPoolOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Ip a m pool o k response has a 4xx status code
func (o *GetIPAMPoolOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Ip a m pool o k response has a 5xx status code
func (o *GetIPAMPoolOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Ip a m pool o k response a status code equal to that given
func (o *GetIPAMPoolOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetIPAMPoolOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/ipampools/{ipampool_name}][%d] getIpAMPoolOK  %+v", 200, o.Payload)
}

func (o *GetIPAMPoolOK) String() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/ipampools/{ipampool_name}][%d] getIpAMPoolOK  %+v", 200, o.Payload)
}

func (o *GetIPAMPoolOK) GetPayload() *models.IPAMPool {
	return o.Payload
}

func (o *GetIPAMPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAMPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPAMPoolUnauthorized creates a GetIPAMPoolUnauthorized with default headers values
func NewGetIPAMPoolUnauthorized() *GetIPAMPoolUnauthorized {
	return &GetIPAMPoolUnauthorized{}
}

/*
GetIPAMPoolUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetIPAMPoolUnauthorized struct {
}

// IsSuccess returns true when this get Ip a m pool unauthorized response has a 2xx status code
func (o *GetIPAMPoolUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Ip a m pool unauthorized response has a 3xx status code
func (o *GetIPAMPoolUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Ip a m pool unauthorized response has a 4xx status code
func (o *GetIPAMPoolUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Ip a m pool unauthorized response has a 5xx status code
func (o *GetIPAMPoolUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get Ip a m pool unauthorized response a status code equal to that given
func (o *GetIPAMPoolUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetIPAMPoolUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/ipampools/{ipampool_name}][%d] getIpAMPoolUnauthorized ", 401)
}

func (o *GetIPAMPoolUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/ipampools/{ipampool_name}][%d] getIpAMPoolUnauthorized ", 401)
}

func (o *GetIPAMPoolUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMPoolForbidden creates a GetIPAMPoolForbidden with default headers values
func NewGetIPAMPoolForbidden() *GetIPAMPoolForbidden {
	return &GetIPAMPoolForbidden{}
}

/*
GetIPAMPoolForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetIPAMPoolForbidden struct {
}

// IsSuccess returns true when this get Ip a m pool forbidden response has a 2xx status code
func (o *GetIPAMPoolForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Ip a m pool forbidden response has a 3xx status code
func (o *GetIPAMPoolForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Ip a m pool forbidden response has a 4xx status code
func (o *GetIPAMPoolForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Ip a m pool forbidden response has a 5xx status code
func (o *GetIPAMPoolForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get Ip a m pool forbidden response a status code equal to that given
func (o *GetIPAMPoolForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetIPAMPoolForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/ipampools/{ipampool_name}][%d] getIpAMPoolForbidden ", 403)
}

func (o *GetIPAMPoolForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/ipampools/{ipampool_name}][%d] getIpAMPoolForbidden ", 403)
}

func (o *GetIPAMPoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMPoolDefault creates a GetIPAMPoolDefault with default headers values
func NewGetIPAMPoolDefault(code int) *GetIPAMPoolDefault {
	return &GetIPAMPoolDefault{
		_statusCode: code,
	}
}

/*
GetIPAMPoolDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetIPAMPoolDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get IP a m pool default response
func (o *GetIPAMPoolDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get IP a m pool default response has a 2xx status code
func (o *GetIPAMPoolDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get IP a m pool default response has a 3xx status code
func (o *GetIPAMPoolDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get IP a m pool default response has a 4xx status code
func (o *GetIPAMPoolDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get IP a m pool default response has a 5xx status code
func (o *GetIPAMPoolDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get IP a m pool default response a status code equal to that given
func (o *GetIPAMPoolDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetIPAMPoolDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/ipampools/{ipampool_name}][%d] getIPAMPool default  %+v", o._statusCode, o.Payload)
}

func (o *GetIPAMPoolDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/ipampools/{ipampool_name}][%d] getIPAMPool default  %+v", o._statusCode, o.Payload)
}

func (o *GetIPAMPoolDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetIPAMPoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
