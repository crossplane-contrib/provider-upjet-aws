// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/upbound/provider-aws/internal/controller/emr/cluster"
	instancefleet "github.com/upbound/provider-aws/internal/controller/emr/instancefleet"
	instancegroup "github.com/upbound/provider-aws/internal/controller/emr/instancegroup"
	securityconfiguration "github.com/upbound/provider-aws/internal/controller/emr/securityconfiguration"
)

// Setup_emr creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_emr(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		instancefleet.Setup,
		instancegroup.Setup,
		securityconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
