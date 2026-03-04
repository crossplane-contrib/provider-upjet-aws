// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	agentruntime "github.com/upbound/provider-aws/v2/internal/controller/cluster/bedrockagentcore/agentruntime"
	agentruntimeendpoint "github.com/upbound/provider-aws/v2/internal/controller/cluster/bedrockagentcore/agentruntimeendpoint"
	apikeycredentialprovider "github.com/upbound/provider-aws/v2/internal/controller/cluster/bedrockagentcore/apikeycredentialprovider"
	browser "github.com/upbound/provider-aws/v2/internal/controller/cluster/bedrockagentcore/browser"
	codeinterpreter "github.com/upbound/provider-aws/v2/internal/controller/cluster/bedrockagentcore/codeinterpreter"
	gateway "github.com/upbound/provider-aws/v2/internal/controller/cluster/bedrockagentcore/gateway"
	gatewaytarget "github.com/upbound/provider-aws/v2/internal/controller/cluster/bedrockagentcore/gatewaytarget"
	memory "github.com/upbound/provider-aws/v2/internal/controller/cluster/bedrockagentcore/memory"
	memorystrategy "github.com/upbound/provider-aws/v2/internal/controller/cluster/bedrockagentcore/memorystrategy"
	oauth2credentialprovider "github.com/upbound/provider-aws/v2/internal/controller/cluster/bedrockagentcore/oauth2credentialprovider"
	tokenvaultcmk "github.com/upbound/provider-aws/v2/internal/controller/cluster/bedrockagentcore/tokenvaultcmk"
	workloadidentity "github.com/upbound/provider-aws/v2/internal/controller/cluster/bedrockagentcore/workloadidentity"
)

// Setup_bedrockagentcore creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_bedrockagentcore(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		agentruntime.Setup,
		agentruntimeendpoint.Setup,
		apikeycredentialprovider.Setup,
		browser.Setup,
		codeinterpreter.Setup,
		gateway.Setup,
		gatewaytarget.Setup,
		memory.Setup,
		memorystrategy.Setup,
		oauth2credentialprovider.Setup,
		tokenvaultcmk.Setup,
		workloadidentity.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_bedrockagentcore creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_bedrockagentcore(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		agentruntime.SetupGated,
		agentruntimeendpoint.SetupGated,
		apikeycredentialprovider.SetupGated,
		browser.SetupGated,
		codeinterpreter.SetupGated,
		gateway.SetupGated,
		gatewaytarget.SetupGated,
		memory.SetupGated,
		memorystrategy.SetupGated,
		oauth2credentialprovider.SetupGated,
		tokenvaultcmk.SetupGated,
		workloadidentity.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
