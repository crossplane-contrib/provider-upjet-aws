// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"context"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/config/conversion"
	"github.com/crossplane/upjet/pkg/registry/reference"
	"github.com/crossplane/upjet/pkg/schema/traverser"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/xpprovider"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/upbound/provider-aws/hack"
)

var (
	//go:embed schema.json
	providerSchema string

	//go:embed provider-metadata.yaml
	providerMetadata []byte

	//go:embed field-rename.yaml
	fieldRename []byte

	// oldSingletonListAPIs is a newline-delimited list of Terraform resource
	// names with converted singleton list APIs with at least CRD API version
	// containing the old singleton list API. This is to prevent the API
	// conversion for the newly added resources whose CRD APIs will already
	// use embedded objects instead of the singleton lists and thus, will
	// not possess a CRD API version with the singleton list. Thus, for
	// the newly added resources (resources added after the singleton lists
	// have been converted), we do not need the CRD API conversion
	// functions that convert between singleton lists and embedded objects,
	// but we need only the Terraform conversion functions.
	// This list is immutable and represents the set of resources with the
	// already generated CRD API versions with now converted singleton lists.
	// Because new resources should never have singleton lists in their
	// generated APIs, there should be no need to add them to this list.
	// However, bugs might result in exceptions in the future.
	// Please see:
	// https://github.com/crossplane-contrib/provider-upjet-aws/pull/1332
	// for more context on singleton list to embedded object conversions.
	//go:embed old-singleton-list-apis.txt
	oldSingletonListAPIs string
)

var skipList = []string{
	"aws_waf_rule_group$",              // Too big CRD schema
	"aws_wafregional_rule_group$",      // Too big CRD schema
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

var (
	reAPIVersion = regexp.MustCompile(`^v(\d+)((alpha|beta)(\d+))?$`)
)

const (
	errFmtCannotBumpSingletonList = "cannot bump the API version for the resource %q containing a singleton list in its API"
	errFmtCannotFindPrev          = "cannot compute the previous API versions for the resource %q containing a singleton list in its API"
	errFmtInvalidAPIVersion       = "cannot parse %q as a Kubernetes API version string"
)

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
func GetProvider(ctx context.Context, generationProvider bool, skipDefaultTags bool) (*config.Provider, error) {
	fwProvider, sdkProvider, err := xpprovider.GetProvider(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get the Terraform framework and SDK providers")
	}

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
		RegionAddition(),
		TagsAllRemoval(),
		IdentifierAssignedByAWS(),
		KnownReferencers(),
		ResourceConfigurator(),
		NamePrefixRemoval(),
		DocumentationForTags(),
		injectFieldRenamingConversionFunctions(),
	}
	if !skipDefaultTags {
		defaultResourceOptions = append(defaultResourceOptions, AddExternalTagsField())
	}

	modulePath := "github.com/upbound/provider-aws"
	pc := config.NewProvider([]byte(providerSchema), "aws",
		modulePath, providerMetadata,
		config.WithShortName("aws"),
		config.WithRootGroup("aws.upbound.io"),
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
	pc.BasePackages.ControllerMap["internal/controller/eks/clusterauth"] = "eks"

	for _, configure := range ProviderConfiguration {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, bumpVersionsWithEmbeddedLists(pc)
}

func bumpVersionsWithEmbeddedLists(pc *config.Provider) error {
	l := strings.Split(strings.TrimSpace(oldSingletonListAPIs), "\n")
	oldSLAPIs := make(map[string]struct{}, len(l))
	for _, n := range l {
		oldSLAPIs[n] = struct{}{}
	}

	for name, r := range pc.Resources {
		r := r
		// nothing to do if no singleton list has been converted to
		// an embedded object
		if len(r.CRDListConversionPaths()) == 0 {
			continue
		}

		if _, ok := oldSLAPIs[name]; ok {
			if err := configureSingletonListAPIConverters(r); err != nil {
				return errors.Wrap(err, "failed to configure singleton list API converters")
			}
		} else {
			// the controller will be reconciling on the CRD API version
			// with the converted API (with embedded objects in place of
			// singleton lists), so we need the appropriate Terraform
			// converter in this case.
			r.TerraformConversions = []config.TerraformConversion{
				config.NewTFSingletonConversion(),
			}
		}
		pc.Resources[name] = r
	}
	return nil
}

func configureSingletonListAPIConverters(r *config.Resource) error {
	bumped := r.Version
	currentVer := "v1beta2"
	// Field renamings for these three resources already bump their versions.
	// Please see config.injectFieldRenamingConversionFunctions().
	// We do not bump their versions again here.
	if !sets.New("aws_connect_hours_of_operation", "aws_connect_queue", "aws_db_instance").Has(r.Name) {
		var err error
		bumped, err = bumpAPIVersion(r.Version)
		if err != nil {
			return errors.Wrapf(err, errFmtCannotBumpSingletonList, r.Name)
		}
		currentVer = r.Version
	}

	r.Version = bumped
	if r.PreviousVersions == nil {
		prev, err := getPreviousVersions(bumped)
		if err != nil {
			return errors.Wrapf(err, errFmtCannotFindPrev, r.Name)
		}
		r.PreviousVersions = prev
	}
	// we would like to set the storage version to v1beta1 to facilitate
	// downgrades.
	r.SetCRDStorageVersion(currentVer)
	// because the controller reconciles on the API version with the singleton list API,
	// no need for a Terraform conversion.
	r.ControllerReconcileVersion = currentVer

	// assumes the first element is the identity conversion from
	// the default resource and removes it because we will register another
	// identity converter below.
	r.Conversions = r.Conversions[1:]
	r.Conversions = append([]conversion.Conversion{
		conversion.NewIdentityConversionExpandPaths(conversion.AllVersions, conversion.AllVersions, conversion.DefaultPathPrefixes(), r.CRDListConversionPaths()...),
		conversion.NewSingletonListConversion(conversion.AllVersions, bumped, conversion.DefaultPathPrefixes(), r.CRDListConversionPaths(), conversion.ToEmbeddedObject),
		conversion.NewSingletonListConversion(bumped, conversion.AllVersions, conversion.DefaultPathPrefixes(), r.CRDListConversionPaths(), conversion.ToSingletonList),
	}, r.Conversions...)

	return nil
}

// returns a new API version by bumping the last number if the
// API version string is a Kubernetes API version string such
// as v1alpha1, v1beta1 or v1. Otherwise, returns an error.
// If the specified version is v1beta1, then the bumped version is v1beta2.
// If the specified version is v1, then the bumped version is v2.
func bumpAPIVersion(v string) (string, error) {
	m := reAPIVersion.FindStringSubmatch(v)
	switch {
	// e.g., v1
	case len(m) == 2:
		n, err := strconv.ParseUint(m[1], 10, 0)
		if err != nil {
			return "", errors.Wrapf(err, errFmtInvalidAPIVersion, v)
		}
		return fmt.Sprintf("v%d", n+1), nil

	// e.g., v1beta1
	case len(m) == 5:
		n, err := strconv.ParseUint(m[4], 10, 0)
		if err != nil {
			return "", errors.Wrapf(err, errFmtInvalidAPIVersion, v)
		}
		return fmt.Sprintf("v%s%s%d", m[1], m[3], n+1), nil

	default:
		// then cannot bump this version string
		return "", errors.Errorf(errFmtInvalidAPIVersion, v)
	}
}

func getPreviousVersions(v string) ([]string, error) {
	p := "v1beta1"
	var result []string
	var err error
	for p != v {
		result = append(result, p)
		p, err = bumpAPIVersion(p)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
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

// TerraformPluginSDKResourceList returns the list of resources that have external
// name configured in ExternalNameConfigs table and to be reconciled under
// the no-fork architecture.
func TerraformPluginSDKResourceList() []string {
	l := make([]string, len(TerraformPluginSDKExternalNameConfigs))
	i := 0
	for name := range TerraformPluginSDKExternalNameConfigs {
		// Expected format is regex, and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

func TerraformPluginFrameworkResourceList() []string {
	l := make([]string, len(TerraformPluginFrameworkExternalNameConfigs))
	i := 0
	for name := range TerraformPluginFrameworkExternalNameConfigs {
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
