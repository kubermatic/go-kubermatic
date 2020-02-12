// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListNodeDeploymentNodesReader is a Reader for the ListNodeDeploymentNodes structure.
type ListNodeDeploymentNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNodeDeploymentNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNodeDeploymentNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListNodeDeploymentNodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListNodeDeploymentNodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListNodeDeploymentNodesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListNodeDeploymentNodesOK creates a ListNodeDeploymentNodesOK with default headers values
func NewListNodeDeploymentNodesOK() *ListNodeDeploymentNodesOK {
	return &ListNodeDeploymentNodesOK{}
}

/*ListNodeDeploymentNodesOK handles this case with default header values.

Node
*/
type ListNodeDeploymentNodesOK struct {
	Payload []*models.Node
}

func (o *ListNodeDeploymentNodesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}/nodes][%d] listNodeDeploymentNodesOK  %+v", 200, o.Payload)
}

func (o *ListNodeDeploymentNodesOK) GetPayload() []*models.Node {
	return o.Payload
}

func (o *ListNodeDeploymentNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNodeDeploymentNodesUnauthorized creates a ListNodeDeploymentNodesUnauthorized with default headers values
func NewListNodeDeploymentNodesUnauthorized() *ListNodeDeploymentNodesUnauthorized {
	return &ListNodeDeploymentNodesUnauthorized{}
}

/*ListNodeDeploymentNodesUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type ListNodeDeploymentNodesUnauthorized struct {
}

func (o *ListNodeDeploymentNodesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}/nodes][%d] listNodeDeploymentNodesUnauthorized ", 401)
}

func (o *ListNodeDeploymentNodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListNodeDeploymentNodesForbidden creates a ListNodeDeploymentNodesForbidden with default headers values
func NewListNodeDeploymentNodesForbidden() *ListNodeDeploymentNodesForbidden {
	return &ListNodeDeploymentNodesForbidden{}
}

/*ListNodeDeploymentNodesForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type ListNodeDeploymentNodesForbidden struct {
}

func (o *ListNodeDeploymentNodesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}/nodes][%d] listNodeDeploymentNodesForbidden ", 403)
}

func (o *ListNodeDeploymentNodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListNodeDeploymentNodesDefault creates a ListNodeDeploymentNodesDefault with default headers values
func NewListNodeDeploymentNodesDefault(code int) *ListNodeDeploymentNodesDefault {
	return &ListNodeDeploymentNodesDefault{
		_statusCode: code,
	}
}

/*ListNodeDeploymentNodesDefault handles this case with default header values.

errorResponse
*/
type ListNodeDeploymentNodesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list node deployment nodes default response
func (o *ListNodeDeploymentNodesDefault) Code() int {
	return o._statusCode
}

func (o *ListNodeDeploymentNodesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}/nodes][%d] listNodeDeploymentNodes default  %+v", o._statusCode, o.Payload)
}

func (o *ListNodeDeploymentNodesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListNodeDeploymentNodesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
