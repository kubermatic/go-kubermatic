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

// ListRuleGroupsReader is a Reader for the ListRuleGroups structure.
type ListRuleGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRuleGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRuleGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListRuleGroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListRuleGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListRuleGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRuleGroupsOK creates a ListRuleGroupsOK with default headers values
func NewListRuleGroupsOK() *ListRuleGroupsOK {
	return &ListRuleGroupsOK{}
}

/*ListRuleGroupsOK handles this case with default header values.

RuleGroup
*/
type ListRuleGroupsOK struct {
	Payload []*models.RuleGroup
}

func (o *ListRuleGroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups][%d] listRuleGroupsOK  %+v", 200, o.Payload)
}

func (o *ListRuleGroupsOK) GetPayload() []*models.RuleGroup {
	return o.Payload
}

func (o *ListRuleGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRuleGroupsUnauthorized creates a ListRuleGroupsUnauthorized with default headers values
func NewListRuleGroupsUnauthorized() *ListRuleGroupsUnauthorized {
	return &ListRuleGroupsUnauthorized{}
}

/*ListRuleGroupsUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type ListRuleGroupsUnauthorized struct {
}

func (o *ListRuleGroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups][%d] listRuleGroupsUnauthorized ", 401)
}

func (o *ListRuleGroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRuleGroupsForbidden creates a ListRuleGroupsForbidden with default headers values
func NewListRuleGroupsForbidden() *ListRuleGroupsForbidden {
	return &ListRuleGroupsForbidden{}
}

/*ListRuleGroupsForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type ListRuleGroupsForbidden struct {
}

func (o *ListRuleGroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups][%d] listRuleGroupsForbidden ", 403)
}

func (o *ListRuleGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRuleGroupsDefault creates a ListRuleGroupsDefault with default headers values
func NewListRuleGroupsDefault(code int) *ListRuleGroupsDefault {
	return &ListRuleGroupsDefault{
		_statusCode: code,
	}
}

/*ListRuleGroupsDefault handles this case with default header values.

errorResponse
*/
type ListRuleGroupsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list rule groups default response
func (o *ListRuleGroupsDefault) Code() int {
	return o._statusCode
}

func (o *ListRuleGroupsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups][%d] listRuleGroups default  %+v", o._statusCode, o.Payload)
}

func (o *ListRuleGroupsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListRuleGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
