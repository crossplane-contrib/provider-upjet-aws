// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	endpointaccess "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/redshiftserverless/endpointaccess"
	redshiftserverlessnamespace "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/redshiftserverless/redshiftserverlessnamespace"
	resourcepolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/redshiftserverless/resourcepolicy"
	snapshot "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/redshiftserverless/snapshot"
	usagelimit "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/redshiftserverless/usagelimit"
	workgroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/redshiftserverless/workgroup"
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
