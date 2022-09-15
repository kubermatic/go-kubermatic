// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/kubermatic/go-kubermatic/models"
)

// CreateClusterTemplateInstanceReader is a Reader for the CreateClusterTemplateInstance structure.
type CreateClusterTemplateInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClusterTemplateInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateClusterTemplateInstanceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateClusterTemplateInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateClusterTemplateInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateClusterTemplateInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateClusterTemplateInstanceCreated creates a CreateClusterTemplateInstanceCreated with default headers values
func NewCreateClusterTemplateInstanceCreated() *CreateClusterTemplateInstanceCreated {
	return &CreateClusterTemplateInstanceCreated{}
}

/*
CreateClusterTemplateInstanceCreated describes a response with status code 201, with default header values.

ClusterTemplateInstance
*/
type CreateClusterTemplateInstanceCreated struct {
	Payload *models.ClusterTemplateInstance
}

// IsSuccess returns true when this create cluster template instance created response has a 2xx status code
func (o *CreateClusterTemplateInstanceCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create cluster template instance created response has a 3xx status code
func (o *CreateClusterTemplateInstanceCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cluster template instance created response has a 4xx status code
func (o *CreateClusterTemplateInstanceCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create cluster template instance created response has a 5xx status code
func (o *CreateClusterTemplateInstanceCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create cluster template instance created response a status code equal to that given
func (o *CreateClusterTemplateInstanceCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateClusterTemplateInstanceCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates/{template_id}/instances][%d] createClusterTemplateInstanceCreated  %+v", 201, o.Payload)
}

func (o *CreateClusterTemplateInstanceCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates/{template_id}/instances][%d] createClusterTemplateInstanceCreated  %+v", 201, o.Payload)
}

func (o *CreateClusterTemplateInstanceCreated) GetPayload() *models.ClusterTemplateInstance {
	return o.Payload
}

func (o *CreateClusterTemplateInstanceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterTemplateInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterTemplateInstanceUnauthorized creates a CreateClusterTemplateInstanceUnauthorized with default headers values
func NewCreateClusterTemplateInstanceUnauthorized() *CreateClusterTemplateInstanceUnauthorized {
	return &CreateClusterTemplateInstanceUnauthorized{}
}

/*
CreateClusterTemplateInstanceUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateClusterTemplateInstanceUnauthorized struct {
}

// IsSuccess returns true when this create cluster template instance unauthorized response has a 2xx status code
func (o *CreateClusterTemplateInstanceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cluster template instance unauthorized response has a 3xx status code
func (o *CreateClusterTemplateInstanceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cluster template instance unauthorized response has a 4xx status code
func (o *CreateClusterTemplateInstanceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create cluster template instance unauthorized response has a 5xx status code
func (o *CreateClusterTemplateInstanceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create cluster template instance unauthorized response a status code equal to that given
func (o *CreateClusterTemplateInstanceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateClusterTemplateInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates/{template_id}/instances][%d] createClusterTemplateInstanceUnauthorized ", 401)
}

func (o *CreateClusterTemplateInstanceUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates/{template_id}/instances][%d] createClusterTemplateInstanceUnauthorized ", 401)
}

func (o *CreateClusterTemplateInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateClusterTemplateInstanceForbidden creates a CreateClusterTemplateInstanceForbidden with default headers values
func NewCreateClusterTemplateInstanceForbidden() *CreateClusterTemplateInstanceForbidden {
	return &CreateClusterTemplateInstanceForbidden{}
}

/*
CreateClusterTemplateInstanceForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateClusterTemplateInstanceForbidden struct {
}

// IsSuccess returns true when this create cluster template instance forbidden response has a 2xx status code
func (o *CreateClusterTemplateInstanceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cluster template instance forbidden response has a 3xx status code
func (o *CreateClusterTemplateInstanceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cluster template instance forbidden response has a 4xx status code
func (o *CreateClusterTemplateInstanceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create cluster template instance forbidden response has a 5xx status code
func (o *CreateClusterTemplateInstanceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create cluster template instance forbidden response a status code equal to that given
func (o *CreateClusterTemplateInstanceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateClusterTemplateInstanceForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates/{template_id}/instances][%d] createClusterTemplateInstanceForbidden ", 403)
}

func (o *CreateClusterTemplateInstanceForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates/{template_id}/instances][%d] createClusterTemplateInstanceForbidden ", 403)
}

func (o *CreateClusterTemplateInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateClusterTemplateInstanceDefault creates a CreateClusterTemplateInstanceDefault with default headers values
func NewCreateClusterTemplateInstanceDefault(code int) *CreateClusterTemplateInstanceDefault {
	return &CreateClusterTemplateInstanceDefault{
		_statusCode: code,
	}
}

/*
CreateClusterTemplateInstanceDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateClusterTemplateInstanceDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create cluster template instance default response
func (o *CreateClusterTemplateInstanceDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create cluster template instance default response has a 2xx status code
func (o *CreateClusterTemplateInstanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create cluster template instance default response has a 3xx status code
func (o *CreateClusterTemplateInstanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create cluster template instance default response has a 4xx status code
func (o *CreateClusterTemplateInstanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create cluster template instance default response has a 5xx status code
func (o *CreateClusterTemplateInstanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create cluster template instance default response a status code equal to that given
func (o *CreateClusterTemplateInstanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateClusterTemplateInstanceDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates/{template_id}/instances][%d] createClusterTemplateInstance default  %+v", o._statusCode, o.Payload)
}

func (o *CreateClusterTemplateInstanceDefault) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates/{template_id}/instances][%d] createClusterTemplateInstance default  %+v", o._statusCode, o.Payload)
}

func (o *CreateClusterTemplateInstanceDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateClusterTemplateInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateClusterTemplateInstanceBody create cluster template instance body
swagger:model CreateClusterTemplateInstanceBody
*/
type CreateClusterTemplateInstanceBody struct {

	// replicas
	Replicas int64 `json:"replicas,omitempty"`
}

// Validate validates this create cluster template instance body
func (o *CreateClusterTemplateInstanceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create cluster template instance body based on context it is used
func (o *CreateClusterTemplateInstanceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateClusterTemplateInstanceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateClusterTemplateInstanceBody) UnmarshalBinary(b []byte) error {
	var res CreateClusterTemplateInstanceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
