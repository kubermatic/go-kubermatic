// Code generated by go-swagger; DO NOT EDIT.

package gcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kubermatic/go-kubermatic/models"
)

// ListGCPSubnetworksNoCredentialsReader is a Reader for the ListGCPSubnetworksNoCredentials structure.
type ListGCPSubnetworksNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGCPSubnetworksNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListGCPSubnetworksNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListGCPSubnetworksNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGCPSubnetworksNoCredentialsOK creates a ListGCPSubnetworksNoCredentialsOK with default headers values
func NewListGCPSubnetworksNoCredentialsOK() *ListGCPSubnetworksNoCredentialsOK {
	return &ListGCPSubnetworksNoCredentialsOK{}
}

/*ListGCPSubnetworksNoCredentialsOK handles this case with default header values.

GCPSubnetworkList
*/
type ListGCPSubnetworksNoCredentialsOK struct {
	Payload models.GCPSubnetworkList
}

func (o *ListGCPSubnetworksNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/subnetworks][%d] listGCPSubnetworksNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListGCPSubnetworksNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGCPSubnetworksNoCredentialsDefault creates a ListGCPSubnetworksNoCredentialsDefault with default headers values
func NewListGCPSubnetworksNoCredentialsDefault(code int) *ListGCPSubnetworksNoCredentialsDefault {
	return &ListGCPSubnetworksNoCredentialsDefault{
		_statusCode: code,
	}
}

/*ListGCPSubnetworksNoCredentialsDefault handles this case with default header values.

errorResponse
*/
type ListGCPSubnetworksNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list g c p subnetworks no credentials default response
func (o *ListGCPSubnetworksNoCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *ListGCPSubnetworksNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/subnetworks][%d] listGCPSubnetworksNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListGCPSubnetworksNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
