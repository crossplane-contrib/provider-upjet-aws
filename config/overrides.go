/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/config"
	"github.com/upbound/upjet/pkg/types/comments"
	"github.com/upbound/upjet/pkg/types/name"

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
func KnownReferencers() config.ResourceOption {
	return func(r *config.Resource) {
		_knownReferencers("", r.TerraformResource, r)
	}
}

func _knownReferencers(prefix string, sr *schema.Resource, cr *config.Resource) { //nolint: gocyclo
	for k, s := range sr.Schema {
		// We shouldn't add referencers for status fields and sensitive fields
		// since they already have secret referencer.
		if (s.Computed && !s.Optional) || s.Sensitive {
			continue
		}
		if sub, ok := s.Elem.(*schema.Resource); ok {
			_knownReferencers(prefix+k+".", sub, cr)
			continue
		}
		switch {
		case strings.HasSuffix(k, "role_arn"):
			cr.References[prefix+k] = config.Reference{
				Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
				Extractor: common.PathARNExtractor,
			}
			fmt.Println(cr.Name, prefix + k, "iam.Role")
		case strings.HasSuffix(k, "security_group_ids"):
			cr.References[prefix+k] = config.Reference{
				Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
				RefFieldName:      name.NewFromSnake(strings.TrimSuffix(k, "s")).Camel + "Refs",
				SelectorFieldName: name.NewFromSnake(strings.TrimSuffix(k, "s")).Camel + "Selector",
			}
			fmt.Println(cr.Name, prefix + k, "ec2.SecurityGroups")
		case cr.ShortGroup == "glue" && k == "database_name":
			cr.References[prefix+"database_name"] = config.Reference{
				Type: "github.com/upbound/provider-aws/apis/glue/v1beta1.CatalogDatabase",
			}
			fmt.Println(cr.Name, prefix + k, "glue.CatalogDatabase")
		case strings.HasSuffix(k, "kms_key_id"):
			cr.References[prefix+k] = config.Reference{
				Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			}
			fmt.Println(cr.Name, prefix + k, "kms.Key")
		case strings.HasSuffix(k, "kms_key_arn"):
			cr.References[prefix+k] = config.Reference{
				Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			}
			fmt.Println(cr.Name, prefix + k, "kms.Key")
		case strings.HasSuffix(k, "kms_key"):
			cr.References[prefix+k] = config.Reference{
				Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			}
			fmt.Println(cr.Name, prefix + k, "kms.Key")
		default:
			switch k {
			case "vpc_id":
				cr.References[prefix+k] = config.Reference{
					Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC",
				}
				fmt.Println(cr.Name, prefix + k, "ec2.VPC")
			case "subnet_ids":
				cr.References[prefix+k] = config.Reference{
					Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
					RefFieldName:      "SubnetIDRefs",
					SelectorFieldName: "SubnetIDSelector",
				}
				fmt.Println(cr.Name, prefix + k, "ec2.Subnets")
			case "subnet_id":
				cr.References[prefix+k] = config.Reference{
					Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
				}
				fmt.Println(cr.Name, prefix + k, "ec2.Subnet")
			case "iam_roles":
				cr.References[prefix+k] = config.Reference{
					Type:              "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
					RefFieldName:      "IAMRoleRefs",
					SelectorFieldName: "IAMRoleSelector",
				}
				fmt.Println(cr.Name, prefix + k, "iam.Role")
			case "security_group_id":
				cr.References[prefix+k] = config.Reference{
					Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
				}
				fmt.Println(cr.Name, prefix + k, "ec2.SecurityGroup")
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
