// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	datasource "github.com/upbound/provider-aws/internal/controller/kendra/datasource"
	experience "github.com/upbound/provider-aws/internal/controller/kendra/experience"
	index "github.com/upbound/provider-aws/internal/controller/kendra/index"
	querysuggestionsblocklist "github.com/upbound/provider-aws/internal/controller/kendra/querysuggestionsblocklist"
	thesaurus "github.com/upbound/provider-aws/internal/controller/kendra/thesaurus"
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
