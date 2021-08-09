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
)

// Sync sync
//
// swagger:model Sync
type Sync struct {

	// If non-empty, entries on this list will be replicated into OPA
	SyncOnly []*GVK `json:"syncOnly"`
}

// Validate validates this sync
func (m *Sync) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSyncOnly(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Sync) validateSyncOnly(formats strfmt.Registry) error {
	if swag.IsZero(m.SyncOnly) { // not required
		return nil
	}

	for i := 0; i < len(m.SyncOnly); i++ {
		if swag.IsZero(m.SyncOnly[i]) { // not required
			continue
		}

		if m.SyncOnly[i] != nil {
			if err := m.SyncOnly[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("syncOnly" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sync based on the context it is used
func (m *Sync) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSyncOnly(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Sync) contextValidateSyncOnly(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SyncOnly); i++ {

		if m.SyncOnly[i] != nil {
			if err := m.SyncOnly[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("syncOnly" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Sync) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Sync) UnmarshalBinary(b []byte) error {
	var res Sync
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
