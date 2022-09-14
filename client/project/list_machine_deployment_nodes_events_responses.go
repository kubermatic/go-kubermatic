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

// ListMachineDeploymentNodesEventsReader is a Reader for the ListMachineDeploymentNodesEvents structure.
type ListMachineDeploymentNodesEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListMachineDeploymentNodesEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListMachineDeploymentNodesEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListMachineDeploymentNodesEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListMachineDeploymentNodesEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListMachineDeploymentNodesEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListMachineDeploymentNodesEventsOK creates a ListMachineDeploymentNodesEventsOK with default headers values
func NewListMachineDeploymentNodesEventsOK() *ListMachineDeploymentNodesEventsOK {
	return &ListMachineDeploymentNodesEventsOK{}
}

/* ListMachineDeploymentNodesEventsOK describes a response with status code 200, with default header values.

Event
*/
type ListMachineDeploymentNodesEventsOK struct {
	Payload []*models.Event
}

// IsSuccess returns true when this list machine deployment nodes events o k response has a 2xx status code
func (o *ListMachineDeploymentNodesEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list machine deployment nodes events o k response has a 3xx status code
func (o *ListMachineDeploymentNodesEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list machine deployment nodes events o k response has a 4xx status code
func (o *ListMachineDeploymentNodesEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list machine deployment nodes events o k response has a 5xx status code
func (o *ListMachineDeploymentNodesEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list machine deployment nodes events o k response a status code equal to that given
func (o *ListMachineDeploymentNodesEventsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListMachineDeploymentNodesEventsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/events][%d] listMachineDeploymentNodesEventsOK  %+v", 200, o.Payload)
}

func (o *ListMachineDeploymentNodesEventsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/events][%d] listMachineDeploymentNodesEventsOK  %+v", 200, o.Payload)
}

func (o *ListMachineDeploymentNodesEventsOK) GetPayload() []*models.Event {
	return o.Payload
}

func (o *ListMachineDeploymentNodesEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMachineDeploymentNodesEventsUnauthorized creates a ListMachineDeploymentNodesEventsUnauthorized with default headers values
func NewListMachineDeploymentNodesEventsUnauthorized() *ListMachineDeploymentNodesEventsUnauthorized {
	return &ListMachineDeploymentNodesEventsUnauthorized{}
}

/* ListMachineDeploymentNodesEventsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListMachineDeploymentNodesEventsUnauthorized struct {
}

// IsSuccess returns true when this list machine deployment nodes events unauthorized response has a 2xx status code
func (o *ListMachineDeploymentNodesEventsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list machine deployment nodes events unauthorized response has a 3xx status code
func (o *ListMachineDeploymentNodesEventsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list machine deployment nodes events unauthorized response has a 4xx status code
func (o *ListMachineDeploymentNodesEventsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list machine deployment nodes events unauthorized response has a 5xx status code
func (o *ListMachineDeploymentNodesEventsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list machine deployment nodes events unauthorized response a status code equal to that given
func (o *ListMachineDeploymentNodesEventsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListMachineDeploymentNodesEventsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/events][%d] listMachineDeploymentNodesEventsUnauthorized ", 401)
}

func (o *ListMachineDeploymentNodesEventsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/events][%d] listMachineDeploymentNodesEventsUnauthorized ", 401)
}

func (o *ListMachineDeploymentNodesEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListMachineDeploymentNodesEventsForbidden creates a ListMachineDeploymentNodesEventsForbidden with default headers values
func NewListMachineDeploymentNodesEventsForbidden() *ListMachineDeploymentNodesEventsForbidden {
	return &ListMachineDeploymentNodesEventsForbidden{}
}

/* ListMachineDeploymentNodesEventsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListMachineDeploymentNodesEventsForbidden struct {
}

// IsSuccess returns true when this list machine deployment nodes events forbidden response has a 2xx status code
func (o *ListMachineDeploymentNodesEventsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list machine deployment nodes events forbidden response has a 3xx status code
func (o *ListMachineDeploymentNodesEventsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list machine deployment nodes events forbidden response has a 4xx status code
func (o *ListMachineDeploymentNodesEventsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list machine deployment nodes events forbidden response has a 5xx status code
func (o *ListMachineDeploymentNodesEventsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list machine deployment nodes events forbidden response a status code equal to that given
func (o *ListMachineDeploymentNodesEventsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListMachineDeploymentNodesEventsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/events][%d] listMachineDeploymentNodesEventsForbidden ", 403)
}

func (o *ListMachineDeploymentNodesEventsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/events][%d] listMachineDeploymentNodesEventsForbidden ", 403)
}

func (o *ListMachineDeploymentNodesEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListMachineDeploymentNodesEventsDefault creates a ListMachineDeploymentNodesEventsDefault with default headers values
func NewListMachineDeploymentNodesEventsDefault(code int) *ListMachineDeploymentNodesEventsDefault {
	return &ListMachineDeploymentNodesEventsDefault{
		_statusCode: code,
	}
}

/* ListMachineDeploymentNodesEventsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListMachineDeploymentNodesEventsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list machine deployment nodes events default response
func (o *ListMachineDeploymentNodesEventsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list machine deployment nodes events default response has a 2xx status code
func (o *ListMachineDeploymentNodesEventsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list machine deployment nodes events default response has a 3xx status code
func (o *ListMachineDeploymentNodesEventsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list machine deployment nodes events default response has a 4xx status code
func (o *ListMachineDeploymentNodesEventsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list machine deployment nodes events default response has a 5xx status code
func (o *ListMachineDeploymentNodesEventsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list machine deployment nodes events default response a status code equal to that given
func (o *ListMachineDeploymentNodesEventsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListMachineDeploymentNodesEventsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/events][%d] listMachineDeploymentNodesEvents default  %+v", o._statusCode, o.Payload)
}

func (o *ListMachineDeploymentNodesEventsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/events][%d] listMachineDeploymentNodesEvents default  %+v", o._statusCode, o.Payload)
}

func (o *ListMachineDeploymentNodesEventsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListMachineDeploymentNodesEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
