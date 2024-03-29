// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KubeOneEquinixCloudSpec KubeOneEquinixCloudSpec specifies access data to a Equinix cloud.
//
// swagger:model KubeOneEquinixCloudSpec
type KubeOneEquinixCloudSpec struct {

	// API key
	APIKey string `json:"apiKey,omitempty"`

	// project ID
	ProjectID string `json:"projectID,omitempty"`
}

// Validate validates this kube one equinix cloud spec
func (m *KubeOneEquinixCloudSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kube one equinix cloud spec based on context it is used
func (m *KubeOneEquinixCloudSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubeOneEquinixCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeOneEquinixCloudSpec) UnmarshalBinary(b []byte) error {
	var res KubeOneEquinixCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
