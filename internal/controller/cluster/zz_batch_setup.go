// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	computeenvironment "github.com/upbound/provider-aws/internal/controller/batch/computeenvironment"
	jobdefinition "github.com/upbound/provider-aws/internal/controller/batch/jobdefinition"
	jobqueue "github.com/upbound/provider-aws/internal/controller/batch/jobqueue"
	schedulingpolicy "github.com/upbound/provider-aws/internal/controller/batch/schedulingpolicy"
)

// Setup_batch creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_batch(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		computeenvironment.Setup,
		jobdefinition.Setup,
		jobqueue.Setup,
		schedulingpolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
