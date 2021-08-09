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

// ListGCPSubnetworksReader is a Reader for the ListGCPSubnetworks structure.
type ListGCPSubnetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGCPSubnetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGCPSubnetworksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGCPSubnetworksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGCPSubnetworksOK creates a ListGCPSubnetworksOK with default headers values
func NewListGCPSubnetworksOK() *ListGCPSubnetworksOK {
	return &ListGCPSubnetworksOK{}
}

/* ListGCPSubnetworksOK describes a response with status code 200, with default header values.

GCPSubnetworkList
*/
type ListGCPSubnetworksOK struct {
	Payload models.GCPSubnetworkList
}

func (o *ListGCPSubnetworksOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/{dc}/subnetworks][%d] listGCPSubnetworksOK  %+v", 200, o.Payload)
}
func (o *ListGCPSubnetworksOK) GetPayload() models.GCPSubnetworkList {
	return o.Payload
}

func (o *ListGCPSubnetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGCPSubnetworksDefault creates a ListGCPSubnetworksDefault with default headers values
func NewListGCPSubnetworksDefault(code int) *ListGCPSubnetworksDefault {
	return &ListGCPSubnetworksDefault{
		_statusCode: code,
	}
}

/* ListGCPSubnetworksDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListGCPSubnetworksDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list g c p subnetworks default response
func (o *ListGCPSubnetworksDefault) Code() int {
	return o._statusCode
}

func (o *ListGCPSubnetworksDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/{dc}/subnetworks][%d] listGCPSubnetworks default  %+v", o._statusCode, o.Payload)
}
func (o *ListGCPSubnetworksDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListGCPSubnetworksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
