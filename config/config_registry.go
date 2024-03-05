/*
Copyright 2024 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// Config configures the specified Provider.
type Config func(provider *config.Provider)

// Configurator is a registry for provider Configs.
type Configurator struct {
	configs []Config
}

// AddConfig adds a Config to the Configurator registry.
func (c *Configurator) AddConfig(conf Config) {
	c.configs = append(c.configs, conf)
}

// ProviderConfiguration is a global registry to be used by
// the resource providers to register their Config functions.
var ProviderConfiguration = &Configurator{}
