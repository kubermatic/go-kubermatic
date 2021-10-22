// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GKECloudSpec g k e cloud spec
//
// swagger:model GKECloudSpec
type GKECloudSpec struct {

	// name
	Name string `json:"name,omitempty"`

	// service account
	ServiceAccount string `json:"serviceAccount,omitempty"`
}

// Validate validates this g k e cloud spec
func (m *GKECloudSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this g k e cloud spec based on context it is used
func (m *GKECloudSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GKECloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GKECloudSpec) UnmarshalBinary(b []byte) error {
	var res GKECloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
