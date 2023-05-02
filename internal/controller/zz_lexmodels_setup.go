/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	bot "github.com/upbound/provider-aws/internal/controller/lexmodels/bot"
	botalias "github.com/upbound/provider-aws/internal/controller/lexmodels/botalias"
	intent "github.com/upbound/provider-aws/internal/controller/lexmodels/intent"
	slottype "github.com/upbound/provider-aws/internal/controller/lexmodels/slottype"
)

// Setup_lexmodels creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_lexmodels(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bot.Setup,
		botalias.Setup,
		intent.Setup,
		slottype.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
