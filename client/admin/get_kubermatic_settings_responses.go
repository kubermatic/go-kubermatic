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

// GetKubermaticSettingsReader is a Reader for the GetKubermaticSettings structure.
type GetKubermaticSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubermaticSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubermaticSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetKubermaticSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKubermaticSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetKubermaticSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetKubermaticSettingsOK creates a GetKubermaticSettingsOK with default headers values
func NewGetKubermaticSettingsOK() *GetKubermaticSettingsOK {
	return &GetKubermaticSettingsOK{}
}

/* GetKubermaticSettingsOK describes a response with status code 200, with default header values.

GlobalSettings
*/
type GetKubermaticSettingsOK struct {
	Payload *models.GlobalSettings
}

func (o *GetKubermaticSettingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/settings][%d] getKubermaticSettingsOK  %+v", 200, o.Payload)
}
func (o *GetKubermaticSettingsOK) GetPayload() *models.GlobalSettings {
	return o.Payload
}

func (o *GetKubermaticSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GlobalSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubermaticSettingsUnauthorized creates a GetKubermaticSettingsUnauthorized with default headers values
func NewGetKubermaticSettingsUnauthorized() *GetKubermaticSettingsUnauthorized {
	return &GetKubermaticSettingsUnauthorized{}
}

/* GetKubermaticSettingsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetKubermaticSettingsUnauthorized struct {
}

func (o *GetKubermaticSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/settings][%d] getKubermaticSettingsUnauthorized ", 401)
}

func (o *GetKubermaticSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubermaticSettingsForbidden creates a GetKubermaticSettingsForbidden with default headers values
func NewGetKubermaticSettingsForbidden() *GetKubermaticSettingsForbidden {
	return &GetKubermaticSettingsForbidden{}
}

/* GetKubermaticSettingsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetKubermaticSettingsForbidden struct {
}

func (o *GetKubermaticSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/settings][%d] getKubermaticSettingsForbidden ", 403)
}

func (o *GetKubermaticSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubermaticSettingsDefault creates a GetKubermaticSettingsDefault with default headers values
func NewGetKubermaticSettingsDefault(code int) *GetKubermaticSettingsDefault {
	return &GetKubermaticSettingsDefault{
		_statusCode: code,
	}
}

/* GetKubermaticSettingsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetKubermaticSettingsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get kubermatic settings default response
func (o *GetKubermaticSettingsDefault) Code() int {
	return o._statusCode
}

func (o *GetKubermaticSettingsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/settings][%d] getKubermaticSettings default  %+v", o._statusCode, o.Payload)
}
func (o *GetKubermaticSettingsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetKubermaticSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
