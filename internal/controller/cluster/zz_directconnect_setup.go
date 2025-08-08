// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	bgppeer "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/bgppeer"
	connection "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/connection"
	connectionassociation "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/connectionassociation"
	gateway "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/gateway"
	gatewayassociation "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/gatewayassociation"
	gatewayassociationproposal "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/gatewayassociationproposal"
	hostedprivatevirtualinterface "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/hostedprivatevirtualinterface"
	hostedprivatevirtualinterfaceaccepter "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/hostedprivatevirtualinterfaceaccepter"
	hostedpublicvirtualinterface "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/hostedpublicvirtualinterface"
	hostedpublicvirtualinterfaceaccepter "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/hostedpublicvirtualinterfaceaccepter"
	hostedtransitvirtualinterface "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/hostedtransitvirtualinterface"
	hostedtransitvirtualinterfaceaccepter "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/hostedtransitvirtualinterfaceaccepter"
	lag "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/lag"
	privatevirtualinterface "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/privatevirtualinterface"
	publicvirtualinterface "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/publicvirtualinterface"
	transitvirtualinterface "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/transitvirtualinterface"
)

// Setup_directconnect creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_directconnect(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bgppeer.Setup,
		connection.Setup,
		connectionassociation.Setup,
		gateway.Setup,
		gatewayassociation.Setup,
		gatewayassociationproposal.Setup,
		hostedprivatevirtualinterface.Setup,
		hostedprivatevirtualinterfaceaccepter.Setup,
		hostedpublicvirtualinterface.Setup,
		hostedpublicvirtualinterfaceaccepter.Setup,
		hostedtransitvirtualinterface.Setup,
		hostedtransitvirtualinterfaceaccepter.Setup,
		lag.Setup,
		privatevirtualinterface.Setup,
		publicvirtualinterface.Setup,
		transitvirtualinterface.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_directconnect creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_directconnect(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bgppeer.SetupGated,
		connection.SetupGated,
		connectionassociation.SetupGated,
		gateway.SetupGated,
		gatewayassociation.SetupGated,
		gatewayassociationproposal.SetupGated,
		hostedprivatevirtualinterface.SetupGated,
		hostedprivatevirtualinterfaceaccepter.SetupGated,
		hostedpublicvirtualinterface.SetupGated,
		hostedpublicvirtualinterfaceaccepter.SetupGated,
		hostedtransitvirtualinterface.SetupGated,
		hostedtransitvirtualinterfaceaccepter.SetupGated,
		lag.SetupGated,
		privatevirtualinterface.SetupGated,
		publicvirtualinterface.SetupGated,
		transitvirtualinterface.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
