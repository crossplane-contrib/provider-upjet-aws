// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	discoverer "github.com/upbound/provider-aws/internal/controller/cluster/schemas/discoverer"
	registry "github.com/upbound/provider-aws/internal/controller/cluster/schemas/registry"
	schema "github.com/upbound/provider-aws/internal/controller/cluster/schemas/schema"
)

// Setup_schemas creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_schemas(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		discoverer.Setup,
		registry.Setup,
		schema.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_schemas creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_schemas(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		discoverer.SetupGated,
		registry.SetupGated,
		schema.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
