package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	endpointaccess "github.com/upbound/provider-aws/internal/controller/redshiftserverless/endpointaccess"
	redshiftserverlessnamespace "github.com/upbound/provider-aws/internal/controller/redshiftserverless/redshiftserverlessnamespace"
	resourcepolicy "github.com/upbound/provider-aws/internal/controller/redshiftserverless/resourcepolicy"
	snapshot "github.com/upbound/provider-aws/internal/controller/redshiftserverless/snapshot"
	usagelimit "github.com/upbound/provider-aws/internal/controller/redshiftserverless/usagelimit"
	workgroup "github.com/upbound/provider-aws/internal/controller/redshiftserverless/workgroup"
)

// Setup_redshiftserverless creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_redshiftserverless(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		endpointaccess.Setup,
		redshiftserverlessnamespace.Setup,
		resourcepolicy.Setup,
		snapshot.Setup,
		usagelimit.Setup,
		workgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
