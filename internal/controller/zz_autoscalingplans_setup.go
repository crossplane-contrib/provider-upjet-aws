/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	scalingplan "github.com/upbound/provider-aws/internal/controller/autoscalingplans/scalingplan"
)

// Setup_autoscalingplans creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_autoscalingplans(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		scalingplan.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
