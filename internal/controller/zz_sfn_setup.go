package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	activity "github.com/upbound/provider-aws/internal/controller/sfn/activity"
	statemachine "github.com/upbound/provider-aws/internal/controller/sfn/statemachine"
)

// Setup_sfn creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_sfn(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		activity.Setup,
		statemachine.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
