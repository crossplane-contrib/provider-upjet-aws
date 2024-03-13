package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	domain "github.com/upbound/provider-aws/internal/controller/simpledb/domain"
)

// Setup_simpledb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_simpledb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		domain.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
