// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Hetzner hetzner
//
// swagger:model Hetzner
type Hetzner struct {

	// If datacenter is set, this preset is only applicable to the
	// configured datacenter.
	Datacenter string `json:"datacenter,omitempty"`

	// Only enabled presets will be available in the KKP dashboard.
	Enabled bool `json:"enabled,omitempty"`

	// Network is the pre-existing Hetzner network in which the machines are running.
	// While machines can be in multiple networks, a single one must be chosen for the
	// HCloud CCM to work.
	// If this is empty, the network configured on the datacenter will be used.
	Network string `json:"network,omitempty"`

	// Token is used to authenticate with the Hetzner API.
	Token string `json:"token,omitempty"`
}

// Validate validates this hetzner
func (m *Hetzner) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hetzner based on context it is used
func (m *Hetzner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Hetzner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Hetzner) UnmarshalBinary(b []byte) error {
	var res Hetzner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
