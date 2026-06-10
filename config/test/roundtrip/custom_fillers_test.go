// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

//nolint:typecheck // due to buildtagger constraints

package roundtrip

import (
	"reflect"

	v1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/upjet/v2/pkg/apitesting/roundtrip"
	"github.com/google/go-cmp/cmp"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/randfill"

	autoscalingv1beta1 "github.com/upbound/provider-aws/v2/apis/cluster/autoscaling/v1beta1"
	connectv1beta1 "github.com/upbound/provider-aws/v2/apis/cluster/connect/v1beta1"
	ec2v1beta1 "github.com/upbound/provider-aws/v2/apis/cluster/ec2/v1beta1"
	ec2v1beta2 "github.com/upbound/provider-aws/v2/apis/cluster/ec2/v1beta2"
	elasticachev1beta1 "github.com/upbound/provider-aws/v2/apis/cluster/elasticache/v1beta1"
	kafkav1beta1 "github.com/upbound/provider-aws/v2/apis/cluster/kafka/v1beta1"
	rdsv1beta1 "github.com/upbound/provider-aws/v2/apis/cluster/rds/v1beta1"
	redshiftv1beta1 "github.com/upbound/provider-aws/v2/apis/cluster/redshift/v1beta1"
)

var awsCustomFuzzers = []roundtrip.FuzzFunc{
	fuzzAutoscalingGroupV1Beta1,
	fuzzAutoscalingAttachmentV1Beta1,
	fuzzConnectHoursOfOperationV1Beta1,
	fuzzConnectQueueV1Beta1,
	fuzzConnectRoutingProfileV1Beta1,
	fuzzEC2RouteV1Beta1,
	fuzzEC2RouteV1Beta2,
	fuzzElasticacheReplicationGroupV1Beta1,
	fuzzKafkaClusterV1Beta1,
	fuzzRDSInstanceV1Beta1,
	fuzzRedshiftClusterV1Beta1,
	fuzzSecretKeySelectors,
	fuzzLocalSecretKeySelectors,
}

var awsCustomCmpOpts = []cmp.Option{
	equateElasticacheReplicationGroup(),
	equateEC2Route(),
	equateConnectRoutingProfile(),
	equateConnectHoursOfOperation(),
	equateConnectQueue(),
}

func keepFirstNonNil(fields ...any) int {
	kept := -1

	for i, f := range fields {
		v := reflect.ValueOf(f)
		if !v.IsValid() || v.Kind() != reflect.Ptr {
			continue
		}
		elem := v.Elem()
		if elem.Kind() != reflect.Ptr && elem.Kind() != reflect.Slice {
			// Expecting pointer-to-pointer (**T) or pointer-to-slice
			continue
		}
		if !elem.CanSet() {
			// Can't set (often unexported field from another package)
			continue
		}

		if elem.IsNil() {
			continue
		}

		if kept == -1 {
			kept = i
			continue
		}

		// More than one non-nil: nil this one out
		elem.Set(reflect.Zero(elem.Type())) // sets to nil
	}

	return kept
}

func fuzzAutoscalingGroupV1Beta1(s *autoscalingv1beta1.AutoscalingGroup, c randfill.Continue) {
	c.Fill(s)
	keepFirstNonNil(&s.Spec.ForProvider.Tags, &s.Spec.ForProvider.Tag)
	keepFirstNonNil(&s.Spec.InitProvider.Tags, &s.Spec.InitProvider.Tag)
	keepFirstNonNil(&s.Status.AtProvider.Tags, &s.Status.AtProvider.Tag)
}

func fuzzAutoscalingAttachmentV1Beta1(s *autoscalingv1beta1.Attachment, c randfill.Continue) {
	c.Fill(s)
	keepFirstNonNil(&s.Spec.ForProvider.ALBTargetGroupArn, &s.Spec.ForProvider.LBTargetGroupArn)
	keepFirstNonNil(&s.Spec.InitProvider.ALBTargetGroupArn, &s.Spec.InitProvider.LBTargetGroupArn)
	keepFirstNonNil(&s.Status.AtProvider.ALBTargetGroupArn, &s.Status.AtProvider.LBTargetGroupArn)
}

func fuzzEC2RouteV1Beta1(s *ec2v1beta1.Route, c randfill.Continue) {
	c.Fill(s)
	keepFirstNonNil(&s.Spec.ForProvider.NetworkInterfaceID, &s.Spec.ForProvider.InstanceID)
	keepFirstNonNil(&s.Spec.InitProvider.NetworkInterfaceID, &s.Spec.InitProvider.InstanceID)
	keepFirstNonNil(&s.Status.AtProvider.NetworkInterfaceID, &s.Status.AtProvider.InstanceID)

	s.Spec.ForProvider.InstanceIDRef = nil
	s.Spec.InitProvider.InstanceIDRef = nil
	s.Spec.ForProvider.InstanceIDSelector = nil
	s.Spec.InitProvider.InstanceIDSelector = nil

}

func fuzzEC2RouteV1Beta2(s *ec2v1beta2.Route, c randfill.Continue) {
	c.Fill(s)
	keepFirstNonNil(&s.Status.AtProvider.NetworkInterfaceID, &s.Status.AtProvider.InstanceID)
}

func fuzzRDSInstanceV1Beta1(s *rdsv1beta1.Instance, c randfill.Continue) {
	c.Fill(s)
	keepFirstNonNil(&s.Spec.ForProvider.DBName, &s.Spec.ForProvider.Name)
	keepFirstNonNil(&s.Spec.InitProvider.DBName, &s.Spec.InitProvider.Name)
	keepFirstNonNil(&s.Status.AtProvider.DBName, &s.Status.AtProvider.Name)
}

func fuzzRedshiftClusterV1Beta1(s *redshiftv1beta1.Cluster, c randfill.Continue) {
	c.Fill(s)
	if s.Spec.ForProvider.Encrypted != nil {
		s.Spec.ForProvider.Encrypted = ptr.To("false")
		if c.Bool() {
			s.Spec.ForProvider.Encrypted = ptr.To("true")
		}
	}
}

func fuzzKafkaClusterV1Beta1(s *kafkav1beta1.Cluster, c randfill.Continue) {
	c.Fill(s)
	if len(s.Spec.ForProvider.BrokerNodeGroupInfo) > 0 &&
		len(s.Spec.ForProvider.BrokerNodeGroupInfo[0].StorageInfo) > 0 &&
		len(s.Spec.ForProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo) > 0 {
		keepFirstNonNil(&s.Spec.ForProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo[0].VolumeSize, &s.Spec.ForProvider.BrokerNodeGroupInfo[0].EBSVolumeSize)
	}
	if len(s.Spec.InitProvider.BrokerNodeGroupInfo) > 0 &&
		len(s.Spec.InitProvider.BrokerNodeGroupInfo[0].StorageInfo) > 0 &&
		len(s.Spec.InitProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo) > 0 {

		keepFirstNonNil(&s.Spec.InitProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo[0].VolumeSize, &s.Spec.InitProvider.BrokerNodeGroupInfo[0].EBSVolumeSize)
	}
	if len(s.Status.AtProvider.BrokerNodeGroupInfo) > 0 &&
		len(s.Status.AtProvider.BrokerNodeGroupInfo[0].StorageInfo) > 0 &&
		len(s.Status.AtProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo) > 0 {

		keepFirstNonNil(&s.Status.AtProvider.BrokerNodeGroupInfo[0].StorageInfo[0].EBSStorageInfo[0].VolumeSize, &s.Status.AtProvider.BrokerNodeGroupInfo[0].EBSVolumeSize)
	}
}

func fuzzConnectHoursOfOperationV1Beta1(s *connectv1beta1.HoursOfOperation, c randfill.Continue) {
	c.FillNoCustom(s)
	keepFirstNonNil(&s.Status.AtProvider.Arn, &s.Status.AtProvider.HoursOfOperationArn)
}

func fuzzConnectQueueV1Beta1(s *connectv1beta1.Queue, c randfill.Continue) {
	c.FillNoCustom(s)
	keepFirstNonNil(&s.Status.AtProvider.QuickConnectIds, &s.Status.AtProvider.QuickConnectIdsAssociated)
}

func fuzzConnectRoutingProfileV1Beta1(s *connectv1beta1.RoutingProfile, c randfill.Continue) {
	c.FillNoCustom(s)
	keepFirstNonNil(&s.Status.AtProvider.QueueConfigsAssociated, &s.Status.AtProvider.QueueConfigs)
}

func fuzzElasticacheReplicationGroupV1Beta1(s *elasticachev1beta1.ReplicationGroup, c randfill.Continue) {
	c.FillNoCustom(s)
	keepFirstNonNil(&s.Spec.ForProvider.NumberCacheClusters, &s.Spec.ForProvider.NumCacheClusters)
	keepFirstNonNil(&s.Spec.ForProvider.ReplicationGroupDescription, &s.Spec.ForProvider.Description)
	keepFirstNonNil(&s.Spec.ForProvider.AvailabilityZones, &s.Spec.ForProvider.PreferredCacheClusterAzs)

	keepFirstNonNil(&s.Spec.InitProvider.NumberCacheClusters, &s.Spec.InitProvider.NumCacheClusters)
	keepFirstNonNil(&s.Spec.InitProvider.ReplicationGroupDescription, &s.Spec.InitProvider.Description)
	keepFirstNonNil(&s.Spec.InitProvider.AvailabilityZones, &s.Spec.InitProvider.PreferredCacheClusterAzs)

	keepFirstNonNil(&s.Status.AtProvider.NumberCacheClusters, &s.Status.AtProvider.NumCacheClusters)
	keepFirstNonNil(&s.Status.AtProvider.ReplicationGroupDescription, &s.Status.AtProvider.Description)
	keepFirstNonNil(&s.Status.AtProvider.AvailabilityZones, &s.Status.AtProvider.PreferredCacheClusterAzs)

	if len(s.Spec.ForProvider.ClusterMode) > 0 {
		s.Spec.ForProvider.NumNodeGroups = nil
		s.Spec.ForProvider.ReplicasPerNodeGroup = nil
	}
	if len(s.Spec.InitProvider.ClusterMode) > 0 {
		s.Spec.InitProvider.NumNodeGroups = nil
		s.Spec.InitProvider.ReplicasPerNodeGroup = nil
	}
	if len(s.Status.AtProvider.ClusterMode) > 0 {
		s.Status.AtProvider.NumNodeGroups = nil
		s.Status.AtProvider.ReplicasPerNodeGroup = nil
	}
}

func fuzzSecretKeySelectors(p *[]v1.SecretKeySelector, c randfill.Continue) {
	c.Fill(p)
	if p != nil && *p == nil {
		*p = []v1.SecretKeySelector{}
	}
}

func fuzzLocalSecretKeySelectors(p *[]v1.LocalSecretKeySelector, c randfill.Continue) {
	c.Fill(p)
	if p != nil && *p == nil {
		*p = []v1.LocalSecretKeySelector{}
	}
}

// Following functions returns a cmp.Option to ignore
// difference for the following case:
// Field `OldFoo` in v1beta1 was renamed to `NewBar` in v1beta2,
// but `NewBar` is also backported to v1beta1. Example:
//    original := v1beta1.MyParameters{ OldFoo: &someValue, NewBar: nil }
// then we roundtrip and end up with:
//    afterRT := v1beta1.MyParameters{ OldFoo: &someValue, NewBar: &someValue }
// These values are semantically equivalent, so we ignore the diff here

func equateConnectHoursOfOperation() cmp.Option {
	return cmp.Transformer("NormalizeHoursOfOperationArn", func(h connectv1beta1.HoursOfOperationObservation) connectv1beta1.HoursOfOperationObservation {
		if h.HoursOfOperationArn == nil && h.Arn != nil {
			h.HoursOfOperationArn = h.Arn
		} else if h.HoursOfOperationArn != nil && h.Arn == nil {
			h.Arn = h.HoursOfOperationArn
		}
		return h
	})
}

func equateConnectQueue() cmp.Option {
	return cmp.Transformer("NormalizeQueueQuickConnectIds", func(h connectv1beta1.QueueObservation) connectv1beta1.QueueObservation {
		if h.QuickConnectIdsAssociated == nil && h.QuickConnectIds != nil {
			h.QuickConnectIdsAssociated = h.QuickConnectIds
		} else if h.QuickConnectIdsAssociated != nil && h.QuickConnectIds == nil {
			h.QuickConnectIds = h.QuickConnectIdsAssociated
		}
		return h
	})
}

func equateConnectRoutingProfile() cmp.Option {
	return cmp.Transformer("NormalizeRoutingProfile", func(h connectv1beta1.RoutingProfileObservation) connectv1beta1.RoutingProfileObservation {
		if h.QueueConfigsAssociated == nil && h.QueueConfigs != nil {
			for _, qc := range h.QueueConfigs {
				rep := connectv1beta1.QueueConfigsAssociatedObservation{ //nolint:staticcheck
					Channel:   qc.Channel,
					Delay:     qc.Delay,
					Priority:  qc.Priority,
					QueueArn:  qc.QueueArn,
					QueueID:   qc.QueueID,
					QueueName: qc.QueueName,
				}
				h.QueueConfigsAssociated = append(h.QueueConfigsAssociated, rep)
			}
		}
		if h.QueueConfigsAssociated != nil && h.QueueConfigs == nil {
			for _, qca := range h.QueueConfigsAssociated {
				rep := connectv1beta1.QueueConfigsObservation{ //nolint:staticcheck
					Channel:   qca.Channel,
					Delay:     qca.Delay,
					Priority:  qca.Priority,
					QueueArn:  qca.QueueArn,
					QueueID:   qca.QueueID,
					QueueName: qca.QueueName,
				}
				h.QueueConfigs = append(h.QueueConfigs, rep)
			}
		}
		return h
	})
}

func equateEC2Route() cmp.Option {
	return cmp.Transformer("NormalizeRoute", func(h ec2v1beta1.Route) ec2v1beta1.Route {
		equateSpecField[ec2v1beta1.Route](&h, "InstanceID", "NetworkInterfaceID", true)
		return h
	})
}

func equateStatusField[T any](h *T, fx, fy string) {
	xval := reflect.ValueOf(h).Elem().FieldByName("Status").FieldByName("AtProvider").FieldByName(fx)
	yval := reflect.ValueOf(h).Elem().FieldByName("Status").FieldByName("AtProvider").FieldByName(fy)

	if xval.IsNil() && !yval.IsNil() {
		xval.Set(yval)
	} else if yval.IsNil() && !xval.IsNil() {
		yval.Set(xval)
	}
}

func equateSpecField[T any](h *T, fx, fy string, init bool) {
	xval := reflect.ValueOf(h).Elem().FieldByName("Spec").FieldByName("ForProvider").FieldByName(fx)
	yval := reflect.ValueOf(h).Elem().FieldByName("Spec").FieldByName("ForProvider").FieldByName(fy)

	if xval.IsNil() && !yval.IsNil() {
		xval.Set(yval)
	} else if yval.IsNil() && !xval.IsNil() {
		yval.Set(xval)
	}

	if !init {
		return
	}
	xival := reflect.ValueOf(h).Elem().FieldByName("Spec").FieldByName("InitProvider").FieldByName(fx)
	yival := reflect.ValueOf(h).Elem().FieldByName("Spec").FieldByName("InitProvider").FieldByName(fy)

	if xival.IsNil() && !yival.IsNil() {
		xival.Set(yival)
	} else if yival.IsNil() && !xival.IsNil() {
		yival.Set(xival)
	}
}

func equateElasticacheReplicationGroup() cmp.Option {
	return cmp.Transformer("NormalizeReplicationGroup", func(h elasticachev1beta1.ReplicationGroup) elasticachev1beta1.ReplicationGroup {
		equateSpecField[elasticachev1beta1.ReplicationGroup](&h, "NumberCacheClusters", "NumCacheClusters", true)
		equateSpecField[elasticachev1beta1.ReplicationGroup](&h, "ReplicationGroupDescription", "Description", true)
		equateSpecField[elasticachev1beta1.ReplicationGroup](&h, "AvailabilityZones", "PreferredCacheClusterAzs", true)

		equateStatusField[elasticachev1beta1.ReplicationGroup](&h, "NumberCacheClusters", "NumCacheClusters")
		equateStatusField[elasticachev1beta1.ReplicationGroup](&h, "ReplicationGroupDescription", "Description")
		equateStatusField[elasticachev1beta1.ReplicationGroup](&h, "AvailabilityZones", "PreferredCacheClusterAzs")

		// ClusterMode in v1beta1 is a structural wrapper around NumNodeGroups/ReplicasPerNodeGroup.
		// After a roundtrip through v1beta2 (where they are top-level fields), both representations
		// end up populated. Normalize to a canonical form so comparison succeeds.
		equateClusterModeForProvider(&h.Spec.ForProvider)
		equateClusterModeInitProvider(&h.Spec.InitProvider)
		equateClusterModeAtProvider(&h.Status.AtProvider)

		return h
	})
}

func equateClusterModeForProvider(fp *elasticachev1beta1.ReplicationGroupParameters) {
	if len(fp.ClusterMode) > 0 {
		if fp.NumNodeGroups == nil {
			fp.NumNodeGroups = fp.ClusterMode[0].NumNodeGroups
		}
		if fp.ReplicasPerNodeGroup == nil {
			fp.ReplicasPerNodeGroup = fp.ClusterMode[0].ReplicasPerNodeGroup
		}
	} else if fp.NumNodeGroups != nil || fp.ReplicasPerNodeGroup != nil {
		fp.ClusterMode = []elasticachev1beta1.ClusterModeParameters{{
			NumNodeGroups:        fp.NumNodeGroups,
			ReplicasPerNodeGroup: fp.ReplicasPerNodeGroup,
		}}
	}
}

func equateClusterModeInitProvider(ip *elasticachev1beta1.ReplicationGroupInitParameters) {
	if len(ip.ClusterMode) > 0 {
		if ip.NumNodeGroups == nil {
			ip.NumNodeGroups = ip.ClusterMode[0].NumNodeGroups
		}
		if ip.ReplicasPerNodeGroup == nil {
			ip.ReplicasPerNodeGroup = ip.ClusterMode[0].ReplicasPerNodeGroup
		}
	} else if ip.NumNodeGroups != nil || ip.ReplicasPerNodeGroup != nil {
		ip.ClusterMode = []elasticachev1beta1.ClusterModeInitParameters{{
			NumNodeGroups:        ip.NumNodeGroups,
			ReplicasPerNodeGroup: ip.ReplicasPerNodeGroup,
		}}
	}
}

func equateClusterModeAtProvider(ap *elasticachev1beta1.ReplicationGroupObservation) {
	if len(ap.ClusterMode) > 0 {
		if ap.NumNodeGroups == nil {
			ap.NumNodeGroups = ap.ClusterMode[0].NumNodeGroups
		}
		if ap.ReplicasPerNodeGroup == nil {
			ap.ReplicasPerNodeGroup = ap.ClusterMode[0].ReplicasPerNodeGroup
		}
	} else if ap.NumNodeGroups != nil || ap.ReplicasPerNodeGroup != nil {
		ap.ClusterMode = []elasticachev1beta1.ClusterModeObservation{{
			NumNodeGroups:        ap.NumNodeGroups,
			ReplicasPerNodeGroup: ap.ReplicasPerNodeGroup,
		}}
	}
}
