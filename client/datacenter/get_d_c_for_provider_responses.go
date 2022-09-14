// Code generated by go-swagger; DO NOT EDIT.

package datacenter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// GetDCForProviderReader is a Reader for the GetDCForProvider structure.
type GetDCForProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDCForProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDCForProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDCForProviderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDCForProviderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDCForProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDCForProviderOK creates a GetDCForProviderOK with default headers values
func NewGetDCForProviderOK() *GetDCForProviderOK {
	return &GetDCForProviderOK{}
}

/* GetDCForProviderOK describes a response with status code 200, with default header values.

Datacenter
*/
type GetDCForProviderOK struct {
	Payload *models.Datacenter
}

// IsSuccess returns true when this get d c for provider o k response has a 2xx status code
func (o *GetDCForProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get d c for provider o k response has a 3xx status code
func (o *GetDCForProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d c for provider o k response has a 4xx status code
func (o *GetDCForProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get d c for provider o k response has a 5xx status code
func (o *GetDCForProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get d c for provider o k response a status code equal to that given
func (o *GetDCForProviderOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDCForProviderOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/{provider_name}/dc/{dc}][%d] getDCForProviderOK  %+v", 200, o.Payload)
}

func (o *GetDCForProviderOK) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/{provider_name}/dc/{dc}][%d] getDCForProviderOK  %+v", 200, o.Payload)
}

func (o *GetDCForProviderOK) GetPayload() *models.Datacenter {
	return o.Payload
}

func (o *GetDCForProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Datacenter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDCForProviderUnauthorized creates a GetDCForProviderUnauthorized with default headers values
func NewGetDCForProviderUnauthorized() *GetDCForProviderUnauthorized {
	return &GetDCForProviderUnauthorized{}
}

/* GetDCForProviderUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetDCForProviderUnauthorized struct {
}

// IsSuccess returns true when this get d c for provider unauthorized response has a 2xx status code
func (o *GetDCForProviderUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d c for provider unauthorized response has a 3xx status code
func (o *GetDCForProviderUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d c for provider unauthorized response has a 4xx status code
func (o *GetDCForProviderUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get d c for provider unauthorized response has a 5xx status code
func (o *GetDCForProviderUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get d c for provider unauthorized response a status code equal to that given
func (o *GetDCForProviderUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetDCForProviderUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/{provider_name}/dc/{dc}][%d] getDCForProviderUnauthorized ", 401)
}

func (o *GetDCForProviderUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/{provider_name}/dc/{dc}][%d] getDCForProviderUnauthorized ", 401)
}

func (o *GetDCForProviderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDCForProviderForbidden creates a GetDCForProviderForbidden with default headers values
func NewGetDCForProviderForbidden() *GetDCForProviderForbidden {
	return &GetDCForProviderForbidden{}
}

/* GetDCForProviderForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetDCForProviderForbidden struct {
}

// IsSuccess returns true when this get d c for provider forbidden response has a 2xx status code
func (o *GetDCForProviderForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d c for provider forbidden response has a 3xx status code
func (o *GetDCForProviderForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d c for provider forbidden response has a 4xx status code
func (o *GetDCForProviderForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get d c for provider forbidden response has a 5xx status code
func (o *GetDCForProviderForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get d c for provider forbidden response a status code equal to that given
func (o *GetDCForProviderForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetDCForProviderForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/{provider_name}/dc/{dc}][%d] getDCForProviderForbidden ", 403)
}

func (o *GetDCForProviderForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/{provider_name}/dc/{dc}][%d] getDCForProviderForbidden ", 403)
}

func (o *GetDCForProviderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDCForProviderDefault creates a GetDCForProviderDefault with default headers values
func NewGetDCForProviderDefault(code int) *GetDCForProviderDefault {
	return &GetDCForProviderDefault{
		_statusCode: code,
	}
}

/* GetDCForProviderDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetDCForProviderDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get d c for provider default response
func (o *GetDCForProviderDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get d c for provider default response has a 2xx status code
func (o *GetDCForProviderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get d c for provider default response has a 3xx status code
func (o *GetDCForProviderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get d c for provider default response has a 4xx status code
func (o *GetDCForProviderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get d c for provider default response has a 5xx status code
func (o *GetDCForProviderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get d c for provider default response a status code equal to that given
func (o *GetDCForProviderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetDCForProviderDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/{provider_name}/dc/{dc}][%d] getDCForProvider default  %+v", o._statusCode, o.Payload)
}

func (o *GetDCForProviderDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/{provider_name}/dc/{dc}][%d] getDCForProvider default  %+v", o._statusCode, o.Payload)
}

func (o *GetDCForProviderDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDCForProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
