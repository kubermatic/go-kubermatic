// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterSpec ClusterSpec defines the cluster specification
//
// swagger:model ClusterSpec
type ClusterSpec struct {

	// Additional Admission Controller plugins
	AdmissionPlugins []string `json:"admissionPlugins"`

	// EnableUserSSHKeyAgent control whether the UserSSHKeyAgent will be deployed in the user cluster or not.
	// If it was enabled, the agent will be deployed and used to sync the user ssh keys, that the user attach
	// to the created cluster. If the agent was disabled, it won't be deployed in the user cluster, thus after
	// the cluster creation any attached ssh keys won't be synced to the worker nodes. Once the agent is enabled/disabled
	// it cannot be changed after the cluster is being created.
	EnableUserSSHKeyAgent bool `json:"enableUserSSHKeyAgent,omitempty"`

	// MachineNetworks optionally specifies the parameters for IPAM.
	MachineNetworks []*MachineNetworkingConfig `json:"machineNetworks"`

	// PodNodeSelectorAdmissionPluginConfig provides the configuration for the PodNodeSelector.
	// It's used by the backend to create a configuration file for this plugin.
	// The key:value from the map is converted to the namespace:<node-selectors-labels> in the file.
	// The format in a file:
	// podNodeSelectorPluginConfig:
	// clusterDefaultNodeSelector: <node-selectors-labels>
	// namespace1: <node-selectors-labels>
	// namespace2: <node-selectors-labels>
	PodNodeSelectorAdmissionPluginConfig map[string]string `json:"podNodeSelectorAdmissionPluginConfig,omitempty"`

	// If active the PodNodeSelector admission plugin is configured at the apiserver
	UsePodNodeSelectorAdmissionPlugin bool `json:"usePodNodeSelectorAdmissionPlugin,omitempty"`

	// If active the PodSecurityPolicy admission plugin is configured at the apiserver
	UsePodSecurityPolicyAdmissionPlugin bool `json:"usePodSecurityPolicyAdmissionPlugin,omitempty"`

	// audit logging
	AuditLogging *AuditLoggingSettings `json:"auditLogging,omitempty"`

	// cloud
	Cloud *CloudSpec `json:"cloud,omitempty"`

	// cluster network
	ClusterNetwork *ClusterNetworkingConfig `json:"clusterNetwork,omitempty"`

	// mla
	Mla *MLASettings `json:"mla,omitempty"`

	// oidc
	Oidc *OIDCSettings `json:"oidc,omitempty"`

	// opa integration
	OpaIntegration *OPAIntegrationSettings `json:"opaIntegration,omitempty"`

	// service account
	ServiceAccount *ServiceAccountSettings `json:"serviceAccount,omitempty"`

	// update window
	UpdateWindow *UpdateWindow `json:"updateWindow,omitempty"`

	// version
	Version Semver `json:"version,omitempty"`
}

// Validate validates this cluster spec
func (m *ClusterSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMachineNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuditLogging(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloud(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMla(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOidc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpaIntegration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateWindow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSpec) validateMachineNetworks(formats strfmt.Registry) error {

	if swag.IsZero(m.MachineNetworks) { // not required
		return nil
	}

	for i := 0; i < len(m.MachineNetworks); i++ {
		if swag.IsZero(m.MachineNetworks[i]) { // not required
			continue
		}

		if m.MachineNetworks[i] != nil {
			if err := m.MachineNetworks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machineNetworks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterSpec) validateAuditLogging(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditLogging) { // not required
		return nil
	}

	if m.AuditLogging != nil {
		if err := m.AuditLogging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auditLogging")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpec) validateCloud(formats strfmt.Registry) error {

	if swag.IsZero(m.Cloud) { // not required
		return nil
	}

	if m.Cloud != nil {
		if err := m.Cloud.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpec) validateClusterNetwork(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterNetwork) { // not required
		return nil
	}

	if m.ClusterNetwork != nil {
		if err := m.ClusterNetwork.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterNetwork")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpec) validateMla(formats strfmt.Registry) error {

	if swag.IsZero(m.Mla) { // not required
		return nil
	}

	if m.Mla != nil {
		if err := m.Mla.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mla")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpec) validateOidc(formats strfmt.Registry) error {

	if swag.IsZero(m.Oidc) { // not required
		return nil
	}

	if m.Oidc != nil {
		if err := m.Oidc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oidc")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpec) validateOpaIntegration(formats strfmt.Registry) error {

	if swag.IsZero(m.OpaIntegration) { // not required
		return nil
	}

	if m.OpaIntegration != nil {
		if err := m.OpaIntegration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("opaIntegration")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpec) validateServiceAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceAccount) { // not required
		return nil
	}

	if m.ServiceAccount != nil {
		if err := m.ServiceAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceAccount")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpec) validateUpdateWindow(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdateWindow) { // not required
		return nil
	}

	if m.UpdateWindow != nil {
		if err := m.UpdateWindow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateWindow")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSpec) UnmarshalBinary(b []byte) error {
	var res ClusterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
