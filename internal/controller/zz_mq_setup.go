/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	broker "github.com/upbound/provider-aws/internal/controller/mq/broker"
	configuration "github.com/upbound/provider-aws/internal/controller/mq/configuration"
)

// Setup_mq creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_mq(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		broker.Setup,
		configuration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
