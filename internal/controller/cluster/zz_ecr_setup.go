// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	lifecyclepolicy "github.com/upbound/provider-aws/internal/controller/cluster/ecr/lifecyclepolicy"
	pullthroughcacherule "github.com/upbound/provider-aws/internal/controller/cluster/ecr/pullthroughcacherule"
	registrypolicy "github.com/upbound/provider-aws/internal/controller/cluster/ecr/registrypolicy"
	registryscanningconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/ecr/registryscanningconfiguration"
	replicationconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/ecr/replicationconfiguration"
	repository "github.com/upbound/provider-aws/internal/controller/cluster/ecr/repository"
	repositorypolicy "github.com/upbound/provider-aws/internal/controller/cluster/ecr/repositorypolicy"
)

// Setup_ecr creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ecr(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		lifecyclepolicy.Setup,
		pullthroughcacherule.Setup,
		registrypolicy.Setup,
		registryscanningconfiguration.Setup,
		replicationconfiguration.Setup,
		repository.Setup,
		repositorypolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_ecr creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_ecr(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		lifecyclepolicy.SetupGated,
		pullthroughcacherule.SetupGated,
		registrypolicy.SetupGated,
		registryscanningconfiguration.SetupGated,
		replicationconfiguration.SetupGated,
		repository.SetupGated,
		repositorypolicy.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
