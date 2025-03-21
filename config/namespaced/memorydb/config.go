// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package memorydb

import (
	"fmt"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the memorydb group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_memorydb_cluster", func(r *config.Resource) {
		r.UseAsync = true

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if clusterendpoints, ok := attr["cluster_endpoint"].([]any); ok {
				for i, cp := range clusterendpoints {
					if clusterendpoint, ok := cp.(map[string]any); ok && len(clusterendpoint) > 0 {
						if address, ok := clusterendpoint["address"].(string); ok {
							key := fmt.Sprintf("cluster_endpoint_%d_address", i)
							conn[key] = []byte(address)
						}
						if port, ok := clusterendpoint["port"].(float64); ok {
							key := fmt.Sprintf("cluster_endpoint_%d_port", i)
							conn[key] = []byte(fmt.Sprintf("%g", port))
						}
					}
				}
			}
			if shards, ok := attr["shards"].([]any); ok {
				for i, shard := range shards {
					if s, ok := shard.(map[string]any); ok {
						if nodes, ok := s["nodes"].([]any); ok {
							for j, node := range nodes {
								if nod, ok := node.(map[string]any); ok {
									if endpoints, ok := nod["endpoint"].([]any); ok && len(endpoints) > 0 {
										for _, endpoint := range endpoints {
											if ep, ok := endpoint.(map[string]any); ok && len(ep) > 0 {
												if address, ok := ep["address"].(string); ok {
													key := fmt.Sprintf("shard_node_%d_%d_endpoint_address", i+1, j+1)
													conn[key] = []byte(address)
												}
												if port, ok := ep["port"].(float64); ok {
													key := fmt.Sprintf("shard_node_%d_%d_endpoint_port", i+1, j+1)
													conn[key] = []byte(fmt.Sprintf("%g", port))
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
			return conn, nil
		}
	})
}
