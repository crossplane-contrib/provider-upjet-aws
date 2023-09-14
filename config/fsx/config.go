/*
Copyright 2022 Upbound Inc.
*/

package fsx

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the fsx group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_fsx_windows_file_system", func(r *config.Resource) {
		r.References["kms_key_id"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			Extractor: common.PathARNExtractor,
		}
	})
}
