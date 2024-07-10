// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AuthorizerInitParameters struct {

	// API identifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigatewayv2/v1beta2.API
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a API in apigatewayv2 to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a API in apigatewayv2 to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// Required credentials as an IAM role for API Gateway to invoke the authorizer.
	// Supported only for REQUEST authorizers.
	AuthorizerCredentialsArn *string `json:"authorizerCredentialsArn,omitempty" tf:"authorizer_credentials_arn,omitempty"`

	// Format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers.
	// Valid values: 1.0, 2.0.
	AuthorizerPayloadFormatVersion *string `json:"authorizerPayloadFormatVersion,omitempty" tf:"authorizer_payload_format_version,omitempty"`

	// Time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled.
	// If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Defaults to 300.
	// Supported only for HTTP API Lambda authorizers.
	AuthorizerResultTTLInSeconds *float64 `json:"authorizerResultTtlInSeconds,omitempty" tf:"authorizer_result_ttl_in_seconds,omitempty"`

	// Authorizer type. Valid values: JWT, REQUEST.
	// Specify REQUEST for a Lambda function using incoming request parameters.
	// For HTTP APIs, specify JWT to use JSON Web Tokens.
	AuthorizerType *string `json:"authorizerType,omitempty" tf:"authorizer_type,omitempty"`

	// Authorizer's Uniform Resource Identifier (URI).
	// For REQUEST authorizers this must be a well-formed Lambda function URI, such as the invoke_arn attribute of the aws_lambda_function resource.
	// Supported only for REQUEST authorizers. Must be between 1 and 2048 characters in length.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta2.Function
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common/apis/lambda.FunctionInvokeARN()
	AuthorizerURI *string `json:"authorizerUri,omitempty" tf:"authorizer_uri,omitempty"`

	// Reference to a Function in lambda to populate authorizerUri.
	// +kubebuilder:validation:Optional
	AuthorizerURIRef *v1.Reference `json:"authorizerUriRef,omitempty" tf:"-"`

	// Selector for a Function in lambda to populate authorizerUri.
	// +kubebuilder:validation:Optional
	AuthorizerURISelector *v1.Selector `json:"authorizerUriSelector,omitempty" tf:"-"`

	// Whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
	// Supported only for HTTP APIs.
	EnableSimpleResponses *bool `json:"enableSimpleResponses,omitempty" tf:"enable_simple_responses,omitempty"`

	// Identity sources for which authorization is requested.
	// For REQUEST authorizers the value is a list of one or more mapping expressions of the specified request parameters.
	// For JWT authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
	// +listType=set
	IdentitySources []*string `json:"identitySources,omitempty" tf:"identity_sources,omitempty"`

	// Configuration of a JWT authorizer. Required for the JWT authorizer type.
	// Supported only for HTTP APIs.
	JwtConfiguration *JwtConfigurationInitParameters `json:"jwtConfiguration,omitempty" tf:"jwt_configuration,omitempty"`

	// Name of the authorizer. Must be between 1 and 128 characters in length.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AuthorizerObservation struct {

	// API identifier.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Required credentials as an IAM role for API Gateway to invoke the authorizer.
	// Supported only for REQUEST authorizers.
	AuthorizerCredentialsArn *string `json:"authorizerCredentialsArn,omitempty" tf:"authorizer_credentials_arn,omitempty"`

	// Format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers.
	// Valid values: 1.0, 2.0.
	AuthorizerPayloadFormatVersion *string `json:"authorizerPayloadFormatVersion,omitempty" tf:"authorizer_payload_format_version,omitempty"`

	// Time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled.
	// If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Defaults to 300.
	// Supported only for HTTP API Lambda authorizers.
	AuthorizerResultTTLInSeconds *float64 `json:"authorizerResultTtlInSeconds,omitempty" tf:"authorizer_result_ttl_in_seconds,omitempty"`

	// Authorizer type. Valid values: JWT, REQUEST.
	// Specify REQUEST for a Lambda function using incoming request parameters.
	// For HTTP APIs, specify JWT to use JSON Web Tokens.
	AuthorizerType *string `json:"authorizerType,omitempty" tf:"authorizer_type,omitempty"`

	// Authorizer's Uniform Resource Identifier (URI).
	// For REQUEST authorizers this must be a well-formed Lambda function URI, such as the invoke_arn attribute of the aws_lambda_function resource.
	// Supported only for REQUEST authorizers. Must be between 1 and 2048 characters in length.
	AuthorizerURI *string `json:"authorizerUri,omitempty" tf:"authorizer_uri,omitempty"`

	// Whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
	// Supported only for HTTP APIs.
	EnableSimpleResponses *bool `json:"enableSimpleResponses,omitempty" tf:"enable_simple_responses,omitempty"`

	// Authorizer identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Identity sources for which authorization is requested.
	// For REQUEST authorizers the value is a list of one or more mapping expressions of the specified request parameters.
	// For JWT authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
	// +listType=set
	IdentitySources []*string `json:"identitySources,omitempty" tf:"identity_sources,omitempty"`

	// Configuration of a JWT authorizer. Required for the JWT authorizer type.
	// Supported only for HTTP APIs.
	JwtConfiguration *JwtConfigurationObservation `json:"jwtConfiguration,omitempty" tf:"jwt_configuration,omitempty"`

	// Name of the authorizer. Must be between 1 and 128 characters in length.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AuthorizerParameters struct {

	// API identifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigatewayv2/v1beta2.API
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a API in apigatewayv2 to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a API in apigatewayv2 to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// Required credentials as an IAM role for API Gateway to invoke the authorizer.
	// Supported only for REQUEST authorizers.
	// +kubebuilder:validation:Optional
	AuthorizerCredentialsArn *string `json:"authorizerCredentialsArn,omitempty" tf:"authorizer_credentials_arn,omitempty"`

	// Format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers.
	// Valid values: 1.0, 2.0.
	// +kubebuilder:validation:Optional
	AuthorizerPayloadFormatVersion *string `json:"authorizerPayloadFormatVersion,omitempty" tf:"authorizer_payload_format_version,omitempty"`

	// Time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled.
	// If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Defaults to 300.
	// Supported only for HTTP API Lambda authorizers.
	// +kubebuilder:validation:Optional
	AuthorizerResultTTLInSeconds *float64 `json:"authorizerResultTtlInSeconds,omitempty" tf:"authorizer_result_ttl_in_seconds,omitempty"`

	// Authorizer type. Valid values: JWT, REQUEST.
	// Specify REQUEST for a Lambda function using incoming request parameters.
	// For HTTP APIs, specify JWT to use JSON Web Tokens.
	// +kubebuilder:validation:Optional
	AuthorizerType *string `json:"authorizerType,omitempty" tf:"authorizer_type,omitempty"`

	// Authorizer's Uniform Resource Identifier (URI).
	// For REQUEST authorizers this must be a well-formed Lambda function URI, such as the invoke_arn attribute of the aws_lambda_function resource.
	// Supported only for REQUEST authorizers. Must be between 1 and 2048 characters in length.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta2.Function
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common/apis/lambda.FunctionInvokeARN()
	// +kubebuilder:validation:Optional
	AuthorizerURI *string `json:"authorizerUri,omitempty" tf:"authorizer_uri,omitempty"`

	// Reference to a Function in lambda to populate authorizerUri.
	// +kubebuilder:validation:Optional
	AuthorizerURIRef *v1.Reference `json:"authorizerUriRef,omitempty" tf:"-"`

	// Selector for a Function in lambda to populate authorizerUri.
	// +kubebuilder:validation:Optional
	AuthorizerURISelector *v1.Selector `json:"authorizerUriSelector,omitempty" tf:"-"`

	// Whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
	// Supported only for HTTP APIs.
	// +kubebuilder:validation:Optional
	EnableSimpleResponses *bool `json:"enableSimpleResponses,omitempty" tf:"enable_simple_responses,omitempty"`

	// Identity sources for which authorization is requested.
	// For REQUEST authorizers the value is a list of one or more mapping expressions of the specified request parameters.
	// For JWT authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentitySources []*string `json:"identitySources,omitempty" tf:"identity_sources,omitempty"`

	// Configuration of a JWT authorizer. Required for the JWT authorizer type.
	// Supported only for HTTP APIs.
	// +kubebuilder:validation:Optional
	JwtConfiguration *JwtConfigurationParameters `json:"jwtConfiguration,omitempty" tf:"jwt_configuration,omitempty"`

	// Name of the authorizer. Must be between 1 and 128 characters in length.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type JwtConfigurationInitParameters struct {

	// List of the intended recipients of the JWT. A valid JWT must provide an aud that matches at least one entry in this list.
	// +listType=set
	Audience []*string `json:"audience,omitempty" tf:"audience,omitempty"`

	// Base domain of the identity provider that issues JSON Web Tokens, such as the endpoint attribute of the aws_cognito_user_pool resource.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`
}

type JwtConfigurationObservation struct {

	// List of the intended recipients of the JWT. A valid JWT must provide an aud that matches at least one entry in this list.
	// +listType=set
	Audience []*string `json:"audience,omitempty" tf:"audience,omitempty"`

	// Base domain of the identity provider that issues JSON Web Tokens, such as the endpoint attribute of the aws_cognito_user_pool resource.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`
}

type JwtConfigurationParameters struct {

	// List of the intended recipients of the JWT. A valid JWT must provide an aud that matches at least one entry in this list.
	// +kubebuilder:validation:Optional
	// +listType=set
	Audience []*string `json:"audience,omitempty" tf:"audience,omitempty"`

	// Base domain of the identity provider that issues JSON Web Tokens, such as the endpoint attribute of the aws_cognito_user_pool resource.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`
}

// AuthorizerSpec defines the desired state of Authorizer
type AuthorizerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthorizerParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AuthorizerInitParameters `json:"initProvider,omitempty"`
}

// AuthorizerStatus defines the observed state of Authorizer.
type AuthorizerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthorizerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Authorizer is the Schema for the Authorizers API. Manages an Amazon API Gateway Version 2 authorizer.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Authorizer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authorizerType) || (has(self.initProvider) && has(self.initProvider.authorizerType))",message="spec.forProvider.authorizerType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   AuthorizerSpec   `json:"spec"`
	Status AuthorizerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthorizerList contains a list of Authorizers
type AuthorizerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Authorizer `json:"items"`
}

// Repository type metadata.
var (
	Authorizer_Kind             = "Authorizer"
	Authorizer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Authorizer_Kind}.String()
	Authorizer_KindAPIVersion   = Authorizer_Kind + "." + CRDGroupVersion.String()
	Authorizer_GroupVersionKind = CRDGroupVersion.WithKind(Authorizer_Kind)
)

func init() {
	SchemeBuilder.Register(&Authorizer{}, &AuthorizerList{})
}