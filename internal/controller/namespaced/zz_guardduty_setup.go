// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	detector "github.com/upbound/provider-aws/v2/internal/controller/namespaced/guardduty/detector"
	filter "github.com/upbound/provider-aws/v2/internal/controller/namespaced/guardduty/filter"
	member "github.com/upbound/provider-aws/v2/internal/controller/namespaced/guardduty/member"
)

// Setup_guardduty creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_guardduty(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		detector.Setup,
		filter.Setup,
		member.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_guardduty creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_guardduty(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		detector.SetupGated,
		filter.SetupGated,
		member.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
