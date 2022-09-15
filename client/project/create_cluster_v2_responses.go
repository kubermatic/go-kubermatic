// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// CreateClusterV2Reader is a Reader for the CreateClusterV2 structure.
type CreateClusterV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClusterV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateClusterV2Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateClusterV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateClusterV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateClusterV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateClusterV2Created creates a CreateClusterV2Created with default headers values
func NewCreateClusterV2Created() *CreateClusterV2Created {
	return &CreateClusterV2Created{}
}

/*
CreateClusterV2Created describes a response with status code 201, with default header values.

Cluster
*/
type CreateClusterV2Created struct {
	Payload *models.Cluster
}

// IsSuccess returns true when this create cluster v2 created response has a 2xx status code
func (o *CreateClusterV2Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create cluster v2 created response has a 3xx status code
func (o *CreateClusterV2Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cluster v2 created response has a 4xx status code
func (o *CreateClusterV2Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create cluster v2 created response has a 5xx status code
func (o *CreateClusterV2Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create cluster v2 created response a status code equal to that given
func (o *CreateClusterV2Created) IsCode(code int) bool {
	return code == 201
}

func (o *CreateClusterV2Created) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters][%d] createClusterV2Created  %+v", 201, o.Payload)
}

func (o *CreateClusterV2Created) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters][%d] createClusterV2Created  %+v", 201, o.Payload)
}

func (o *CreateClusterV2Created) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *CreateClusterV2Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterV2Unauthorized creates a CreateClusterV2Unauthorized with default headers values
func NewCreateClusterV2Unauthorized() *CreateClusterV2Unauthorized {
	return &CreateClusterV2Unauthorized{}
}

/*
CreateClusterV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateClusterV2Unauthorized struct {
}

// IsSuccess returns true when this create cluster v2 unauthorized response has a 2xx status code
func (o *CreateClusterV2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cluster v2 unauthorized response has a 3xx status code
func (o *CreateClusterV2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cluster v2 unauthorized response has a 4xx status code
func (o *CreateClusterV2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create cluster v2 unauthorized response has a 5xx status code
func (o *CreateClusterV2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create cluster v2 unauthorized response a status code equal to that given
func (o *CreateClusterV2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateClusterV2Unauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters][%d] createClusterV2Unauthorized ", 401)
}

func (o *CreateClusterV2Unauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters][%d] createClusterV2Unauthorized ", 401)
}

func (o *CreateClusterV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateClusterV2Forbidden creates a CreateClusterV2Forbidden with default headers values
func NewCreateClusterV2Forbidden() *CreateClusterV2Forbidden {
	return &CreateClusterV2Forbidden{}
}

/*
CreateClusterV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateClusterV2Forbidden struct {
}

// IsSuccess returns true when this create cluster v2 forbidden response has a 2xx status code
func (o *CreateClusterV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cluster v2 forbidden response has a 3xx status code
func (o *CreateClusterV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cluster v2 forbidden response has a 4xx status code
func (o *CreateClusterV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create cluster v2 forbidden response has a 5xx status code
func (o *CreateClusterV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create cluster v2 forbidden response a status code equal to that given
func (o *CreateClusterV2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateClusterV2Forbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters][%d] createClusterV2Forbidden ", 403)
}

func (o *CreateClusterV2Forbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters][%d] createClusterV2Forbidden ", 403)
}

func (o *CreateClusterV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateClusterV2Default creates a CreateClusterV2Default with default headers values
func NewCreateClusterV2Default(code int) *CreateClusterV2Default {
	return &CreateClusterV2Default{
		_statusCode: code,
	}
}

/*
CreateClusterV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type CreateClusterV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create cluster v2 default response
func (o *CreateClusterV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create cluster v2 default response has a 2xx status code
func (o *CreateClusterV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create cluster v2 default response has a 3xx status code
func (o *CreateClusterV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create cluster v2 default response has a 4xx status code
func (o *CreateClusterV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create cluster v2 default response has a 5xx status code
func (o *CreateClusterV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create cluster v2 default response a status code equal to that given
func (o *CreateClusterV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateClusterV2Default) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters][%d] createClusterV2 default  %+v", o._statusCode, o.Payload)
}

func (o *CreateClusterV2Default) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters][%d] createClusterV2 default  %+v", o._statusCode, o.Payload)
}

func (o *CreateClusterV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateClusterV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
