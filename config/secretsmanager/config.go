package secretsmanager

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for the secretsmanager group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_secretsmanager_secret", func(r *config.Resource) {
		// Use aws_secretsmanager_secret_rotation.
		config.MoveToStatus(r.TerraformResource, "rotation_rules", "rotation_lambda_arn")
		// aws_secretsmanager_secret_policy.
		config.MoveToStatus(r.TerraformResource, "policy")
		delete(r.TerraformResource.Schema, "name_prefix")
	})
}
