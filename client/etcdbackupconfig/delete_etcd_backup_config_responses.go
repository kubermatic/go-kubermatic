// Code generated by go-swagger; DO NOT EDIT.

package etcdbackupconfig

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// DeleteEtcdBackupConfigReader is a Reader for the DeleteEtcdBackupConfig structure.
type DeleteEtcdBackupConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEtcdBackupConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEtcdBackupConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteEtcdBackupConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteEtcdBackupConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteEtcdBackupConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteEtcdBackupConfigOK creates a DeleteEtcdBackupConfigOK with default headers values
func NewDeleteEtcdBackupConfigOK() *DeleteEtcdBackupConfigOK {
	return &DeleteEtcdBackupConfigOK{}
}

/* DeleteEtcdBackupConfigOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteEtcdBackupConfigOK struct {
}

func (o *DeleteEtcdBackupConfigOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] deleteEtcdBackupConfigOK ", 200)
}

func (o *DeleteEtcdBackupConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEtcdBackupConfigUnauthorized creates a DeleteEtcdBackupConfigUnauthorized with default headers values
func NewDeleteEtcdBackupConfigUnauthorized() *DeleteEtcdBackupConfigUnauthorized {
	return &DeleteEtcdBackupConfigUnauthorized{}
}

/* DeleteEtcdBackupConfigUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteEtcdBackupConfigUnauthorized struct {
}

func (o *DeleteEtcdBackupConfigUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] deleteEtcdBackupConfigUnauthorized ", 401)
}

func (o *DeleteEtcdBackupConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEtcdBackupConfigForbidden creates a DeleteEtcdBackupConfigForbidden with default headers values
func NewDeleteEtcdBackupConfigForbidden() *DeleteEtcdBackupConfigForbidden {
	return &DeleteEtcdBackupConfigForbidden{}
}

/* DeleteEtcdBackupConfigForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteEtcdBackupConfigForbidden struct {
}

func (o *DeleteEtcdBackupConfigForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] deleteEtcdBackupConfigForbidden ", 403)
}

func (o *DeleteEtcdBackupConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEtcdBackupConfigDefault creates a DeleteEtcdBackupConfigDefault with default headers values
func NewDeleteEtcdBackupConfigDefault(code int) *DeleteEtcdBackupConfigDefault {
	return &DeleteEtcdBackupConfigDefault{
		_statusCode: code,
	}
}

/* DeleteEtcdBackupConfigDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteEtcdBackupConfigDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete etcd backup config default response
func (o *DeleteEtcdBackupConfigDefault) Code() int {
	return o._statusCode
}

func (o *DeleteEtcdBackupConfigDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] deleteEtcdBackupConfig default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteEtcdBackupConfigDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteEtcdBackupConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
