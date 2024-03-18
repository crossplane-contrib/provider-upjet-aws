// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"context"
	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"
	"reflect"
	"unsafe"

	"github.com/crossplane/upjet/pkg/registry/reference"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/xpprovider"
	"github.com/pkg/errors"

	"github.com/upbound/provider-aws/hack"

	"github.com/crossplane/upjet/pkg/config"
)

var (
	//go:embed schema.json
	providerSchema string

	//go:embed provider-metadata.yaml
	providerMetadata []byte
)

var skipList = []string{
	"aws_waf_rule_group$",              // Too big CRD schema
	"aws_wafregional_rule_group$",      // Too big CRD schema
	"aws_mwaa_environment$",            // See https://github.com/crossplane-contrib/terrajet/issues/100
	"aws_ecs_tag$",                     // tags are already managed by ecs resources.
	"aws_alb$",                         // identical with aws_lb
	"aws_alb_target_group_attachment$", // identical with aws_lb_target_group_attachment
	"aws_iam_policy_attachment$",       // identical with aws_iam_*_policy_attachment resources.
	"aws_iam_group_policy$",            // identical with aws_iam_*_policy_attachment resources.
	"aws_iam_user_policy$",             // identical with aws_iam_*_policy_attachment resources.
	"aws_alb$",                         // identical with aws_lb.
	"aws_alb_listener$",                // identical with aws_lb_listener.
	"aws_alb_target_group$",            // identical with aws_lb_target_group.
	"aws_alb_target_group_attachment$", // identical with aws_lb_target_group_attachment.
	"aws_iot_authorizer$",              // failure with unknown reason.
	"aws_location_map$",                // failure with unknown reason.
	"aws_appflow_connector_profile$",   // failure with unknown reason.
	"aws_rds_reserved_instance",        // Expense of testing
}

// workaround for the TF AWS v4.67.0-based no-fork release: We would like to
// keep the types in the generated CRDs intact
// (prevent number->int type replacements).
func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns the provider configuration.
// The `generationProvider` argument specifies whether the provider
// configuration is being read for the code generation pipelines.
// In that case, we will only use the JSON schema for generating
// the CRDs.
func GetProvider(ctx context.Context, generationProvider bool) (*config.Provider, *xpprovider.AWSClient, error) {
	var p *schema.Provider
	var err error
	if generationProvider {
		p, err = getProviderSchema(providerSchema)
	} else {
		p, err = xpprovider.GetProviderSchema(ctx)
	}
	if err != nil {
		return nil, nil, errors.Wrapf(err, "cannot get the Terraform provider schema with generation mode set to %t", generationProvider)
	}
	// we set schema.Provider's meta to nil because p.Configure modifies
	// a singleton pointer. This further assumes that the
	// schema.Provider.Configure calls do not modify the global state
	// represented by the config.Provider.TerraformProvider.
	var awsClient *xpprovider.AWSClient
	if !generationProvider {
		// #nosec G103
		awsClient = (*xpprovider.AWSClient)(unsafe.Pointer(reflect.ValueOf(p.Meta()).Pointer()))
	}
	p.SetMeta(nil)
	modulePath := "github.com/upbound/provider-aws"
	pc := config.NewProvider([]byte(providerSchema), "aws",
		modulePath, providerMetadata,
		config.WithShortName("aws"),
		config.WithRootGroup("aws.upbound.io"),
		config.WithIncludeList(CLIReconciledResourceList()),
		config.WithTerraformPluginSDKIncludeList(NoForkResourceList()),
		config.WithReferenceInjectors([]config.ReferenceInjector{reference.NewInjector(modulePath)}),
		config.WithSkipList(skipList),
		config.WithFeaturesPackage("internal/features"),
		config.WithMainTemplate(hack.MainTemplate),
		config.WithTerraformProvider(p),
		config.WithDefaultResourceOptions(
			GroupKindOverrides(),
			KindOverrides(),
			RegionAddition(),
			TagsAllRemoval(),
			IdentifierAssignedByAWS(),
			KnownReferencers(),
			AddExternalTagsField(),
			ResourceConfigurator(),
			NamePrefixRemoval(),
			DocumentationForTags(),
		),
	)
	pc.BasePackages.ControllerMap["internal/controller/eks/clusterauth"] = "eks"

	for _, configure := range ProviderConfiguration {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, awsClient, nil
}

// CLIReconciledResourceList returns the list of resources that have external
// name configured in ExternalNameConfigs table and to be reconciled under
// the TF CLI based architecture.
func CLIReconciledResourceList() []string {
	l := make([]string, len(CLIReconciledExternalNameConfigs))
	i := 0
	for name := range CLIReconciledExternalNameConfigs {
		// Expected format is regex, and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

// NoForkResourceList returns the list of resources that have external
// name configured in ExternalNameConfigs table and to be reconciled under
// the no-fork architecture.
func NoForkResourceList() []string {
	l := make([]string, len(NoForkExternalNameConfigs))
	i := 0
	for name := range NoForkExternalNameConfigs {
		// Expected format is regex, and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

// Configure configures the specified Provider.
type Configure func(provider *config.Provider)

// Configurator is a registry for provider Configs.
type Configurator []Configure

// AddConfig adds a Config to the Configurator registry.
func (c *Configurator) AddConfig(conf Configure) {
	*c = append(*c, conf)
}

// ProviderConfiguration is a global registry to be used by
// the resource providers to register their Config functions.
var ProviderConfiguration = Configurator{}
