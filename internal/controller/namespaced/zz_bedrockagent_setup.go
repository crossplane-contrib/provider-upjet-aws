// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	agent "github.com/upbound/provider-aws/v2/internal/controller/namespaced/bedrockagent/agent"
	agentknowledgebaseassociation "github.com/upbound/provider-aws/v2/internal/controller/namespaced/bedrockagent/agentknowledgebaseassociation"
	datasource "github.com/upbound/provider-aws/v2/internal/controller/namespaced/bedrockagent/datasource"
	flow "github.com/upbound/provider-aws/v2/internal/controller/namespaced/bedrockagent/flow"
	knowledgebase "github.com/upbound/provider-aws/v2/internal/controller/namespaced/bedrockagent/knowledgebase"
	prompt "github.com/upbound/provider-aws/v2/internal/controller/namespaced/bedrockagent/prompt"
)

// Setup_bedrockagent creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_bedrockagent(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		agent.Setup,
		agentknowledgebaseassociation.Setup,
		datasource.Setup,
		flow.Setup,
		knowledgebase.Setup,
		prompt.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_bedrockagent creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_bedrockagent(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		agent.SetupGated,
		agentknowledgebaseassociation.SetupGated,
		datasource.SetupGated,
		flow.SetupGated,
		knowledgebase.SetupGated,
		prompt.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_bedrockagent registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_bedrockagent(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		agent.SetupWebhookWithManager,
		agentknowledgebaseassociation.SetupWebhookWithManager,
		datasource.SetupWebhookWithManager,
		flow.SetupWebhookWithManager,
		knowledgebase.SetupWebhookWithManager,
		prompt.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
