/*
Copyright 2022 Upbound Inc.
*/

package mq

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for rds group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_mq_broker", func(r *config.Resource) {
		r.UseAsync = true
		// TODO(aru): looks like currently angryjet cannot handle references
		//  for non-string struct fields. `configuration.revision` is a
		//  float64 field. Thus here we remove the automatically injected
		//  cross-resource reference from example manifests.
		delete(r.References, "configuration.revision")
	})
}
