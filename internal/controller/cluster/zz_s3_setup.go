// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	bucket "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucket"
	bucketaccelerateconfiguration "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketaccelerateconfiguration"
	bucketacl "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketacl"
	bucketanalyticsconfiguration "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketanalyticsconfiguration"
	bucketcorsconfiguration "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketcorsconfiguration"
	bucketintelligenttieringconfiguration "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketintelligenttieringconfiguration"
	bucketinventory "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketinventory"
	bucketlifecycleconfiguration "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketlifecycleconfiguration"
	bucketlogging "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketlogging"
	bucketmetric "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketmetric"
	bucketnotification "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketnotification"
	bucketobject "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketobject"
	bucketobjectlockconfiguration "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketobjectlockconfiguration"
	bucketownershipcontrols "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketownershipcontrols"
	bucketpolicy "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketpolicy"
	bucketpublicaccessblock "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketpublicaccessblock"
	bucketreplicationconfiguration "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketreplicationconfiguration"
	bucketrequestpaymentconfiguration "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketrequestpaymentconfiguration"
	bucketserversideencryptionconfiguration "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketserversideencryptionconfiguration"
	bucketversioning "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketversioning"
	bucketwebsiteconfiguration "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/bucketwebsiteconfiguration"
	directorybucket "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/directorybucket"
	object "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/object"
	objectcopy "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3/objectcopy"
)

// Setup_s3 creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_s3(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bucket.Setup,
		bucketaccelerateconfiguration.Setup,
		bucketacl.Setup,
		bucketanalyticsconfiguration.Setup,
		bucketcorsconfiguration.Setup,
		bucketintelligenttieringconfiguration.Setup,
		bucketinventory.Setup,
		bucketlifecycleconfiguration.Setup,
		bucketlogging.Setup,
		bucketmetric.Setup,
		bucketnotification.Setup,
		bucketobject.Setup,
		bucketobjectlockconfiguration.Setup,
		bucketownershipcontrols.Setup,
		bucketpolicy.Setup,
		bucketpublicaccessblock.Setup,
		bucketreplicationconfiguration.Setup,
		bucketrequestpaymentconfiguration.Setup,
		bucketserversideencryptionconfiguration.Setup,
		bucketversioning.Setup,
		bucketwebsiteconfiguration.Setup,
		directorybucket.Setup,
		object.Setup,
		objectcopy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_s3 creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_s3(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bucket.SetupGated,
		bucketaccelerateconfiguration.SetupGated,
		bucketacl.SetupGated,
		bucketanalyticsconfiguration.SetupGated,
		bucketcorsconfiguration.SetupGated,
		bucketintelligenttieringconfiguration.SetupGated,
		bucketinventory.SetupGated,
		bucketlifecycleconfiguration.SetupGated,
		bucketlogging.SetupGated,
		bucketmetric.SetupGated,
		bucketnotification.SetupGated,
		bucketobject.SetupGated,
		bucketobjectlockconfiguration.SetupGated,
		bucketownershipcontrols.SetupGated,
		bucketpolicy.SetupGated,
		bucketpublicaccessblock.SetupGated,
		bucketreplicationconfiguration.SetupGated,
		bucketrequestpaymentconfiguration.SetupGated,
		bucketserversideencryptionconfiguration.SetupGated,
		bucketversioning.SetupGated,
		bucketwebsiteconfiguration.SetupGated,
		directorybucket.SetupGated,
		object.SetupGated,
		objectcopy.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
