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

// NodeSettings NodeSettings are node specific flags which can be configured on datacenter level.
//
// swagger:model NodeSettings
type NodeSettings struct {

	// Optional: These image registries will be configured as insecure
	// on the container runtime.
	InsecureRegistries []string `json:"insecureRegistries"`

	// Optional: Translates to --pod-infra-container-image on the kubelet.
	// If not set, the kubelet will default it.
	PauseImage string `json:"pauseImage,omitempty"`

	// Optional: These image registries will be configured as registry mirrors
	// on the container runtime.
	RegistryMirrors []string `json:"registryMirrors"`

	// http proxy
	HTTPProxy ProxyValue `json:"httpProxy,omitempty"`

	// no proxy
	NoProxy ProxyValue `json:"noProxy,omitempty"`
}

// Validate validates this node settings
func (m *NodeSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTTPProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNoProxy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeSettings) validateHTTPProxy(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPProxy) { // not required
		return nil
	}

	if err := m.HTTPProxy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("httpProxy")
		}
		return err
	}

	return nil
}

func (m *NodeSettings) validateNoProxy(formats strfmt.Registry) error {
	if swag.IsZero(m.NoProxy) { // not required
		return nil
	}

	if err := m.NoProxy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("noProxy")
		}
		return err
	}

	return nil
}

// ContextValidate validate this node settings based on the context it is used
func (m *NodeSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHTTPProxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNoProxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeSettings) contextValidateHTTPProxy(ctx context.Context, formats strfmt.Registry) error {

	if err := m.HTTPProxy.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("httpProxy")
		}
		return err
	}

	return nil
}

func (m *NodeSettings) contextValidateNoProxy(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NoProxy.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("noProxy")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeSettings) UnmarshalBinary(b []byte) error {
	var res NodeSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
