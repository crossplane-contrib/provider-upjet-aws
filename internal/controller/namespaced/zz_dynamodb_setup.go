// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	contributorinsights "github.com/upbound/provider-aws/internal/controller/namespaced/dynamodb/contributorinsights"
	globaltable "github.com/upbound/provider-aws/internal/controller/namespaced/dynamodb/globaltable"
	kinesisstreamingdestination "github.com/upbound/provider-aws/internal/controller/namespaced/dynamodb/kinesisstreamingdestination"
	resourcepolicy "github.com/upbound/provider-aws/internal/controller/namespaced/dynamodb/resourcepolicy"
	table "github.com/upbound/provider-aws/internal/controller/namespaced/dynamodb/table"
	tableitem "github.com/upbound/provider-aws/internal/controller/namespaced/dynamodb/tableitem"
	tablereplica "github.com/upbound/provider-aws/internal/controller/namespaced/dynamodb/tablereplica"
	tag "github.com/upbound/provider-aws/internal/controller/namespaced/dynamodb/tag"
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
