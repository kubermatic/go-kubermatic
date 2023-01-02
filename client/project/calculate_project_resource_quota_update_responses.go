// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/kubermatic/go-kubermatic/models"
)

// CalculateProjectResourceQuotaUpdateReader is a Reader for the CalculateProjectResourceQuotaUpdate structure.
type CalculateProjectResourceQuotaUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CalculateProjectResourceQuotaUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCalculateProjectResourceQuotaUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCalculateProjectResourceQuotaUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCalculateProjectResourceQuotaUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCalculateProjectResourceQuotaUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCalculateProjectResourceQuotaUpdateOK creates a CalculateProjectResourceQuotaUpdateOK with default headers values
func NewCalculateProjectResourceQuotaUpdateOK() *CalculateProjectResourceQuotaUpdateOK {
	return &CalculateProjectResourceQuotaUpdateOK{}
}

/*
CalculateProjectResourceQuotaUpdateOK describes a response with status code 200, with default header values.

ResourceQuotaUpdateCalculation
*/
type CalculateProjectResourceQuotaUpdateOK struct {
	Payload *models.ResourceQuotaUpdateCalculation
}

// IsSuccess returns true when this calculate project resource quota update o k response has a 2xx status code
func (o *CalculateProjectResourceQuotaUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this calculate project resource quota update o k response has a 3xx status code
func (o *CalculateProjectResourceQuotaUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this calculate project resource quota update o k response has a 4xx status code
func (o *CalculateProjectResourceQuotaUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this calculate project resource quota update o k response has a 5xx status code
func (o *CalculateProjectResourceQuotaUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this calculate project resource quota update o k response a status code equal to that given
func (o *CalculateProjectResourceQuotaUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *CalculateProjectResourceQuotaUpdateOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/quotacalculation][%d] calculateProjectResourceQuotaUpdateOK  %+v", 200, o.Payload)
}

func (o *CalculateProjectResourceQuotaUpdateOK) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/quotacalculation][%d] calculateProjectResourceQuotaUpdateOK  %+v", 200, o.Payload)
}

func (o *CalculateProjectResourceQuotaUpdateOK) GetPayload() *models.ResourceQuotaUpdateCalculation {
	return o.Payload
}

func (o *CalculateProjectResourceQuotaUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceQuotaUpdateCalculation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCalculateProjectResourceQuotaUpdateUnauthorized creates a CalculateProjectResourceQuotaUpdateUnauthorized with default headers values
func NewCalculateProjectResourceQuotaUpdateUnauthorized() *CalculateProjectResourceQuotaUpdateUnauthorized {
	return &CalculateProjectResourceQuotaUpdateUnauthorized{}
}

/*
CalculateProjectResourceQuotaUpdateUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CalculateProjectResourceQuotaUpdateUnauthorized struct {
}

// IsSuccess returns true when this calculate project resource quota update unauthorized response has a 2xx status code
func (o *CalculateProjectResourceQuotaUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this calculate project resource quota update unauthorized response has a 3xx status code
func (o *CalculateProjectResourceQuotaUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this calculate project resource quota update unauthorized response has a 4xx status code
func (o *CalculateProjectResourceQuotaUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this calculate project resource quota update unauthorized response has a 5xx status code
func (o *CalculateProjectResourceQuotaUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this calculate project resource quota update unauthorized response a status code equal to that given
func (o *CalculateProjectResourceQuotaUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CalculateProjectResourceQuotaUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/quotacalculation][%d] calculateProjectResourceQuotaUpdateUnauthorized ", 401)
}

func (o *CalculateProjectResourceQuotaUpdateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/quotacalculation][%d] calculateProjectResourceQuotaUpdateUnauthorized ", 401)
}

func (o *CalculateProjectResourceQuotaUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCalculateProjectResourceQuotaUpdateForbidden creates a CalculateProjectResourceQuotaUpdateForbidden with default headers values
func NewCalculateProjectResourceQuotaUpdateForbidden() *CalculateProjectResourceQuotaUpdateForbidden {
	return &CalculateProjectResourceQuotaUpdateForbidden{}
}

/*
CalculateProjectResourceQuotaUpdateForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CalculateProjectResourceQuotaUpdateForbidden struct {
}

// IsSuccess returns true when this calculate project resource quota update forbidden response has a 2xx status code
func (o *CalculateProjectResourceQuotaUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this calculate project resource quota update forbidden response has a 3xx status code
func (o *CalculateProjectResourceQuotaUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this calculate project resource quota update forbidden response has a 4xx status code
func (o *CalculateProjectResourceQuotaUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this calculate project resource quota update forbidden response has a 5xx status code
func (o *CalculateProjectResourceQuotaUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this calculate project resource quota update forbidden response a status code equal to that given
func (o *CalculateProjectResourceQuotaUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CalculateProjectResourceQuotaUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/quotacalculation][%d] calculateProjectResourceQuotaUpdateForbidden ", 403)
}

func (o *CalculateProjectResourceQuotaUpdateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/quotacalculation][%d] calculateProjectResourceQuotaUpdateForbidden ", 403)
}

func (o *CalculateProjectResourceQuotaUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCalculateProjectResourceQuotaUpdateDefault creates a CalculateProjectResourceQuotaUpdateDefault with default headers values
func NewCalculateProjectResourceQuotaUpdateDefault(code int) *CalculateProjectResourceQuotaUpdateDefault {
	return &CalculateProjectResourceQuotaUpdateDefault{
		_statusCode: code,
	}
}

/*
CalculateProjectResourceQuotaUpdateDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CalculateProjectResourceQuotaUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the calculate project resource quota update default response
func (o *CalculateProjectResourceQuotaUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this calculate project resource quota update default response has a 2xx status code
func (o *CalculateProjectResourceQuotaUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this calculate project resource quota update default response has a 3xx status code
func (o *CalculateProjectResourceQuotaUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this calculate project resource quota update default response has a 4xx status code
func (o *CalculateProjectResourceQuotaUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this calculate project resource quota update default response has a 5xx status code
func (o *CalculateProjectResourceQuotaUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this calculate project resource quota update default response a status code equal to that given
func (o *CalculateProjectResourceQuotaUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CalculateProjectResourceQuotaUpdateDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/quotacalculation][%d] calculateProjectResourceQuotaUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *CalculateProjectResourceQuotaUpdateDefault) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/quotacalculation][%d] calculateProjectResourceQuotaUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *CalculateProjectResourceQuotaUpdateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CalculateProjectResourceQuotaUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CalculateProjectResourceQuotaUpdateBody calculate project resource quota update body
swagger:model CalculateProjectResourceQuotaUpdateBody
*/
type CalculateProjectResourceQuotaUpdateBody struct {

	// DiskSizeGB will be processed only for those providers which don't have the disk size in their API objects, like AWS, Alibabla and GCP.
	DiskSizeGB int64 `json:"diskSizeGB,omitempty"`

	// replicas
	Replicas int64 `json:"replicas,omitempty"`

	// alibaba instance type
	AlibabaInstanceType *models.AlibabaInstanceType `json:"alibabaInstanceType,omitempty"`

	// anexia node spec
	AnexiaNodeSpec *models.AnexiaNodeSpec `json:"anexiaNodeSpec,omitempty"`

	// aws size
	AwsSize *models.AWSSize `json:"awsSize,omitempty"`

	// azure size
	AzureSize *models.AzureSize `json:"azureSize,omitempty"`

	// do size
	DoSize *models.DigitaloceanSize `json:"doSize,omitempty"`

	// equinix size
	EquinixSize *models.PacketSize `json:"equinixSize,omitempty"`

	// gcp size
	GcpSize *models.GCPMachineSize `json:"gcpSize,omitempty"`

	// hetzner size
	HetznerSize *models.HetznerSize `json:"hetznerSize,omitempty"`

	// kubevirt node size
	KubevirtNodeSize *models.KubevirtNodeSize `json:"kubevirtNodeSize,omitempty"`

	// nutanix node spec
	NutanixNodeSpec *models.NutanixNodeSpec `json:"nutanixNodeSpec,omitempty"`

	// openstack size
	OpenstackSize *models.OpenstackSize `json:"openstackSize,omitempty"`

	// v sphere node spec
	VSphereNodeSpec *models.VSphereNodeSpec `json:"vSphereNodeSpec,omitempty"`

	// vm director node spec
	VMDirectorNodeSpec *models.VMwareCloudDirectorNodeSpec `json:"vmDirectorNodeSpec,omitempty"`
}

// Validate validates this calculate project resource quota update body
func (o *CalculateProjectResourceQuotaUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAlibabaInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAnexiaNodeSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAwsSize(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAzureSize(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDoSize(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEquinixSize(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGcpSize(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHetznerSize(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateKubevirtNodeSize(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNutanixNodeSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOpenstackSize(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVSphereNodeSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVMDirectorNodeSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) validateAlibabaInstanceType(formats strfmt.Registry) error {
	if swag.IsZero(o.AlibabaInstanceType) { // not required
		return nil
	}

	if o.AlibabaInstanceType != nil {
		if err := o.AlibabaInstanceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "alibabaInstanceType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "alibabaInstanceType")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) validateAnexiaNodeSpec(formats strfmt.Registry) error {
	if swag.IsZero(o.AnexiaNodeSpec) { // not required
		return nil
	}

	if o.AnexiaNodeSpec != nil {
		if err := o.AnexiaNodeSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "anexiaNodeSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "anexiaNodeSpec")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) validateAwsSize(formats strfmt.Registry) error {
	if swag.IsZero(o.AwsSize) { // not required
		return nil
	}

	if o.AwsSize != nil {
		if err := o.AwsSize.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "awsSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "awsSize")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) validateAzureSize(formats strfmt.Registry) error {
	if swag.IsZero(o.AzureSize) { // not required
		return nil
	}

	if o.AzureSize != nil {
		if err := o.AzureSize.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "azureSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "azureSize")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) validateDoSize(formats strfmt.Registry) error {
	if swag.IsZero(o.DoSize) { // not required
		return nil
	}

	if o.DoSize != nil {
		if err := o.DoSize.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "doSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "doSize")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) validateEquinixSize(formats strfmt.Registry) error {
	if swag.IsZero(o.EquinixSize) { // not required
		return nil
	}

	if o.EquinixSize != nil {
		if err := o.EquinixSize.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "equinixSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "equinixSize")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) validateGcpSize(formats strfmt.Registry) error {
	if swag.IsZero(o.GcpSize) { // not required
		return nil
	}

	if o.GcpSize != nil {
		if err := o.GcpSize.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "gcpSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "gcpSize")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) validateHetznerSize(formats strfmt.Registry) error {
	if swag.IsZero(o.HetznerSize) { // not required
		return nil
	}

	if o.HetznerSize != nil {
		if err := o.HetznerSize.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "hetznerSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "hetznerSize")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) validateKubevirtNodeSize(formats strfmt.Registry) error {
	if swag.IsZero(o.KubevirtNodeSize) { // not required
		return nil
	}

	if o.KubevirtNodeSize != nil {
		if err := o.KubevirtNodeSize.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "kubevirtNodeSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "kubevirtNodeSize")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) validateNutanixNodeSpec(formats strfmt.Registry) error {
	if swag.IsZero(o.NutanixNodeSpec) { // not required
		return nil
	}

	if o.NutanixNodeSpec != nil {
		if err := o.NutanixNodeSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "nutanixNodeSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "nutanixNodeSpec")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) validateOpenstackSize(formats strfmt.Registry) error {
	if swag.IsZero(o.OpenstackSize) { // not required
		return nil
	}

	if o.OpenstackSize != nil {
		if err := o.OpenstackSize.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "openstackSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "openstackSize")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) validateVSphereNodeSpec(formats strfmt.Registry) error {
	if swag.IsZero(o.VSphereNodeSpec) { // not required
		return nil
	}

	if o.VSphereNodeSpec != nil {
		if err := o.VSphereNodeSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "vSphereNodeSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "vSphereNodeSpec")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) validateVMDirectorNodeSpec(formats strfmt.Registry) error {
	if swag.IsZero(o.VMDirectorNodeSpec) { // not required
		return nil
	}

	if o.VMDirectorNodeSpec != nil {
		if err := o.VMDirectorNodeSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "vmDirectorNodeSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "vmDirectorNodeSpec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this calculate project resource quota update body based on the context it is used
func (o *CalculateProjectResourceQuotaUpdateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAlibabaInstanceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateAnexiaNodeSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateAwsSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateAzureSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDoSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateEquinixSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateGcpSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateHetznerSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateKubevirtNodeSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateNutanixNodeSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOpenstackSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateVSphereNodeSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateVMDirectorNodeSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) contextValidateAlibabaInstanceType(ctx context.Context, formats strfmt.Registry) error {

	if o.AlibabaInstanceType != nil {
		if err := o.AlibabaInstanceType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "alibabaInstanceType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "alibabaInstanceType")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) contextValidateAnexiaNodeSpec(ctx context.Context, formats strfmt.Registry) error {

	if o.AnexiaNodeSpec != nil {
		if err := o.AnexiaNodeSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "anexiaNodeSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "anexiaNodeSpec")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) contextValidateAwsSize(ctx context.Context, formats strfmt.Registry) error {

	if o.AwsSize != nil {
		if err := o.AwsSize.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "awsSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "awsSize")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) contextValidateAzureSize(ctx context.Context, formats strfmt.Registry) error {

	if o.AzureSize != nil {
		if err := o.AzureSize.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "azureSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "azureSize")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) contextValidateDoSize(ctx context.Context, formats strfmt.Registry) error {

	if o.DoSize != nil {
		if err := o.DoSize.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "doSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "doSize")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) contextValidateEquinixSize(ctx context.Context, formats strfmt.Registry) error {

	if o.EquinixSize != nil {
		if err := o.EquinixSize.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "equinixSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "equinixSize")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) contextValidateGcpSize(ctx context.Context, formats strfmt.Registry) error {

	if o.GcpSize != nil {
		if err := o.GcpSize.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "gcpSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "gcpSize")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) contextValidateHetznerSize(ctx context.Context, formats strfmt.Registry) error {

	if o.HetznerSize != nil {
		if err := o.HetznerSize.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "hetznerSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "hetznerSize")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) contextValidateKubevirtNodeSize(ctx context.Context, formats strfmt.Registry) error {

	if o.KubevirtNodeSize != nil {
		if err := o.KubevirtNodeSize.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "kubevirtNodeSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "kubevirtNodeSize")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) contextValidateNutanixNodeSpec(ctx context.Context, formats strfmt.Registry) error {

	if o.NutanixNodeSpec != nil {
		if err := o.NutanixNodeSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "nutanixNodeSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "nutanixNodeSpec")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) contextValidateOpenstackSize(ctx context.Context, formats strfmt.Registry) error {

	if o.OpenstackSize != nil {
		if err := o.OpenstackSize.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "openstackSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "openstackSize")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) contextValidateVSphereNodeSpec(ctx context.Context, formats strfmt.Registry) error {

	if o.VSphereNodeSpec != nil {
		if err := o.VSphereNodeSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "vSphereNodeSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "vSphereNodeSpec")
			}
			return err
		}
	}

	return nil
}

func (o *CalculateProjectResourceQuotaUpdateBody) contextValidateVMDirectorNodeSpec(ctx context.Context, formats strfmt.Registry) error {

	if o.VMDirectorNodeSpec != nil {
		if err := o.VMDirectorNodeSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "vmDirectorNodeSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "vmDirectorNodeSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CalculateProjectResourceQuotaUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CalculateProjectResourceQuotaUpdateBody) UnmarshalBinary(b []byte) error {
	var res CalculateProjectResourceQuotaUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}