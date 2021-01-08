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

// ListAzureResourceGroupsReader is a Reader for the ListAzureResourceGroups structure.
type ListAzureResourceGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAzureResourceGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAzureResourceGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAzureResourceGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAzureResourceGroupsOK creates a ListAzureResourceGroupsOK with default headers values
func NewListAzureResourceGroupsOK() *ListAzureResourceGroupsOK {
	return &ListAzureResourceGroupsOK{}
}

/*ListAzureResourceGroupsOK handles this case with default header values.

AzureResourceGroupsList
*/
type ListAzureResourceGroupsOK struct {
	Payload *models.AzureResourceGroupsList
}

func (o *ListAzureResourceGroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/azure/resourcegroups][%d] listAzureResourceGroupsOK  %+v", 200, o.Payload)
}

func (o *ListAzureResourceGroupsOK) GetPayload() *models.AzureResourceGroupsList {
	return o.Payload
}

func (o *ListAzureResourceGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureResourceGroupsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAzureResourceGroupsDefault creates a ListAzureResourceGroupsDefault with default headers values
func NewListAzureResourceGroupsDefault(code int) *ListAzureResourceGroupsDefault {
	return &ListAzureResourceGroupsDefault{
		_statusCode: code,
	}
}

/*ListAzureResourceGroupsDefault handles this case with default header values.

errorResponse
*/
type ListAzureResourceGroupsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list azure resource groups default response
func (o *ListAzureResourceGroupsDefault) Code() int {
	return o._statusCode
}

func (o *ListAzureResourceGroupsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/azure/resourcegroups][%d] listAzureResourceGroups default  %+v", o._statusCode, o.Payload)
}

func (o *ListAzureResourceGroupsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAzureResourceGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
