// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	platformapplication "github.com/upbound/provider-aws/internal/controller/namespaced/sns/platformapplication"
	smspreferences "github.com/upbound/provider-aws/internal/controller/namespaced/sns/smspreferences"
	topic "github.com/upbound/provider-aws/internal/controller/namespaced/sns/topic"
	topicpolicy "github.com/upbound/provider-aws/internal/controller/namespaced/sns/topicpolicy"
	topicsubscription "github.com/upbound/provider-aws/internal/controller/namespaced/sns/topicsubscription"
)

// Setup_sns creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_sns(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		platformapplication.Setup,
		smspreferences.Setup,
		topic.Setup,
		topicpolicy.Setup,
		topicsubscription.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
