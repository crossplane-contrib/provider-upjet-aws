//go:build !custombootcheck
// +build !custombootcheck

// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package bootcheck

func CheckEnv() error {
	// No-op by default. Use build tags for build-time isolation of custom preflight checks.
	// Ensure to update the build tags on L1-L2 so that they are mutually exclusive across implementations.
	return nil
}
