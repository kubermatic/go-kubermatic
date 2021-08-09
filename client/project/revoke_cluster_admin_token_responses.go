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

// RevokeClusterAdminTokenReader is a Reader for the RevokeClusterAdminToken structure.
type RevokeClusterAdminTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeClusterAdminTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRevokeClusterAdminTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRevokeClusterAdminTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRevokeClusterAdminTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRevokeClusterAdminTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRevokeClusterAdminTokenOK creates a RevokeClusterAdminTokenOK with default headers values
func NewRevokeClusterAdminTokenOK() *RevokeClusterAdminTokenOK {
	return &RevokeClusterAdminTokenOK{}
}

/* RevokeClusterAdminTokenOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type RevokeClusterAdminTokenOK struct {
}

func (o *RevokeClusterAdminTokenOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/token][%d] revokeClusterAdminTokenOK ", 200)
}

func (o *RevokeClusterAdminTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeClusterAdminTokenUnauthorized creates a RevokeClusterAdminTokenUnauthorized with default headers values
func NewRevokeClusterAdminTokenUnauthorized() *RevokeClusterAdminTokenUnauthorized {
	return &RevokeClusterAdminTokenUnauthorized{}
}

/* RevokeClusterAdminTokenUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type RevokeClusterAdminTokenUnauthorized struct {
}

func (o *RevokeClusterAdminTokenUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/token][%d] revokeClusterAdminTokenUnauthorized ", 401)
}

func (o *RevokeClusterAdminTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeClusterAdminTokenForbidden creates a RevokeClusterAdminTokenForbidden with default headers values
func NewRevokeClusterAdminTokenForbidden() *RevokeClusterAdminTokenForbidden {
	return &RevokeClusterAdminTokenForbidden{}
}

/* RevokeClusterAdminTokenForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type RevokeClusterAdminTokenForbidden struct {
}

func (o *RevokeClusterAdminTokenForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/token][%d] revokeClusterAdminTokenForbidden ", 403)
}

func (o *RevokeClusterAdminTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeClusterAdminTokenDefault creates a RevokeClusterAdminTokenDefault with default headers values
func NewRevokeClusterAdminTokenDefault(code int) *RevokeClusterAdminTokenDefault {
	return &RevokeClusterAdminTokenDefault{
		_statusCode: code,
	}
}

/* RevokeClusterAdminTokenDefault describes a response with status code -1, with default header values.

errorResponse
*/
type RevokeClusterAdminTokenDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the revoke cluster admin token default response
func (o *RevokeClusterAdminTokenDefault) Code() int {
	return o._statusCode
}

func (o *RevokeClusterAdminTokenDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/token][%d] revokeClusterAdminToken default  %+v", o._statusCode, o.Payload)
}
func (o *RevokeClusterAdminTokenDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RevokeClusterAdminTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
