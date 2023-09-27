/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	jobdefinition "github.com/upbound/provider-aws/internal/controller/batch/jobdefinition"
	schedulingpolicy "github.com/upbound/provider-aws/internal/controller/batch/schedulingpolicy"
)

// Setup_batch creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_batch(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		jobdefinition.Setup,
		schedulingpolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
