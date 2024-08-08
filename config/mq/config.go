// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package mq

import (
	"fmt"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/registry"
)

// Configure adds configurations for the mq group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_mq_broker", func(r *config.Resource) {
		r.References["security_groups"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SecurityGroupRefs",
			SelectorFieldName: "SecurityGroupSelector",
		}
		r.UseAsync = true
		// TODO(aru): looks like currently angryjet cannot handle references
		//  for non-string struct fields. `configuration.revision` is a
		//  float64 field. Thus here we remove the automatically injected
		//  cross-resource reference from example manifests.
		delete(r.References, "configuration.revision")

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if instances, ok := attr["instances"].([]any); ok {
				for i, inst := range instances {
					if instance, ok := inst.(map[string]any); ok {
						if cu, ok := instance["console_url"].(string); ok {
							key := fmt.Sprintf("instance_%d_console_url", i)
							conn[key] = []byte(cu)
						}
						if ip, ok := instance["ip_address"].(string); ok {
							key := fmt.Sprintf("instance_%d_ip_address", i)
							conn[key] = []byte(ip)
						}
						if endpoints, ok := instance["endpoints"].([]any); ok && len(endpoints) > 0 {
							for j, endpoint := range endpoints {
								key := fmt.Sprintf("instance_%d_endpoint_%d", i, j)
								conn[key] = []byte(endpoint.(string))
							}
						}
					}
				}
			}
			return conn, nil
		}
		// we need to prevent a race between the Broker.mq & User.mq
		// controllers when the users are specified under the spec.forProvider.
		// This configuration will prevent the upjet runtime from
		// late-initializing spec.forProvider.user when the bootstrap users
		// are specified under spec.initProvider. Without this configuration,
		// spec.forProvider gets initialized even if the bootstrap users are
		// specified under spec.initProvider.
		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{
				"user",
			},
		}
	})

	p.AddResourceConfigurator("aws_mq_user", func(r *config.Resource) {
		r.References["broker_id"] = config.Reference{
			TerraformName: "aws_mq_broker",
		}
		r.Version = "v1alpha1"
		r.MetaResource = &registry.Resource{
			ArgumentDocs: make(map[string]string),
		}
		r.MetaResource.ArgumentDocs["console_access"] = `- (Optional) Setting consoleAccess will result in an update loop till the MQ Broker to which this user belongs is restarted.`
	})
}
