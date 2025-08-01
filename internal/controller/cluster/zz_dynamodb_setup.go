// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	contributorinsights "github.com/upbound/provider-aws/internal/controller/cluster/dynamodb/contributorinsights"
	globaltable "github.com/upbound/provider-aws/internal/controller/cluster/dynamodb/globaltable"
	kinesisstreamingdestination "github.com/upbound/provider-aws/internal/controller/cluster/dynamodb/kinesisstreamingdestination"
	resourcepolicy "github.com/upbound/provider-aws/internal/controller/cluster/dynamodb/resourcepolicy"
	table "github.com/upbound/provider-aws/internal/controller/cluster/dynamodb/table"
	tableitem "github.com/upbound/provider-aws/internal/controller/cluster/dynamodb/tableitem"
	tablereplica "github.com/upbound/provider-aws/internal/controller/cluster/dynamodb/tablereplica"
	tag "github.com/upbound/provider-aws/internal/controller/cluster/dynamodb/tag"
)

// Setup_dynamodb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dynamodb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		contributorinsights.Setup,
		globaltable.Setup,
		kinesisstreamingdestination.Setup,
		resourcepolicy.Setup,
		table.Setup,
		tableitem.Setup,
		tablereplica.Setup,
		tag.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_dynamodb creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dynamodb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		contributorinsights.SetupGated,
		globaltable.SetupGated,
		kinesisstreamingdestination.SetupGated,
		resourcepolicy.SetupGated,
		table.SetupGated,
		tableitem.SetupGated,
		tablereplica.SetupGated,
		tag.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
