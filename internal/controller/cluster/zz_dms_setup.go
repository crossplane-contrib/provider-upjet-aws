// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	certificate "github.com/upbound/provider-aws/v2/internal/controller/cluster/dms/certificate"
	endpoint "github.com/upbound/provider-aws/v2/internal/controller/cluster/dms/endpoint"
	eventsubscription "github.com/upbound/provider-aws/v2/internal/controller/cluster/dms/eventsubscription"
	replicationinstance "github.com/upbound/provider-aws/v2/internal/controller/cluster/dms/replicationinstance"
	replicationsubnetgroup "github.com/upbound/provider-aws/v2/internal/controller/cluster/dms/replicationsubnetgroup"
	replicationtask "github.com/upbound/provider-aws/v2/internal/controller/cluster/dms/replicationtask"
	s3endpoint "github.com/upbound/provider-aws/v2/internal/controller/cluster/dms/s3endpoint"
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

// SetupGated_dms creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dms(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		certificate.SetupGated,
		endpoint.SetupGated,
		eventsubscription.SetupGated,
		replicationinstance.SetupGated,
		replicationsubnetgroup.SetupGated,
		replicationtask.SetupGated,
		s3endpoint.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
