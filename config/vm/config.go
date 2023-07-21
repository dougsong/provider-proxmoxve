package vm

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("proxmox_vm_qemu", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "VM"
	})
}
