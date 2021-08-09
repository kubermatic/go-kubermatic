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

// UnbindUserFromClusterRoleBindingV2Reader is a Reader for the UnbindUserFromClusterRoleBindingV2 structure.
type UnbindUserFromClusterRoleBindingV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnbindUserFromClusterRoleBindingV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnbindUserFromClusterRoleBindingV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUnbindUserFromClusterRoleBindingV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUnbindUserFromClusterRoleBindingV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUnbindUserFromClusterRoleBindingV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnbindUserFromClusterRoleBindingV2OK creates a UnbindUserFromClusterRoleBindingV2OK with default headers values
func NewUnbindUserFromClusterRoleBindingV2OK() *UnbindUserFromClusterRoleBindingV2OK {
	return &UnbindUserFromClusterRoleBindingV2OK{}
}

/* UnbindUserFromClusterRoleBindingV2OK describes a response with status code 200, with default header values.

ClusterRoleBinding
*/
type UnbindUserFromClusterRoleBindingV2OK struct {
	Payload *models.ClusterRoleBinding
}

func (o *UnbindUserFromClusterRoleBindingV2OK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] unbindUserFromClusterRoleBindingV2OK  %+v", 200, o.Payload)
}
func (o *UnbindUserFromClusterRoleBindingV2OK) GetPayload() *models.ClusterRoleBinding {
	return o.Payload
}

func (o *UnbindUserFromClusterRoleBindingV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterRoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnbindUserFromClusterRoleBindingV2Unauthorized creates a UnbindUserFromClusterRoleBindingV2Unauthorized with default headers values
func NewUnbindUserFromClusterRoleBindingV2Unauthorized() *UnbindUserFromClusterRoleBindingV2Unauthorized {
	return &UnbindUserFromClusterRoleBindingV2Unauthorized{}
}

/* UnbindUserFromClusterRoleBindingV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type UnbindUserFromClusterRoleBindingV2Unauthorized struct {
}

func (o *UnbindUserFromClusterRoleBindingV2Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] unbindUserFromClusterRoleBindingV2Unauthorized ", 401)
}

func (o *UnbindUserFromClusterRoleBindingV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnbindUserFromClusterRoleBindingV2Forbidden creates a UnbindUserFromClusterRoleBindingV2Forbidden with default headers values
func NewUnbindUserFromClusterRoleBindingV2Forbidden() *UnbindUserFromClusterRoleBindingV2Forbidden {
	return &UnbindUserFromClusterRoleBindingV2Forbidden{}
}

/* UnbindUserFromClusterRoleBindingV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type UnbindUserFromClusterRoleBindingV2Forbidden struct {
}

func (o *UnbindUserFromClusterRoleBindingV2Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] unbindUserFromClusterRoleBindingV2Forbidden ", 403)
}

func (o *UnbindUserFromClusterRoleBindingV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnbindUserFromClusterRoleBindingV2Default creates a UnbindUserFromClusterRoleBindingV2Default with default headers values
func NewUnbindUserFromClusterRoleBindingV2Default(code int) *UnbindUserFromClusterRoleBindingV2Default {
	return &UnbindUserFromClusterRoleBindingV2Default{
		_statusCode: code,
	}
}

/* UnbindUserFromClusterRoleBindingV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type UnbindUserFromClusterRoleBindingV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the unbind user from cluster role binding v2 default response
func (o *UnbindUserFromClusterRoleBindingV2Default) Code() int {
	return o._statusCode
}

func (o *UnbindUserFromClusterRoleBindingV2Default) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] unbindUserFromClusterRoleBindingV2 default  %+v", o._statusCode, o.Payload)
}
func (o *UnbindUserFromClusterRoleBindingV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnbindUserFromClusterRoleBindingV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
