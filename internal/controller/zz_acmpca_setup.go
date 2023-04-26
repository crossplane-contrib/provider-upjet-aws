/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	certificate "github.com/upbound/provider-aws/internal/controller/acmpca/certificate"
	certificateauthority "github.com/upbound/provider-aws/internal/controller/acmpca/certificateauthority"
	certificateauthoritycertificate "github.com/upbound/provider-aws/internal/controller/acmpca/certificateauthoritycertificate"
	permission "github.com/upbound/provider-aws/internal/controller/acmpca/permission"
	policy "github.com/upbound/provider-aws/internal/controller/acmpca/policy"
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
