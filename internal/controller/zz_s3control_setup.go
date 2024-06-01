// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	accesspoint "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/s3control/accesspoint"
	accesspointpolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/s3control/accesspointpolicy"
	accountpublicaccessblock "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/s3control/accountpublicaccessblock"
	multiregionaccesspoint "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/s3control/multiregionaccesspoint"
	multiregionaccesspointpolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/s3control/multiregionaccesspointpolicy"
	objectlambdaaccesspoint "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/s3control/objectlambdaaccesspoint"
	objectlambdaaccesspointpolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/s3control/objectlambdaaccesspointpolicy"
	storagelensconfiguration "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/s3control/storagelensconfiguration"
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
