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

// ListAzureSizesReader is a Reader for the ListAzureSizes structure.
type ListAzureSizesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAzureSizesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAzureSizesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAzureSizesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAzureSizesOK creates a ListAzureSizesOK with default headers values
func NewListAzureSizesOK() *ListAzureSizesOK {
	return &ListAzureSizesOK{}
}

/* ListAzureSizesOK describes a response with status code 200, with default header values.

AzureSizeList
*/
type ListAzureSizesOK struct {
	Payload models.AzureSizeList
}

func (o *ListAzureSizesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/azure/sizes][%d] listAzureSizesOK  %+v", 200, o.Payload)
}
func (o *ListAzureSizesOK) GetPayload() models.AzureSizeList {
	return o.Payload
}

func (o *ListAzureSizesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAzureSizesDefault creates a ListAzureSizesDefault with default headers values
func NewListAzureSizesDefault(code int) *ListAzureSizesDefault {
	return &ListAzureSizesDefault{
		_statusCode: code,
	}
}

/* ListAzureSizesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListAzureSizesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list azure sizes default response
func (o *ListAzureSizesDefault) Code() int {
	return o._statusCode
}

func (o *ListAzureSizesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/azure/sizes][%d] listAzureSizes default  %+v", o._statusCode, o.Payload)
}
func (o *ListAzureSizesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAzureSizesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
