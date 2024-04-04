#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

echo "obtain kubeconfig from ClusterAuth connection secret"
${KUBECTL} -n upbound-system get secret sample-eks-cluster-conn -o go-template='{{ .data.kubeconfig | base64decode }}' > sampleclusterkube
echo "checking kubectl version"
${KUBECTL} --kubeconfig ./sampleclusterkube version
echo "checking cluster-info"
${KUBECTL} --kubeconfig ./sampleclusterkube cluster-info
echo "listing nodes"
${KUBECTL} --kubeconfig ./sampleclusterkube get nodes
echo "listing pods"
${KUBECTL} --kubeconfig ./sampleclusterkube get pods

