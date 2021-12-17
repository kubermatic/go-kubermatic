// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CPUFeature CPUFeature allows specifying a CPU feature.
//
// swagger:model CPUFeature
type CPUFeature struct {

	// Name of the CPU feature
	Name string `json:"name,omitempty"`

	// Policy is the CPU feature attribute which can have the following attributes:
	// force    - The virtual CPU will claim the feature is supported regardless of it being supported by host CPU.
	// require  - Guest creation will fail unless the feature is supported by the host CPU or the hypervisor is able to emulate it.
	// optional - The feature will be supported by virtual CPU if and only if it is supported by host CPU.
	// disable  - The feature will not be supported by virtual CPU.
	// forbid   - Guest creation will fail if the feature is supported by host CPU.
	// Defaults to require
	// +optional
	Policy string `json:"policy,omitempty"`
}

// Validate validates this CPU feature
func (m *CPUFeature) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this CPU feature based on context it is used
func (m *CPUFeature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CPUFeature) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CPUFeature) UnmarshalBinary(b []byte) error {
	var res CPUFeature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
