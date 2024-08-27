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

// DatacenterSpecTinkerbell DatacenterSepcTinkerbell contains spec for tinkerbell provider.
//
// swagger:model DatacenterSpecTinkerbell
type DatacenterSpecTinkerbell struct {

	// images
	Images *TinkerbellImageSources `json:"images,omitempty"`
}

// Validate validates this datacenter spec tinkerbell
func (m *DatacenterSpecTinkerbell) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatacenterSpecTinkerbell) validateImages(formats strfmt.Registry) error {
	if swag.IsZero(m.Images) { // not required
		return nil
	}

	if m.Images != nil {
		if err := m.Images.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("images")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("images")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this datacenter spec tinkerbell based on the context it is used
func (m *DatacenterSpecTinkerbell) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatacenterSpecTinkerbell) contextValidateImages(ctx context.Context, formats strfmt.Registry) error {

	if m.Images != nil {
		if err := m.Images.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("images")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("images")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatacenterSpecTinkerbell) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatacenterSpecTinkerbell) UnmarshalBinary(b []byte) error {
	var res DatacenterSpecTinkerbell
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}