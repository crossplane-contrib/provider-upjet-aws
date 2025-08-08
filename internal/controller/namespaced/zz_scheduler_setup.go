// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	schedule "github.com/upbound/provider-aws/internal/controller/namespaced/scheduler/schedule"
	schedulegroup "github.com/upbound/provider-aws/internal/controller/namespaced/scheduler/schedulegroup"
)

// Setup_scheduler creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_scheduler(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		schedule.Setup,
		schedulegroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_scheduler creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_scheduler(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		schedule.SetupGated,
		schedulegroup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
