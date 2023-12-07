package identitystore

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the identitystore group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_identitystore_group", func(r *config.Resource) {
		// Display name is required by terraform, and while it's not part of the external name or terraform id, it is
		// how the group is displayed, and it's immutable.
		r.ExternalName.IdentifierFields = append(r.ExternalName.IdentifierFields, "display_name")
	})
}
