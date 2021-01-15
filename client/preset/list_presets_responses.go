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

// ListPresetsReader is a Reader for the ListPresets structure.
type ListPresetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPresetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPresetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListPresetsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListPresetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListPresetsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListPresetsOK creates a ListPresetsOK with default headers values
func NewListPresetsOK() *ListPresetsOK {
	return &ListPresetsOK{}
}

/*ListPresetsOK handles this case with default header values.

PresetList
*/
type ListPresetsOK struct {
	Payload *models.PresetList
}

func (o *ListPresetsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/presets][%d] listPresetsOK  %+v", 200, o.Payload)
}

func (o *ListPresetsOK) GetPayload() *models.PresetList {
	return o.Payload
}

func (o *ListPresetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PresetList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPresetsUnauthorized creates a ListPresetsUnauthorized with default headers values
func NewListPresetsUnauthorized() *ListPresetsUnauthorized {
	return &ListPresetsUnauthorized{}
}

/*ListPresetsUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type ListPresetsUnauthorized struct {
}

func (o *ListPresetsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/presets][%d] listPresetsUnauthorized ", 401)
}

func (o *ListPresetsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListPresetsForbidden creates a ListPresetsForbidden with default headers values
func NewListPresetsForbidden() *ListPresetsForbidden {
	return &ListPresetsForbidden{}
}

/*ListPresetsForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type ListPresetsForbidden struct {
}

func (o *ListPresetsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/presets][%d] listPresetsForbidden ", 403)
}

func (o *ListPresetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListPresetsDefault creates a ListPresetsDefault with default headers values
func NewListPresetsDefault(code int) *ListPresetsDefault {
	return &ListPresetsDefault{
		_statusCode: code,
	}
}

/*ListPresetsDefault handles this case with default header values.

errorResponse
*/
type ListPresetsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list presets default response
func (o *ListPresetsDefault) Code() int {
	return o._statusCode
}

func (o *ListPresetsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/presets][%d] listPresets default  %+v", o._statusCode, o.Payload)
}

func (o *ListPresetsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListPresetsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
