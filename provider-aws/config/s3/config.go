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
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["bucket"] = name
			},
			OmittedFields: []string{
				"bucket",
				"bucket_prefix",
			},
			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
		r.References = config.References{
			"server_side_encryption_configuration.rule.apply_server_side_encryption_by_default.kms_master_key_id": {
				Type:      "github.com/upbound/official-providers/provider-aws/apis/kms/v1alpha2.Key",
				Extractor: common.PathARNExtractor,
			},
		}
	})
}
