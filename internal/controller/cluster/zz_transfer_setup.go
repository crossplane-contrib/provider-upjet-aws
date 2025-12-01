// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	connector "github.com/upbound/provider-aws/v2/internal/controller/cluster/transfer/connector"
	server "github.com/upbound/provider-aws/v2/internal/controller/cluster/transfer/server"
	sshkey "github.com/upbound/provider-aws/v2/internal/controller/cluster/transfer/sshkey"
	tag "github.com/upbound/provider-aws/v2/internal/controller/cluster/transfer/tag"
	user "github.com/upbound/provider-aws/v2/internal/controller/cluster/transfer/user"
	workflow "github.com/upbound/provider-aws/v2/internal/controller/cluster/transfer/workflow"
)

// Setup_transfer creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_transfer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connector.Setup,
		server.Setup,
		sshkey.Setup,
		tag.Setup,
		user.Setup,
		workflow.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_transfer creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_transfer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connector.SetupGated,
		server.SetupGated,
		sshkey.SetupGated,
		tag.SetupGated,
		user.SetupGated,
		workflow.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
