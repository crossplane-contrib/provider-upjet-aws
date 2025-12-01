// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	policy "github.com/upbound/provider-aws/v2/internal/controller/namespaced/appautoscaling/policy"
	scheduledaction "github.com/upbound/provider-aws/v2/internal/controller/namespaced/appautoscaling/scheduledaction"
	target "github.com/upbound/provider-aws/v2/internal/controller/namespaced/appautoscaling/target"
)

// Setup_appautoscaling creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_appautoscaling(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		policy.Setup,
		scheduledaction.Setup,
		target.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_appautoscaling creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_appautoscaling(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		policy.SetupGated,
		scheduledaction.SetupGated,
		target.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
