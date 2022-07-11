// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/client/addon"
	"github.com/kubermatic/go-kubermatic/client/admin"
	"github.com/kubermatic/go-kubermatic/client/aks"
	"github.com/kubermatic/go-kubermatic/client/alibaba"
	"github.com/kubermatic/go-kubermatic/client/allowedregistries"
	"github.com/kubermatic/go-kubermatic/client/allowedregistry"
	"github.com/kubermatic/go-kubermatic/client/anexia"
	"github.com/kubermatic/go-kubermatic/client/applications"
	"github.com/kubermatic/go-kubermatic/client/aws"
	"github.com/kubermatic/go-kubermatic/client/azure"
	"github.com/kubermatic/go-kubermatic/client/backupcredentials"
	"github.com/kubermatic/go-kubermatic/client/cniversion"
	"github.com/kubermatic/go-kubermatic/client/constraint"
	"github.com/kubermatic/go-kubermatic/client/constraints"
	"github.com/kubermatic/go-kubermatic/client/constrainttemplates"
	"github.com/kubermatic/go-kubermatic/client/credentials"
	"github.com/kubermatic/go-kubermatic/client/datacenter"
	"github.com/kubermatic/go-kubermatic/client/digitalocean"
	"github.com/kubermatic/go-kubermatic/client/eks"
	"github.com/kubermatic/go-kubermatic/client/etcdbackupconfig"
	"github.com/kubermatic/go-kubermatic/client/etcdrestore"
	"github.com/kubermatic/go-kubermatic/client/gcp"
	"github.com/kubermatic/go-kubermatic/client/get"
	"github.com/kubermatic/go-kubermatic/client/gke"
	"github.com/kubermatic/go-kubermatic/client/hetzner"
	"github.com/kubermatic/go-kubermatic/client/ipampool"
	"github.com/kubermatic/go-kubermatic/client/kubevirt"
	"github.com/kubermatic/go-kubermatic/client/metering"
	"github.com/kubermatic/go-kubermatic/client/metric"
	"github.com/kubermatic/go-kubermatic/client/mlaadminsetting"
	"github.com/kubermatic/go-kubermatic/client/networkdefaults"
	"github.com/kubermatic/go-kubermatic/client/nutanix"
	"github.com/kubermatic/go-kubermatic/client/openstack"
	"github.com/kubermatic/go-kubermatic/client/operations"
	"github.com/kubermatic/go-kubermatic/client/packet"
	"github.com/kubermatic/go-kubermatic/client/preset"
	"github.com/kubermatic/go-kubermatic/client/project"
	"github.com/kubermatic/go-kubermatic/client/resource_quota"
	"github.com/kubermatic/go-kubermatic/client/rulegroup"
	"github.com/kubermatic/go-kubermatic/client/seed"
	"github.com/kubermatic/go-kubermatic/client/serviceaccounts"
	"github.com/kubermatic/go-kubermatic/client/settings"
	"github.com/kubermatic/go-kubermatic/client/tokens"
	"github.com/kubermatic/go-kubermatic/client/user"
	"github.com/kubermatic/go-kubermatic/client/users"
	"github.com/kubermatic/go-kubermatic/client/version"
	"github.com/kubermatic/go-kubermatic/client/versions"
	"github.com/kubermatic/go-kubermatic/client/vmwareclouddirector"
	"github.com/kubermatic/go-kubermatic/client/vsphere"
)

// Default kubermatic kubernetes platform API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new kubermatic kubernetes platform API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *KubermaticKubernetesPlatformAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new kubermatic kubernetes platform API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *KubermaticKubernetesPlatformAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new kubermatic kubernetes platform API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *KubermaticKubernetesPlatformAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(KubermaticKubernetesPlatformAPI)
	cli.Transport = transport
	cli.Addon = addon.New(transport, formats)
	cli.Admin = admin.New(transport, formats)
	cli.Aks = aks.New(transport, formats)
	cli.Alibaba = alibaba.New(transport, formats)
	cli.Allowedregistries = allowedregistries.New(transport, formats)
	cli.Allowedregistry = allowedregistry.New(transport, formats)
	cli.Anexia = anexia.New(transport, formats)
	cli.Applications = applications.New(transport, formats)
	cli.Aws = aws.New(transport, formats)
	cli.Azure = azure.New(transport, formats)
	cli.Backupcredentials = backupcredentials.New(transport, formats)
	cli.Cniversion = cniversion.New(transport, formats)
	cli.Constraint = constraint.New(transport, formats)
	cli.Constraints = constraints.New(transport, formats)
	cli.Constrainttemplates = constrainttemplates.New(transport, formats)
	cli.Credentials = credentials.New(transport, formats)
	cli.Datacenter = datacenter.New(transport, formats)
	cli.Digitalocean = digitalocean.New(transport, formats)
	cli.Eks = eks.New(transport, formats)
	cli.Etcdbackupconfig = etcdbackupconfig.New(transport, formats)
	cli.Etcdrestore = etcdrestore.New(transport, formats)
	cli.Gcp = gcp.New(transport, formats)
	cli.Get = get.New(transport, formats)
	cli.Gke = gke.New(transport, formats)
	cli.Hetzner = hetzner.New(transport, formats)
	cli.Ipampool = ipampool.New(transport, formats)
	cli.Kubevirt = kubevirt.New(transport, formats)
	cli.Metering = metering.New(transport, formats)
	cli.Metric = metric.New(transport, formats)
	cli.Mlaadminsetting = mlaadminsetting.New(transport, formats)
	cli.Networkdefaults = networkdefaults.New(transport, formats)
	cli.Nutanix = nutanix.New(transport, formats)
	cli.Openstack = openstack.New(transport, formats)
	cli.Operations = operations.New(transport, formats)
	cli.Packet = packet.New(transport, formats)
	cli.Preset = preset.New(transport, formats)
	cli.Project = project.New(transport, formats)
	cli.ResourceQuota = resource_quota.New(transport, formats)
	cli.Rulegroup = rulegroup.New(transport, formats)
	cli.Seed = seed.New(transport, formats)
	cli.Serviceaccounts = serviceaccounts.New(transport, formats)
	cli.Settings = settings.New(transport, formats)
	cli.Tokens = tokens.New(transport, formats)
	cli.User = user.New(transport, formats)
	cli.Users = users.New(transport, formats)
	cli.Version = version.New(transport, formats)
	cli.Versions = versions.New(transport, formats)
	cli.Vmwareclouddirector = vmwareclouddirector.New(transport, formats)
	cli.Vsphere = vsphere.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// KubermaticKubernetesPlatformAPI is a client for kubermatic kubernetes platform API
type KubermaticKubernetesPlatformAPI struct {
	Addon addon.ClientService

	Admin admin.ClientService

	Aks aks.ClientService

	Alibaba alibaba.ClientService

	Allowedregistries allowedregistries.ClientService

	Allowedregistry allowedregistry.ClientService

	Anexia anexia.ClientService

	Applications applications.ClientService

	Aws aws.ClientService

	Azure azure.ClientService

	Backupcredentials backupcredentials.ClientService

	Cniversion cniversion.ClientService

	Constraint constraint.ClientService

	Constraints constraints.ClientService

	Constrainttemplates constrainttemplates.ClientService

	Credentials credentials.ClientService

	Datacenter datacenter.ClientService

	Digitalocean digitalocean.ClientService

	Eks eks.ClientService

	Etcdbackupconfig etcdbackupconfig.ClientService

	Etcdrestore etcdrestore.ClientService

	Gcp gcp.ClientService

	Get get.ClientService

	Gke gke.ClientService

	Hetzner hetzner.ClientService

	Ipampool ipampool.ClientService

	Kubevirt kubevirt.ClientService

	Metering metering.ClientService

	Metric metric.ClientService

	Mlaadminsetting mlaadminsetting.ClientService

	Networkdefaults networkdefaults.ClientService

	Nutanix nutanix.ClientService

	Openstack openstack.ClientService

	Operations operations.ClientService

	Packet packet.ClientService

	Preset preset.ClientService

	Project project.ClientService

	ResourceQuota resource_quota.ClientService

	Rulegroup rulegroup.ClientService

	Seed seed.ClientService

	Serviceaccounts serviceaccounts.ClientService

	Settings settings.ClientService

	Tokens tokens.ClientService

	User user.ClientService

	Users users.ClientService

	Version version.ClientService

	Versions versions.ClientService

	Vmwareclouddirector vmwareclouddirector.ClientService

	Vsphere vsphere.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *KubermaticKubernetesPlatformAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Addon.SetTransport(transport)
	c.Admin.SetTransport(transport)
	c.Aks.SetTransport(transport)
	c.Alibaba.SetTransport(transport)
	c.Allowedregistries.SetTransport(transport)
	c.Allowedregistry.SetTransport(transport)
	c.Anexia.SetTransport(transport)
	c.Applications.SetTransport(transport)
	c.Aws.SetTransport(transport)
	c.Azure.SetTransport(transport)
	c.Backupcredentials.SetTransport(transport)
	c.Cniversion.SetTransport(transport)
	c.Constraint.SetTransport(transport)
	c.Constraints.SetTransport(transport)
	c.Constrainttemplates.SetTransport(transport)
	c.Credentials.SetTransport(transport)
	c.Datacenter.SetTransport(transport)
	c.Digitalocean.SetTransport(transport)
	c.Eks.SetTransport(transport)
	c.Etcdbackupconfig.SetTransport(transport)
	c.Etcdrestore.SetTransport(transport)
	c.Gcp.SetTransport(transport)
	c.Get.SetTransport(transport)
	c.Gke.SetTransport(transport)
	c.Hetzner.SetTransport(transport)
	c.Ipampool.SetTransport(transport)
	c.Kubevirt.SetTransport(transport)
	c.Metering.SetTransport(transport)
	c.Metric.SetTransport(transport)
	c.Mlaadminsetting.SetTransport(transport)
	c.Networkdefaults.SetTransport(transport)
	c.Nutanix.SetTransport(transport)
	c.Openstack.SetTransport(transport)
	c.Operations.SetTransport(transport)
	c.Packet.SetTransport(transport)
	c.Preset.SetTransport(transport)
	c.Project.SetTransport(transport)
	c.ResourceQuota.SetTransport(transport)
	c.Rulegroup.SetTransport(transport)
	c.Seed.SetTransport(transport)
	c.Serviceaccounts.SetTransport(transport)
	c.Settings.SetTransport(transport)
	c.Tokens.SetTransport(transport)
	c.User.SetTransport(transport)
	c.Users.SetTransport(transport)
	c.Version.SetTransport(transport)
	c.Versions.SetTransport(transport)
	c.Vmwareclouddirector.SetTransport(transport)
	c.Vsphere.SetTransport(transport)
}
