// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenstackTenant OpenstackTenant is the object representing a openstack tenant.
// swagger:model OpenstackTenant
type OpenstackTenant struct {

	// Id uniquely identifies the current tenant
	ID string `json:"id,omitempty"`

	// Name is the name of the tenant
	Name string `json:"name,omitempty"`
}

// Validate validates this openstack tenant
func (m *OpenstackTenant) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenstackTenant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenstackTenant) UnmarshalBinary(b []byte) error {
	var res OpenstackTenant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
