// Code generated by go-swagger; DO NOT EDIT.

package resource_quota

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// PutResourceQuotaReader is a Reader for the PutResourceQuota structure.
type PutResourceQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutResourceQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutResourceQuotaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutResourceQuotaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutResourceQuotaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutResourceQuotaDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutResourceQuotaOK creates a PutResourceQuotaOK with default headers values
func NewPutResourceQuotaOK() *PutResourceQuotaOK {
	return &PutResourceQuotaOK{}
}

/* PutResourceQuotaOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type PutResourceQuotaOK struct {
}

func (o *PutResourceQuotaOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quotas/{quota_name}][%d] putResourceQuotaOK ", 200)
}

func (o *PutResourceQuotaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutResourceQuotaUnauthorized creates a PutResourceQuotaUnauthorized with default headers values
func NewPutResourceQuotaUnauthorized() *PutResourceQuotaUnauthorized {
	return &PutResourceQuotaUnauthorized{}
}

/* PutResourceQuotaUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PutResourceQuotaUnauthorized struct {
}

func (o *PutResourceQuotaUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quotas/{quota_name}][%d] putResourceQuotaUnauthorized ", 401)
}

func (o *PutResourceQuotaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutResourceQuotaForbidden creates a PutResourceQuotaForbidden with default headers values
func NewPutResourceQuotaForbidden() *PutResourceQuotaForbidden {
	return &PutResourceQuotaForbidden{}
}

/* PutResourceQuotaForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PutResourceQuotaForbidden struct {
}

func (o *PutResourceQuotaForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quotas/{quota_name}][%d] putResourceQuotaForbidden ", 403)
}

func (o *PutResourceQuotaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutResourceQuotaDefault creates a PutResourceQuotaDefault with default headers values
func NewPutResourceQuotaDefault(code int) *PutResourceQuotaDefault {
	return &PutResourceQuotaDefault{
		_statusCode: code,
	}
}

/* PutResourceQuotaDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PutResourceQuotaDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the put resource quota default response
func (o *PutResourceQuotaDefault) Code() int {
	return o._statusCode
}

func (o *PutResourceQuotaDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quotas/{quota_name}][%d] putResourceQuota default  %+v", o._statusCode, o.Payload)
}
func (o *PutResourceQuotaDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutResourceQuotaDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
