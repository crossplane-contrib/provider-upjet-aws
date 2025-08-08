// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	authenticationprofile "github.com/upbound/provider-aws/internal/controller/namespaced/redshift/authenticationprofile"
	cluster "github.com/upbound/provider-aws/internal/controller/namespaced/redshift/cluster"
	endpointaccess "github.com/upbound/provider-aws/internal/controller/namespaced/redshift/endpointaccess"
	eventsubscription "github.com/upbound/provider-aws/internal/controller/namespaced/redshift/eventsubscription"
	hsmclientcertificate "github.com/upbound/provider-aws/internal/controller/namespaced/redshift/hsmclientcertificate"
	hsmconfiguration "github.com/upbound/provider-aws/internal/controller/namespaced/redshift/hsmconfiguration"
	parametergroup "github.com/upbound/provider-aws/internal/controller/namespaced/redshift/parametergroup"
	scheduledaction "github.com/upbound/provider-aws/internal/controller/namespaced/redshift/scheduledaction"
	snapshotcopygrant "github.com/upbound/provider-aws/internal/controller/namespaced/redshift/snapshotcopygrant"
	snapshotschedule "github.com/upbound/provider-aws/internal/controller/namespaced/redshift/snapshotschedule"
	snapshotscheduleassociation "github.com/upbound/provider-aws/internal/controller/namespaced/redshift/snapshotscheduleassociation"
	subnetgroup "github.com/upbound/provider-aws/internal/controller/namespaced/redshift/subnetgroup"
	usagelimit "github.com/upbound/provider-aws/internal/controller/namespaced/redshift/usagelimit"
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

// SetupGated_redshift creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_redshift(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authenticationprofile.SetupGated,
		cluster.SetupGated,
		endpointaccess.SetupGated,
		eventsubscription.SetupGated,
		hsmclientcertificate.SetupGated,
		hsmconfiguration.SetupGated,
		parametergroup.SetupGated,
		scheduledaction.SetupGated,
		snapshotcopygrant.SetupGated,
		snapshotschedule.SetupGated,
		snapshotscheduleassociation.SetupGated,
		subnetgroup.SetupGated,
		usagelimit.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
