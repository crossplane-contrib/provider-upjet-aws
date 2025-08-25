// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	catalogdatabase "github.com/upbound/provider-aws/internal/controller/namespaced/glue/catalogdatabase"
	catalogtable "github.com/upbound/provider-aws/internal/controller/namespaced/glue/catalogtable"
	catalogtableoptimizer "github.com/upbound/provider-aws/internal/controller/namespaced/glue/catalogtableoptimizer"
	classifier "github.com/upbound/provider-aws/internal/controller/namespaced/glue/classifier"
	connection "github.com/upbound/provider-aws/internal/controller/namespaced/glue/connection"
	crawler "github.com/upbound/provider-aws/internal/controller/namespaced/glue/crawler"
	datacatalogencryptionsettings "github.com/upbound/provider-aws/internal/controller/namespaced/glue/datacatalogencryptionsettings"
	job "github.com/upbound/provider-aws/internal/controller/namespaced/glue/job"
	registry "github.com/upbound/provider-aws/internal/controller/namespaced/glue/registry"
	resourcepolicy "github.com/upbound/provider-aws/internal/controller/namespaced/glue/resourcepolicy"
	schema "github.com/upbound/provider-aws/internal/controller/namespaced/glue/schema"
	securityconfiguration "github.com/upbound/provider-aws/internal/controller/namespaced/glue/securityconfiguration"
	trigger "github.com/upbound/provider-aws/internal/controller/namespaced/glue/trigger"
	userdefinedfunction "github.com/upbound/provider-aws/internal/controller/namespaced/glue/userdefinedfunction"
	workflow "github.com/upbound/provider-aws/internal/controller/namespaced/glue/workflow"
)

// Setup_glue creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_glue(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		catalogdatabase.Setup,
		catalogtable.Setup,
		catalogtableoptimizer.Setup,
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

// SetupGated_glue creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_glue(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		catalogdatabase.SetupGated,
		catalogtable.SetupGated,
		catalogtableoptimizer.SetupGated,
		classifier.SetupGated,
		connection.SetupGated,
		crawler.SetupGated,
		datacatalogencryptionsettings.SetupGated,
		job.SetupGated,
		registry.SetupGated,
		resourcepolicy.SetupGated,
		schema.SetupGated,
		securityconfiguration.SetupGated,
		trigger.SetupGated,
		userdefinedfunction.SetupGated,
		workflow.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
