// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigMapKeySelector Selects a key from a ConfigMap.
//
// +structType=atomic
//
// swagger:model ConfigMapKeySelector
type ConfigMapKeySelector struct {

	// The key to select.
	Key string `json:"key,omitempty"`

	// Name of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// TODO: Add other useful fields. apiVersion, kind, uid?
	// +optional
	Name string `json:"name,omitempty"`

	// Specify whether the ConfigMap or its key must be defined
	// +optional
	Optional bool `json:"optional,omitempty"`
}

// Validate validates this config map key selector
func (m *ConfigMapKeySelector) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this config map key selector based on context it is used
func (m *ConfigMapKeySelector) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigMapKeySelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigMapKeySelector) UnmarshalBinary(b []byte) error {
	var res ConfigMapKeySelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
