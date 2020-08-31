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

// GetExternalClusterNodeReader is a Reader for the GetExternalClusterNode structure.
type GetExternalClusterNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalClusterNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalClusterNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetExternalClusterNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalClusterNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetExternalClusterNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetExternalClusterNodeOK creates a GetExternalClusterNodeOK with default headers values
func NewGetExternalClusterNodeOK() *GetExternalClusterNodeOK {
	return &GetExternalClusterNodeOK{}
}

/*GetExternalClusterNodeOK handles this case with default header values.

Node
*/
type GetExternalClusterNodeOK struct {
	Payload *models.Node
}

func (o *GetExternalClusterNodeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes/{node_id}][%d] getExternalClusterNodeOK  %+v", 200, o.Payload)
}

func (o *GetExternalClusterNodeOK) GetPayload() *models.Node {
	return o.Payload
}

func (o *GetExternalClusterNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Node)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalClusterNodeUnauthorized creates a GetExternalClusterNodeUnauthorized with default headers values
func NewGetExternalClusterNodeUnauthorized() *GetExternalClusterNodeUnauthorized {
	return &GetExternalClusterNodeUnauthorized{}
}

/*GetExternalClusterNodeUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type GetExternalClusterNodeUnauthorized struct {
}

func (o *GetExternalClusterNodeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes/{node_id}][%d] getExternalClusterNodeUnauthorized ", 401)
}

func (o *GetExternalClusterNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExternalClusterNodeForbidden creates a GetExternalClusterNodeForbidden with default headers values
func NewGetExternalClusterNodeForbidden() *GetExternalClusterNodeForbidden {
	return &GetExternalClusterNodeForbidden{}
}

/*GetExternalClusterNodeForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type GetExternalClusterNodeForbidden struct {
}

func (o *GetExternalClusterNodeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes/{node_id}][%d] getExternalClusterNodeForbidden ", 403)
}

func (o *GetExternalClusterNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExternalClusterNodeDefault creates a GetExternalClusterNodeDefault with default headers values
func NewGetExternalClusterNodeDefault(code int) *GetExternalClusterNodeDefault {
	return &GetExternalClusterNodeDefault{
		_statusCode: code,
	}
}

/*GetExternalClusterNodeDefault handles this case with default header values.

errorResponse
*/
type GetExternalClusterNodeDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get external cluster node default response
func (o *GetExternalClusterNodeDefault) Code() int {
	return o._statusCode
}

func (o *GetExternalClusterNodeDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes/{node_id}][%d] getExternalClusterNode default  %+v", o._statusCode, o.Payload)
}

func (o *GetExternalClusterNodeDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetExternalClusterNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
