/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	instance "github.com/upbound/provider-aws/internal/controller/rds/instance"
	optiongroup "github.com/upbound/provider-aws/internal/controller/rds/optiongroup"
	parametergroup "github.com/upbound/provider-aws/internal/controller/rds/parametergroup"
	subnetgroup "github.com/upbound/provider-aws/internal/controller/rds/subnetgroup"
)

// Setup_rds creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_rds(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		instance.Setup,
		optiongroup.Setup,
		parametergroup.Setup,
		subnetgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
