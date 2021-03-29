// Code generated by go-swagger; DO NOT EDIT.

package tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// PatchMainServiceAccountTokenReader is a Reader for the PatchMainServiceAccountToken structure.
type PatchMainServiceAccountTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchMainServiceAccountTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchMainServiceAccountTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchMainServiceAccountTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchMainServiceAccountTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchMainServiceAccountTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchMainServiceAccountTokenOK creates a PatchMainServiceAccountTokenOK with default headers values
func NewPatchMainServiceAccountTokenOK() *PatchMainServiceAccountTokenOK {
	return &PatchMainServiceAccountTokenOK{}
}

/*PatchMainServiceAccountTokenOK handles this case with default header values.

PublicServiceAccountToken
*/
type PatchMainServiceAccountTokenOK struct {
	Payload *models.PublicServiceAccountToken
}

func (o *PatchMainServiceAccountTokenOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/serviceaccounts/{serviceaccount_id}/tokens/{token_id}][%d] patchMainServiceAccountTokenOK  %+v", 200, o.Payload)
}

func (o *PatchMainServiceAccountTokenOK) GetPayload() *models.PublicServiceAccountToken {
	return o.Payload
}

func (o *PatchMainServiceAccountTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicServiceAccountToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMainServiceAccountTokenUnauthorized creates a PatchMainServiceAccountTokenUnauthorized with default headers values
func NewPatchMainServiceAccountTokenUnauthorized() *PatchMainServiceAccountTokenUnauthorized {
	return &PatchMainServiceAccountTokenUnauthorized{}
}

/*PatchMainServiceAccountTokenUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type PatchMainServiceAccountTokenUnauthorized struct {
}

func (o *PatchMainServiceAccountTokenUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/serviceaccounts/{serviceaccount_id}/tokens/{token_id}][%d] patchMainServiceAccountTokenUnauthorized ", 401)
}

func (o *PatchMainServiceAccountTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchMainServiceAccountTokenForbidden creates a PatchMainServiceAccountTokenForbidden with default headers values
func NewPatchMainServiceAccountTokenForbidden() *PatchMainServiceAccountTokenForbidden {
	return &PatchMainServiceAccountTokenForbidden{}
}

/*PatchMainServiceAccountTokenForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type PatchMainServiceAccountTokenForbidden struct {
}

func (o *PatchMainServiceAccountTokenForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/serviceaccounts/{serviceaccount_id}/tokens/{token_id}][%d] patchMainServiceAccountTokenForbidden ", 403)
}

func (o *PatchMainServiceAccountTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchMainServiceAccountTokenDefault creates a PatchMainServiceAccountTokenDefault with default headers values
func NewPatchMainServiceAccountTokenDefault(code int) *PatchMainServiceAccountTokenDefault {
	return &PatchMainServiceAccountTokenDefault{
		_statusCode: code,
	}
}

/*PatchMainServiceAccountTokenDefault handles this case with default header values.

errorResponse
*/
type PatchMainServiceAccountTokenDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch main service account token default response
func (o *PatchMainServiceAccountTokenDefault) Code() int {
	return o._statusCode
}

func (o *PatchMainServiceAccountTokenDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/serviceaccounts/{serviceaccount_id}/tokens/{token_id}][%d] patchMainServiceAccountToken default  %+v", o._statusCode, o.Payload)
}

func (o *PatchMainServiceAccountTokenDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchMainServiceAccountTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
