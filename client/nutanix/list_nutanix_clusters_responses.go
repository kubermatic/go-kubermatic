// Code generated by go-swagger; DO NOT EDIT.

package nutanix

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListNutanixClustersReader is a Reader for the ListNutanixClusters structure.
type ListNutanixClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNutanixClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNutanixClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListNutanixClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListNutanixClustersOK creates a ListNutanixClustersOK with default headers values
func NewListNutanixClustersOK() *ListNutanixClustersOK {
	return &ListNutanixClustersOK{}
}

/*
ListNutanixClustersOK describes a response with status code 200, with default header values.

NutanixClusterList
*/
type ListNutanixClustersOK struct {
	Payload models.NutanixClusterList
}

// IsSuccess returns true when this list nutanix clusters o k response has a 2xx status code
func (o *ListNutanixClustersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list nutanix clusters o k response has a 3xx status code
func (o *ListNutanixClustersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list nutanix clusters o k response has a 4xx status code
func (o *ListNutanixClustersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list nutanix clusters o k response has a 5xx status code
func (o *ListNutanixClustersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list nutanix clusters o k response a status code equal to that given
func (o *ListNutanixClustersOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListNutanixClustersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/nutanix/{dc}/clusters][%d] listNutanixClustersOK  %+v", 200, o.Payload)
}

func (o *ListNutanixClustersOK) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/nutanix/{dc}/clusters][%d] listNutanixClustersOK  %+v", 200, o.Payload)
}

func (o *ListNutanixClustersOK) GetPayload() models.NutanixClusterList {
	return o.Payload
}

func (o *ListNutanixClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNutanixClustersDefault creates a ListNutanixClustersDefault with default headers values
func NewListNutanixClustersDefault(code int) *ListNutanixClustersDefault {
	return &ListNutanixClustersDefault{
		_statusCode: code,
	}
}

/*
ListNutanixClustersDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListNutanixClustersDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list nutanix clusters default response
func (o *ListNutanixClustersDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list nutanix clusters default response has a 2xx status code
func (o *ListNutanixClustersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list nutanix clusters default response has a 3xx status code
func (o *ListNutanixClustersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list nutanix clusters default response has a 4xx status code
func (o *ListNutanixClustersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list nutanix clusters default response has a 5xx status code
func (o *ListNutanixClustersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list nutanix clusters default response a status code equal to that given
func (o *ListNutanixClustersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListNutanixClustersDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/nutanix/{dc}/clusters][%d] listNutanixClusters default  %+v", o._statusCode, o.Payload)
}

func (o *ListNutanixClustersDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/nutanix/{dc}/clusters][%d] listNutanixClusters default  %+v", o._statusCode, o.Payload)
}

func (o *ListNutanixClustersDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListNutanixClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
