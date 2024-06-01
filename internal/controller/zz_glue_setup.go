// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	catalogdatabase "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/glue/catalogdatabase"
	catalogtable "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/glue/catalogtable"
	classifier "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/glue/classifier"
	connection "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/glue/connection"
	crawler "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/glue/crawler"
	datacatalogencryptionsettings "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/glue/datacatalogencryptionsettings"
	job "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/glue/job"
	registry "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/glue/registry"
	resourcepolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/glue/resourcepolicy"
	schema "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/glue/schema"
	securityconfiguration "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/glue/securityconfiguration"
	trigger "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/glue/trigger"
	userdefinedfunction "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/glue/userdefinedfunction"
	workflow "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/glue/workflow"
)

// Setup_glue creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_glue(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		catalogdatabase.Setup,
		catalogtable.Setup,
		classifier.Setup,
		connection.Setup,
		crawler.Setup,
		datacatalogencryptionsettings.Setup,
		job.Setup,
		registry.Setup,
		resourcepolicy.Setup,
		schema.Setup,
		securityconfiguration.Setup,
		trigger.Setup,
		userdefinedfunction.Setup,
		workflow.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
