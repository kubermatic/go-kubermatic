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

// ListClusterRoleBindingV2Reader is a Reader for the ListClusterRoleBindingV2 structure.
type ListClusterRoleBindingV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClusterRoleBindingV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClusterRoleBindingV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListClusterRoleBindingV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListClusterRoleBindingV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListClusterRoleBindingV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListClusterRoleBindingV2OK creates a ListClusterRoleBindingV2OK with default headers values
func NewListClusterRoleBindingV2OK() *ListClusterRoleBindingV2OK {
	return &ListClusterRoleBindingV2OK{}
}

/* ListClusterRoleBindingV2OK describes a response with status code 200, with default header values.

ClusterRoleBinding
*/
type ListClusterRoleBindingV2OK struct {
	Payload []*models.ClusterRoleBinding
}

// IsSuccess returns true when this list cluster role binding v2 o k response has a 2xx status code
func (o *ListClusterRoleBindingV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list cluster role binding v2 o k response has a 3xx status code
func (o *ListClusterRoleBindingV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list cluster role binding v2 o k response has a 4xx status code
func (o *ListClusterRoleBindingV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list cluster role binding v2 o k response has a 5xx status code
func (o *ListClusterRoleBindingV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list cluster role binding v2 o k response a status code equal to that given
func (o *ListClusterRoleBindingV2OK) IsCode(code int) bool {
	return code == 200
}

func (o *ListClusterRoleBindingV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterbindings][%d] listClusterRoleBindingV2OK  %+v", 200, o.Payload)
}

func (o *ListClusterRoleBindingV2OK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterbindings][%d] listClusterRoleBindingV2OK  %+v", 200, o.Payload)
}

func (o *ListClusterRoleBindingV2OK) GetPayload() []*models.ClusterRoleBinding {
	return o.Payload
}

func (o *ListClusterRoleBindingV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterRoleBindingV2Unauthorized creates a ListClusterRoleBindingV2Unauthorized with default headers values
func NewListClusterRoleBindingV2Unauthorized() *ListClusterRoleBindingV2Unauthorized {
	return &ListClusterRoleBindingV2Unauthorized{}
}

/* ListClusterRoleBindingV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListClusterRoleBindingV2Unauthorized struct {
}

// IsSuccess returns true when this list cluster role binding v2 unauthorized response has a 2xx status code
func (o *ListClusterRoleBindingV2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list cluster role binding v2 unauthorized response has a 3xx status code
func (o *ListClusterRoleBindingV2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list cluster role binding v2 unauthorized response has a 4xx status code
func (o *ListClusterRoleBindingV2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list cluster role binding v2 unauthorized response has a 5xx status code
func (o *ListClusterRoleBindingV2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list cluster role binding v2 unauthorized response a status code equal to that given
func (o *ListClusterRoleBindingV2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListClusterRoleBindingV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterbindings][%d] listClusterRoleBindingV2Unauthorized ", 401)
}

func (o *ListClusterRoleBindingV2Unauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterbindings][%d] listClusterRoleBindingV2Unauthorized ", 401)
}

func (o *ListClusterRoleBindingV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClusterRoleBindingV2Forbidden creates a ListClusterRoleBindingV2Forbidden with default headers values
func NewListClusterRoleBindingV2Forbidden() *ListClusterRoleBindingV2Forbidden {
	return &ListClusterRoleBindingV2Forbidden{}
}

/* ListClusterRoleBindingV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListClusterRoleBindingV2Forbidden struct {
}

// IsSuccess returns true when this list cluster role binding v2 forbidden response has a 2xx status code
func (o *ListClusterRoleBindingV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list cluster role binding v2 forbidden response has a 3xx status code
func (o *ListClusterRoleBindingV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list cluster role binding v2 forbidden response has a 4xx status code
func (o *ListClusterRoleBindingV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list cluster role binding v2 forbidden response has a 5xx status code
func (o *ListClusterRoleBindingV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list cluster role binding v2 forbidden response a status code equal to that given
func (o *ListClusterRoleBindingV2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListClusterRoleBindingV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterbindings][%d] listClusterRoleBindingV2Forbidden ", 403)
}

func (o *ListClusterRoleBindingV2Forbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterbindings][%d] listClusterRoleBindingV2Forbidden ", 403)
}

func (o *ListClusterRoleBindingV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClusterRoleBindingV2Default creates a ListClusterRoleBindingV2Default with default headers values
func NewListClusterRoleBindingV2Default(code int) *ListClusterRoleBindingV2Default {
	return &ListClusterRoleBindingV2Default{
		_statusCode: code,
	}
}

/* ListClusterRoleBindingV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type ListClusterRoleBindingV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list cluster role binding v2 default response
func (o *ListClusterRoleBindingV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list cluster role binding v2 default response has a 2xx status code
func (o *ListClusterRoleBindingV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list cluster role binding v2 default response has a 3xx status code
func (o *ListClusterRoleBindingV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list cluster role binding v2 default response has a 4xx status code
func (o *ListClusterRoleBindingV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list cluster role binding v2 default response has a 5xx status code
func (o *ListClusterRoleBindingV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list cluster role binding v2 default response a status code equal to that given
func (o *ListClusterRoleBindingV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListClusterRoleBindingV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterbindings][%d] listClusterRoleBindingV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListClusterRoleBindingV2Default) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterbindings][%d] listClusterRoleBindingV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListClusterRoleBindingV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListClusterRoleBindingV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
