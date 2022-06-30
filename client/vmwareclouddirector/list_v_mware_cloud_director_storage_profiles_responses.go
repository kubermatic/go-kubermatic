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

// ListVMwareCloudDirectorStorageProfilesReader is a Reader for the ListVMwareCloudDirectorStorageProfiles structure.
type ListVMwareCloudDirectorStorageProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVMwareCloudDirectorStorageProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVMwareCloudDirectorStorageProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListVMwareCloudDirectorStorageProfilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVMwareCloudDirectorStorageProfilesOK creates a ListVMwareCloudDirectorStorageProfilesOK with default headers values
func NewListVMwareCloudDirectorStorageProfilesOK() *ListVMwareCloudDirectorStorageProfilesOK {
	return &ListVMwareCloudDirectorStorageProfilesOK{}
}

/* ListVMwareCloudDirectorStorageProfilesOK describes a response with status code 200, with default header values.

VMwareCloudDirectorStorageProfileList
*/
type ListVMwareCloudDirectorStorageProfilesOK struct {
	Payload models.VMwareCloudDirectorStorageProfileList
}

func (o *ListVMwareCloudDirectorStorageProfilesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/vmwareclouddirector/{dc}/storageprofiles][%d] listVMwareCloudDirectorStorageProfilesOK  %+v", 200, o.Payload)
}
func (o *ListVMwareCloudDirectorStorageProfilesOK) GetPayload() models.VMwareCloudDirectorStorageProfileList {
	return o.Payload
}

func (o *ListVMwareCloudDirectorStorageProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVMwareCloudDirectorStorageProfilesDefault creates a ListVMwareCloudDirectorStorageProfilesDefault with default headers values
func NewListVMwareCloudDirectorStorageProfilesDefault(code int) *ListVMwareCloudDirectorStorageProfilesDefault {
	return &ListVMwareCloudDirectorStorageProfilesDefault{
		_statusCode: code,
	}
}

/* ListVMwareCloudDirectorStorageProfilesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListVMwareCloudDirectorStorageProfilesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list v mware cloud director storage profiles default response
func (o *ListVMwareCloudDirectorStorageProfilesDefault) Code() int {
	return o._statusCode
}

func (o *ListVMwareCloudDirectorStorageProfilesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/vmwareclouddirector/{dc}/storageprofiles][%d] listVMwareCloudDirectorStorageProfiles default  %+v", o._statusCode, o.Payload)
}
func (o *ListVMwareCloudDirectorStorageProfilesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListVMwareCloudDirectorStorageProfilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
