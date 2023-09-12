package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

// Lovingly ripped off from config/common/ARNExtractor. Is there a better way?
func GetConfigurationRevision() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		r, err := paved.GetString("status.atProvider.revision")
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		return r
	}
}
