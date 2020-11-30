// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListAWSSubnetsNoCredentialsV2Reader is a Reader for the ListAWSSubnetsNoCredentialsV2 structure.
type ListAWSSubnetsNoCredentialsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAWSSubnetsNoCredentialsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAWSSubnetsNoCredentialsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAWSSubnetsNoCredentialsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAWSSubnetsNoCredentialsV2OK creates a ListAWSSubnetsNoCredentialsV2OK with default headers values
func NewListAWSSubnetsNoCredentialsV2OK() *ListAWSSubnetsNoCredentialsV2OK {
	return &ListAWSSubnetsNoCredentialsV2OK{}
}

/*ListAWSSubnetsNoCredentialsV2OK handles this case with default header values.

AWSSubnetList
*/
type ListAWSSubnetsNoCredentialsV2OK struct {
	Payload models.AWSSubnetList
}

func (o *ListAWSSubnetsNoCredentialsV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/aws/subnets][%d] listAWSSubnetsNoCredentialsV2OK  %+v", 200, o.Payload)
}

func (o *ListAWSSubnetsNoCredentialsV2OK) GetPayload() models.AWSSubnetList {
	return o.Payload
}

func (o *ListAWSSubnetsNoCredentialsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAWSSubnetsNoCredentialsV2Default creates a ListAWSSubnetsNoCredentialsV2Default with default headers values
func NewListAWSSubnetsNoCredentialsV2Default(code int) *ListAWSSubnetsNoCredentialsV2Default {
	return &ListAWSSubnetsNoCredentialsV2Default{
		_statusCode: code,
	}
}

/*ListAWSSubnetsNoCredentialsV2Default handles this case with default header values.

errorResponse
*/
type ListAWSSubnetsNoCredentialsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list a w s subnets no credentials v2 default response
func (o *ListAWSSubnetsNoCredentialsV2Default) Code() int {
	return o._statusCode
}

func (o *ListAWSSubnetsNoCredentialsV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/aws/subnets][%d] listAWSSubnetsNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListAWSSubnetsNoCredentialsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAWSSubnetsNoCredentialsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
