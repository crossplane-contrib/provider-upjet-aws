// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/pipeline"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/upbound/provider-aws/config"
)

func main() {
	var (
		app                   = kingpin.New("generator", "Run Upjet code generation pipelines for provider-aws").DefaultEnvars()
		repoRoot              = app.Arg("repo-root", "Root directory for the provider repository").Required().String()
		skippedResourcesCSV   = app.Flag("skipped-resources-csv", "File path where a list of skipped (not-generated) Terraform resource names will be stored as a CSV").Envar("SKIPPED_RESOURCES_CSV").String()
		generatedResourceList = app.Flag("generated-resource-list", "File path where a list of the generated resources will be stored.").Envar("GENERATED_RESOURCE_LIST").Default("../config/generated.lst").String()
	)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	absRootDir, err := filepath.Abs(*repoRoot)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", *repoRoot))
	}
	p, err := config.GetProvider(context.Background(), true, false)
	kingpin.FatalIfError(err, "Cannot initialize the provider configuration")
	dumpGeneratedResourceList(p, generatedResourceList)
	dumpSkippedResourcesCSV(p, skippedResourcesCSV)
	pipeline.Run(p, absRootDir)
}

func dumpGeneratedResourceList(p *ujconfig.Provider, targetPath *string) {
	if len(*targetPath) == 0 {
		return
	}
	generatedResources := make([]string, 0, len(p.Resources))
	for name := range p.Resources {
		generatedResources = append(generatedResources, name)
	}
	sort.Strings(generatedResources)
	// Indentation is not necessary, as it's a flat JSON array, but newlines prevent git conflicts from concurrent PRs
	// adding new resources that are not alphabetically adjacent.
	buff, err := json.MarshalIndent(generatedResources, "", "")
	if err != nil {
		panic(fmt.Sprintf("Cannot marshal native schema versions to JSON: %s", err.Error()))
	}
	if err := os.WriteFile(*targetPath, buff, 0o600); err != nil {
		panic(fmt.Sprintf("Cannot write native schema versions of generated resources to file %s: %s", *targetPath, err.Error()))
	}
}

func dumpSkippedResourcesCSV(p *ujconfig.Provider, targetPath *string) {
	if len(*targetPath) == 0 {
		return
	}
	skippedCount := len(p.GetSkippedResourceNames())
	totalCount := skippedCount + len(p.Resources)
	summaryLine := fmt.Sprintf("Available, skipped, total, coverage: %d, %d, %d, %.1f%%", len(p.Resources), skippedCount, totalCount, (float64(len(p.Resources))/float64(totalCount))*100)
	if err := os.WriteFile(*targetPath, []byte(strings.Join(append([]string{summaryLine}, p.GetSkippedResourceNames()...), "\n")), 0o600); err != nil {
		panic(fmt.Sprintf("Cannot write skipped resources CSV to file %s: %s", *targetPath, err.Error()))
	}
}
