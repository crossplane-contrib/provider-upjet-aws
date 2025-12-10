// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	channel "github.com/upbound/provider-aws/v2/internal/controller/cluster/ivs/channel"
	recordingconfiguration "github.com/upbound/provider-aws/v2/internal/controller/cluster/ivs/recordingconfiguration"
)

// Setup_ivs creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ivs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		channel.Setup,
		recordingconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_ivs creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_ivs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		channel.SetupGated,
		recordingconfiguration.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
