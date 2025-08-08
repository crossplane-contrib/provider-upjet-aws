// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	authorizer "github.com/upbound/provider-aws/internal/controller/cluster/iot/authorizer"
	certificate "github.com/upbound/provider-aws/internal/controller/cluster/iot/certificate"
	domainconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/iot/domainconfiguration"
	indexingconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/iot/indexingconfiguration"
	loggingoptions "github.com/upbound/provider-aws/internal/controller/cluster/iot/loggingoptions"
	policy "github.com/upbound/provider-aws/internal/controller/cluster/iot/policy"
	policyattachment "github.com/upbound/provider-aws/internal/controller/cluster/iot/policyattachment"
	provisioningtemplate "github.com/upbound/provider-aws/internal/controller/cluster/iot/provisioningtemplate"
	rolealias "github.com/upbound/provider-aws/internal/controller/cluster/iot/rolealias"
	thing "github.com/upbound/provider-aws/internal/controller/cluster/iot/thing"
	thinggroup "github.com/upbound/provider-aws/internal/controller/cluster/iot/thinggroup"
	thinggroupmembership "github.com/upbound/provider-aws/internal/controller/cluster/iot/thinggroupmembership"
	thingprincipalattachment "github.com/upbound/provider-aws/internal/controller/cluster/iot/thingprincipalattachment"
	thingtype "github.com/upbound/provider-aws/internal/controller/cluster/iot/thingtype"
	topicrule "github.com/upbound/provider-aws/internal/controller/cluster/iot/topicrule"
	topicruledestination "github.com/upbound/provider-aws/internal/controller/cluster/iot/topicruledestination"
)

// Setup_iot creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_iot(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authorizer.Setup,
		certificate.Setup,
		domainconfiguration.Setup,
		indexingconfiguration.Setup,
		loggingoptions.Setup,
		policy.Setup,
		policyattachment.Setup,
		provisioningtemplate.Setup,
		rolealias.Setup,
		thing.Setup,
		thinggroup.Setup,
		thinggroupmembership.Setup,
		thingprincipalattachment.Setup,
		thingtype.Setup,
		topicrule.Setup,
		topicruledestination.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_iot creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_iot(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authorizer.SetupGated,
		certificate.SetupGated,
		domainconfiguration.SetupGated,
		indexingconfiguration.SetupGated,
		loggingoptions.SetupGated,
		policy.SetupGated,
		policyattachment.SetupGated,
		provisioningtemplate.SetupGated,
		rolealias.SetupGated,
		thing.SetupGated,
		thinggroup.SetupGated,
		thinggroupmembership.SetupGated,
		thingprincipalattachment.SetupGated,
		thingtype.SetupGated,
		topicrule.SetupGated,
		topicruledestination.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
