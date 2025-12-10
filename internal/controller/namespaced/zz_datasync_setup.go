// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	locations3 "github.com/upbound/provider-aws/v2/internal/controller/namespaced/datasync/locations3"
	task "github.com/upbound/provider-aws/v2/internal/controller/namespaced/datasync/task"
)

// Setup_datasync creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datasync(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		locations3.Setup,
		task.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_datasync creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_datasync(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		locations3.SetupGated,
		task.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
