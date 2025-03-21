// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	approvalruletemplate "github.com/upbound/provider-aws/internal/controller/cluster/codecommit/approvalruletemplate"
	approvalruletemplateassociation "github.com/upbound/provider-aws/internal/controller/cluster/codecommit/approvalruletemplateassociation"
	repository "github.com/upbound/provider-aws/internal/controller/cluster/codecommit/repository"
	trigger "github.com/upbound/provider-aws/internal/controller/cluster/codecommit/trigger"
)

// Setup_codecommit creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_codecommit(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		approvalruletemplate.Setup,
		approvalruletemplateassociation.Setup,
		repository.Setup,
		trigger.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
