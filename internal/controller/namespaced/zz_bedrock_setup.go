// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	custommodel "github.com/upbound/provider-aws/v2/internal/controller/namespaced/bedrock/custommodel"
	evaluationjob "github.com/upbound/provider-aws/v2/internal/controller/namespaced/bedrock/evaluationjob"
	foundationmodelagreement "github.com/upbound/provider-aws/v2/internal/controller/namespaced/bedrock/foundationmodelagreement"
	guardrail "github.com/upbound/provider-aws/v2/internal/controller/namespaced/bedrock/guardrail"
	guardrailversion "github.com/upbound/provider-aws/v2/internal/controller/namespaced/bedrock/guardrailversion"
	inferenceprofile "github.com/upbound/provider-aws/v2/internal/controller/namespaced/bedrock/inferenceprofile"
	modelinvocationjob "github.com/upbound/provider-aws/v2/internal/controller/namespaced/bedrock/modelinvocationjob"
	modelinvocationloggingconfiguration "github.com/upbound/provider-aws/v2/internal/controller/namespaced/bedrock/modelinvocationloggingconfiguration"
	provisionedmodelthroughput "github.com/upbound/provider-aws/v2/internal/controller/namespaced/bedrock/provisionedmodelthroughput"
	usecaseformodelaccess "github.com/upbound/provider-aws/v2/internal/controller/namespaced/bedrock/usecaseformodelaccess"
)

// Setup_bedrock creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_bedrock(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		custommodel.Setup,
		evaluationjob.Setup,
		foundationmodelagreement.Setup,
		guardrail.Setup,
		guardrailversion.Setup,
		inferenceprofile.Setup,
		modelinvocationjob.Setup,
		modelinvocationloggingconfiguration.Setup,
		provisionedmodelthroughput.Setup,
		usecaseformodelaccess.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_bedrock creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_bedrock(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		custommodel.SetupGated,
		evaluationjob.SetupGated,
		foundationmodelagreement.SetupGated,
		guardrail.SetupGated,
		guardrailversion.SetupGated,
		inferenceprofile.SetupGated,
		modelinvocationjob.SetupGated,
		modelinvocationloggingconfiguration.SetupGated,
		provisionedmodelthroughput.SetupGated,
		usecaseformodelaccess.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_bedrock registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_bedrock(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		custommodel.SetupWebhookWithManager,
		evaluationjob.SetupWebhookWithManager,
		foundationmodelagreement.SetupWebhookWithManager,
		guardrail.SetupWebhookWithManager,
		guardrailversion.SetupWebhookWithManager,
		inferenceprofile.SetupWebhookWithManager,
		modelinvocationjob.SetupWebhookWithManager,
		modelinvocationloggingconfiguration.SetupWebhookWithManager,
		provisionedmodelthroughput.SetupWebhookWithManager,
		usecaseformodelaccess.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
