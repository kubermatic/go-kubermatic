// Code generated by go-swagger; DO NOT EDIT.

package vmwareclouddirector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListProjectVMwareCloudDirectorNetworksReader is a Reader for the ListProjectVMwareCloudDirectorNetworks structure.
type ListProjectVMwareCloudDirectorNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectVMwareCloudDirectorNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectVMwareCloudDirectorNetworksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectVMwareCloudDirectorNetworksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectVMwareCloudDirectorNetworksOK creates a ListProjectVMwareCloudDirectorNetworksOK with default headers values
func NewListProjectVMwareCloudDirectorNetworksOK() *ListProjectVMwareCloudDirectorNetworksOK {
	return &ListProjectVMwareCloudDirectorNetworksOK{}
}

/*
ListProjectVMwareCloudDirectorNetworksOK describes a response with status code 200, with default header values.

VMwareCloudDirectorNetworkList
*/
type ListProjectVMwareCloudDirectorNetworksOK struct {
	Payload models.VMwareCloudDirectorNetworkList
}

// IsSuccess returns true when this list project v mware cloud director networks o k response has a 2xx status code
func (o *ListProjectVMwareCloudDirectorNetworksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project v mware cloud director networks o k response has a 3xx status code
func (o *ListProjectVMwareCloudDirectorNetworksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project v mware cloud director networks o k response has a 4xx status code
func (o *ListProjectVMwareCloudDirectorNetworksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project v mware cloud director networks o k response has a 5xx status code
func (o *ListProjectVMwareCloudDirectorNetworksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project v mware cloud director networks o k response a status code equal to that given
func (o *ListProjectVMwareCloudDirectorNetworksOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListProjectVMwareCloudDirectorNetworksOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/vmwareclouddirector/{dc}/networks][%d] listProjectVMwareCloudDirectorNetworksOK  %+v", 200, o.Payload)
}

func (o *ListProjectVMwareCloudDirectorNetworksOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/vmwareclouddirector/{dc}/networks][%d] listProjectVMwareCloudDirectorNetworksOK  %+v", 200, o.Payload)
}

func (o *ListProjectVMwareCloudDirectorNetworksOK) GetPayload() models.VMwareCloudDirectorNetworkList {
	return o.Payload
}

func (o *ListProjectVMwareCloudDirectorNetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectVMwareCloudDirectorNetworksDefault creates a ListProjectVMwareCloudDirectorNetworksDefault with default headers values
func NewListProjectVMwareCloudDirectorNetworksDefault(code int) *ListProjectVMwareCloudDirectorNetworksDefault {
	return &ListProjectVMwareCloudDirectorNetworksDefault{
		_statusCode: code,
	}
}

/*
ListProjectVMwareCloudDirectorNetworksDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectVMwareCloudDirectorNetworksDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list project v mware cloud director networks default response
func (o *ListProjectVMwareCloudDirectorNetworksDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list project v mware cloud director networks default response has a 2xx status code
func (o *ListProjectVMwareCloudDirectorNetworksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project v mware cloud director networks default response has a 3xx status code
func (o *ListProjectVMwareCloudDirectorNetworksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project v mware cloud director networks default response has a 4xx status code
func (o *ListProjectVMwareCloudDirectorNetworksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project v mware cloud director networks default response has a 5xx status code
func (o *ListProjectVMwareCloudDirectorNetworksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project v mware cloud director networks default response a status code equal to that given
func (o *ListProjectVMwareCloudDirectorNetworksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListProjectVMwareCloudDirectorNetworksDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/vmwareclouddirector/{dc}/networks][%d] listProjectVMwareCloudDirectorNetworks default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectVMwareCloudDirectorNetworksDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/vmwareclouddirector/{dc}/networks][%d] listProjectVMwareCloudDirectorNetworks default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectVMwareCloudDirectorNetworksDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectVMwareCloudDirectorNetworksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
