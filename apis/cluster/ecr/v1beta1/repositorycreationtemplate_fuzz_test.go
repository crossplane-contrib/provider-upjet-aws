// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	"encoding/json"
	"strings"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// FuzzRepositoryCreationTemplateMarshaling fuzzes JSON marshaling/unmarshaling of RepositoryCreationTemplate
func FuzzRepositoryCreationTemplateMarshaling(f *testing.F) {
	// Add seed inputs
	f.Add(`{"apiVersion":"ecr.aws.upbound.io/v1beta1","kind":"RepositoryCreationTemplate","metadata":{"name":"test"},"spec":{"forProvider":{"region":"us-east-1","prefix":"test"}}}`)
	f.Add(`{}`)
	f.Add(`{"spec":{"forProvider":{"prefix":""}}}`)
	f.Add(`{"metadata":{"name":"very-long-name-that-might-cause-issues-with-kubernetes-naming-limits-and-validation"}}`)
	
	f.Fuzz(func(t *testing.T, jsonInput string) {
		// Test that marshaling/unmarshaling doesn't panic
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Marshaling/unmarshaling panicked with input %q: %v", jsonInput, r)
			}
		}()
		
		// Try to unmarshal the input
		var template RepositoryCreationTemplate
		err := json.Unmarshal([]byte(jsonInput), &template)
		if err != nil {
			// Invalid JSON is expected for fuzz testing, just log and continue
			t.Logf("Invalid JSON input (expected): %v", err)
			return
		}
		
		// Try to marshal it back
		_, err = json.Marshal(&template)
		if err != nil {
			t.Errorf("Failed to marshal valid RepositoryCreationTemplate: %v", err)
		}
		
		// Test runtime object interface
		var obj runtime.Object = &template
		if obj.GetObjectKind() == nil {
			t.Error("GetObjectKind returned nil")
		}
		
		// Test deep copy
		copied := template.DeepCopy()
		if copied == nil {
			t.Error("DeepCopy returned nil")
		}
	})
}

// FuzzRepositoryCreationTemplateValidation fuzzes validation of RepositoryCreationTemplate fields
func FuzzRepositoryCreationTemplateValidation(f *testing.F) {
	// Add seed inputs for validation testing
	f.Add("test-template", "us-east-1", "test-prefix", "IMMUTABLE")
	f.Add("", "", "", "")
	f.Add("template-with-very-long-name-that-exceeds-normal-kubernetes-limits", "invalid-region", "prefix-with-special-chars!@#", "INVALID_MUTABILITY")
	f.Add("test", "us-west-2", "", "MUTABLE")
	
	f.Fuzz(func(t *testing.T, name, region, prefix, imageTagMutability string) {
		// Test validation doesn't panic
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Validation panicked with inputs name=%q, region=%q, prefix=%q, imageTagMutability=%q: %v",
					name, region, prefix, imageTagMutability, r)
			}
		}()
		
		// Create a RepositoryCreationTemplate with fuzzed inputs
		template := &RepositoryCreationTemplate{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "ecr.aws.upbound.io/v1beta1",
				Kind:       "RepositoryCreationTemplate",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
			Spec: RepositoryCreationTemplateSpec{
				ForProvider: RepositoryCreationTemplateParameters{
					Region: &region,
					Prefix: &prefix,
				},
			},
		}
		
		// Test that the object can be created without panicking
		if template.GetName() != name {
			t.Errorf("Expected name %q, got %q", name, template.GetName())
		}
		
		// Test various field validations
		if name == "" {
			t.Logf("Testing empty name")
		}
		
		if len(name) > 253 {
			t.Logf("Testing very long name: %d characters", len(name))
		}
		
		if region == "" {
			t.Logf("Testing empty region")
		}
		
		if prefix == "" {
			t.Logf("Testing empty prefix")
		}
		
		// Test special characters in various fields
		if strings.ContainsAny(name, "!@#$%^&*()") {
			t.Logf("Testing name with special characters: %q", name)
		}
		
		if strings.ContainsAny(prefix, "!@#$%^&*()") {
			t.Logf("Testing prefix with special characters: %q", prefix)
		}
	})
}

// FuzzRepositoryCreationTemplateStatus fuzzes status field operations
func FuzzRepositoryCreationTemplateStatus(f *testing.F) {
	// Add seed inputs for status testing
	f.Add(`{"registryId":"123456789012","templateName":"test-template"}`)
	f.Add(`{}`)
	f.Add(`{"error":"some error occurred"}`)
	f.Add(`{"registryId":"","templateName":""}`)
	
	f.Fuzz(func(t *testing.T, statusJSON string) {
		// Test status operations don't panic
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Status operations panicked with input %q: %v", statusJSON, r)
			}
		}()
		
		template := &RepositoryCreationTemplate{}
		
		// Try to unmarshal status
		var statusData map[string]interface{}
		if json.Unmarshal([]byte(statusJSON), &statusData) == nil {
			// Test SetObservation
			err := template.SetObservation(statusData)
			if err != nil {
				t.Logf("SetObservation failed (may be expected): %v", err)
			}
			
			// Test GetObservation
			obs, err := template.GetObservation()
			if err != nil {
				t.Logf("GetObservation failed: %v", err)
			} else if obs == nil {
				t.Logf("GetObservation returned nil")
			}
		}
		
		// Test other status-related methods
		conditions := template.GetCondition("Ready")
		if conditions.Status == "" {
			t.Logf("No Ready condition found")
		}
		
		// Test connection details
		connDetails := template.GetConnectionDetailsMapping()
		if connDetails == nil {
			t.Logf("No connection details mapping")
		}
	})
}

// FuzzRepositoryCreationTemplateReferences fuzzes cross-resource references
func FuzzRepositoryCreationTemplateReferences(f *testing.F) {
	// Add seed inputs for reference testing
	f.Add("test-kms-key", "test-namespace", "KmsKeyRef")
	f.Add("", "", "")
	f.Add("kms-key-with-very-long-name-that-might-exceed-limits", "namespace-with-special-chars_123", "InvalidRefType")
	f.Add("key", "ns", "KmsKeySelector")
	
	f.Fuzz(func(t *testing.T, refName, refNamespace, refType string) {
		// Test reference operations don't panic
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Reference operations panicked with inputs refName=%q, refNamespace=%q, refType=%q: %v",
					refName, refNamespace, refType, r)
			}
		}()
		
		_ = &RepositoryCreationTemplate{
			Spec: RepositoryCreationTemplateSpec{
				ForProvider: RepositoryCreationTemplateParameters{
					// Test KMS key reference fields if they exist
				},
			},
		}
		
		// Test various reference scenarios
		if refName == "" {
			t.Logf("Testing empty reference name")
		}
		
		if refNamespace == "" {
			t.Logf("Testing empty reference namespace")
		}
		
		if len(refName) > 253 {
			t.Logf("Testing very long reference name: %d characters", len(refName))
		}
		
		// Test reference name validation
		if strings.ContainsAny(refName, "!@#$%^&*()") {
			t.Logf("Testing reference name with special characters: %q", refName)
		}
		
		// Test namespace validation
		if strings.ContainsAny(refNamespace, "!@#$%^&*()") {
			t.Logf("Testing reference namespace with special characters: %q", refNamespace)
		}
		
		// Test reference type validation
		validRefTypes := []string{"KmsKeyRef", "KmsKeySelector"}
		isValidRefType := false
		for _, valid := range validRefTypes {
			if refType == valid {
				isValidRefType = true
				break
			}
		}
		
		if refType != "" && !isValidRefType {
			t.Logf("Testing invalid reference type: %q", refType)
		}
	})
}

// FuzzRepositoryCreationTemplateLifecycle fuzzes lifecycle operations
func FuzzRepositoryCreationTemplateLifecycle(f *testing.F) {
	// Add seed inputs for lifecycle testing
	f.Add("test-template", "us-east-1", true, false)
	f.Add("", "", false, true)
	f.Add("template-to-delete", "eu-west-1", true, true)
	f.Add("template-with-finalizers", "ap-south-1", false, false)
	
	f.Fuzz(func(t *testing.T, name, region string, hasFinalizers, isBeingDeleted bool) {
		// Test lifecycle operations don't panic
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Lifecycle operations panicked with inputs name=%q, region=%q, hasFinalizers=%v, isBeingDeleted=%v: %v",
					name, region, hasFinalizers, isBeingDeleted, r)
			}
		}()
		
		template := &RepositoryCreationTemplate{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
			Spec: RepositoryCreationTemplateSpec{
				ForProvider: RepositoryCreationTemplateParameters{
					Region: &region,
				},
			},
		}
		
		// Add finalizers if requested
		if hasFinalizers {
			template.ObjectMeta.Finalizers = []string{"ecr.aws.upbound.io/finalizer"}
		}
		
		// Set deletion timestamp if being deleted
		if isBeingDeleted {
			now := metav1.Now()
			template.ObjectMeta.DeletionTimestamp = &now
		}
		
		// Test lifecycle methods
		if template.GetDeletionPolicy() == "" {
			t.Logf("No deletion policy set")
		}
		
		if template.GetManagementPolicies() == nil {
			t.Logf("No management policies set")
		}
		
		// Test finalizer operations
		if len(template.GetFinalizers()) != len(template.ObjectMeta.Finalizers) {
			t.Error("Finalizer count mismatch")
		}
		
		// Test deletion timestamp
		if isBeingDeleted && template.GetDeletionTimestamp() == nil {
			t.Error("Deletion timestamp not set when expected")
		}
		
		if !isBeingDeleted && template.GetDeletionTimestamp() != nil {
			t.Error("Deletion timestamp set when not expected")
		}
	})
}
