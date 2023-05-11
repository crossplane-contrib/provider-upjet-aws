/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	pipeline "github.com/upbound/provider-aws/internal/controller/datapipeline/pipeline"
)

// Setup_datapipeline creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datapipeline(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		pipeline.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
