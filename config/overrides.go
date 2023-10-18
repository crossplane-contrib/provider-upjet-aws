/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/types"
	"github.com/crossplane/upjet/pkg/types/comments"
	"github.com/crossplane/upjet/pkg/types/name"

	"github.com/upbound/provider-aws/config/common"
)

// RegionAddition adds region to the spec of all resources except iam group which
// does not have a region notion.
func RegionAddition() config.ResourceOption {
	return func(r *config.Resource) {
		if r.ShortGroup == "iam" || r.ShortGroup == "opsworks" {
			return
		}
		c := "Region is the region you'd like your resource to be created in.\n"
		comment, err := comments.New(c, comments.WithTFTag("-"))
		if err != nil {
			panic(errors.Wrap(err, "cannot build comment for region"))
		}

		// check if the underlying Terraform resource already has "region"
		// as a (state) attribute
		if s, ok := r.TerraformResource.Schema["region"]; ok && types.IsObservation(s) {
			r.SchemaElementOptions.SetAddToObservation("region")
		}

		r.TerraformResource.Schema["region"] = &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: comment.String(),
		}
		if r.MetaResource == nil {
			return
		}
		for _, ex := range r.MetaResource.Examples {
			defaultRegion := "us-west-1"
			if err := ex.SetPathValue("region", defaultRegion); err != nil {
				panic(err)
			}
			for k := range ex.Dependencies {
				if strings.HasPrefix(k, "aws_iam") {
					continue
				}
				if err := ex.Dependencies.SetPathValue(k, "region", defaultRegion); err != nil {
					panic(err)
				}
			}
		}
	}
}

// TagsAllRemoval removes the tags_all field that is used only in tfstate to
// accumulate provider-wide default tags in TF, which is not something we support.
// So, we don't need it as a parameter while "tags" is already in place.
func TagsAllRemoval() config.ResourceOption {
	return func(r *config.Resource) {
		if t, ok := r.TerraformResource.Schema["tags_all"]; ok {
			t.Computed = true
			t.Optional = false
		}
	}
}

// IdentifierAssignedByAWS will work for all AWS types because even if the ID
// is assigned by user, we'll see it in the TF State ID.
// The resource-specific configurations should override this whenever possible.
func IdentifierAssignedByAWS() config.ResourceOption {
	return func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	}
}

// NamePrefixRemoval makes sure we remove name_prefix from all since it is mostly
// for Terraform functionality.
func NamePrefixRemoval() config.ResourceOption {
	return func(r *config.Resource) {
		for _, f := range r.ExternalName.OmittedFields {
			if f == "name_prefix" {
				return
			}
		}
		r.ExternalName.OmittedFields = append(r.ExternalName.OmittedFields, "name_prefix")
	}
}

// KnownReferencers adds referencers for fields that are known and common among
// more than a few resources.
func KnownReferencers() config.ResourceOption { //nolint:gocyclo
	return func(r *config.Resource) {
		for k, s := range r.TerraformResource.Schema {
			// We shouldn't add referencers for status fields and sensitive fields
			// since they already have secret referencer.
			if (s.Computed && !s.Optional) || s.Sensitive {
				continue
			}
			switch {
			case strings.HasSuffix(k, "role_arn"):
				r.References[k] = config.Reference{
					Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
					Extractor: common.PathARNExtractor,
				}
			case strings.HasSuffix(k, "security_group_ids"):
				r.References[k] = config.Reference{
					Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
					RefFieldName:      name.NewFromSnake(strings.TrimSuffix(k, "s")).Camel + "Refs",
					SelectorFieldName: name.NewFromSnake(strings.TrimSuffix(k, "s")).Camel + "Selector",
				}
			case r.ShortGroup == "glue" && k == "database_name":
				r.References["database_name"] = config.Reference{
					Type: "github.com/upbound/provider-aws/apis/glue/v1beta1.CatalogDatabase",
				}
			}
			switch k {
			case "vpc_id":
				r.References["vpc_id"] = config.Reference{
					Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC",
				}
			case "subnet_ids":
				r.References["subnet_ids"] = config.Reference{
					Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
					RefFieldName:      "SubnetIDRefs",
					SelectorFieldName: "SubnetIDSelector",
				}
			case "subnet_id":
				r.References["subnet_id"] = config.Reference{
					Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
				}
			case "iam_roles":
				r.References["iam_roles"] = config.Reference{
					Type:              "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
					RefFieldName:      "IAMRoleRefs",
					SelectorFieldName: "IAMRoleSelector",
				}
			case "security_group_id":
				r.References["security_group_id"] = config.Reference{
					Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
				}
			case "kms_key_id":
				r.References["kms_key_id"] = config.Reference{
					Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
				}
			case "kms_key_arn":
				r.References["kms_key_arn"] = config.Reference{
					Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
				}
			case "kms_key":
				r.References["kms_key"] = config.Reference{
					Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
				}
			}
		}
	}
}

// AddExternalTagsField adds ExternalTagsFieldName configuration for resources that have tags field.
func AddExternalTagsField() config.ResourceOption {
	return func(r *config.Resource) {
		if s, ok := r.TerraformResource.Schema["tags"]; ok && s.Type == schema.TypeMap {
			r.InitializerFns = append(r.InitializerFns, config.TagInitializer)
		}
	}
}

// DocumentationForTags overrides the API documentation of the tags fields since
// it contains Terraform-specific feature call out.
func DocumentationForTags() config.ResourceOption {
	return func(r *config.Resource) {
		if r.MetaResource == nil {
			return
		}
		if _, ok := r.MetaResource.ArgumentDocs["tags"]; ok {
			r.MetaResource.ArgumentDocs["tags"] = "- (Optional) Key-value map of resource tags."
		}
	}
}

func noForkClientConfig() config.ResourceOption {
	return func(r *config.Resource) {
		r.UseAsync = false
		r.UseNoForkClient = true
	}
}
