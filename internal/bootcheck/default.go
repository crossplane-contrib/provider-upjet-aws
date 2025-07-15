//go:build !custombootcheck
// +build !custombootcheck

package bootcheck

func CheckEnv() error {
	// No-op by default. Use build tags for build-time isolation of custom preflight checks.
	// Ensure to update the build tags on L1-L2 so that they are mutually exclusive across implementations.
	return nil
}
