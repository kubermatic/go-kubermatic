// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// GetExternalClusterUpgradesReader is a Reader for the GetExternalClusterUpgrades structure.
type GetExternalClusterUpgradesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalClusterUpgradesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalClusterUpgradesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetExternalClusterUpgradesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalClusterUpgradesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetExternalClusterUpgradesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetExternalClusterUpgradesOK creates a GetExternalClusterUpgradesOK with default headers values
func NewGetExternalClusterUpgradesOK() *GetExternalClusterUpgradesOK {
	return &GetExternalClusterUpgradesOK{}
}

/* GetExternalClusterUpgradesOK describes a response with status code 200, with default header values.

MasterVersion
*/
type GetExternalClusterUpgradesOK struct {
	Payload []*models.MasterVersion
}

// IsSuccess returns true when this get external cluster upgrades o k response has a 2xx status code
func (o *GetExternalClusterUpgradesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get external cluster upgrades o k response has a 3xx status code
func (o *GetExternalClusterUpgradesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get external cluster upgrades o k response has a 4xx status code
func (o *GetExternalClusterUpgradesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get external cluster upgrades o k response has a 5xx status code
func (o *GetExternalClusterUpgradesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get external cluster upgrades o k response a status code equal to that given
func (o *GetExternalClusterUpgradesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetExternalClusterUpgradesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/upgrades][%d] getExternalClusterUpgradesOK  %+v", 200, o.Payload)
}

func (o *GetExternalClusterUpgradesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/upgrades][%d] getExternalClusterUpgradesOK  %+v", 200, o.Payload)
}

func (o *GetExternalClusterUpgradesOK) GetPayload() []*models.MasterVersion {
	return o.Payload
}

func (o *GetExternalClusterUpgradesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalClusterUpgradesUnauthorized creates a GetExternalClusterUpgradesUnauthorized with default headers values
func NewGetExternalClusterUpgradesUnauthorized() *GetExternalClusterUpgradesUnauthorized {
	return &GetExternalClusterUpgradesUnauthorized{}
}

/* GetExternalClusterUpgradesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetExternalClusterUpgradesUnauthorized struct {
}

// IsSuccess returns true when this get external cluster upgrades unauthorized response has a 2xx status code
func (o *GetExternalClusterUpgradesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get external cluster upgrades unauthorized response has a 3xx status code
func (o *GetExternalClusterUpgradesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get external cluster upgrades unauthorized response has a 4xx status code
func (o *GetExternalClusterUpgradesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get external cluster upgrades unauthorized response has a 5xx status code
func (o *GetExternalClusterUpgradesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get external cluster upgrades unauthorized response a status code equal to that given
func (o *GetExternalClusterUpgradesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetExternalClusterUpgradesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/upgrades][%d] getExternalClusterUpgradesUnauthorized ", 401)
}

func (o *GetExternalClusterUpgradesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/upgrades][%d] getExternalClusterUpgradesUnauthorized ", 401)
}

func (o *GetExternalClusterUpgradesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExternalClusterUpgradesForbidden creates a GetExternalClusterUpgradesForbidden with default headers values
func NewGetExternalClusterUpgradesForbidden() *GetExternalClusterUpgradesForbidden {
	return &GetExternalClusterUpgradesForbidden{}
}

/* GetExternalClusterUpgradesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetExternalClusterUpgradesForbidden struct {
}

// IsSuccess returns true when this get external cluster upgrades forbidden response has a 2xx status code
func (o *GetExternalClusterUpgradesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get external cluster upgrades forbidden response has a 3xx status code
func (o *GetExternalClusterUpgradesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get external cluster upgrades forbidden response has a 4xx status code
func (o *GetExternalClusterUpgradesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get external cluster upgrades forbidden response has a 5xx status code
func (o *GetExternalClusterUpgradesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get external cluster upgrades forbidden response a status code equal to that given
func (o *GetExternalClusterUpgradesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetExternalClusterUpgradesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/upgrades][%d] getExternalClusterUpgradesForbidden ", 403)
}

func (o *GetExternalClusterUpgradesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/upgrades][%d] getExternalClusterUpgradesForbidden ", 403)
}

func (o *GetExternalClusterUpgradesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExternalClusterUpgradesDefault creates a GetExternalClusterUpgradesDefault with default headers values
func NewGetExternalClusterUpgradesDefault(code int) *GetExternalClusterUpgradesDefault {
	return &GetExternalClusterUpgradesDefault{
		_statusCode: code,
	}
}

/* GetExternalClusterUpgradesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetExternalClusterUpgradesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get external cluster upgrades default response
func (o *GetExternalClusterUpgradesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get external cluster upgrades default response has a 2xx status code
func (o *GetExternalClusterUpgradesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get external cluster upgrades default response has a 3xx status code
func (o *GetExternalClusterUpgradesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get external cluster upgrades default response has a 4xx status code
func (o *GetExternalClusterUpgradesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get external cluster upgrades default response has a 5xx status code
func (o *GetExternalClusterUpgradesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get external cluster upgrades default response a status code equal to that given
func (o *GetExternalClusterUpgradesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetExternalClusterUpgradesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/upgrades][%d] getExternalClusterUpgrades default  %+v", o._statusCode, o.Payload)
}

func (o *GetExternalClusterUpgradesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/upgrades][%d] getExternalClusterUpgrades default  %+v", o._statusCode, o.Payload)
}

func (o *GetExternalClusterUpgradesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetExternalClusterUpgradesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
