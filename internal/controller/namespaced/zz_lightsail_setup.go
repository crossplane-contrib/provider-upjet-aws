// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	bucket "github.com/upbound/provider-aws/internal/controller/namespaced/lightsail/bucket"
	certificate "github.com/upbound/provider-aws/internal/controller/namespaced/lightsail/certificate"
	containerservice "github.com/upbound/provider-aws/internal/controller/namespaced/lightsail/containerservice"
	disk "github.com/upbound/provider-aws/internal/controller/namespaced/lightsail/disk"
	diskattachment "github.com/upbound/provider-aws/internal/controller/namespaced/lightsail/diskattachment"
	domain "github.com/upbound/provider-aws/internal/controller/namespaced/lightsail/domain"
	domainentry "github.com/upbound/provider-aws/internal/controller/namespaced/lightsail/domainentry"
	instance "github.com/upbound/provider-aws/internal/controller/namespaced/lightsail/instance"
	instancepublicports "github.com/upbound/provider-aws/internal/controller/namespaced/lightsail/instancepublicports"
	keypair "github.com/upbound/provider-aws/internal/controller/namespaced/lightsail/keypair"
	lb "github.com/upbound/provider-aws/internal/controller/namespaced/lightsail/lb"
	lbattachment "github.com/upbound/provider-aws/internal/controller/namespaced/lightsail/lbattachment"
	lbcertificate "github.com/upbound/provider-aws/internal/controller/namespaced/lightsail/lbcertificate"
	lbstickinesspolicy "github.com/upbound/provider-aws/internal/controller/namespaced/lightsail/lbstickinesspolicy"
	staticip "github.com/upbound/provider-aws/internal/controller/namespaced/lightsail/staticip"
	staticipattachment "github.com/upbound/provider-aws/internal/controller/namespaced/lightsail/staticipattachment"
)

// Setup_lightsail creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_lightsail(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bucket.Setup,
		certificate.Setup,
		containerservice.Setup,
		disk.Setup,
		diskattachment.Setup,
		domain.Setup,
		domainentry.Setup,
		instance.Setup,
		instancepublicports.Setup,
		keypair.Setup,
		lb.Setup,
		lbattachment.Setup,
		lbcertificate.Setup,
		lbstickinesspolicy.Setup,
		staticip.Setup,
		staticipattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_lightsail creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_lightsail(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bucket.SetupGated,
		certificate.SetupGated,
		containerservice.SetupGated,
		disk.SetupGated,
		diskattachment.SetupGated,
		domain.SetupGated,
		domainentry.SetupGated,
		instance.SetupGated,
		instancepublicports.SetupGated,
		keypair.SetupGated,
		lb.SetupGated,
		lbattachment.SetupGated,
		lbcertificate.SetupGated,
		lbstickinesspolicy.SetupGated,
		staticip.SetupGated,
		staticipattachment.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
