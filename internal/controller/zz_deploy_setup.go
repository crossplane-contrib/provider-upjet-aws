package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	app "github.com/upbound/provider-aws/internal/controller/deploy/app"
	deploymentconfig "github.com/upbound/provider-aws/internal/controller/deploy/deploymentconfig"
	deploymentgroup "github.com/upbound/provider-aws/internal/controller/deploy/deploymentgroup"
)

// Setup_deploy creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_deploy(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		app.Setup,
		deploymentconfig.Setup,
		deploymentgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
