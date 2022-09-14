// Code generated by go-swagger; DO NOT EDIT.

package resource_quota

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// GetResourceQuotaReader is a Reader for the GetResourceQuota structure.
type GetResourceQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceQuotaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetResourceQuotaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetResourceQuotaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetResourceQuotaDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetResourceQuotaOK creates a GetResourceQuotaOK with default headers values
func NewGetResourceQuotaOK() *GetResourceQuotaOK {
	return &GetResourceQuotaOK{}
}

/* GetResourceQuotaOK describes a response with status code 200, with default header values.

ResourceQuota
*/
type GetResourceQuotaOK struct {
	Payload *models.ResourceQuota
}

// IsSuccess returns true when this get resource quota o k response has a 2xx status code
func (o *GetResourceQuotaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get resource quota o k response has a 3xx status code
func (o *GetResourceQuotaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource quota o k response has a 4xx status code
func (o *GetResourceQuotaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resource quota o k response has a 5xx status code
func (o *GetResourceQuotaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource quota o k response a status code equal to that given
func (o *GetResourceQuotaOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetResourceQuotaOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/quotas/{quota_name}][%d] getResourceQuotaOK  %+v", 200, o.Payload)
}

func (o *GetResourceQuotaOK) String() string {
	return fmt.Sprintf("[GET /api/v2/quotas/{quota_name}][%d] getResourceQuotaOK  %+v", 200, o.Payload)
}

func (o *GetResourceQuotaOK) GetPayload() *models.ResourceQuota {
	return o.Payload
}

func (o *GetResourceQuotaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceQuota)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceQuotaUnauthorized creates a GetResourceQuotaUnauthorized with default headers values
func NewGetResourceQuotaUnauthorized() *GetResourceQuotaUnauthorized {
	return &GetResourceQuotaUnauthorized{}
}

/* GetResourceQuotaUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetResourceQuotaUnauthorized struct {
}

// IsSuccess returns true when this get resource quota unauthorized response has a 2xx status code
func (o *GetResourceQuotaUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource quota unauthorized response has a 3xx status code
func (o *GetResourceQuotaUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource quota unauthorized response has a 4xx status code
func (o *GetResourceQuotaUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resource quota unauthorized response has a 5xx status code
func (o *GetResourceQuotaUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource quota unauthorized response a status code equal to that given
func (o *GetResourceQuotaUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetResourceQuotaUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/quotas/{quota_name}][%d] getResourceQuotaUnauthorized ", 401)
}

func (o *GetResourceQuotaUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/quotas/{quota_name}][%d] getResourceQuotaUnauthorized ", 401)
}

func (o *GetResourceQuotaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetResourceQuotaForbidden creates a GetResourceQuotaForbidden with default headers values
func NewGetResourceQuotaForbidden() *GetResourceQuotaForbidden {
	return &GetResourceQuotaForbidden{}
}

/* GetResourceQuotaForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetResourceQuotaForbidden struct {
}

// IsSuccess returns true when this get resource quota forbidden response has a 2xx status code
func (o *GetResourceQuotaForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource quota forbidden response has a 3xx status code
func (o *GetResourceQuotaForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource quota forbidden response has a 4xx status code
func (o *GetResourceQuotaForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resource quota forbidden response has a 5xx status code
func (o *GetResourceQuotaForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource quota forbidden response a status code equal to that given
func (o *GetResourceQuotaForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetResourceQuotaForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/quotas/{quota_name}][%d] getResourceQuotaForbidden ", 403)
}

func (o *GetResourceQuotaForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/quotas/{quota_name}][%d] getResourceQuotaForbidden ", 403)
}

func (o *GetResourceQuotaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetResourceQuotaDefault creates a GetResourceQuotaDefault with default headers values
func NewGetResourceQuotaDefault(code int) *GetResourceQuotaDefault {
	return &GetResourceQuotaDefault{
		_statusCode: code,
	}
}

/* GetResourceQuotaDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetResourceQuotaDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get resource quota default response
func (o *GetResourceQuotaDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get resource quota default response has a 2xx status code
func (o *GetResourceQuotaDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get resource quota default response has a 3xx status code
func (o *GetResourceQuotaDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get resource quota default response has a 4xx status code
func (o *GetResourceQuotaDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get resource quota default response has a 5xx status code
func (o *GetResourceQuotaDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get resource quota default response a status code equal to that given
func (o *GetResourceQuotaDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetResourceQuotaDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/quotas/{quota_name}][%d] getResourceQuota default  %+v", o._statusCode, o.Payload)
}

func (o *GetResourceQuotaDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/quotas/{quota_name}][%d] getResourceQuota default  %+v", o._statusCode, o.Payload)
}

func (o *GetResourceQuotaDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetResourceQuotaDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
