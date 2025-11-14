// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	alertmanagerdefinition "github.com/upbound/provider-aws/internal/controller/cluster/amp/alertmanagerdefinition"
	rulegroupnamespace "github.com/upbound/provider-aws/internal/controller/cluster/amp/rulegroupnamespace"
	scraper "github.com/upbound/provider-aws/internal/controller/cluster/amp/scraper"
	workspace "github.com/upbound/provider-aws/internal/controller/cluster/amp/workspace"
)

// Setup_amp creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_amp(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alertmanagerdefinition.Setup,
		rulegroupnamespace.Setup,
		scraper.Setup,
		workspace.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_amp creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_amp(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alertmanagerdefinition.SetupGated,
		rulegroupnamespace.SetupGated,
		scraper.SetupGated,
		workspace.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
