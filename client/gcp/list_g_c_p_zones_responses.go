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

// ListGCPZonesReader is a Reader for the ListGCPZones structure.
type ListGCPZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGCPZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGCPZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGCPZonesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGCPZonesOK creates a ListGCPZonesOK with default headers values
func NewListGCPZonesOK() *ListGCPZonesOK {
	return &ListGCPZonesOK{}
}

/* ListGCPZonesOK describes a response with status code 200, with default header values.

GCPZoneList
*/
type ListGCPZonesOK struct {
	Payload models.GCPZoneList
}

func (o *ListGCPZonesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/{dc}/zones][%d] listGCPZonesOK  %+v", 200, o.Payload)
}
func (o *ListGCPZonesOK) GetPayload() models.GCPZoneList {
	return o.Payload
}

func (o *ListGCPZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGCPZonesDefault creates a ListGCPZonesDefault with default headers values
func NewListGCPZonesDefault(code int) *ListGCPZonesDefault {
	return &ListGCPZonesDefault{
		_statusCode: code,
	}
}

/* ListGCPZonesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListGCPZonesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list g c p zones default response
func (o *ListGCPZonesDefault) Code() int {
	return o._statusCode
}

func (o *ListGCPZonesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/{dc}/zones][%d] listGCPZones default  %+v", o._statusCode, o.Payload)
}
func (o *ListGCPZonesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListGCPZonesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
