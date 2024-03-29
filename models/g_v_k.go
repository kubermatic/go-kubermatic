// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GVK GVK group version kind of a resource.
//
// swagger:model GVK
type GVK struct {

	// group
	Group string `json:"group,omitempty"`

	// kind
	Kind string `json:"kind,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this g v k
func (m *GVK) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this g v k based on context it is used
func (m *GVK) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GVK) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GVK) UnmarshalBinary(b []byte) error {
	var res GVK
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
