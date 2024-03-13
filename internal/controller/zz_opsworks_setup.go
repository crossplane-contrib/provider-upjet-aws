package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	application "github.com/upbound/provider-aws/internal/controller/opsworks/application"
	customlayer "github.com/upbound/provider-aws/internal/controller/opsworks/customlayer"
	ecsclusterlayer "github.com/upbound/provider-aws/internal/controller/opsworks/ecsclusterlayer"
	ganglialayer "github.com/upbound/provider-aws/internal/controller/opsworks/ganglialayer"
	haproxylayer "github.com/upbound/provider-aws/internal/controller/opsworks/haproxylayer"
	instance "github.com/upbound/provider-aws/internal/controller/opsworks/instance"
	javaapplayer "github.com/upbound/provider-aws/internal/controller/opsworks/javaapplayer"
	memcachedlayer "github.com/upbound/provider-aws/internal/controller/opsworks/memcachedlayer"
	mysqllayer "github.com/upbound/provider-aws/internal/controller/opsworks/mysqllayer"
	nodejsapplayer "github.com/upbound/provider-aws/internal/controller/opsworks/nodejsapplayer"
	permission "github.com/upbound/provider-aws/internal/controller/opsworks/permission"
	phpapplayer "github.com/upbound/provider-aws/internal/controller/opsworks/phpapplayer"
	railsapplayer "github.com/upbound/provider-aws/internal/controller/opsworks/railsapplayer"
	rdsdbinstance "github.com/upbound/provider-aws/internal/controller/opsworks/rdsdbinstance"
	stack "github.com/upbound/provider-aws/internal/controller/opsworks/stack"
	staticweblayer "github.com/upbound/provider-aws/internal/controller/opsworks/staticweblayer"
	userprofile "github.com/upbound/provider-aws/internal/controller/opsworks/userprofile"
)

// Setup_opsworks creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_opsworks(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		customlayer.Setup,
		ecsclusterlayer.Setup,
		ganglialayer.Setup,
		haproxylayer.Setup,
		instance.Setup,
		javaapplayer.Setup,
		memcachedlayer.Setup,
		mysqllayer.Setup,
		nodejsapplayer.Setup,
		permission.Setup,
		phpapplayer.Setup,
		railsapplayer.Setup,
		rdsdbinstance.Setup,
		stack.Setup,
		staticweblayer.Setup,
		userprofile.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
