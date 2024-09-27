package batch

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the batch group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_batch_compute_environment", func(r *config.Resource) {
		r.References = config.References{
			"compute_resources.instance_role": config.Reference{
				TerraformName:     "aws_iam_instance_profile",
				RefFieldName:      "InstanceRoleRef",
				SelectorFieldName: "InstanceRoleSelector",
			},
			"compute_resources.placement_group": config.Reference{
				TerraformName:     "aws_placement_group",
				RefFieldName:      "PlacementGroupRef",
				SelectorFieldName: "PlacementGroupSelector",
			},
			"service_role": config.Reference{
				TerraformName:     "aws_iam_role",
				RefFieldName:      "ServiceRoleRef",
				SelectorFieldName: "ServiceRoleSelector",
			},
			"compute_resources.subnets": config.Reference{
				TerraformName:     "aws_subnet",
				RefFieldName:      "SubnetRefs",
				SelectorFieldName: "SubnetSelector",
			},
			"compute_resources.security_group_ids": config.Reference{
				TerraformName:     "aws_security_group",
				RefFieldName:      "SecurityGroupRefs",
				SelectorFieldName: "SecurityGroupSelector",
			},
		}
		r.UseAsync = true
	})
}
