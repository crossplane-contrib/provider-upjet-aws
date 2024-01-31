/*
Copyright 2022 Upbound Inc.
*/

package apis

import (
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IntegrationIDPrefixed returns an extractor that returns the ID
// of an Integration with "integrations/" prefix which is the format
// expected by Route.
func IntegrationIDPrefixed() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		return func(mg metav1.Object) string {
			if meta.GetExternalName(mg) == "" {
				return ""
			}
			return "integrations/" + meta.GetExternalName(mg)
		}(mg)
	}
}
