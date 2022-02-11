// Code generated by go-swagger; DO NOT EDIT.

package eks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListEKSVPCSReader is a Reader for the ListEKSVPCS structure.
type ListEKSVPCSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEKSVPCSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEKSVPCSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListEKSVPCSDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListEKSVPCSOK creates a ListEKSVPCSOK with default headers values
func NewListEKSVPCSOK() *ListEKSVPCSOK {
	return &ListEKSVPCSOK{}
}

/* ListEKSVPCSOK describes a response with status code 200, with default header values.

EKSVPCList
*/
type ListEKSVPCSOK struct {
	Payload models.EKSVPCList
}

func (o *ListEKSVPCSOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/vpcs][%d] listEKSVPCSOK  %+v", 200, o.Payload)
}
func (o *ListEKSVPCSOK) GetPayload() models.EKSVPCList {
	return o.Payload
}

func (o *ListEKSVPCSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEKSVPCSDefault creates a ListEKSVPCSDefault with default headers values
func NewListEKSVPCSDefault(code int) *ListEKSVPCSDefault {
	return &ListEKSVPCSDefault{
		_statusCode: code,
	}
}

/* ListEKSVPCSDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListEKSVPCSDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list e k s v p c s default response
func (o *ListEKSVPCSDefault) Code() int {
	return o._statusCode
}

func (o *ListEKSVPCSDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/vpcs][%d] listEKSVPCS default  %+v", o._statusCode, o.Payload)
}
func (o *ListEKSVPCSDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListEKSVPCSDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
