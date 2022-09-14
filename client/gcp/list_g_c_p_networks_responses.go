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

// ListGCPNetworksReader is a Reader for the ListGCPNetworks structure.
type ListGCPNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGCPNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGCPNetworksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGCPNetworksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGCPNetworksOK creates a ListGCPNetworksOK with default headers values
func NewListGCPNetworksOK() *ListGCPNetworksOK {
	return &ListGCPNetworksOK{}
}

/* ListGCPNetworksOK describes a response with status code 200, with default header values.

GCPNetworkList
*/
type ListGCPNetworksOK struct {
	Payload models.GCPNetworkList
}

// IsSuccess returns true when this list g c p networks o k response has a 2xx status code
func (o *ListGCPNetworksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list g c p networks o k response has a 3xx status code
func (o *ListGCPNetworksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list g c p networks o k response has a 4xx status code
func (o *ListGCPNetworksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list g c p networks o k response has a 5xx status code
func (o *ListGCPNetworksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list g c p networks o k response a status code equal to that given
func (o *ListGCPNetworksOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListGCPNetworksOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/networks][%d] listGCPNetworksOK  %+v", 200, o.Payload)
}

func (o *ListGCPNetworksOK) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/networks][%d] listGCPNetworksOK  %+v", 200, o.Payload)
}

func (o *ListGCPNetworksOK) GetPayload() models.GCPNetworkList {
	return o.Payload
}

func (o *ListGCPNetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGCPNetworksDefault creates a ListGCPNetworksDefault with default headers values
func NewListGCPNetworksDefault(code int) *ListGCPNetworksDefault {
	return &ListGCPNetworksDefault{
		_statusCode: code,
	}
}

/* ListGCPNetworksDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListGCPNetworksDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list g c p networks default response
func (o *ListGCPNetworksDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list g c p networks default response has a 2xx status code
func (o *ListGCPNetworksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list g c p networks default response has a 3xx status code
func (o *ListGCPNetworksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list g c p networks default response has a 4xx status code
func (o *ListGCPNetworksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list g c p networks default response has a 5xx status code
func (o *ListGCPNetworksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list g c p networks default response a status code equal to that given
func (o *ListGCPNetworksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListGCPNetworksDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/networks][%d] listGCPNetworks default  %+v", o._statusCode, o.Payload)
}

func (o *ListGCPNetworksDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/networks][%d] listGCPNetworks default  %+v", o._statusCode, o.Payload)
}

func (o *ListGCPNetworksDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListGCPNetworksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
