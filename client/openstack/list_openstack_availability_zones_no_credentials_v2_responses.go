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

// ListOpenstackAvailabilityZonesNoCredentialsV2Reader is a Reader for the ListOpenstackAvailabilityZonesNoCredentialsV2 structure.
type ListOpenstackAvailabilityZonesNoCredentialsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOpenstackAvailabilityZonesNoCredentialsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOpenstackAvailabilityZonesNoCredentialsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListOpenstackAvailabilityZonesNoCredentialsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOpenstackAvailabilityZonesNoCredentialsV2OK creates a ListOpenstackAvailabilityZonesNoCredentialsV2OK with default headers values
func NewListOpenstackAvailabilityZonesNoCredentialsV2OK() *ListOpenstackAvailabilityZonesNoCredentialsV2OK {
	return &ListOpenstackAvailabilityZonesNoCredentialsV2OK{}
}

/* ListOpenstackAvailabilityZonesNoCredentialsV2OK describes a response with status code 200, with default header values.

OpenstackAvailabilityZone
*/
type ListOpenstackAvailabilityZonesNoCredentialsV2OK struct {
	Payload []*models.OpenstackAvailabilityZone
}

// IsSuccess returns true when this list openstack availability zones no credentials v2 o k response has a 2xx status code
func (o *ListOpenstackAvailabilityZonesNoCredentialsV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list openstack availability zones no credentials v2 o k response has a 3xx status code
func (o *ListOpenstackAvailabilityZonesNoCredentialsV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list openstack availability zones no credentials v2 o k response has a 4xx status code
func (o *ListOpenstackAvailabilityZonesNoCredentialsV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list openstack availability zones no credentials v2 o k response has a 5xx status code
func (o *ListOpenstackAvailabilityZonesNoCredentialsV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list openstack availability zones no credentials v2 o k response a status code equal to that given
func (o *ListOpenstackAvailabilityZonesNoCredentialsV2OK) IsCode(code int) bool {
	return code == 200
}

func (o *ListOpenstackAvailabilityZonesNoCredentialsV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/availabilityzones][%d] listOpenstackAvailabilityZonesNoCredentialsV2OK  %+v", 200, o.Payload)
}

func (o *ListOpenstackAvailabilityZonesNoCredentialsV2OK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/availabilityzones][%d] listOpenstackAvailabilityZonesNoCredentialsV2OK  %+v", 200, o.Payload)
}

func (o *ListOpenstackAvailabilityZonesNoCredentialsV2OK) GetPayload() []*models.OpenstackAvailabilityZone {
	return o.Payload
}

func (o *ListOpenstackAvailabilityZonesNoCredentialsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOpenstackAvailabilityZonesNoCredentialsV2Default creates a ListOpenstackAvailabilityZonesNoCredentialsV2Default with default headers values
func NewListOpenstackAvailabilityZonesNoCredentialsV2Default(code int) *ListOpenstackAvailabilityZonesNoCredentialsV2Default {
	return &ListOpenstackAvailabilityZonesNoCredentialsV2Default{
		_statusCode: code,
	}
}

/* ListOpenstackAvailabilityZonesNoCredentialsV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type ListOpenstackAvailabilityZonesNoCredentialsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list openstack availability zones no credentials v2 default response
func (o *ListOpenstackAvailabilityZonesNoCredentialsV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list openstack availability zones no credentials v2 default response has a 2xx status code
func (o *ListOpenstackAvailabilityZonesNoCredentialsV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list openstack availability zones no credentials v2 default response has a 3xx status code
func (o *ListOpenstackAvailabilityZonesNoCredentialsV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list openstack availability zones no credentials v2 default response has a 4xx status code
func (o *ListOpenstackAvailabilityZonesNoCredentialsV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list openstack availability zones no credentials v2 default response has a 5xx status code
func (o *ListOpenstackAvailabilityZonesNoCredentialsV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list openstack availability zones no credentials v2 default response a status code equal to that given
func (o *ListOpenstackAvailabilityZonesNoCredentialsV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListOpenstackAvailabilityZonesNoCredentialsV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/availabilityzones][%d] listOpenstackAvailabilityZonesNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListOpenstackAvailabilityZonesNoCredentialsV2Default) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/availabilityzones][%d] listOpenstackAvailabilityZonesNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListOpenstackAvailabilityZonesNoCredentialsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListOpenstackAvailabilityZonesNoCredentialsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
