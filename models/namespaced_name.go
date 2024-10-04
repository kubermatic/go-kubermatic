// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NamespacedName namespaced name
//
// swagger:model NamespacedName
type NamespacedName struct {

	// name
	Name string `json:"Name,omitempty"`

	// namespace
	Namespace string `json:"Namespace,omitempty"`
}

// Validate validates this namespaced name
func (m *NamespacedName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this namespaced name based on context it is used
func (m *NamespacedName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NamespacedName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamespacedName) UnmarshalBinary(b []byte) error {
	var res NamespacedName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
