// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	stream "github.com/upbound/provider-aws/internal/controller/cluster/kinesis/stream"
	streamconsumer "github.com/upbound/provider-aws/internal/controller/cluster/kinesis/streamconsumer"
)

// Setup_kinesis creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_kinesis(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		stream.Setup,
		streamconsumer.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
