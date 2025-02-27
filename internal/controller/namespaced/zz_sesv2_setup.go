// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	configurationset "github.com/upbound/provider-aws/internal/controller/sesv2/configurationset"
	configurationseteventdestination "github.com/upbound/provider-aws/internal/controller/sesv2/configurationseteventdestination"
	dedicatedippool "github.com/upbound/provider-aws/internal/controller/sesv2/dedicatedippool"
	emailidentity "github.com/upbound/provider-aws/internal/controller/sesv2/emailidentity"
	emailidentityfeedbackattributes "github.com/upbound/provider-aws/internal/controller/sesv2/emailidentityfeedbackattributes"
	emailidentitymailfromattributes "github.com/upbound/provider-aws/internal/controller/sesv2/emailidentitymailfromattributes"
)

// Setup_sesv2 creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_sesv2(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		configurationset.Setup,
		configurationseteventdestination.Setup,
		dedicatedippool.Setup,
		emailidentity.Setup,
		emailidentityfeedbackattributes.Setup,
		emailidentitymailfromattributes.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
