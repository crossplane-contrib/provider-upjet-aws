// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	principalassociation "github.com/upbound/provider-aws/internal/controller/cluster/ram/principalassociation"
	resourceassociation "github.com/upbound/provider-aws/internal/controller/cluster/ram/resourceassociation"
	resourceshare "github.com/upbound/provider-aws/internal/controller/cluster/ram/resourceshare"
	resourceshareaccepter "github.com/upbound/provider-aws/internal/controller/cluster/ram/resourceshareaccepter"
)

// Setup_ram creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ram(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		principalassociation.Setup,
		resourceassociation.Setup,
		resourceshare.Setup,
		resourceshareaccepter.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
