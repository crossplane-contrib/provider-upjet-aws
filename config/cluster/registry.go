// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cluster

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the specified Provider.
type Configure func(provider *config.Provider)

// Configurator is a registry for provider Configs.
type Configurator []Configure

// AddConfig adds a Config to the Configurator registry.
func (c *Configurator) AddConfig(conf Configure) {
	*c = append(*c, conf)
}

// ProviderConfiguration is a global registry to be used by
// the resource providers to register their Config functions.
var ProviderConfiguration = Configurator{}
