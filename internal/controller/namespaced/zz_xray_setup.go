// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	encryptionconfig "github.com/upbound/provider-aws/internal/controller/xray/encryptionconfig"
	group "github.com/upbound/provider-aws/internal/controller/xray/group"
	samplingrule "github.com/upbound/provider-aws/internal/controller/xray/samplingrule"
)

// Setup_xray creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_xray(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		encryptionconfig.Setup,
		group.Setup,
		samplingrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
