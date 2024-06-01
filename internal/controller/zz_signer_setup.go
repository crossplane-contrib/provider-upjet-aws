// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	signingjob "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/signer/signingjob"
	signingprofile "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/signer/signingprofile"
	signingprofilepermission "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/signer/signingprofilepermission"
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
