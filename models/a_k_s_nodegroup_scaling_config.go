// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AKSNodegroupScalingConfig a k s nodegroup scaling config
//
// swagger:model AKSNodegroupScalingConfig
type AKSNodegroupScalingConfig struct {

	// MaxCount - The maximum number of nodes for auto-scaling
	MaxCount int32 `json:"maxCount,omitempty"`

	// MinCount - The minimum number of nodes for auto-scaling
	MinCount int32 `json:"minCount,omitempty"`
}

// Validate validates this a k s nodegroup scaling config
func (m *AKSNodegroupScalingConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this a k s nodegroup scaling config based on context it is used
func (m *AKSNodegroupScalingConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AKSNodegroupScalingConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AKSNodegroupScalingConfig) UnmarshalBinary(b []byte) error {
	var res AKSNodegroupScalingConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
