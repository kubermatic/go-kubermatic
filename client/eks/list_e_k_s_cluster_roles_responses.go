// Code generated by go-swagger; DO NOT EDIT.

package eks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListEKSClusterRolesReader is a Reader for the ListEKSClusterRoles structure.
type ListEKSClusterRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEKSClusterRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEKSClusterRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListEKSClusterRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListEKSClusterRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListEKSClusterRolesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListEKSClusterRolesOK creates a ListEKSClusterRolesOK with default headers values
func NewListEKSClusterRolesOK() *ListEKSClusterRolesOK {
	return &ListEKSClusterRolesOK{}
}

/* ListEKSClusterRolesOK describes a response with status code 200, with default header values.

EKSClusterRolesList
*/
type ListEKSClusterRolesOK struct {
	Payload models.EKSClusterRolesList
}

func (o *ListEKSClusterRolesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/clusterroles][%d] listEKSClusterRolesOK  %+v", 200, o.Payload)
}
func (o *ListEKSClusterRolesOK) GetPayload() models.EKSClusterRolesList {
	return o.Payload
}

func (o *ListEKSClusterRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEKSClusterRolesUnauthorized creates a ListEKSClusterRolesUnauthorized with default headers values
func NewListEKSClusterRolesUnauthorized() *ListEKSClusterRolesUnauthorized {
	return &ListEKSClusterRolesUnauthorized{}
}

/* ListEKSClusterRolesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListEKSClusterRolesUnauthorized struct {
}

func (o *ListEKSClusterRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/clusterroles][%d] listEKSClusterRolesUnauthorized ", 401)
}

func (o *ListEKSClusterRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSClusterRolesForbidden creates a ListEKSClusterRolesForbidden with default headers values
func NewListEKSClusterRolesForbidden() *ListEKSClusterRolesForbidden {
	return &ListEKSClusterRolesForbidden{}
}

/* ListEKSClusterRolesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListEKSClusterRolesForbidden struct {
}

func (o *ListEKSClusterRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/clusterroles][%d] listEKSClusterRolesForbidden ", 403)
}

func (o *ListEKSClusterRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSClusterRolesDefault creates a ListEKSClusterRolesDefault with default headers values
func NewListEKSClusterRolesDefault(code int) *ListEKSClusterRolesDefault {
	return &ListEKSClusterRolesDefault{
		_statusCode: code,
	}
}

/* ListEKSClusterRolesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListEKSClusterRolesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list e k s cluster roles default response
func (o *ListEKSClusterRolesDefault) Code() int {
	return o._statusCode
}

func (o *ListEKSClusterRolesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/clusterroles][%d] listEKSClusterRoles default  %+v", o._statusCode, o.Payload)
}
func (o *ListEKSClusterRolesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListEKSClusterRolesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}