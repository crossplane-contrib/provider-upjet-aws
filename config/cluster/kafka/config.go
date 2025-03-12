// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package kafka

import (
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/config/conversion"

	"github.com/upbound/provider-aws/apis/cluster/kafka/v1beta1"
	"github.com/upbound/provider-aws/apis/cluster/kafka/v1beta2"
	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the kafka group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_msk_cluster", func(r *config.Resource) {
		r.References["encryption_info.encryption_at_rest_kms_key_arn"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
		r.References["logging_info.broker_logs.s3.bucket"] = config.Reference{
			TerraformName: "aws_s3_bucket",
		}
		r.References["logging_info.broker_logs.cloudwatch_logs.log_group"] = config.Reference{
			TerraformName: "aws_cloudwatch_log_group",
		}
		r.References["broker_node_group_info.client_subnets"] = config.Reference{
			TerraformName: "aws_subnet",
		}
		r.References["broker_node_group_info.security_groups"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		r.References["configuration_info.arn"] = config.Reference{
			TerraformName: "aws_msk_configuration",
			Extractor:     common.PathARNExtractor,
		}
		r.UseAsync = true

		r.Version = "v1beta2"
		r.Conversions = append(r.Conversions,
			conversion.NewCustomConverter("v1beta1", "v1beta2", func(src, target xpresource.Managed) error {
				srcTyped := src.(*v1beta1.Cluster)
				targetTyped := target.(*v1beta2.Cluster)
				if len(srcTyped.Spec.ForProvider.BrokerNodeGroupInfo) > 0 {
					if srcTyped.Spec.ForProvider.BrokerNodeGroupInfo[0].EBSVolumeSize != nil {
						if targetTyped.Spec.ForProvider.BrokerNodeGroupInfo != nil {
							if targetTyped.Spec.ForProvider.BrokerNodeGroupInfo[0].StorageInfo != nil {
								if targetTyped.Spec.ForProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo != nil {
									targetTyped.Spec.ForProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo[0].VolumeSize = srcTyped.Spec.ForProvider.BrokerNodeGroupInfo[0].EBSVolumeSize
								} else {
									targetTyped.Spec.ForProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo = []v1beta2.EBSStorageInfoParameters{
										{
											VolumeSize: srcTyped.Spec.ForProvider.BrokerNodeGroupInfo[0].EBSVolumeSize,
										},
									}
								}
							} else {
								targetTyped.Spec.ForProvider.BrokerNodeGroupInfo[0].StorageInfo = []v1beta2.StorageInfoParameters{
									{
										EBSStorageInfo: []v1beta2.EBSStorageInfoParameters{
											{
												VolumeSize: srcTyped.Spec.ForProvider.BrokerNodeGroupInfo[0].EBSVolumeSize,
											},
										},
									},
								}
							}
						} else {
							targetTyped.Spec.ForProvider.BrokerNodeGroupInfo = []v1beta2.BrokerNodeGroupInfoParameters{
								{
									StorageInfo: []v1beta2.StorageInfoParameters{
										{
											EBSStorageInfo: []v1beta2.EBSStorageInfoParameters{
												{
													VolumeSize: srcTyped.Spec.ForProvider.BrokerNodeGroupInfo[0].EBSVolumeSize,
												},
											},
										},
									},
								},
							}
						}
					}
				}

				if len(srcTyped.Spec.InitProvider.BrokerNodeGroupInfo) > 0 {
					if srcTyped.Spec.InitProvider.BrokerNodeGroupInfo[0].EBSVolumeSize != nil {
						if targetTyped.Spec.InitProvider.BrokerNodeGroupInfo != nil {
							if targetTyped.Spec.InitProvider.BrokerNodeGroupInfo[0].StorageInfo != nil {
								if targetTyped.Spec.InitProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo != nil {
									targetTyped.Spec.InitProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo[0].VolumeSize = srcTyped.Spec.InitProvider.BrokerNodeGroupInfo[0].EBSVolumeSize
								} else {
									targetTyped.Spec.InitProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo = []v1beta2.EBSStorageInfoInitParameters{
										{
											VolumeSize: srcTyped.Spec.InitProvider.BrokerNodeGroupInfo[0].EBSVolumeSize,
										},
									}
								}
							} else {
								targetTyped.Spec.InitProvider.BrokerNodeGroupInfo[0].StorageInfo = []v1beta2.StorageInfoInitParameters{
									{
										EBSStorageInfo: []v1beta2.EBSStorageInfoInitParameters{
											{
												VolumeSize: srcTyped.Spec.InitProvider.BrokerNodeGroupInfo[0].EBSVolumeSize,
											},
										},
									},
								}
							}
						} else {
							targetTyped.Spec.InitProvider.BrokerNodeGroupInfo = []v1beta2.BrokerNodeGroupInfoInitParameters{
								{
									StorageInfo: []v1beta2.StorageInfoInitParameters{
										{
											EBSStorageInfo: []v1beta2.EBSStorageInfoInitParameters{
												{
													VolumeSize: srcTyped.Spec.InitProvider.BrokerNodeGroupInfo[0].EBSVolumeSize,
												},
											},
										},
									},
								},
							}
						}
					}
				}

				if len(srcTyped.Status.AtProvider.BrokerNodeGroupInfo) > 0 {
					if srcTyped.Status.AtProvider.BrokerNodeGroupInfo[0].EBSVolumeSize != nil {
						if targetTyped.Status.AtProvider.BrokerNodeGroupInfo != nil {
							if targetTyped.Status.AtProvider.BrokerNodeGroupInfo[0].StorageInfo != nil {
								if targetTyped.Status.AtProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo != nil {
									targetTyped.Status.AtProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo[0].VolumeSize = srcTyped.Status.AtProvider.BrokerNodeGroupInfo[0].EBSVolumeSize
								} else {
									targetTyped.Status.AtProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo = []v1beta2.EBSStorageInfoObservation{
										{
											VolumeSize: srcTyped.Status.AtProvider.BrokerNodeGroupInfo[0].EBSVolumeSize,
										},
									}
								}
							} else {
								targetTyped.Status.AtProvider.BrokerNodeGroupInfo[0].StorageInfo = []v1beta2.StorageInfoObservation{
									{
										EBSStorageInfo: []v1beta2.EBSStorageInfoObservation{
											{
												VolumeSize: srcTyped.Status.AtProvider.BrokerNodeGroupInfo[0].EBSVolumeSize,
											},
										},
									},
								}
							}
						} else {
							targetTyped.Status.AtProvider.BrokerNodeGroupInfo = []v1beta2.BrokerNodeGroupInfoObservation{
								{
									StorageInfo: []v1beta2.StorageInfoObservation{
										{
											EBSStorageInfo: []v1beta2.EBSStorageInfoObservation{
												{
													VolumeSize: srcTyped.Status.AtProvider.BrokerNodeGroupInfo[0].EBSVolumeSize,
												},
											},
										},
									},
								},
							}
						}
					}
				}
				return nil
			}),
			conversion.NewCustomConverter("v1beta2", "v1beta1", func(src, target xpresource.Managed) error {
				srcTyped := src.(*v1beta2.Cluster)
				targetTyped := target.(*v1beta1.Cluster)
				if srcTyped.Spec.ForProvider.BrokerNodeGroupInfo != nil {
					if srcTyped.Spec.ForProvider.BrokerNodeGroupInfo[0].StorageInfo != nil {
						if srcTyped.Spec.ForProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo != nil {
							if targetTyped.Spec.ForProvider.BrokerNodeGroupInfo != nil {
								targetTyped.Spec.ForProvider.BrokerNodeGroupInfo[0].EBSVolumeSize = srcTyped.Spec.ForProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo[0].VolumeSize
							} else {
								targetTyped.Spec.ForProvider.BrokerNodeGroupInfo = []v1beta1.BrokerNodeGroupInfoParameters{
									{
										EBSVolumeSize: srcTyped.Spec.ForProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo[0].VolumeSize,
									},
								}
							}
						}
					}
				}

				if srcTyped.Spec.InitProvider.BrokerNodeGroupInfo != nil {
					if srcTyped.Spec.InitProvider.BrokerNodeGroupInfo[0].StorageInfo != nil {
						if srcTyped.Spec.InitProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo != nil {
							if targetTyped.Spec.InitProvider.BrokerNodeGroupInfo != nil {
								targetTyped.Spec.InitProvider.BrokerNodeGroupInfo[0].EBSVolumeSize = srcTyped.Spec.InitProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo[0].VolumeSize
							} else {
								targetTyped.Spec.InitProvider.BrokerNodeGroupInfo = []v1beta1.BrokerNodeGroupInfoInitParameters{
									{
										EBSVolumeSize: srcTyped.Spec.InitProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo[0].VolumeSize,
									},
								}
							}
						}
					}
				}

				if srcTyped.Status.AtProvider.BrokerNodeGroupInfo != nil {
					if srcTyped.Status.AtProvider.BrokerNodeGroupInfo[0].StorageInfo != nil {
						if srcTyped.Status.AtProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo != nil {
							if targetTyped.Status.AtProvider.BrokerNodeGroupInfo != nil {
								targetTyped.Status.AtProvider.BrokerNodeGroupInfo[0].EBSVolumeSize = srcTyped.Status.AtProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo[0].VolumeSize
							} else {
								targetTyped.Status.AtProvider.BrokerNodeGroupInfo = []v1beta1.BrokerNodeGroupInfoObservation{
									{
										EBSVolumeSize: srcTyped.Status.AtProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo[0].VolumeSize,
									},
								}
							}
						}
					}
				}
				return nil
			}),
		)
	})
	p.AddResourceConfigurator("aws_msk_scram_secret_association", func(r *config.Resource) {
		r.References["secret_arn_list"] = config.Reference{
			TerraformName:     "aws_secretsmanager_secret",
			RefFieldName:      "SecretArnRefs",
			SelectorFieldName: "SecretArnSelector",
		}
		r.MetaResource.ArgumentDocs["secret_arn_list"] = "- (Required) List of all AWS Secrets Manager secret ARNs to associate with the cluster. Secrets not referenced, selected or listed here will be disassociated from the cluster."
	})
	p.AddResourceConfigurator("aws_msk_serverless_cluster", func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_config.security_group_ids"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SecurityGroupIDRefs",
			SelectorFieldName: "SecurityGroupIDSelector",
		}
		r.References["vpc_config.subnet_ids"] = config.Reference{
			TerraformName:     "aws_subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
		r.OverrideFieldNames = map[string]string{
			"ClientAuthenticationParameters":     "ServerlessClusterClientAuthenticationParameters",
			"ClientAuthenticationInitParameters": "ServerlessClusterClientAuthenticationInitParameters",
			"ClientAuthenticationObservation":    "ServerlessClusterClientAuthenticationObservation",
			"SaslParameters":                     "ClientAuthenticationSaslParameters",
			"SaslInitParameters":                 "ClientAuthenticationSaslInitParameters",
			"SaslObservation":                    "ClientAuthenticationSaslObservation",
		}
	})
}
