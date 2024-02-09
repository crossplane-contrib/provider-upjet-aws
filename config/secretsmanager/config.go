package secretsmanager

import (
	"fmt"
	"strconv"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/resource/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
)

// Configure adds configurations for the secretsmanager group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_secretsmanager_secret", func(r *config.Resource) {
		// Use aws_secretsmanager_secret_rotation.
		config.MoveToStatus(r.TerraformResource, "rotation_rules", "rotation_lambda_arn")
		// aws_secretsmanager_secret_policy.
		config.MoveToStatus(r.TerraformResource, "policy")
		// TODO: we had better do this for all resources...
		r.TerraformConfigurationInjector = func(_ map[string]any, params map[string]any) {
			params["name_prefix"] = ""
		}
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			// skip diff customization on create
			if state == nil || state.Empty() {
				return diff, nil
			}
			if config == nil {
				return nil, errors.New("resource config cannot be nil")
			}
			// skip no diff or destroy diffs
			if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
				return diff, nil
			}

			resData, err := schema.InternalMap(r.TerraformResource.SchemaMap()).Data(state, diff)
			if err != nil {
				return nil, errors.New("could not construct resource data")
			}

			// do not customize diff if replica field has no change
			if !resData.HasChange("replica") {
				return diff, nil
			}
			currentReplicaSet, ok := r.TerraformResource.Data(state).Get("replica").(*schema.Set)
			if !ok {
				return nil, errors.New("could not read \"replica\" from state")
			}

			nisReplica, ok := config.Get("replica")
			if !ok {
				// config is empty for replica, no need for custom diff logic
				// this is already handled correctly with the built-in diff logic
				return diff, nil
			}
			desiredReplicaList := nisReplica.([]interface{})
			// this is the hash implementation of *schema.Set, which is unexported
			hashFunc := func(val interface{}) string {
				code := currentReplicaSet.F(val)
				if code < 0 {
					code = -code
				}
				return strconv.Itoa(code)
			}
			// the number of
			desiredReplicasWithRegionOnly := make(map[string]int)
			//
			// map[region]map[kms_key_id]hash
			regionsInCurrentState := make(map[string]map[string]string)
			type replica struct {
				KMSKeyID string `json:"kms_key_id"`
				Region   string `json:"region"`
			}

			// traverse state
			for _, v := range currentReplicaSet.List() {
				// v is an interface{} type, which is a replica
				// marshal then unmarshal to convert it to the internal type replica for easier access to fields
				replicaBytes, err := json.JSParser.Marshal(v)
				if err != nil {
					return nil, errors.Wrap(err, "cannot serialize replica")
				}
				cReplica := &replica{}
				if err := json.JSParser.Unmarshal(replicaBytes, cReplica); err != nil {
					return nil, err
				}
				if cReplica.Region == "" || cReplica.KMSKeyID == "" {
					// we should not be here, replicas at current state always have their region set and kms_key_id computed
					return nil, errors.New("replica in current state does not have region or kms key id set")
				}

				// add the kms_key_id to the region's kms_key_id map. For convenience, store the hash of (region,kms_key_id) pair
				if replicaHashesByKMSKeysOfRegion, ok := regionsInCurrentState[cReplica.Region]; ok {
					replicaHashesByKMSKeysOfRegion[cReplica.KMSKeyID] = hashFunc(v)
				} else {
					// if we are adding a kms_key_id for the region for the first time, initialize the kms_key_id map
					regionsInCurrentState[cReplica.Region] = map[string]string{
						cReplica.KMSKeyID: hashFunc(v),
					}
				}
			}

			// a convenience function for deleting diff entries for a particular replica entry
			removeReplicaFromDiffViaHash := func(hash string) {
				delete(diff.Attributes, fmt.Sprintf("replica.%s.kms_key_id", hash))
				delete(diff.Attributes, fmt.Sprintf("replica.%s.region", hash))
				delete(diff.Attributes, fmt.Sprintf("replica.%s.status", hash))
				delete(diff.Attributes, fmt.Sprintf("replica.%s.status_message", hash))
				delete(diff.Attributes, fmt.Sprintf("replica.%s.last_accessed_date", hash))
			}
			// traverse the desired Replica list at resource config (params)
			// we want to count the Replicas, that has only region specified (no explicit kms_key_id specified, kms_key_id is left to cloud API to be automatically selected)
			// then record this per region
			for _, v := range desiredReplicaList {
				// v is an interface{} type, which is a replica
				// marshal then unmarshal to convert it to the internal type replica for easier access to fields
				replicaBytes, err := json.JSParser.Marshal(v)
				if err != nil {
					return nil, err
				}
				dReplica := &replica{}
				if err := json.JSParser.Unmarshal(replicaBytes, dReplica); err != nil {
					return nil, err
				}

				// count the region-only replicas (i.e. with automatically assigned KMS Key IDs) at parameters
				if dReplica.KMSKeyID == "" {
					if count, ok := desiredReplicasWithRegionOnly[dReplica.Region]; ok {
						desiredReplicasWithRegionOnly[dReplica.Region] = count + 1
					} else {
						desiredReplicasWithRegionOnly[dReplica.Region] = 1
					}
				} else {
					// this is a Replica at params, with explicit KMS Key ID specified in region
					// check whether we have an exact match in current state
					if replicaHashesByKMSKeyOfRegion, ok := regionsInCurrentState[dReplica.Region]; ok {
						// we have an exact matching region,kms_key_id pair in the current state
						// there should be no diff involved, remove the diff if it got calculated somehow
						removeReplicaFromDiffViaHash(replicaHashesByKMSKeyOfRegion[dReplica.KMSKeyID])
						createdHash := hashFunc(map[string]interface{}{
							"kms_key_id": dReplica.KMSKeyID,
							"region":     dReplica.Region,
						})
						removeReplicaFromDiffViaHash(createdHash)
						delete(replicaHashesByKMSKeyOfRegion, dReplica.KMSKeyID)
					}
				}
			}

			// now try to match the region-only desired KMS Key IDs with the ones in the left in current state (after filtering out the explicit matches above).
			for region, unmatchedDesiredKMSKeyCount := range desiredReplicasWithRegionOnly {
				kmsKeysOfRegionInCurrentState, ok := regionsInCurrentState[region]
				if !ok {
					// this is a Replica with brand-new region, no action needed. it already shows up on diff
					continue
				}

				switch {
				case len(kmsKeysOfRegionInCurrentState) > unmatchedDesiredKMSKeyCount:
					// for the particular region, we have more KMS Key IDs present in the current state than we desire, e.g.
					// current state for region1 = { region1_kmsKeyA, region1_kmsKeyB, region1_kmsKeyC }
					// desired state for region1 = { region1_kmsKeyANY }
					// due to set difference implementation in TF, this will show up as 3 deletions, 1 creation in DIFF
					// instead, in this case, we want to have 2 deletions, 0 creation DIFF
					// thus, remove all (unmatchedDesiredKMSKeyCount=1) creation diffs and remove (unmatchedDesiredKMSKeyCount=1) deletion diff
					// Arbitrarily choose which replica to delete, since they're indistinguishable
					i := 0
					for _, hash := range kmsKeysOfRegionInCurrentState {
						if i >= unmatchedDesiredKMSKeyCount {
							break
						}
						removeReplicaFromDiffViaHash(hash)
						i++
					}
					creationHash := hashFunc(map[string]interface{}{
						"kms_key_id": "",
						"region":     region,
					})
					removeReplicaFromDiffViaHash(creationHash)
				case len(kmsKeysOfRegionInCurrentState) < unmatchedDesiredKMSKeyCount:
					// this might not be possible at all, due to Replica hash function
					for _, hash := range kmsKeysOfRegionInCurrentState {
						removeReplicaFromDiffViaHash(hash)
					}
				default:
					// for the particular region, we have matching number of KMS Key IDs to desired, i.e. there should be no diff for these
					// example
					// current state for region2 KMS Key IDs = { region2_kmsKeyX}
					// desired state for region2 KMS Key IDs = { region2_kmsKeyANY }
					// due to set difference implementation in TF, this will show up as 1 deletion, 1 creation in DIFF
					// instead, in this case, we want to have no diff
					// thus, remove all creation diffs and remove all deletion diffs
					for _, hash := range kmsKeysOfRegionInCurrentState {
						removeReplicaFromDiffViaHash(hash)
					}
					creationHash := hashFunc(map[string]interface{}{
						"kms_key_id": "",
						"region":     region,
					})
					removeReplicaFromDiffViaHash(creationHash)
				}
			}
			// compare the total desired Replica count and current Replica count
			// adjust the diff for replica.#
			if len(desiredReplicaList) == len(currentReplicaSet.List()) {
				// no diff, therefore remove diff if exists
				delete(diff.Attributes, "replica.#")
			} else if replicaCount, ok := diff.Attributes["replica.#"]; ok {
				// there is a diff in unmodified diff, make sure it is correct after modifications
				replicaCount.Old = strconv.Itoa(len(currentReplicaSet.List()))
				replicaCount.New = strconv.Itoa(len(desiredReplicaList))
			} else {
				// there was no diff in unmodified diff, but there is on the customized. Add this diff
				diff.Attributes["replica.#"] = &terraform.ResourceAttrDiff{
					Old: strconv.Itoa(len(currentReplicaSet.List())),
					New: strconv.Itoa(len(desiredReplicaList)),
				}
			}
			return diff, nil
		}
	})
}
