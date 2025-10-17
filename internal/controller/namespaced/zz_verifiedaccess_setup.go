// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	endpoint "github.com/upbound/provider-aws/internal/controller/namespaced/verifiedaccess/endpoint"
	group "github.com/upbound/provider-aws/internal/controller/namespaced/verifiedaccess/group"
	instance "github.com/upbound/provider-aws/internal/controller/namespaced/verifiedaccess/instance"
	instanceloggingconfiguration "github.com/upbound/provider-aws/internal/controller/namespaced/verifiedaccess/instanceloggingconfiguration"
	instancetrustproviderattachment "github.com/upbound/provider-aws/internal/controller/namespaced/verifiedaccess/instancetrustproviderattachment"
	trustprovider "github.com/upbound/provider-aws/internal/controller/namespaced/verifiedaccess/trustprovider"
)

// Setup_verifiedaccess creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_verifiedaccess(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		endpoint.Setup,
		group.Setup,
		instance.Setup,
		instanceloggingconfiguration.Setup,
		instancetrustproviderattachment.Setup,
		trustprovider.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_verifiedaccess creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_verifiedaccess(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		endpoint.SetupGated,
		group.SetupGated,
		instance.SetupGated,
		instanceloggingconfiguration.SetupGated,
		instancetrustproviderattachment.SetupGated,
		trustprovider.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
