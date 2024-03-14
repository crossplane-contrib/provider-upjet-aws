package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	account "github.com/upbound/provider-aws/internal/controller/securityhub/account"
	actiontarget "github.com/upbound/provider-aws/internal/controller/securityhub/actiontarget"
	findingaggregator "github.com/upbound/provider-aws/internal/controller/securityhub/findingaggregator"
	insight "github.com/upbound/provider-aws/internal/controller/securityhub/insight"
	inviteaccepter "github.com/upbound/provider-aws/internal/controller/securityhub/inviteaccepter"
	member "github.com/upbound/provider-aws/internal/controller/securityhub/member"
	productsubscription "github.com/upbound/provider-aws/internal/controller/securityhub/productsubscription"
	standardssubscription "github.com/upbound/provider-aws/internal/controller/securityhub/standardssubscription"
)

// Setup_securityhub creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_securityhub(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		actiontarget.Setup,
		findingaggregator.Setup,
		insight.Setup,
		inviteaccepter.Setup,
		member.Setup,
		productsubscription.Setup,
		standardssubscription.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
