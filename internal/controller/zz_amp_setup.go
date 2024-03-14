package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	alertmanagerdefinition "github.com/upbound/provider-aws/internal/controller/amp/alertmanagerdefinition"
	rulegroupnamespace "github.com/upbound/provider-aws/internal/controller/amp/rulegroupnamespace"
	workspace "github.com/upbound/provider-aws/internal/controller/amp/workspace"
)

// Setup_amp creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_amp(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alertmanagerdefinition.Setup,
		rulegroupnamespace.Setup,
		workspace.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
