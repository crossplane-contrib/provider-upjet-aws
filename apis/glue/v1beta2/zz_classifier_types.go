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

type ClassifierInitParameters struct {

	// A classifier for CSV content. Defined below.
	CsvClassifier *CsvClassifierInitParameters `json:"csvClassifier,omitempty" tf:"csv_classifier,omitempty"`

	// –  A classifier that uses grok patterns. Defined below.
	GrokClassifier *GrokClassifierInitParameters `json:"grokClassifier,omitempty" tf:"grok_classifier,omitempty"`

	// –  A classifier for JSON content. Defined below.
	JSONClassifier *JSONClassifierInitParameters `json:"jsonClassifier,omitempty" tf:"json_classifier,omitempty"`

	// –  A classifier for XML content. Defined below.
	XMLClassifier *XMLClassifierInitParameters `json:"xmlClassifier,omitempty" tf:"xml_classifier,omitempty"`
}

type ClassifierObservation struct {

	// A classifier for CSV content. Defined below.
	CsvClassifier *CsvClassifierObservation `json:"csvClassifier,omitempty" tf:"csv_classifier,omitempty"`

	// –  A classifier that uses grok patterns. Defined below.
	GrokClassifier *GrokClassifierObservation `json:"grokClassifier,omitempty" tf:"grok_classifier,omitempty"`

	// Name of the classifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// –  A classifier for JSON content. Defined below.
	JSONClassifier *JSONClassifierObservation `json:"jsonClassifier,omitempty" tf:"json_classifier,omitempty"`

	// –  A classifier for XML content. Defined below.
	XMLClassifier *XMLClassifierObservation `json:"xmlClassifier,omitempty" tf:"xml_classifier,omitempty"`
}

type ClassifierParameters struct {

	// A classifier for CSV content. Defined below.
	// +kubebuilder:validation:Optional
	CsvClassifier *CsvClassifierParameters `json:"csvClassifier,omitempty" tf:"csv_classifier,omitempty"`

	// –  A classifier that uses grok patterns. Defined below.
	// +kubebuilder:validation:Optional
	GrokClassifier *GrokClassifierParameters `json:"grokClassifier,omitempty" tf:"grok_classifier,omitempty"`

	// –  A classifier for JSON content. Defined below.
	// +kubebuilder:validation:Optional
	JSONClassifier *JSONClassifierParameters `json:"jsonClassifier,omitempty" tf:"json_classifier,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// –  A classifier for XML content. Defined below.
	// +kubebuilder:validation:Optional
	XMLClassifier *XMLClassifierParameters `json:"xmlClassifier,omitempty" tf:"xml_classifier,omitempty"`
}

type CsvClassifierInitParameters struct {

	// Enables the processing of files that contain only one column.
	AllowSingleColumn *bool `json:"allowSingleColumn,omitempty" tf:"allow_single_column,omitempty"`

	// Indicates whether the CSV file contains a header. This can be one of "ABSENT", "PRESENT", or "UNKNOWN".
	ContainsHeader *string `json:"containsHeader,omitempty" tf:"contains_header,omitempty"`

	// Enables the custom datatype to be configured.
	CustomDatatypeConfigured *bool `json:"customDatatypeConfigured,omitempty" tf:"custom_datatype_configured,omitempty"`

	// A list of supported custom datatypes. Valid values are BINARY, BOOLEAN, DATE, DECIMAL, DOUBLE, FLOAT, INT, LONG, SHORT, STRING, TIMESTAMP.
	CustomDatatypes []*string `json:"customDatatypes,omitempty" tf:"custom_datatypes,omitempty"`

	// The delimiter used in the CSV to separate columns.
	Delimiter *string `json:"delimiter,omitempty" tf:"delimiter,omitempty"`

	// Specifies whether to trim column values.
	DisableValueTrimming *bool `json:"disableValueTrimming,omitempty" tf:"disable_value_trimming,omitempty"`

	// A list of strings representing column names.
	Header []*string `json:"header,omitempty" tf:"header,omitempty"`

	// A custom symbol to denote what combines content into a single column value. It must be different from the column delimiter.
	QuoteSymbol *string `json:"quoteSymbol,omitempty" tf:"quote_symbol,omitempty"`

	// –  The SerDe for processing CSV. Valid values are OpenCSVSerDe, LazySimpleSerDe, None.
	Serde *string `json:"serde,omitempty" tf:"serde,omitempty"`
}

type CsvClassifierObservation struct {

	// Enables the processing of files that contain only one column.
	AllowSingleColumn *bool `json:"allowSingleColumn,omitempty" tf:"allow_single_column,omitempty"`

	// Indicates whether the CSV file contains a header. This can be one of "ABSENT", "PRESENT", or "UNKNOWN".
	ContainsHeader *string `json:"containsHeader,omitempty" tf:"contains_header,omitempty"`

	// Enables the custom datatype to be configured.
	CustomDatatypeConfigured *bool `json:"customDatatypeConfigured,omitempty" tf:"custom_datatype_configured,omitempty"`

	// A list of supported custom datatypes. Valid values are BINARY, BOOLEAN, DATE, DECIMAL, DOUBLE, FLOAT, INT, LONG, SHORT, STRING, TIMESTAMP.
	CustomDatatypes []*string `json:"customDatatypes,omitempty" tf:"custom_datatypes,omitempty"`

	// The delimiter used in the CSV to separate columns.
	Delimiter *string `json:"delimiter,omitempty" tf:"delimiter,omitempty"`

	// Specifies whether to trim column values.
	DisableValueTrimming *bool `json:"disableValueTrimming,omitempty" tf:"disable_value_trimming,omitempty"`

	// A list of strings representing column names.
	Header []*string `json:"header,omitempty" tf:"header,omitempty"`

	// A custom symbol to denote what combines content into a single column value. It must be different from the column delimiter.
	QuoteSymbol *string `json:"quoteSymbol,omitempty" tf:"quote_symbol,omitempty"`

	// –  The SerDe for processing CSV. Valid values are OpenCSVSerDe, LazySimpleSerDe, None.
	Serde *string `json:"serde,omitempty" tf:"serde,omitempty"`
}

type CsvClassifierParameters struct {

	// Enables the processing of files that contain only one column.
	// +kubebuilder:validation:Optional
	AllowSingleColumn *bool `json:"allowSingleColumn,omitempty" tf:"allow_single_column,omitempty"`

	// Indicates whether the CSV file contains a header. This can be one of "ABSENT", "PRESENT", or "UNKNOWN".
	// +kubebuilder:validation:Optional
	ContainsHeader *string `json:"containsHeader,omitempty" tf:"contains_header,omitempty"`

	// Enables the custom datatype to be configured.
	// +kubebuilder:validation:Optional
	CustomDatatypeConfigured *bool `json:"customDatatypeConfigured,omitempty" tf:"custom_datatype_configured,omitempty"`

	// A list of supported custom datatypes. Valid values are BINARY, BOOLEAN, DATE, DECIMAL, DOUBLE, FLOAT, INT, LONG, SHORT, STRING, TIMESTAMP.
	// +kubebuilder:validation:Optional
	CustomDatatypes []*string `json:"customDatatypes,omitempty" tf:"custom_datatypes,omitempty"`

	// The delimiter used in the CSV to separate columns.
	// +kubebuilder:validation:Optional
	Delimiter *string `json:"delimiter,omitempty" tf:"delimiter,omitempty"`

	// Specifies whether to trim column values.
	// +kubebuilder:validation:Optional
	DisableValueTrimming *bool `json:"disableValueTrimming,omitempty" tf:"disable_value_trimming,omitempty"`

	// A list of strings representing column names.
	// +kubebuilder:validation:Optional
	Header []*string `json:"header,omitempty" tf:"header,omitempty"`

	// A custom symbol to denote what combines content into a single column value. It must be different from the column delimiter.
	// +kubebuilder:validation:Optional
	QuoteSymbol *string `json:"quoteSymbol,omitempty" tf:"quote_symbol,omitempty"`

	// –  The SerDe for processing CSV. Valid values are OpenCSVSerDe, LazySimpleSerDe, None.
	// +kubebuilder:validation:Optional
	Serde *string `json:"serde,omitempty" tf:"serde,omitempty"`
}

type GrokClassifierInitParameters struct {

	// An identifier of the data format that the classifier matches, such as Twitter, JSON, Omniture logs, Amazon CloudWatch Logs, and so on.
	Classification *string `json:"classification,omitempty" tf:"classification,omitempty"`

	// Custom grok patterns used by this classifier.
	CustomPatterns *string `json:"customPatterns,omitempty" tf:"custom_patterns,omitempty"`

	// The grok pattern used by this classifier.
	GrokPattern *string `json:"grokPattern,omitempty" tf:"grok_pattern,omitempty"`
}

type GrokClassifierObservation struct {

	// An identifier of the data format that the classifier matches, such as Twitter, JSON, Omniture logs, Amazon CloudWatch Logs, and so on.
	Classification *string `json:"classification,omitempty" tf:"classification,omitempty"`

	// Custom grok patterns used by this classifier.
	CustomPatterns *string `json:"customPatterns,omitempty" tf:"custom_patterns,omitempty"`

	// The grok pattern used by this classifier.
	GrokPattern *string `json:"grokPattern,omitempty" tf:"grok_pattern,omitempty"`
}

type GrokClassifierParameters struct {

	// An identifier of the data format that the classifier matches, such as Twitter, JSON, Omniture logs, Amazon CloudWatch Logs, and so on.
	// +kubebuilder:validation:Optional
	Classification *string `json:"classification" tf:"classification,omitempty"`

	// Custom grok patterns used by this classifier.
	// +kubebuilder:validation:Optional
	CustomPatterns *string `json:"customPatterns,omitempty" tf:"custom_patterns,omitempty"`

	// The grok pattern used by this classifier.
	// +kubebuilder:validation:Optional
	GrokPattern *string `json:"grokPattern" tf:"grok_pattern,omitempty"`
}

type JSONClassifierInitParameters struct {

	// A JsonPath string defining the JSON data for the classifier to classify. AWS Glue supports a subset of JsonPath, as described in Writing JsonPath Custom Classifiers.
	JSONPath *string `json:"jsonPath,omitempty" tf:"json_path,omitempty"`
}

type JSONClassifierObservation struct {

	// A JsonPath string defining the JSON data for the classifier to classify. AWS Glue supports a subset of JsonPath, as described in Writing JsonPath Custom Classifiers.
	JSONPath *string `json:"jsonPath,omitempty" tf:"json_path,omitempty"`
}

type JSONClassifierParameters struct {

	// A JsonPath string defining the JSON data for the classifier to classify. AWS Glue supports a subset of JsonPath, as described in Writing JsonPath Custom Classifiers.
	// +kubebuilder:validation:Optional
	JSONPath *string `json:"jsonPath" tf:"json_path,omitempty"`
}

type XMLClassifierInitParameters struct {

	// An identifier of the data format that the classifier matches.
	Classification *string `json:"classification,omitempty" tf:"classification,omitempty"`

	// The XML tag designating the element that contains each record in an XML document being parsed. Note that this cannot identify a self-closing element (closed by />). An empty row element that contains only attributes can be parsed as long as it ends with a closing tag (for example, <row item_a="A" item_b="B"></row> is okay, but <row item_a="A" item_b="B" /> is not).
	RowTag *string `json:"rowTag,omitempty" tf:"row_tag,omitempty"`
}

type XMLClassifierObservation struct {

	// An identifier of the data format that the classifier matches.
	Classification *string `json:"classification,omitempty" tf:"classification,omitempty"`

	// The XML tag designating the element that contains each record in an XML document being parsed. Note that this cannot identify a self-closing element (closed by />). An empty row element that contains only attributes can be parsed as long as it ends with a closing tag (for example, <row item_a="A" item_b="B"></row> is okay, but <row item_a="A" item_b="B" /> is not).
	RowTag *string `json:"rowTag,omitempty" tf:"row_tag,omitempty"`
}

type XMLClassifierParameters struct {

	// An identifier of the data format that the classifier matches.
	// +kubebuilder:validation:Optional
	Classification *string `json:"classification" tf:"classification,omitempty"`

	// The XML tag designating the element that contains each record in an XML document being parsed. Note that this cannot identify a self-closing element (closed by />). An empty row element that contains only attributes can be parsed as long as it ends with a closing tag (for example, <row item_a="A" item_b="B"></row> is okay, but <row item_a="A" item_b="B" /> is not).
	// +kubebuilder:validation:Optional
	RowTag *string `json:"rowTag" tf:"row_tag,omitempty"`
}

// ClassifierSpec defines the desired state of Classifier
type ClassifierSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClassifierParameters `json:"forProvider"`
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
	InitProvider ClassifierInitParameters `json:"initProvider,omitempty"`
}

// ClassifierStatus defines the observed state of Classifier.
type ClassifierStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClassifierObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Classifier is the Schema for the Classifiers API. Provides an Glue Classifier resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Classifier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClassifierSpec   `json:"spec"`
	Status            ClassifierStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClassifierList contains a list of Classifiers
type ClassifierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Classifier `json:"items"`
}

// Repository type metadata.
var (
	Classifier_Kind             = "Classifier"
	Classifier_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Classifier_Kind}.String()
	Classifier_KindAPIVersion   = Classifier_Kind + "." + CRDGroupVersion.String()
	Classifier_GroupVersionKind = CRDGroupVersion.WithKind(Classifier_Kind)
)

func init() {
	SchemeBuilder.Register(&Classifier{}, &ClassifierList{})
}
