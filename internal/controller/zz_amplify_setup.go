/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	app "github.com/upbound/provider-aws/internal/controller/amplify/app"
	backendenvironment "github.com/upbound/provider-aws/internal/controller/amplify/backendenvironment"
	branch "github.com/upbound/provider-aws/internal/controller/amplify/branch"
	webhook "github.com/upbound/provider-aws/internal/controller/amplify/webhook"
)

// Setup_amplify creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_amplify(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		app.Setup,
		backendenvironment.Setup,
		branch.Setup,
		webhook.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
