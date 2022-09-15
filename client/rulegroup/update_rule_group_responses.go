// Code generated by go-swagger; DO NOT EDIT.

package rulegroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// UpdateRuleGroupReader is a Reader for the UpdateRuleGroup structure.
type UpdateRuleGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRuleGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRuleGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRuleGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRuleGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateRuleGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateRuleGroupOK creates a UpdateRuleGroupOK with default headers values
func NewUpdateRuleGroupOK() *UpdateRuleGroupOK {
	return &UpdateRuleGroupOK{}
}

/*
UpdateRuleGroupOK describes a response with status code 200, with default header values.

RuleGroup
*/
type UpdateRuleGroupOK struct {
	Payload *models.RuleGroup
}

// IsSuccess returns true when this update rule group o k response has a 2xx status code
func (o *UpdateRuleGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update rule group o k response has a 3xx status code
func (o *UpdateRuleGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rule group o k response has a 4xx status code
func (o *UpdateRuleGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update rule group o k response has a 5xx status code
func (o *UpdateRuleGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update rule group o k response a status code equal to that given
func (o *UpdateRuleGroupOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateRuleGroupOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}][%d] updateRuleGroupOK  %+v", 200, o.Payload)
}

func (o *UpdateRuleGroupOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}][%d] updateRuleGroupOK  %+v", 200, o.Payload)
}

func (o *UpdateRuleGroupOK) GetPayload() *models.RuleGroup {
	return o.Payload
}

func (o *UpdateRuleGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuleGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRuleGroupUnauthorized creates a UpdateRuleGroupUnauthorized with default headers values
func NewUpdateRuleGroupUnauthorized() *UpdateRuleGroupUnauthorized {
	return &UpdateRuleGroupUnauthorized{}
}

/*
UpdateRuleGroupUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type UpdateRuleGroupUnauthorized struct {
}

// IsSuccess returns true when this update rule group unauthorized response has a 2xx status code
func (o *UpdateRuleGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rule group unauthorized response has a 3xx status code
func (o *UpdateRuleGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rule group unauthorized response has a 4xx status code
func (o *UpdateRuleGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update rule group unauthorized response has a 5xx status code
func (o *UpdateRuleGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update rule group unauthorized response a status code equal to that given
func (o *UpdateRuleGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateRuleGroupUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}][%d] updateRuleGroupUnauthorized ", 401)
}

func (o *UpdateRuleGroupUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}][%d] updateRuleGroupUnauthorized ", 401)
}

func (o *UpdateRuleGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRuleGroupForbidden creates a UpdateRuleGroupForbidden with default headers values
func NewUpdateRuleGroupForbidden() *UpdateRuleGroupForbidden {
	return &UpdateRuleGroupForbidden{}
}

/*
UpdateRuleGroupForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type UpdateRuleGroupForbidden struct {
}

// IsSuccess returns true when this update rule group forbidden response has a 2xx status code
func (o *UpdateRuleGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rule group forbidden response has a 3xx status code
func (o *UpdateRuleGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rule group forbidden response has a 4xx status code
func (o *UpdateRuleGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update rule group forbidden response has a 5xx status code
func (o *UpdateRuleGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update rule group forbidden response a status code equal to that given
func (o *UpdateRuleGroupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateRuleGroupForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}][%d] updateRuleGroupForbidden ", 403)
}

func (o *UpdateRuleGroupForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}][%d] updateRuleGroupForbidden ", 403)
}

func (o *UpdateRuleGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRuleGroupDefault creates a UpdateRuleGroupDefault with default headers values
func NewUpdateRuleGroupDefault(code int) *UpdateRuleGroupDefault {
	return &UpdateRuleGroupDefault{
		_statusCode: code,
	}
}

/*
UpdateRuleGroupDefault describes a response with status code -1, with default header values.

errorResponse
*/
type UpdateRuleGroupDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update rule group default response
func (o *UpdateRuleGroupDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update rule group default response has a 2xx status code
func (o *UpdateRuleGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update rule group default response has a 3xx status code
func (o *UpdateRuleGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update rule group default response has a 4xx status code
func (o *UpdateRuleGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update rule group default response has a 5xx status code
func (o *UpdateRuleGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update rule group default response a status code equal to that given
func (o *UpdateRuleGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateRuleGroupDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}][%d] updateRuleGroup default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRuleGroupDefault) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}][%d] updateRuleGroup default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRuleGroupDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateRuleGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
