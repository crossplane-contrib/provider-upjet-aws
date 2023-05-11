/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	profile "github.com/upbound/provider-aws/internal/controller/rolesanywhere/profile"
)

// Setup_rolesanywhere creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_rolesanywhere(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		profile.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
