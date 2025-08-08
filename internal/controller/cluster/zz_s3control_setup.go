// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accesspoint "github.com/upbound/provider-aws/internal/controller/cluster/s3control/accesspoint"
	accesspointpolicy "github.com/upbound/provider-aws/internal/controller/cluster/s3control/accesspointpolicy"
	accountpublicaccessblock "github.com/upbound/provider-aws/internal/controller/cluster/s3control/accountpublicaccessblock"
	multiregionaccesspoint "github.com/upbound/provider-aws/internal/controller/cluster/s3control/multiregionaccesspoint"
	multiregionaccesspointpolicy "github.com/upbound/provider-aws/internal/controller/cluster/s3control/multiregionaccesspointpolicy"
	objectlambdaaccesspoint "github.com/upbound/provider-aws/internal/controller/cluster/s3control/objectlambdaaccesspoint"
	objectlambdaaccesspointpolicy "github.com/upbound/provider-aws/internal/controller/cluster/s3control/objectlambdaaccesspointpolicy"
	storagelensconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/s3control/storagelensconfiguration"
)

// Setup_s3control creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_s3control(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesspoint.Setup,
		accesspointpolicy.Setup,
		accountpublicaccessblock.Setup,
		multiregionaccesspoint.Setup,
		multiregionaccesspointpolicy.Setup,
		objectlambdaaccesspoint.Setup,
		objectlambdaaccesspointpolicy.Setup,
		storagelensconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_s3control creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_s3control(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesspoint.SetupGated,
		accesspointpolicy.SetupGated,
		accountpublicaccessblock.SetupGated,
		multiregionaccesspoint.SetupGated,
		multiregionaccesspointpolicy.SetupGated,
		objectlambdaaccesspoint.SetupGated,
		objectlambdaaccesspointpolicy.SetupGated,
		storagelensconfiguration.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
