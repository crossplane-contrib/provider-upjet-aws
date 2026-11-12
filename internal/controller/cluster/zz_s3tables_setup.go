// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	s3tablesnamespace "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3tables/s3tablesnamespace"
	table "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3tables/table"
	tablebucket "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3tables/tablebucket"
	tablebucketpolicy "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3tables/tablebucketpolicy"
	tablebucketreplication "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3tables/tablebucketreplication"
	tablepolicy "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3tables/tablepolicy"
	tablereplication "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3tables/tablereplication"
)

// Setup_s3tables creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_s3tables(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		s3tablesnamespace.Setup,
		table.Setup,
		tablebucket.Setup,
		tablebucketpolicy.Setup,
		tablebucketreplication.Setup,
		tablepolicy.Setup,
		tablereplication.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_s3tables creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_s3tables(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		s3tablesnamespace.SetupGated,
		table.SetupGated,
		tablebucket.SetupGated,
		tablebucketpolicy.SetupGated,
		tablebucketreplication.SetupGated,
		tablepolicy.SetupGated,
		tablereplication.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_s3tables registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_s3tables(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		s3tablesnamespace.SetupWebhookWithManager,
		table.SetupWebhookWithManager,
		tablebucket.SetupWebhookWithManager,
		tablebucketpolicy.SetupWebhookWithManager,
		tablebucketreplication.SetupWebhookWithManager,
		tablepolicy.SetupWebhookWithManager,
		tablereplication.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
