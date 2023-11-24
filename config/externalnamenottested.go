/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/crossplane/upjet/pkg/config"
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

	// ds
	//
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
	// aws_guardduty_invite_accepter can be imported using the member GuardDuty detector ID
	"aws_guardduty_invite_accepter": FormattedIdentifierFromProvider("", "detector_id"),
	// GuardDuty IPSet can be imported using the primary GuardDuty detector ID and IPSet ID
	// 00b00fd5aecc0ab60a708659477e9617:123456789012
	"aws_guardduty_ipset": config.IdentifierFromProvider,
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

	// elasticbeanstalk
	//
	// Elastic Beanstalk Environments can be imported using the id
	"aws_elastic_beanstalk_environment": config.IdentifierFromProvider,

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
	// Endpoints can be imported using the name
	"aws_sagemaker_endpoint": config.NameAsIdentifier,
	// SageMaker Flow Definitions can be imported using the flow_definition_name
	"aws_sagemaker_flow_definition": config.ParameterAsIdentifier("flow_definition_name"),
	// SageMaker Human Task UIs can be imported using the human_task_ui_name
	"aws_sagemaker_human_task_ui": config.ParameterAsIdentifier("human_task_ui_name"),
	// SageMaker Projects can be imported using the project_name
	"aws_sagemaker_project": config.ParameterAsIdentifier("project_name"),

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

	// inspector
	//
	// Inspector Assessment Targets can be imported via their Amazon Resource Name (ARN)
	// Example: arn:aws:inspector:us-east-1:123456789012:target/0-xxxxxxx
	"aws_inspector_assessment_target": config.TemplatedStringAsIdentifier("name", "arn:aws:inspector:{{ .parameters.region }}:{{ .setup.client_metadata.account_id }}:target/{{ .external_name }}"),
	// aws_inspector_assessment_template can be imported by using the template assessment ARN
	// Example: arn:aws:inspector:us-west-2:123456789012:target/0-9IaAzhGR/template/0-WEcjR8CH
	"aws_inspector_assessment_template": config.TemplatedStringAsIdentifier("name", "arn:aws:inspector:{{ .parameters.region }}:{{ .setup.client_metadata.account_id }}:target/{{ .parameters.target_arn }}/template/{{ .external_name }}"),
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_inspector_resource_group": config.IdentifierFromProvider,

	// wafregional
	//
	// WAF Regional Web ACL Association can be imported using their web_acl_id:resource_arn
	"aws_wafregional_web_acl_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.web_acl_id }}:{{ .parameters.resource_arn }}"),

	// synthetics
	//
	// Synthetics Canaries can be imported using the name
	"aws_synthetics_canary": config.NameAsIdentifier,

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
	// Transfer Accesses can be imported using the server_id and external_id
	// Example: s-12345678/S-1-1-12-1234567890-123456789-1234567890-1234
	"aws_transfer_access": config.TemplatedStringAsIdentifier("", "{{ .parameters.server_id }}/{{ .parameters.external_id }}"),

	// wafv2
	//
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
	// Workspaces can be imported using their ID
	"aws_workspaces_workspace": config.IdentifierFromProvider,

	// apprunner
	//
	// App Runner VPC Ingress Connection can be imported by using the arn
	// Example: arn:aws:apprunner:us-west-2:837424938642:vpcingressconnection/example/b379f86381d74825832c2e82080342fa
	// TODO: We just normalized the external-name but still kept the naming argument spec.forProvider.name. Need further normalization.
	"aws_apprunner_vpc_ingress_connection": TemplatedStringAsIdentifierWithNoName("arn:aws:apprunner:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:vpcingressconnection/{{ .parameters.name }}/{{ .external_name }}"),

	// appsync
	//
	// Appsync Types can be imported using the id (api-id:format:name)
	// TODO: Need further normalization spec.forProvider
	"aws_appsync_type": config.TemplatedStringAsIdentifier("", "{{ .parameters.api_id }}:{{ .parameters.format }}:{{ .external_name }}"),

	// auditmanager
	//
	// Audit Manager Account Registration resources can be imported using the id
	"aws_auditmanager_account_registration": config.IdentifierFromProvider,
	// Audit Manager Assessments can be imported using the assessment id (abc123-de45)
	// TODO: While testing check is name argument appear in the ID for these resource. If so, then normalize spec.forProvider.name.
	"aws_auditmanager_assessment": config.IdentifierFromProvider,
	// Audit Manager Assessment Reports can be imported using the assessment report id (abc123-de45)
	// TODO: While testing check is name argument appear in the ID for these resource. If so, then normalize spec.forProvider.name.
	"aws_auditmanager_assessment_report": config.IdentifierFromProvider,
	// An Audit Manager Control can be imported using the id (abc123-de45)
	// TODO: While testing check is name argument appear in the ID for these resource. If so, then normalize spec.forProvider.name.
	"aws_auditmanager_control": config.IdentifierFromProvider,
	// Audit Manager Framework can be imported using the framework id (abc123-de45)
	// TODO: While testing check is name argument appear in the ID for these resource. If so, then normalize spec.forProvider.name.
	"aws_auditmanager_framework": config.IdentifierFromProvider,

	// ce
	//
	// aws_ce_anomaly_subscription can be imported using the id
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_ce_anomaly_subscription": config.IdentifierFromProvider,
	// aws_ce_cost_allocation_tag can be imported using the id
	"aws_ce_cost_allocation_tag": config.ParameterAsIdentifier("tag_key"),

	// cloudwatch
	//
	// This resource can be imported using the log_group_name
	"aws_cloudwatch_log_data_protection_policy": config.ParameterAsIdentifier("log_group_name"),

	// comprehend
	//
	// Comprehend Document Classifier can be imported using the ARN
	// Example: arn:aws:comprehend:us-west-2:123456789012:document_classifier/example
	"aws_comprehend_document_classifier": config.TemplatedStringAsIdentifier("name", "arn:aws:comprehend:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:document_classifier/{{ .external_name }}"),
	// Comprehend Entity Recognizer can be imported using the ARN
	// Example: arn:aws:comprehend:us-west-2:123456789012:entity-recognizer/example
	"aws_comprehend_entity_recognizer": config.TemplatedStringAsIdentifier("name", "arn:aws:comprehend:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:entity-recognizer/{{ .external_name }}"),

	// controltower
	//
	// Control Tower Controls can be imported using their organizational_unit_arn/control_identifier
	// Example: arn:aws:organizations::123456789101:ou/o-qqaejywet/ou-qg5o-ufbhdtv3,arn:aws:controltower:us-east-1::control/WTDSMKDKDNLE
	"aws_controltower_control": config.TemplatedStringAsIdentifier("", "{{ .parameters.target_identifier }},{{ .external_name }}"),

	// datasync
	//
	// aws_datasync_location_object_storage can be imported by using the Amazon Resource Name (ARN)
	// Example: arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
	"aws_datasync_location_object_storage": config.TemplatedStringAsIdentifier("", "arn:aws:datasync:{{ .parameters.region }}:{{ .setup.client_metadata.account_id }}:location/{{ .external_name }}"),

	// directory_service
	//
	// RADIUS settings can be imported using the directory ID
	"aws_directory_service_radius_settings": config.IdentifierFromProvider,
	// Replicated Regions can be imported using directory ID,Region name
	"aws_directory_service_region": config.IdentifierFromProvider,
	// Directory Service Shared Directories can be imported using the shared directory ID
	"aws_directory_service_shared_directory_accepter": config.IdentifierFromProvider,

	// dx
	//
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_dx_macsec_key_association": config.IdentifierFromProvider,

	// ec2
	//
	// aws_ec2_transit_gateway_policy_table_association can be imported by using the EC2 Transit Gateway Policy Table identifier, an underscore, and the EC2 Transit Gateway Attachment identifier
	"aws_ec2_transit_gateway_policy_table_association": config.IdentifierFromProvider,

	// fsx
	//
	// Amazon File Cache cache can be imported using the resource id
	"aws_fsx_file_cache": config.IdentifierFromProvider,

	// identitystore
	//
	// An Identity Store Group can be imported using the combination identity_store_id/group_id
	"aws_identitystore_group": config.TemplatedStringAsIdentifier("", "{{ .parameters.identity_store_id }}/{{ .external_name }}"),
	// aws_identitystore_group_membership can be imported using the identity_store_id/membership_id
	"aws_identitystore_group_membership": config.TemplatedStringAsIdentifier("", "{{ .parameters.identity_store_id }}/{{ .external_name }}"),
	// An Identity Store User can be imported using the combination identity_store_id/user_id
	"aws_identitystore_user": config.TemplatedStringAsIdentifier("", "{{ .parameters.identity_store_id }}/{{ .external_name }}"),

	// inspector2
	//
	// Inspector V2 Delegated Admin Account can be imported using the account_id
	"aws_inspector2_delegated_admin_account": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	// TODO: Check if we need privilege to test this resource. If yes - split it with "Need privilege" label.
	"aws_inspector2_organization_configuration": config.IdentifierFromProvider,

	// ivs
	//
	// IVS (Interactive Video) Playback Key Pair can be imported using the ARN
	// Example: arn:aws:ivs:us-west-2:326937407773:playback-key/KDJRJNQhiQzA
	"aws_ivs_playback_key_pair": config.TemplatedStringAsIdentifier("", "arn:aws:ivs:{{ .parameters.region }}:{{ .setup.client_metadata.account_id }}:playback-key/{{ .external_name }}"),

	// ivschat
	//
	// IVS (Interactive Video) Chat Logging Configuration can be imported using the ARN
	// Example: arn:aws:ivschat:us-west-2:326937407773:logging-configuration/MMUQc8wcqZmC
	"aws_ivschat_logging_configuration": config.TemplatedStringAsIdentifier("", "arn:aws:ivschat:{{ .parameters.region }}:{{ .setup.configuration.account_id }}:logging-configuration/{{ .external_name }}"),
	// IVS (Interactive Video) Chat Room can be imported using the ARN
	// Example: arn:aws:ivschat:us-west-2:326937407773:room/GoXEXyB4VwHb
	"aws_ivschat_room": config.TemplatedStringAsIdentifier("", "arn:aws:ivschat:{{ .parameters.region }}:{{ .setup.configuration.account_id }}:room/{{ .external_name }}"),

	// kms
	//
	// KMS (Key Management) Custom Key Store can be imported using the id
	"aws_kms_custom_key_store": config.IdentifierFromProvider,

	// lakeformation
	//
	// Lake Formation LF-Tags can be imported using the catalog_id:key
	"aws_lakeformation_lf_tag": config.TemplatedStringAsIdentifier("", "{{ .external_name }}:{{ .parameters.key }}}"),
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_lakeformation_resource_lf_tags": config.IdentifierFromProvider,

	// lightsail
	//
	// Lightsail Container Service Deployment Version can be imported using the service_name and version separated by a slash (/)
	"aws_lightsail_container_service_deployment_version": config.TemplatedStringAsIdentifier("", "{{ .parameters.service_name }}/{{ .external_name }}"),
	// aws_lightsail_lb_certificate_attachment can be imported by using the id attribute
	// ID: lb_name,certificate_name
	"aws_lightsail_lb_certificate_attachment": config.IdentifierFromProvider,
	// aws_lightsail_lb_https_redirection_policy can be imported by using the lb_name attribute
	"aws_lightsail_lb_https_redirection_policy": config.ParameterAsIdentifier("lb_name"),
	// Lightsail Databases can be imported using their name
	"aws_lightsail_database": config.IdentifierFromProvider,

	// macie2
	//
	// aws_macie2_classification_export_configuration can be imported using the account ID and region
	"aws_macie2_classification_export_configuration": config.IdentifierFromProvider,

	// kendra
	//
	// aws_kendra_faq can be imported using the unique identifiers of the FAQ and index separated by a slash (/)
	// "aws_kendra_faq": config.TemplatedStringAsIdentifier("", "{{ .external_name }}/{{ .parameters.index_id }}}"),
	"aws_kendra_faq": config.IdentifierFromProvider,

	// medialive
	//
	// MediaLive MultiplexProgram can be imported using the id, or a combination of "program_name/multiplex_id"
	"aws_medialive_multiplex_program": config.IdentifierFromProvider,

	// networkmanager
	//
	// aws_networkmanager_site_to_site_vpn_attachment can be imported using the attachment ID
	"aws_networkmanager_site_to_site_vpn_attachment": config.IdentifierFromProvider,
	// aws_networkmanager_transit_gateway_peering can be imported using the peering ID
	"aws_networkmanager_transit_gateway_peering": config.IdentifierFromProvider,
	// aws_networkmanager_transit_gateway_route_table_attachment can be imported using the attachment ID
	"aws_networkmanager_transit_gateway_route_table_attachment": config.IdentifierFromProvider,

	// opensearch
	//
	// AWS Opensearch Inbound Connection Accepters can be imported by using the Inbound Connection ID
	"aws_opensearch_inbound_connection_accepter": config.ParameterAsIdentifier("connection_id"),
	// AWS Opensearch Outbound Connections can be imported by using the Outbound Connection ID
	"aws_opensearch_outbound_connection": config.IdentifierFromProvider,

	// rds
	//
	// A RDS (Relational Database) Export Task can be imported using the export_task_identifier
	"aws_rds_export_task": config.ParameterAsIdentifier("export_task_identifier"),
	// Due to the expense of testing this resource, we move it to skiplist.
	// RDS DB Instance Reservations can be imported using the instance_id
	"aws_rds_reserved_instance": config.IdentifierFromProvider,

	// redshift
	//

	// Redshift Cluster IAM Roless can be imported using the cluster_identifier
	"aws_redshift_cluster_iam_roles": config.ParameterAsIdentifier("cluster_identifier"),
	// Redshift endpoint access can be imported using the name
	"aws_redshift_endpoint_access": config.ParameterAsIdentifier("endpoint_name"),
	// Redshift endpoint authorization can be imported using the id
	// Example: 01234567910:cluster-example-id
	"aws_redshift_endpoint_authorization": config.TemplatedStringAsIdentifier("", "{{ .parameters.account }}:{{ .external_name }}"),
	// Redshift usage limits can be imported using the id
	// Example: 01234567910:cluster-example-id:example:example
	"aws_redshift_partner": config.TemplatedStringAsIdentifier("", "{{ .parameters.account_id }}:{{ .parameters.cluster_identifier }}:{{ .external_name }}"),
	// Redshift Data Statements can be imported using the id
	"aws_redshiftdata_statement": config.IdentifierFromProvider,

	// resourceexplorer2
	//
	// Resource Explorer indexes can be imported using the arn
	// Example: arn:aws:resource-explorer-2:us-east-1:123456789012:index/6047ac4e-207e-4487-9bcf-cb53bb0ff5cc
	"aws_resourceexplorer2_index": config.IdentifierFromProvider,
	// Resource Explorer views can be imported using the arn
	// Example: arn:aws:resource-explorer-2:us-west-2:123456789012:view/exampleview/e0914f6c-6c27-4b47-b5d4-6b28381a2421
	"aws_resourceexplorer2_view": config.IdentifierFromProvider,

	// rolesanywhere
	//
	// aws_rolesanywhere_trust_anchor can be imported using its id
	"aws_rolesanywhere_trust_anchor": config.IdentifierFromProvider,

	// schemas
	//
	// EventBridge Schema Registry Policy can be imported using the registry_name
	"aws_schemas_registry_policy": config.ParameterAsIdentifier("registry_name"),

	// sesv2
	//
	// SESv2 (Simple Email V2) Dedicated IP Assignment can be imported using the id, which is a comma-separated string made up of ip and destination_pool_name
	"aws_sesv2_dedicated_ip_assignment": config.IdentifierFromProvider,

	// ssoadmin
	//
	// SSO Managed Policy Attachments can be imported using the name, path, permission_set_arn, and instance_arn separated by a comma (,)
	// Example: TestPolicy,/,arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
	"aws_ssoadmin_customer_managed_policy_attachment": config.IdentifierFromProvider,
	// SSO Account Assignments can be imported using the instance_arn
	"aws_ssoadmin_instance_access_control_attributes": config.ParameterAsIdentifier("instance_arn"),
	// SSO Admin Permissions Boundary Attachments can be imported using the permission_set_arn and instance_arn, separated by a comma (,)
	"aws_ssoadmin_permissions_boundary_attachment": config.IdentifierFromProvider,

	// transcribe
	//
	// Transcribe MedicalVocabulary can be imported using the vocabulary_name
	"aws_transcribe_medical_vocabulary": config.ParameterAsIdentifier("vocabulary_name"),
}
