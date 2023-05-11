/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	cloudformationstack "github.com/upbound/provider-aws/internal/controller/serverlessrepo/cloudformationstack"
)

// Setup_serverlessrepo creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_serverlessrepo(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cloudformationstack.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
