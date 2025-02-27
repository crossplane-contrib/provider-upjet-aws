// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	locations3 "github.com/upbound/provider-aws/internal/controller/datasync/locations3"
	task "github.com/upbound/provider-aws/internal/controller/datasync/task"
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
