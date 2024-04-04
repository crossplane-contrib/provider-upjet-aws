#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

echo "obtain kubeconfig from ClusterAuth connection secret"
${KUBECTL} -n upbound-system get secret sample-eks-cluster-conn -o go-template='{{ .data.kubeconfig | base64decode }}' > /tmp/sampleclusterkube
echo "checking kubectl version"
${KUBECTL} --kubeconfig /tmp/sampleclusterkube version
echo "checking cluster-info"
${KUBECTL} --kubeconfig /tmp/sampleclusterkube cluster-info
echo "listing nodes"
${KUBECTL} --kubeconfig /tmp/sampleclusterkube get nodes
echo "listing pods"
${KUBECTL} --kubeconfig /tmp/sampleclusterkube get pods

