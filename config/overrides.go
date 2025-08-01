// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.in/yaml.v3"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/config/conversion"
	"github.com/crossplane/upjet/pkg/types/name"

	"github.com/upbound/provider-aws/config/common"
)

// RegionRequired makes the region field required for resources that have it in their schema.
func RegionRequired() config.ResourceOption { //nolint:gocyclo
	return func(r *config.Resource) {
		// Only process resources that already have a region field
		if s, ok := r.TerraformResource.Schema["region"]; ok {
			// Make existing region field required
			s.Required = true
			s.Optional = false
			s.Computed = false
			s.Description = "Region is the region you'd like your resource to be created in."

			// Add region to generated examples only if schema has region field
			if r.MetaResource != nil {
				for _, ex := range r.MetaResource.Examples {
					defaultRegion := "us-west-1"
					if err := ex.SetPathValue("region", defaultRegion); err != nil {
						panic(err)
					}
					for k := range ex.Dependencies {
						if err := ex.Dependencies.SetPathValue(k, "region", defaultRegion); err != nil {
							panic(err)
						}
					}
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
					TerraformName: "aws_iam_role",
					Extractor:     common.PathARNExtractor,
				}
			case strings.HasSuffix(k, "security_group_ids"):
				r.References[k] = config.Reference{
					TerraformName:     "aws_security_group",
					RefFieldName:      name.NewFromSnake(strings.TrimSuffix(k, "s")).Camel + "Refs",
					SelectorFieldName: name.NewFromSnake(strings.TrimSuffix(k, "s")).Camel + "Selector",
				}
			case r.ShortGroup == "glue" && k == "database_name":
				r.References["database_name"] = config.Reference{
					TerraformName: "aws_glue_catalog_database",
				}
			}
			switch k {
			case "vpc_id":
				r.References["vpc_id"] = config.Reference{
					TerraformName: "aws_vpc",
				}
			case "subnet_ids":
				r.References["subnet_ids"] = config.Reference{
					TerraformName:     "aws_subnet",
					RefFieldName:      "SubnetIDRefs",
					SelectorFieldName: "SubnetIDSelector",
				}
			case "subnet_id":
				r.References["subnet_id"] = config.Reference{
					TerraformName: "aws_subnet",
				}
			case "iam_roles":
				r.References["iam_roles"] = config.Reference{
					TerraformName:     "aws_iam_role",
					RefFieldName:      "IAMRoleRefs",
					SelectorFieldName: "IAMRoleSelector",
				}
			case "security_group_id":
				r.References["security_group_id"] = config.Reference{
					TerraformName: "aws_security_group",
				}
			case "kms_key_id":
				r.References["kms_key_id"] = config.Reference{
					TerraformName: "aws_kms_key",
				}
			case "kms_key_arn":
				r.References["kms_key_arn"] = config.Reference{
					TerraformName: "aws_kms_key",
				}
			case "kms_key":
				r.References["kms_key"] = config.Reference{
					TerraformName: "aws_kms_key",
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

func injectFieldRenamingConversionFunctions() config.ResourceOption {
	type fieldRenameConversionData struct {
		SourceVersion string                       `yaml:"sourceVersion"`
		TargetVersion string                       `yaml:"targetVersion"`
		Data          map[string]map[string]string `yaml:"data"`
	}
	var data []fieldRenameConversionData
	if err := yaml.Unmarshal(fieldRename, &data); err != nil {
		panic(err)
	}
	return func(r *config.Resource) {
		for _, cd := range data {
			if d, ok := cd.Data[r.Name]; ok {
				for s, t := range d {
					r.Version = cd.TargetVersion
					r.Conversions = append(r.Conversions,
						conversion.NewFieldRenameConversion(cd.SourceVersion, s, cd.TargetVersion, t),
						conversion.NewFieldRenameConversion(cd.TargetVersion, t, cd.SourceVersion, s),
					)
				}
			}
		}
	}
}

func injectPluginFrameworkCustomStateEmptyCheck() config.ResourceOption {
	return func(r *config.Resource) {
		if r.TerraformPluginFrameworkResource != nil {
			r.TerraformPluginFrameworkIsStateEmptyFn = tfStateVlueIsEmpty
		}
	}
}

func tfStateVlueIsEmpty(ctx context.Context, tfStateValue tftypes.Value, resourceSchema rschema.Schema) (bool, error) { //nolint:gocyclo // easier to follow as a unit
	sdkState := tfsdk.State{
		Raw:    tfStateValue.Copy(),
		Schema: resourceSchema,
	}
	var isEmpty bool
	var region string
	sdkState.GetAttribute(ctx, path.Root("region"), &region)
	if region != "" {
		sdkState.SetAttribute(ctx, path.Root("region"), (*string)(nil))
		isEmpty = true
		if err := tftypes.Walk(sdkState.Raw, func(attributePath *tftypes.AttributePath, value tftypes.Value) (bool, error) {
			if len(attributePath.Steps()) != 1 {
				return true, nil
			}

			valType := value.Type()
			switch {
			case valType.Is(tftypes.Set{}), valType.Is(tftypes.List{}), valType.Is(tftypes.Tuple{}):
				if value.IsKnown() && !value.IsNull() {
					destVal := make([]tftypes.Value, 0)
					if err := value.As(&destVal); err != nil {
						return true, err
					}
					if len(destVal) > 0 {
						isEmpty = false
						return false, nil
					}
				}
			default:
				if value.IsKnown() && !value.IsNull() {
					isEmpty = false
					return false, nil
				}
			}
			return true, nil
		}); err != nil {
			return false, err
		}
	}
	return isEmpty, nil
}
