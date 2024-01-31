/*
Copyright 2021 Upbound Inc.
*/

package budgets

import (
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/crossplane-runtime/pkg/test"
	"github.com/google/go-cmp/cmp"
	"github.com/upbound/provider-aws/apis/budgets/v1beta1"
	"github.com/upbound/provider-aws/apis/budgets/v1beta2"
	"testing"
)

var (
	costFilterKey       = "Service"
	costFilterVal       = "Amazon Elastic Compute Cloud - Compute"
	costFilterOperation = "Operation"
)

func TestBudgetConverterFromv1beta1Tov1beta2(t *testing.T) {
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
				src: &v1beta1.Budget{
					Spec: v1beta1.BudgetSpec{
						ForProvider: v1beta1.BudgetParameters{
							CostFilters: map[string]*string{
								costFilterKey: &costFilterVal,
							},
						},
						InitProvider: v1beta1.BudgetInitParameters{
							CostFilters: map[string]*string{
								costFilterKey:       &costFilterVal,
								costFilterOperation: &costFilterOperation,
							},
						},
					},
					Status: v1beta1.BudgetStatus{
						AtProvider: v1beta1.BudgetObservation{
							CostFilters: map[string]*string{
								costFilterKey: &costFilterVal,
							},
						},
					},
				},
				target: &v1beta2.Budget{
					Spec: v1beta2.BudgetSpec{
						ForProvider: v1beta2.BudgetParameters{},
					},
				},
			},
			want: want{
				target: &v1beta2.Budget{
					Spec: v1beta2.BudgetSpec{
						ForProvider: v1beta2.BudgetParameters{
							CostFilter: []v1beta2.CostFilterParameters{
								{
									Name:   &costFilterKey,
									Values: []*string{&costFilterVal},
								},
							},
						},
						InitProvider: v1beta2.BudgetInitParameters{
							CostFilter: []v1beta2.CostFilterInitParameters{
								{
									Name:   &costFilterKey,
									Values: []*string{&costFilterVal},
								},
								{
									Name:   &costFilterOperation,
									Values: []*string{&costFilterOperation},
								},
							},
						},
					},
					Status: v1beta2.BudgetStatus{
						AtProvider: v1beta2.BudgetObservation{
							CostFilter: []v1beta2.CostFilterObservation{
								{
									Name:   &costFilterKey,
									Values: []*string{&costFilterVal},
								},
							},
						},
					},
				},
			},
		},
		"Addition": {
			args: args{
				src: &v1beta1.Budget{
					Spec: v1beta1.BudgetSpec{
						ForProvider: v1beta1.BudgetParameters{
							CostFilters: map[string]*string{
								costFilterKey: &costFilterVal,
							},
						},
					},
				},
				target: &v1beta2.Budget{
					Spec: v1beta2.BudgetSpec{
						ForProvider: v1beta2.BudgetParameters{
							CostFilter: []v1beta2.CostFilterParameters{
								{
									Name:   &costFilterOperation,
									Values: []*string{&costFilterOperation},
								},
							},
						},
					},
				},
			},
			want: want{
				target: &v1beta2.Budget{
					Spec: v1beta2.BudgetSpec{
						ForProvider: v1beta2.BudgetParameters{
							CostFilter: []v1beta2.CostFilterParameters{
								{
									Name:   &costFilterOperation,
									Values: []*string{&costFilterOperation},
								},
								{
									Name:   &costFilterKey,
									Values: []*string{&costFilterVal},
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
			err := budgetConverterFromv1beta1Tov1beta2(tc.args.src, tc.args.target)
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("autoScalingGroupConverterFromv1beta1Tov1beta2(...): -want error, +got error:\n%s", diff)
			}
			if diff := cmp.Diff(tc.want.target, tc.args.target, test.EquateErrors()); diff != "" {
				t.Errorf("autoScalingGroupConverterFromv1beta1Tov1beta2(...): -want target, +got target:\n%s", diff)
			}
		})
	}
}

func TestBudgetConverterFromv1beta2Tov1beta1(t *testing.T) {
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
				src: &v1beta2.Budget{
					Spec: v1beta2.BudgetSpec{
						ForProvider: v1beta2.BudgetParameters{
							CostFilter: []v1beta2.CostFilterParameters{
								{
									Name:   &costFilterKey,
									Values: []*string{&costFilterVal},
								},
							},
						},
						InitProvider: v1beta2.BudgetInitParameters{
							CostFilter: []v1beta2.CostFilterInitParameters{
								{
									Name:   &costFilterKey,
									Values: []*string{&costFilterVal},
								},
								{
									Name:   &costFilterOperation,
									Values: []*string{&costFilterOperation},
								},
							},
						},
					},
					Status: v1beta2.BudgetStatus{
						AtProvider: v1beta2.BudgetObservation{
							CostFilter: []v1beta2.CostFilterObservation{
								{
									Name:   &costFilterKey,
									Values: []*string{&costFilterVal},
								},
							},
						},
					},
				},
				target: &v1beta1.Budget{
					Spec: v1beta1.BudgetSpec{
						ForProvider: v1beta1.BudgetParameters{},
					},
				},
			},
			want: want{
				target: &v1beta1.Budget{
					Spec: v1beta1.BudgetSpec{
						ForProvider: v1beta1.BudgetParameters{
							CostFilters: map[string]*string{
								costFilterKey: &costFilterVal,
							},
						},
						InitProvider: v1beta1.BudgetInitParameters{
							CostFilters: map[string]*string{
								costFilterKey:       &costFilterVal,
								costFilterOperation: &costFilterOperation,
							},
						},
					},
					Status: v1beta1.BudgetStatus{
						AtProvider: v1beta1.BudgetObservation{
							CostFilters: map[string]*string{
								costFilterKey: &costFilterVal,
							},
						},
					},
				},
			},
		},
		"Addition": {
			args: args{
				src: &v1beta2.Budget{
					Spec: v1beta2.BudgetSpec{
						ForProvider: v1beta2.BudgetParameters{
							CostFilter: []v1beta2.CostFilterParameters{
								{
									Name:   &costFilterKey,
									Values: []*string{&costFilterVal},
								},
							},
						},
					},
				},
				target: &v1beta1.Budget{
					Spec: v1beta1.BudgetSpec{
						ForProvider: v1beta1.BudgetParameters{
							CostFilters: map[string]*string{
								costFilterOperation: &costFilterOperation,
							},
						},
					},
				},
			},
			want: want{
				target: &v1beta1.Budget{
					Spec: v1beta1.BudgetSpec{
						ForProvider: v1beta1.BudgetParameters{
							CostFilters: map[string]*string{
								costFilterOperation: &costFilterOperation,
								costFilterKey:       &costFilterVal,
							},
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := budgetConverterFromv1beta2Tov1beta1(tc.args.src, tc.args.target)
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("autoScalingGroupConverterFromv1beta1Tov1beta2(...): -want error, +got error:\n%s", diff)
			}
			if diff := cmp.Diff(tc.want.target, tc.args.target, test.EquateErrors()); diff != "" {
				t.Errorf("autoScalingGroupConverterFromv1beta1Tov1beta2(...): -want target, +got target:\n%s", diff)
			}
		})
	}
}
