// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package autoscaling

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/crossplane-runtime/pkg/test"

	"github.com/upbound/provider-aws/apis/namespaced/autoscaling/v1beta1"
	"github.com/upbound/provider-aws/apis/namespaced/autoscaling/v1beta2"
)

var (
	key                   = "key"
	value                 = "value"
	propagateAtLaunch     = "true"
	propagateAtLaunchBool = true
)

func TestAutoScalingGroupConverterFromv1beta1Tov1beta2(t *testing.T) {
	type args struct {
		src    xpresource.Managed
		target xpresource.Managed
	}
	type want struct {
		target xpresource.Managed
		err    error
	}
	cases := map[string]struct {
		args args
		want want
	}{
		"Successful": {
			args: args{
				src: &v1beta1.AutoscalingGroup{
					Spec: v1beta1.AutoscalingGroupSpec{
						ForProvider: v1beta1.AutoscalingGroupParameters{
							Tags: []map[string]*string{
								{
									key:                   &key,
									value:                 &value,
									"propagate_at_launch": &propagateAtLaunch,
								},
							},
						},
						InitProvider: v1beta1.AutoscalingGroupInitParameters{
							Tags: []map[string]*string{
								{
									key:                   &key,
									value:                 &value,
									"propagate_at_launch": &propagateAtLaunch,
								},
							},
						},
					},
					Status: v1beta1.AutoscalingGroupStatus{
						AtProvider: v1beta1.AutoscalingGroupObservation{
							Tags: []map[string]*string{
								{
									key:                   &key,
									value:                 &value,
									"propagate_at_launch": &propagateAtLaunch,
								},
							},
						},
					},
				},
				target: &v1beta2.AutoscalingGroup{
					Spec: v1beta2.AutoscalingGroupSpec{
						ForProvider: v1beta2.AutoscalingGroupParameters{},
					},
				},
			},
			want: want{
				target: &v1beta2.AutoscalingGroup{
					Spec: v1beta2.AutoscalingGroupSpec{
						ForProvider: v1beta2.AutoscalingGroupParameters{
							Tag: []v1beta2.TagParameters{
								{
									Key:               &key,
									Value:             &value,
									PropagateAtLaunch: &propagateAtLaunchBool,
								},
							},
						},
						InitProvider: v1beta2.AutoscalingGroupInitParameters{
							Tag: []v1beta2.TagInitParameters{
								{
									Key:               &key,
									Value:             &value,
									PropagateAtLaunch: &propagateAtLaunchBool,
								},
							},
						},
					},
					Status: v1beta2.AutoscalingGroupStatus{
						AtProvider: v1beta2.AutoscalingGroupObservation{
							Tag: []v1beta2.TagObservation{
								{
									Key:               &key,
									Value:             &value,
									PropagateAtLaunch: &propagateAtLaunchBool,
								},
							},
						},
					},
				},
				err: nil,
			},
		},
		"Unsuccessful": {
			args: args{
				src: &v1beta1.AutoscalingGroup{
					Spec: v1beta1.AutoscalingGroupSpec{
						ForProvider: v1beta1.AutoscalingGroupParameters{
							Tags: []map[string]*string{
								{
									key:                   &key,
									value:                 &value,
									"propagate_at_launch": &value,
								},
							},
						},
					},
				},
				target: &v1beta2.AutoscalingGroup{
					Spec: v1beta2.AutoscalingGroupSpec{
						ForProvider: v1beta2.AutoscalingGroupParameters{},
					},
				},
			},
			want: want{
				err:    &strconv.NumError{Func: "ParseBool", Num: value, Err: errors.New("invalid syntax")},
				target: &v1beta2.AutoscalingGroup{},
			},
		},
		"MissingKey": {
			args: args{
				src: &v1beta1.AutoscalingGroup{
					Spec: v1beta1.AutoscalingGroupSpec{
						ForProvider: v1beta1.AutoscalingGroupParameters{
							Tags: []map[string]*string{
								{
									value:                 &value,
									"propagate_at_launch": &propagateAtLaunch,
								},
							},
						},
					},
				},
				target: &v1beta2.AutoscalingGroup{
					Spec: v1beta2.AutoscalingGroupSpec{
						ForProvider: v1beta2.AutoscalingGroupParameters{},
					},
				},
			},
			want: want{
				target: &v1beta2.AutoscalingGroup{
					Spec: v1beta2.AutoscalingGroupSpec{
						ForProvider: v1beta2.AutoscalingGroupParameters{
							Tag: []v1beta2.TagParameters{
								{
									Value:             &value,
									PropagateAtLaunch: &propagateAtLaunchBool,
								},
							},
						},
					},
				},
				err: nil,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := autoScalingGroupConverterFromv1beta1Tov1beta2(tc.args.src, tc.args.target)
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("autoScalingGroupConverterFromv1beta1Tov1beta2(...): -want error, +got error:\n%s", diff)
			}
			if diff := cmp.Diff(tc.want.target, tc.args.target, test.EquateErrors()); diff != "" {
				t.Errorf("autoScalingGroupConverterFromv1beta1Tov1beta2(...): -want target, +got target:\n%s", diff)
			}
		})
	}
}

func TestAutoScalingGroupConverterFromv1beta2Tov1beta1(t *testing.T) {
	type args struct {
		src    xpresource.Managed
		target xpresource.Managed
	}
	type want struct {
		target xpresource.Managed
		err    error
	}
	cases := map[string]struct {
		args args
		want want
	}{
		"Successful": {
			args: args{
				src: &v1beta2.AutoscalingGroup{
					Spec: v1beta2.AutoscalingGroupSpec{
						ForProvider: v1beta2.AutoscalingGroupParameters{
							Tag: []v1beta2.TagParameters{
								{
									Key:               &key,
									Value:             &value,
									PropagateAtLaunch: &propagateAtLaunchBool,
								},
							},
						},
						InitProvider: v1beta2.AutoscalingGroupInitParameters{
							Tag: []v1beta2.TagInitParameters{
								{
									Key:               &key,
									Value:             &value,
									PropagateAtLaunch: &propagateAtLaunchBool,
								},
							},
						},
					},
					Status: v1beta2.AutoscalingGroupStatus{
						AtProvider: v1beta2.AutoscalingGroupObservation{
							Tag: []v1beta2.TagObservation{
								{
									Key:               &key,
									Value:             &value,
									PropagateAtLaunch: &propagateAtLaunchBool,
								},
							},
						},
					},
				},
				target: &v1beta1.AutoscalingGroup{
					Spec: v1beta1.AutoscalingGroupSpec{
						ForProvider: v1beta1.AutoscalingGroupParameters{},
					},
				},
			},
			want: want{
				target: &v1beta1.AutoscalingGroup{
					Spec: v1beta1.AutoscalingGroupSpec{
						ForProvider: v1beta1.AutoscalingGroupParameters{
							Tags: []map[string]*string{
								{
									key:                   &key,
									value:                 &value,
									"propagate_at_launch": &propagateAtLaunch,
								},
							},
						},
						InitProvider: v1beta1.AutoscalingGroupInitParameters{
							Tags: []map[string]*string{
								{
									key:                   &key,
									value:                 &value,
									"propagate_at_launch": &propagateAtLaunch,
								},
							},
						},
					},
					Status: v1beta1.AutoscalingGroupStatus{
						AtProvider: v1beta1.AutoscalingGroupObservation{
							Tags: []map[string]*string{
								{
									key:                   &key,
									value:                 &value,
									"propagate_at_launch": &propagateAtLaunch,
								},
							},
						},
					},
				},
			},
		},
		"MissingKey": {
			args: args{
				src: &v1beta2.AutoscalingGroup{
					Spec: v1beta2.AutoscalingGroupSpec{
						ForProvider: v1beta2.AutoscalingGroupParameters{
							Tag: []v1beta2.TagParameters{
								{
									Key:   &key,
									Value: &value,
								},
							},
						},
						InitProvider: v1beta2.AutoscalingGroupInitParameters{
							Tag: []v1beta2.TagInitParameters{
								{
									Key:               &key,
									PropagateAtLaunch: &propagateAtLaunchBool,
								},
							},
						},
					},
					Status: v1beta2.AutoscalingGroupStatus{
						AtProvider: v1beta2.AutoscalingGroupObservation{
							Tag: []v1beta2.TagObservation{
								{
									Value:             &value,
									PropagateAtLaunch: &propagateAtLaunchBool,
								},
							},
						},
					},
				},
				target: &v1beta1.AutoscalingGroup{
					Spec: v1beta1.AutoscalingGroupSpec{
						ForProvider: v1beta1.AutoscalingGroupParameters{},
					},
				},
			},
			want: want{
				target: &v1beta1.AutoscalingGroup{
					Spec: v1beta1.AutoscalingGroupSpec{
						ForProvider: v1beta1.AutoscalingGroupParameters{
							Tags: []map[string]*string{
								{
									key:   &key,
									value: &value,
								},
							},
						},
						InitProvider: v1beta1.AutoscalingGroupInitParameters{
							Tags: []map[string]*string{
								{
									key:                   &key,
									"propagate_at_launch": &propagateAtLaunch,
								},
							},
						},
					},
					Status: v1beta1.AutoscalingGroupStatus{
						AtProvider: v1beta1.AutoscalingGroupObservation{
							Tags: []map[string]*string{
								{
									value:                 &value,
									"propagate_at_launch": &propagateAtLaunch,
								},
							},
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := autoScalingGroupConverterFromv1beta2Tov1beta1(tc.args.src, tc.args.target)
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("autoScalingGroupConverterFromv1beta1Tov1beta2(...): -want error, +got error:\n%s", diff)
			}
			if diff := cmp.Diff(tc.want.target, tc.args.target, test.EquateErrors()); diff != "" {
				t.Errorf("autoScalingGroupConverterFromv1beta1Tov1beta2(...): -want target, +got target:\n%s", diff)
			}
		})
	}
}
