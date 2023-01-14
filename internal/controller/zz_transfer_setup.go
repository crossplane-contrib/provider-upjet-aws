/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	server "github.com/upbound/provider-aws/internal/controller/transfer/server"
	sshkey "github.com/upbound/provider-aws/internal/controller/transfer/sshkey"
	user "github.com/upbound/provider-aws/internal/controller/transfer/user"
)

// Setup_transfer creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_transfer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		server.Setup,
		sshkey.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
