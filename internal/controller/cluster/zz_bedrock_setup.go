// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	inferenceprofile "github.com/upbound/provider-aws/v2/internal/controller/cluster/bedrock/inferenceprofile"
)

// Setup_bedrock creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_bedrock(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		inferenceprofile.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_bedrock creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_bedrock(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		inferenceprofile.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
