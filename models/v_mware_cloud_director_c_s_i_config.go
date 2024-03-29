// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VMwareCloudDirectorCSIConfig v mware cloud director c s i config
//
// swagger:model VMwareCloudDirectorCSIConfig
type VMwareCloudDirectorCSIConfig struct {

	// Filesystem to use for named disks, defaults to "ext4"
	// +optional
	Filesystem string `json:"filesystem,omitempty"`

	// The name of the storage profile to use for disks created by CSI driver
	StorageProfile string `json:"storageProfile,omitempty"`
}

// Validate validates this v mware cloud director c s i config
func (m *VMwareCloudDirectorCSIConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v mware cloud director c s i config based on context it is used
func (m *VMwareCloudDirectorCSIConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMwareCloudDirectorCSIConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMwareCloudDirectorCSIConfig) UnmarshalBinary(b []byte) error {
	var res VMwareCloudDirectorCSIConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
