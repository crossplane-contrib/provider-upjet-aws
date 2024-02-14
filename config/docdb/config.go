package docdb

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/types/comments"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the docdb group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_docdb_cluster", func(r *config.Resource) {
		config.MoveToStatus(r.TerraformResource, "cluster_members")
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["endpoint"].(string); ok {
				conn["endpoint"] = []byte(a)
			}
			if a, ok := attr["arn"].(string); ok {
				conn["arn"] = []byte(a)
			}
			if a, ok := attr["master_password"].(string); ok {
				conn["password"] = []byte(a)
			}
			return conn, nil
		}
		desc, _ := comments.New("If true, the password will be auto-generated and"+
			" stored in the Secret referenced by the masterPasswordSecretRef field.",
			comments.WithTFTag("-"))
		r.TerraformResource.Schema["auto_generate_password"] = &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Description: desc.String(),
		}
		r.InitializerFns = append(r.InitializerFns,
			common.PasswordGenerator(
				"spec.forProvider.masterPasswordSecretRef",
				"spec.forProvider.autoGeneratePassword",
			))
		r.TerraformResource.Schema["master_password"].Description = "Password for the " +
			"master DB user. If you set autoGeneratePassword to true, the Secret" +
			" referenced here will be created or updated with generated password" +
			" if it does not already contain one."
		r.References["db_cluster_parameter_group_name"] = config.Reference{
			TerraformName: "aws_docdb_cluster_parameter_group",
		}
	})

	p.AddResourceConfigurator("aws_docdb_cluster_instance", func(r *config.Resource) {
		r.References["cluster_identifier"] = config.Reference{
			Type: "Cluster",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_docdb_subnet_group", func(r *config.Resource) {
		r.References["subnet_ids"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
		}
	})
}
