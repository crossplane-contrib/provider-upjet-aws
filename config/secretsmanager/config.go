package secretsmanager

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for the secretsmanager group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_secretsmanager_secret", func(r *config.Resource) {
		// Use aws_secretsmanager_secret_rotation.
		config.MoveToStatus(r.TerraformResource, "rotation_rules", "rotation_lambda_arn")
		// aws_secretsmanager_secret_policy.
		config.MoveToStatus(r.TerraformResource, "policy")
		// TODO: we had better do this for all resources...
		r.TerraformConfigurationInjector = func(_ map[string]any, params map[string]any) error {
			params["name_prefix"] = ""
			return nil
		}
	})
}
