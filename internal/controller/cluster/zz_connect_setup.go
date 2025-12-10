// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	botassociation "github.com/upbound/provider-aws/v2/internal/controller/cluster/connect/botassociation"
	contactflow "github.com/upbound/provider-aws/v2/internal/controller/cluster/connect/contactflow"
	contactflowmodule "github.com/upbound/provider-aws/v2/internal/controller/cluster/connect/contactflowmodule"
	hoursofoperation "github.com/upbound/provider-aws/v2/internal/controller/cluster/connect/hoursofoperation"
	instance "github.com/upbound/provider-aws/v2/internal/controller/cluster/connect/instance"
	instancestorageconfig "github.com/upbound/provider-aws/v2/internal/controller/cluster/connect/instancestorageconfig"
	lambdafunctionassociation "github.com/upbound/provider-aws/v2/internal/controller/cluster/connect/lambdafunctionassociation"
	phonenumber "github.com/upbound/provider-aws/v2/internal/controller/cluster/connect/phonenumber"
	queue "github.com/upbound/provider-aws/v2/internal/controller/cluster/connect/queue"
	quickconnect "github.com/upbound/provider-aws/v2/internal/controller/cluster/connect/quickconnect"
	routingprofile "github.com/upbound/provider-aws/v2/internal/controller/cluster/connect/routingprofile"
	securityprofile "github.com/upbound/provider-aws/v2/internal/controller/cluster/connect/securityprofile"
	user "github.com/upbound/provider-aws/v2/internal/controller/cluster/connect/user"
	userhierarchystructure "github.com/upbound/provider-aws/v2/internal/controller/cluster/connect/userhierarchystructure"
	vocabulary "github.com/upbound/provider-aws/v2/internal/controller/cluster/connect/vocabulary"
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

// SetupGated_connect creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_connect(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		botassociation.SetupGated,
		contactflow.SetupGated,
		contactflowmodule.SetupGated,
		hoursofoperation.SetupGated,
		instance.SetupGated,
		instancestorageconfig.SetupGated,
		lambdafunctionassociation.SetupGated,
		phonenumber.SetupGated,
		queue.SetupGated,
		quickconnect.SetupGated,
		routingprofile.SetupGated,
		securityprofile.SetupGated,
		user.SetupGated,
		userhierarchystructure.SetupGated,
		vocabulary.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
