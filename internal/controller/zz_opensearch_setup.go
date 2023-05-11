/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	domain "github.com/upbound/provider-aws/internal/controller/opensearch/domain"
	domainpolicy "github.com/upbound/provider-aws/internal/controller/opensearch/domainpolicy"
	domainsamloptions "github.com/upbound/provider-aws/internal/controller/opensearch/domainsamloptions"
)

// Setup_opensearch creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_opensearch(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		domain.Setup,
		domainpolicy.Setup,
		domainsamloptions.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
