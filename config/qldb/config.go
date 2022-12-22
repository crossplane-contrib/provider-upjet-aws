package qldb

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for kinesis group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_qldb_stream", func(r *config.Resource) {
		r.References["kinesis_configuration.stream_arn"] = config.Reference{
			TerraformName: "aws_kinesis_stream",
			Extractor:     common.PathTerraformIDExtractor,
		}
		r.References["ledger_name"] = config.Reference{
			TerraformName: "aws_qldb_ledger",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})
}
