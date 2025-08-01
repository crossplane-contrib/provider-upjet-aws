// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	secret "github.com/upbound/provider-aws/internal/controller/cluster/secretsmanager/secret"
	secretpolicy "github.com/upbound/provider-aws/internal/controller/cluster/secretsmanager/secretpolicy"
	secretrotation "github.com/upbound/provider-aws/internal/controller/cluster/secretsmanager/secretrotation"
	secretversion "github.com/upbound/provider-aws/internal/controller/cluster/secretsmanager/secretversion"
)

// Setup_secretsmanager creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_secretsmanager(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		secret.Setup,
		secretpolicy.Setup,
		secretrotation.Setup,
		secretversion.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_secretsmanager creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_secretsmanager(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		secret.SetupGated,
		secretpolicy.SetupGated,
		secretrotation.SetupGated,
		secretversion.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
