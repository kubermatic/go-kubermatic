// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// CreateOrUpdateMeteringCredentialsReader is a Reader for the CreateOrUpdateMeteringCredentials structure.
type CreateOrUpdateMeteringCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrUpdateMeteringCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOrUpdateMeteringCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateOrUpdateMeteringCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateOrUpdateMeteringCredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateOrUpdateMeteringCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateOrUpdateMeteringCredentialsOK creates a CreateOrUpdateMeteringCredentialsOK with default headers values
func NewCreateOrUpdateMeteringCredentialsOK() *CreateOrUpdateMeteringCredentialsOK {
	return &CreateOrUpdateMeteringCredentialsOK{}
}

/* CreateOrUpdateMeteringCredentialsOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type CreateOrUpdateMeteringCredentialsOK struct {
}

func (o *CreateOrUpdateMeteringCredentialsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/credentials][%d] createOrUpdateMeteringCredentialsOK ", 200)
}

func (o *CreateOrUpdateMeteringCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrUpdateMeteringCredentialsUnauthorized creates a CreateOrUpdateMeteringCredentialsUnauthorized with default headers values
func NewCreateOrUpdateMeteringCredentialsUnauthorized() *CreateOrUpdateMeteringCredentialsUnauthorized {
	return &CreateOrUpdateMeteringCredentialsUnauthorized{}
}

/* CreateOrUpdateMeteringCredentialsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateOrUpdateMeteringCredentialsUnauthorized struct {
}

func (o *CreateOrUpdateMeteringCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/credentials][%d] createOrUpdateMeteringCredentialsUnauthorized ", 401)
}

func (o *CreateOrUpdateMeteringCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrUpdateMeteringCredentialsForbidden creates a CreateOrUpdateMeteringCredentialsForbidden with default headers values
func NewCreateOrUpdateMeteringCredentialsForbidden() *CreateOrUpdateMeteringCredentialsForbidden {
	return &CreateOrUpdateMeteringCredentialsForbidden{}
}

/* CreateOrUpdateMeteringCredentialsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateOrUpdateMeteringCredentialsForbidden struct {
}

func (o *CreateOrUpdateMeteringCredentialsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/credentials][%d] createOrUpdateMeteringCredentialsForbidden ", 403)
}

func (o *CreateOrUpdateMeteringCredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrUpdateMeteringCredentialsDefault creates a CreateOrUpdateMeteringCredentialsDefault with default headers values
func NewCreateOrUpdateMeteringCredentialsDefault(code int) *CreateOrUpdateMeteringCredentialsDefault {
	return &CreateOrUpdateMeteringCredentialsDefault{
		_statusCode: code,
	}
}

/* CreateOrUpdateMeteringCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateOrUpdateMeteringCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create or update metering credentials default response
func (o *CreateOrUpdateMeteringCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *CreateOrUpdateMeteringCredentialsDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/credentials][%d] createOrUpdateMeteringCredentials default  %+v", o._statusCode, o.Payload)
}
func (o *CreateOrUpdateMeteringCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateOrUpdateMeteringCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
