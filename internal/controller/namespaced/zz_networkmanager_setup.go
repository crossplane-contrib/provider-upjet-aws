// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	attachmentaccepter "github.com/upbound/provider-aws/v2/internal/controller/namespaced/networkmanager/attachmentaccepter"
	connectattachment "github.com/upbound/provider-aws/v2/internal/controller/namespaced/networkmanager/connectattachment"
	connection "github.com/upbound/provider-aws/v2/internal/controller/namespaced/networkmanager/connection"
	corenetwork "github.com/upbound/provider-aws/v2/internal/controller/namespaced/networkmanager/corenetwork"
	customergatewayassociation "github.com/upbound/provider-aws/v2/internal/controller/namespaced/networkmanager/customergatewayassociation"
	device "github.com/upbound/provider-aws/v2/internal/controller/namespaced/networkmanager/device"
	globalnetwork "github.com/upbound/provider-aws/v2/internal/controller/namespaced/networkmanager/globalnetwork"
	link "github.com/upbound/provider-aws/v2/internal/controller/namespaced/networkmanager/link"
	linkassociation "github.com/upbound/provider-aws/v2/internal/controller/namespaced/networkmanager/linkassociation"
	site "github.com/upbound/provider-aws/v2/internal/controller/namespaced/networkmanager/site"
	transitgatewayconnectpeerassociation "github.com/upbound/provider-aws/v2/internal/controller/namespaced/networkmanager/transitgatewayconnectpeerassociation"
	transitgatewayregistration "github.com/upbound/provider-aws/v2/internal/controller/namespaced/networkmanager/transitgatewayregistration"
	vpcattachment "github.com/upbound/provider-aws/v2/internal/controller/namespaced/networkmanager/vpcattachment"
)

// Setup_networkmanager creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_networkmanager(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		attachmentaccepter.Setup,
		connectattachment.Setup,
		connection.Setup,
		corenetwork.Setup,
		customergatewayassociation.Setup,
		device.Setup,
		globalnetwork.Setup,
		link.Setup,
		linkassociation.Setup,
		site.Setup,
		transitgatewayconnectpeerassociation.Setup,
		transitgatewayregistration.Setup,
		vpcattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_networkmanager creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_networkmanager(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		attachmentaccepter.SetupGated,
		connectattachment.SetupGated,
		connection.SetupGated,
		corenetwork.SetupGated,
		customergatewayassociation.SetupGated,
		device.SetupGated,
		globalnetwork.SetupGated,
		link.SetupGated,
		linkassociation.SetupGated,
		site.SetupGated,
		transitgatewayconnectpeerassociation.SetupGated,
		transitgatewayregistration.SetupGated,
		vpcattachment.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
