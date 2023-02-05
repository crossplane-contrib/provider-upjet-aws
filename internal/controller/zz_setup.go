/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	providerconfig "github.com/dkb-bank/official-provider-aws/internal/controller/providerconfig"
	activereceiptruleset "github.com/dkb-bank/official-provider-aws/internal/controller/ses/activereceiptruleset"
	configurationset "github.com/dkb-bank/official-provider-aws/internal/controller/ses/configurationset"
	domaindkim "github.com/dkb-bank/official-provider-aws/internal/controller/ses/domaindkim"
	domainidentity "github.com/dkb-bank/official-provider-aws/internal/controller/ses/domainidentity"
	domainmailfrom "github.com/dkb-bank/official-provider-aws/internal/controller/ses/domainmailfrom"
	emailidentity "github.com/dkb-bank/official-provider-aws/internal/controller/ses/emailidentity"
	eventdestination "github.com/dkb-bank/official-provider-aws/internal/controller/ses/eventdestination"
	identitynotificationtopic "github.com/dkb-bank/official-provider-aws/internal/controller/ses/identitynotificationtopic"
	identitypolicy "github.com/dkb-bank/official-provider-aws/internal/controller/ses/identitypolicy"
	receiptfilter "github.com/dkb-bank/official-provider-aws/internal/controller/ses/receiptfilter"
	receiptrule "github.com/dkb-bank/official-provider-aws/internal/controller/ses/receiptrule"
	receiptruleset "github.com/dkb-bank/official-provider-aws/internal/controller/ses/receiptruleset"
	template "github.com/dkb-bank/official-provider-aws/internal/controller/ses/template"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		activereceiptruleset.Setup,
		configurationset.Setup,
		domaindkim.Setup,
		domainidentity.Setup,
		domainmailfrom.Setup,
		emailidentity.Setup,
		eventdestination.Setup,
		identitynotificationtopic.Setup,
		identitypolicy.Setup,
		receiptfilter.Setup,
		receiptrule.Setup,
		receiptruleset.Setup,
		template.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
