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

// ListAzureSizesNoCredentialsV2Reader is a Reader for the ListAzureSizesNoCredentialsV2 structure.
type ListAzureSizesNoCredentialsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAzureSizesNoCredentialsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAzureSizesNoCredentialsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAzureSizesNoCredentialsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAzureSizesNoCredentialsV2OK creates a ListAzureSizesNoCredentialsV2OK with default headers values
func NewListAzureSizesNoCredentialsV2OK() *ListAzureSizesNoCredentialsV2OK {
	return &ListAzureSizesNoCredentialsV2OK{}
}

/* ListAzureSizesNoCredentialsV2OK describes a response with status code 200, with default header values.

AzureSizeList
*/
type ListAzureSizesNoCredentialsV2OK struct {
	Payload models.AzureSizeList
}

// IsSuccess returns true when this list azure sizes no credentials v2 o k response has a 2xx status code
func (o *ListAzureSizesNoCredentialsV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list azure sizes no credentials v2 o k response has a 3xx status code
func (o *ListAzureSizesNoCredentialsV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list azure sizes no credentials v2 o k response has a 4xx status code
func (o *ListAzureSizesNoCredentialsV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list azure sizes no credentials v2 o k response has a 5xx status code
func (o *ListAzureSizesNoCredentialsV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list azure sizes no credentials v2 o k response a status code equal to that given
func (o *ListAzureSizesNoCredentialsV2OK) IsCode(code int) bool {
	return code == 200
}

func (o *ListAzureSizesNoCredentialsV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/azure/sizes][%d] listAzureSizesNoCredentialsV2OK  %+v", 200, o.Payload)
}

func (o *ListAzureSizesNoCredentialsV2OK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/azure/sizes][%d] listAzureSizesNoCredentialsV2OK  %+v", 200, o.Payload)
}

func (o *ListAzureSizesNoCredentialsV2OK) GetPayload() models.AzureSizeList {
	return o.Payload
}

func (o *ListAzureSizesNoCredentialsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAzureSizesNoCredentialsV2Default creates a ListAzureSizesNoCredentialsV2Default with default headers values
func NewListAzureSizesNoCredentialsV2Default(code int) *ListAzureSizesNoCredentialsV2Default {
	return &ListAzureSizesNoCredentialsV2Default{
		_statusCode: code,
	}
}

/* ListAzureSizesNoCredentialsV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type ListAzureSizesNoCredentialsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list azure sizes no credentials v2 default response
func (o *ListAzureSizesNoCredentialsV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list azure sizes no credentials v2 default response has a 2xx status code
func (o *ListAzureSizesNoCredentialsV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list azure sizes no credentials v2 default response has a 3xx status code
func (o *ListAzureSizesNoCredentialsV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list azure sizes no credentials v2 default response has a 4xx status code
func (o *ListAzureSizesNoCredentialsV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list azure sizes no credentials v2 default response has a 5xx status code
func (o *ListAzureSizesNoCredentialsV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list azure sizes no credentials v2 default response a status code equal to that given
func (o *ListAzureSizesNoCredentialsV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListAzureSizesNoCredentialsV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/azure/sizes][%d] listAzureSizesNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListAzureSizesNoCredentialsV2Default) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/azure/sizes][%d] listAzureSizesNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListAzureSizesNoCredentialsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAzureSizesNoCredentialsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
