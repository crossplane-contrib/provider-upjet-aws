// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accountsettingdefault "github.com/upbound/provider-aws/internal/controller/namespaced/ecs/accountsettingdefault"
	capacityprovider "github.com/upbound/provider-aws/internal/controller/namespaced/ecs/capacityprovider"
	cluster "github.com/upbound/provider-aws/internal/controller/namespaced/ecs/cluster"
	clustercapacityproviders "github.com/upbound/provider-aws/internal/controller/namespaced/ecs/clustercapacityproviders"
	service "github.com/upbound/provider-aws/internal/controller/namespaced/ecs/service"
	taskdefinition "github.com/upbound/provider-aws/internal/controller/namespaced/ecs/taskdefinition"
)

// Setup_ecs creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ecs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accountsettingdefault.Setup,
		capacityprovider.Setup,
		cluster.Setup,
		clustercapacityproviders.Setup,
		service.Setup,
		taskdefinition.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_ecs creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_ecs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accountsettingdefault.SetupGated,
		capacityprovider.SetupGated,
		cluster.SetupGated,
		clustercapacityproviders.SetupGated,
		service.SetupGated,
		taskdefinition.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
