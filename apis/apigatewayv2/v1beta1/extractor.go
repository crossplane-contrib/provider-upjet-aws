/*
Copyright 2021 Upbound Inc.
*/

package v1beta1

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
			integration, ok := mg.(*Integration)
			if !ok {
				return ""
			}
			if meta.GetExternalName(integration) == "" {
				return ""
			}
			return "integrations/" + meta.GetExternalName(integration)
		}(mg)
	}
}
