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

// RevokeClusterAdminTokenV2Reader is a Reader for the RevokeClusterAdminTokenV2 structure.
type RevokeClusterAdminTokenV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeClusterAdminTokenV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRevokeClusterAdminTokenV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRevokeClusterAdminTokenV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRevokeClusterAdminTokenV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRevokeClusterAdminTokenV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRevokeClusterAdminTokenV2OK creates a RevokeClusterAdminTokenV2OK with default headers values
func NewRevokeClusterAdminTokenV2OK() *RevokeClusterAdminTokenV2OK {
	return &RevokeClusterAdminTokenV2OK{}
}

/* RevokeClusterAdminTokenV2OK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type RevokeClusterAdminTokenV2OK struct {
}

// IsSuccess returns true when this revoke cluster admin token v2 o k response has a 2xx status code
func (o *RevokeClusterAdminTokenV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this revoke cluster admin token v2 o k response has a 3xx status code
func (o *RevokeClusterAdminTokenV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revoke cluster admin token v2 o k response has a 4xx status code
func (o *RevokeClusterAdminTokenV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this revoke cluster admin token v2 o k response has a 5xx status code
func (o *RevokeClusterAdminTokenV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this revoke cluster admin token v2 o k response a status code equal to that given
func (o *RevokeClusterAdminTokenV2OK) IsCode(code int) bool {
	return code == 200
}

func (o *RevokeClusterAdminTokenV2OK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/token][%d] revokeClusterAdminTokenV2OK ", 200)
}

func (o *RevokeClusterAdminTokenV2OK) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/token][%d] revokeClusterAdminTokenV2OK ", 200)
}

func (o *RevokeClusterAdminTokenV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeClusterAdminTokenV2Unauthorized creates a RevokeClusterAdminTokenV2Unauthorized with default headers values
func NewRevokeClusterAdminTokenV2Unauthorized() *RevokeClusterAdminTokenV2Unauthorized {
	return &RevokeClusterAdminTokenV2Unauthorized{}
}

/* RevokeClusterAdminTokenV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type RevokeClusterAdminTokenV2Unauthorized struct {
}

// IsSuccess returns true when this revoke cluster admin token v2 unauthorized response has a 2xx status code
func (o *RevokeClusterAdminTokenV2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this revoke cluster admin token v2 unauthorized response has a 3xx status code
func (o *RevokeClusterAdminTokenV2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revoke cluster admin token v2 unauthorized response has a 4xx status code
func (o *RevokeClusterAdminTokenV2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this revoke cluster admin token v2 unauthorized response has a 5xx status code
func (o *RevokeClusterAdminTokenV2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this revoke cluster admin token v2 unauthorized response a status code equal to that given
func (o *RevokeClusterAdminTokenV2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *RevokeClusterAdminTokenV2Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/token][%d] revokeClusterAdminTokenV2Unauthorized ", 401)
}

func (o *RevokeClusterAdminTokenV2Unauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/token][%d] revokeClusterAdminTokenV2Unauthorized ", 401)
}

func (o *RevokeClusterAdminTokenV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeClusterAdminTokenV2Forbidden creates a RevokeClusterAdminTokenV2Forbidden with default headers values
func NewRevokeClusterAdminTokenV2Forbidden() *RevokeClusterAdminTokenV2Forbidden {
	return &RevokeClusterAdminTokenV2Forbidden{}
}

/* RevokeClusterAdminTokenV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type RevokeClusterAdminTokenV2Forbidden struct {
}

// IsSuccess returns true when this revoke cluster admin token v2 forbidden response has a 2xx status code
func (o *RevokeClusterAdminTokenV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this revoke cluster admin token v2 forbidden response has a 3xx status code
func (o *RevokeClusterAdminTokenV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revoke cluster admin token v2 forbidden response has a 4xx status code
func (o *RevokeClusterAdminTokenV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this revoke cluster admin token v2 forbidden response has a 5xx status code
func (o *RevokeClusterAdminTokenV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this revoke cluster admin token v2 forbidden response a status code equal to that given
func (o *RevokeClusterAdminTokenV2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *RevokeClusterAdminTokenV2Forbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/token][%d] revokeClusterAdminTokenV2Forbidden ", 403)
}

func (o *RevokeClusterAdminTokenV2Forbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/token][%d] revokeClusterAdminTokenV2Forbidden ", 403)
}

func (o *RevokeClusterAdminTokenV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeClusterAdminTokenV2Default creates a RevokeClusterAdminTokenV2Default with default headers values
func NewRevokeClusterAdminTokenV2Default(code int) *RevokeClusterAdminTokenV2Default {
	return &RevokeClusterAdminTokenV2Default{
		_statusCode: code,
	}
}

/* RevokeClusterAdminTokenV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type RevokeClusterAdminTokenV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the revoke cluster admin token v2 default response
func (o *RevokeClusterAdminTokenV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this revoke cluster admin token v2 default response has a 2xx status code
func (o *RevokeClusterAdminTokenV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this revoke cluster admin token v2 default response has a 3xx status code
func (o *RevokeClusterAdminTokenV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this revoke cluster admin token v2 default response has a 4xx status code
func (o *RevokeClusterAdminTokenV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this revoke cluster admin token v2 default response has a 5xx status code
func (o *RevokeClusterAdminTokenV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this revoke cluster admin token v2 default response a status code equal to that given
func (o *RevokeClusterAdminTokenV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *RevokeClusterAdminTokenV2Default) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/token][%d] revokeClusterAdminTokenV2 default  %+v", o._statusCode, o.Payload)
}

func (o *RevokeClusterAdminTokenV2Default) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/token][%d] revokeClusterAdminTokenV2 default  %+v", o._statusCode, o.Payload)
}

func (o *RevokeClusterAdminTokenV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RevokeClusterAdminTokenV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
