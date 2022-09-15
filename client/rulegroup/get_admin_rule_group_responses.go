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

// GetAdminRuleGroupReader is a Reader for the GetAdminRuleGroup structure.
type GetAdminRuleGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdminRuleGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAdminRuleGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAdminRuleGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAdminRuleGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAdminRuleGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAdminRuleGroupOK creates a GetAdminRuleGroupOK with default headers values
func NewGetAdminRuleGroupOK() *GetAdminRuleGroupOK {
	return &GetAdminRuleGroupOK{}
}

/*
GetAdminRuleGroupOK describes a response with status code 200, with default header values.

RuleGroup
*/
type GetAdminRuleGroupOK struct {
	Payload *models.RuleGroup
}

// IsSuccess returns true when this get admin rule group o k response has a 2xx status code
func (o *GetAdminRuleGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get admin rule group o k response has a 3xx status code
func (o *GetAdminRuleGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get admin rule group o k response has a 4xx status code
func (o *GetAdminRuleGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get admin rule group o k response has a 5xx status code
func (o *GetAdminRuleGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get admin rule group o k response a status code equal to that given
func (o *GetAdminRuleGroupOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAdminRuleGroupOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}][%d] getAdminRuleGroupOK  %+v", 200, o.Payload)
}

func (o *GetAdminRuleGroupOK) String() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}][%d] getAdminRuleGroupOK  %+v", 200, o.Payload)
}

func (o *GetAdminRuleGroupOK) GetPayload() *models.RuleGroup {
	return o.Payload
}

func (o *GetAdminRuleGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuleGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdminRuleGroupUnauthorized creates a GetAdminRuleGroupUnauthorized with default headers values
func NewGetAdminRuleGroupUnauthorized() *GetAdminRuleGroupUnauthorized {
	return &GetAdminRuleGroupUnauthorized{}
}

/*
GetAdminRuleGroupUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetAdminRuleGroupUnauthorized struct {
}

// IsSuccess returns true when this get admin rule group unauthorized response has a 2xx status code
func (o *GetAdminRuleGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get admin rule group unauthorized response has a 3xx status code
func (o *GetAdminRuleGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get admin rule group unauthorized response has a 4xx status code
func (o *GetAdminRuleGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get admin rule group unauthorized response has a 5xx status code
func (o *GetAdminRuleGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get admin rule group unauthorized response a status code equal to that given
func (o *GetAdminRuleGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAdminRuleGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}][%d] getAdminRuleGroupUnauthorized ", 401)
}

func (o *GetAdminRuleGroupUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}][%d] getAdminRuleGroupUnauthorized ", 401)
}

func (o *GetAdminRuleGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdminRuleGroupForbidden creates a GetAdminRuleGroupForbidden with default headers values
func NewGetAdminRuleGroupForbidden() *GetAdminRuleGroupForbidden {
	return &GetAdminRuleGroupForbidden{}
}

/*
GetAdminRuleGroupForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetAdminRuleGroupForbidden struct {
}

// IsSuccess returns true when this get admin rule group forbidden response has a 2xx status code
func (o *GetAdminRuleGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get admin rule group forbidden response has a 3xx status code
func (o *GetAdminRuleGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get admin rule group forbidden response has a 4xx status code
func (o *GetAdminRuleGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get admin rule group forbidden response has a 5xx status code
func (o *GetAdminRuleGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get admin rule group forbidden response a status code equal to that given
func (o *GetAdminRuleGroupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAdminRuleGroupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}][%d] getAdminRuleGroupForbidden ", 403)
}

func (o *GetAdminRuleGroupForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}][%d] getAdminRuleGroupForbidden ", 403)
}

func (o *GetAdminRuleGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdminRuleGroupDefault creates a GetAdminRuleGroupDefault with default headers values
func NewGetAdminRuleGroupDefault(code int) *GetAdminRuleGroupDefault {
	return &GetAdminRuleGroupDefault{
		_statusCode: code,
	}
}

/*
GetAdminRuleGroupDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetAdminRuleGroupDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get admin rule group default response
func (o *GetAdminRuleGroupDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get admin rule group default response has a 2xx status code
func (o *GetAdminRuleGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get admin rule group default response has a 3xx status code
func (o *GetAdminRuleGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get admin rule group default response has a 4xx status code
func (o *GetAdminRuleGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get admin rule group default response has a 5xx status code
func (o *GetAdminRuleGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get admin rule group default response a status code equal to that given
func (o *GetAdminRuleGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetAdminRuleGroupDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}][%d] getAdminRuleGroup default  %+v", o._statusCode, o.Payload)
}

func (o *GetAdminRuleGroupDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/rulegroups/{rulegroup_id}][%d] getAdminRuleGroup default  %+v", o._statusCode, o.Payload)
}

func (o *GetAdminRuleGroupDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAdminRuleGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
