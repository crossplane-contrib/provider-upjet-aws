// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	voiceconnector "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/chime/voiceconnector"
	voiceconnectorgroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/chime/voiceconnectorgroup"
	voiceconnectorlogging "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/chime/voiceconnectorlogging"
	voiceconnectororigination "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/chime/voiceconnectororigination"
	voiceconnectorstreaming "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/chime/voiceconnectorstreaming"
	voiceconnectortermination "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/chime/voiceconnectortermination"
	voiceconnectorterminationcredentials "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/chime/voiceconnectorterminationcredentials"
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
