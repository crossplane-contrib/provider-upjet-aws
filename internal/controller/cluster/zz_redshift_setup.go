// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	authenticationprofile "github.com/upbound/provider-aws/internal/controller/redshift/authenticationprofile"
	cluster "github.com/upbound/provider-aws/internal/controller/redshift/cluster"
	endpointaccess "github.com/upbound/provider-aws/internal/controller/redshift/endpointaccess"
	eventsubscription "github.com/upbound/provider-aws/internal/controller/redshift/eventsubscription"
	hsmclientcertificate "github.com/upbound/provider-aws/internal/controller/redshift/hsmclientcertificate"
	hsmconfiguration "github.com/upbound/provider-aws/internal/controller/redshift/hsmconfiguration"
	parametergroup "github.com/upbound/provider-aws/internal/controller/redshift/parametergroup"
	scheduledaction "github.com/upbound/provider-aws/internal/controller/redshift/scheduledaction"
	snapshotcopygrant "github.com/upbound/provider-aws/internal/controller/redshift/snapshotcopygrant"
	snapshotschedule "github.com/upbound/provider-aws/internal/controller/redshift/snapshotschedule"
	snapshotscheduleassociation "github.com/upbound/provider-aws/internal/controller/redshift/snapshotscheduleassociation"
	subnetgroup "github.com/upbound/provider-aws/internal/controller/redshift/subnetgroup"
	usagelimit "github.com/upbound/provider-aws/internal/controller/redshift/usagelimit"
)

// Setup_redshift creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_redshift(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authenticationprofile.Setup,
		cluster.Setup,
		endpointaccess.Setup,
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
