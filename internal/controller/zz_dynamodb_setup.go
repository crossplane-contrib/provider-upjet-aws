// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	contributorinsights "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/dynamodb/contributorinsights"
	globaltable "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/dynamodb/globaltable"
	kinesisstreamingdestination "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/dynamodb/kinesisstreamingdestination"
	table "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/dynamodb/table"
	tableitem "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/dynamodb/tableitem"
	tablereplica "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/dynamodb/tablereplica"
	tag "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/dynamodb/tag"
)

// Setup_dynamodb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dynamodb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		contributorinsights.Setup,
		globaltable.Setup,
		kinesisstreamingdestination.Setup,
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
