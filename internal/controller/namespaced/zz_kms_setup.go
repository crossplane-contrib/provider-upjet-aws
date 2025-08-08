// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	alias "github.com/upbound/provider-aws/internal/controller/namespaced/kms/alias"
	ciphertext "github.com/upbound/provider-aws/internal/controller/namespaced/kms/ciphertext"
	externalkey "github.com/upbound/provider-aws/internal/controller/namespaced/kms/externalkey"
	grant "github.com/upbound/provider-aws/internal/controller/namespaced/kms/grant"
	key "github.com/upbound/provider-aws/internal/controller/namespaced/kms/key"
	replicaexternalkey "github.com/upbound/provider-aws/internal/controller/namespaced/kms/replicaexternalkey"
	replicakey "github.com/upbound/provider-aws/internal/controller/namespaced/kms/replicakey"
)

// Setup_kms creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_kms(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alias.Setup,
		ciphertext.Setup,
		externalkey.Setup,
		grant.Setup,
		key.Setup,
		replicaexternalkey.Setup,
		replicakey.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_kms creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_kms(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alias.SetupGated,
		ciphertext.SetupGated,
		externalkey.SetupGated,
		grant.SetupGated,
		key.SetupGated,
		replicaexternalkey.SetupGated,
		replicakey.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
