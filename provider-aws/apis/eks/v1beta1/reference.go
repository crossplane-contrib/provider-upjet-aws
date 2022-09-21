package v1beta1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

// TODO(muvaf): Move this to crossplane-runtime.

// ExternalNameIfReady returns the external name only if the targeted resource
// reports itself as ready.
func ExternalNameIfReady() reference.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		if !mr.GetCondition(xpv1.TypeReady).Equal(xpv1.Available()) {
			return ""
		}
		return reference.ExternalName()(mr)
	}
}
