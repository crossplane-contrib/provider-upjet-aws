// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	attachment "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/autoscaling/attachment"
	autoscalinggroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/autoscaling/autoscalinggroup"
	grouptag "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/autoscaling/grouptag"
	launchconfiguration "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/autoscaling/launchconfiguration"
	lifecyclehook "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/autoscaling/lifecyclehook"
	notification "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/autoscaling/notification"
	policy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/autoscaling/policy"
	schedule "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/autoscaling/schedule"
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
