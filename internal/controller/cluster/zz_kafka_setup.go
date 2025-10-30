// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cluster "github.com/upbound/provider-aws/internal/controller/cluster/kafka/cluster"
	clusterpolicy "github.com/upbound/provider-aws/internal/controller/cluster/kafka/clusterpolicy"
	configuration "github.com/upbound/provider-aws/internal/controller/cluster/kafka/configuration"
	replicator "github.com/upbound/provider-aws/internal/controller/cluster/kafka/replicator"
	scramsecretassociation "github.com/upbound/provider-aws/internal/controller/cluster/kafka/scramsecretassociation"
	serverlesscluster "github.com/upbound/provider-aws/internal/controller/cluster/kafka/serverlesscluster"
	singlescramsecretassociation "github.com/upbound/provider-aws/internal/controller/cluster/kafka/singlescramsecretassociation"
)

// Setup_kafka creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_kafka(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		clusterpolicy.Setup,
		configuration.Setup,
		replicator.Setup,
		scramsecretassociation.Setup,
		serverlesscluster.Setup,
		singlescramsecretassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_kafka creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_kafka(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.SetupGated,
		clusterpolicy.SetupGated,
		configuration.SetupGated,
		replicator.SetupGated,
		scramsecretassociation.SetupGated,
		serverlesscluster.SetupGated,
		singlescramsecretassociation.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
