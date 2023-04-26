/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	connection "github.com/upbound/provider-aws/internal/controller/codestarconnections/connection"
	host "github.com/upbound/provider-aws/internal/controller/codestarconnections/host"
)

// Setup_codestarconnections creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_codestarconnections(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connection.Setup,
		host.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
