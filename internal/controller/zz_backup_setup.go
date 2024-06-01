// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	framework "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/backup/framework"
	globalsettings "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/backup/globalsettings"
	plan "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/backup/plan"
	regionsettings "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/backup/regionsettings"
	reportplan "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/backup/reportplan"
	selection "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/backup/selection"
	vault "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/backup/vault"
	vaultlockconfiguration "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/backup/vaultlockconfiguration"
	vaultnotifications "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/backup/vaultnotifications"
	vaultpolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/backup/vaultpolicy"
)

// Setup_backup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_backup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		framework.Setup,
		globalsettings.Setup,
		plan.Setup,
		regionsettings.Setup,
		reportplan.Setup,
		selection.Setup,
		vault.Setup,
		vaultlockconfiguration.Setup,
		vaultnotifications.Setup,
		vaultpolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
