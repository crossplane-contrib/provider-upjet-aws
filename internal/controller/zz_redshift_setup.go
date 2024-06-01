// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	authenticationprofile "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/redshift/authenticationprofile"
	cluster "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/redshift/cluster"
	eventsubscription "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/redshift/eventsubscription"
	hsmclientcertificate "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/redshift/hsmclientcertificate"
	hsmconfiguration "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/redshift/hsmconfiguration"
	parametergroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/redshift/parametergroup"
	scheduledaction "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/redshift/scheduledaction"
	snapshotcopygrant "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/redshift/snapshotcopygrant"
	snapshotschedule "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/redshift/snapshotschedule"
	snapshotscheduleassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/redshift/snapshotscheduleassociation"
	subnetgroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/redshift/subnetgroup"
	usagelimit "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/redshift/usagelimit"
)

// Setup_redshift creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_redshift(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authenticationprofile.Setup,
		cluster.Setup,
		eventsubscription.Setup,
		hsmclientcertificate.Setup,
		hsmconfiguration.Setup,
		parametergroup.Setup,
		scheduledaction.Setup,
		snapshotcopygrant.Setup,
		snapshotschedule.Setup,
		snapshotscheduleassociation.Setup,
		subnetgroup.Setup,
		usagelimit.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
