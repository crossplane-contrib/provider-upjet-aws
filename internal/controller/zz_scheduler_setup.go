// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	schedule "github.com/upbound/provider-aws/internal/controller/scheduler/schedule"
	schedulegroup "github.com/upbound/provider-aws/internal/controller/scheduler/schedulegroup"
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
