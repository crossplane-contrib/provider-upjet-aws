/*
Copyright 2021 Upbound Inc.
*/

package common

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/upbound/provider-aws/config/common"

	// PathARNExtractor is the golang path to ARNExtractor function
	// in this package.
	PathARNExtractor = SelfPackagePath + ".ARNExtractor()"

	// VersionV1Alpha2 is used as minimum version for all manually configured
	// resources.
	VersionV1Alpha2 = "v1alpha2"

	// ExternalTagsFieldName is used as field name to set external resource tags.
	ExternalTagsFieldName = "Tags"
)

// ARNExtractor extracts ARN of the resources from "status.atProvider.arn" which
// is quite common among all AWS resources.
func ARNExtractor() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		r, err := paved.GetString("status.atProvider.arn")
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		return r
	}
}
