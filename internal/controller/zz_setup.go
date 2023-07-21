/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	lxc "github.com/dougsong/provider-proxmoxve/internal/controller/lxc/lxc"
	providerconfig "github.com/dougsong/provider-proxmoxve/internal/controller/providerconfig"
	qemu "github.com/dougsong/provider-proxmoxve/internal/controller/vm/qemu"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		lxc.Setup,
		providerconfig.Setup,
		qemu.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
