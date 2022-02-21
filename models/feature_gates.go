// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FeatureGates FeatureGates represents an object holding feature gate settings
//
// swagger:model FeatureGates
type FeatureGates struct {

	// konnectivity service
	KonnectivityService bool `json:"konnectivityService,omitempty"`

	// o ID c kube cfg endpoint
	OIDCKubeCfgEndpoint bool `json:"oidcKubeCfgEndpoint,omitempty"`
}

// Validate validates this feature gates
func (m *FeatureGates) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this feature gates based on context it is used
func (m *FeatureGates) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FeatureGates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeatureGates) UnmarshalBinary(b []byte) error {
	var res FeatureGates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
