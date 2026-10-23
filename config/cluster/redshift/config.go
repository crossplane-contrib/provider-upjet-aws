// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package redshift

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/cluster/common"
)

// Configure adds configurations for the redshift group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_redshift_cluster", func(r *config.Resource) {
		r.UseAsync = true
		r.Version = "v1beta2"
		r.SetCRDStorageVersion("v1beta1")
		r.ControllerReconcileVersion = "v1beta1" //nolint:staticcheck // still handling the deprecated behavior until rollout
	})
	p.AddResourceConfigurator("aws_redshift_resource_policy", func(r *config.Resource) {
		// The resource policy is used to authorize zero-ETL integrations into a
		// target Redshift data warehouse. The target can be a Redshift
		// Serverless namespace, whose ARN lives in the status, so reference it
		// instead of the auto-injected provisioned cluster reference.
		r.References["resource_arn"] = config.Reference{
			TerraformName: "aws_redshiftserverless_namespace",
			Extractor:     common.PathARNExtractor,
		}
	})
}
