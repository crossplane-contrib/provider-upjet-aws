// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	eventdatastore "github.com/upbound/provider-aws/internal/controller/cloudtrail/eventdatastore"
	trail "github.com/upbound/provider-aws/internal/controller/cloudtrail/trail"
)

// Setup_cloudtrail creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudtrail(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		eventdatastore.Setup,
		trail.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
