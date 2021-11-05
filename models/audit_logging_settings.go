// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuditLoggingSettings audit logging settings
//
// swagger:model AuditLoggingSettings
type AuditLoggingSettings struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// policy preset
	PolicyPreset AuditPolicyPreset `json:"policyPreset,omitempty"`
}

// Validate validates this audit logging settings
func (m *AuditLoggingSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicyPreset(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditLoggingSettings) validatePolicyPreset(formats strfmt.Registry) error {
	if swag.IsZero(m.PolicyPreset) { // not required
		return nil
	}

	if err := m.PolicyPreset.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("policyPreset")
		}
		return err
	}

	return nil
}

// ContextValidate validate this audit logging settings based on the context it is used
func (m *AuditLoggingSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePolicyPreset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditLoggingSettings) contextValidatePolicyPreset(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PolicyPreset.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("policyPreset")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditLoggingSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditLoggingSettings) UnmarshalBinary(b []byte) error {
	var res AuditLoggingSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
