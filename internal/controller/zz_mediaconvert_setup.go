package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	queue "github.com/upbound/provider-aws/internal/controller/mediaconvert/queue"
)

// Setup_mediaconvert creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_mediaconvert(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		queue.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
