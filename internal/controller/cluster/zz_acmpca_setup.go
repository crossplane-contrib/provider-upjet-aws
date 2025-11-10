// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	certificate "github.com/upbound/provider-aws/v2/internal/controller/cluster/acmpca/certificate"
	certificateauthority "github.com/upbound/provider-aws/v2/internal/controller/cluster/acmpca/certificateauthority"
	certificateauthoritycertificate "github.com/upbound/provider-aws/v2/internal/controller/cluster/acmpca/certificateauthoritycertificate"
	permission "github.com/upbound/provider-aws/v2/internal/controller/cluster/acmpca/permission"
	policy "github.com/upbound/provider-aws/v2/internal/controller/cluster/acmpca/policy"
)

// Setup_acmpca creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_acmpca(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		certificate.Setup,
		certificateauthority.Setup,
		certificateauthoritycertificate.Setup,
		permission.Setup,
		policy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_acmpca creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_acmpca(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		certificate.SetupGated,
		certificateauthority.SetupGated,
		certificateauthoritycertificate.SetupGated,
		permission.SetupGated,
		policy.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
