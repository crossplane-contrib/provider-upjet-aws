//go:build generate
// +build generate

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// NOTE: See the below link for details on what is happening here.
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

// Remove existing CRDs
//go:generate rm -rf ../package/crds

// Remove generated files
//go:generate bash -c "find ../apis \\( -iname 'zz_generated.conversion_hubs.go' -o -iname 'zz_generated.conversion_spokes.go' \\) -delete"
//go:generate bash -c "find ../apis -type d -empty -delete"
//go:generate bash -c "find ../internal/controller -iname 'zz_*' -delete"
//go:generate bash -c "find ../internal/controller -type d -empty -delete"
//go:generate bash -c "find ../cmd/provider -name 'zz_*' -type f -delete"
//go:generate bash -c "find ../cmd/provider -type d -maxdepth 1 -mindepth 1 -empty -delete"

// Scrape metadata from Terraform registry
//go:generate go run github.com/crossplane/upjet/cmd/scraper -n hashicorp/terraform-provider-aws -r ../.work/terraform-provider-aws/website/docs/r -o ../config/provider-metadata.yaml

// NOTE(muvaf): Some of Terraform AWS provider files have "!generate" build tag
// that prevent us from using it for generator program.

// Run Terrajet generator
//go:generate go run ../cmd/generator/main.go ..

// Generate deepcopy methodsets and CRD manifests
//go:generate go run -tags generate sigs.k8s.io/controller-tools/cmd/controller-gen object:headerFile=../hack/boilerplate.go.txt paths=../apis/cluster/... crd:allowDangerousTypes=true,crdVersions=v1 output:artifacts:config=../package/crds
//go:generate go run -tags generate sigs.k8s.io/controller-tools/cmd/controller-gen object:headerFile=../hack/boilerplate.go.txt paths=../apis/namespaced/... crd:allowDangerousTypes=true,crdVersions=v1 output:artifacts:config=../package/crds

// Generate crossplane-runtime methodsets (resource.Claim, etc)
//go:generate go run -tags generate github.com/crossplane/crossplane-tools/cmd/angryjet generate-methodsets --header-file=../hack/boilerplate.go.txt ../apis/...

// Run upjet's transformer for the generated resolvers to get rid of the cross
// API-group imports and to prevent import cycles
//go:generate go run github.com/crossplane/upjet/cmd/resolver -g aws.upbound.io -a github.com/upbound/provider-aws/internal/apis -s -p ../apis/cluster/...
//go:generate go run github.com/crossplane/upjet/cmd/resolver -g aws.m.upbound.io -a github.com/upbound/provider-aws/internal/apis -s -p ../apis/namespaced/...

package apis

import (
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen" //nolint:typecheck

	_ "github.com/crossplane/crossplane-tools/cmd/angryjet" //nolint:typecheck

	_ "github.com/crossplane/upjet/cmd/scraper"

	_ "github.com/crossplane/upjet/cmd/resolver"
)
