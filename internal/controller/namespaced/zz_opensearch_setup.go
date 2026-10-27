// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	authorizevpcendpointaccess "github.com/upbound/provider-aws/v2/internal/controller/namespaced/opensearch/authorizevpcendpointaccess"
	domain "github.com/upbound/provider-aws/v2/internal/controller/namespaced/opensearch/domain"
	domainpolicy "github.com/upbound/provider-aws/v2/internal/controller/namespaced/opensearch/domainpolicy"
	domainsamloptions "github.com/upbound/provider-aws/v2/internal/controller/namespaced/opensearch/domainsamloptions"
	vpcendpoint "github.com/upbound/provider-aws/v2/internal/controller/namespaced/opensearch/vpcendpoint"
)

// Setup_opensearch creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_opensearch(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authorizevpcendpointaccess.Setup,
		domain.Setup,
		domainpolicy.Setup,
		domainsamloptions.Setup,
		vpcendpoint.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_opensearch creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_opensearch(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authorizevpcendpointaccess.SetupGated,
		domain.SetupGated,
		domainpolicy.SetupGated,
		domainsamloptions.SetupGated,
		vpcendpoint.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_opensearch registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_opensearch(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		authorizevpcendpointaccess.SetupWebhookWithManager,
		domain.SetupWebhookWithManager,
		domainpolicy.SetupWebhookWithManager,
		domainsamloptions.SetupWebhookWithManager,
		vpcendpoint.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
