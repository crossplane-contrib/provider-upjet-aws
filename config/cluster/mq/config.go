// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package mq

import (
	"encoding/base64"
	"fmt"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/crossplane-runtime/v2/pkg/fieldpath"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/registry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Configure adds configurations for the mq group.
func Configure(p *config.Provider) { //nolint:gocyclo
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
		r.LateInitializer.IgnoredFields = append(r.LateInitializer.IgnoredFields, "maintenance_window_start_time")
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

	p.AddResourceConfigurator("aws_mq_configuration", func(r *config.Resource) {
		c := &repeatedDiffCheckerForData{}
		r.UpdateLoopPrevention = c
	})
}

// repeatedDiffCheckerForData implements the UpdateLoopPrevention interface.
// It is responsible for checking if the same diff related to the "data"
// argument (in this case, XML content) appears repeatedly, which may indicate
// a schema violation or invalid data scenario. If a repeated diff is detected,
// it blocks the update process.
type repeatedDiffCheckerForData struct {
	// previousDiff stores the base64-encoded string of the last diff.
	// This is used to compare against the current diff to detect repeated
	// updates.
	previousDiff *string
}

// UpdateLoopPreventionFunc checks for repeated diffs in the resource's "data"
// attribute. If the diff has not changed since the previous reconciliation
// loop, it blocks the update by returning an appropriate result.
func (c *repeatedDiffCheckerForData) UpdateLoopPreventionFunc(diff *terraform.InstanceDiff, mg xpresource.Managed) (*config.UpdateLoopPreventResult, error) { //nolint:gocyclo // easier to follow as a unit
	// Skip processing if there is no diff, the diff is empty, or it is a destroy operation.
	if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
		return nil, nil
	}
	paved, err := fieldpath.PaveObject(mg)
	if err != nil {
		return nil, errors.Wrap(err, "cannot pave object")
	}
	// Retrieve the value of the "spec.forProvider.engineType" field from the paved object.
	// This field is used to determine the engine type of the MQ configuration.
	engineType, err := paved.GetString("spec.forProvider.engineType")
	if err != nil {
		return nil, errors.Wrap(err, "cannot get value of spec.forProvider.engineType")
	}
	// Check if the engine type is "ActiveMQ". If it is not, skip further checks and return nil.
	// This block ensures that the diff check logic only applies to resources with ActiveMQ engine type,
	// avoiding unnecessary diff processing for other engine types.
	if engineType != "ActiveMQ" {
		return nil, nil
	}
	// Encode the "data" attribute of the diff into a base64 string for comparison.
	var encodedDiff string
	if dataDiff, ok := diff.Attributes["data"]; ok {
		// Use GoString to get a string representation of the attribute.
		encodedDiff = base64.StdEncoding.EncodeToString([]byte(dataDiff.GoString()))
	}
	// If there is no previous diff recorded, store the current diff and allow the update to proceed.
	if c.previousDiff == nil {
		c.previousDiff = &encodedDiff
		return nil, nil
	}
	// If the current diff matches the previous diff, block the update and return a reason.
	if encodedDiff == *c.previousDiff {
		return &config.UpdateLoopPreventResult{Reason: "Repeated diff for the provided XML data, please check the XML content you have provided. " +
			"It may contain invalid or schema violating content."}, nil
	}
	// Update the previous diff with the current diff for the next reconciliation loop.
	c.previousDiff = &encodedDiff
	return nil, nil
}
