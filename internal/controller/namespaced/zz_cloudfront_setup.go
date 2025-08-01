// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cachepolicy "github.com/upbound/provider-aws/internal/controller/namespaced/cloudfront/cachepolicy"
	distribution "github.com/upbound/provider-aws/internal/controller/namespaced/cloudfront/distribution"
	fieldlevelencryptionconfig "github.com/upbound/provider-aws/internal/controller/namespaced/cloudfront/fieldlevelencryptionconfig"
	fieldlevelencryptionprofile "github.com/upbound/provider-aws/internal/controller/namespaced/cloudfront/fieldlevelencryptionprofile"
	function "github.com/upbound/provider-aws/internal/controller/namespaced/cloudfront/function"
	keygroup "github.com/upbound/provider-aws/internal/controller/namespaced/cloudfront/keygroup"
	monitoringsubscription "github.com/upbound/provider-aws/internal/controller/namespaced/cloudfront/monitoringsubscription"
	originaccesscontrol "github.com/upbound/provider-aws/internal/controller/namespaced/cloudfront/originaccesscontrol"
	originaccessidentity "github.com/upbound/provider-aws/internal/controller/namespaced/cloudfront/originaccessidentity"
	originrequestpolicy "github.com/upbound/provider-aws/internal/controller/namespaced/cloudfront/originrequestpolicy"
	publickey "github.com/upbound/provider-aws/internal/controller/namespaced/cloudfront/publickey"
	realtimelogconfig "github.com/upbound/provider-aws/internal/controller/namespaced/cloudfront/realtimelogconfig"
	responseheaderspolicy "github.com/upbound/provider-aws/internal/controller/namespaced/cloudfront/responseheaderspolicy"
)

// Setup_cloudfront creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudfront(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cachepolicy.Setup,
		distribution.Setup,
		fieldlevelencryptionconfig.Setup,
		fieldlevelencryptionprofile.Setup,
		function.Setup,
		keygroup.Setup,
		monitoringsubscription.Setup,
		originaccesscontrol.Setup,
		originaccessidentity.Setup,
		originrequestpolicy.Setup,
		publickey.Setup,
		realtimelogconfig.Setup,
		responseheaderspolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cloudfront creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cloudfront(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cachepolicy.SetupGated,
		distribution.SetupGated,
		fieldlevelencryptionconfig.SetupGated,
		fieldlevelencryptionprofile.SetupGated,
		function.SetupGated,
		keygroup.SetupGated,
		monitoringsubscription.SetupGated,
		originaccesscontrol.SetupGated,
		originaccessidentity.SetupGated,
		originrequestpolicy.SetupGated,
		publickey.SetupGated,
		realtimelogconfig.SetupGated,
		responseheaderspolicy.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
