/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	alias "github.com/upbound/provider-aws/internal/controller/lambda/alias"
	function "github.com/upbound/provider-aws/internal/controller/lambda/function"
	functioneventinvokeconfig "github.com/upbound/provider-aws/internal/controller/lambda/functioneventinvokeconfig"
	permission "github.com/upbound/provider-aws/internal/controller/lambda/permission"
)

// Setup_lambda creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_lambda(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alias.Setup,
		function.Setup,
		functioneventinvokeconfig.Setup,
		permission.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
