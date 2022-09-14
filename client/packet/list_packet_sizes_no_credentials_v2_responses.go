// Code generated by go-swagger; DO NOT EDIT.

package packet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListPacketSizesNoCredentialsV2Reader is a Reader for the ListPacketSizesNoCredentialsV2 structure.
type ListPacketSizesNoCredentialsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPacketSizesNoCredentialsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPacketSizesNoCredentialsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListPacketSizesNoCredentialsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListPacketSizesNoCredentialsV2OK creates a ListPacketSizesNoCredentialsV2OK with default headers values
func NewListPacketSizesNoCredentialsV2OK() *ListPacketSizesNoCredentialsV2OK {
	return &ListPacketSizesNoCredentialsV2OK{}
}

/* ListPacketSizesNoCredentialsV2OK describes a response with status code 200, with default header values.

PacketSizeList
*/
type ListPacketSizesNoCredentialsV2OK struct {
	Payload models.PacketSizeList
}

// IsSuccess returns true when this list packet sizes no credentials v2 o k response has a 2xx status code
func (o *ListPacketSizesNoCredentialsV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list packet sizes no credentials v2 o k response has a 3xx status code
func (o *ListPacketSizesNoCredentialsV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list packet sizes no credentials v2 o k response has a 4xx status code
func (o *ListPacketSizesNoCredentialsV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list packet sizes no credentials v2 o k response has a 5xx status code
func (o *ListPacketSizesNoCredentialsV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list packet sizes no credentials v2 o k response a status code equal to that given
func (o *ListPacketSizesNoCredentialsV2OK) IsCode(code int) bool {
	return code == 200
}

func (o *ListPacketSizesNoCredentialsV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/packet/sizes][%d] listPacketSizesNoCredentialsV2OK  %+v", 200, o.Payload)
}

func (o *ListPacketSizesNoCredentialsV2OK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/packet/sizes][%d] listPacketSizesNoCredentialsV2OK  %+v", 200, o.Payload)
}

func (o *ListPacketSizesNoCredentialsV2OK) GetPayload() models.PacketSizeList {
	return o.Payload
}

func (o *ListPacketSizesNoCredentialsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPacketSizesNoCredentialsV2Default creates a ListPacketSizesNoCredentialsV2Default with default headers values
func NewListPacketSizesNoCredentialsV2Default(code int) *ListPacketSizesNoCredentialsV2Default {
	return &ListPacketSizesNoCredentialsV2Default{
		_statusCode: code,
	}
}

/* ListPacketSizesNoCredentialsV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type ListPacketSizesNoCredentialsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list packet sizes no credentials v2 default response
func (o *ListPacketSizesNoCredentialsV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list packet sizes no credentials v2 default response has a 2xx status code
func (o *ListPacketSizesNoCredentialsV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list packet sizes no credentials v2 default response has a 3xx status code
func (o *ListPacketSizesNoCredentialsV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list packet sizes no credentials v2 default response has a 4xx status code
func (o *ListPacketSizesNoCredentialsV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list packet sizes no credentials v2 default response has a 5xx status code
func (o *ListPacketSizesNoCredentialsV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list packet sizes no credentials v2 default response a status code equal to that given
func (o *ListPacketSizesNoCredentialsV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListPacketSizesNoCredentialsV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/packet/sizes][%d] listPacketSizesNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListPacketSizesNoCredentialsV2Default) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/packet/sizes][%d] listPacketSizesNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListPacketSizesNoCredentialsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListPacketSizesNoCredentialsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
