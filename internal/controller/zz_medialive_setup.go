// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	input "github.com/upbound/provider-aws/internal/controller/medialive/input"
	inputsecuritygroup "github.com/upbound/provider-aws/internal/controller/medialive/inputsecuritygroup"
	multiplex "github.com/upbound/provider-aws/internal/controller/medialive/multiplex"
)

// Setup_medialive creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_medialive(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		input.Setup,
		inputsecuritygroup.Setup,
		multiplex.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
