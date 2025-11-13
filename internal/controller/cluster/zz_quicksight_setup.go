// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accountsubscription "github.com/upbound/provider-aws/v2/internal/controller/cluster/quicksight/accountsubscription"
	analysis "github.com/upbound/provider-aws/v2/internal/controller/cluster/quicksight/analysis"
	dashboard "github.com/upbound/provider-aws/v2/internal/controller/cluster/quicksight/dashboard"
	dataset "github.com/upbound/provider-aws/v2/internal/controller/cluster/quicksight/dataset"
	datasource "github.com/upbound/provider-aws/v2/internal/controller/cluster/quicksight/datasource"
	group "github.com/upbound/provider-aws/v2/internal/controller/cluster/quicksight/group"
	groupmembership "github.com/upbound/provider-aws/v2/internal/controller/cluster/quicksight/groupmembership"
	user "github.com/upbound/provider-aws/v2/internal/controller/cluster/quicksight/user"
	vpcconnection "github.com/upbound/provider-aws/v2/internal/controller/cluster/quicksight/vpcconnection"
)

// Setup_quicksight creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_quicksight(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accountsubscription.Setup,
		analysis.Setup,
		dashboard.Setup,
		dataset.Setup,
		datasource.Setup,
		group.Setup,
		groupmembership.Setup,
		user.Setup,
		vpcconnection.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_quicksight creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_quicksight(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accountsubscription.SetupGated,
		analysis.SetupGated,
		dashboard.SetupGated,
		dataset.SetupGated,
		datasource.SetupGated,
		group.SetupGated,
		groupmembership.SetupGated,
		user.SetupGated,
		vpcconnection.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
