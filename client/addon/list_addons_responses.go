// Code generated by go-swagger; DO NOT EDIT.

package addon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListAddonsReader is a Reader for the ListAddons structure.
type ListAddonsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAddonsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAddonsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAddonsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAddonsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListAddonsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAddonsOK creates a ListAddonsOK with default headers values
func NewListAddonsOK() *ListAddonsOK {
	return &ListAddonsOK{}
}

/*ListAddonsOK handles this case with default header values.

Addon
*/
type ListAddonsOK struct {
	Payload []*models.Addon
}

func (o *ListAddonsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons][%d] listAddonsOK  %+v", 200, o.Payload)
}

func (o *ListAddonsOK) GetPayload() []*models.Addon {
	return o.Payload
}

func (o *ListAddonsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAddonsUnauthorized creates a ListAddonsUnauthorized with default headers values
func NewListAddonsUnauthorized() *ListAddonsUnauthorized {
	return &ListAddonsUnauthorized{}
}

/*ListAddonsUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type ListAddonsUnauthorized struct {
}

func (o *ListAddonsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons][%d] listAddonsUnauthorized ", 401)
}

func (o *ListAddonsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAddonsForbidden creates a ListAddonsForbidden with default headers values
func NewListAddonsForbidden() *ListAddonsForbidden {
	return &ListAddonsForbidden{}
}

/*ListAddonsForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type ListAddonsForbidden struct {
}

func (o *ListAddonsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons][%d] listAddonsForbidden ", 403)
}

func (o *ListAddonsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAddonsDefault creates a ListAddonsDefault with default headers values
func NewListAddonsDefault(code int) *ListAddonsDefault {
	return &ListAddonsDefault{
		_statusCode: code,
	}
}

/*ListAddonsDefault handles this case with default header values.

errorResponse
*/
type ListAddonsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list addons default response
func (o *ListAddonsDefault) Code() int {
	return o._statusCode
}

func (o *ListAddonsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons][%d] listAddons default  %+v", o._statusCode, o.Payload)
}

func (o *ListAddonsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAddonsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
