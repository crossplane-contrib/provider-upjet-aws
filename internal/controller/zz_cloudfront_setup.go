/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	cachepolicy "github.com/upbound/provider-aws/internal/controller/cloudfront/cachepolicy"
	distribution "github.com/upbound/provider-aws/internal/controller/cloudfront/distribution"
	fieldlevelencryptionconfig "github.com/upbound/provider-aws/internal/controller/cloudfront/fieldlevelencryptionconfig"
	fieldlevelencryptionprofile "github.com/upbound/provider-aws/internal/controller/cloudfront/fieldlevelencryptionprofile"
	function "github.com/upbound/provider-aws/internal/controller/cloudfront/function"
	keygroup "github.com/upbound/provider-aws/internal/controller/cloudfront/keygroup"
	monitoringsubscription "github.com/upbound/provider-aws/internal/controller/cloudfront/monitoringsubscription"
	originaccesscontrol "github.com/upbound/provider-aws/internal/controller/cloudfront/originaccesscontrol"
	originaccessidentity "github.com/upbound/provider-aws/internal/controller/cloudfront/originaccessidentity"
	originrequestpolicy "github.com/upbound/provider-aws/internal/controller/cloudfront/originrequestpolicy"
	publickey "github.com/upbound/provider-aws/internal/controller/cloudfront/publickey"
	realtimelogconfig "github.com/upbound/provider-aws/internal/controller/cloudfront/realtimelogconfig"
	responseheaderspolicy "github.com/upbound/provider-aws/internal/controller/cloudfront/responseheaderspolicy"
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
