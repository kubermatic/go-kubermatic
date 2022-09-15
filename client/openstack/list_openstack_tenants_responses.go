// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListOpenstackTenantsReader is a Reader for the ListOpenstackTenants structure.
type ListOpenstackTenantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOpenstackTenantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOpenstackTenantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListOpenstackTenantsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOpenstackTenantsOK creates a ListOpenstackTenantsOK with default headers values
func NewListOpenstackTenantsOK() *ListOpenstackTenantsOK {
	return &ListOpenstackTenantsOK{}
}

/*
ListOpenstackTenantsOK describes a response with status code 200, with default header values.

OpenstackTenant
*/
type ListOpenstackTenantsOK struct {
	Payload []*models.OpenstackTenant
}

// IsSuccess returns true when this list openstack tenants o k response has a 2xx status code
func (o *ListOpenstackTenantsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list openstack tenants o k response has a 3xx status code
func (o *ListOpenstackTenantsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list openstack tenants o k response has a 4xx status code
func (o *ListOpenstackTenantsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list openstack tenants o k response has a 5xx status code
func (o *ListOpenstackTenantsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list openstack tenants o k response a status code equal to that given
func (o *ListOpenstackTenantsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListOpenstackTenantsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/openstack/tenants][%d] listOpenstackTenantsOK  %+v", 200, o.Payload)
}

func (o *ListOpenstackTenantsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/openstack/tenants][%d] listOpenstackTenantsOK  %+v", 200, o.Payload)
}

func (o *ListOpenstackTenantsOK) GetPayload() []*models.OpenstackTenant {
	return o.Payload
}

func (o *ListOpenstackTenantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOpenstackTenantsDefault creates a ListOpenstackTenantsDefault with default headers values
func NewListOpenstackTenantsDefault(code int) *ListOpenstackTenantsDefault {
	return &ListOpenstackTenantsDefault{
		_statusCode: code,
	}
}

/*
ListOpenstackTenantsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListOpenstackTenantsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list openstack tenants default response
func (o *ListOpenstackTenantsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list openstack tenants default response has a 2xx status code
func (o *ListOpenstackTenantsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list openstack tenants default response has a 3xx status code
func (o *ListOpenstackTenantsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list openstack tenants default response has a 4xx status code
func (o *ListOpenstackTenantsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list openstack tenants default response has a 5xx status code
func (o *ListOpenstackTenantsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list openstack tenants default response a status code equal to that given
func (o *ListOpenstackTenantsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListOpenstackTenantsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/openstack/tenants][%d] listOpenstackTenants default  %+v", o._statusCode, o.Payload)
}

func (o *ListOpenstackTenantsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/openstack/tenants][%d] listOpenstackTenants default  %+v", o._statusCode, o.Payload)
}

func (o *ListOpenstackTenantsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListOpenstackTenantsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
