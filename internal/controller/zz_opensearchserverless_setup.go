// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	accesspolicy "github.com/upbound/provider-aws/internal/controller/opensearchserverless/accesspolicy"
	collection "github.com/upbound/provider-aws/internal/controller/opensearchserverless/collection"
	lifecyclepolicy "github.com/upbound/provider-aws/internal/controller/opensearchserverless/lifecyclepolicy"
	securityconfig "github.com/upbound/provider-aws/internal/controller/opensearchserverless/securityconfig"
	securitypolicy "github.com/upbound/provider-aws/internal/controller/opensearchserverless/securitypolicy"
	vpcendpoint "github.com/upbound/provider-aws/internal/controller/opensearchserverless/vpcendpoint"
)

// Setup_opensearchserverless creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_opensearchserverless(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesspolicy.Setup,
		collection.Setup,
		lifecyclepolicy.Setup,
		securityconfig.Setup,
		securitypolicy.Setup,
		vpcendpoint.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
