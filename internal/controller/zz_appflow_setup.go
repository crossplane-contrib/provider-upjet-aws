package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	flow "github.com/upbound/provider-aws/internal/controller/appflow/flow"
)

// Setup_appflow creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_appflow(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		flow.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
