// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TaintSpec TaintSpec defines a node taint
//
// swagger:model TaintSpec
type TaintSpec struct {

	// effect
	Effect string `json:"effect,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this taint spec
func (m *TaintSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this taint spec based on context it is used
func (m *TaintSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaintSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaintSpec) UnmarshalBinary(b []byte) error {
	var res TaintSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
