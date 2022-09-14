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

// IsSuccess returns true when this revoke cluster admin token o k response has a 2xx status code
func (o *RevokeClusterAdminTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this revoke cluster admin token o k response has a 3xx status code
func (o *RevokeClusterAdminTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revoke cluster admin token o k response has a 4xx status code
func (o *RevokeClusterAdminTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this revoke cluster admin token o k response has a 5xx status code
func (o *RevokeClusterAdminTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this revoke cluster admin token o k response a status code equal to that given
func (o *RevokeClusterAdminTokenOK) IsCode(code int) bool {
	return code == 200
}

func (o *RevokeClusterAdminTokenOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/token][%d] revokeClusterAdminTokenOK ", 200)
}

func (o *RevokeClusterAdminTokenOK) String() string {
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

// IsSuccess returns true when this revoke cluster admin token unauthorized response has a 2xx status code
func (o *RevokeClusterAdminTokenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this revoke cluster admin token unauthorized response has a 3xx status code
func (o *RevokeClusterAdminTokenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revoke cluster admin token unauthorized response has a 4xx status code
func (o *RevokeClusterAdminTokenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this revoke cluster admin token unauthorized response has a 5xx status code
func (o *RevokeClusterAdminTokenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this revoke cluster admin token unauthorized response a status code equal to that given
func (o *RevokeClusterAdminTokenUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *RevokeClusterAdminTokenUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/token][%d] revokeClusterAdminTokenUnauthorized ", 401)
}

func (o *RevokeClusterAdminTokenUnauthorized) String() string {
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

// IsSuccess returns true when this revoke cluster admin token forbidden response has a 2xx status code
func (o *RevokeClusterAdminTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this revoke cluster admin token forbidden response has a 3xx status code
func (o *RevokeClusterAdminTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revoke cluster admin token forbidden response has a 4xx status code
func (o *RevokeClusterAdminTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this revoke cluster admin token forbidden response has a 5xx status code
func (o *RevokeClusterAdminTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this revoke cluster admin token forbidden response a status code equal to that given
func (o *RevokeClusterAdminTokenForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *RevokeClusterAdminTokenForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/token][%d] revokeClusterAdminTokenForbidden ", 403)
}

func (o *RevokeClusterAdminTokenForbidden) String() string {
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

// IsSuccess returns true when this revoke cluster admin token default response has a 2xx status code
func (o *RevokeClusterAdminTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this revoke cluster admin token default response has a 3xx status code
func (o *RevokeClusterAdminTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this revoke cluster admin token default response has a 4xx status code
func (o *RevokeClusterAdminTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this revoke cluster admin token default response has a 5xx status code
func (o *RevokeClusterAdminTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this revoke cluster admin token default response a status code equal to that given
func (o *RevokeClusterAdminTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *RevokeClusterAdminTokenDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/token][%d] revokeClusterAdminToken default  %+v", o._statusCode, o.Payload)
}

func (o *RevokeClusterAdminTokenDefault) String() string {
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
