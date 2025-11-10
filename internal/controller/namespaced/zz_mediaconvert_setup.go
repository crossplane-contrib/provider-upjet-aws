// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	queue "github.com/upbound/provider-aws/v2/internal/controller/namespaced/mediaconvert/queue"
)

// Setup_mediaconvert creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_mediaconvert(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		queue.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_mediaconvert creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_mediaconvert(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		queue.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
