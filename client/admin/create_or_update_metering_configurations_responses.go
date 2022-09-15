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

// CreateOrUpdateMeteringConfigurationsReader is a Reader for the CreateOrUpdateMeteringConfigurations structure.
type CreateOrUpdateMeteringConfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrUpdateMeteringConfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOrUpdateMeteringConfigurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateOrUpdateMeteringConfigurationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateOrUpdateMeteringConfigurationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateOrUpdateMeteringConfigurationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateOrUpdateMeteringConfigurationsOK creates a CreateOrUpdateMeteringConfigurationsOK with default headers values
func NewCreateOrUpdateMeteringConfigurationsOK() *CreateOrUpdateMeteringConfigurationsOK {
	return &CreateOrUpdateMeteringConfigurationsOK{}
}

/*
CreateOrUpdateMeteringConfigurationsOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type CreateOrUpdateMeteringConfigurationsOK struct {
}

// IsSuccess returns true when this create or update metering configurations o k response has a 2xx status code
func (o *CreateOrUpdateMeteringConfigurationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create or update metering configurations o k response has a 3xx status code
func (o *CreateOrUpdateMeteringConfigurationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create or update metering configurations o k response has a 4xx status code
func (o *CreateOrUpdateMeteringConfigurationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create or update metering configurations o k response has a 5xx status code
func (o *CreateOrUpdateMeteringConfigurationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create or update metering configurations o k response a status code equal to that given
func (o *CreateOrUpdateMeteringConfigurationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateOrUpdateMeteringConfigurationsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/configurations][%d] createOrUpdateMeteringConfigurationsOK ", 200)
}

func (o *CreateOrUpdateMeteringConfigurationsOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/configurations][%d] createOrUpdateMeteringConfigurationsOK ", 200)
}

func (o *CreateOrUpdateMeteringConfigurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrUpdateMeteringConfigurationsUnauthorized creates a CreateOrUpdateMeteringConfigurationsUnauthorized with default headers values
func NewCreateOrUpdateMeteringConfigurationsUnauthorized() *CreateOrUpdateMeteringConfigurationsUnauthorized {
	return &CreateOrUpdateMeteringConfigurationsUnauthorized{}
}

/*
CreateOrUpdateMeteringConfigurationsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateOrUpdateMeteringConfigurationsUnauthorized struct {
}

// IsSuccess returns true when this create or update metering configurations unauthorized response has a 2xx status code
func (o *CreateOrUpdateMeteringConfigurationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create or update metering configurations unauthorized response has a 3xx status code
func (o *CreateOrUpdateMeteringConfigurationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create or update metering configurations unauthorized response has a 4xx status code
func (o *CreateOrUpdateMeteringConfigurationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create or update metering configurations unauthorized response has a 5xx status code
func (o *CreateOrUpdateMeteringConfigurationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create or update metering configurations unauthorized response a status code equal to that given
func (o *CreateOrUpdateMeteringConfigurationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateOrUpdateMeteringConfigurationsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/configurations][%d] createOrUpdateMeteringConfigurationsUnauthorized ", 401)
}

func (o *CreateOrUpdateMeteringConfigurationsUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/configurations][%d] createOrUpdateMeteringConfigurationsUnauthorized ", 401)
}

func (o *CreateOrUpdateMeteringConfigurationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrUpdateMeteringConfigurationsForbidden creates a CreateOrUpdateMeteringConfigurationsForbidden with default headers values
func NewCreateOrUpdateMeteringConfigurationsForbidden() *CreateOrUpdateMeteringConfigurationsForbidden {
	return &CreateOrUpdateMeteringConfigurationsForbidden{}
}

/*
CreateOrUpdateMeteringConfigurationsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateOrUpdateMeteringConfigurationsForbidden struct {
}

// IsSuccess returns true when this create or update metering configurations forbidden response has a 2xx status code
func (o *CreateOrUpdateMeteringConfigurationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create or update metering configurations forbidden response has a 3xx status code
func (o *CreateOrUpdateMeteringConfigurationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create or update metering configurations forbidden response has a 4xx status code
func (o *CreateOrUpdateMeteringConfigurationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create or update metering configurations forbidden response has a 5xx status code
func (o *CreateOrUpdateMeteringConfigurationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create or update metering configurations forbidden response a status code equal to that given
func (o *CreateOrUpdateMeteringConfigurationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateOrUpdateMeteringConfigurationsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/configurations][%d] createOrUpdateMeteringConfigurationsForbidden ", 403)
}

func (o *CreateOrUpdateMeteringConfigurationsForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/configurations][%d] createOrUpdateMeteringConfigurationsForbidden ", 403)
}

func (o *CreateOrUpdateMeteringConfigurationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrUpdateMeteringConfigurationsDefault creates a CreateOrUpdateMeteringConfigurationsDefault with default headers values
func NewCreateOrUpdateMeteringConfigurationsDefault(code int) *CreateOrUpdateMeteringConfigurationsDefault {
	return &CreateOrUpdateMeteringConfigurationsDefault{
		_statusCode: code,
	}
}

/*
CreateOrUpdateMeteringConfigurationsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateOrUpdateMeteringConfigurationsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create or update metering configurations default response
func (o *CreateOrUpdateMeteringConfigurationsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create or update metering configurations default response has a 2xx status code
func (o *CreateOrUpdateMeteringConfigurationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create or update metering configurations default response has a 3xx status code
func (o *CreateOrUpdateMeteringConfigurationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create or update metering configurations default response has a 4xx status code
func (o *CreateOrUpdateMeteringConfigurationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create or update metering configurations default response has a 5xx status code
func (o *CreateOrUpdateMeteringConfigurationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create or update metering configurations default response a status code equal to that given
func (o *CreateOrUpdateMeteringConfigurationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateOrUpdateMeteringConfigurationsDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/configurations][%d] createOrUpdateMeteringConfigurations default  %+v", o._statusCode, o.Payload)
}

func (o *CreateOrUpdateMeteringConfigurationsDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/configurations][%d] createOrUpdateMeteringConfigurations default  %+v", o._statusCode, o.Payload)
}

func (o *CreateOrUpdateMeteringConfigurationsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateOrUpdateMeteringConfigurationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
