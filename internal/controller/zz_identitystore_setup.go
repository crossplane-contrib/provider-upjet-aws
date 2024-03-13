package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	group "github.com/upbound/provider-aws/internal/controller/identitystore/group"
	groupmembership "github.com/upbound/provider-aws/internal/controller/identitystore/groupmembership"
	user "github.com/upbound/provider-aws/internal/controller/identitystore/user"
)

// Setup_identitystore creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_identitystore(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		group.Setup,
		groupmembership.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
