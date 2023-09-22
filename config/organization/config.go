// Copyright 2022 Upbound Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package organization

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for the organization group.
func Configure(p *config.Provider) {
	// please see: https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/organizations_account
	// If `role_name` is used, it's stated that Terraform will always
	// show a difference. Thus, that argument is removed here.
	// TODO: There may be similar other arguments.
	p.AddResourceConfigurator("aws_organizations_account", func(r *config.Resource) {
		delete(r.TerraformResource.Schema, "role_name")
	})
	p.AddResourceConfigurator("aws_organizations_delegated_administrator", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/organizations/v1beta1.Account",
		}
	})
	// We are deleting this reference as we have three different types: Organization Account,
	// Organization Root, Organization Unit.
	p.AddResourceConfigurator("aws_organizations_policy_attachment", func(r *config.Resource) {
		delete(r.References, "target_id")
	})
}
