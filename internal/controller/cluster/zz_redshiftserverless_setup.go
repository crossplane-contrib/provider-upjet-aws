// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	endpointaccess "github.com/upbound/provider-aws/internal/controller/cluster/redshiftserverless/endpointaccess"
	redshiftserverlessnamespace "github.com/upbound/provider-aws/internal/controller/cluster/redshiftserverless/redshiftserverlessnamespace"
	resourcepolicy "github.com/upbound/provider-aws/internal/controller/cluster/redshiftserverless/resourcepolicy"
	snapshot "github.com/upbound/provider-aws/internal/controller/cluster/redshiftserverless/snapshot"
	usagelimit "github.com/upbound/provider-aws/internal/controller/cluster/redshiftserverless/usagelimit"
	workgroup "github.com/upbound/provider-aws/internal/controller/cluster/redshiftserverless/workgroup"
)

// Setup_redshiftserverless creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_redshiftserverless(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		endpointaccess.Setup,
		redshiftserverlessnamespace.Setup,
		resourcepolicy.Setup,
		snapshot.Setup,
		usagelimit.Setup,
		workgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_redshiftserverless creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_redshiftserverless(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		endpointaccess.SetupGated,
		redshiftserverlessnamespace.SetupGated,
		resourcepolicy.SetupGated,
		snapshot.SetupGated,
		usagelimit.SetupGated,
		workgroup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
