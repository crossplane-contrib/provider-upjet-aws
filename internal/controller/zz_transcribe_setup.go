// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	languagemodel "github.com/upbound/provider-aws/internal/controller/transcribe/languagemodel"
	vocabulary "github.com/upbound/provider-aws/internal/controller/transcribe/vocabulary"
	vocabularyfilter "github.com/upbound/provider-aws/internal/controller/transcribe/vocabularyfilter"
)

// Setup_transcribe creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_transcribe(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		languagemodel.Setup,
		vocabulary.Setup,
		vocabularyfilter.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
