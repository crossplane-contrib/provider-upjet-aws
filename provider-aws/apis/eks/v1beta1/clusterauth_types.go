/*
Copyright 2022 Upbound Inc.
*/

package v1beta1

import (
	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ClusterAuthParameters struct {
	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region"`

	// ClusterName is the name of the cluster you'd like to fetch Kubeconfig of.
	// Either ClusterName, ClusterNameRef or ClusterNameSelector has to be given.
	// +crossplane:generate:reference:type=Cluster
	ClusterName string `json:"clusterName,omitempty"`

	// Reference to a Cluster to populate clusterName.
	// Either ClusterName, ClusterNameRef or ClusterNameSelector has to be given.
	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty"`

	// Selector for a Cluster to populate clusterName.
	// Either ClusterName, ClusterNameRef or ClusterNameSelector has to be given.
	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty"`

	// RefreshPeriod is how frequently you'd like the token in the published
	// Kubeconfig to be refreshed. The maximum is 10m0s.
	// The default is 10m0s.
	// +kubebuilder:default:="10m0s"
	RefreshPeriod *metav1.Duration `json:"refreshPeriod,omitempty"`
}

type ClusterAuthObservation struct {

	// LastRefreshTime is the time when the token was refreshed.
	LastRefreshTime *metav1.Time `json:"lastRefreshTime,omitempty"`
}

// ClusterAuthSpec defines the desired state of ClusterAuth
type ClusterAuthSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterAuthParameters `json:"forProvider"`
}

// ClusterAuthStatus defines the observed state of ClusterAuth.
type ClusterAuthStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterAuthObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterAuth is used to retrieve Kubeconfig of given EKS cluster.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ClusterAuth struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterAuthSpec   `json:"spec"`
	Status            ClusterAuthStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterAuthList contains a list of ClusterAuths
type ClusterAuthList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterAuth `json:"items"`
}

// Repository type metadata.
var (
	ClusterAuth_Kind             = "ClusterAuth"
	ClusterAuth_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterAuth_Kind}.String()
	ClusterAuth_KindAPIVersion   = ClusterAuth_Kind + "." + CRDGroupVersion.String()
	ClusterAuth_GroupVersionKind = CRDGroupVersion.WithKind(ClusterAuth_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterAuth{}, &ClusterAuthList{})
}
