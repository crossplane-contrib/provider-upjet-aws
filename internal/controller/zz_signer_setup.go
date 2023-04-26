/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	signingjob "github.com/upbound/provider-aws/internal/controller/signer/signingjob"
	signingprofile "github.com/upbound/provider-aws/internal/controller/signer/signingprofile"
	signingprofilepermission "github.com/upbound/provider-aws/internal/controller/signer/signingprofilepermission"
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
