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

// DatacenterSpec DatacenterSpec specifies the data for a datacenter.
//
// swagger:model DatacenterSpec
type DatacenterSpec struct {

	// Optional: Country of the seed as ISO-3166 two-letter code, e.g. DE or UK.
	// It is used for informational purposes.
	Country string `json:"country,omitempty"`

	// Optional: DisableCSIDriver disables the installation of CSI driver on every clusters within the DC
	// If true it can't be over-written in the cluster configuration
	DisableCSIDriver bool `json:"disableCsiDriver,omitempty"`

	// EnforceAuditLogging enforces audit logging on every cluster within the DC,
	// ignoring cluster-specific settings.
	EnforceAuditLogging bool `json:"enforceAuditLogging,omitempty"`

	// EnforcePodSecurityPolicy enforces pod security policy plugin on every clusters within the DC,
	// ignoring cluster-specific settings
	EnforcePodSecurityPolicy bool `json:"enforcePodSecurityPolicy,omitempty"`

	// IPv6Enabled is a flag to indicate if the ipv6 is enabled for the datacenter.
	IPV6Enabled bool `json:"ipv6Enabled,omitempty"`

	// Optional: Detailed location of the cluster, like "Hamburg" or "Datacenter 7".
	// It is used for informational purposes.
	Location string `json:"location,omitempty"`

	// Name of the datacenter provider. Extracted based on which provider is defined in the spec.
	// It is used for informational purposes.
	Provider string `json:"provider,omitempty"`

	// required emails
	RequiredEmails []string `json:"requiredEmails"`

	// Name of the seed this datacenter belongs to.
	Seed string `json:"seed,omitempty"`

	// alibaba
	Alibaba *DatacenterSpecAlibaba `json:"alibaba,omitempty"`

	// anexia
	Anexia *DatacenterSpecAnexia `json:"anexia,omitempty"`

	// aws
	Aws *DatacenterSpecAWS `json:"aws,omitempty"`

	// azure
	Azure *DatacenterSpecAzure `json:"azure,omitempty"`

	// baremetal
	Baremetal *DatacenterSpecBaremetal `json:"baremetal,omitempty"`

	// bringyourown
	Bringyourown DatacenterSpecBringYourOwn `json:"bringyourown,omitempty"`

	// digitalocean
	Digitalocean *DatacenterSpecDigitalocean `json:"digitalocean,omitempty"`

	// enforced audit webhook settings
	EnforcedAuditWebhookSettings *AuditWebhookBackendSettings `json:"enforcedAuditWebhookSettings,omitempty"`

	// fake
	Fake *DatacenterSpecFake `json:"fake,omitempty"`

	// gcp
	Gcp *DatacenterSpecGCP `json:"gcp,omitempty"`

	// hetzner
	Hetzner *DatacenterSpecHetzner `json:"hetzner,omitempty"`

	// kubelb
	Kubelb *KubeLBDatacenterSettings `json:"kubelb,omitempty"`

	// kubevirt
	Kubevirt *DatacenterSpecKubevirt `json:"kubevirt,omitempty"`

	// machine flavor filter
	MachineFlavorFilter *MachineFlavorFilter `json:"machineFlavorFilter,omitempty"`

	// node
	Node *NodeSettings `json:"node,omitempty"`

	// nutanix
	Nutanix *DatacenterSpecNutanix `json:"nutanix,omitempty"`

	// openstack
	Openstack *DatacenterSpecOpenstack `json:"openstack,omitempty"`

	// operating system profiles
	OperatingSystemProfiles OperatingSystemProfileList `json:"operatingSystemProfiles,omitempty"`

	// packet
	Packet *DatacenterSpecPacket `json:"packet,omitempty"`

	// vmwareclouddirector
	Vmwareclouddirector *DatacenterSpecVMwareCloudDirector `json:"vmwareclouddirector,omitempty"`

	// vsphere
	Vsphere *DatacenterSpecVSphere `json:"vsphere,omitempty"`
}

// Validate validates this datacenter spec
func (m *DatacenterSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlibaba(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnexia(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaremetal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDigitalocean(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnforcedAuditWebhookSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFake(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHetzner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubelb(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubevirt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineFlavorFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNutanix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenstack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatingSystemProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePacket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVmwareclouddirector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVsphere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatacenterSpec) validateAlibaba(formats strfmt.Registry) error {
	if swag.IsZero(m.Alibaba) { // not required
		return nil
	}

	if m.Alibaba != nil {
		if err := m.Alibaba.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alibaba")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alibaba")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validateAnexia(formats strfmt.Registry) error {
	if swag.IsZero(m.Anexia) { // not required
		return nil
	}

	if m.Anexia != nil {
		if err := m.Anexia.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("anexia")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("anexia")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validateAws(formats strfmt.Registry) error {
	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validateAzure(formats strfmt.Registry) error {
	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validateBaremetal(formats strfmt.Registry) error {
	if swag.IsZero(m.Baremetal) { // not required
		return nil
	}

	if m.Baremetal != nil {
		if err := m.Baremetal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baremetal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("baremetal")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validateDigitalocean(formats strfmt.Registry) error {
	if swag.IsZero(m.Digitalocean) { // not required
		return nil
	}

	if m.Digitalocean != nil {
		if err := m.Digitalocean.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("digitalocean")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("digitalocean")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validateEnforcedAuditWebhookSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.EnforcedAuditWebhookSettings) { // not required
		return nil
	}

	if m.EnforcedAuditWebhookSettings != nil {
		if err := m.EnforcedAuditWebhookSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enforcedAuditWebhookSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enforcedAuditWebhookSettings")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validateFake(formats strfmt.Registry) error {
	if swag.IsZero(m.Fake) { // not required
		return nil
	}

	if m.Fake != nil {
		if err := m.Fake.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fake")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fake")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validateGcp(formats strfmt.Registry) error {
	if swag.IsZero(m.Gcp) { // not required
		return nil
	}

	if m.Gcp != nil {
		if err := m.Gcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validateHetzner(formats strfmt.Registry) error {
	if swag.IsZero(m.Hetzner) { // not required
		return nil
	}

	if m.Hetzner != nil {
		if err := m.Hetzner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hetzner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hetzner")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validateKubelb(formats strfmt.Registry) error {
	if swag.IsZero(m.Kubelb) { // not required
		return nil
	}

	if m.Kubelb != nil {
		if err := m.Kubelb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubelb")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubelb")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validateKubevirt(formats strfmt.Registry) error {
	if swag.IsZero(m.Kubevirt) { // not required
		return nil
	}

	if m.Kubevirt != nil {
		if err := m.Kubevirt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubevirt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubevirt")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validateMachineFlavorFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.MachineFlavorFilter) { // not required
		return nil
	}

	if m.MachineFlavorFilter != nil {
		if err := m.MachineFlavorFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machineFlavorFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machineFlavorFilter")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validateNutanix(formats strfmt.Registry) error {
	if swag.IsZero(m.Nutanix) { // not required
		return nil
	}

	if m.Nutanix != nil {
		if err := m.Nutanix.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nutanix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nutanix")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validateOpenstack(formats strfmt.Registry) error {
	if swag.IsZero(m.Openstack) { // not required
		return nil
	}

	if m.Openstack != nil {
		if err := m.Openstack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openstack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("openstack")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validateOperatingSystemProfiles(formats strfmt.Registry) error {
	if swag.IsZero(m.OperatingSystemProfiles) { // not required
		return nil
	}

	if m.OperatingSystemProfiles != nil {
		if err := m.OperatingSystemProfiles.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operatingSystemProfiles")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operatingSystemProfiles")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validatePacket(formats strfmt.Registry) error {
	if swag.IsZero(m.Packet) { // not required
		return nil
	}

	if m.Packet != nil {
		if err := m.Packet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("packet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("packet")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validateVmwareclouddirector(formats strfmt.Registry) error {
	if swag.IsZero(m.Vmwareclouddirector) { // not required
		return nil
	}

	if m.Vmwareclouddirector != nil {
		if err := m.Vmwareclouddirector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmwareclouddirector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmwareclouddirector")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) validateVsphere(formats strfmt.Registry) error {
	if swag.IsZero(m.Vsphere) { // not required
		return nil
	}

	if m.Vsphere != nil {
		if err := m.Vsphere.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsphere")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vsphere")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this datacenter spec based on the context it is used
func (m *DatacenterSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlibaba(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAnexia(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAws(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBaremetal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDigitalocean(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnforcedAuditWebhookSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFake(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHetzner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKubelb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKubevirt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMachineFlavorFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNutanix(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpenstack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperatingSystemProfiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePacket(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVmwareclouddirector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVsphere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatacenterSpec) contextValidateAlibaba(ctx context.Context, formats strfmt.Registry) error {

	if m.Alibaba != nil {
		if err := m.Alibaba.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alibaba")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alibaba")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) contextValidateAnexia(ctx context.Context, formats strfmt.Registry) error {

	if m.Anexia != nil {
		if err := m.Anexia.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("anexia")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("anexia")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) contextValidateAws(ctx context.Context, formats strfmt.Registry) error {

	if m.Aws != nil {
		if err := m.Aws.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) contextValidateAzure(ctx context.Context, formats strfmt.Registry) error {

	if m.Azure != nil {
		if err := m.Azure.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) contextValidateBaremetal(ctx context.Context, formats strfmt.Registry) error {

	if m.Baremetal != nil {
		if err := m.Baremetal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baremetal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("baremetal")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) contextValidateDigitalocean(ctx context.Context, formats strfmt.Registry) error {

	if m.Digitalocean != nil {
		if err := m.Digitalocean.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("digitalocean")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("digitalocean")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) contextValidateEnforcedAuditWebhookSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.EnforcedAuditWebhookSettings != nil {
		if err := m.EnforcedAuditWebhookSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enforcedAuditWebhookSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enforcedAuditWebhookSettings")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) contextValidateFake(ctx context.Context, formats strfmt.Registry) error {

	if m.Fake != nil {
		if err := m.Fake.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fake")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fake")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) contextValidateGcp(ctx context.Context, formats strfmt.Registry) error {

	if m.Gcp != nil {
		if err := m.Gcp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) contextValidateHetzner(ctx context.Context, formats strfmt.Registry) error {

	if m.Hetzner != nil {
		if err := m.Hetzner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hetzner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hetzner")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) contextValidateKubelb(ctx context.Context, formats strfmt.Registry) error {

	if m.Kubelb != nil {
		if err := m.Kubelb.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubelb")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubelb")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) contextValidateKubevirt(ctx context.Context, formats strfmt.Registry) error {

	if m.Kubevirt != nil {
		if err := m.Kubevirt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubevirt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubevirt")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) contextValidateMachineFlavorFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.MachineFlavorFilter != nil {
		if err := m.MachineFlavorFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machineFlavorFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machineFlavorFilter")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {
		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) contextValidateNutanix(ctx context.Context, formats strfmt.Registry) error {

	if m.Nutanix != nil {
		if err := m.Nutanix.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nutanix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nutanix")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) contextValidateOpenstack(ctx context.Context, formats strfmt.Registry) error {

	if m.Openstack != nil {
		if err := m.Openstack.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openstack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("openstack")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) contextValidateOperatingSystemProfiles(ctx context.Context, formats strfmt.Registry) error {

	if err := m.OperatingSystemProfiles.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operatingSystemProfiles")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("operatingSystemProfiles")
		}
		return err
	}

	return nil
}

func (m *DatacenterSpec) contextValidatePacket(ctx context.Context, formats strfmt.Registry) error {

	if m.Packet != nil {
		if err := m.Packet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("packet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("packet")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) contextValidateVmwareclouddirector(ctx context.Context, formats strfmt.Registry) error {

	if m.Vmwareclouddirector != nil {
		if err := m.Vmwareclouddirector.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmwareclouddirector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmwareclouddirector")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpec) contextValidateVsphere(ctx context.Context, formats strfmt.Registry) error {

	if m.Vsphere != nil {
		if err := m.Vsphere.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsphere")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vsphere")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatacenterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatacenterSpec) UnmarshalBinary(b []byte) error {
	var res DatacenterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
