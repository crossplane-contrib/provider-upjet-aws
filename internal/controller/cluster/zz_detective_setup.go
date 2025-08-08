// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	graph "github.com/upbound/provider-aws/internal/controller/cluster/detective/graph"
	invitationaccepter "github.com/upbound/provider-aws/internal/controller/cluster/detective/invitationaccepter"
	member "github.com/upbound/provider-aws/internal/controller/cluster/detective/member"
)

// Setup_detective creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_detective(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		graph.Setup,
		invitationaccepter.Setup,
		member.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_detective creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_detective(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		graph.SetupGated,
		invitationaccepter.SetupGated,
		member.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
