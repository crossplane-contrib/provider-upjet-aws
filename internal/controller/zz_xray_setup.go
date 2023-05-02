/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	encryptionconfig "github.com/upbound/provider-aws/internal/controller/xray/encryptionconfig"
	group "github.com/upbound/provider-aws/internal/controller/xray/group"
	samplingrule "github.com/upbound/provider-aws/internal/controller/xray/samplingrule"
)

// Setup_xray creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_xray(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		encryptionconfig.Setup,
		group.Setup,
		samplingrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
