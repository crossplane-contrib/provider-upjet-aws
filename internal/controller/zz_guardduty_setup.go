/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	detector "github.com/upbound/provider-aws/internal/controller/guardduty/detector"
	filter "github.com/upbound/provider-aws/internal/controller/guardduty/filter"
	member "github.com/upbound/provider-aws/internal/controller/guardduty/member"
)

// Setup_guardduty creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_guardduty(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		detector.Setup,
		filter.Setup,
		member.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
