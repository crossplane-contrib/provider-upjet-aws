// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cell "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/route53recoveryreadiness/cell"
	readinesscheck "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/route53recoveryreadiness/readinesscheck"
	recoverygroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/route53recoveryreadiness/recoverygroup"
	resourceset "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/route53recoveryreadiness/resourceset"
)

// Setup_route53recoveryreadiness creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_route53recoveryreadiness(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cell.Setup,
		readinesscheck.Setup,
		recoverygroup.Setup,
		resourceset.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
