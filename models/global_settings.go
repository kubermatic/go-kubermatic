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

// GlobalSettings GlobalSettings defines global settings
//
// swagger:model GlobalSettings
type GlobalSettings struct {

	// AllowedOperatingSystems shows the available operating systems to use in the machine deployment.
	AllowedOperatingSystems map[string]bool `json:"allowedOperatingSystems,omitempty"`

	// DefaultNodeCount is the default number of replicas for the initial MachineDeployment.
	DefaultNodeCount int8 `json:"defaultNodeCount,omitempty"`

	// DisableAdminKubeconfig disables the admin kubeconfig functionality on the dashboard.
	DisableAdminKubeconfig bool `json:"disableAdminKubeconfig,omitempty"`

	// DisableChangelogPopup disables the changelog popup in KKP dashboard.
	DisableChangelogPopup bool `json:"disableChangelogPopup,omitempty"`

	// DisplayDemoInfo controls whether a a link to the KKP API documentation is shown in the footer.
	DisplayAPIDocs bool `json:"displayAPIDocs,omitempty"`

	// DisplayDemoInfo controls whether a "Demo System" hint is shown in the footer.
	DisplayDemoInfo bool `json:"displayDemoInfo,omitempty"`

	// DisplayDemoInfo controls whether a a link to TOS is shown in the footer.
	DisplayTermsOfService bool `json:"displayTermsOfService,omitempty"`

	// EnableClusterBackups enables the Cluster Backup feature in the dashboard.
	EnableClusterBackups bool `json:"enableClusterBackups,omitempty"`

	// EnableDashboard enables the link to the Kubernetes dashboard for a user cluster.
	EnableDashboard bool `json:"enableDashboard,omitempty"`

	// enable external cluster import
	EnableExternalClusterImport bool `json:"enableExternalClusterImport,omitempty"`

	// enable o ID c kubeconfig
	EnableOIDCKubeconfig bool `json:"enableOIDCKubeconfig,omitempty"`

	// EnableShareCluster enables the Share Cluster feature for the user clusters.
	EnableShareCluster bool `json:"enableShareCluster,omitempty"`

	// EnableWebTerminal enables the Web Terminal feature for the user clusters.
	EnableWebTerminal bool `json:"enableWebTerminal,omitempty"`

	// mla alertmanager prefix
	MlaAlertmanagerPrefix string `json:"mlaAlertmanagerPrefix,omitempty"`

	// mla grafana prefix
	MlaGrafanaPrefix string `json:"mlaGrafanaPrefix,omitempty"`

	// restrict project creation
	RestrictProjectCreation bool `json:"restrictProjectCreation,omitempty"`

	// restrict project deletion
	RestrictProjectDeletion bool `json:"restrictProjectDeletion,omitempty"`

	// UserProjectsLimit is the maximum number of projects a user can create.
	UserProjectsLimit int64 `json:"userProjectsLimit,omitempty"`

	// cleanup options
	CleanupOptions *CleanupOptions `json:"cleanupOptions,omitempty"`

	// custom links
	CustomLinks CustomLinks `json:"customLinks,omitempty"`

	// default quota
	DefaultQuota *ProjectResourceQuota `json:"defaultQuota,omitempty"`

	// machine deployment options
	MachineDeploymentOptions *MachineDeploymentOptions `json:"machineDeploymentOptions,omitempty"`

	// machine deployment VM resource quota
	MachineDeploymentVMResourceQuota *MachineFlavorFilter `json:"machineDeploymentVMResourceQuota,omitempty"`

	// mla options
	MlaOptions *MlaOptions `json:"mlaOptions,omitempty"`

	// notifications
	Notifications *NotificationsOptions `json:"notifications,omitempty"`

	// opa options
	OpaOptions *OpaOptions `json:"opaOptions,omitempty"`

	// provider configuration
	ProviderConfiguration *ProviderConfiguration `json:"providerConfiguration,omitempty"`
}

// Validate validates this global settings
func (m *GlobalSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCleanupOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultQuota(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineDeploymentOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineDeploymentVMResourceQuota(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMlaOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpaOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GlobalSettings) validateCleanupOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.CleanupOptions) { // not required
		return nil
	}

	if m.CleanupOptions != nil {
		if err := m.CleanupOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cleanupOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cleanupOptions")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalSettings) validateCustomLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomLinks) { // not required
		return nil
	}

	if err := m.CustomLinks.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("customLinks")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("customLinks")
		}
		return err
	}

	return nil
}

func (m *GlobalSettings) validateDefaultQuota(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultQuota) { // not required
		return nil
	}

	if m.DefaultQuota != nil {
		if err := m.DefaultQuota.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultQuota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultQuota")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalSettings) validateMachineDeploymentOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.MachineDeploymentOptions) { // not required
		return nil
	}

	if m.MachineDeploymentOptions != nil {
		if err := m.MachineDeploymentOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machineDeploymentOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machineDeploymentOptions")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalSettings) validateMachineDeploymentVMResourceQuota(formats strfmt.Registry) error {
	if swag.IsZero(m.MachineDeploymentVMResourceQuota) { // not required
		return nil
	}

	if m.MachineDeploymentVMResourceQuota != nil {
		if err := m.MachineDeploymentVMResourceQuota.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machineDeploymentVMResourceQuota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machineDeploymentVMResourceQuota")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalSettings) validateMlaOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.MlaOptions) { // not required
		return nil
	}

	if m.MlaOptions != nil {
		if err := m.MlaOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mlaOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mlaOptions")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalSettings) validateNotifications(formats strfmt.Registry) error {
	if swag.IsZero(m.Notifications) { // not required
		return nil
	}

	if m.Notifications != nil {
		if err := m.Notifications.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notifications")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalSettings) validateOpaOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.OpaOptions) { // not required
		return nil
	}

	if m.OpaOptions != nil {
		if err := m.OpaOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("opaOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("opaOptions")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalSettings) validateProviderConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.ProviderConfiguration) { // not required
		return nil
	}

	if m.ProviderConfiguration != nil {
		if err := m.ProviderConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("providerConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("providerConfiguration")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this global settings based on the context it is used
func (m *GlobalSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCleanupOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultQuota(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMachineDeploymentOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMachineDeploymentVMResourceQuota(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMlaOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotifications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpaOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProviderConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GlobalSettings) contextValidateCleanupOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.CleanupOptions != nil {
		if err := m.CleanupOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cleanupOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cleanupOptions")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalSettings) contextValidateCustomLinks(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CustomLinks.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("customLinks")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("customLinks")
		}
		return err
	}

	return nil
}

func (m *GlobalSettings) contextValidateDefaultQuota(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultQuota != nil {
		if err := m.DefaultQuota.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultQuota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultQuota")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalSettings) contextValidateMachineDeploymentOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.MachineDeploymentOptions != nil {
		if err := m.MachineDeploymentOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machineDeploymentOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machineDeploymentOptions")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalSettings) contextValidateMachineDeploymentVMResourceQuota(ctx context.Context, formats strfmt.Registry) error {

	if m.MachineDeploymentVMResourceQuota != nil {
		if err := m.MachineDeploymentVMResourceQuota.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machineDeploymentVMResourceQuota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machineDeploymentVMResourceQuota")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalSettings) contextValidateMlaOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.MlaOptions != nil {
		if err := m.MlaOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mlaOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mlaOptions")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalSettings) contextValidateNotifications(ctx context.Context, formats strfmt.Registry) error {

	if m.Notifications != nil {
		if err := m.Notifications.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notifications")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalSettings) contextValidateOpaOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.OpaOptions != nil {
		if err := m.OpaOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("opaOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("opaOptions")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalSettings) contextValidateProviderConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.ProviderConfiguration != nil {
		if err := m.ProviderConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("providerConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("providerConfiguration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GlobalSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GlobalSettings) UnmarshalBinary(b []byte) error {
	var res GlobalSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
