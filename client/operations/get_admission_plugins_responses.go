// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// GetAdmissionPluginsReader is a Reader for the GetAdmissionPlugins structure.
type GetAdmissionPluginsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdmissionPluginsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAdmissionPluginsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAdmissionPluginsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAdmissionPluginsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAdmissionPluginsOK creates a GetAdmissionPluginsOK with default headers values
func NewGetAdmissionPluginsOK() *GetAdmissionPluginsOK {
	return &GetAdmissionPluginsOK{}
}

/*
GetAdmissionPluginsOK describes a response with status code 200, with default header values.

AdmissionPluginList
*/
type GetAdmissionPluginsOK struct {
	Payload models.AdmissionPluginList
}

// IsSuccess returns true when this get admission plugins o k response has a 2xx status code
func (o *GetAdmissionPluginsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get admission plugins o k response has a 3xx status code
func (o *GetAdmissionPluginsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get admission plugins o k response has a 4xx status code
func (o *GetAdmissionPluginsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get admission plugins o k response has a 5xx status code
func (o *GetAdmissionPluginsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get admission plugins o k response a status code equal to that given
func (o *GetAdmissionPluginsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAdmissionPluginsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/admission/plugins/{version}][%d] getAdmissionPluginsOK  %+v", 200, o.Payload)
}

func (o *GetAdmissionPluginsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/admission/plugins/{version}][%d] getAdmissionPluginsOK  %+v", 200, o.Payload)
}

func (o *GetAdmissionPluginsOK) GetPayload() models.AdmissionPluginList {
	return o.Payload
}

func (o *GetAdmissionPluginsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdmissionPluginsUnauthorized creates a GetAdmissionPluginsUnauthorized with default headers values
func NewGetAdmissionPluginsUnauthorized() *GetAdmissionPluginsUnauthorized {
	return &GetAdmissionPluginsUnauthorized{}
}

/*
GetAdmissionPluginsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetAdmissionPluginsUnauthorized struct {
}

// IsSuccess returns true when this get admission plugins unauthorized response has a 2xx status code
func (o *GetAdmissionPluginsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get admission plugins unauthorized response has a 3xx status code
func (o *GetAdmissionPluginsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get admission plugins unauthorized response has a 4xx status code
func (o *GetAdmissionPluginsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get admission plugins unauthorized response has a 5xx status code
func (o *GetAdmissionPluginsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get admission plugins unauthorized response a status code equal to that given
func (o *GetAdmissionPluginsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAdmissionPluginsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admission/plugins/{version}][%d] getAdmissionPluginsUnauthorized ", 401)
}

func (o *GetAdmissionPluginsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/admission/plugins/{version}][%d] getAdmissionPluginsUnauthorized ", 401)
}

func (o *GetAdmissionPluginsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdmissionPluginsDefault creates a GetAdmissionPluginsDefault with default headers values
func NewGetAdmissionPluginsDefault(code int) *GetAdmissionPluginsDefault {
	return &GetAdmissionPluginsDefault{
		_statusCode: code,
	}
}

/*
GetAdmissionPluginsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetAdmissionPluginsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get admission plugins default response
func (o *GetAdmissionPluginsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get admission plugins default response has a 2xx status code
func (o *GetAdmissionPluginsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get admission plugins default response has a 3xx status code
func (o *GetAdmissionPluginsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get admission plugins default response has a 4xx status code
func (o *GetAdmissionPluginsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get admission plugins default response has a 5xx status code
func (o *GetAdmissionPluginsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get admission plugins default response a status code equal to that given
func (o *GetAdmissionPluginsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetAdmissionPluginsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/admission/plugins/{version}][%d] getAdmissionPlugins default  %+v", o._statusCode, o.Payload)
}

func (o *GetAdmissionPluginsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/admission/plugins/{version}][%d] getAdmissionPlugins default  %+v", o._statusCode, o.Payload)
}

func (o *GetAdmissionPluginsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAdmissionPluginsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
