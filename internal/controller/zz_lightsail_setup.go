/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	bucket "github.com/upbound/provider-aws/internal/controller/lightsail/bucket"
	certificate "github.com/upbound/provider-aws/internal/controller/lightsail/certificate"
	containerservice "github.com/upbound/provider-aws/internal/controller/lightsail/containerservice"
	disk "github.com/upbound/provider-aws/internal/controller/lightsail/disk"
	diskattachment "github.com/upbound/provider-aws/internal/controller/lightsail/diskattachment"
	domain "github.com/upbound/provider-aws/internal/controller/lightsail/domain"
	domainentry "github.com/upbound/provider-aws/internal/controller/lightsail/domainentry"
	instance "github.com/upbound/provider-aws/internal/controller/lightsail/instance"
	instancepublicports "github.com/upbound/provider-aws/internal/controller/lightsail/instancepublicports"
	keypair "github.com/upbound/provider-aws/internal/controller/lightsail/keypair"
	lb "github.com/upbound/provider-aws/internal/controller/lightsail/lb"
	lbattachment "github.com/upbound/provider-aws/internal/controller/lightsail/lbattachment"
	lbcertificate "github.com/upbound/provider-aws/internal/controller/lightsail/lbcertificate"
	lbstickinesspolicy "github.com/upbound/provider-aws/internal/controller/lightsail/lbstickinesspolicy"
	staticip "github.com/upbound/provider-aws/internal/controller/lightsail/staticip"
	staticipattachment "github.com/upbound/provider-aws/internal/controller/lightsail/staticipattachment"
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
