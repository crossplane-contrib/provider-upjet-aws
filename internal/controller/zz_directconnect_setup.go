// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	bgppeer "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/directconnect/bgppeer"
	connection "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/directconnect/connection"
	connectionassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/directconnect/connectionassociation"
	gateway "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/directconnect/gateway"
	gatewayassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/directconnect/gatewayassociation"
	gatewayassociationproposal "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/directconnect/gatewayassociationproposal"
	hostedprivatevirtualinterface "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/directconnect/hostedprivatevirtualinterface"
	hostedprivatevirtualinterfaceaccepter "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/directconnect/hostedprivatevirtualinterfaceaccepter"
	hostedpublicvirtualinterface "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/directconnect/hostedpublicvirtualinterface"
	hostedpublicvirtualinterfaceaccepter "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/directconnect/hostedpublicvirtualinterfaceaccepter"
	hostedtransitvirtualinterface "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/directconnect/hostedtransitvirtualinterface"
	hostedtransitvirtualinterfaceaccepter "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/directconnect/hostedtransitvirtualinterfaceaccepter"
	lag "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/directconnect/lag"
	privatevirtualinterface "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/directconnect/privatevirtualinterface"
	publicvirtualinterface "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/directconnect/publicvirtualinterface"
	transitvirtualinterface "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/directconnect/transitvirtualinterface"
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
