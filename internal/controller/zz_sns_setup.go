/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	topic "github.com/upbound/provider-aws/internal/controller/sns/topic"
)

// Setup_sns creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_sns(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		topic.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
