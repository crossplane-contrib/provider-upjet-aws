/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	certificate "github.com/upbound/provider-aws/internal/controller/dms/certificate"
	endpoint "github.com/upbound/provider-aws/internal/controller/dms/endpoint"
	eventsubscription "github.com/upbound/provider-aws/internal/controller/dms/eventsubscription"
	replicationinstance "github.com/upbound/provider-aws/internal/controller/dms/replicationinstance"
	replicationsubnetgroup "github.com/upbound/provider-aws/internal/controller/dms/replicationsubnetgroup"
	replicationtask "github.com/upbound/provider-aws/internal/controller/dms/replicationtask"
	s3endpoint "github.com/upbound/provider-aws/internal/controller/dms/s3endpoint"
)

// Setup_dms creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dms(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		certificate.Setup,
		endpoint.Setup,
		eventsubscription.Setup,
		replicationinstance.Setup,
		replicationsubnetgroup.Setup,
		replicationtask.Setup,
		s3endpoint.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
