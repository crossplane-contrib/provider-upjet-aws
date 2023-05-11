/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	stream "github.com/upbound/provider-aws/internal/controller/kinesisvideo/stream"
)

// Setup_kinesisvideo creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_kinesisvideo(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		stream.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
