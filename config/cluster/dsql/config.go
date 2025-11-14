// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package dsql

import (
	"context"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/terraform/errors"
	"github.com/hashicorp/terraform-plugin-framework/path"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Configure adds configurations for the ds group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_dsql_cluster_peering", func(r *config.Resource) {
		r.TerraformPluginFrameworkIsStateEmptyFn = func(ctx context.Context, tfStateValue tftypes.Value, resourceSchema rschema.Schema) (bool, error) {
			sdkState := tfsdk.State{
				Raw:    tfStateValue,
				Schema: resourceSchema,
			}
			var clusters []string
			if diags := sdkState.GetAttribute(ctx, path.Root("clusters"), &clusters); diags.HasError() {
				return false, errors.FrameworkDiagnosticsError("reading clusters attribute", diags)
			}
			return len(clusters) == 0, nil
		}
	})
}
