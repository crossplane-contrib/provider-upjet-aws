// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	repository "github.com/upbound/provider-aws/internal/controller/cluster/ecrpublic/repository"
	repositorypolicy "github.com/upbound/provider-aws/internal/controller/cluster/ecrpublic/repositorypolicy"
)

// Setup_ecrpublic creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ecrpublic(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		repository.Setup,
		repositorypolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
