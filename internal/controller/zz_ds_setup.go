/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	conditionalforwarder "github.com/upbound/provider-aws/internal/controller/ds/conditionalforwarder"
	directory "github.com/upbound/provider-aws/internal/controller/ds/directory"
	shareddirectory "github.com/upbound/provider-aws/internal/controller/ds/shareddirectory"
)

// Setup_ds creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ds(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		conditionalforwarder.Setup,
		directory.Setup,
		shareddirectory.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
