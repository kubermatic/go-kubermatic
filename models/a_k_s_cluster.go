// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AKSCluster AKSCluster represents an object of AKS cluster.
//
// swagger:model AKSCluster
type AKSCluster struct {

	// is imported
	IsImported bool `json:"imported,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// resource group
	ResourceGroup string `json:"resourceGroup,omitempty"`
}

// Validate validates this a k s cluster
func (m *AKSCluster) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this a k s cluster based on context it is used
func (m *AKSCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AKSCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AKSCluster) UnmarshalBinary(b []byte) error {
	var res AKSCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
