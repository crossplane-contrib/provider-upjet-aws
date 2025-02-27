//go:build linter_run

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package cluster

import "k8s.io/apimachinery/pkg/runtime"

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	panic(`Must not be called in provider runtime. The provider should not have been built with the "linter_run" build constraint.`)
}
