/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	appmonitor "github.com/upbound/provider-aws/internal/controller/rum/appmonitor"
	metricsdestination "github.com/upbound/provider-aws/internal/controller/rum/metricsdestination"
)

// Setup_rum creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_rum(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appmonitor.Setup,
		metricsdestination.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
