// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accesspoint "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3files/accesspoint"
	filesystem "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3files/filesystem"
	filesystempolicy "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3files/filesystempolicy"
	mounttarget "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3files/mounttarget"
	synchronizationconfiguration "github.com/upbound/provider-aws/v2/internal/controller/cluster/s3files/synchronizationconfiguration"
)

// Setup_s3files creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_s3files(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesspoint.Setup,
		filesystem.Setup,
		filesystempolicy.Setup,
		mounttarget.Setup,
		synchronizationconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_s3files creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_s3files(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesspoint.SetupGated,
		filesystem.SetupGated,
		filesystempolicy.SetupGated,
		mounttarget.SetupGated,
		synchronizationconfiguration.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_s3files registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_s3files(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		accesspoint.SetupWebhookWithManager,
		filesystem.SetupWebhookWithManager,
		filesystempolicy.SetupWebhookWithManager,
		mounttarget.SetupWebhookWithManager,
		synchronizationconfiguration.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
