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

// CbslBody cbsl body
//
// swagger:model CbslBody
type CbslBody struct {

	// Name of the cluster backup
	Name string `json:"name,omitempty"`

	// cbsl spec
	CbslSpec *BackupStorageLocationSpec `json:"cbslSpec,omitempty"`

	// credentials
	Credentials *S3BackupCredentials `json:"credentials,omitempty"`
}

// Validate validates this cbsl body
func (m *CbslBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCbslSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CbslBody) validateCbslSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.CbslSpec) { // not required
		return nil
	}

	if m.CbslSpec != nil {
		if err := m.CbslSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cbslSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cbslSpec")
			}
			return err
		}
	}

	return nil
}

func (m *CbslBody) validateCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cbsl body based on the context it is used
func (m *CbslBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCbslSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CbslBody) contextValidateCbslSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.CbslSpec != nil {
		if err := m.CbslSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cbslSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cbslSpec")
			}
			return err
		}
	}

	return nil
}

func (m *CbslBody) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.Credentials != nil {
		if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CbslBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CbslBody) UnmarshalBinary(b []byte) error {
	var res CbslBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
