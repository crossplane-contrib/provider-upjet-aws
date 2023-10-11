// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	botassociation "github.com/upbound/provider-aws/internal/controller/connect/botassociation"
	contactflow "github.com/upbound/provider-aws/internal/controller/connect/contactflow"
	contactflowmodule "github.com/upbound/provider-aws/internal/controller/connect/contactflowmodule"
	hoursofoperation "github.com/upbound/provider-aws/internal/controller/connect/hoursofoperation"
	instance "github.com/upbound/provider-aws/internal/controller/connect/instance"
	instancestorageconfig "github.com/upbound/provider-aws/internal/controller/connect/instancestorageconfig"
	lambdafunctionassociation "github.com/upbound/provider-aws/internal/controller/connect/lambdafunctionassociation"
	phonenumber "github.com/upbound/provider-aws/internal/controller/connect/phonenumber"
	queue "github.com/upbound/provider-aws/internal/controller/connect/queue"
	quickconnect "github.com/upbound/provider-aws/internal/controller/connect/quickconnect"
	routingprofile "github.com/upbound/provider-aws/internal/controller/connect/routingprofile"
	securityprofile "github.com/upbound/provider-aws/internal/controller/connect/securityprofile"
	user "github.com/upbound/provider-aws/internal/controller/connect/user"
	userhierarchystructure "github.com/upbound/provider-aws/internal/controller/connect/userhierarchystructure"
	vocabulary "github.com/upbound/provider-aws/internal/controller/connect/vocabulary"
)

// Setup_connect creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_connect(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		botassociation.Setup,
		contactflow.Setup,
		contactflowmodule.Setup,
		hoursofoperation.Setup,
		instance.Setup,
		instancestorageconfig.Setup,
		lambdafunctionassociation.Setup,
		phonenumber.Setup,
		queue.Setup,
		quickconnect.Setup,
		routingprofile.Setup,
		securityprofile.Setup,
		user.Setup,
		userhierarchystructure.Setup,
		vocabulary.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
