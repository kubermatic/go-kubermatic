// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AzureNodeSpec AzureNodeSpec describes settings for an Azure node
//
// swagger:model AzureNodeSpec
type AzureNodeSpec struct {

	// AssignAvailabilitySet is used to check if an availability set should be created and assigned to the cluster.
	AssignAvailabilitySet bool `json:"assignAvailabilitySet,omitempty"`

	// should the machine have a publicly accessible IP address
	AssignPublicIP bool `json:"assignPublicIP,omitempty"`

	// Data disk size in GB
	DataDiskSize int32 `json:"dataDiskSize,omitempty"`

	// EnableAcceleratedNetworking is used to check if an accelerating networking should be used for azure vms.
	EnableAcceleratedNetworking bool `json:"enableAcceleratedNetworking,omitempty"`

	// ImageID represents the ID of the image that should be used to run the node
	ImageID string `json:"imageID,omitempty"`

	// OS disk size in GB
	OSDiskSize int32 `json:"osDiskSize,omitempty"`

	// VM size
	// Required: true
	Size *string `json:"size"`

	// Additional metadata to set
	Tags map[string]string `json:"tags,omitempty"`

	// Zones represents the availability zones for azure vms
	Zones []string `json:"zones"`
}

// Validate validates this azure node spec
func (m *AzureNodeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureNodeSpec) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this azure node spec based on context it is used
func (m *AzureNodeSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureNodeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureNodeSpec) UnmarshalBinary(b []byte) error {
	var res AzureNodeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
