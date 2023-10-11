// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	codepipeline "github.com/upbound/provider-aws/internal/controller/codepipeline/codepipeline"
	customactiontype "github.com/upbound/provider-aws/internal/controller/codepipeline/customactiontype"
	webhook "github.com/upbound/provider-aws/internal/controller/codepipeline/webhook"
)

// Setup_codepipeline creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_codepipeline(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		codepipeline.Setup,
		customactiontype.Setup,
		webhook.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
