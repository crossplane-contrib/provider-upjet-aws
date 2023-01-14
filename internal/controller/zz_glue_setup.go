/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	catalogdatabase "github.com/upbound/provider-aws/internal/controller/glue/catalogdatabase"
	crawler "github.com/upbound/provider-aws/internal/controller/glue/crawler"
)

// Setup_glue creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_glue(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		catalogdatabase.Setup,
		crawler.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
