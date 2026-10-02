// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package agentregistry

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/cluster/common"
)

// Configure adds configurations for the agentregistry group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_agentregistry_registry", func(r *config.Resource) {
		r.References["encryption_configuration.kms_key_arn"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
		r.AddSingletonListConversion("approval_configuration", "approvalConfiguration")
		r.AddSingletonListConversion("auto_detection_configuration", "autoDetectionConfiguration")
		r.AddSingletonListConversion("discovery_configuration", "discoveryConfiguration")
		r.AddSingletonListConversion("discovery_configuration[*].authorizer_configuration", "discoveryConfiguration[*].authorizerConfiguration")
		r.AddSingletonListConversion("discovery_configuration[*].authorizer_configuration[*].custom_jwt_authorizer", "discoveryConfiguration[*].authorizerConfiguration[*].customJwtAuthorizer")
		r.AddSingletonListConversion("discovery_configuration[*].authorizer_configuration[*].custom_jwt_authorizer[*].custom_claim[*].authorizing_claim_match_value", "discoveryConfiguration[*].authorizerConfiguration[*].customJwtAuthorizer[*].customClaim[*].authorizingClaimMatchValue")
		r.AddSingletonListConversion("discovery_configuration[*].authorizer_configuration[*].custom_jwt_authorizer[*].custom_claim[*].authorizing_claim_match_value[*].claim_match_value", "discoveryConfiguration[*].authorizerConfiguration[*].customJwtAuthorizer[*].customClaim[*].authorizingClaimMatchValue[*].claimMatchValue")
		r.AddSingletonListConversion("discovery_configuration[*].authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint", "discoveryConfiguration[*].authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpoint")
		r.AddSingletonListConversion("discovery_configuration[*].authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint[*].managed_vpc_resource", "discoveryConfiguration[*].authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpoint[*].managedVpcResource")
		r.AddSingletonListConversion("discovery_configuration[*].authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint[*].self_managed_lattice_resource", "discoveryConfiguration[*].authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpoint[*].selfManagedLatticeResource")
		r.AddSingletonListConversion("discovery_configuration[*].authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint_override[*].private_endpoint", "discoveryConfiguration[*].authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpointOverride[*].privateEndpoint")
		r.AddSingletonListConversion("discovery_configuration[*].authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint_override[*].private_endpoint[*].managed_vpc_resource", "discoveryConfiguration[*].authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpointOverride[*].privateEndpoint[*].managedVpcResource")
		r.AddSingletonListConversion("discovery_configuration[*].authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint_override[*].private_endpoint[*].self_managed_lattice_resource", "discoveryConfiguration[*].authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpointOverride[*].privateEndpoint[*].selfManagedLatticeResource")
		r.AddSingletonListConversion("encryption_configuration", "encryptionConfiguration")
	})
}
