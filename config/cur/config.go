/*
Copyright 2022 Upbound Inc.
*/

package cur

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the cur group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_cur_report_definition", func(r *config.Resource) {
		r.References["s3_bucket"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket",
		}
	})
}
