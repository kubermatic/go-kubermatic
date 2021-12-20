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

// ListGKEClustersReader is a Reader for the ListGKEClusters structure.
type ListGKEClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGKEClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGKEClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGKEClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGKEClustersOK creates a ListGKEClustersOK with default headers values
func NewListGKEClustersOK() *ListGKEClustersOK {
	return &ListGKEClustersOK{}
}

/* ListGKEClustersOK describes a response with status code 200, with default header values.

GKEClusterList
*/
type ListGKEClustersOK struct {
	Payload models.GKEClusterList
}

func (o *ListGKEClustersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gke/clusters][%d] listGKEClustersOK  %+v", 200, o.Payload)
}
func (o *ListGKEClustersOK) GetPayload() models.GKEClusterList {
	return o.Payload
}

func (o *ListGKEClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGKEClustersDefault creates a ListGKEClustersDefault with default headers values
func NewListGKEClustersDefault(code int) *ListGKEClustersDefault {
	return &ListGKEClustersDefault{
		_statusCode: code,
	}
}

/* ListGKEClustersDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListGKEClustersDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list g k e clusters default response
func (o *ListGKEClustersDefault) Code() int {
	return o._statusCode
}

func (o *ListGKEClustersDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gke/clusters][%d] listGKEClusters default  %+v", o._statusCode, o.Payload)
}
func (o *ListGKEClustersDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListGKEClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}