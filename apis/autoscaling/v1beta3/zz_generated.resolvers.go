// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta3

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-aws/config/common"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *AutoscalingGroup) ResolveReferences( // ResolveReferences of this AutoscalingGroup.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("autoscaling.aws.upbound.io", "v1beta2", "LaunchConfiguration", "LaunchConfigurationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LaunchConfiguration),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LaunchConfigurationRef,
			Selector:     mg.Spec.ForProvider.LaunchConfigurationSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LaunchConfiguration")
	}
	mg.Spec.ForProvider.LaunchConfiguration = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LaunchConfigurationRef = rsp.ResolvedReference

	if mg.Spec.ForProvider.LaunchTemplate != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta2", "LaunchTemplate", "LaunchTemplateList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LaunchTemplate.ID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.LaunchTemplate.IDRef,
				Selector:     mg.Spec.ForProvider.LaunchTemplate.IDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.LaunchTemplate.ID")
		}
		mg.Spec.ForProvider.LaunchTemplate.ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.LaunchTemplate.IDRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.MixedInstancesPolicy != nil {
		if mg.Spec.ForProvider.MixedInstancesPolicy.LaunchTemplate != nil {
			if mg.Spec.ForProvider.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification != nil {
				{
					m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta2", "LaunchTemplate", "LaunchTemplateList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification.LaunchTemplateID),
						Extract:      resource.ExtractResourceID(),
						Reference:    mg.Spec.ForProvider.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification.LaunchTemplateIDRef,
						Selector:     mg.Spec.ForProvider.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification.LaunchTemplateIDSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification.LaunchTemplateID")
				}
				mg.Spec.ForProvider.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification.LaunchTemplateID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification.LaunchTemplateIDRef = rsp.ResolvedReference

			}
		}
	}
	if mg.Spec.ForProvider.MixedInstancesPolicy != nil {
		if mg.Spec.ForProvider.MixedInstancesPolicy.LaunchTemplate != nil {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.MixedInstancesPolicy.LaunchTemplate.Override); i5++ {
				if mg.Spec.ForProvider.MixedInstancesPolicy.LaunchTemplate.Override[i5].LaunchTemplateSpecification != nil {
					{
						m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta2", "LaunchTemplate", "LaunchTemplateList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MixedInstancesPolicy.LaunchTemplate.Override[i5].LaunchTemplateSpecification.LaunchTemplateID),
							Extract:      resource.ExtractResourceID(),
							Reference:    mg.Spec.ForProvider.MixedInstancesPolicy.LaunchTemplate.Override[i5].LaunchTemplateSpecification.LaunchTemplateIDRef,
							Selector:     mg.Spec.ForProvider.MixedInstancesPolicy.LaunchTemplate.Override[i5].LaunchTemplateSpecification.LaunchTemplateIDSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.MixedInstancesPolicy.LaunchTemplate.Override[i5].LaunchTemplateSpecification.LaunchTemplateID")
					}
					mg.Spec.ForProvider.MixedInstancesPolicy.LaunchTemplate.Override[i5].LaunchTemplateSpecification.LaunchTemplateID = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.MixedInstancesPolicy.LaunchTemplate.Override[i5].LaunchTemplateSpecification.LaunchTemplateIDRef = rsp.ResolvedReference

				}
			}
		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "PlacementGroup", "PlacementGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PlacementGroup),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.PlacementGroupRef,
			Selector:     mg.Spec.ForProvider.PlacementGroupSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PlacementGroup")
	}
	mg.Spec.ForProvider.PlacementGroup = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PlacementGroupRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceLinkedRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.ServiceLinkedRoleArnRef,
			Selector:     mg.Spec.ForProvider.ServiceLinkedRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceLinkedRoleArn")
	}
	mg.Spec.ForProvider.ServiceLinkedRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceLinkedRoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCZoneIdentifier),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.VPCZoneIdentifierRefs,
			Selector:      mg.Spec.ForProvider.VPCZoneIdentifierSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCZoneIdentifier")
	}
	mg.Spec.ForProvider.VPCZoneIdentifier = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VPCZoneIdentifierRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("autoscaling.aws.upbound.io", "v1beta2", "LaunchConfiguration", "LaunchConfigurationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LaunchConfiguration),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.LaunchConfigurationRef,
			Selector:     mg.Spec.InitProvider.LaunchConfigurationSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LaunchConfiguration")
	}
	mg.Spec.InitProvider.LaunchConfiguration = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LaunchConfigurationRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.LaunchTemplate != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta2", "LaunchTemplate", "LaunchTemplateList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LaunchTemplate.ID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.LaunchTemplate.IDRef,
				Selector:     mg.Spec.InitProvider.LaunchTemplate.IDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.LaunchTemplate.ID")
		}
		mg.Spec.InitProvider.LaunchTemplate.ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.LaunchTemplate.IDRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.MixedInstancesPolicy != nil {
		if mg.Spec.InitProvider.MixedInstancesPolicy.LaunchTemplate != nil {
			if mg.Spec.InitProvider.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification != nil {
				{
					m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta2", "LaunchTemplate", "LaunchTemplateList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification.LaunchTemplateID),
						Extract:      resource.ExtractResourceID(),
						Reference:    mg.Spec.InitProvider.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification.LaunchTemplateIDRef,
						Selector:     mg.Spec.InitProvider.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification.LaunchTemplateIDSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification.LaunchTemplateID")
				}
				mg.Spec.InitProvider.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification.LaunchTemplateID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification.LaunchTemplateIDRef = rsp.ResolvedReference

			}
		}
	}
	if mg.Spec.InitProvider.MixedInstancesPolicy != nil {
		if mg.Spec.InitProvider.MixedInstancesPolicy.LaunchTemplate != nil {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.MixedInstancesPolicy.LaunchTemplate.Override); i5++ {
				if mg.Spec.InitProvider.MixedInstancesPolicy.LaunchTemplate.Override[i5].LaunchTemplateSpecification != nil {
					{
						m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta2", "LaunchTemplate", "LaunchTemplateList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MixedInstancesPolicy.LaunchTemplate.Override[i5].LaunchTemplateSpecification.LaunchTemplateID),
							Extract:      resource.ExtractResourceID(),
							Reference:    mg.Spec.InitProvider.MixedInstancesPolicy.LaunchTemplate.Override[i5].LaunchTemplateSpecification.LaunchTemplateIDRef,
							Selector:     mg.Spec.InitProvider.MixedInstancesPolicy.LaunchTemplate.Override[i5].LaunchTemplateSpecification.LaunchTemplateIDSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.MixedInstancesPolicy.LaunchTemplate.Override[i5].LaunchTemplateSpecification.LaunchTemplateID")
					}
					mg.Spec.InitProvider.MixedInstancesPolicy.LaunchTemplate.Override[i5].LaunchTemplateSpecification.LaunchTemplateID = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.MixedInstancesPolicy.LaunchTemplate.Override[i5].LaunchTemplateSpecification.LaunchTemplateIDRef = rsp.ResolvedReference

				}
			}
		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "PlacementGroup", "PlacementGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PlacementGroup),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.PlacementGroupRef,
			Selector:     mg.Spec.InitProvider.PlacementGroupSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PlacementGroup")
	}
	mg.Spec.InitProvider.PlacementGroup = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PlacementGroupRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceLinkedRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.ServiceLinkedRoleArnRef,
			Selector:     mg.Spec.InitProvider.ServiceLinkedRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceLinkedRoleArn")
	}
	mg.Spec.InitProvider.ServiceLinkedRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceLinkedRoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.VPCZoneIdentifier),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.VPCZoneIdentifierRefs,
			Selector:      mg.Spec.InitProvider.VPCZoneIdentifierSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCZoneIdentifier")
	}
	mg.Spec.InitProvider.VPCZoneIdentifier = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.VPCZoneIdentifierRefs = mrsp.ResolvedReferences

	return nil
}
