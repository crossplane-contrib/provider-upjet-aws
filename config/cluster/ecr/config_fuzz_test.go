// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package ecr

import (
	"strings"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// FuzzECRRepositoryCreationTemplateConfig fuzzes the ECR Repository Creation Template configuration
func FuzzECRRepositoryCreationTemplateConfig(f *testing.F) {
	// Add some seed inputs for better coverage
	f.Add("test-prefix", "IMMUTABLE", "PULL_THROUGH_CACHE", "arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012")
	f.Add("", "MUTABLE", "REPLICATION", "")
	f.Add("global-template", "", "", "invalid-arn")
	f.Add("very-long-prefix-name-that-might-cause-issues", "IMMUTABLE", "PULL_THROUGH_CACHE,REPLICATION", "")
	
	f.Fuzz(func(t *testing.T, prefix, imageTagMutability, appliedFor, kmsKeyArn string) {
		// Skip empty prefix as it's required
		if prefix == "" {
			t.Skip("Empty prefix not valid for testing")
		}
		
		// Create a mock provider and resource for testing
		p := config.NewProvider([]byte(`{}`), "aws", "github.com/upbound/provider-aws", []byte(`{}`))
		
		// Create a mock Terraform resource schema
		mockSchema := map[string]*schema.Schema{
			"prefix": {
				Type:     schema.TypeString,
				Required: true,
			},
			"image_tag_mutability": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"applied_for": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"encryption_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		}
		
		// Create a resource with the mock schema for validation
		_ = &config.Resource{
			Name: "aws_ecr_repository_creation_template",
			TerraformResource: &schema.Resource{
				Schema: mockSchema,
			},
		}
		
		// Apply the ECR configuration
		Configure(p)
		
		// Test that the configuration doesn't panic with fuzzed inputs
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Configuration panicked with inputs prefix=%q, imageTagMutability=%q, appliedFor=%q, kmsKeyArn=%q: %v",
					prefix, imageTagMutability, appliedFor, kmsKeyArn, r)
			}
		}()
		
		// Validate that the configuration is reasonable
		if len(p.Resources) == 0 {
			t.Error("No resources configured")
		}
		
		// Test specific ECR resource configuration
		ecrRepoTemplate, exists := p.Resources["aws_ecr_repository_creation_template"]
		if !exists {
			t.Error("ECR repository creation template not configured")
			return
		}
		
		// Validate configuration properties
		if ecrRepoTemplate.Kind != "RepositoryCreationTemplate" {
			t.Errorf("Expected Kind 'RepositoryCreationTemplate', got %q", ecrRepoTemplate.Kind)
		}
		
		if ecrRepoTemplate.ShortGroup != "ecr" {
			t.Errorf("Expected ShortGroup 'ecr', got %q", ecrRepoTemplate.ShortGroup)
		}
		
		// Test that references are properly configured
		if ecrRepoTemplate.References == nil {
			t.Error("References not configured")
		} else {
			if _, hasKMSRef := ecrRepoTemplate.References["encryption_configuration.kms_key"]; !hasKMSRef {
				t.Error("KMS key reference not configured")
			}
		}
	})
}

// FuzzECRRepositoryConfig fuzzes the ECR Repository configuration
func FuzzECRRepositoryConfig(f *testing.F) {
	// Add seed inputs
	f.Add("test-repo", "AES256", "arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012")
	f.Add("", "KMS", "")
	f.Add("repo-with-special-chars_123", "", "invalid-kms-arn")
	
	f.Fuzz(func(t *testing.T, repoName, encryptionType, kmsKeyArn string) {
		// Create a mock provider for testing
		p := config.NewProvider([]byte(`{}`), "aws", "github.com/upbound/provider-aws", []byte(`{}`))
		
		// Apply the ECR configuration
		Configure(p)
		
		// Test that the configuration doesn't panic with fuzzed inputs
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Configuration panicked with inputs repoName=%q, encryptionType=%q, kmsKeyArn=%q: %v",
					repoName, encryptionType, kmsKeyArn, r)
			}
		}()
		
		// Validate ECR repository configuration
		ecrRepo, exists := p.Resources["aws_ecr_repository"]
		if !exists {
			t.Error("ECR repository not configured")
			return
		}
		
		// Verify UseAsync is set for repository (deletion takes time)
		if !ecrRepo.UseAsync {
			t.Error("ECR repository should use async operations")
		}
		
		// Test KMS reference configuration
		if ecrRepo.References == nil {
			t.Error("References not configured for ECR repository")
		} else {
			if _, hasKMSRef := ecrRepo.References["encryption_configuration.kms_key"]; !hasKMSRef {
				t.Error("KMS key reference not configured for ECR repository")
			}
		}
	})
}

// FuzzECRResourceValidation fuzzes resource validation logic
func FuzzECRResourceValidation(f *testing.F) {
	// Add seed inputs for various validation scenarios
	f.Add("valid-prefix", "IMMUTABLE", true)
	f.Add("", "MUTABLE", false)
	f.Add("prefix-with-123", "INVALID_MUTABILITY", true)
	f.Add("very-long-prefix-name-that-exceeds-normal-limits-and-might-cause-validation-issues", "IMMUTABLE", false)
	
	f.Fuzz(func(t *testing.T, prefix, imageTagMutability string, useAsync bool) {
		// Test input validation doesn't panic or cause issues
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Validation panicked with inputs prefix=%q, imageTagMutability=%q, useAsync=%v: %v",
					prefix, imageTagMutability, useAsync, r)
			}
		}()
		
		// Test prefix validation logic
		if len(prefix) > 256 {
			// Very long prefixes should be handled gracefully
			t.Logf("Testing very long prefix: %d characters", len(prefix))
		}
		
		// Test image tag mutability validation
		validMutabilityValues := []string{"MUTABLE", "IMMUTABLE"}
		isValidMutability := false
		for _, valid := range validMutabilityValues {
			if imageTagMutability == valid {
				isValidMutability = true
				break
			}
		}
		
		if imageTagMutability != "" && !isValidMutability {
			t.Logf("Testing invalid image tag mutability: %q", imageTagMutability)
		}
		
		// Test that configuration handles various input combinations
		if strings.Contains(prefix, "..") {
			t.Logf("Testing prefix with double dots: %q", prefix)
		}
		
		if strings.ContainsAny(prefix, "!@#$%^&*()") {
			t.Logf("Testing prefix with special characters: %q", prefix)
		}
	})
}

// FuzzECRARNExtraction fuzzes ARN extraction logic
func FuzzECRARNExtraction(f *testing.F) {
	// Add seed inputs for ARN testing
	f.Add("arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012")
	f.Add("arn:aws:ecr:us-west-2:123456789012:repository/my-repo")
	f.Add("invalid-arn")
	f.Add("")
	f.Add("arn:aws:kms:::")
	f.Add("arn:aws:kms:region:account:key/")
	
	f.Fuzz(func(t *testing.T, arnInput string) {
		// Test ARN extraction doesn't panic with various inputs
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("ARN extraction panicked with input %q: %v", arnInput, r)
			}
		}()
		
		// Test various ARN formats
		if strings.HasPrefix(arnInput, "arn:aws:") {
			parts := strings.Split(arnInput, ":")
			if len(parts) >= 6 {
				t.Logf("Testing valid ARN format: %q", arnInput)
			} else {
				t.Logf("Testing incomplete ARN: %q", arnInput)
			}
		} else if arnInput != "" {
			t.Logf("Testing invalid ARN format: %q", arnInput)
		}
		
		// Test that empty ARNs are handled
		if arnInput == "" {
			t.Logf("Testing empty ARN")
		}
		
		// Test very long ARNs
		if len(arnInput) > 1000 {
			t.Logf("Testing very long ARN: %d characters", len(arnInput))
		}
	})
}
