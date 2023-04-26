/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	stack "github.com/upbound/provider-aws/internal/controller/cloudformation/stack"
	stackset "github.com/upbound/provider-aws/internal/controller/cloudformation/stackset"
)

// Setup_cloudformation creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudformation(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		stack.Setup,
		stackset.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
