// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	domain "github.com/upbound/provider-aws/internal/controller/namespaced/elasticsearch/domain"
	domainpolicy "github.com/upbound/provider-aws/internal/controller/namespaced/elasticsearch/domainpolicy"
	domainsamloptions "github.com/upbound/provider-aws/internal/controller/namespaced/elasticsearch/domainsamloptions"
)

// Setup_elasticsearch creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_elasticsearch(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		domain.Setup,
		domainpolicy.Setup,
		domainsamloptions.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_elasticsearch creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_elasticsearch(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		domain.SetupGated,
		domainpolicy.SetupGated,
		domainsamloptions.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
