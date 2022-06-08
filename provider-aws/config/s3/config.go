/*
Copyright 2021 Upbound Inc.
*/

package s3

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for s3 group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_s3_bucket", func(r *config.Resource) {
		r.References = config.References{
			"server_side_encryption_configuration.rule.apply_server_side_encryption_by_default.kms_master_key_id": {
				Type:      "github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key",
				Extractor: common.PathARNExtractor,
			},
		}
	})
}
