// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package elasticache

import (
	"strings"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/registry"
)

func TestConfigureUserGroup(t *testing.T) {
	r := &config.Resource{
		Name:       "aws_elasticache_user_group",
		References: map[string]config.Reference{},
		MetaResource: &registry.Resource{
			ArgumentDocs: map[string]string{},
		},
	}
	configureUserGroup(r)

	if !containsString(r.LateInitializer.IgnoredFields, "user_ids") {
		t.Fatalf("expected LateInitializer.IgnoredFields to contain user_ids, got %v", r.LateInitializer.IgnoredFields)
	}
	ref, ok := r.References["user_ids"]
	if !ok || ref.TerraformName != "aws_elasticache_user" {
		t.Fatalf("expected user_ids reference to aws_elasticache_user, got %#v", ref)
	}
	if r.MetaResource.ArgumentDocs["user_ids"] != userGroupUserIdsConflictDocs {
		t.Fatalf("unexpected user_ids ArgumentDocs: %q", r.MetaResource.ArgumentDocs["user_ids"])
	}
}

func TestConfigureUserGroupAssociation(t *testing.T) {
	r := &config.Resource{
		Name:       "aws_elasticache_user_group_association",
		References: map[string]config.Reference{},
		MetaResource: &registry.Resource{
			Description: "Associate an ElastiCache user and user group.",
		},
	}
	configureUserGroupAssociation(r)

	if got := r.References["user_group_id"].TerraformName; got != "aws_elasticache_user_group" {
		t.Fatalf("user_group_id ref = %q, want aws_elasticache_user_group", got)
	}
	if got := r.References["user_id"].TerraformName; got != "aws_elasticache_user" {
		t.Fatalf("user_id ref = %q, want aws_elasticache_user", got)
	}
	if !strings.Contains(r.MetaResource.Description, "conflicts with UserGroup") {
		t.Fatalf("Description missing conflict warning: %q", r.MetaResource.Description)
	}
}

func containsString(list []string, want string) bool {
	for _, s := range list {
		if s == want {
			return true
		}
	}
	return false
}
