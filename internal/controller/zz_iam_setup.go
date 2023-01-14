/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	accesskey "github.com/upbound/provider-aws/internal/controller/iam/accesskey"
	openidconnectprovider "github.com/upbound/provider-aws/internal/controller/iam/openidconnectprovider"
	policy "github.com/upbound/provider-aws/internal/controller/iam/policy"
	role "github.com/upbound/provider-aws/internal/controller/iam/role"
	rolepolicyattachment "github.com/upbound/provider-aws/internal/controller/iam/rolepolicyattachment"
	user "github.com/upbound/provider-aws/internal/controller/iam/user"
	userpolicyattachment "github.com/upbound/provider-aws/internal/controller/iam/userpolicyattachment"
)

// Setup_iam creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_iam(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesskey.Setup,
		openidconnectprovider.Setup,
		policy.Setup,
		role.Setup,
		rolepolicyattachment.Setup,
		user.Setup,
		userpolicyattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
