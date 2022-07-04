package backup

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for backup group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_backup_vault", func(r *config.Resource) {
		r.References = config.References{
			"kms_key_arn": config.Reference{
				Type:      "github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key",
				Extractor: common.PathARNExtractor,
			},
		}
	})

	p.AddResourceConfigurator("aws_backup_selection", func(r *config.Resource) {
		r.References = config.References{
			"iam_role_arn": config.Reference{
				Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role",
				Extractor: common.PathARNExtractor,
			},
			"plan_id": config.Reference{
				Type: "Plan",
			},
		}
	})

	p.AddResourceConfigurator("aws_backup_vault_notifications", func(r *config.Resource) {
		r.References = config.References{
			"sns_topic_arn": config.Reference{
				Type:      "github.com/upbound/official-providers/provider-aws/apis/sns/v1beta1.Topic",
				Extractor: common.PathARNExtractor,
			},
		}
	})
}
