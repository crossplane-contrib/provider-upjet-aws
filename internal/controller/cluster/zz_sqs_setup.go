// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	queue "github.com/upbound/provider-aws/internal/controller/sqs/queue"
	queuepolicy "github.com/upbound/provider-aws/internal/controller/sqs/queuepolicy"
	queueredriveallowpolicy "github.com/upbound/provider-aws/internal/controller/sqs/queueredriveallowpolicy"
	queueredrivepolicy "github.com/upbound/provider-aws/internal/controller/sqs/queueredrivepolicy"
)

// Setup_sqs creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_sqs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		queue.Setup,
		queuepolicy.Setup,
		queueredriveallowpolicy.Setup,
		queueredrivepolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
