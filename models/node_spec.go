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

// NodeSpec NodeSpec node specification
//
// swagger:model NodeSpec
type NodeSpec struct {

	// Map of string keys and values that can be used to organize and categorize (scope and select) objects.
	// It will be applied to Nodes allowing users run their apps on specific Node using labelSelector.
	Labels map[string]string `json:"labels,omitempty"`

	// SSH user name
	SSHUserName string `json:"sshUserName,omitempty"`

	// List of taints to set on new nodes
	Taints []*TaintSpec `json:"taints"`

	// cloud
	// Required: true
	Cloud *NodeCloudSpec `json:"cloud"`

	// network
	Network *NetworkSpec `json:"network,omitempty"`

	// operating system
	// Required: true
	OperatingSystem *OperatingSystemSpec `json:"operatingSystem"`

	// versions
	// Required: true
	Versions *NodeVersionInfo `json:"versions"`
}

// Validate validates this node spec
func (m *NodeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloud(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatingSystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeSpec) validateTaints(formats strfmt.Registry) error {
	if swag.IsZero(m.Taints) { // not required
		return nil
	}

	for i := 0; i < len(m.Taints); i++ {
		if swag.IsZero(m.Taints[i]) { // not required
			continue
		}

		if m.Taints[i] != nil {
			if err := m.Taints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("taints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NodeSpec) validateCloud(formats strfmt.Registry) error {

	if err := validate.Required("cloud", "body", m.Cloud); err != nil {
		return err
	}

	if m.Cloud != nil {
		if err := m.Cloud.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloud")
			}
			return err
		}
	}

	return nil
}

func (m *NodeSpec) validateNetwork(formats strfmt.Registry) error {
	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *NodeSpec) validateOperatingSystem(formats strfmt.Registry) error {

	if err := validate.Required("operatingSystem", "body", m.OperatingSystem); err != nil {
		return err
	}

	if m.OperatingSystem != nil {
		if err := m.OperatingSystem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operatingSystem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operatingSystem")
			}
			return err
		}
	}

	return nil
}

func (m *NodeSpec) validateVersions(formats strfmt.Registry) error {

	if err := validate.Required("versions", "body", m.Versions); err != nil {
		return err
	}

	if m.Versions != nil {
		if err := m.Versions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("versions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this node spec based on the context it is used
func (m *NodeSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTaints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCloud(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperatingSystem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeSpec) contextValidateTaints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Taints); i++ {

		if m.Taints[i] != nil {
			if err := m.Taints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("taints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NodeSpec) contextValidateCloud(ctx context.Context, formats strfmt.Registry) error {

	if m.Cloud != nil {
		if err := m.Cloud.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloud")
			}
			return err
		}
	}

	return nil
}

func (m *NodeSpec) contextValidateNetwork(ctx context.Context, formats strfmt.Registry) error {

	if m.Network != nil {
		if err := m.Network.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *NodeSpec) contextValidateOperatingSystem(ctx context.Context, formats strfmt.Registry) error {

	if m.OperatingSystem != nil {
		if err := m.OperatingSystem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operatingSystem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operatingSystem")
			}
			return err
		}
	}

	return nil
}

func (m *NodeSpec) contextValidateVersions(ctx context.Context, formats strfmt.Registry) error {

	if m.Versions != nil {
		if err := m.Versions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("versions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeSpec) UnmarshalBinary(b []byte) error {
	var res NodeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
