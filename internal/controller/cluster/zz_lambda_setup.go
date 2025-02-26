// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	alias "github.com/upbound/provider-aws/internal/controller/lambda/alias"
	codesigningconfig "github.com/upbound/provider-aws/internal/controller/lambda/codesigningconfig"
	eventsourcemapping "github.com/upbound/provider-aws/internal/controller/lambda/eventsourcemapping"
	function "github.com/upbound/provider-aws/internal/controller/lambda/function"
	functioneventinvokeconfig "github.com/upbound/provider-aws/internal/controller/lambda/functioneventinvokeconfig"
	functionurl "github.com/upbound/provider-aws/internal/controller/lambda/functionurl"
	invocation "github.com/upbound/provider-aws/internal/controller/lambda/invocation"
	layerversion "github.com/upbound/provider-aws/internal/controller/lambda/layerversion"
	layerversionpermission "github.com/upbound/provider-aws/internal/controller/lambda/layerversionpermission"
	permission "github.com/upbound/provider-aws/internal/controller/lambda/permission"
	provisionedconcurrencyconfig "github.com/upbound/provider-aws/internal/controller/lambda/provisionedconcurrencyconfig"
)

// Setup_lambda creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_lambda(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alias.Setup,
		codesigningconfig.Setup,
		eventsourcemapping.Setup,
		function.Setup,
		functioneventinvokeconfig.Setup,
		functionurl.Setup,
		invocation.Setup,
		layerversion.Setup,
		layerversionpermission.Setup,
		permission.Setup,
		provisionedconcurrencyconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
