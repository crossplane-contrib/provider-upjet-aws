/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	ledger "github.com/upbound/provider-aws/internal/controller/qldb/ledger"
	stream "github.com/upbound/provider-aws/internal/controller/qldb/stream"
)

// Setup_qldb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_qldb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		ledger.Setup,
		stream.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
