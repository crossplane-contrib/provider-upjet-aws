package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	reportdefinition "github.com/upbound/provider-aws/internal/controller/cur/reportdefinition"
)

// Setup_cur creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cur(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		reportdefinition.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
