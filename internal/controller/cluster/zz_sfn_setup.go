// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	activity "github.com/upbound/provider-aws/v2/internal/controller/cluster/sfn/activity"
	statemachine "github.com/upbound/provider-aws/v2/internal/controller/cluster/sfn/statemachine"
)

// Setup_sfn creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_sfn(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		activity.Setup,
		statemachine.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_sfn creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_sfn(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		activity.SetupGated,
		statemachine.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
