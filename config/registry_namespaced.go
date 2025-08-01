// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"context"
	_ "embed"
	"github.com/upbound/provider-aws/config/namespaced"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/registry/reference"
	"github.com/crossplane/upjet/v2/pkg/schema/traverser"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"github.com/upbound/provider-aws/hack"
)

// GetProviderNamespaced returns the provider configuration.
// The `generationProvider` argument specifies whether the provider
// configuration is being read for the code generation pipelines.
// In that case, we will only use the JSON schema for generating
// the CRDs.
func GetProviderNamespaced(ctx context.Context, fwProvider fwprovider.Provider, sdkProvider *schema.Provider, generationProvider bool, skipDefaultTags bool) (*config.Provider, error) {
	if generationProvider {
		p, err := getProviderSchema(providerSchema)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read the Terraform SDK provider from the JSON schema for code generation")
		}
		if err := traverser.TFResourceSchema(sdkProvider.ResourcesMap).Traverse(traverser.NewMaxItemsSync(p.ResourcesMap)); err != nil {
			return nil, errors.Wrap(err, "cannot sync the MaxItems constraints between the Go schema and the JSON schema")
		}
		// use the JSON schema to temporarily prevent float64->int64
		// conversions in the CRD APIs.
		// We would like to convert to int64s with the next major release of
		// the provider.
		sdkProvider = p
	}

	defaultResourceOptions := []config.ResourceOption{
		GroupKindOverrides(),
		KindOverrides(),
		RegionRequired(),
		TagsAllRemoval(),
		IdentifierAssignedByAWS(),
		KnownReferencers(),
		ResourceConfigurator(),
		NamePrefixRemoval(),
		DocumentationForTags(),
		injectPluginFrameworkCustomStateEmptyCheck(),
	}
	if !skipDefaultTags {
		defaultResourceOptions = append(defaultResourceOptions, AddExternalTagsField())
	}

	modulePath := "github.com/upbound/provider-aws"
	pc := config.NewProvider([]byte(providerSchema), "aws",
		modulePath, providerMetadata,
		config.WithShortName("aws"),
		config.WithRootGroup("aws.m.upbound.io"),
		config.WithIncludeList(CLIReconciledResourceList()),
		config.WithTerraformPluginSDKIncludeList(TerraformPluginSDKResourceList()),
		config.WithTerraformPluginFrameworkIncludeList(TerraformPluginFrameworkResourceList()),
		config.WithReferenceInjectors([]config.ReferenceInjector{reference.NewInjector(modulePath)}),
		config.WithSkipList(skipList),
		config.WithFeaturesPackage("internal/features"),
		config.WithMainTemplate(hack.MainTemplate),
		config.WithTerraformProvider(sdkProvider),
		config.WithTerraformPluginFrameworkProvider(fwProvider),
		config.WithSchemaTraversers(&config.SingletonListEmbedder{}),
		config.WithDefaultResourceOptions(defaultResourceOptions...),
	)
	pc.BasePackages.ControllerMap["eks/clusterauth"] = "eks"

	for _, configure := range namespaced.ProviderConfiguration {
		configure(pc)
	}

	pc.ConfigureResources()
	registerTFSingletonListConversions(pc)
	return pc, nil
}

func registerTFSingletonListConversions(pc *config.Provider) {
	for name, r := range pc.Resources {
		r := r
		// nothing to do if no singleton list has been converted to
		// an embedded object
		if len(r.CRDListConversionPaths()) == 0 {
			continue
		}

		// the resource has at least one singleton list converted, so we need
		// the appropriate Terraform converter in this case.
		r.TerraformConversions = []config.TerraformConversion{
			config.NewTFSingletonConversion(),
		}

		pc.Resources[name] = r
	}
}
