// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	authenticationprofile "github.com/upbound/provider-aws/internal/controller/cluster/redshift/authenticationprofile"
	cluster "github.com/upbound/provider-aws/internal/controller/cluster/redshift/cluster"
	endpointaccess "github.com/upbound/provider-aws/internal/controller/cluster/redshift/endpointaccess"
	eventsubscription "github.com/upbound/provider-aws/internal/controller/cluster/redshift/eventsubscription"
	hsmclientcertificate "github.com/upbound/provider-aws/internal/controller/cluster/redshift/hsmclientcertificate"
	hsmconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/redshift/hsmconfiguration"
	parametergroup "github.com/upbound/provider-aws/internal/controller/cluster/redshift/parametergroup"
	scheduledaction "github.com/upbound/provider-aws/internal/controller/cluster/redshift/scheduledaction"
	snapshotcopygrant "github.com/upbound/provider-aws/internal/controller/cluster/redshift/snapshotcopygrant"
	snapshotschedule "github.com/upbound/provider-aws/internal/controller/cluster/redshift/snapshotschedule"
	snapshotscheduleassociation "github.com/upbound/provider-aws/internal/controller/cluster/redshift/snapshotscheduleassociation"
	subnetgroup "github.com/upbound/provider-aws/internal/controller/cluster/redshift/subnetgroup"
	usagelimit "github.com/upbound/provider-aws/internal/controller/cluster/redshift/usagelimit"
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
