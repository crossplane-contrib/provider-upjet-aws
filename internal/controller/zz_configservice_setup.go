/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	awsconfigurationrecorderstatus "github.com/upbound/provider-aws/internal/controller/configservice/awsconfigurationrecorderstatus"
	configrule "github.com/upbound/provider-aws/internal/controller/configservice/configrule"
	configurationaggregator "github.com/upbound/provider-aws/internal/controller/configservice/configurationaggregator"
	configurationrecorder "github.com/upbound/provider-aws/internal/controller/configservice/configurationrecorder"
	conformancepack "github.com/upbound/provider-aws/internal/controller/configservice/conformancepack"
	deliverychannel "github.com/upbound/provider-aws/internal/controller/configservice/deliverychannel"
	remediationconfiguration "github.com/upbound/provider-aws/internal/controller/configservice/remediationconfiguration"
)

// Setup_configservice creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_configservice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		awsconfigurationrecorderstatus.Setup,
		configrule.Setup,
		configurationaggregator.Setup,
		configurationrecorder.Setup,
		conformancepack.Setup,
		deliverychannel.Setup,
		remediationconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
