/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	feature "github.com/upbound/provider-aws/internal/controller/evidently/feature"
	project "github.com/upbound/provider-aws/internal/controller/evidently/project"
	segment "github.com/upbound/provider-aws/internal/controller/evidently/segment"
)

// Setup_evidently creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_evidently(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		feature.Setup,
		project.Setup,
		segment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
