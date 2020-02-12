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

// ListNodeDeploymentsReader is a Reader for the ListNodeDeployments structure.
type ListNodeDeploymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNodeDeploymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNodeDeploymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListNodeDeploymentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListNodeDeploymentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListNodeDeploymentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListNodeDeploymentsOK creates a ListNodeDeploymentsOK with default headers values
func NewListNodeDeploymentsOK() *ListNodeDeploymentsOK {
	return &ListNodeDeploymentsOK{}
}

/*ListNodeDeploymentsOK handles this case with default header values.

NodeDeployment
*/
type ListNodeDeploymentsOK struct {
	Payload []*models.NodeDeployment
}

func (o *ListNodeDeploymentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments][%d] listNodeDeploymentsOK  %+v", 200, o.Payload)
}

func (o *ListNodeDeploymentsOK) GetPayload() []*models.NodeDeployment {
	return o.Payload
}

func (o *ListNodeDeploymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNodeDeploymentsUnauthorized creates a ListNodeDeploymentsUnauthorized with default headers values
func NewListNodeDeploymentsUnauthorized() *ListNodeDeploymentsUnauthorized {
	return &ListNodeDeploymentsUnauthorized{}
}

/*ListNodeDeploymentsUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type ListNodeDeploymentsUnauthorized struct {
}

func (o *ListNodeDeploymentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments][%d] listNodeDeploymentsUnauthorized ", 401)
}

func (o *ListNodeDeploymentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListNodeDeploymentsForbidden creates a ListNodeDeploymentsForbidden with default headers values
func NewListNodeDeploymentsForbidden() *ListNodeDeploymentsForbidden {
	return &ListNodeDeploymentsForbidden{}
}

/*ListNodeDeploymentsForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type ListNodeDeploymentsForbidden struct {
}

func (o *ListNodeDeploymentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments][%d] listNodeDeploymentsForbidden ", 403)
}

func (o *ListNodeDeploymentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListNodeDeploymentsDefault creates a ListNodeDeploymentsDefault with default headers values
func NewListNodeDeploymentsDefault(code int) *ListNodeDeploymentsDefault {
	return &ListNodeDeploymentsDefault{
		_statusCode: code,
	}
}

/*ListNodeDeploymentsDefault handles this case with default header values.

errorResponse
*/
type ListNodeDeploymentsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list node deployments default response
func (o *ListNodeDeploymentsDefault) Code() int {
	return o._statusCode
}

func (o *ListNodeDeploymentsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments][%d] listNodeDeployments default  %+v", o._statusCode, o.Payload)
}

func (o *ListNodeDeploymentsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListNodeDeploymentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
