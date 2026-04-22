// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package quicksight

import (
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the quicksight group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_quicksight_dashboard", func(r *config.Resource) {
		delete(r.TerraformResource.Schema, "definition")
		l := r.TFListConversionPaths()
		for _, e := range l {
			if strings.HasPrefix(e, "definition[*].") {
				r.RemoveSingletonListConversion(e)
			}
		}
		// r.MetaResource.ArgumentDocs["definition_json"] = "A raw JSON string used to define the dashboard structure. When this field is used, Crossplane cannot observe changes in the configuration through the AWS API; therefore, drift detection cannot be performed. Refer to the AWS documentation for the expected JSON structure: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CreateDashboard.html"
		r.MetaResource.Description = "Creates a QuickSight Dashboard resource. The 'definition' field is not supported due to Kubernetes CRD size limitations with deeply nested fields." // Please use the 'definitionJson' field to define the dashboard structure."
	})

	p.AddResourceConfigurator("aws_quicksight_analysis", func(r *config.Resource) {
		delete(r.TerraformResource.Schema, "definition")
		l := r.TFListConversionPaths()
		for _, e := range l {
			if strings.HasPrefix(e, "definition[*].") {
				r.RemoveSingletonListConversion(e)
			}
		}
		r.MetaResource.Description = "Creates a QuickSight Analysis resource. The 'definition' field is not supported due to Kubernetes CRD size limitations with deeply nested fields."
	})
}
