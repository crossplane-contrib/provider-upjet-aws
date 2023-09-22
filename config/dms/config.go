/*
Copyright 2021 Upbound Inc.
*/

package dms

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the dms group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_dms_endpoint", func(r *config.Resource) {
		r.References = config.References{
			"secrets_manager_access_role_arn": {
				Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
				Extractor: common.PathARNExtractor,
			},
			"service_access_role": {
				Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
				Extractor: common.PathARNExtractor,
			},
			"kms_key_arn": {
				Type:      "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
				Extractor: common.PathARNExtractor,
			},
		}
	})
}
