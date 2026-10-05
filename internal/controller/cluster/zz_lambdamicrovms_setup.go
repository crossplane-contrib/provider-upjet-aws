// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	image "github.com/upbound/provider-aws/v2/internal/controller/cluster/lambdamicrovms/image"
	microvm "github.com/upbound/provider-aws/v2/internal/controller/cluster/lambdamicrovms/microvm"
)

// Setup_lambdamicrovms creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_lambdamicrovms(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		image.Setup,
		microvm.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_lambdamicrovms creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_lambdamicrovms(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		image.SetupGated,
		microvm.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_lambdamicrovms registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_lambdamicrovms(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		image.SetupWebhookWithManager,
		microvm.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
