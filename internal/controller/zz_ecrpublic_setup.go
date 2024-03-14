package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	repository "github.com/upbound/provider-aws/internal/controller/ecrpublic/repository"
	repositorypolicy "github.com/upbound/provider-aws/internal/controller/ecrpublic/repositorypolicy"
)

// Setup_ecrpublic creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ecrpublic(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		repository.Setup,
		repositorypolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
