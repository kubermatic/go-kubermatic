// Code generated by go-swagger; DO NOT EDIT.

package addon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// PatchAddonReader is a Reader for the PatchAddon structure.
type PatchAddonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAddonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAddonOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchAddonUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchAddonForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchAddonDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchAddonOK creates a PatchAddonOK with default headers values
func NewPatchAddonOK() *PatchAddonOK {
	return &PatchAddonOK{}
}

/* PatchAddonOK describes a response with status code 200, with default header values.

Addon
*/
type PatchAddonOK struct {
	Payload *models.Addon
}

func (o *PatchAddonOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] patchAddonOK  %+v", 200, o.Payload)
}
func (o *PatchAddonOK) GetPayload() *models.Addon {
	return o.Payload
}

func (o *PatchAddonOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Addon)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAddonUnauthorized creates a PatchAddonUnauthorized with default headers values
func NewPatchAddonUnauthorized() *PatchAddonUnauthorized {
	return &PatchAddonUnauthorized{}
}

/* PatchAddonUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchAddonUnauthorized struct {
}

func (o *PatchAddonUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] patchAddonUnauthorized ", 401)
}

func (o *PatchAddonUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchAddonForbidden creates a PatchAddonForbidden with default headers values
func NewPatchAddonForbidden() *PatchAddonForbidden {
	return &PatchAddonForbidden{}
}

/* PatchAddonForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchAddonForbidden struct {
}

func (o *PatchAddonForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] patchAddonForbidden ", 403)
}

func (o *PatchAddonForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchAddonDefault creates a PatchAddonDefault with default headers values
func NewPatchAddonDefault(code int) *PatchAddonDefault {
	return &PatchAddonDefault{
		_statusCode: code,
	}
}

/* PatchAddonDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PatchAddonDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch addon default response
func (o *PatchAddonDefault) Code() int {
	return o._statusCode
}

func (o *PatchAddonDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] patchAddon default  %+v", o._statusCode, o.Payload)
}
func (o *PatchAddonDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchAddonDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
