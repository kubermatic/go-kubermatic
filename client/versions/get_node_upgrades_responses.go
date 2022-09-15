// Code generated by go-swagger; DO NOT EDIT.

package versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// GetNodeUpgradesReader is a Reader for the GetNodeUpgrades structure.
type GetNodeUpgradesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeUpgradesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNodeUpgradesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetNodeUpgradesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNodeUpgradesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetNodeUpgradesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNodeUpgradesOK creates a GetNodeUpgradesOK with default headers values
func NewGetNodeUpgradesOK() *GetNodeUpgradesOK {
	return &GetNodeUpgradesOK{}
}

/*
GetNodeUpgradesOK describes a response with status code 200, with default header values.

MasterVersion
*/
type GetNodeUpgradesOK struct {
	Payload []*models.MasterVersion
}

// IsSuccess returns true when this get node upgrades o k response has a 2xx status code
func (o *GetNodeUpgradesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get node upgrades o k response has a 3xx status code
func (o *GetNodeUpgradesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get node upgrades o k response has a 4xx status code
func (o *GetNodeUpgradesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get node upgrades o k response has a 5xx status code
func (o *GetNodeUpgradesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get node upgrades o k response a status code equal to that given
func (o *GetNodeUpgradesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNodeUpgradesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/upgrades/node][%d] getNodeUpgradesOK  %+v", 200, o.Payload)
}

func (o *GetNodeUpgradesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/upgrades/node][%d] getNodeUpgradesOK  %+v", 200, o.Payload)
}

func (o *GetNodeUpgradesOK) GetPayload() []*models.MasterVersion {
	return o.Payload
}

func (o *GetNodeUpgradesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeUpgradesUnauthorized creates a GetNodeUpgradesUnauthorized with default headers values
func NewGetNodeUpgradesUnauthorized() *GetNodeUpgradesUnauthorized {
	return &GetNodeUpgradesUnauthorized{}
}

/*
GetNodeUpgradesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetNodeUpgradesUnauthorized struct {
}

// IsSuccess returns true when this get node upgrades unauthorized response has a 2xx status code
func (o *GetNodeUpgradesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get node upgrades unauthorized response has a 3xx status code
func (o *GetNodeUpgradesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get node upgrades unauthorized response has a 4xx status code
func (o *GetNodeUpgradesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get node upgrades unauthorized response has a 5xx status code
func (o *GetNodeUpgradesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get node upgrades unauthorized response a status code equal to that given
func (o *GetNodeUpgradesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetNodeUpgradesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/upgrades/node][%d] getNodeUpgradesUnauthorized ", 401)
}

func (o *GetNodeUpgradesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/upgrades/node][%d] getNodeUpgradesUnauthorized ", 401)
}

func (o *GetNodeUpgradesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNodeUpgradesForbidden creates a GetNodeUpgradesForbidden with default headers values
func NewGetNodeUpgradesForbidden() *GetNodeUpgradesForbidden {
	return &GetNodeUpgradesForbidden{}
}

/*
GetNodeUpgradesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetNodeUpgradesForbidden struct {
}

// IsSuccess returns true when this get node upgrades forbidden response has a 2xx status code
func (o *GetNodeUpgradesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get node upgrades forbidden response has a 3xx status code
func (o *GetNodeUpgradesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get node upgrades forbidden response has a 4xx status code
func (o *GetNodeUpgradesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get node upgrades forbidden response has a 5xx status code
func (o *GetNodeUpgradesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get node upgrades forbidden response a status code equal to that given
func (o *GetNodeUpgradesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetNodeUpgradesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/upgrades/node][%d] getNodeUpgradesForbidden ", 403)
}

func (o *GetNodeUpgradesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/upgrades/node][%d] getNodeUpgradesForbidden ", 403)
}

func (o *GetNodeUpgradesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNodeUpgradesDefault creates a GetNodeUpgradesDefault with default headers values
func NewGetNodeUpgradesDefault(code int) *GetNodeUpgradesDefault {
	return &GetNodeUpgradesDefault{
		_statusCode: code,
	}
}

/*
GetNodeUpgradesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetNodeUpgradesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get node upgrades default response
func (o *GetNodeUpgradesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get node upgrades default response has a 2xx status code
func (o *GetNodeUpgradesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get node upgrades default response has a 3xx status code
func (o *GetNodeUpgradesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get node upgrades default response has a 4xx status code
func (o *GetNodeUpgradesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get node upgrades default response has a 5xx status code
func (o *GetNodeUpgradesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get node upgrades default response a status code equal to that given
func (o *GetNodeUpgradesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetNodeUpgradesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/upgrades/node][%d] getNodeUpgrades default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodeUpgradesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/upgrades/node][%d] getNodeUpgrades default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodeUpgradesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetNodeUpgradesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
