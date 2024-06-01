// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/dax/cluster"
	parametergroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/dax/parametergroup"
	subnetgroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/dax/subnetgroup"
)

// Setup_dax creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dax(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		parametergroup.Setup,
		subnetgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
