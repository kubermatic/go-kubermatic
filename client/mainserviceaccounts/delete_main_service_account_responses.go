// Code generated by go-swagger; DO NOT EDIT.

package mainserviceaccounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// DeleteMainServiceAccountReader is a Reader for the DeleteMainServiceAccount structure.
type DeleteMainServiceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMainServiceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMainServiceAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteMainServiceAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteMainServiceAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteMainServiceAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMainServiceAccountOK creates a DeleteMainServiceAccountOK with default headers values
func NewDeleteMainServiceAccountOK() *DeleteMainServiceAccountOK {
	return &DeleteMainServiceAccountOK{}
}

/*DeleteMainServiceAccountOK handles this case with default header values.

EmptyResponse is a empty response
*/
type DeleteMainServiceAccountOK struct {
}

func (o *DeleteMainServiceAccountOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/serviceaccounts/{serviceaccount_id}][%d] deleteMainServiceAccountOK ", 200)
}

func (o *DeleteMainServiceAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMainServiceAccountUnauthorized creates a DeleteMainServiceAccountUnauthorized with default headers values
func NewDeleteMainServiceAccountUnauthorized() *DeleteMainServiceAccountUnauthorized {
	return &DeleteMainServiceAccountUnauthorized{}
}

/*DeleteMainServiceAccountUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type DeleteMainServiceAccountUnauthorized struct {
}

func (o *DeleteMainServiceAccountUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/serviceaccounts/{serviceaccount_id}][%d] deleteMainServiceAccountUnauthorized ", 401)
}

func (o *DeleteMainServiceAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMainServiceAccountForbidden creates a DeleteMainServiceAccountForbidden with default headers values
func NewDeleteMainServiceAccountForbidden() *DeleteMainServiceAccountForbidden {
	return &DeleteMainServiceAccountForbidden{}
}

/*DeleteMainServiceAccountForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type DeleteMainServiceAccountForbidden struct {
}

func (o *DeleteMainServiceAccountForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/serviceaccounts/{serviceaccount_id}][%d] deleteMainServiceAccountForbidden ", 403)
}

func (o *DeleteMainServiceAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMainServiceAccountDefault creates a DeleteMainServiceAccountDefault with default headers values
func NewDeleteMainServiceAccountDefault(code int) *DeleteMainServiceAccountDefault {
	return &DeleteMainServiceAccountDefault{
		_statusCode: code,
	}
}

/*DeleteMainServiceAccountDefault handles this case with default header values.

errorResponse
*/
type DeleteMainServiceAccountDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete main service account default response
func (o *DeleteMainServiceAccountDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMainServiceAccountDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/serviceaccounts/{serviceaccount_id}][%d] deleteMainServiceAccount default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMainServiceAccountDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteMainServiceAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
