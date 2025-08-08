// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	signingjob "github.com/upbound/provider-aws/internal/controller/cluster/signer/signingjob"
	signingprofile "github.com/upbound/provider-aws/internal/controller/cluster/signer/signingprofile"
	signingprofilepermission "github.com/upbound/provider-aws/internal/controller/cluster/signer/signingprofilepermission"
)

// Setup_signer creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_signer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		signingjob.Setup,
		signingprofile.Setup,
		signingprofilepermission.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_signer creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_signer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		signingjob.SetupGated,
		signingprofile.SetupGated,
		signingprofilepermission.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
