// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	awsconfigurationrecorderstatus "github.com/upbound/provider-aws/internal/controller/cluster/configservice/awsconfigurationrecorderstatus"
	configrule "github.com/upbound/provider-aws/internal/controller/cluster/configservice/configrule"
	configurationaggregator "github.com/upbound/provider-aws/internal/controller/cluster/configservice/configurationaggregator"
	configurationrecorder "github.com/upbound/provider-aws/internal/controller/cluster/configservice/configurationrecorder"
	conformancepack "github.com/upbound/provider-aws/internal/controller/cluster/configservice/conformancepack"
	deliverychannel "github.com/upbound/provider-aws/internal/controller/cluster/configservice/deliverychannel"
	remediationconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/configservice/remediationconfiguration"
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

// SetupGated_configservice creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_configservice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		awsconfigurationrecorderstatus.SetupGated,
		configrule.SetupGated,
		configurationaggregator.SetupGated,
		configurationrecorder.SetupGated,
		conformancepack.SetupGated,
		deliverychannel.SetupGated,
		remediationconfiguration.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
