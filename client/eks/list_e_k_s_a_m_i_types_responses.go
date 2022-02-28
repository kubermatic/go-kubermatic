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

// ListEKSAMITypesReader is a Reader for the ListEKSAMITypes structure.
type ListEKSAMITypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEKSAMITypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEKSAMITypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListEKSAMITypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListEKSAMITypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListEKSAMITypesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListEKSAMITypesOK creates a ListEKSAMITypesOK with default headers values
func NewListEKSAMITypesOK() *ListEKSAMITypesOK {
	return &ListEKSAMITypesOK{}
}

/* ListEKSAMITypesOK describes a response with status code 200, with default header values.

EKSAMITypes
*/
type ListEKSAMITypesOK struct {
	Payload models.EKSAMITypes
}

func (o *ListEKSAMITypesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/eks/amitypes][%d] listEKSAMITypesOK  %+v", 200, o.Payload)
}
func (o *ListEKSAMITypesOK) GetPayload() models.EKSAMITypes {
	return o.Payload
}

func (o *ListEKSAMITypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEKSAMITypesUnauthorized creates a ListEKSAMITypesUnauthorized with default headers values
func NewListEKSAMITypesUnauthorized() *ListEKSAMITypesUnauthorized {
	return &ListEKSAMITypesUnauthorized{}
}

/* ListEKSAMITypesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListEKSAMITypesUnauthorized struct {
}

func (o *ListEKSAMITypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/eks/amitypes][%d] listEKSAMITypesUnauthorized ", 401)
}

func (o *ListEKSAMITypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSAMITypesForbidden creates a ListEKSAMITypesForbidden with default headers values
func NewListEKSAMITypesForbidden() *ListEKSAMITypesForbidden {
	return &ListEKSAMITypesForbidden{}
}

/* ListEKSAMITypesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListEKSAMITypesForbidden struct {
}

func (o *ListEKSAMITypesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/eks/amitypes][%d] listEKSAMITypesForbidden ", 403)
}

func (o *ListEKSAMITypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSAMITypesDefault creates a ListEKSAMITypesDefault with default headers values
func NewListEKSAMITypesDefault(code int) *ListEKSAMITypesDefault {
	return &ListEKSAMITypesDefault{
		_statusCode: code,
	}
}

/* ListEKSAMITypesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListEKSAMITypesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list e k s a m i types default response
func (o *ListEKSAMITypesDefault) Code() int {
	return o._statusCode
}

func (o *ListEKSAMITypesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/eks/amitypes][%d] listEKSAMITypes default  %+v", o._statusCode, o.Payload)
}
func (o *ListEKSAMITypesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListEKSAMITypesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}