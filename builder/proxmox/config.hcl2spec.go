// Code generated by "mapstructure-to-hcl2 -type Config,nicConfig,diskConfig"; DO NOT EDIT.
package proxmox

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType         *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug               *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce               *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError             *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars            map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars       []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	HTTPDir                   *string           `mapstructure:"http_directory" cty:"http_directory"`
	HTTPPortMin               *int              `mapstructure:"http_port_min" cty:"http_port_min"`
	HTTPPortMax               *int              `mapstructure:"http_port_max" cty:"http_port_max"`
	ISOChecksum               *string           `mapstructure:"iso_checksum" required:"true" cty:"iso_checksum"`
	ISOChecksumURL            *string           `mapstructure:"iso_checksum_url" cty:"iso_checksum_url"`
	ISOChecksumType           *string           `mapstructure:"iso_checksum_type" cty:"iso_checksum_type"`
	RawSingleISOUrl           *string           `mapstructure:"iso_url" required:"true" cty:"iso_url"`
	ISOUrls                   []string          `mapstructure:"iso_urls" cty:"iso_urls"`
	TargetPath                *string           `mapstructure:"iso_target_path" cty:"iso_target_path"`
	TargetExtension           *string           `mapstructure:"iso_target_extension" cty:"iso_target_extension"`
	BootGroupInterval         *string           `mapstructure:"boot_keygroup_interval" cty:"boot_keygroup_interval"`
	BootWait                  *string           `mapstructure:"boot_wait" cty:"boot_wait"`
	BootCommand               []string          `mapstructure:"boot_command" cty:"boot_command"`
	BootKeyInterval           *string           `mapstructure:"boot_key_interval" cty:"boot_key_interval"`
	Type                      *string           `mapstructure:"communicator" cty:"communicator"`
	PauseBeforeConnect        *string           `mapstructure:"pause_before_connecting" cty:"pause_before_connecting"`
	SSHHost                   *string           `mapstructure:"ssh_host" cty:"ssh_host"`
	SSHPort                   *int              `mapstructure:"ssh_port" cty:"ssh_port"`
	SSHUsername               *string           `mapstructure:"ssh_username" cty:"ssh_username"`
	SSHPassword               *string           `mapstructure:"ssh_password" cty:"ssh_password"`
	SSHKeyPairName            *string           `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   *string           `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys    *bool             `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile         *string           `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file"`
	SSHPty                    *bool             `mapstructure:"ssh_pty" cty:"ssh_pty"`
	SSHTimeout                *string           `mapstructure:"ssh_timeout" cty:"ssh_timeout"`
	SSHAgentAuth              *bool             `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth"`
	SSHDisableAgentForwarding *bool             `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      *int              `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts"`
	SSHBastionHost            *string           `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host"`
	SSHBastionPort            *int              `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port"`
	SSHBastionAgentAuth       *bool             `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth"`
	SSHBastionUsername        *string           `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username"`
	SSHBastionPassword        *string           `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password"`
	SSHBastionPrivateKeyFile  *string           `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod     *string           `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method"`
	SSHProxyHost              *string           `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host"`
	SSHProxyPort              *int              `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port"`
	SSHProxyUsername          *string           `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username"`
	SSHProxyPassword          *string           `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password"`
	SSHKeepAliveInterval      *string           `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       *string           `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string          `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string          `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels"`
	SSHPublicKey              []byte            `mapstructure:"ssh_public_key" cty:"ssh_public_key"`
	SSHPrivateKey             []byte            `mapstructure:"ssh_private_key" cty:"ssh_private_key"`
	WinRMUser                 *string           `mapstructure:"winrm_username" cty:"winrm_username"`
	WinRMPassword             *string           `mapstructure:"winrm_password" cty:"winrm_password"`
	WinRMHost                 *string           `mapstructure:"winrm_host" cty:"winrm_host"`
	WinRMPort                 *int              `mapstructure:"winrm_port" cty:"winrm_port"`
	WinRMTimeout              *string           `mapstructure:"winrm_timeout" cty:"winrm_timeout"`
	WinRMUseSSL               *bool             `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl"`
	WinRMInsecure             *bool             `mapstructure:"winrm_insecure" cty:"winrm_insecure"`
	WinRMUseNTLM              *bool             `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm"`
	ProxmoxURLRaw             *string           `mapstructure:"proxmox_url" cty:"proxmox_url"`
	SkipCertValidation        *bool             `mapstructure:"insecure_skip_tls_verify" cty:"insecure_skip_tls_verify"`
	Username                  *string           `mapstructure:"username" cty:"username"`
	Password                  *string           `mapstructure:"password" cty:"password"`
	Node                      *string           `mapstructure:"node" cty:"node"`
	Pool                      *string           `mapstructure:"pool" cty:"pool"`
	VMName                    *string           `mapstructure:"vm_name" cty:"vm_name"`
	VMID                      *int              `mapstructure:"vm_id" cty:"vm_id"`
	Memory                    *int              `mapstructure:"memory" cty:"memory"`
	Cores                     *int              `mapstructure:"cores" cty:"cores"`
	CPUType                   *string           `mapstructure:"cpu_type" cty:"cpu_type"`
	Sockets                   *int              `mapstructure:"sockets" cty:"sockets"`
	OS                        *string           `mapstructure:"os" cty:"os"`
	NICs                      []FlatnicConfig   `mapstructure:"network_adapters" cty:"network_adapters"`
	Disks                     []FlatdiskConfig  `mapstructure:"disks" cty:"disks"`
	ISOFile                   *string           `mapstructure:"iso_file" cty:"iso_file"`
	ISOStoragePool            *string           `mapstructure:"iso_storage_pool" cty:"iso_storage_pool"`
	Agent                     *bool             `mapstructure:"qemu_agent" cty:"qemu_agent"`
	SCSIController            *string           `mapstructure:"scsi_controller" cty:"scsi_controller"`
	TemplateName              *string           `mapstructure:"template_name" cty:"template_name"`
	TemplateDescription       *string           `mapstructure:"template_description" cty:"template_description"`
	UnmountISO                *bool             `mapstructure:"unmount_iso" cty:"unmount_iso"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":            &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":          &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                 &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                 &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":              &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":        &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables":   &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"http_directory":               &hcldec.AttrSpec{Name: "http_directory", Type: cty.String, Required: false},
		"http_port_min":                &hcldec.AttrSpec{Name: "http_port_min", Type: cty.Number, Required: false},
		"http_port_max":                &hcldec.AttrSpec{Name: "http_port_max", Type: cty.Number, Required: false},
		"iso_checksum":                 &hcldec.AttrSpec{Name: "iso_checksum", Type: cty.String, Required: false},
		"iso_checksum_url":             &hcldec.AttrSpec{Name: "iso_checksum_url", Type: cty.String, Required: false},
		"iso_checksum_type":            &hcldec.AttrSpec{Name: "iso_checksum_type", Type: cty.String, Required: false},
		"iso_url":                      &hcldec.AttrSpec{Name: "iso_url", Type: cty.String, Required: false},
		"iso_urls":                     &hcldec.AttrSpec{Name: "iso_urls", Type: cty.List(cty.String), Required: false},
		"iso_target_path":              &hcldec.AttrSpec{Name: "iso_target_path", Type: cty.String, Required: false},
		"iso_target_extension":         &hcldec.AttrSpec{Name: "iso_target_extension", Type: cty.String, Required: false},
		"boot_keygroup_interval":       &hcldec.AttrSpec{Name: "boot_keygroup_interval", Type: cty.String, Required: false},
		"boot_wait":                    &hcldec.AttrSpec{Name: "boot_wait", Type: cty.String, Required: false},
		"boot_command":                 &hcldec.AttrSpec{Name: "boot_command", Type: cty.List(cty.String), Required: false},
		"boot_key_interval":            &hcldec.AttrSpec{Name: "boot_key_interval", Type: cty.String, Required: false},
		"communicator":                 &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":      &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                     &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                     &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                 &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                 &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":             &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":      &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"ssh_clear_authorized_keys":    &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_private_key_file":         &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_pty":                      &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                  &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":               &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding": &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":       &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":             &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":             &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":       &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":         &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":         &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_private_key_file": &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":     &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":               &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":               &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":           &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":           &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":      &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":       &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":           &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":            &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":               &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":              &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":               &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":               &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                   &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_port":                   &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":               &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":               &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"proxmox_url":                  &hcldec.AttrSpec{Name: "proxmox_url", Type: cty.String, Required: false},
		"insecure_skip_tls_verify":     &hcldec.AttrSpec{Name: "insecure_skip_tls_verify", Type: cty.Bool, Required: false},
		"username":                     &hcldec.AttrSpec{Name: "username", Type: cty.String, Required: false},
		"password":                     &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"node":                         &hcldec.AttrSpec{Name: "node", Type: cty.String, Required: false},
		"pool":                         &hcldec.AttrSpec{Name: "pool", Type: cty.String, Required: false},
		"vm_name":                      &hcldec.AttrSpec{Name: "vm_name", Type: cty.String, Required: false},
		"vm_id":                        &hcldec.AttrSpec{Name: "vm_id", Type: cty.Number, Required: false},
		"memory":                       &hcldec.AttrSpec{Name: "memory", Type: cty.Number, Required: false},
		"cores":                        &hcldec.AttrSpec{Name: "cores", Type: cty.Number, Required: false},
		"cpu_type":                     &hcldec.AttrSpec{Name: "cpu_type", Type: cty.String, Required: false},
		"sockets":                      &hcldec.AttrSpec{Name: "sockets", Type: cty.Number, Required: false},
		"os":                           &hcldec.AttrSpec{Name: "os", Type: cty.String, Required: false},
		"network_adapters":             &hcldec.BlockListSpec{TypeName: "network_adapters", Nested: hcldec.ObjectSpec((*FlatnicConfig)(nil).HCL2Spec())},
		"disks":                        &hcldec.BlockListSpec{TypeName: "disks", Nested: hcldec.ObjectSpec((*FlatdiskConfig)(nil).HCL2Spec())},
		"iso_file":                     &hcldec.AttrSpec{Name: "iso_file", Type: cty.String, Required: false},
		"iso_storage_pool":             &hcldec.AttrSpec{Name: "iso_storage_pool", Type: cty.String, Required: false},
		"qemu_agent":                   &hcldec.AttrSpec{Name: "qemu_agent", Type: cty.Bool, Required: false},
		"scsi_controller":              &hcldec.AttrSpec{Name: "scsi_controller", Type: cty.String, Required: false},
		"template_name":                &hcldec.AttrSpec{Name: "template_name", Type: cty.String, Required: false},
		"template_description":         &hcldec.AttrSpec{Name: "template_description", Type: cty.String, Required: false},
		"unmount_iso":                  &hcldec.AttrSpec{Name: "unmount_iso", Type: cty.Bool, Required: false},
	}
	return s
}

// FlatdiskConfig is an auto-generated flat version of diskConfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatdiskConfig struct {
	Type            *string `mapstructure:"type" cty:"type"`
	StoragePool     *string `mapstructure:"storage_pool" cty:"storage_pool"`
	StoragePoolType *string `mapstructure:"storage_pool_type" cty:"storage_pool_type"`
	Size            *string `mapstructure:"disk_size" cty:"disk_size"`
	CacheMode       *string `mapstructure:"cache_mode" cty:"cache_mode"`
	DiskFormat      *string `mapstructure:"format" cty:"format"`
}

// FlatMapstructure returns a new FlatdiskConfig.
// FlatdiskConfig is an auto-generated flat version of diskConfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*diskConfig) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatdiskConfig)
}

// HCL2Spec returns the hcl spec of a diskConfig.
// This spec is used by HCL to read the fields of diskConfig.
// The decoded values from this spec will then be applied to a FlatdiskConfig.
func (*FlatdiskConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"type":              &hcldec.AttrSpec{Name: "type", Type: cty.String, Required: false},
		"storage_pool":      &hcldec.AttrSpec{Name: "storage_pool", Type: cty.String, Required: false},
		"storage_pool_type": &hcldec.AttrSpec{Name: "storage_pool_type", Type: cty.String, Required: false},
		"disk_size":         &hcldec.AttrSpec{Name: "disk_size", Type: cty.String, Required: false},
		"cache_mode":        &hcldec.AttrSpec{Name: "cache_mode", Type: cty.String, Required: false},
		"format":            &hcldec.AttrSpec{Name: "format", Type: cty.String, Required: false},
	}
	return s
}

// FlatnicConfig is an auto-generated flat version of nicConfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatnicConfig struct {
	Model      *string `mapstructure:"model" cty:"model"`
	MACAddress *string `mapstructure:"mac_address" cty:"mac_address"`
	Bridge     *string `mapstructure:"bridge" cty:"bridge"`
	VLANTag    *string `mapstructure:"vlan_tag" cty:"vlan_tag"`
}

// FlatMapstructure returns a new FlatnicConfig.
// FlatnicConfig is an auto-generated flat version of nicConfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*nicConfig) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatnicConfig)
}

// HCL2Spec returns the hcl spec of a nicConfig.
// This spec is used by HCL to read the fields of nicConfig.
// The decoded values from this spec will then be applied to a FlatnicConfig.
func (*FlatnicConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"model":       &hcldec.AttrSpec{Name: "model", Type: cty.String, Required: false},
		"mac_address": &hcldec.AttrSpec{Name: "mac_address", Type: cty.String, Required: false},
		"bridge":      &hcldec.AttrSpec{Name: "bridge", Type: cty.String, Required: false},
		"vlan_tag":    &hcldec.AttrSpec{Name: "vlan_tag", Type: cty.String, Required: false},
	}
	return s
}
