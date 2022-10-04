package grafana

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for grafana group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_grafana_workspace", func(r *config.Resource) {
		r.References["role_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_grafana_role_association", func(r *config.Resource) {
		r.References["workspace_id"] = config.Reference{
			Type: "Workspace",
		}
	})

	p.AddResourceConfigurator("aws_grafana_workspace_saml_configuration", func(r *config.Resource) {
		r.References["workspace_id"] = config.Reference{
			Type: "Workspace",
		}
	})
}
