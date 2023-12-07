// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	certificate "github.com/upbound/provider-aws/internal/controller/iot/certificate"
	indexingconfiguration "github.com/upbound/provider-aws/internal/controller/iot/indexingconfiguration"
	loggingoptions "github.com/upbound/provider-aws/internal/controller/iot/loggingoptions"
	policy "github.com/upbound/provider-aws/internal/controller/iot/policy"
	policyattachment "github.com/upbound/provider-aws/internal/controller/iot/policyattachment"
	provisioningtemplate "github.com/upbound/provider-aws/internal/controller/iot/provisioningtemplate"
	rolealias "github.com/upbound/provider-aws/internal/controller/iot/rolealias"
	thing "github.com/upbound/provider-aws/internal/controller/iot/thing"
	thinggroup "github.com/upbound/provider-aws/internal/controller/iot/thinggroup"
	thinggroupmembership "github.com/upbound/provider-aws/internal/controller/iot/thinggroupmembership"
	thingprincipalattachment "github.com/upbound/provider-aws/internal/controller/iot/thingprincipalattachment"
	thingtype "github.com/upbound/provider-aws/internal/controller/iot/thingtype"
	topicrule "github.com/upbound/provider-aws/internal/controller/iot/topicrule"
	topicruledestination "github.com/upbound/provider-aws/internal/controller/iot/topicruledestination"
)

// Setup_iot creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_iot(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		certificate.Setup,
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
