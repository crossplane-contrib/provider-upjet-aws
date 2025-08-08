// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	attachment "github.com/upbound/provider-aws/internal/controller/namespaced/autoscaling/attachment"
	autoscalinggroup "github.com/upbound/provider-aws/internal/controller/namespaced/autoscaling/autoscalinggroup"
	grouptag "github.com/upbound/provider-aws/internal/controller/namespaced/autoscaling/grouptag"
	launchconfiguration "github.com/upbound/provider-aws/internal/controller/namespaced/autoscaling/launchconfiguration"
	lifecyclehook "github.com/upbound/provider-aws/internal/controller/namespaced/autoscaling/lifecyclehook"
	notification "github.com/upbound/provider-aws/internal/controller/namespaced/autoscaling/notification"
	policy "github.com/upbound/provider-aws/internal/controller/namespaced/autoscaling/policy"
	schedule "github.com/upbound/provider-aws/internal/controller/namespaced/autoscaling/schedule"
)

// Setup_autoscaling creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_autoscaling(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		attachment.Setup,
		autoscalinggroup.Setup,
		grouptag.Setup,
		launchconfiguration.Setup,
		lifecyclehook.Setup,
		notification.Setup,
		policy.Setup,
		schedule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_autoscaling creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_autoscaling(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		attachment.SetupGated,
		autoscalinggroup.SetupGated,
		grouptag.SetupGated,
		launchconfiguration.SetupGated,
		lifecyclehook.SetupGated,
		notification.SetupGated,
		policy.SetupGated,
		schedule.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
