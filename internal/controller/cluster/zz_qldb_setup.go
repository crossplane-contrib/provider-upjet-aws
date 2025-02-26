// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

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
