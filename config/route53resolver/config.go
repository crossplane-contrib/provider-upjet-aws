/*
Copyright 2022 Upbound Inc.
*/

package route53resolver

import (
	"fmt"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/resource/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
	"strconv"
)

// Configure adds configurations for the route53resolver group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_route53_resolver_query_log_config", func(r *config.Resource) {
		delete(r.References, "destination_arn")
	})

	p.AddResourceConfigurator("aws_route53_resolver_endpoint", func(r *config.Resource) {
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

			currentIPAddressSet, ok := r.TerraformResource.Data(state).Get("ip_address").(*schema.Set)
			if !ok {
				return nil, errors.New("could not read \"ip_address\" from state")
			}
			nisIp, ok := config.Get("ip_address")
			if !ok {
				return nil, errors.New("could not read ip_address block from config")
			}
			desiredIPAddressList := nisIp.([]interface{})
			// this is the hash implementation of *schema.Set, which is unexported
			hashFunc := func(val interface{}) string {
				code := currentIPAddressSet.F(val)
				if code < 0 {
					code = -code
				}
				return strconv.Itoa(code)
			}
			// the number of
			desiredIpAddressesWithSubnetOnly := make(map[string]int)
			//
			// map[subnet_id]map[ip]hash
			subnetsInCurrentState := make(map[string]map[string]string)
			type ipAddress struct {
				IP       string `json:"ip"`
				SubnetId string `json:"subnet_id"`
				IpId     string `json:"ip_id"`
			}

			// traverse state
			for _, v := range currentIPAddressSet.List() {
				// v is an interface{} type, which is an ip_address
				// marshal then unmarshal to convert it to the internal type ipAddress for easier access to fields
				ipadBytes, err := json.JSParser.Marshal(v)
				if err != nil {
					continue
				}
				ipAddr := &ipAddress{}
				if err := json.JSParser.Unmarshal(ipadBytes, ipAddr); err != nil {
					return nil, err
				}
				if ipAddr.SubnetId == "" || ipAddr.IP == "" {
					// we should not be here, ipAddresses at current state always have their subnet_id set and ip computed
					continue
				}
				// if _, ok := desiredIpAddressesWithSubnetOnly[ipAddr.SubnetId]; !ok {
				//	desiredIpAddressesWithSubnetOnly[ipAddr.SubnetId] = 0
				// }

				// add the IP to the subnet's IP map. For convenience, store the hash of (subnet_id,ip) pair
				if subnet, ok := subnetsInCurrentState[ipAddr.SubnetId]; ok {
					subnet[ipAddr.IP] = hashFunc(v)
				} else {
					// if we are adding an IP for the subnet for the first time, initialize the IP map
					subnetsInCurrentState[ipAddr.SubnetId] = map[string]string{
						ipAddr.IP: hashFunc(v),
					}
				}
			}

			// a convenience function for deleting diff entries for a particular ipAddress entry
			removeIPAddressFromDiffViaHash := func(hash string) {
				delete(diff.Attributes, fmt.Sprintf("ip_address.%s.subnet_id", hash))
				delete(diff.Attributes, fmt.Sprintf("ip_address.%s.ip_id", hash))
				delete(diff.Attributes, fmt.Sprintf("ip_address.%s.ip", hash))
			}
			// traverse the desired IP address list at resource config (params)
			// we want to count the ip addresses, that has only subnet specified (no explicit ip specified, ip is left to cloud API to be automatically selected)
			// then record this per subnet
			for _, v := range desiredIPAddressList {
				// v is an interface{} type, which is an ip_address
				// marshal then unmarshal to convert it to the internal type ipAddress for easier access to fields
				ipadBytes, err := json.JSParser.Marshal(v)
				if err != nil {
					continue
				}
				ipAddr := &ipAddress{}
				if err := json.JSParser.Unmarshal(ipadBytes, ipAddr); err != nil {
					return nil, err
				}

				// count the subnet-only IP address (automatically assigned IPs) at parameters
				if ipAddr.IP == "" {
					if count, ok := desiredIpAddressesWithSubnetOnly[ipAddr.SubnetId]; ok {
						desiredIpAddressesWithSubnetOnly[ipAddr.SubnetId] = count + 1
					} else {
						desiredIpAddressesWithSubnetOnly[ipAddr.SubnetId] = 1
					}
				} else {
					// this is an IP address at params, with explicit IP specified in subnet
					// check whether we have an exact match in current state
					if subnet, ok := subnetsInCurrentState[ipAddr.SubnetId]; ok {
						// we have an exact matching subnet,IP pair in the current state
						// there should be no diff involved, remove the diff if it got calculated somehow
						removeIPAddressFromDiffViaHash(subnet[ipAddr.IP])
						createdHash := hashFunc(map[string]interface{}{
							"ip":        ipAddr.IP,
							"subnet_id": ipAddr.SubnetId,
						})
						removeIPAddressFromDiffViaHash(createdHash)
						delete(subnet, ipAddr.IP)
					} else {
						// we have a new subnet,IP pair introduced, this will show up in diff
						// no action needed on diff
					}
				}
			}

			// now try to match the subnet-only desired IPs with the ones in the left in current state (after filtering out the explicit matches above).
			for subnetId, unmatchedDesiredIPCount := range desiredIpAddressesWithSubnetOnly {
				ipsOfSubnetInCurrentState, ok := subnetsInCurrentState[subnetId]
				if !ok {
					// this is an ip address with brand-new subnet, no action needed. it already shows up on diff
					continue
				}

				if len(ipsOfSubnetInCurrentState) > unmatchedDesiredIPCount {
					// for the particular subnet, we have more IPs present in the current state than we desire, e.g.
					// current state for subnet1 = { subnet1_ipA, subnet1_ipB, subnet1_ipC }
					// desired state for subnet1 = { subnet1_ipANY }
					// due to set difference implementation in TF, this will show up as 3 deletions, 1 creation in DIFF
					// instead, in this case, we want to have 2 deletions, 0 creation DIFF
					// thus, remove all (unmatchedDesiredIPCount=1) creation diffs and remove (unmatchedDesiredIPCount=1) deletion diff
					// choose IP address to delete randomly
					i := 0
					for _, hash := range ipsOfSubnetInCurrentState {
						if i >= unmatchedDesiredIPCount {
							break
						}
						removeIPAddressFromDiffViaHash(hash)
						i++
					}
					creationHash := hashFunc(map[string]interface{}{
						"ip":        "",
						"ip_id":     "",
						"subnet_id": subnetId,
					})
					removeIPAddressFromDiffViaHash(creationHash)
				} else if len(ipsOfSubnetInCurrentState) < unmatchedDesiredIPCount {
					// this might not be possible at all, due to endpoint hash function
					for _, hash := range ipsOfSubnetInCurrentState {
						removeIPAddressFromDiffViaHash(hash)
					}
				} else {
					// for the particular subnet, we have matching number of IPs to desired, i.e there should be no diff for these
					// example
					// current state for subnet2 IPs = { subnet2_ipX}
					// desired state for subnet2 IPS = { subnet2_ipANY }
					// due to set difference implementation in TF, this will show up as 1 deletion, 1 creation in DIFF
					// instead, in this case, we want to have no diff
					// thus, remove all creation diffs and remove all deletion diffs
					// choose IP address to delete randomly
					for _, hash := range ipsOfSubnetInCurrentState {
						removeIPAddressFromDiffViaHash(hash)
					}
					creationHash := hashFunc(map[string]interface{}{
						"ip":        "",
						"ip_id":     "",
						"subnet_id": subnetId,
					})
					removeIPAddressFromDiffViaHash(creationHash)
				}

			}
			// compare the total desired IP count and current IP count
			// adjust the diff for ipAddress.#
			if len(desiredIPAddressList) == len(currentIPAddressSet.List()) {
				// no diff, therefore remove diff if exists
				delete(diff.Attributes, "ip_address.#")
			} else if ipAddressCount, ok := diff.Attributes["ip_address.#"]; ok {
				// there is a diff in unmodified diff, make sure it is correct after modifications
				ipAddressCount.Old = strconv.Itoa(len(currentIPAddressSet.List()))
				ipAddressCount.New = strconv.Itoa(len(desiredIPAddressList))
			} else {
				// there was no diff in unmodified diff, but there is on the customized. Add this diff
				diff.Attributes["ip_address.#"] = &terraform.ResourceAttrDiff{
					Old: strconv.Itoa(len(currentIPAddressSet.List())),
					New: strconv.Itoa(len(desiredIPAddressList)),
				}
			}
			return diff, nil
		}
	})

}
