package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	assessmenttarget "github.com/upbound/provider-aws/internal/controller/inspector/assessmenttarget"
	assessmenttemplate "github.com/upbound/provider-aws/internal/controller/inspector/assessmenttemplate"
	resourcegroup "github.com/upbound/provider-aws/internal/controller/inspector/resourcegroup"
)

// Setup_inspector creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_inspector(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		assessmenttarget.Setup,
		assessmenttemplate.Setup,
		resourcegroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
