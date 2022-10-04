package gamelift

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for gamelift group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_gamelift_build", func(r *config.Resource) {
		r.References["storage_location.role_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
		r.References["storage_location.bucket"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket",
		}
	})

	p.AddResourceConfigurator("aws_gamelift_fleet", func(r *config.Resource) {
		r.References["build_id"] = config.Reference{
			Type: "Build",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_gamelift_game_server_group", func(r *config.Resource) {
		r.References["role_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
		r.References["launch_template.id"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/ec2/v1beta1.LaunchTemplate",
			Extractor: common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_gamelift_game_session_queue", func(r *config.Resource) {
		r.References["notification_target"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/sns/v1beta1.Topic",
			Extractor: common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_gamelift_script", func(r *config.Resource) {
		r.References["storage_location.role_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
		r.References["storage_location.bucket"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket",
		}
	})
}
