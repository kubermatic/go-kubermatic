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

// ListEKSVersionsReader is a Reader for the ListEKSVersions structure.
type ListEKSVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEKSVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEKSVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListEKSVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListEKSVersionsOK creates a ListEKSVersionsOK with default headers values
func NewListEKSVersionsOK() *ListEKSVersionsOK {
	return &ListEKSVersionsOK{}
}

/* ListEKSVersionsOK describes a response with status code 200, with default header values.

MasterVersion
*/
type ListEKSVersionsOK struct {
	Payload []*models.MasterVersion
}

// IsSuccess returns true when this list e k s versions o k response has a 2xx status code
func (o *ListEKSVersionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list e k s versions o k response has a 3xx status code
func (o *ListEKSVersionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s versions o k response has a 4xx status code
func (o *ListEKSVersionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list e k s versions o k response has a 5xx status code
func (o *ListEKSVersionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s versions o k response a status code equal to that given
func (o *ListEKSVersionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListEKSVersionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/versions][%d] listEKSVersionsOK  %+v", 200, o.Payload)
}

func (o *ListEKSVersionsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/versions][%d] listEKSVersionsOK  %+v", 200, o.Payload)
}

func (o *ListEKSVersionsOK) GetPayload() []*models.MasterVersion {
	return o.Payload
}

func (o *ListEKSVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEKSVersionsDefault creates a ListEKSVersionsDefault with default headers values
func NewListEKSVersionsDefault(code int) *ListEKSVersionsDefault {
	return &ListEKSVersionsDefault{
		_statusCode: code,
	}
}

/* ListEKSVersionsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListEKSVersionsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list e k s versions default response
func (o *ListEKSVersionsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list e k s versions default response has a 2xx status code
func (o *ListEKSVersionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list e k s versions default response has a 3xx status code
func (o *ListEKSVersionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list e k s versions default response has a 4xx status code
func (o *ListEKSVersionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list e k s versions default response has a 5xx status code
func (o *ListEKSVersionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list e k s versions default response a status code equal to that given
func (o *ListEKSVersionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListEKSVersionsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/versions][%d] listEKSVersions default  %+v", o._statusCode, o.Payload)
}

func (o *ListEKSVersionsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/versions][%d] listEKSVersions default  %+v", o._statusCode, o.Payload)
}

func (o *ListEKSVersionsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListEKSVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
