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

// DeleteClusterV2Reader is a Reader for the DeleteClusterV2 structure.
type DeleteClusterV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteClusterV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteClusterV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteClusterV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteClusterV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteClusterV2OK creates a DeleteClusterV2OK with default headers values
func NewDeleteClusterV2OK() *DeleteClusterV2OK {
	return &DeleteClusterV2OK{}
}

/* DeleteClusterV2OK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteClusterV2OK struct {
}

func (o *DeleteClusterV2OK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}][%d] deleteClusterV2OK ", 200)
}

func (o *DeleteClusterV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterV2Unauthorized creates a DeleteClusterV2Unauthorized with default headers values
func NewDeleteClusterV2Unauthorized() *DeleteClusterV2Unauthorized {
	return &DeleteClusterV2Unauthorized{}
}

/* DeleteClusterV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteClusterV2Unauthorized struct {
}

func (o *DeleteClusterV2Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}][%d] deleteClusterV2Unauthorized ", 401)
}

func (o *DeleteClusterV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterV2Forbidden creates a DeleteClusterV2Forbidden with default headers values
func NewDeleteClusterV2Forbidden() *DeleteClusterV2Forbidden {
	return &DeleteClusterV2Forbidden{}
}

/* DeleteClusterV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteClusterV2Forbidden struct {
}

func (o *DeleteClusterV2Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}][%d] deleteClusterV2Forbidden ", 403)
}

func (o *DeleteClusterV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterV2Default creates a DeleteClusterV2Default with default headers values
func NewDeleteClusterV2Default(code int) *DeleteClusterV2Default {
	return &DeleteClusterV2Default{
		_statusCode: code,
	}
}

/* DeleteClusterV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteClusterV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete cluster v2 default response
func (o *DeleteClusterV2Default) Code() int {
	return o._statusCode
}

func (o *DeleteClusterV2Default) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}][%d] deleteClusterV2 default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteClusterV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteClusterV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
