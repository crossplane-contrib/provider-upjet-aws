// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/upbound/provider-aws/internal/controller/cluster/kafka/cluster"
	configuration "github.com/upbound/provider-aws/internal/controller/cluster/kafka/configuration"
	scramsecretassociation "github.com/upbound/provider-aws/internal/controller/cluster/kafka/scramsecretassociation"
	serverlesscluster "github.com/upbound/provider-aws/internal/controller/cluster/kafka/serverlesscluster"
)

// Setup_kafka creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_kafka(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		configuration.Setup,
		scramsecretassociation.Setup,
		serverlesscluster.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
