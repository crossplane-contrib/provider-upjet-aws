// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type WorkerConfigurationInitParameters_2 struct {

	// A summary description of the worker configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Contents of connect-distributed.properties file. The value can be either base64 encoded or in raw format.
	PropertiesFileContent *string `json:"propertiesFileContent,omitempty" tf:"properties_file_content,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type WorkerConfigurationObservation_2 struct {

	// the Amazon Resource Name (ARN) of the worker configuration.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A summary description of the worker configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// an ID of the latest successfully created revision of the worker configuration.
	LatestRevision *float64 `json:"latestRevision,omitempty" tf:"latest_revision,omitempty"`

	// The name of the worker configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Contents of connect-distributed.properties file. The value can be either base64 encoded or in raw format.
	PropertiesFileContent *string `json:"propertiesFileContent,omitempty" tf:"properties_file_content,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type WorkerConfigurationParameters_2 struct {

	// A summary description of the worker configuration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the worker configuration.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Contents of connect-distributed.properties file. The value can be either base64 encoded or in raw format.
	// +kubebuilder:validation:Optional
	PropertiesFileContent *string `json:"propertiesFileContent,omitempty" tf:"properties_file_content,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WorkerConfigurationSpec defines the desired state of WorkerConfiguration
type WorkerConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkerConfigurationParameters_2 `json:"forProvider"`
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
	InitProvider WorkerConfigurationInitParameters_2 `json:"initProvider,omitempty"`
}

// WorkerConfigurationStatus defines the observed state of WorkerConfiguration.
type WorkerConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkerConfigurationObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// WorkerConfiguration is the Schema for the WorkerConfigurations API. Provides an Amazon MSK Connect worker configuration resource. This resource is create-only, and requires a unique "name" parameter. AWS does not currently provide update or delete APIs.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type WorkerConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.propertiesFileContent) || (has(self.initProvider) && has(self.initProvider.propertiesFileContent))",message="spec.forProvider.propertiesFileContent is a required parameter"
	Spec   WorkerConfigurationSpec   `json:"spec"`
	Status WorkerConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkerConfigurationList contains a list of WorkerConfigurations
type WorkerConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkerConfiguration `json:"items"`
}

// Repository type metadata.
var (
	WorkerConfiguration_Kind             = "WorkerConfiguration"
	WorkerConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WorkerConfiguration_Kind}.String()
	WorkerConfiguration_KindAPIVersion   = WorkerConfiguration_Kind + "." + CRDGroupVersion.String()
	WorkerConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(WorkerConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&WorkerConfiguration{}, &WorkerConfigurationList{})
}
