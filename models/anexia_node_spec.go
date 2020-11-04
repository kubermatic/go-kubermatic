// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AnexiaNodeSpec AnexiaNodeSpec anexia specific node settings
//
// swagger:model AnexiaNodeSpec
type AnexiaNodeSpec struct {

	// CPUs states how many cpus the node will have.
	// Required: true
	CPUs *int64 `json:"cpus"`

	// DiskSize states the disk size that node will have.
	// Required: true
	DiskSize *int64 `json:"diskSize"`

	// Memory states the memory that node will have.
	// Required: true
	Memory *int64 `json:"memory"`

	// TemplateID instance template
	// Required: true
	TemplateID *string `json:"templateID"`

	// VlanID Instance vlanID
	// Required: true
	VlanID *string `json:"vlanID"`
}

// Validate validates this anexia node spec
func (m *AnexiaNodeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnexiaNodeSpec) validateCPUs(formats strfmt.Registry) error {

	if err := validate.Required("cpus", "body", m.CPUs); err != nil {
		return err
	}

	return nil
}

func (m *AnexiaNodeSpec) validateDiskSize(formats strfmt.Registry) error {

	if err := validate.Required("diskSize", "body", m.DiskSize); err != nil {
		return err
	}

	return nil
}

func (m *AnexiaNodeSpec) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	return nil
}

func (m *AnexiaNodeSpec) validateTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("templateID", "body", m.TemplateID); err != nil {
		return err
	}

	return nil
}

func (m *AnexiaNodeSpec) validateVlanID(formats strfmt.Registry) error {

	if err := validate.Required("vlanID", "body", m.VlanID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnexiaNodeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnexiaNodeSpec) UnmarshalBinary(b []byte) error {
	var res AnexiaNodeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
