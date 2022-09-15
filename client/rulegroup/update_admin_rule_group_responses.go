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

// UpdateAdminRuleGroupReader is a Reader for the UpdateAdminRuleGroup structure.
type UpdateAdminRuleGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAdminRuleGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAdminRuleGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateAdminRuleGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAdminRuleGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateAdminRuleGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAdminRuleGroupOK creates a UpdateAdminRuleGroupOK with default headers values
func NewUpdateAdminRuleGroupOK() *UpdateAdminRuleGroupOK {
	return &UpdateAdminRuleGroupOK{}
}

/*
UpdateAdminRuleGroupOK describes a response with status code 200, with default header values.

RuleGroup
*/
type UpdateAdminRuleGroupOK struct {
	Payload *models.RuleGroup
}

// IsSuccess returns true when this update admin rule group o k response has a 2xx status code
func (o *UpdateAdminRuleGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update admin rule group o k response has a 3xx status code
func (o *UpdateAdminRuleGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update admin rule group o k response has a 4xx status code
func (o *UpdateAdminRuleGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update admin rule group o k response has a 5xx status code
func (o *UpdateAdminRuleGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update admin rule group o k response a status code equal to that given
func (o *UpdateAdminRuleGroupOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateAdminRuleGroupOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}][%d] updateAdminRuleGroupOK  %+v", 200, o.Payload)
}

func (o *UpdateAdminRuleGroupOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}][%d] updateAdminRuleGroupOK  %+v", 200, o.Payload)
}

func (o *UpdateAdminRuleGroupOK) GetPayload() *models.RuleGroup {
	return o.Payload
}

func (o *UpdateAdminRuleGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuleGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAdminRuleGroupUnauthorized creates a UpdateAdminRuleGroupUnauthorized with default headers values
func NewUpdateAdminRuleGroupUnauthorized() *UpdateAdminRuleGroupUnauthorized {
	return &UpdateAdminRuleGroupUnauthorized{}
}

/*
UpdateAdminRuleGroupUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type UpdateAdminRuleGroupUnauthorized struct {
}

// IsSuccess returns true when this update admin rule group unauthorized response has a 2xx status code
func (o *UpdateAdminRuleGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update admin rule group unauthorized response has a 3xx status code
func (o *UpdateAdminRuleGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update admin rule group unauthorized response has a 4xx status code
func (o *UpdateAdminRuleGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update admin rule group unauthorized response has a 5xx status code
func (o *UpdateAdminRuleGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update admin rule group unauthorized response a status code equal to that given
func (o *UpdateAdminRuleGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateAdminRuleGroupUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}][%d] updateAdminRuleGroupUnauthorized ", 401)
}

func (o *UpdateAdminRuleGroupUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}][%d] updateAdminRuleGroupUnauthorized ", 401)
}

func (o *UpdateAdminRuleGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAdminRuleGroupForbidden creates a UpdateAdminRuleGroupForbidden with default headers values
func NewUpdateAdminRuleGroupForbidden() *UpdateAdminRuleGroupForbidden {
	return &UpdateAdminRuleGroupForbidden{}
}

/*
UpdateAdminRuleGroupForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type UpdateAdminRuleGroupForbidden struct {
}

// IsSuccess returns true when this update admin rule group forbidden response has a 2xx status code
func (o *UpdateAdminRuleGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update admin rule group forbidden response has a 3xx status code
func (o *UpdateAdminRuleGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update admin rule group forbidden response has a 4xx status code
func (o *UpdateAdminRuleGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update admin rule group forbidden response has a 5xx status code
func (o *UpdateAdminRuleGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update admin rule group forbidden response a status code equal to that given
func (o *UpdateAdminRuleGroupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateAdminRuleGroupForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}][%d] updateAdminRuleGroupForbidden ", 403)
}

func (o *UpdateAdminRuleGroupForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}][%d] updateAdminRuleGroupForbidden ", 403)
}

func (o *UpdateAdminRuleGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAdminRuleGroupDefault creates a UpdateAdminRuleGroupDefault with default headers values
func NewUpdateAdminRuleGroupDefault(code int) *UpdateAdminRuleGroupDefault {
	return &UpdateAdminRuleGroupDefault{
		_statusCode: code,
	}
}

/*
UpdateAdminRuleGroupDefault describes a response with status code -1, with default header values.

errorResponse
*/
type UpdateAdminRuleGroupDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update admin rule group default response
func (o *UpdateAdminRuleGroupDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update admin rule group default response has a 2xx status code
func (o *UpdateAdminRuleGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update admin rule group default response has a 3xx status code
func (o *UpdateAdminRuleGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update admin rule group default response has a 4xx status code
func (o *UpdateAdminRuleGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update admin rule group default response has a 5xx status code
func (o *UpdateAdminRuleGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update admin rule group default response a status code equal to that given
func (o *UpdateAdminRuleGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateAdminRuleGroupDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}][%d] updateAdminRuleGroup default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAdminRuleGroupDefault) String() string {
	return fmt.Sprintf("[PUT /api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}][%d] updateAdminRuleGroup default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAdminRuleGroupDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateAdminRuleGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
