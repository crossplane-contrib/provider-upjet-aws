/*
Copyright 2021 Upbound Inc.
*/

package dms

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for dms group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_dms_endpoint", func(r *config.Resource) {
		r.References["service_access_role"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
	})
}
