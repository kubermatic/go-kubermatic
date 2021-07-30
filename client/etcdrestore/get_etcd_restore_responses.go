// Code generated by go-swagger; DO NOT EDIT.

package etcdrestore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// GetEtcdRestoreReader is a Reader for the GetEtcdRestore structure.
type GetEtcdRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEtcdRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEtcdRestoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEtcdRestoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEtcdRestoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetEtcdRestoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEtcdRestoreOK creates a GetEtcdRestoreOK with default headers values
func NewGetEtcdRestoreOK() *GetEtcdRestoreOK {
	return &GetEtcdRestoreOK{}
}

/*GetEtcdRestoreOK handles this case with default header values.

EtcdRestore
*/
type GetEtcdRestoreOK struct {
	Payload *models.EtcdRestore
}

func (o *GetEtcdRestoreOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] getEtcdRestoreOK  %+v", 200, o.Payload)
}

func (o *GetEtcdRestoreOK) GetPayload() *models.EtcdRestore {
	return o.Payload
}

func (o *GetEtcdRestoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EtcdRestore)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEtcdRestoreUnauthorized creates a GetEtcdRestoreUnauthorized with default headers values
func NewGetEtcdRestoreUnauthorized() *GetEtcdRestoreUnauthorized {
	return &GetEtcdRestoreUnauthorized{}
}

/*GetEtcdRestoreUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type GetEtcdRestoreUnauthorized struct {
}

func (o *GetEtcdRestoreUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] getEtcdRestoreUnauthorized ", 401)
}

func (o *GetEtcdRestoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEtcdRestoreForbidden creates a GetEtcdRestoreForbidden with default headers values
func NewGetEtcdRestoreForbidden() *GetEtcdRestoreForbidden {
	return &GetEtcdRestoreForbidden{}
}

/*GetEtcdRestoreForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type GetEtcdRestoreForbidden struct {
}

func (o *GetEtcdRestoreForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] getEtcdRestoreForbidden ", 403)
}

func (o *GetEtcdRestoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEtcdRestoreDefault creates a GetEtcdRestoreDefault with default headers values
func NewGetEtcdRestoreDefault(code int) *GetEtcdRestoreDefault {
	return &GetEtcdRestoreDefault{
		_statusCode: code,
	}
}

/*GetEtcdRestoreDefault handles this case with default header values.

errorResponse
*/
type GetEtcdRestoreDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get etcd restore default response
func (o *GetEtcdRestoreDefault) Code() int {
	return o._statusCode
}

func (o *GetEtcdRestoreDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] getEtcdRestore default  %+v", o._statusCode, o.Payload)
}

func (o *GetEtcdRestoreDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetEtcdRestoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
