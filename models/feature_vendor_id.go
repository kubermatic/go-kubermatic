// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FeatureVendorID feature vendor ID
//
// swagger:model FeatureVendorID
type FeatureVendorID struct {

	// Enabled determines if the feature should be enabled or disabled on the guest.
	// Defaults to true.
	// +optional
	Enabled bool `json:"enabled,omitempty"`

	// VendorID sets the hypervisor vendor id, visible to the vmi.
	// String up to twelve characters.
	VendorID string `json:"vendorid,omitempty"`
}

// Validate validates this feature vendor ID
func (m *FeatureVendorID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this feature vendor ID based on context it is used
func (m *FeatureVendorID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FeatureVendorID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeatureVendorID) UnmarshalBinary(b []byte) error {
	var res FeatureVendorID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
