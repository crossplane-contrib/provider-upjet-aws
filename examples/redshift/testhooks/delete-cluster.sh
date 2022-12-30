#!/usr/bin/env bash
set -aeuo pipefail

# Delete the cluster resource before deleting the resource itself
${KUBECTL} delete cluster.redshift.aws.upbound.io --all