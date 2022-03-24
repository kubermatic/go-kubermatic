// Code generated by go-swagger; DO NOT EDIT.

package datacenter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/kubermatic/go-kubermatic/models"
)

// CreateDCReader is a Reader for the CreateDC structure.
type CreateDCReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDCReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDCCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateDCUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateDCForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateDCDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateDCCreated creates a CreateDCCreated with default headers values
func NewCreateDCCreated() *CreateDCCreated {
	return &CreateDCCreated{}
}

/* CreateDCCreated describes a response with status code 201, with default header values.

Datacenter
*/
type CreateDCCreated struct {
	Payload *models.Datacenter
}

func (o *CreateDCCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/seed/{seed_name}/dc][%d] createDCCreated  %+v", 201, o.Payload)
}
func (o *CreateDCCreated) GetPayload() *models.Datacenter {
	return o.Payload
}

func (o *CreateDCCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Datacenter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDCUnauthorized creates a CreateDCUnauthorized with default headers values
func NewCreateDCUnauthorized() *CreateDCUnauthorized {
	return &CreateDCUnauthorized{}
}

/* CreateDCUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateDCUnauthorized struct {
}

func (o *CreateDCUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/seed/{seed_name}/dc][%d] createDCUnauthorized ", 401)
}

func (o *CreateDCUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDCForbidden creates a CreateDCForbidden with default headers values
func NewCreateDCForbidden() *CreateDCForbidden {
	return &CreateDCForbidden{}
}

/* CreateDCForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateDCForbidden struct {
}

func (o *CreateDCForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/seed/{seed_name}/dc][%d] createDCForbidden ", 403)
}

func (o *CreateDCForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDCDefault creates a CreateDCDefault with default headers values
func NewCreateDCDefault(code int) *CreateDCDefault {
	return &CreateDCDefault{
		_statusCode: code,
	}
}

/* CreateDCDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateDCDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create d c default response
func (o *CreateDCDefault) Code() int {
	return o._statusCode
}

func (o *CreateDCDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/seed/{seed_name}/dc][%d] createDC default  %+v", o._statusCode, o.Payload)
}
func (o *CreateDCDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateDCDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateDCBody create d c body
swagger:model CreateDCBody
*/
type CreateDCBody struct {

	// name
	Name string `json:"name,omitempty"`

	// spec
	Spec *models.DatacenterSpec `json:"spec,omitempty"`
}

// Validate validates this create d c body
func (o *CreateDCBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateDCBody) validateSpec(formats strfmt.Registry) error {
	if swag.IsZero(o.Spec) { // not required
		return nil
	}

	if o.Spec != nil {
		if err := o.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "spec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create d c body based on the context it is used
func (o *CreateDCBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateDCBody) contextValidateSpec(ctx context.Context, formats strfmt.Registry) error {

	if o.Spec != nil {
		if err := o.Spec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateDCBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateDCBody) UnmarshalBinary(b []byte) error {
	var res CreateDCBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
