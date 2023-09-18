/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	identityprovider "github.com/upbound/provider-aws/internal/controller/cognitoidp/identityprovider"
	resourceserver "github.com/upbound/provider-aws/internal/controller/cognitoidp/resourceserver"
	riskconfiguration "github.com/upbound/provider-aws/internal/controller/cognitoidp/riskconfiguration"
	user "github.com/upbound/provider-aws/internal/controller/cognitoidp/user"
	usergroup "github.com/upbound/provider-aws/internal/controller/cognitoidp/usergroup"
	useringroup "github.com/upbound/provider-aws/internal/controller/cognitoidp/useringroup"
	userpool "github.com/upbound/provider-aws/internal/controller/cognitoidp/userpool"
	userpooldomain "github.com/upbound/provider-aws/internal/controller/cognitoidp/userpooldomain"
	userpooluicustomization "github.com/upbound/provider-aws/internal/controller/cognitoidp/userpooluicustomization"
)

// Setup_cognitoidp creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cognitoidp(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		identityprovider.Setup,
		resourceserver.Setup,
		riskconfiguration.Setup,
		user.Setup,
		usergroup.Setup,
		useringroup.Setup,
		userpool.Setup,
		userpooldomain.Setup,
		userpooluicustomization.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
