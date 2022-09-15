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

// ListVMwareCloudDirectorCatalogsReader is a Reader for the ListVMwareCloudDirectorCatalogs structure.
type ListVMwareCloudDirectorCatalogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVMwareCloudDirectorCatalogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVMwareCloudDirectorCatalogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListVMwareCloudDirectorCatalogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVMwareCloudDirectorCatalogsOK creates a ListVMwareCloudDirectorCatalogsOK with default headers values
func NewListVMwareCloudDirectorCatalogsOK() *ListVMwareCloudDirectorCatalogsOK {
	return &ListVMwareCloudDirectorCatalogsOK{}
}

/*
ListVMwareCloudDirectorCatalogsOK describes a response with status code 200, with default header values.

VMwareCloudDirectorCatalogList
*/
type ListVMwareCloudDirectorCatalogsOK struct {
	Payload models.VMwareCloudDirectorCatalogList
}

// IsSuccess returns true when this list v mware cloud director catalogs o k response has a 2xx status code
func (o *ListVMwareCloudDirectorCatalogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list v mware cloud director catalogs o k response has a 3xx status code
func (o *ListVMwareCloudDirectorCatalogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list v mware cloud director catalogs o k response has a 4xx status code
func (o *ListVMwareCloudDirectorCatalogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list v mware cloud director catalogs o k response has a 5xx status code
func (o *ListVMwareCloudDirectorCatalogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list v mware cloud director catalogs o k response a status code equal to that given
func (o *ListVMwareCloudDirectorCatalogsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListVMwareCloudDirectorCatalogsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/vmwareclouddirector/{dc}/catalogs][%d] listVMwareCloudDirectorCatalogsOK  %+v", 200, o.Payload)
}

func (o *ListVMwareCloudDirectorCatalogsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/vmwareclouddirector/{dc}/catalogs][%d] listVMwareCloudDirectorCatalogsOK  %+v", 200, o.Payload)
}

func (o *ListVMwareCloudDirectorCatalogsOK) GetPayload() models.VMwareCloudDirectorCatalogList {
	return o.Payload
}

func (o *ListVMwareCloudDirectorCatalogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVMwareCloudDirectorCatalogsDefault creates a ListVMwareCloudDirectorCatalogsDefault with default headers values
func NewListVMwareCloudDirectorCatalogsDefault(code int) *ListVMwareCloudDirectorCatalogsDefault {
	return &ListVMwareCloudDirectorCatalogsDefault{
		_statusCode: code,
	}
}

/*
ListVMwareCloudDirectorCatalogsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListVMwareCloudDirectorCatalogsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list v mware cloud director catalogs default response
func (o *ListVMwareCloudDirectorCatalogsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list v mware cloud director catalogs default response has a 2xx status code
func (o *ListVMwareCloudDirectorCatalogsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list v mware cloud director catalogs default response has a 3xx status code
func (o *ListVMwareCloudDirectorCatalogsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list v mware cloud director catalogs default response has a 4xx status code
func (o *ListVMwareCloudDirectorCatalogsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list v mware cloud director catalogs default response has a 5xx status code
func (o *ListVMwareCloudDirectorCatalogsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list v mware cloud director catalogs default response a status code equal to that given
func (o *ListVMwareCloudDirectorCatalogsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListVMwareCloudDirectorCatalogsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/vmwareclouddirector/{dc}/catalogs][%d] listVMwareCloudDirectorCatalogs default  %+v", o._statusCode, o.Payload)
}

func (o *ListVMwareCloudDirectorCatalogsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/vmwareclouddirector/{dc}/catalogs][%d] listVMwareCloudDirectorCatalogs default  %+v", o._statusCode, o.Payload)
}

func (o *ListVMwareCloudDirectorCatalogsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListVMwareCloudDirectorCatalogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
