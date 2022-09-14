// Code generated by go-swagger; DO NOT EDIT.

package alibaba

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListAlibabaZonesNoCredentialsV2Reader is a Reader for the ListAlibabaZonesNoCredentialsV2 structure.
type ListAlibabaZonesNoCredentialsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAlibabaZonesNoCredentialsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAlibabaZonesNoCredentialsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAlibabaZonesNoCredentialsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAlibabaZonesNoCredentialsV2OK creates a ListAlibabaZonesNoCredentialsV2OK with default headers values
func NewListAlibabaZonesNoCredentialsV2OK() *ListAlibabaZonesNoCredentialsV2OK {
	return &ListAlibabaZonesNoCredentialsV2OK{}
}

/* ListAlibabaZonesNoCredentialsV2OK describes a response with status code 200, with default header values.

AlibabaZoneList
*/
type ListAlibabaZonesNoCredentialsV2OK struct {
	Payload models.AlibabaZoneList
}

// IsSuccess returns true when this list alibaba zones no credentials v2 o k response has a 2xx status code
func (o *ListAlibabaZonesNoCredentialsV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list alibaba zones no credentials v2 o k response has a 3xx status code
func (o *ListAlibabaZonesNoCredentialsV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list alibaba zones no credentials v2 o k response has a 4xx status code
func (o *ListAlibabaZonesNoCredentialsV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list alibaba zones no credentials v2 o k response has a 5xx status code
func (o *ListAlibabaZonesNoCredentialsV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list alibaba zones no credentials v2 o k response a status code equal to that given
func (o *ListAlibabaZonesNoCredentialsV2OK) IsCode(code int) bool {
	return code == 200
}

func (o *ListAlibabaZonesNoCredentialsV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/alibaba/zones][%d] listAlibabaZonesNoCredentialsV2OK  %+v", 200, o.Payload)
}

func (o *ListAlibabaZonesNoCredentialsV2OK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/alibaba/zones][%d] listAlibabaZonesNoCredentialsV2OK  %+v", 200, o.Payload)
}

func (o *ListAlibabaZonesNoCredentialsV2OK) GetPayload() models.AlibabaZoneList {
	return o.Payload
}

func (o *ListAlibabaZonesNoCredentialsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAlibabaZonesNoCredentialsV2Default creates a ListAlibabaZonesNoCredentialsV2Default with default headers values
func NewListAlibabaZonesNoCredentialsV2Default(code int) *ListAlibabaZonesNoCredentialsV2Default {
	return &ListAlibabaZonesNoCredentialsV2Default{
		_statusCode: code,
	}
}

/* ListAlibabaZonesNoCredentialsV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type ListAlibabaZonesNoCredentialsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list alibaba zones no credentials v2 default response
func (o *ListAlibabaZonesNoCredentialsV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list alibaba zones no credentials v2 default response has a 2xx status code
func (o *ListAlibabaZonesNoCredentialsV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list alibaba zones no credentials v2 default response has a 3xx status code
func (o *ListAlibabaZonesNoCredentialsV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list alibaba zones no credentials v2 default response has a 4xx status code
func (o *ListAlibabaZonesNoCredentialsV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list alibaba zones no credentials v2 default response has a 5xx status code
func (o *ListAlibabaZonesNoCredentialsV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list alibaba zones no credentials v2 default response a status code equal to that given
func (o *ListAlibabaZonesNoCredentialsV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListAlibabaZonesNoCredentialsV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/alibaba/zones][%d] listAlibabaZonesNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListAlibabaZonesNoCredentialsV2Default) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/alibaba/zones][%d] listAlibabaZonesNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListAlibabaZonesNoCredentialsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAlibabaZonesNoCredentialsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
