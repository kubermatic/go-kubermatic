// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VirtualMachinePreferenceList VirtualMachinePreferenceList represents a list of VirtualMachinePreference.
//
// VirtualMachinePreference are divided into 2 categories: "custom" or "kubermatic".
//
// swagger:model VirtualMachinePreferenceList
type VirtualMachinePreferenceList struct {

	// preferences
	Preferences map[string][]VirtualMachinePreference `json:"preferences,omitempty"`
}

// Validate validates this virtual machine preference list
func (m *VirtualMachinePreferenceList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePreferences(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachinePreferenceList) validatePreferences(formats strfmt.Registry) error {
	if swag.IsZero(m.Preferences) { // not required
		return nil
	}

	for k := range m.Preferences {

		if err := validate.Required("preferences"+"."+k, "body", m.Preferences[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.Preferences[k]); i++ {

			if err := m.Preferences[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("preferences" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("preferences" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

// ContextValidate validate this virtual machine preference list based on the context it is used
func (m *VirtualMachinePreferenceList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePreferences(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachinePreferenceList) contextValidatePreferences(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Preferences {

		for i := 0; i < len(m.Preferences[k]); i++ {

			if err := m.Preferences[k][i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("preferences" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("preferences" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualMachinePreferenceList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualMachinePreferenceList) UnmarshalBinary(b []byte) error {
	var res VirtualMachinePreferenceList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
