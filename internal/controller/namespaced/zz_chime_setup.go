// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	voiceconnector "github.com/upbound/provider-aws/v2/internal/controller/namespaced/chime/voiceconnector"
	voiceconnectorgroup "github.com/upbound/provider-aws/v2/internal/controller/namespaced/chime/voiceconnectorgroup"
	voiceconnectorlogging "github.com/upbound/provider-aws/v2/internal/controller/namespaced/chime/voiceconnectorlogging"
	voiceconnectororigination "github.com/upbound/provider-aws/v2/internal/controller/namespaced/chime/voiceconnectororigination"
	voiceconnectorstreaming "github.com/upbound/provider-aws/v2/internal/controller/namespaced/chime/voiceconnectorstreaming"
	voiceconnectortermination "github.com/upbound/provider-aws/v2/internal/controller/namespaced/chime/voiceconnectortermination"
	voiceconnectorterminationcredentials "github.com/upbound/provider-aws/v2/internal/controller/namespaced/chime/voiceconnectorterminationcredentials"
)

// Setup_chime creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_chime(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		voiceconnector.Setup,
		voiceconnectorgroup.Setup,
		voiceconnectorlogging.Setup,
		voiceconnectororigination.Setup,
		voiceconnectorstreaming.Setup,
		voiceconnectortermination.Setup,
		voiceconnectorterminationcredentials.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_chime creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_chime(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		voiceconnector.SetupGated,
		voiceconnectorgroup.SetupGated,
		voiceconnectorlogging.SetupGated,
		voiceconnectororigination.SetupGated,
		voiceconnectorstreaming.SetupGated,
		voiceconnectortermination.SetupGated,
		voiceconnectorterminationcredentials.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
