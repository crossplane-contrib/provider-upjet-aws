// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	index "github.com/upbound/provider-aws/v2/internal/controller/namespaced/s3vectors/index"
	vectorbucket "github.com/upbound/provider-aws/v2/internal/controller/namespaced/s3vectors/vectorbucket"
	vectorbucketpolicy "github.com/upbound/provider-aws/v2/internal/controller/namespaced/s3vectors/vectorbucketpolicy"
)

// Setup_s3vectors creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_s3vectors(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		index.Setup,
		vectorbucket.Setup,
		vectorbucketpolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_s3vectors creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_s3vectors(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		index.SetupGated,
		vectorbucket.SetupGated,
		vectorbucketpolicy.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
