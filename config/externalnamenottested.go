/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/upbound/upjet/pkg/config"
)

// ExternalNameNotTestedConfigs contains no-tested configurations for this
// provider.
var ExternalNameNotTestedConfigs = map[string]config.ExternalName{

	// amp
	//

	// amplify
	//
	// Amplify domain association can be imported using app_id and domain_name: d2ypk4k47z8u6/example.com
	"aws_amplify_domain_association": config.TemplatedStringAsIdentifier("domain_name", "{{ .parameters.app_id }}/{{ .external_name }}"),

	// apprunner
	//
	// App Runner Custom Domain Associations can be imported by using the domain_name and service_arn separated by a comma (,)
	"aws_apprunner_custom_domain_association": config.TemplatedStringAsIdentifier("domain_name", "{{ .external_name }},{{ .parameters.service_arn }}"),

	// aws_appsync_domain_name can be imported using the AppSync domain name
	"aws_appsync_domain_name": config.ParameterAsIdentifier("domain_name"),
	// aws_appsync_domain_name_api_association can be imported using the AppSync domain name
	"aws_appsync_domain_name_api_association": config.ParameterAsIdentifier("domain_name"),

	// batch
	//
	// AWS Batch compute can be imported using the compute_environment_name
	"aws_batch_compute_environment": config.ParameterAsIdentifier("compute_environment_name"),
	// Batch Job Definition can be imported using the arn: arn:aws:batch:us-east-1:123456789012:job-definition/sample
	"aws_batch_job_definition": config.TemplatedStringAsIdentifier("name", "arn:aws:batch:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:job-definition/{{ .external_name }}"),
	// Batch Job Queue can be imported using the arn: arn:aws:batch:us-east-1:123456789012:job-queue/sample
	"aws_batch_job_queue": config.TemplatedStringAsIdentifier("name", "arn:aws:batch:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:job-queue/{{ .external_name }}"),

	// ce
	//
	// aws_ce_cost_category can be imported using the id
	"aws_ce_cost_category": config.IdentifierFromProvider,

	// cloudformation
	//
	// Cloudformation Stacks Instances imported using the StackSet name, target AWS account ID, and target AWS: example,123456789012,us-east-1
	"aws_cloudformation_stack_set_instance": config.IdentifierFromProvider,
	// aws_cloudformation_type can be imported with their type version Amazon Resource Name (ARN)
	"aws_cloudformation_type": config.IdentifierFromProvider,

	// cloudhsmv2
	//
	// CloudHSM v2 Clusters can be imported using the cluster id
	"aws_cloudhsm_v2_cluster": config.IdentifierFromProvider,
	// HSM modules can be imported using their HSM ID
	"aws_cloudhsm_v2_hsm": config.IdentifierFromProvider,

	// codeartifact
	//
	// CodeArtifact Domain can be imported using the CodeArtifact Domain arn
	"aws_codeartifact_domain": config.IdentifierFromProvider,
	// CodeArtifact Domain Permissions Policies can be imported using the CodeArtifact Domain ARN
	"aws_codeartifact_domain_permissions_policy": config.IdentifierFromProvider,
	// CodeArtifact Repository can be imported using the CodeArtifact Repository ARN
	"aws_codeartifact_repository": config.IdentifierFromProvider,
	// CodeArtifact Repository Permissions Policies can be imported using the CodeArtifact Repository ARN
	"aws_codeartifact_repository_permissions_policy": config.IdentifierFromProvider,

	// codebuild
	//
	// CodeBuild Project can be imported using the name
	"aws_codebuild_project": config.NameAsIdentifier,
	// CodeBuild Report Group can be imported using the CodeBuild Report Group arn: arn:aws:codebuild:us-west-2:123456789:report-group/report-group-name
	"aws_codebuild_report_group": config.TemplatedStringAsIdentifier("name", "arn:aws:codebuild:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:report-group/{{ .external_name }}"),
	// CodeBuild Resource Policy can be imported using the CodeBuild Resource Policy arn
	"aws_codebuild_resource_policy": config.IdentifierFromProvider,
	// CodeBuild Source Credential can be imported using the CodeBuild Source Credential arn: arn:aws:codebuild:us-west-2:123456789:token:github
	"aws_codebuild_source_credential": config.TemplatedStringAsIdentifier("", "arn:aws:codebuild:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:token:{{ .parameters.token }}"),
	// CodeBuild Webhooks can be imported using the CodeBuild Project name
	"aws_codebuild_webhook": config.ParameterAsIdentifier("project_name"),

	// cognitoidp
	//
	// Cognito User Groups can be imported using the user_pool_id/name attributes concatenated
	"aws_cognito_user_group": config.TemplatedStringAsIdentifier("name", "{{ .parameters.user_pool_id }}/{{ .external_name }}"),

	// configservice
	//
	// Config aggregate authorizations can be imported using account_id:region
	"aws_config_aggregate_authorization": config.TemplatedStringAsIdentifier("", "{{ .parameters.account_id }}:{{ .parameters.region }}"),
	// Config Organization Conformance Packs can be imported using the name
	"aws_config_organization_conformance_pack": config.NameAsIdentifier,
	// Config Organization Custom Rules can be imported using the name
	"aws_config_organization_custom_rule": config.NameAsIdentifier,
	// Config Organization Managed Rules can be imported using the name
	"aws_config_organization_managed_rule": config.NameAsIdentifier,

	// connect
	//
	// Amazon Connect User Hierarchy Groups can be imported using the instance_id and hierarchy_group_id separated by a colon (:)
	"aws_connect_user_hierarchy_group": config.IdentifierFromProvider,

	// datapipeline
	//
	// aws_datapipeline_pipeline_definition can be imported using the id
	"aws_datapipeline_pipeline_definition": config.IdentifierFromProvider,

	// datasync
	//
	// aws_datasync_agent can be imported by using the DataSync Agent Amazon Resource Name (ARN)
	"aws_datasync_agent": config.IdentifierFromProvider,
	// aws_datasync_location_efs can be imported by using the DataSync Task Amazon Resource Name (ARN)
	"aws_datasync_location_efs": config.IdentifierFromProvider,
	// aws_datasync_location_fsx_lustre_file_system can be imported by using the DataSync-ARN#FSx-Lustre-ARN
	"aws_datasync_location_fsx_lustre_file_system": config.IdentifierFromProvider,
	// aws_datasync_location_fsx_openzfs_file_system can be imported by using the DataSync-ARN#FSx-openzfs-ARN
	"aws_datasync_location_fsx_openzfs_file_system": config.IdentifierFromProvider,
	// aws_datasync_location_fsx_windows_file_system can be imported by using the DataSync-ARN#FSx-Windows-ARN
	"aws_datasync_location_fsx_windows_file_system": config.IdentifierFromProvider,
	// aws_datasync_location_hdfs can be imported by using the Amazon Resource Name (ARN)
	"aws_datasync_location_hdfs": config.IdentifierFromProvider,
	// aws_datasync_location_nfs can be imported by using the DataSync Task Amazon Resource Name (ARN)
	"aws_datasync_location_nfs": config.IdentifierFromProvider,
	// aws_datasync_location_s3 can be imported by using the DataSync Task Amazon Resource Name (ARN)
	"aws_datasync_location_s3": config.IdentifierFromProvider,
	// aws_datasync_location_smb can be imported by using the Amazon Resource Name (ARN)
	"aws_datasync_location_smb": config.IdentifierFromProvider,
	// aws_datasync_task can be imported by using the DataSync Task Amazon Resource Name (ARN)
	"aws_datasync_task": config.IdentifierFromProvider,

	// directconnect
	//
	// No import
	"aws_dx_connection_confirmation": config.IdentifierFromProvider,
	// No import
	"aws_dx_hosted_connection": config.IdentifierFromProvider,

	// dms
	//
	// Event subscriptions can be imported using the name
	"aws_dms_event_subscription": config.NameAsIdentifier,
	// Replication instances can be imported using the replication_instance_id
	"aws_dms_replication_instance": config.ParameterAsIdentifier("replication_instance_id"),
	// Replication tasks can be imported using the replication_task_id
	"aws_dms_replication_task": config.ParameterAsIdentifier("replication_task_id"),

	// ds
	//
	// Conditional forwarders can be imported using the directory id and remote_domain_name: d-1234567890:example.com
	"aws_directory_service_conditional_forwarder": config.TemplatedStringAsIdentifier("", "{{ .parameters.directory_id }}:{{ .parameters.remote_domain_name }}"),
	// Directory Service Log Subscriptions can be imported using the directory id
	"aws_directory_service_log_subscription": config.ParameterAsIdentifier("directory_id"),

	// ec2
	//
	// No import
	"aws_ami_from_instance": config.IdentifierFromProvider,
	//
	"aws_ec2_client_vpn_authorization_rule": config.IdentifierFromProvider,
	// AWS Client VPN endpoints can be imported using the id value found via aws ec2 describe-client-vpn-endpoints
	"aws_ec2_client_vpn_endpoint": config.IdentifierFromProvider,
	// AWS Client VPN network associations can be imported using the endpoint ID and the association ID. Values are separated by a ,
	"aws_ec2_client_vpn_network_association": config.IdentifierFromProvider,
	// AWS Client VPN routes can be imported using the endpoint ID, target subnet ID, and destination CIDR block. All values are separated by a ,
	"aws_ec2_client_vpn_route": config.TemplatedStringAsIdentifier("", "{{ .parameters.client_vpn_endpoint_id }},{{ .parameters.target_vpc_subnet_id }},{{ .parameters.destination_cidr_block }}"),
	// aws_ec2_fleet can be imported by using the Fleet identifier
	"aws_ec2_fleet": config.IdentifierFromProvider,
	// aws_ec2_local_gateway_route can be imported by using the EC2 Local Gateway Route Table identifier and destination CIDR block separated by underscores (_)
	"aws_ec2_local_gateway_route": config.TemplatedStringAsIdentifier("", "{{ .parameters.local_gateway_route_table_id }}_{{ .parameters.destination_cidr_block }}"),
	// aws_ec2_local_gateway_route_table_vpc_association can be imported by using the Local Gateway Route Table VPC Association identifier
	"aws_ec2_local_gateway_route_table_vpc_association": config.IdentifierFromProvider,
	// aws_ec2_tag can be imported by using the EC2 resource identifier and key, separated by a comma (,)
	"aws_ec2_tag": config.TemplatedStringAsIdentifier("", "{{ .parameters.resource_id }}_{{ .parameters.key }}"),
	// Traffic mirror sessions can be imported using the id
	"aws_ec2_traffic_mirror_session": config.IdentifierFromProvider,
	// Traffic mirror targets can be imported using the id
	"aws_ec2_traffic_mirror_target": config.IdentifierFromProvider,
	// Internet Gateway Attachments can be imported using the id
	"aws_internet_gateway_attachment": config.IdentifierFromProvider,
	// No import
	"aws_network_acl_association": config.IdentifierFromProvider,
	// VPC Endpoint Services can be imported using ID of the connection, which is the VPC Endpoint Service ID and VPC Endpoint ID separated by underscore (_)
	"aws_vpc_endpoint_connection_accepter": config.TemplatedStringAsIdentifier("", "{{ .parameters.vpc_endpoint_service_id }}_{{ .parameters.vpc_endpoint_id }}"),
	// VPC Endpoint Policies can be imported using the id
	"aws_vpc_endpoint_policy": config.IdentifierFromProvider,
	// No import
	"aws_vpc_endpoint_security_group_association": config.IdentifierFromProvider,
	// IPAMs can be imported using the delegate account id
	"aws_vpc_ipam_organization_admin_account": config.ParameterAsIdentifier("delegated_admin_account_id"),
	// IPAMs can be imported using the <cidr>_<ipam-pool-id>
	"aws_vpc_ipam_pool_aws_default_network_acl": config.IdentifierFromProvider,
	// No import
	"aws_vpc_ipam_preview_next_cidr": config.IdentifierFromProvider,
	// aws_vpc_ipv6_cidr_block_association can be imported by using the VPC CIDR Association ID
	"aws_vpc_ipv6_cidr_block_association": config.IdentifierFromProvider,

	// securityhub
	//
	// imported using the AWS account ID
	"aws_securityhub_organization_admin_account": FormattedIdentifierFromProvider("", "admin_account_id"),
	// imported using the AWS account ID
	// no Terraform argument specifies the AWS account ID and
	// Terraform resource ID is the AWS account ID for the resource
	"aws_securityhub_organization_configuration": config.IdentifierFromProvider,
	// no import documentation
	"aws_securityhub_standards_control": config.IdentifierFromProvider,

	// servicecatalog
	//
	// no import documentation
	"aws_servicecatalog_organizations_access": config.IdentifierFromProvider,
	// imported using the provisioned product ID,
	// which has provider-generated random parts:
	// pp-dnigbtea24ste
	"aws_servicecatalog_provisioned_product": config.IdentifierFromProvider,

	// servicediscovery
	//
	// imported using the service ID and instance ID:
	// 0123456789/i-0123
	"aws_service_discovery_instance": FormattedIdentifierFromProvider("/", "service_id", "instance_id"),

	// elasticache
	//
	// ElastiCache Security Groups can be imported by name
	"aws_elasticache_security_group": config.NameAsIdentifier,
	// ElastiCache Global Replication Groups can be imported using the global_replication_group_id,
	// which is an attribute reported in the state.
	// TODO: we need to check the value of a global_replication_group_id to
	// see if further normalization is possible
	"aws_elasticache_global_replication_group": config.IdentifierFromProvider,
	// ElastiCache user group associations can be imported using the user_group_id and user_id:
	// userGoupId1,userId
	"aws_elasticache_user_group_association": FormattedIdentifierFromProvider(",", "user_group_id", "user_id"),

	// ram
	//
	// RAM Principal Associations can be imported using their Resource Share ARN and the principal separated by a comma:
	// arn:aws:ram:eu-west-1:123456789012:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12,123456789012
	"aws_ram_principal_association": FormattedIdentifierFromProvider(",", "resource_share_arn", "principal"),
	// RAM Resource Associations can be imported using their Resource Share ARN and Resource ARN separated by a comma:
	// arn:aws:ram:eu-west-1:123456789012:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12,arn:aws:ec2:eu-west-1:123456789012:subnet/subnet-12345678
	"aws_ram_resource_association": FormattedIdentifierFromProvider(",", "resource_share_arn", "resource_arn"),
	// Resource shares can be imported using the arn of the resource share:
	// aws_ram_resource_share.example arn:aws:ram:eu-west-1:123456789012:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12
	// TODO: validation may kick in, in which case we can use config.IdentifierFromProvider
	"aws_ram_resource_share": TemplatedStringAsIdentifierWithNoName("arn:aws:ram:{{ .parameters.region }}:{{ .setup.client_metadata.account_id }}:resource-share/{{ .external_name }}"),
	// Resource share accepters can be imported using the resource share ARN:
	// arn:aws:ram:us-east-1:123456789012:resource-share/c4b56393-e8d9-89d9-6dc9-883752de4767
	"aws_ram_resource_share_accepter": FormattedIdentifierFromProvider("", "share_arn"),

	// ecs
	//
	// ECS Task Sets can be imported via the task_set_id, service, and cluster separated by commas (,):
	// ecs-svc/7177320696926227436,arn:aws:ecs:us-west-2:123456789101:service/example/example-1234567890,arn:aws:ecs:us-west-2:123456789101:cluster/example
	// TODO: validation may kick in, in which case we can use config.IdentifierFromProvider
	"aws_ecs_task_set": TemplatedStringAsIdentifierWithNoName("{{ .external_name }},{{ .parameters.service }},{{ .parameters.cluster }}"),

	// gamelift
	//
	// GameLift Game Server Group can be imported using the name
	"aws_gamelift_game_server_group": config.ParameterAsIdentifier("game_server_group_name"),

	// guardduty
	//
	// GuardDuty detectors can be imported using the detector ID
	"aws_guardduty_detector": config.IdentifierFromProvider,
	// GuardDuty filters can be imported using the detector ID and filter's name separated by a colon
	// 00b00fd5aecc0ab60a708659477e9617:MyFilter
	"aws_guardduty_filter": config.TemplatedStringAsIdentifier("name", "{{ .parameters.detector_id }}:{{ .external_name }}"),
	// aws_guardduty_invite_accepter can be imported using the member GuardDuty detector ID
	"aws_guardduty_invite_accepter": FormattedIdentifierFromProvider("", "detector_id"),
	// GuardDuty IPSet can be imported using the primary GuardDuty detector ID and IPSet ID
	// 00b00fd5aecc0ab60a708659477e9617:123456789012
	"aws_guardduty_ipset": config.IdentifierFromProvider,
	// GuardDuty members can be imported using the primary GuardDuty detector ID and member AWS account ID
	// 00b00fd5aecc0ab60a708659477e9617:123456789012
	"aws_guardduty_member": config.IdentifierFromProvider,
	// GuardDuty Organization Admin Account can be imported using the AWS account ID
	"aws_guardduty_organization_admin_account": FormattedIdentifierFromProvider("", "admin_account_id"),
	// GuardDuty Organization Configurations can be imported using the GuardDuty Detector ID
	"aws_guardduty_organization_configuration": FormattedIdentifierFromProvider("", "detector_id"),
	// GuardDuty PublishingDestination can be imported using the master GuardDuty detector ID and PublishingDestinationID
	// a4b86f26fa42e7e7cf0d1c333ea77777:a4b86f27a0e464e4a7e0516d242f1234
	"aws_guardduty_publishing_destination": config.IdentifierFromProvider,
	// GuardDuty ThreatIntelSet can be imported using the primary GuardDuty detector ID and ThreatIntelSetID
	// 00b00fd5aecc0ab60a708659477e9617:123456789012
	"aws_guardduty_threatintelset": config.IdentifierFromProvider,

	// s3control
	//
	// S3 Control Buckets can be imported using Amazon Resource Name (ARN)
	// arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-12345678/bucket/example
	"aws_s3control_bucket": config.IdentifierFromProvider,
	// S3 Control Bucket Lifecycle Configurations can be imported using the Amazon Resource Name (ARN)
	// arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-12345678/bucket/example
	"aws_s3control_bucket_lifecycle_configuration": config.IdentifierFromProvider,
	// S3 Control Bucket Policies can be imported using the Amazon Resource Name (ARN)
	// arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-12345678/bucket/example
	"aws_s3control_bucket_policy": config.IdentifierFromProvider,
	// Multi-Region Access Points can be imported using the account_id and name of the Multi-Region Access Point separated by a colon (:)
	// Example: 123456789012:example
	"aws_s3control_multi_region_access_point": config.TemplatedStringAsIdentifier("", "{{ .parameters.account_id }}:{{ .parameters.details.name }}"),
	// Multi-Region Access Point Policies can be imported using the account_id and name of the Multi-Region Access Point separated by a colon (:)
	// Example: 123456789012:example
	"aws_s3control_multi_region_access_point_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.account_id }}:{{ .parameters.details.name }}"),
	// Object Lambda Access Points can be imported using the account_id and name, separated by a colon (:)
	// Example: 123456789012:example
	"aws_s3control_object_lambda_access_point": config.TemplatedStringAsIdentifier("name", "{{ .parameters.account_id }}:{{ .external_name }}"),
	// Object Lambda Access Point policies can be imported using the account_id and name, separated by a colon (:)
	// Example: 123456789012:example
	"aws_s3control_object_lambda_access_point_policy": config.TemplatedStringAsIdentifier("name", "{{ .parameters.account_id }}:{{ .external_name }}"),

	// elasticbeanstalk
	//
	// Elastic Beanstalk Applications can be imported using the name
	"aws_elastic_beanstalk_application_version": config.NameAsIdentifier,
	// Elastic Beanstalk Environments can be imported using the id
	"aws_elastic_beanstalk_environment": config.IdentifierFromProvider,

	// elasticsearch
	//
	// Elasticsearch domains can be imported using the domain_name
	"aws_elasticsearch_domain": config.ParameterAsIdentifier("domain_name"),
	// No import
	"aws_elasticsearch_domain_policy": config.IdentifierFromProvider,
	// Elasticsearch domains can be imported using the domain_name
	"aws_elasticsearch_domain_saml_options": config.ParameterAsIdentifier("domain_name"),

	// elbv2
	//
	// Listener Certificates can be imported by using the listener arn and certificate arn, separated by an underscore (_)
	// arn:aws:elasticloadbalancing:us-west-2:123456789012:listener/app/test/8e4497da625e2d8a/9ab28ade35828f96/67b3d2d36dd7c26b_arn:aws:iam::123456789012:server-certificate/tf-acc-test-6453083910015726063
	"aws_lb_listener_certificate": config.IdentifierFromProvider,

	// emr
	//
	// EMR clusters can be imported using the id
	"aws_emr_cluster": config.IdentifierFromProvider,
	// EMR Instance Fleet can be imported with the EMR Cluster identifier and Instance Fleet identifier separated by a forward slash (/)
	// j-123456ABCDEF/if-15EK4O09RZLNR
	"aws_emr_instance_fleet": config.IdentifierFromProvider,
	// EMR task instance group can be imported using their EMR Cluster id and Instance Group id separated by a forward-slash /
	// j-123456ABCDEF/ig-15EK4O09RZLNR
	"aws_emr_instance_group": config.IdentifierFromProvider,
	// EMR Managed Scaling Policies can be imported via the EMR Cluster identifier
	"aws_emr_managed_scaling_policy": FormattedIdentifierFromProvider("", "cluster_id"),
	// EMR studios can be imported using the id
	"aws_emr_studio": config.IdentifierFromProvider,
	// EMR studio session mappings can be imported using the id, e.g., studio-id:identity-type:identity-id
	"aws_emr_studio_session_mapping": config.IdentifierFromProvider,

	// emrcontainers
	//
	// EKS Clusters can be imported using the id
	"aws_emrcontainers_virtual_cluster": config.IdentifierFromProvider,

	// fms
	//
	// Firewall Manager administrator account association can be imported using the account ID
	// TODO: account_id parameter is not `Required` in TF schema. But we use this field in id construction. So, please mark as required this field while configuration
	"aws_fms_admin_account": FormattedIdentifierFromProvider("", "account_id"),
	// Firewall Manager policies can be imported using the policy ID
	"aws_fms_policy": config.IdentifierFromProvider,

	// fsx
	//
	// FSx File Systems can be imported using the id
	"aws_fsx_ontap_file_system": config.IdentifierFromProvider,
	// FSx Storage Virtual Machine can be imported using the id
	"aws_fsx_ontap_storage_virtual_machine": config.IdentifierFromProvider,
	// FSx ONTAP volume can be imported using the id
	"aws_fsx_ontap_volume": config.IdentifierFromProvider,
	// FSx File Systems can be imported using the id
	"aws_fsx_openzfs_file_system": config.IdentifierFromProvider,
	// FSx OpenZFS snapshot can be imported using the id
	"aws_fsx_openzfs_snapshot": config.IdentifierFromProvider,
	// FSx Volumes can be imported using the id
	"aws_fsx_openzfs_volume": config.IdentifierFromProvider,

	// iot
	//
	// IoT topic rule destinations can be imported using the arn
	// arn:aws:iot:us-west-2:123456789012:ruledestination/vpc/2ce781c8-68a6-4c52-9c62-63fe489ecc60
	"aws_iot_topic_rule_destination": config.IdentifierFromProvider,

	// kafka
	//
	// MSK SCRAM Secret Associations can be imported using the id
	"aws_msk_scram_secret_association": config.IdentifierFromProvider,

	// macie
	//
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_macie_member_account_association": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_macie_s3_bucket_association": config.IdentifierFromProvider,

	// macie2
	//
	// aws_macie2_organization_admin_account can be imported using the id
	"aws_macie2_organization_admin_account": config.IdentifierFromProvider,

	// memorydb
	//
	// Use the user_name to import a user
	"aws_memorydb_user": config.ParameterAsIdentifier("user_name"),

	// opsworks
	//
	// Opsworks Application can be imported using the id
	"aws_opsworks_application": config.IdentifierFromProvider,
	// OpsWorks Custom Layers can be imported using the id
	"aws_opsworks_custom_layer": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_opsworks_ecs_cluster_layer": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_opsworks_ganglia_layer": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_opsworks_haproxy_layer": config.IdentifierFromProvider,
	// Opsworks Instances can be imported using the instance id
	"aws_opsworks_instance": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_opsworks_java_app_layer": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_opsworks_memcached_layer": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_opsworks_mysql_layer": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_opsworks_nodejs_app_layer": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_opsworks_permission": config.IdentifierFromProvider,
	// OpsWorks PHP Application Layers can be imported using the id
	"aws_opsworks_php_app_layer": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_opsworks_rails_app_layer": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_opsworks_rds_db_instance": config.IdentifierFromProvider,
	// OpsWorks stacks can be imported using the id
	"aws_opsworks_stack": config.IdentifierFromProvider,
	// OpsWorks static web server Layers can be imported using the id
	"aws_opsworks_static_web_layer": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_opsworks_user_profile": config.IdentifierFromProvider,

	// ssm
	//
	// AWS Maintenance Window Task can be imported using the window_id and window_task_id separated by /
	"aws_ssm_maintenance_window_task": config.IdentifierFromProvider,
	// SSM Parameters can be imported using the parameter store name
	"aws_ssm_parameter": config.NameAsIdentifier,
	// SSM resource data sync can be imported using the name
	"aws_ssm_resource_data_sync": config.NameAsIdentifier,

	// pinpoint
	//
	// Pinpoint ADM Channel can be imported using the application-id
	"aws_pinpoint_adm_channel": FormattedIdentifierFromProvider("", "application_id"),
	// Pinpoint APNs Channel can be imported using the application-id
	"aws_pinpoint_apns_channel": FormattedIdentifierFromProvider("", "application_id"),
	// Pinpoint APNs Sandbox Channel can be imported using the application-id
	"aws_pinpoint_apns_sandbox_channel": FormattedIdentifierFromProvider("", "application_id"),
	// Pinpoint APNs VoIP Channel can be imported using the application-id
	"aws_pinpoint_apns_voip_channel": FormattedIdentifierFromProvider("", "application_id"),
	// Pinpoint APNs VoIP Sandbox Channel can be imported using the application-id
	"aws_pinpoint_apns_voip_sandbox_channel": FormattedIdentifierFromProvider("", "application_id"),
	// Pinpoint Baidu Channel can be imported using the application-id
	"aws_pinpoint_baidu_channel": FormattedIdentifierFromProvider("", "application_id"),
	// Pinpoint Email Channel can be imported using the application-id
	"aws_pinpoint_email_channel": FormattedIdentifierFromProvider("", "application_id"),
	// Pinpoint Event Stream can be imported using the application-id
	"aws_pinpoint_event_stream": FormattedIdentifierFromProvider("", "application_id"),
	// Pinpoint GCM Channel can be imported using the application-id
	"aws_pinpoint_gcm_channel": FormattedIdentifierFromProvider("", "application_id"),

	// quicksight
	//
	// A QuickSight data source can be imported using the AWS account ID, and data source ID name separated by a slash (/)
	// 123456789123/my-data-source-id
	"aws_quicksight_data_source": FormattedIdentifierFromProvider("/", "aws_account_id", "data_source_id"),
	// QuickSight Group membership can be imported using the AWS account ID, namespace, group name and member name separated by /
	// 123456789123/default/all-access-users/john_smith
	"aws_quicksight_group_membership": FormattedIdentifierFromProvider("/", "aws_account_id", "namespace", "group_name", "member_name"),

	// redshift
	//
	// Redshift security groups can be imported using the name
	"aws_redshift_security_group": config.NameAsIdentifier,
	// Redshift Snapshot Copy Grants support import by name
	"aws_redshift_snapshot_copy_grant": config.IdentifierFromProvider,

	// route53domains
	//
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_route53domains_registered_domain": config.IdentifierFromProvider,

	// s3outposts
	//
	// S3 Outposts Endpoints can be imported using Amazon Resource Name (ARN), EC2 Security Group identifier, and EC2 Subnet identifier, separated by commas (,)
	// arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-12345678/endpoint/0123456789abcdef,sg-12345678,subnet-12345678
	"aws_s3outposts_endpoint": config.IdentifierFromProvider,

	// sagemaker
	//
	// SageMaker Apps can be imported using the id
	"aws_sagemaker_app": config.IdentifierFromProvider,
	// SageMaker Devices can be imported using the device-fleet-name/device-name
	// my-fleet/my-device
	"aws_sagemaker_device": FormattedIdentifierFromProvider("/", "device_fleet_name", "device.device_name"),
	// SageMaker Device Fleets can be imported using the name
	"aws_sagemaker_device_fleet": config.ParameterAsIdentifier("device_fleet_name"),
	// Endpoints can be imported using the name
	"aws_sagemaker_endpoint": config.NameAsIdentifier,
	// Endpoint configurations can be imported using the name
	"aws_sagemaker_endpoint_configuration": config.NameAsIdentifier,
	// SageMaker Flow Definitions can be imported using the flow_definition_name
	"aws_sagemaker_flow_definition": config.ParameterAsIdentifier("flow_definition_name"),
	// SageMaker Human Task UIs can be imported using the human_task_ui_name
	"aws_sagemaker_human_task_ui": config.ParameterAsIdentifier("human_task_ui_name"),
	// SageMaker Code Images can be imported using the name
	"aws_sagemaker_image_version": config.ParameterAsIdentifier("image_name"),
	// Models can be imported using the name
	"aws_sagemaker_model": config.NameAsIdentifier,
	// SageMaker Model Package Groups can be imported using the name
	"aws_sagemaker_model_package_group_policy": config.ParameterAsIdentifier("model_package_group_name"),
	// SageMaker Projects can be imported using the project_name
	"aws_sagemaker_project": config.ParameterAsIdentifier("project_name"),
	// SageMaker Workforces can be imported using the workforce_name
	"aws_sagemaker_workforce": config.ParameterAsIdentifier("workforce_name"),
	// SageMaker Workteams can be imported using the workteam_name
	"aws_sagemaker_workteam": config.ParameterAsIdentifier("workteam_name"),

	// storagegateway
	//
	// aws_storagegateway_cache can be imported by using the gateway Amazon Resource Name (ARN) and local disk identifier separated with a colon (:)
	// Example: arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678:pci-0000:03:00.0-scsi-0:0:0:0
	"aws_storagegateway_cache": config.TemplatedStringAsIdentifier("", "{{ .parameters.gateway_arn }}:{{ .parameters.disk_id }}"),
	// aws_storagegateway_cached_iscsi_volume can be imported by using the volume Amazon Resource Name (ARN)
	// Example: arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678
	"aws_storagegateway_cached_iscsi_volume": config.TemplatedStringAsIdentifier("", "{{ .parameters.gateway_arn }}/volume/{{ .external_name }}"),
	// aws_storagegateway_file_system_association can be imported by using the FSx file system association Amazon Resource Name (ARN)
	// Example: arn:aws:storagegateway:us-east-1:123456789012:fs-association/fsa-0DA347732FDB40125
	"aws_storagegateway_file_system_association": config.TemplatedStringAsIdentifier("", "arn:aws:storagegateway:{{ .parameters.region }}:{{ .setup.client_metadata.account_id }}:fs-association/{{ .external_name }}"),
	// aws_storagegateway_gateway can be imported by using the gateway Amazon Resource Name (ARN)
	// Example: arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678
	"aws_storagegateway_gateway": config.TemplatedStringAsIdentifier("", "arn:aws:storagegateway:{{ .parameters.region }}:{{ .setup.client_metadata.account_id }}:gateway/{{ .external_name }}"),
	// aws_storagegateway_nfs_file_share can be imported by using the NFS File Share Amazon Resource Name (ARN)
	// Example: arn:aws:storagegateway:us-east-1:123456789012:share/share-12345678
	"aws_storagegateway_nfs_file_share": config.TemplatedStringAsIdentifier("", "arn:aws:storagegateway:{{ .parameters.region }}:{{ .setup.client_metadata.account_id }}:share/{{ .external_name }}"),
	// aws_storagegateway_smb_file_share can be imported by using the SMB File Share Amazon Resource Name (ARN)
	// Example: arn:aws:storagegateway:us-east-1:123456789012:share/share-12345678
	"aws_storagegateway_smb_file_share": config.TemplatedStringAsIdentifier("", "arn:aws:storagegateway:{{ .parameters.region }}:{{ .setup.client_metadata.account_id }}:share/{{ .external_name }}"),
	// aws_storagegateway_stored_iscsi_volume can be imported by using the volume Amazon Resource Name (ARN)
	// Example: arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678
	"aws_storagegateway_stored_iscsi_volume": config.TemplatedStringAsIdentifier("", "{{ .parameters.gateway_arn }}/volume/{{ .external_name }}"),
	// aws_storagegateway_tape_pool can be imported by using the volume Amazon Resource Name (ARN)
	// Example: arn:aws:storagegateway:us-east-1:123456789012:tapepool/pool-12345678
	"aws_storagegateway_tape_pool": config.TemplatedStringAsIdentifier("", "arn:aws:storagegateway:{{ .parameters.region }}:{{ .setup.client_metadata.account_id }}:tapepool/{{ .external_name }}"),
	// aws_storagegateway_upload_buffer can be imported by using the gateway Amazon Resource Name (ARN) and local disk identifier separated with a colon (:)
	// Example: arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678:pci-0000:03:00.0-scsi-0:0:0:0
	"aws_storagegateway_upload_buffer": config.TemplatedStringAsIdentifier("", "{{ .parameters.gateway_arn }}:{{ .parameters.disk_id }}"),
	// aws_storagegateway_working_storage can be imported by using the gateway Amazon Resource Name (ARN) and local disk identifier separated with a colon (:)
	// Example: arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678:pci-0000:03:00.0-scsi-0:0:0:0
	"aws_storagegateway_working_storage": config.TemplatedStringAsIdentifier("", "{{ .parameters.gateway_arn }}:{{ .parameters.disk_id }}"),

	// location
	//
	// aws_location_map resources can be imported using the map name
	"aws_location_map": config.ParameterAsIdentifier("map_name"),

	// mskconnect
	//
	// MSK Connect Connector can be imported using the connector's arn
	// Example: arn:aws:kafkaconnect:eu-central-1:123456789012:connector/example/264edee4-17a3-412e-bd76-6681cfc93805-3
	// TODO: Normalize external_name while testing resource
	"aws_mskconnect_connector": config.IdentifierFromProvider,
	// MSK Connect Custom Plugin can be imported using the plugin's arn
	// Example: arn:aws:kafkaconnect:eu-central-1:123456789012:custom-plugin/debezium-example/abcdefgh-1234-5678-9abc-defghijklmno-4
	// TODO: Normalize external_name while testing resource
	"aws_mskconnect_custom_plugin": config.IdentifierFromProvider,
	// MSK Connect Worker Configuration can be imported using the plugin's arn
	// Example: arn:aws:kafkaconnect:eu-central-1:123456789012:worker-configuration/example/8848493b-7fcc-478c-a646-4a52634e3378-4
	// TODO: Normalize external_name while testing resource
	"aws_mskconnect_worker_configuration": config.IdentifierFromProvider,

	// waf
	//
	// WAF Byte Match Set can be imported using the id
	"aws_waf_byte_match_set": config.IdentifierFromProvider,
	// WAF Geo Match Set can be imported using their ID
	"aws_waf_geo_match_set": config.IdentifierFromProvider,
	// WAF IPSets can be imported using their ID
	"aws_waf_ipset": config.IdentifierFromProvider,
	// WAF Rated Based Rule can be imported using the id
	"aws_waf_rate_based_rule": config.IdentifierFromProvider,
	// WAF Regex Match Set can be imported using their ID
	"aws_waf_regex_match_set": config.IdentifierFromProvider,
	// AWS WAF Regex Pattern Set can be imported using their ID
	"aws_waf_regex_pattern_set": config.IdentifierFromProvider,
	// WAF rules can be imported using the id
	"aws_waf_rule": config.IdentifierFromProvider,
	// AWS WAF Size Constraint Set can be imported using their ID
	"aws_waf_size_constraint_set": config.IdentifierFromProvider,
	// AWS WAF SQL Injection Match Set can be imported using their ID
	"aws_waf_sql_injection_match_set": config.IdentifierFromProvider,
	// WAF Web ACL can be imported using the id
	"aws_waf_web_acl": config.IdentifierFromProvider,
	// WAF XSS Match Set can be imported using their ID
	"aws_waf_xss_match_set": config.IdentifierFromProvider,

	// wafregional
	//
	// WAF Regional Byte Match Set can be imported using the id
	"aws_wafregional_byte_match_set": config.IdentifierFromProvider,
	// WAF Regional Geo Match Set can be imported using the id
	"aws_wafregional_geo_match_set": config.IdentifierFromProvider,
	// WAF Regional IPSets can be imported using their ID
	"aws_wafregional_ipset": config.IdentifierFromProvider,
	// WAF Regional Rate Based Rule can be imported using the id
	"aws_wafregional_rate_based_rule": config.IdentifierFromProvider,
	// WAF Regional Regex Match Set can be imported using the id
	"aws_wafregional_regex_match_set": config.IdentifierFromProvider,
	// WAF Regional Regex Pattern Set can be imported using the id
	"aws_wafregional_regex_pattern_set": config.IdentifierFromProvider,
	// WAF Regional Rule can be imported using the id
	"aws_wafregional_rule": config.IdentifierFromProvider,
	// WAF Size Constraint Set can be imported using the id
	"aws_wafregional_size_constraint_set": config.IdentifierFromProvider,
	// WAF Regional Sql Injection Match Set can be imported using the id
	"aws_wafregional_sql_injection_match_set": config.IdentifierFromProvider,
	// WAF Regional Web ACL can be imported using the id
	"aws_wafregional_web_acl": config.IdentifierFromProvider,
	// WAF Regional Web ACL Association can be imported using their web_acl_id:resource_arn
	"aws_wafregional_web_acl_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.web_acl_id }}:{{ .parameters.resource_arn }}"),
	// AWS WAF Regional XSS Match can be imported using the id
	"aws_wafregional_xss_match_set": config.IdentifierFromProvider,

	// ssoadmin
	//
	// SSO Account Assignments can be imported using the principal_id, principal_type, target_id, target_type, permission_set_arn, instance_arn separated by commas (,)
	// Example: f81d4fae-7dec-11d0-a765-00a0c91e6bf6,GROUP,1234567890,AWS_ACCOUNT,arn:aws:sso:::permissionSet/ssoins-0123456789abcdef/ps-0123456789abcdef,arn:aws:sso:::instance/ssoins-0123456789abcdef
	"aws_ssoadmin_account_assignment": config.TemplatedStringAsIdentifier("", "{{ .parameters.principal_id }},{{ .parameters.principal_type }},{{ .parameters.target_id }},{{ .parameters.target_type }},{{ .parameters.permission_set_arn }},{{ .parameters.instance_arn }}"),
	// SSO Managed Policy Attachments can be imported using the managed_policy_arn, permission_set_arn, and instance_arn separated by a comma (,)
	// Example: arn:aws:iam::aws:policy/AlexaForBusinessDeviceSetup,arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
	"aws_ssoadmin_managed_policy_attachment": config.TemplatedStringAsIdentifier("", "{{ .parameters.managed_policy_arn }},{{ .parameters.permission_set_arn }},{{ .parameters.instance_arn}}"),
	// SSO Permission Sets can be imported using the arn and instance_arn separated by a comma (,)
	// Example: arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
	// TODO: Normalize external_name while testing
	"aws_ssoadmin_permission_set": config.IdentifierFromProvider,
	// SSO Permission Set Inline Policies can be imported using the permission_set_arn and instance_arn separated by a comma (,)
	// Example: arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
	"aws_ssoadmin_permission_set_inline_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.permission_set_arn }},{{ .parameters.instance_arn }}"),

	// swf
	//
	// SWF Domains can be imported using the name
	"aws_swf_domain": config.NameAsIdentifier,

	// synthetics
	//
	// Synthetics Canaries can be imported using the name
	"aws_synthetics_canary": config.NameAsIdentifier,

	// timestreamwrite
	//
	// Timestream databases can be imported using the database_name
	"aws_timestreamwrite_database": config.ParameterAsIdentifier("database_name"),
	// Timestream tables can be imported using the table_name and database_name separate by a colon (:)
	// Example: ExampleTable:ExampleDatabase
	"aws_timestreamwrite_table": config.TemplatedStringAsIdentifier("", "{{ .parameters.table_name }}:{{ .parameters.database_name }}"),

	// networkfirewall
	//
	// Network Firewall Logging Configurations can be imported using the firewall_arn
	// Example: arn:aws:network-firewall:us-west-1:123456789012:firewall/example
	"aws_networkfirewall_logging_configuration": config.TemplatedStringAsIdentifier("", "arn:aws:network-firewall:{{ .parameters.region }}:{{ .setup.client_metadata.account_id }}:firewall/{{ .external_name }}"),
	// Network Firewall Resource Policies can be imported using the resource_arn
	// Example: arn:aws:network-firewall:us-west-1:123456789012:stateful-rulegroup/example
	"aws_networkfirewall_resource_policy": config.TemplatedStringAsIdentifier("", "arn:aws:network-firewall:{{ .parameters.region }}:{{ .setup.client_metadata.account_id }}:stateful-rulegroup/{{ .external_name }}"),

	// ses
	//
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_ses_domain_identity_verification": config.IdentifierFromProvider,

	// shield
	//
	// Shield protection resources can be imported by specifying their ID
	"aws_shield_protection": config.IdentifierFromProvider,
	// Shield protection group resources can be imported by specifying their protection group id
	"aws_shield_protection_group": config.ParameterAsIdentifier("protection_group_id"),
	// Shield protection health check association resources can be imported by specifying the shield_protection_id and health_check_arn
	// Example: ff9592dc-22f3-4e88-afa1-7b29fde9669a+arn:aws:route53:::healthcheck/3742b175-edb9-46bc-9359-f53e3b794b1b
	"aws_shield_protection_health_check_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.shield_protection_id }}+{{ .parameters.health_check_arn }}"),

	// transfer
	//
	// Transfer SSH Public Key can be imported using the server_id and user_name and ssh_public_key_id separated by /
	// Example: s-12345678/test-username/key-12345
	"aws_transfer_ssh_key": config.IdentifierFromProvider,
	// Transfer Workflows can be imported using the worflow_id
	"aws_transfer_workflow": config.IdentifierFromProvider,
	// Transfer Accesses can be imported using the server_id and external_id
	// Example: s-12345678/S-1-1-12-1234567890-123456789-1234567890-1234
	"aws_transfer_access": config.TemplatedStringAsIdentifier("", "{{ .parameters.server_id }}/{{ .parameters.external_id }}"),

	// s3
	//
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_s3_object_copy": config.IdentifierFromProvider,

	// wafv2
	//
	// WAFv2 IP Sets can be imported using ID/name/scope
	"aws_wafv2_ip_set": config.IdentifierFromProvider,
	// WAFv2 Regex Pattern Sets can be imported using ID/name/scope
	"aws_wafv2_regex_pattern_set": config.IdentifierFromProvider,
	// WAFv2 Rule Group can be imported using ID/name/scope
	"aws_wafv2_rule_group": config.IdentifierFromProvider,
	// WAFv2 Web ACLs can be imported using ID/Name/Scope
	"aws_wafv2_web_acl": config.IdentifierFromProvider,
	// WAFv2 Web ACL Association can be imported using WEB_ACL_ARN,RESOURCE_ARN
	// Example: arn:aws:wafv2:...7ce849ea,arn:aws:apigateway:...ages/name
	"aws_wafv2_web_acl_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.web_acl_arn }},{{ .parameters.resource_arn }}"),
	// WAFv2 Web ACL Logging Configurations can be imported using the WAFv2 Web ACL ARN
	// Example: arn:aws:wafv2:us-west-2:123456789012:regional/webacl/test-logs/a1b2c3d4-5678-90ab-cdef
	"aws_wafv2_web_acl_logging_configuration": config.IdentifierFromProvider,

	// worklink
	//
	// WorkLink can be imported using the ARN
	// Example: arn:aws:worklink::123456789012:fleet/example
	"aws_worklink_fleet": config.TemplatedStringAsIdentifier("name", "arn:aws:worklink::{{ .setup.client_metadata.account_id }}:fleet/{{ .external_name }}"),
	// WorkLink Website Certificate Authority can be imported using FLEET-ARN,WEBSITE-CA-ID
	// Example: arn:aws:worklink::123456789012:fleet/example,abcdefghijk
	"aws_worklink_website_certificate_authority_association": config.IdentifierFromProvider,

	// workspaces
	//
	// Workspaces directory can be imported using the directory ID
	"aws_workspaces_directory": config.IdentifierFromProvider,
	// WorkSpaces IP groups can be imported using their GroupID
	"aws_workspaces_ip_group": config.IdentifierFromProvider,
	// Workspaces can be imported using their ID
	"aws_workspaces_workspace": config.IdentifierFromProvider,

	// xray
	//
	// XRay Encryption Config can be imported using the region name
	"aws_xray_encryption_config": config.IdentifierFromProvider,
	// XRay Groups can be imported using the ARN
	// Example: arn:aws:xray:us-west-2:1234567890:group/example-group/TNGX7SW5U6QY36T4ZMOUA3HVLBYCZTWDIOOXY3CJAXTHSS3YCWUA
	"aws_xray_group": config.IdentifierFromProvider,
	// XRay Sampling Rules can be imported using the name
	"aws_xray_sampling_rule": config.ParameterAsIdentifier("rule_name"),

	// imagebuilder
	//
	// aws_imagebuilder_components resources can be imported by using the Amazon Resource Name (ARN)
	"aws_imagebuilder_component": config.IdentifierFromProvider,
}
