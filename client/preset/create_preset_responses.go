// Code generated by go-swagger; DO NOT EDIT.

package preset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// CreatePresetReader is a Reader for the CreatePreset structure.
type CreatePresetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePresetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePresetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreatePresetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreatePresetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreatePresetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePresetOK creates a CreatePresetOK with default headers values
func NewCreatePresetOK() *CreatePresetOK {
	return &CreatePresetOK{}
}

/* CreatePresetOK describes a response with status code 200, with default header values.

Preset
*/
type CreatePresetOK struct {
	Payload *models.Preset
}

// IsSuccess returns true when this create preset o k response has a 2xx status code
func (o *CreatePresetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create preset o k response has a 3xx status code
func (o *CreatePresetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create preset o k response has a 4xx status code
func (o *CreatePresetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create preset o k response has a 5xx status code
func (o *CreatePresetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create preset o k response a status code equal to that given
func (o *CreatePresetOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreatePresetOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/providers/{provider_name}/presets][%d] createPresetOK  %+v", 200, o.Payload)
}

func (o *CreatePresetOK) String() string {
	return fmt.Sprintf("[POST /api/v2/providers/{provider_name}/presets][%d] createPresetOK  %+v", 200, o.Payload)
}

func (o *CreatePresetOK) GetPayload() *models.Preset {
	return o.Payload
}

func (o *CreatePresetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Preset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePresetUnauthorized creates a CreatePresetUnauthorized with default headers values
func NewCreatePresetUnauthorized() *CreatePresetUnauthorized {
	return &CreatePresetUnauthorized{}
}

/* CreatePresetUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreatePresetUnauthorized struct {
}

// IsSuccess returns true when this create preset unauthorized response has a 2xx status code
func (o *CreatePresetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create preset unauthorized response has a 3xx status code
func (o *CreatePresetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create preset unauthorized response has a 4xx status code
func (o *CreatePresetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create preset unauthorized response has a 5xx status code
func (o *CreatePresetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create preset unauthorized response a status code equal to that given
func (o *CreatePresetUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreatePresetUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/providers/{provider_name}/presets][%d] createPresetUnauthorized ", 401)
}

func (o *CreatePresetUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/providers/{provider_name}/presets][%d] createPresetUnauthorized ", 401)
}

func (o *CreatePresetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePresetForbidden creates a CreatePresetForbidden with default headers values
func NewCreatePresetForbidden() *CreatePresetForbidden {
	return &CreatePresetForbidden{}
}

/* CreatePresetForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreatePresetForbidden struct {
}

// IsSuccess returns true when this create preset forbidden response has a 2xx status code
func (o *CreatePresetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create preset forbidden response has a 3xx status code
func (o *CreatePresetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create preset forbidden response has a 4xx status code
func (o *CreatePresetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create preset forbidden response has a 5xx status code
func (o *CreatePresetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create preset forbidden response a status code equal to that given
func (o *CreatePresetForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreatePresetForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/providers/{provider_name}/presets][%d] createPresetForbidden ", 403)
}

func (o *CreatePresetForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/providers/{provider_name}/presets][%d] createPresetForbidden ", 403)
}

func (o *CreatePresetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePresetDefault creates a CreatePresetDefault with default headers values
func NewCreatePresetDefault(code int) *CreatePresetDefault {
	return &CreatePresetDefault{
		_statusCode: code,
	}
}

/* CreatePresetDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreatePresetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create preset default response
func (o *CreatePresetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create preset default response has a 2xx status code
func (o *CreatePresetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create preset default response has a 3xx status code
func (o *CreatePresetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create preset default response has a 4xx status code
func (o *CreatePresetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create preset default response has a 5xx status code
func (o *CreatePresetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create preset default response a status code equal to that given
func (o *CreatePresetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreatePresetDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/providers/{provider_name}/presets][%d] createPreset default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePresetDefault) String() string {
	return fmt.Sprintf("[POST /api/v2/providers/{provider_name}/presets][%d] createPreset default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePresetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreatePresetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
