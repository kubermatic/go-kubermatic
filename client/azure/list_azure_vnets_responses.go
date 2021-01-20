// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListAzureVnetsReader is a Reader for the ListAzureVnets structure.
type ListAzureVnetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAzureVnetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAzureVnetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAzureVnetsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAzureVnetsOK creates a ListAzureVnetsOK with default headers values
func NewListAzureVnetsOK() *ListAzureVnetsOK {
	return &ListAzureVnetsOK{}
}

/*ListAzureVnetsOK handles this case with default header values.

AzureVirtualNetworksList
*/
type ListAzureVnetsOK struct {
	Payload *models.AzureVirtualNetworksList
}

func (o *ListAzureVnetsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/azure/vnets][%d] listAzureVnetsOK  %+v", 200, o.Payload)
}

func (o *ListAzureVnetsOK) GetPayload() *models.AzureVirtualNetworksList {
	return o.Payload
}

func (o *ListAzureVnetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureVirtualNetworksList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAzureVnetsDefault creates a ListAzureVnetsDefault with default headers values
func NewListAzureVnetsDefault(code int) *ListAzureVnetsDefault {
	return &ListAzureVnetsDefault{
		_statusCode: code,
	}
}

/*ListAzureVnetsDefault handles this case with default header values.

errorResponse
*/
type ListAzureVnetsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list azure vnets default response
func (o *ListAzureVnetsDefault) Code() int {
	return o._statusCode
}

func (o *ListAzureVnetsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/azure/vnets][%d] listAzureVnets default  %+v", o._statusCode, o.Payload)
}

func (o *ListAzureVnetsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAzureVnetsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}