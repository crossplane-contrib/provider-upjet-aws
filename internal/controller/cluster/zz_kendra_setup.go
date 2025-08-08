// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	datasource "github.com/upbound/provider-aws/internal/controller/cluster/kendra/datasource"
	experience "github.com/upbound/provider-aws/internal/controller/cluster/kendra/experience"
	index "github.com/upbound/provider-aws/internal/controller/cluster/kendra/index"
	querysuggestionsblocklist "github.com/upbound/provider-aws/internal/controller/cluster/kendra/querysuggestionsblocklist"
	thesaurus "github.com/upbound/provider-aws/internal/controller/cluster/kendra/thesaurus"
)

// Setup_kendra creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_kendra(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		datasource.Setup,
		experience.Setup,
		index.Setup,
		querysuggestionsblocklist.Setup,
		thesaurus.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_kendra creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_kendra(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		datasource.SetupGated,
		experience.SetupGated,
		index.SetupGated,
		querysuggestionsblocklist.SetupGated,
		thesaurus.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
