/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	discoverer "github.com/upbound/provider-aws/internal/controller/schemas/discoverer"
	registry "github.com/upbound/provider-aws/internal/controller/schemas/registry"
	schema "github.com/upbound/provider-aws/internal/controller/schemas/schema"
)

// Setup_schemas creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_schemas(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		discoverer.Setup,
		registry.Setup,
		schema.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
