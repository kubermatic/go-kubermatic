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

// PatchServiceAccountTokenReader is a Reader for the PatchServiceAccountToken structure.
type PatchServiceAccountTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchServiceAccountTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchServiceAccountTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchServiceAccountTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchServiceAccountTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchServiceAccountTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchServiceAccountTokenOK creates a PatchServiceAccountTokenOK with default headers values
func NewPatchServiceAccountTokenOK() *PatchServiceAccountTokenOK {
	return &PatchServiceAccountTokenOK{}
}

/* PatchServiceAccountTokenOK describes a response with status code 200, with default header values.

PublicServiceAccountToken
*/
type PatchServiceAccountTokenOK struct {
	Payload *models.PublicServiceAccountToken
}

// IsSuccess returns true when this patch service account token o k response has a 2xx status code
func (o *PatchServiceAccountTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch service account token o k response has a 3xx status code
func (o *PatchServiceAccountTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch service account token o k response has a 4xx status code
func (o *PatchServiceAccountTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch service account token o k response has a 5xx status code
func (o *PatchServiceAccountTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch service account token o k response a status code equal to that given
func (o *PatchServiceAccountTokenOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchServiceAccountTokenOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens/{token_id}][%d] patchServiceAccountTokenOK  %+v", 200, o.Payload)
}

func (o *PatchServiceAccountTokenOK) String() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens/{token_id}][%d] patchServiceAccountTokenOK  %+v", 200, o.Payload)
}

func (o *PatchServiceAccountTokenOK) GetPayload() *models.PublicServiceAccountToken {
	return o.Payload
}

func (o *PatchServiceAccountTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicServiceAccountToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchServiceAccountTokenUnauthorized creates a PatchServiceAccountTokenUnauthorized with default headers values
func NewPatchServiceAccountTokenUnauthorized() *PatchServiceAccountTokenUnauthorized {
	return &PatchServiceAccountTokenUnauthorized{}
}

/* PatchServiceAccountTokenUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchServiceAccountTokenUnauthorized struct {
}

// IsSuccess returns true when this patch service account token unauthorized response has a 2xx status code
func (o *PatchServiceAccountTokenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch service account token unauthorized response has a 3xx status code
func (o *PatchServiceAccountTokenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch service account token unauthorized response has a 4xx status code
func (o *PatchServiceAccountTokenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch service account token unauthorized response has a 5xx status code
func (o *PatchServiceAccountTokenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch service account token unauthorized response a status code equal to that given
func (o *PatchServiceAccountTokenUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchServiceAccountTokenUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens/{token_id}][%d] patchServiceAccountTokenUnauthorized ", 401)
}

func (o *PatchServiceAccountTokenUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens/{token_id}][%d] patchServiceAccountTokenUnauthorized ", 401)
}

func (o *PatchServiceAccountTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchServiceAccountTokenForbidden creates a PatchServiceAccountTokenForbidden with default headers values
func NewPatchServiceAccountTokenForbidden() *PatchServiceAccountTokenForbidden {
	return &PatchServiceAccountTokenForbidden{}
}

/* PatchServiceAccountTokenForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchServiceAccountTokenForbidden struct {
}

// IsSuccess returns true when this patch service account token forbidden response has a 2xx status code
func (o *PatchServiceAccountTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch service account token forbidden response has a 3xx status code
func (o *PatchServiceAccountTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch service account token forbidden response has a 4xx status code
func (o *PatchServiceAccountTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch service account token forbidden response has a 5xx status code
func (o *PatchServiceAccountTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch service account token forbidden response a status code equal to that given
func (o *PatchServiceAccountTokenForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchServiceAccountTokenForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens/{token_id}][%d] patchServiceAccountTokenForbidden ", 403)
}

func (o *PatchServiceAccountTokenForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens/{token_id}][%d] patchServiceAccountTokenForbidden ", 403)
}

func (o *PatchServiceAccountTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchServiceAccountTokenDefault creates a PatchServiceAccountTokenDefault with default headers values
func NewPatchServiceAccountTokenDefault(code int) *PatchServiceAccountTokenDefault {
	return &PatchServiceAccountTokenDefault{
		_statusCode: code,
	}
}

/* PatchServiceAccountTokenDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PatchServiceAccountTokenDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch service account token default response
func (o *PatchServiceAccountTokenDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this patch service account token default response has a 2xx status code
func (o *PatchServiceAccountTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch service account token default response has a 3xx status code
func (o *PatchServiceAccountTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch service account token default response has a 4xx status code
func (o *PatchServiceAccountTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch service account token default response has a 5xx status code
func (o *PatchServiceAccountTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch service account token default response a status code equal to that given
func (o *PatchServiceAccountTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PatchServiceAccountTokenDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens/{token_id}][%d] patchServiceAccountToken default  %+v", o._statusCode, o.Payload)
}

func (o *PatchServiceAccountTokenDefault) String() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens/{token_id}][%d] patchServiceAccountToken default  %+v", o._statusCode, o.Payload)
}

func (o *PatchServiceAccountTokenDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchServiceAccountTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
