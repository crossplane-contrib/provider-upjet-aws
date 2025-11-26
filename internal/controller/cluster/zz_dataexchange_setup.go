// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	dataset "github.com/upbound/provider-aws/v2/internal/controller/cluster/dataexchange/dataset"
	revision "github.com/upbound/provider-aws/v2/internal/controller/cluster/dataexchange/revision"
)

// Setup_dataexchange creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dataexchange(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		dataset.Setup,
		revision.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_dataexchange creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dataexchange(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		dataset.SetupGated,
		revision.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
