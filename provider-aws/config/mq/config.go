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
	})
}
