// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// DeleteBackupDestinationReader is a Reader for the DeleteBackupDestination structure.
type DeleteBackupDestinationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBackupDestinationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteBackupDestinationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteBackupDestinationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteBackupDestinationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteBackupDestinationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteBackupDestinationOK creates a DeleteBackupDestinationOK with default headers values
func NewDeleteBackupDestinationOK() *DeleteBackupDestinationOK {
	return &DeleteBackupDestinationOK{}
}

/* DeleteBackupDestinationOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteBackupDestinationOK struct {
}

// IsSuccess returns true when this delete backup destination o k response has a 2xx status code
func (o *DeleteBackupDestinationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete backup destination o k response has a 3xx status code
func (o *DeleteBackupDestinationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete backup destination o k response has a 4xx status code
func (o *DeleteBackupDestinationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete backup destination o k response has a 5xx status code
func (o *DeleteBackupDestinationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete backup destination o k response a status code equal to that given
func (o *DeleteBackupDestinationOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteBackupDestinationOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}/backupdestinations/{backup_destination}][%d] deleteBackupDestinationOK ", 200)
}

func (o *DeleteBackupDestinationOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}/backupdestinations/{backup_destination}][%d] deleteBackupDestinationOK ", 200)
}

func (o *DeleteBackupDestinationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBackupDestinationUnauthorized creates a DeleteBackupDestinationUnauthorized with default headers values
func NewDeleteBackupDestinationUnauthorized() *DeleteBackupDestinationUnauthorized {
	return &DeleteBackupDestinationUnauthorized{}
}

/* DeleteBackupDestinationUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteBackupDestinationUnauthorized struct {
}

// IsSuccess returns true when this delete backup destination unauthorized response has a 2xx status code
func (o *DeleteBackupDestinationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete backup destination unauthorized response has a 3xx status code
func (o *DeleteBackupDestinationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete backup destination unauthorized response has a 4xx status code
func (o *DeleteBackupDestinationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete backup destination unauthorized response has a 5xx status code
func (o *DeleteBackupDestinationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete backup destination unauthorized response a status code equal to that given
func (o *DeleteBackupDestinationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteBackupDestinationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}/backupdestinations/{backup_destination}][%d] deleteBackupDestinationUnauthorized ", 401)
}

func (o *DeleteBackupDestinationUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}/backupdestinations/{backup_destination}][%d] deleteBackupDestinationUnauthorized ", 401)
}

func (o *DeleteBackupDestinationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBackupDestinationForbidden creates a DeleteBackupDestinationForbidden with default headers values
func NewDeleteBackupDestinationForbidden() *DeleteBackupDestinationForbidden {
	return &DeleteBackupDestinationForbidden{}
}

/* DeleteBackupDestinationForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteBackupDestinationForbidden struct {
}

// IsSuccess returns true when this delete backup destination forbidden response has a 2xx status code
func (o *DeleteBackupDestinationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete backup destination forbidden response has a 3xx status code
func (o *DeleteBackupDestinationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete backup destination forbidden response has a 4xx status code
func (o *DeleteBackupDestinationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete backup destination forbidden response has a 5xx status code
func (o *DeleteBackupDestinationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete backup destination forbidden response a status code equal to that given
func (o *DeleteBackupDestinationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteBackupDestinationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}/backupdestinations/{backup_destination}][%d] deleteBackupDestinationForbidden ", 403)
}

func (o *DeleteBackupDestinationForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}/backupdestinations/{backup_destination}][%d] deleteBackupDestinationForbidden ", 403)
}

func (o *DeleteBackupDestinationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBackupDestinationDefault creates a DeleteBackupDestinationDefault with default headers values
func NewDeleteBackupDestinationDefault(code int) *DeleteBackupDestinationDefault {
	return &DeleteBackupDestinationDefault{
		_statusCode: code,
	}
}

/* DeleteBackupDestinationDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteBackupDestinationDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete backup destination default response
func (o *DeleteBackupDestinationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete backup destination default response has a 2xx status code
func (o *DeleteBackupDestinationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete backup destination default response has a 3xx status code
func (o *DeleteBackupDestinationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete backup destination default response has a 4xx status code
func (o *DeleteBackupDestinationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete backup destination default response has a 5xx status code
func (o *DeleteBackupDestinationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete backup destination default response a status code equal to that given
func (o *DeleteBackupDestinationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteBackupDestinationDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}/backupdestinations/{backup_destination}][%d] deleteBackupDestination default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteBackupDestinationDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}/backupdestinations/{backup_destination}][%d] deleteBackupDestination default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteBackupDestinationDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteBackupDestinationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
