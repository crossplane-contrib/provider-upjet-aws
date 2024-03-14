package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	securityconfiguration "github.com/upbound/provider-aws/internal/controller/emr/securityconfiguration"
)

// Setup_emr creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_emr(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		securityconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
