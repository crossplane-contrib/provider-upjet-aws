// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	conditionalforwarder "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ds/conditionalforwarder"
	directory "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ds/directory"
	shareddirectory "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ds/shareddirectory"
)

// Setup_ds creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ds(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		conditionalforwarder.Setup,
		directory.Setup,
		shareddirectory.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
