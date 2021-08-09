// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KubevirtNodeSpec KubevirtNodeSpec kubevirt specific node settings
//
// swagger:model KubevirtNodeSpec
type KubevirtNodeSpec struct {

	// CPUs states how many cpus the kubevirt node will have.
	// Required: true
	CPUs *string `json:"cpus"`

	// Memory states the memory that kubevirt node will have.
	// Required: true
	Memory *string `json:"memory"`

	// Namespace states in which namespace kubevirt node will be provisioned.
	// Required: true
	Namespace *string `json:"namespace"`

	// PVCSize states the size of the provisioned pvc per node.
	// Required: true
	PVCSize *string `json:"pvcSize"`

	// SourceURL states the url from which the imported image will be downloaded.
	// Required: true
	SourceURL *string `json:"sourceURL"`

	// StorageClassName states the storage class name for the provisioned PVCs.
	// Required: true
	StorageClassName *string `json:"storageClassName"`
}

// Validate validates this kubevirt node spec
func (m *KubevirtNodeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePVCSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageClassName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubevirtNodeSpec) validateCPUs(formats strfmt.Registry) error {

	if err := validate.Required("cpus", "body", m.CPUs); err != nil {
		return err
	}

	return nil
}

func (m *KubevirtNodeSpec) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	return nil
}

func (m *KubevirtNodeSpec) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *KubevirtNodeSpec) validatePVCSize(formats strfmt.Registry) error {

	if err := validate.Required("pvcSize", "body", m.PVCSize); err != nil {
		return err
	}

	return nil
}

func (m *KubevirtNodeSpec) validateSourceURL(formats strfmt.Registry) error {

	if err := validate.Required("sourceURL", "body", m.SourceURL); err != nil {
		return err
	}

	return nil
}

func (m *KubevirtNodeSpec) validateStorageClassName(formats strfmt.Registry) error {

	if err := validate.Required("storageClassName", "body", m.StorageClassName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this kubevirt node spec based on context it is used
func (m *KubevirtNodeSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubevirtNodeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubevirtNodeSpec) UnmarshalBinary(b []byte) error {
	var res KubevirtNodeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
