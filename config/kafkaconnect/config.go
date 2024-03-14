package kafkaconnect

import (
	"time"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the kafkaconnect group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_mskconnect_connector", func(r *config.Resource) {
		// This will always refer to a Cluster in the kafka api group, if it refers to any managed resource at all,
		// but which property from the status of that cluster to use depends on the authentication mechanism chosen.
		delete(r.References, "kafka_cluster.apache_kafka_cluster.bootstrap_servers")
		r.References["kafka_cluster.apache_kafka_cluster.vpc.security_groups"] = config.Reference{
			TerraformName:     "aws_security_group",
			SelectorFieldName: "SecurityGroupSelector",
			RefFieldName:      "SecurityGroupRefs",
		}
		r.References["kafka_cluster.apache_kafka_cluster.vpc.subnets"] = config.Reference{
			TerraformName:     "aws_subnet",
			SelectorFieldName: "SubnetSelector",
			RefFieldName:      "SubnetRefs",
		}
		r.References["log_delivery.worker_log_delivery.s3.bucket"] = config.Reference{
			TerraformName: "aws_s3_bucket",
		}
		r.References["log_delivery.worker_log_delivery.cloudwatch_logs.log_group"] = config.Reference{
			TerraformName: "aws_cloudwatch_log_group",
		}
		r.References["log_delivery.worker_log_delivery.firehose.delivery_stream"] = config.Reference{
			TerraformName: "aws_kinesis_firehose_delivery_stream",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)`,
		}
		r.References["worker_configuration.arn"] = config.Reference{
			TerraformName: "aws_mskconnect_worker_configuration",
			Extractor:     common.PathARNExtractor,
		}
		r.References["plugin.custom_plugin.arn"] = config.Reference{
			TerraformName: "aws_mskconnect_custom_plugin",
			Extractor:     common.PathARNExtractor,
		}
		// References only work to string fields.
		delete(r.References, "plugin.custom_plugin.revision")
		r.UseAsync = true
		r.OperationTimeouts.Create = 30 * time.Minute
		r.MetaResource.Description += ` Changes to any parameter besides "scaling" will be rejected. Instead you must create a new resource.`
	})
	p.AddResourceConfigurator("aws_mskconnect_custom_plugin", func(r *config.Resource) {
		r.MetaResource.Description += ` This resource can be Created, Observed and Deleted, but not Updated. AWS does not currently provide update APIs.`
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_mskconnect_worker_configuration", func(r *config.Resource) {
		r.MetaResource.Description += ` This resource is create-only, and requires a unique "name" parameter. AWS does not currently provide update or delete APIs.`
	})
}
