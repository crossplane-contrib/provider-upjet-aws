//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapacitySpecificationInitParameters) DeepCopyInto(out *CapacitySpecificationInitParameters) {
	*out = *in
	if in.ReadCapacityUnits != nil {
		in, out := &in.ReadCapacityUnits, &out.ReadCapacityUnits
		*out = new(float64)
		**out = **in
	}
	if in.ThroughputMode != nil {
		in, out := &in.ThroughputMode, &out.ThroughputMode
		*out = new(string)
		**out = **in
	}
	if in.WriteCapacityUnits != nil {
		in, out := &in.WriteCapacityUnits, &out.WriteCapacityUnits
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapacitySpecificationInitParameters.
func (in *CapacitySpecificationInitParameters) DeepCopy() *CapacitySpecificationInitParameters {
	if in == nil {
		return nil
	}
	out := new(CapacitySpecificationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapacitySpecificationObservation) DeepCopyInto(out *CapacitySpecificationObservation) {
	*out = *in
	if in.ReadCapacityUnits != nil {
		in, out := &in.ReadCapacityUnits, &out.ReadCapacityUnits
		*out = new(float64)
		**out = **in
	}
	if in.ThroughputMode != nil {
		in, out := &in.ThroughputMode, &out.ThroughputMode
		*out = new(string)
		**out = **in
	}
	if in.WriteCapacityUnits != nil {
		in, out := &in.WriteCapacityUnits, &out.WriteCapacityUnits
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapacitySpecificationObservation.
func (in *CapacitySpecificationObservation) DeepCopy() *CapacitySpecificationObservation {
	if in == nil {
		return nil
	}
	out := new(CapacitySpecificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapacitySpecificationParameters) DeepCopyInto(out *CapacitySpecificationParameters) {
	*out = *in
	if in.ReadCapacityUnits != nil {
		in, out := &in.ReadCapacityUnits, &out.ReadCapacityUnits
		*out = new(float64)
		**out = **in
	}
	if in.ThroughputMode != nil {
		in, out := &in.ThroughputMode, &out.ThroughputMode
		*out = new(string)
		**out = **in
	}
	if in.WriteCapacityUnits != nil {
		in, out := &in.WriteCapacityUnits, &out.WriteCapacityUnits
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapacitySpecificationParameters.
func (in *CapacitySpecificationParameters) DeepCopy() *CapacitySpecificationParameters {
	if in == nil {
		return nil
	}
	out := new(CapacitySpecificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientSideTimestampsInitParameters) DeepCopyInto(out *ClientSideTimestampsInitParameters) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientSideTimestampsInitParameters.
func (in *ClientSideTimestampsInitParameters) DeepCopy() *ClientSideTimestampsInitParameters {
	if in == nil {
		return nil
	}
	out := new(ClientSideTimestampsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientSideTimestampsObservation) DeepCopyInto(out *ClientSideTimestampsObservation) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientSideTimestampsObservation.
func (in *ClientSideTimestampsObservation) DeepCopy() *ClientSideTimestampsObservation {
	if in == nil {
		return nil
	}
	out := new(ClientSideTimestampsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientSideTimestampsParameters) DeepCopyInto(out *ClientSideTimestampsParameters) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientSideTimestampsParameters.
func (in *ClientSideTimestampsParameters) DeepCopy() *ClientSideTimestampsParameters {
	if in == nil {
		return nil
	}
	out := new(ClientSideTimestampsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusteringKeyInitParameters) DeepCopyInto(out *ClusteringKeyInitParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrderBy != nil {
		in, out := &in.OrderBy, &out.OrderBy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusteringKeyInitParameters.
func (in *ClusteringKeyInitParameters) DeepCopy() *ClusteringKeyInitParameters {
	if in == nil {
		return nil
	}
	out := new(ClusteringKeyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusteringKeyObservation) DeepCopyInto(out *ClusteringKeyObservation) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrderBy != nil {
		in, out := &in.OrderBy, &out.OrderBy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusteringKeyObservation.
func (in *ClusteringKeyObservation) DeepCopy() *ClusteringKeyObservation {
	if in == nil {
		return nil
	}
	out := new(ClusteringKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusteringKeyParameters) DeepCopyInto(out *ClusteringKeyParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrderBy != nil {
		in, out := &in.OrderBy, &out.OrderBy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusteringKeyParameters.
func (in *ClusteringKeyParameters) DeepCopy() *ClusteringKeyParameters {
	if in == nil {
		return nil
	}
	out := new(ClusteringKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ColumnInitParameters) DeepCopyInto(out *ColumnInitParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ColumnInitParameters.
func (in *ColumnInitParameters) DeepCopy() *ColumnInitParameters {
	if in == nil {
		return nil
	}
	out := new(ColumnInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ColumnObservation) DeepCopyInto(out *ColumnObservation) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ColumnObservation.
func (in *ColumnObservation) DeepCopy() *ColumnObservation {
	if in == nil {
		return nil
	}
	out := new(ColumnObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ColumnParameters) DeepCopyInto(out *ColumnParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ColumnParameters.
func (in *ColumnParameters) DeepCopy() *ColumnParameters {
	if in == nil {
		return nil
	}
	out := new(ColumnParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommentInitParameters) DeepCopyInto(out *CommentInitParameters) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommentInitParameters.
func (in *CommentInitParameters) DeepCopy() *CommentInitParameters {
	if in == nil {
		return nil
	}
	out := new(CommentInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommentObservation) DeepCopyInto(out *CommentObservation) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommentObservation.
func (in *CommentObservation) DeepCopy() *CommentObservation {
	if in == nil {
		return nil
	}
	out := new(CommentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommentParameters) DeepCopyInto(out *CommentParameters) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommentParameters.
func (in *CommentParameters) DeepCopy() *CommentParameters {
	if in == nil {
		return nil
	}
	out := new(CommentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSpecificationInitParameters) DeepCopyInto(out *EncryptionSpecificationInitParameters) {
	*out = *in
	if in.KMSKeyIdentifier != nil {
		in, out := &in.KMSKeyIdentifier, &out.KMSKeyIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSpecificationInitParameters.
func (in *EncryptionSpecificationInitParameters) DeepCopy() *EncryptionSpecificationInitParameters {
	if in == nil {
		return nil
	}
	out := new(EncryptionSpecificationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSpecificationObservation) DeepCopyInto(out *EncryptionSpecificationObservation) {
	*out = *in
	if in.KMSKeyIdentifier != nil {
		in, out := &in.KMSKeyIdentifier, &out.KMSKeyIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSpecificationObservation.
func (in *EncryptionSpecificationObservation) DeepCopy() *EncryptionSpecificationObservation {
	if in == nil {
		return nil
	}
	out := new(EncryptionSpecificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSpecificationParameters) DeepCopyInto(out *EncryptionSpecificationParameters) {
	*out = *in
	if in.KMSKeyIdentifier != nil {
		in, out := &in.KMSKeyIdentifier, &out.KMSKeyIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSpecificationParameters.
func (in *EncryptionSpecificationParameters) DeepCopy() *EncryptionSpecificationParameters {
	if in == nil {
		return nil
	}
	out := new(EncryptionSpecificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionKeyInitParameters) DeepCopyInto(out *PartitionKeyInitParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionKeyInitParameters.
func (in *PartitionKeyInitParameters) DeepCopy() *PartitionKeyInitParameters {
	if in == nil {
		return nil
	}
	out := new(PartitionKeyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionKeyObservation) DeepCopyInto(out *PartitionKeyObservation) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionKeyObservation.
func (in *PartitionKeyObservation) DeepCopy() *PartitionKeyObservation {
	if in == nil {
		return nil
	}
	out := new(PartitionKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionKeyParameters) DeepCopyInto(out *PartitionKeyParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionKeyParameters.
func (in *PartitionKeyParameters) DeepCopy() *PartitionKeyParameters {
	if in == nil {
		return nil
	}
	out := new(PartitionKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PointInTimeRecoveryInitParameters) DeepCopyInto(out *PointInTimeRecoveryInitParameters) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PointInTimeRecoveryInitParameters.
func (in *PointInTimeRecoveryInitParameters) DeepCopy() *PointInTimeRecoveryInitParameters {
	if in == nil {
		return nil
	}
	out := new(PointInTimeRecoveryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PointInTimeRecoveryObservation) DeepCopyInto(out *PointInTimeRecoveryObservation) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PointInTimeRecoveryObservation.
func (in *PointInTimeRecoveryObservation) DeepCopy() *PointInTimeRecoveryObservation {
	if in == nil {
		return nil
	}
	out := new(PointInTimeRecoveryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PointInTimeRecoveryParameters) DeepCopyInto(out *PointInTimeRecoveryParameters) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PointInTimeRecoveryParameters.
func (in *PointInTimeRecoveryParameters) DeepCopy() *PointInTimeRecoveryParameters {
	if in == nil {
		return nil
	}
	out := new(PointInTimeRecoveryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchemaDefinitionInitParameters) DeepCopyInto(out *SchemaDefinitionInitParameters) {
	*out = *in
	if in.ClusteringKey != nil {
		in, out := &in.ClusteringKey, &out.ClusteringKey
		*out = make([]ClusteringKeyInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Column != nil {
		in, out := &in.Column, &out.Column
		*out = make([]ColumnInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PartitionKey != nil {
		in, out := &in.PartitionKey, &out.PartitionKey
		*out = make([]PartitionKeyInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StaticColumn != nil {
		in, out := &in.StaticColumn, &out.StaticColumn
		*out = make([]StaticColumnInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchemaDefinitionInitParameters.
func (in *SchemaDefinitionInitParameters) DeepCopy() *SchemaDefinitionInitParameters {
	if in == nil {
		return nil
	}
	out := new(SchemaDefinitionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchemaDefinitionObservation) DeepCopyInto(out *SchemaDefinitionObservation) {
	*out = *in
	if in.ClusteringKey != nil {
		in, out := &in.ClusteringKey, &out.ClusteringKey
		*out = make([]ClusteringKeyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Column != nil {
		in, out := &in.Column, &out.Column
		*out = make([]ColumnObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PartitionKey != nil {
		in, out := &in.PartitionKey, &out.PartitionKey
		*out = make([]PartitionKeyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StaticColumn != nil {
		in, out := &in.StaticColumn, &out.StaticColumn
		*out = make([]StaticColumnObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchemaDefinitionObservation.
func (in *SchemaDefinitionObservation) DeepCopy() *SchemaDefinitionObservation {
	if in == nil {
		return nil
	}
	out := new(SchemaDefinitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchemaDefinitionParameters) DeepCopyInto(out *SchemaDefinitionParameters) {
	*out = *in
	if in.ClusteringKey != nil {
		in, out := &in.ClusteringKey, &out.ClusteringKey
		*out = make([]ClusteringKeyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Column != nil {
		in, out := &in.Column, &out.Column
		*out = make([]ColumnParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PartitionKey != nil {
		in, out := &in.PartitionKey, &out.PartitionKey
		*out = make([]PartitionKeyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StaticColumn != nil {
		in, out := &in.StaticColumn, &out.StaticColumn
		*out = make([]StaticColumnParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchemaDefinitionParameters.
func (in *SchemaDefinitionParameters) DeepCopy() *SchemaDefinitionParameters {
	if in == nil {
		return nil
	}
	out := new(SchemaDefinitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticColumnInitParameters) DeepCopyInto(out *StaticColumnInitParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticColumnInitParameters.
func (in *StaticColumnInitParameters) DeepCopy() *StaticColumnInitParameters {
	if in == nil {
		return nil
	}
	out := new(StaticColumnInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticColumnObservation) DeepCopyInto(out *StaticColumnObservation) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticColumnObservation.
func (in *StaticColumnObservation) DeepCopy() *StaticColumnObservation {
	if in == nil {
		return nil
	}
	out := new(StaticColumnObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticColumnParameters) DeepCopyInto(out *StaticColumnParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticColumnParameters.
func (in *StaticColumnParameters) DeepCopy() *StaticColumnParameters {
	if in == nil {
		return nil
	}
	out := new(StaticColumnParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TTLInitParameters) DeepCopyInto(out *TTLInitParameters) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TTLInitParameters.
func (in *TTLInitParameters) DeepCopy() *TTLInitParameters {
	if in == nil {
		return nil
	}
	out := new(TTLInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TTLObservation) DeepCopyInto(out *TTLObservation) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TTLObservation.
func (in *TTLObservation) DeepCopy() *TTLObservation {
	if in == nil {
		return nil
	}
	out := new(TTLObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TTLParameters) DeepCopyInto(out *TTLParameters) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TTLParameters.
func (in *TTLParameters) DeepCopy() *TTLParameters {
	if in == nil {
		return nil
	}
	out := new(TTLParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Table) DeepCopyInto(out *Table) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Table.
func (in *Table) DeepCopy() *Table {
	if in == nil {
		return nil
	}
	out := new(Table)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Table) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableInitParameters) DeepCopyInto(out *TableInitParameters) {
	*out = *in
	if in.CapacitySpecification != nil {
		in, out := &in.CapacitySpecification, &out.CapacitySpecification
		*out = new(CapacitySpecificationInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientSideTimestamps != nil {
		in, out := &in.ClientSideTimestamps, &out.ClientSideTimestamps
		*out = new(ClientSideTimestampsInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(CommentInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultTimeToLive != nil {
		in, out := &in.DefaultTimeToLive, &out.DefaultTimeToLive
		*out = new(float64)
		**out = **in
	}
	if in.EncryptionSpecification != nil {
		in, out := &in.EncryptionSpecification, &out.EncryptionSpecification
		*out = new(EncryptionSpecificationInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyspaceName != nil {
		in, out := &in.KeyspaceName, &out.KeyspaceName
		*out = new(string)
		**out = **in
	}
	if in.KeyspaceNameRef != nil {
		in, out := &in.KeyspaceNameRef, &out.KeyspaceNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyspaceNameSelector != nil {
		in, out := &in.KeyspaceNameSelector, &out.KeyspaceNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.PointInTimeRecovery != nil {
		in, out := &in.PointInTimeRecovery, &out.PointInTimeRecovery
		*out = new(PointInTimeRecoveryInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.SchemaDefinition != nil {
		in, out := &in.SchemaDefinition, &out.SchemaDefinition
		*out = new(SchemaDefinitionInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(TTLInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.TableName != nil {
		in, out := &in.TableName, &out.TableName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableInitParameters.
func (in *TableInitParameters) DeepCopy() *TableInitParameters {
	if in == nil {
		return nil
	}
	out := new(TableInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableList) DeepCopyInto(out *TableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Table, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableList.
func (in *TableList) DeepCopy() *TableList {
	if in == nil {
		return nil
	}
	out := new(TableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableObservation) DeepCopyInto(out *TableObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.CapacitySpecification != nil {
		in, out := &in.CapacitySpecification, &out.CapacitySpecification
		*out = new(CapacitySpecificationObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientSideTimestamps != nil {
		in, out := &in.ClientSideTimestamps, &out.ClientSideTimestamps
		*out = new(ClientSideTimestampsObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(CommentObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultTimeToLive != nil {
		in, out := &in.DefaultTimeToLive, &out.DefaultTimeToLive
		*out = new(float64)
		**out = **in
	}
	if in.EncryptionSpecification != nil {
		in, out := &in.EncryptionSpecification, &out.EncryptionSpecification
		*out = new(EncryptionSpecificationObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.KeyspaceName != nil {
		in, out := &in.KeyspaceName, &out.KeyspaceName
		*out = new(string)
		**out = **in
	}
	if in.PointInTimeRecovery != nil {
		in, out := &in.PointInTimeRecovery, &out.PointInTimeRecovery
		*out = new(PointInTimeRecoveryObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.SchemaDefinition != nil {
		in, out := &in.SchemaDefinition, &out.SchemaDefinition
		*out = new(SchemaDefinitionObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(TTLObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.TableName != nil {
		in, out := &in.TableName, &out.TableName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableObservation.
func (in *TableObservation) DeepCopy() *TableObservation {
	if in == nil {
		return nil
	}
	out := new(TableObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableParameters) DeepCopyInto(out *TableParameters) {
	*out = *in
	if in.CapacitySpecification != nil {
		in, out := &in.CapacitySpecification, &out.CapacitySpecification
		*out = new(CapacitySpecificationParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientSideTimestamps != nil {
		in, out := &in.ClientSideTimestamps, &out.ClientSideTimestamps
		*out = new(ClientSideTimestampsParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(CommentParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultTimeToLive != nil {
		in, out := &in.DefaultTimeToLive, &out.DefaultTimeToLive
		*out = new(float64)
		**out = **in
	}
	if in.EncryptionSpecification != nil {
		in, out := &in.EncryptionSpecification, &out.EncryptionSpecification
		*out = new(EncryptionSpecificationParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyspaceName != nil {
		in, out := &in.KeyspaceName, &out.KeyspaceName
		*out = new(string)
		**out = **in
	}
	if in.KeyspaceNameRef != nil {
		in, out := &in.KeyspaceNameRef, &out.KeyspaceNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyspaceNameSelector != nil {
		in, out := &in.KeyspaceNameSelector, &out.KeyspaceNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.PointInTimeRecovery != nil {
		in, out := &in.PointInTimeRecovery, &out.PointInTimeRecovery
		*out = new(PointInTimeRecoveryParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SchemaDefinition != nil {
		in, out := &in.SchemaDefinition, &out.SchemaDefinition
		*out = new(SchemaDefinitionParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(TTLParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.TableName != nil {
		in, out := &in.TableName, &out.TableName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableParameters.
func (in *TableParameters) DeepCopy() *TableParameters {
	if in == nil {
		return nil
	}
	out := new(TableParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableSpec) DeepCopyInto(out *TableSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableSpec.
func (in *TableSpec) DeepCopy() *TableSpec {
	if in == nil {
		return nil
	}
	out := new(TableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableStatus) DeepCopyInto(out *TableStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableStatus.
func (in *TableStatus) DeepCopy() *TableStatus {
	if in == nil {
		return nil
	}
	out := new(TableStatus)
	in.DeepCopyInto(out)
	return out
}