// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	vault "github.com/upbound/provider-aws/internal/controller/cluster/glacier/vault"
	vaultlock "github.com/upbound/provider-aws/internal/controller/cluster/glacier/vaultlock"
)

// Setup_glacier creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_glacier(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		vault.Setup,
		vaultlock.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_glacier creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_glacier(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		vault.SetupGated,
		vaultlock.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
