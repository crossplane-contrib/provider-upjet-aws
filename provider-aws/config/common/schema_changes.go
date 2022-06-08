/*
Copyright 2022 Upbound Inc.
*/

package common

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// MutuallyExclusiveFields is used for fields that are represented as standalone
// CRDs already. It moves those fields to status so that they are managed only
// in those other CR instances.
func MutuallyExclusiveFields(s *schema.Resource, fields ...string) {
	for _, f := range fields {
		if _, ok := s.Schema[f]; !ok {
			panic(fmt.Sprintf("field %s does not exist in schema", f))
		}
		s.Schema[f].Optional = false
		s.Schema[f].Computed = true
	}
}
