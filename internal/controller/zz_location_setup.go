package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	geofencecollection "github.com/upbound/provider-aws/internal/controller/location/geofencecollection"
	placeindex "github.com/upbound/provider-aws/internal/controller/location/placeindex"
	routecalculator "github.com/upbound/provider-aws/internal/controller/location/routecalculator"
	tracker "github.com/upbound/provider-aws/internal/controller/location/tracker"
	trackerassociation "github.com/upbound/provider-aws/internal/controller/location/trackerassociation"
)

// Setup_location creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_location(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		geofencecollection.Setup,
		placeindex.Setup,
		routecalculator.Setup,
		tracker.Setup,
		trackerassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
