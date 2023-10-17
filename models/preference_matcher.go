// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PreferenceMatcher PreferenceMatcher references a set of preference that is used to fill fields in the VMI template.
//
// swagger:model PreferenceMatcher
type PreferenceMatcher struct {

	// InferFromVolume lists the name of a volume that should be used to infer or discover the preference
	// to be used through known annotations on the underlying resource. Once applied to the PreferenceMatcher
	// this field is removed.
	//
	// +optional
	InferFromVolume string `json:"inferFromVolume,omitempty"`

	// Kind specifies which preference resource is referenced.
	// Allowed values are: "VirtualMachinePreference" and "VirtualMachineClusterPreference".
	// If not specified, "VirtualMachineClusterPreference" is used by default.
	//
	// +optional
	Kind string `json:"kind,omitempty"`

	// Name is the name of the VirtualMachinePreference or VirtualMachineClusterPreference
	//
	// +optional
	Name string `json:"name,omitempty"`

	// RevisionName specifies a ControllerRevision containing a specific copy of the
	// VirtualMachinePreference or VirtualMachineClusterPreference to be used. This is
	// initially captured the first time the instancetype is applied to the VirtualMachineInstance.
	//
	// +optional
	RevisionName string `json:"revisionName,omitempty"`
}

// Validate validates this preference matcher
func (m *PreferenceMatcher) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this preference matcher based on context it is used
func (m *PreferenceMatcher) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PreferenceMatcher) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PreferenceMatcher) UnmarshalBinary(b []byte) error {
	var res PreferenceMatcher
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
