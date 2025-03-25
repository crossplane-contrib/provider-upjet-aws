// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	botassociation "github.com/upbound/provider-aws/internal/controller/namespaced/connect/botassociation"
	contactflow "github.com/upbound/provider-aws/internal/controller/namespaced/connect/contactflow"
	contactflowmodule "github.com/upbound/provider-aws/internal/controller/namespaced/connect/contactflowmodule"
	hoursofoperation "github.com/upbound/provider-aws/internal/controller/namespaced/connect/hoursofoperation"
	instance "github.com/upbound/provider-aws/internal/controller/namespaced/connect/instance"
	instancestorageconfig "github.com/upbound/provider-aws/internal/controller/namespaced/connect/instancestorageconfig"
	lambdafunctionassociation "github.com/upbound/provider-aws/internal/controller/namespaced/connect/lambdafunctionassociation"
	phonenumber "github.com/upbound/provider-aws/internal/controller/namespaced/connect/phonenumber"
	queue "github.com/upbound/provider-aws/internal/controller/namespaced/connect/queue"
	quickconnect "github.com/upbound/provider-aws/internal/controller/namespaced/connect/quickconnect"
	routingprofile "github.com/upbound/provider-aws/internal/controller/namespaced/connect/routingprofile"
	securityprofile "github.com/upbound/provider-aws/internal/controller/namespaced/connect/securityprofile"
	user "github.com/upbound/provider-aws/internal/controller/namespaced/connect/user"
	userhierarchystructure "github.com/upbound/provider-aws/internal/controller/namespaced/connect/userhierarchystructure"
	vocabulary "github.com/upbound/provider-aws/internal/controller/namespaced/connect/vocabulary"
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
