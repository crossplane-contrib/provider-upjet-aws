/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ControlPanel.
func (mg *ControlPanel) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterArn),
		Extract:      common.TerraformID(),
		Reference:    mg.Spec.ForProvider.ClusterArnRef,
		Selector:     mg.Spec.ForProvider.ClusterArnSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterArn")
	}
	mg.Spec.ForProvider.ClusterArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterArnRef = rsp.ResolvedReference

	return nil
}