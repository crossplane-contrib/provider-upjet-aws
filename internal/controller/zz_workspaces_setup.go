package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	directory "github.com/upbound/provider-aws/internal/controller/workspaces/directory"
	ipgroup "github.com/upbound/provider-aws/internal/controller/workspaces/ipgroup"
)

// Setup_workspaces creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_workspaces(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		directory.Setup,
		ipgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
