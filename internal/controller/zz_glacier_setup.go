package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	vault "github.com/upbound/provider-aws/internal/controller/glacier/vault"
	vaultlock "github.com/upbound/provider-aws/internal/controller/glacier/vaultlock"
)

// Setup_glacier creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_glacier(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		vault.Setup,
		vaultlock.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
