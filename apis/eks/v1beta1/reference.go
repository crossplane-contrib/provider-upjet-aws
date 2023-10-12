package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"k8s.io/utils/ptr"
)

// ExternalNameIfClusterActive returns the external name only if the EKS cluster
// is in ACTIVE state.
func ExternalNameIfClusterActive() reference.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		cl, ok := mr.(*Cluster)
		if !ok {
			return ""
		}
		if ptr.Deref(cl.Status.AtProvider.Status, "") != "ACTIVE" {
			return ""
		}
		return reference.ExternalName()(mr)
	}
}
