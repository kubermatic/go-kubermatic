// Code generated by go-swagger; DO NOT EDIT.

package preset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/kubermatic/go-kubermatic/models"
)

// UpdatePresetStatusReader is a Reader for the UpdatePresetStatus structure.
type UpdatePresetStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePresetStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePresetStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdatePresetStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdatePresetStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdatePresetStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdatePresetStatusOK creates a UpdatePresetStatusOK with default headers values
func NewUpdatePresetStatusOK() *UpdatePresetStatusOK {
	return &UpdatePresetStatusOK{}
}

/*UpdatePresetStatusOK handles this case with default header values.

EmptyResponse is a empty response
*/
type UpdatePresetStatusOK struct {
}

func (o *UpdatePresetStatusOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/presets/{preset_name}/status][%d] updatePresetStatusOK ", 200)
}

func (o *UpdatePresetStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePresetStatusUnauthorized creates a UpdatePresetStatusUnauthorized with default headers values
func NewUpdatePresetStatusUnauthorized() *UpdatePresetStatusUnauthorized {
	return &UpdatePresetStatusUnauthorized{}
}

/*UpdatePresetStatusUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type UpdatePresetStatusUnauthorized struct {
}

func (o *UpdatePresetStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/presets/{preset_name}/status][%d] updatePresetStatusUnauthorized ", 401)
}

func (o *UpdatePresetStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePresetStatusForbidden creates a UpdatePresetStatusForbidden with default headers values
func NewUpdatePresetStatusForbidden() *UpdatePresetStatusForbidden {
	return &UpdatePresetStatusForbidden{}
}

/*UpdatePresetStatusForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type UpdatePresetStatusForbidden struct {
}

func (o *UpdatePresetStatusForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/presets/{preset_name}/status][%d] updatePresetStatusForbidden ", 403)
}

func (o *UpdatePresetStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePresetStatusDefault creates a UpdatePresetStatusDefault with default headers values
func NewUpdatePresetStatusDefault(code int) *UpdatePresetStatusDefault {
	return &UpdatePresetStatusDefault{
		_statusCode: code,
	}
}

/*UpdatePresetStatusDefault handles this case with default header values.

errorResponse
*/
type UpdatePresetStatusDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update preset status default response
func (o *UpdatePresetStatusDefault) Code() int {
	return o._statusCode
}

func (o *UpdatePresetStatusDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v2/presets/{preset_name}/status][%d] updatePresetStatus default  %+v", o._statusCode, o.Payload)
}

func (o *UpdatePresetStatusDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdatePresetStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdatePresetStatusBody update preset status body
swagger:model UpdatePresetStatusBody
*/
type UpdatePresetStatusBody struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this update preset status body
func (o *UpdatePresetStatusBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdatePresetStatusBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdatePresetStatusBody) UnmarshalBinary(b []byte) error {
	var res UpdatePresetStatusBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
