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

type DefaultTagsInitParameters struct {

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DefaultTagsObservation struct {

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DefaultTagsParameters struct {

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ObjectInitParameters struct {

	// Canned ACL to apply. Valid values are private, public-read, public-read-write, aws-exec-read, authenticated-read, bucket-owner-read, and bucket-owner-full-control.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// Name of the bucket to put the file in. Alternatively, an S3 access point ARN can be specified.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta2.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Whether or not to use Amazon S3 Bucket Keys for SSE-KMS.
	BucketKeyEnabled *bool `json:"bucketKeyEnabled,omitempty" tf:"bucket_key_enabled,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Caching behavior along the request/reply chain Read w3c cache_control for further details.
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// Indicates the algorithm used to create the checksum for the object. If a value is specified and the object is encrypted with KMS, you must have permission to use the kms:Decrypt action. Valid values: CRC32, CRC32C, SHA1, SHA256.
	ChecksumAlgorithm *string `json:"checksumAlgorithm,omitempty" tf:"checksum_algorithm,omitempty"`

	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the gzipbase64 function with small text strings. For larger objects, use source to stream the content from a disk file.
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// Presentational information for the object. Read w3c content_disposition for further information.
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// Content encodings that have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read w3c content encoding for further information.
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// Language the content is in e.g., en-US or en-GB.
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// Standard MIME type describing the format of the object data, e.g., application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Triggers updates when the value changes.11.11.11 or earlier). This attribute is not compatible with KMS encryption, kms_key_id or server_side_encryption = "aws:kms", also if an object is larger than 16 MB, the AWS Management Console will upload or copy that object as a Multipart Upload, and therefore the ETag will not be an MD5 digest (see source_hash instead).
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// Whether to allow the object to be deleted by removing any legal hold on any object version. Default is false. This value should be set to true only if the bucket has S3 object lock enabled.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// ARN of the KMS Key to use for object encryption. If the S3 Bucket has server-side encryption enabled, that value will automatically be used. If referencing the aws_kms_key resource, use the arn attribute. If referencing the aws_kms_alias data source or resource, use the target_key_arn attribute.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Name of the object once it is in the bucket.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Map of keys/values to provision metadata (will be automatically prefixed by x-amz-meta-, note that only lowercase label are currently supported by the AWS Go API).
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Legal hold status that you want to apply to the specified object. Valid values are ON and OFF.
	ObjectLockLegalHoldStatus *string `json:"objectLockLegalHoldStatus,omitempty" tf:"object_lock_legal_hold_status,omitempty"`

	// Object lock retention mode that you want to apply to this object. Valid values are GOVERNANCE and COMPLIANCE.
	ObjectLockMode *string `json:"objectLockMode,omitempty" tf:"object_lock_mode,omitempty"`

	// Date and time, in RFC3339 format, when this object's object lock will expire.
	ObjectLockRetainUntilDate *string `json:"objectLockRetainUntilDate,omitempty" tf:"object_lock_retain_until_date,omitempty"`

	// Override provider-level configuration options. See Override Provider below for more details.
	OverrideProvider *OverrideProviderInitParameters `json:"overrideProvider,omitempty" tf:"override_provider,omitempty"`

	// Server-side encryption of the object in S3. Valid values are "AES256" and "aws:kms".
	ServerSideEncryption *string `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

	// Path to a file that will be read and uploaded as raw bytes for the object content.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Triggers updates like etag but useful to address etag encryption limitations.11.12 or later). (The value is only stored in state and not saved by AWS.)
	SourceHash *string `json:"sourceHash,omitempty" tf:"source_hash,omitempty"`

	// Storage Class for the object. Defaults to "STANDARD".
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Target URL for website redirect.
	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect,omitempty"`
}

type ObjectObservation struct {

	// Canned ACL to apply. Valid values are private, public-read, public-read-write, aws-exec-read, authenticated-read, bucket-owner-read, and bucket-owner-full-control.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// ARN of the object.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Name of the bucket to put the file in. Alternatively, an S3 access point ARN can be specified.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Whether or not to use Amazon S3 Bucket Keys for SSE-KMS.
	BucketKeyEnabled *bool `json:"bucketKeyEnabled,omitempty" tf:"bucket_key_enabled,omitempty"`

	// Caching behavior along the request/reply chain Read w3c cache_control for further details.
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// Indicates the algorithm used to create the checksum for the object. If a value is specified and the object is encrypted with KMS, you must have permission to use the kms:Decrypt action. Valid values: CRC32, CRC32C, SHA1, SHA256.
	ChecksumAlgorithm *string `json:"checksumAlgorithm,omitempty" tf:"checksum_algorithm,omitempty"`

	// The base64-encoded, 32-bit CRC32 checksum of the object.
	ChecksumCrc32 *string `json:"checksumCrc32,omitempty" tf:"checksum_crc32,omitempty"`

	// The base64-encoded, 32-bit CRC32C checksum of the object.
	ChecksumCrc32C *string `json:"checksumCrc32C,omitempty" tf:"checksum_crc32c,omitempty"`

	// The base64-encoded, 160-bit SHA-1 digest of the object.
	ChecksumSha1 *string `json:"checksumSha1,omitempty" tf:"checksum_sha1,omitempty"`

	// The base64-encoded, 256-bit SHA-256 digest of the object.
	ChecksumSha256 *string `json:"checksumSha256,omitempty" tf:"checksum_sha256,omitempty"`

	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the gzipbase64 function with small text strings. For larger objects, use source to stream the content from a disk file.
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// Presentational information for the object. Read w3c content_disposition for further information.
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// Content encodings that have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read w3c content encoding for further information.
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// Language the content is in e.g., en-US or en-GB.
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// Standard MIME type describing the format of the object data, e.g., application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Triggers updates when the value changes.11.11.11 or earlier). This attribute is not compatible with KMS encryption, kms_key_id or server_side_encryption = "aws:kms", also if an object is larger than 16 MB, the AWS Management Console will upload or copy that object as a Multipart Upload, and therefore the ETag will not be an MD5 digest (see source_hash instead).
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// Whether to allow the object to be deleted by removing any legal hold on any object version. Default is false. This value should be set to true only if the bucket has S3 object lock enabled.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN of the KMS Key to use for object encryption. If the S3 Bucket has server-side encryption enabled, that value will automatically be used. If referencing the aws_kms_key resource, use the arn attribute. If referencing the aws_kms_alias data source or resource, use the target_key_arn attribute.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Name of the object once it is in the bucket.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Map of keys/values to provision metadata (will be automatically prefixed by x-amz-meta-, note that only lowercase label are currently supported by the AWS Go API).
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Legal hold status that you want to apply to the specified object. Valid values are ON and OFF.
	ObjectLockLegalHoldStatus *string `json:"objectLockLegalHoldStatus,omitempty" tf:"object_lock_legal_hold_status,omitempty"`

	// Object lock retention mode that you want to apply to this object. Valid values are GOVERNANCE and COMPLIANCE.
	ObjectLockMode *string `json:"objectLockMode,omitempty" tf:"object_lock_mode,omitempty"`

	// Date and time, in RFC3339 format, when this object's object lock will expire.
	ObjectLockRetainUntilDate *string `json:"objectLockRetainUntilDate,omitempty" tf:"object_lock_retain_until_date,omitempty"`

	// Override provider-level configuration options. See Override Provider below for more details.
	OverrideProvider *OverrideProviderObservation `json:"overrideProvider,omitempty" tf:"override_provider,omitempty"`

	// Server-side encryption of the object in S3. Valid values are "AES256" and "aws:kms".
	ServerSideEncryption *string `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

	// Path to a file that will be read and uploaded as raw bytes for the object content.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Triggers updates like etag but useful to address etag encryption limitations.11.12 or later). (The value is only stored in state and not saved by AWS.)
	SourceHash *string `json:"sourceHash,omitempty" tf:"source_hash,omitempty"`

	// Storage Class for the object. Defaults to "STANDARD".
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Unique version ID value for the object, if bucket versioning is enabled.
	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`

	// Target URL for website redirect.
	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect,omitempty"`
}

type ObjectParameters struct {

	// Canned ACL to apply. Valid values are private, public-read, public-read-write, aws-exec-read, authenticated-read, bucket-owner-read, and bucket-owner-full-control.
	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// Name of the bucket to put the file in. Alternatively, an S3 access point ARN can be specified.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta2.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Whether or not to use Amazon S3 Bucket Keys for SSE-KMS.
	// +kubebuilder:validation:Optional
	BucketKeyEnabled *bool `json:"bucketKeyEnabled,omitempty" tf:"bucket_key_enabled,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Caching behavior along the request/reply chain Read w3c cache_control for further details.
	// +kubebuilder:validation:Optional
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// Indicates the algorithm used to create the checksum for the object. If a value is specified and the object is encrypted with KMS, you must have permission to use the kms:Decrypt action. Valid values: CRC32, CRC32C, SHA1, SHA256.
	// +kubebuilder:validation:Optional
	ChecksumAlgorithm *string `json:"checksumAlgorithm,omitempty" tf:"checksum_algorithm,omitempty"`

	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the gzipbase64 function with small text strings. For larger objects, use source to stream the content from a disk file.
	// +kubebuilder:validation:Optional
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// Presentational information for the object. Read w3c content_disposition for further information.
	// +kubebuilder:validation:Optional
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// Content encodings that have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read w3c content encoding for further information.
	// +kubebuilder:validation:Optional
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// Language the content is in e.g., en-US or en-GB.
	// +kubebuilder:validation:Optional
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// Standard MIME type describing the format of the object data, e.g., application/octet-stream. All Valid MIME Types are valid for this input.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Triggers updates when the value changes.11.11.11 or earlier). This attribute is not compatible with KMS encryption, kms_key_id or server_side_encryption = "aws:kms", also if an object is larger than 16 MB, the AWS Management Console will upload or copy that object as a Multipart Upload, and therefore the ETag will not be an MD5 digest (see source_hash instead).
	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// Whether to allow the object to be deleted by removing any legal hold on any object version. Default is false. This value should be set to true only if the bucket has S3 object lock enabled.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// ARN of the KMS Key to use for object encryption. If the S3 Bucket has server-side encryption enabled, that value will automatically be used. If referencing the aws_kms_key resource, use the arn attribute. If referencing the aws_kms_alias data source or resource, use the target_key_arn attribute.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Name of the object once it is in the bucket.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Map of keys/values to provision metadata (will be automatically prefixed by x-amz-meta-, note that only lowercase label are currently supported by the AWS Go API).
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Legal hold status that you want to apply to the specified object. Valid values are ON and OFF.
	// +kubebuilder:validation:Optional
	ObjectLockLegalHoldStatus *string `json:"objectLockLegalHoldStatus,omitempty" tf:"object_lock_legal_hold_status,omitempty"`

	// Object lock retention mode that you want to apply to this object. Valid values are GOVERNANCE and COMPLIANCE.
	// +kubebuilder:validation:Optional
	ObjectLockMode *string `json:"objectLockMode,omitempty" tf:"object_lock_mode,omitempty"`

	// Date and time, in RFC3339 format, when this object's object lock will expire.
	// +kubebuilder:validation:Optional
	ObjectLockRetainUntilDate *string `json:"objectLockRetainUntilDate,omitempty" tf:"object_lock_retain_until_date,omitempty"`

	// Override provider-level configuration options. See Override Provider below for more details.
	// +kubebuilder:validation:Optional
	OverrideProvider *OverrideProviderParameters `json:"overrideProvider,omitempty" tf:"override_provider,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Server-side encryption of the object in S3. Valid values are "AES256" and "aws:kms".
	// +kubebuilder:validation:Optional
	ServerSideEncryption *string `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

	// Path to a file that will be read and uploaded as raw bytes for the object content.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Triggers updates like etag but useful to address etag encryption limitations.11.12 or later). (The value is only stored in state and not saved by AWS.)
	// +kubebuilder:validation:Optional
	SourceHash *string `json:"sourceHash,omitempty" tf:"source_hash,omitempty"`

	// Storage Class for the object. Defaults to "STANDARD".
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Target URL for website redirect.
	// +kubebuilder:validation:Optional
	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect,omitempty"`
}

type OverrideProviderInitParameters struct {

	// Override the provider default_tags configuration block.
	DefaultTags *DefaultTagsInitParameters `json:"defaultTags,omitempty" tf:"default_tags,omitempty"`
}

type OverrideProviderObservation struct {

	// Override the provider default_tags configuration block.
	DefaultTags *DefaultTagsObservation `json:"defaultTags,omitempty" tf:"default_tags,omitempty"`
}

type OverrideProviderParameters struct {

	// Override the provider default_tags configuration block.
	// +kubebuilder:validation:Optional
	DefaultTags *DefaultTagsParameters `json:"defaultTags,omitempty" tf:"default_tags,omitempty"`
}

// ObjectSpec defines the desired state of Object
type ObjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ObjectParameters `json:"forProvider"`
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
	InitProvider ObjectInitParameters `json:"initProvider,omitempty"`
}

// ObjectStatus defines the observed state of Object.
type ObjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ObjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Object is the Schema for the Objects API. Provides an S3 object resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Object struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.key) || (has(self.initProvider) && has(self.initProvider.key))",message="spec.forProvider.key is a required parameter"
	Spec   ObjectSpec   `json:"spec"`
	Status ObjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectList contains a list of Objects
type ObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Object `json:"items"`
}

// Repository type metadata.
var (
	Object_Kind             = "Object"
	Object_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Object_Kind}.String()
	Object_KindAPIVersion   = Object_Kind + "." + CRDGroupVersion.String()
	Object_GroupVersionKind = CRDGroupVersion.WithKind(Object_Kind)
)

func init() {
	SchemeBuilder.Register(&Object{}, &ObjectList{})
}