// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package gamelift

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/namespaced/common"
)

// Configure adds configurations for the gamelift group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_gamelift_build", func(r *config.Resource) {
		r.References["storage_location.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		r.References["storage_location.bucket"] = config.Reference{
			TerraformName: "aws_s3_bucket",
		}
	})

	p.AddResourceConfigurator("aws_gamelift_fleet", func(r *config.Resource) {
		r.References["build_id"] = config.Reference{
			TerraformName: "aws_gamelift_build",
		}
		r.UseAsync = true
		r.Path = "fleet"
	})

	p.AddResourceConfigurator("aws_gamelift_game_server_group", func(r *config.Resource) {
		r.References["role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		r.References["launch_template.id"] = config.Reference{
			TerraformName: "aws_launch_template",
			Extractor:     common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_gamelift_game_session_queue", func(r *config.Resource) {
		r.References["notification_target"] = config.Reference{
			TerraformName: "aws_sns_topic",
			Extractor:     common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_gamelift_script", func(r *config.Resource) {
		r.References["storage_location.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		r.References["storage_location.bucket"] = config.Reference{
			TerraformName: "aws_s3_bucket",
		}
	})
}
